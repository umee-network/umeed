package oracle

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"math"
	"sync"
	"time"

	rpcClient "github.com/cosmos/cosmos-sdk/client/rpc"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umee-network/umee/price-feeder/oracle/client"
	"github.com/umee-network/umee/price-feeder/oracle/provider"
	pfsync "github.com/umee-network/umee/price-feeder/pkg/sync"
	umeetypes "github.com/umee-network/umee/x/oracle/types"
	"google.golang.org/grpc"
)

type Oracle struct {
	logger          zerolog.Logger
	closer          *pfsync.Closer
	mtx             sync.RWMutex
	lastPriceSyncTS time.Time
	prices          map[string]sdk.Dec
	oracleClient    *client.OracleClient
}

func New(oc *client.OracleClient) *Oracle {
	return &Oracle{
		logger:       log.With().Str("module", "oracle").Logger(),
		closer:       pfsync.NewCloser(),
		oracleClient: oc,
	}
}

func (o *Oracle) Stop() {
	o.closer.Close()
	<-o.closer.Done()
}

var denoms = []string{"ATOMUSDT"}

// Should return a prices map that we can set
// The oracle to. Needs to touch multiple providers,
// And then average out.
func GetPrices() map[string]sdk.Dec {

	binanceProvider := provider.NewBinanceProvider()
	krakenProvider := provider.NewKrakenProvider()

	wg := new(sync.WaitGroup)

	wg.Add(1)

	binancePrices, err := binanceProvider.GetTickerPrices(denoms...)

	if binancePrices == nil || err != nil {
		panic("Unable to get binance prices")
	}

	wg.Done()

	wg.Add(1)

	krackenPrices, bigErr := krakenProvider.GetTickerPrices(denoms...)

	if krackenPrices == nil || bigErr != nil {
		panic("Unable to get kracken prices")
	}

	wg.Done()

	wg.Wait()

	// Create new map for averages

	var averages = make(map[string]sdk.Dec, len(denoms))

	half := sdk.MustNewDecFromStr("0.50")

	for _, v := range denoms {
		averages[v] = binancePrices[v]
		averages[v].Add(krackenPrices[v])
		averages[v].Mul(half)
	}

	return averages
}

func (o *Oracle) GetParams() (*umeetypes.QueryParamsResponse, error) {

	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		o.oracleClient.GRPCEndpoint, // your gRPC server address.
		grpc.WithInsecure(),         // The Cosmos SDK doesn't support any transport security mechanism.
	)

	if err != nil {
		panic(err)
	}

	defer grpcConn.Close()
	queryClient := umeetypes.NewQueryClient(grpcConn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	queryResponse, err := queryClient.Params(ctx, &umeetypes.QueryParamsRequest{})

	if err != nil {
		return nil, err
	} else if queryResponse == nil {
		return nil, err
	}
	return queryResponse, nil
}

func (o *Oracle) generateSalt(length int) (string, error) {
	if length == 0 {
		panic("Cannot generate empty salt")
	}
	n := length / 2
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

var previousVotePeriod float64

type PreviousPrevote struct {
	ExchangeRates     string
	Salt              string
	SubmitBlockHeight int64
}

func NewPreviousPrevote() *PreviousPrevote {
	return &PreviousPrevote{
		Salt:              "",
		ExchangeRates:     "",
		SubmitBlockHeight: 0,
	}
}

var previousPrevote *PreviousPrevote = nil

func (o *Oracle) tick() {

	ctx, err := o.oracleClient.CreateContext()

	if err != nil {
		panic(err)
	}

	pricesArray := GetPrices()

	if pricesArray == nil {
		panic("Unable to get prices")
	}

	o.prices = pricesArray

	oracleParams, err := o.GetParams()

	if err != nil || oracleParams == nil {
		panic(err)
	}

	blockHeight, err := rpcClient.GetChainHeight(*ctx)

	if err != nil || blockHeight == 0 {
		panic(err)
	}

	// Get oracle vote period, next block height,
	// Current vote period, index in vote period

	oracleVotePeriod := int64(oracleParams.Params.VotePeriod)
	nextBlockHeight := blockHeight + 1
	currentVotePeriod := math.Floor(float64(nextBlockHeight) / float64(oracleVotePeriod))
	indexInVotePeriod := nextBlockHeight % oracleVotePeriod
	// Skip until new voting period
	// Skip when index [0, oracleVotePeriod - 1] is bigger than oracleVotePeriod - 2 or index is 0
	if (previousVotePeriod != 0 && currentVotePeriod == previousVotePeriod) ||
		oracleVotePeriod-indexInVotePeriod < 2 {
		return
	}

	// If we're past the voting period we needed to hit,
	// Reset and submit another pre-vote
	if previousVotePeriod != 0 && currentVotePeriod-previousVotePeriod != 1 {
		// Reset
		previousVotePeriod = 0
		previousPrevote = nil
		return
	}

	isPrevoteOnlyTx := previousPrevote == nil

	exchangeRates := ""

	for k, v := range o.prices {
		if exchangeRates != "" {
			exchangeRates = exchangeRates + "," + v.String() + k
		}
		if exchangeRates == "" {
			exchangeRates = v.String() + k
		}
	}

	salt, err := o.generateSalt(2)

	if err != nil {
		panic(err)
	}

	valAddr, err := sdk.ValAddressFromBech32(o.oracleClient.ValidatorAddrString)

	if err != nil {
		panic(err)
	}

	hash := umeetypes.GetAggregateVoteHash(salt, exchangeRates, o.oracleClient.ValidatorAddr)

	msg := &umeetypes.MsgAggregateExchangeRatePrevote{
		Hash:      hash.String(), // Hash of prices from the oracle
		Feeder:    o.oracleClient.OracleAddrString,
		Validator: valAddr.String(), //Hash accepts the actual addr
	}

	if isPrevoteOnlyTx {
		// Broadcast message

		err := o.oracleClient.BroadcastPrevote(msg)
		if err != nil {
			panic(err)
		}

		if err != nil {
			panic(err)
		}

		currentHeight, err := rpcClient.GetChainHeight(*ctx)

		if err != nil {
			panic(err)
		}

		previousVotePeriod = math.Floor(float64(currentHeight) / float64(oracleVotePeriod))
		previousPrevote = NewPreviousPrevote()
		previousPrevote.Salt = salt
		previousPrevote.ExchangeRates = exchangeRates
		previousPrevote.SubmitBlockHeight = int64(currentHeight)
	}

	// Is next voting period

	if !isPrevoteOnlyTx {

		// Vote

		voteMsg := &umeetypes.MsgAggregateExchangeRateVote{
			Salt:          previousPrevote.Salt,
			ExchangeRates: previousPrevote.ExchangeRates,
			Feeder:        o.oracleClient.OracleAddrString,
			Validator:     valAddr.String(),
		}

		// Broadcast message

		err := o.oracleClient.BroadcastVote(nextBlockHeight,
			oracleVotePeriod-indexInVotePeriod,
			voteMsg)

		if err != nil {
			// This can happen if the voting is off-timed,
			// Or the voting denoms are not currently on whitelist.
			// We want to just reset and handle this silently :
			previousPrevote = nil
			previousVotePeriod = 0
			return
		}

		previousPrevote = nil

	}

}

func (o *Oracle) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			o.closer.Close()
			return

		default:
			o.tick()
			o.lastPriceSyncTS = time.Now()
			time.Sleep(10 * time.Millisecond)

		}
	}
}

// GetLastPriceSyncTimestamp returns the latest timestamp at which prices where
// fetched from the oracle's set of exchange rate providers.
func (o *Oracle) GetLastPriceSyncTimestamp() time.Time {
	o.mtx.RLock()
	defer o.mtx.RUnlock()

	return o.lastPriceSyncTS
}

// GetPrices returns a copy of the current prices fetched from the oracle's
// set of exchange rate providers.
func (o *Oracle) GetPrices() map[string]sdk.Dec {
	o.mtx.RLock()
	defer o.mtx.RUnlock()

	// Creates a new array for the prices in the oracle
	prices := make(map[string]sdk.Dec, len(o.prices))
	for k, v := range o.prices {
		// Fills in the prices with each value in the oracle
		prices[k] = v
	}

	return prices
}
