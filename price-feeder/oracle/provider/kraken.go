package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
	"github.com/umee-network/umee/price-feeder/oracle/types"
)

const (
	krakenHost                    = "ws.kraken.com"
	krakenEventSystemStatus       = "systemStatus"
	krakenEventSubscriptionStatus = "subscriptionStatus"
)

var _ Provider = (*KrakenProvider)(nil)

type (
	// KrakenProvider defines an Oracle provider implemented by the Kraken public
	// API.
	//
	// REF: https://docs.kraken.com/websockets/#overview
	KrakenProvider struct {
		wsURL           url.URL
		wsClient        *websocket.Conn
		logger          zerolog.Logger
		mtx             sync.Mutex
		tickers         map[string]TickerPrice        // Symbol => TickerPrice
		candles         map[string]KrakenCandle       // Symbol => KrakenCandle
		subscribedPairs map[string]types.CurrencyPair // Symbol => types.CurrencyPair
	}

	// KrakenTicker ticker price response from Kraken ticker channel.
	// https://docs.kraken.com/websockets/#message-ticker
	KrakenTicker struct {
		C []string `json:"c"` // Close with Price in the first position
		V []string `json:"v"` // Volume with the value over last 24 hours in the second position
	}

	KrakenCandle struct {
		Open   string
		High   string
		Low    string
		Close  string
		Volume string
	}

	// KrakenSubscriptionMsg Msg to subscribe to all the pairs at once.
	KrakenSubscriptionMsg struct {
		Event        string                    `json:"event"`        // subscribe/unsubscribe
		Pair         []string                  `json:"pair"`         // Array of currency pairs ex.: "BTC/USDT",
		Subscription KrakenSubscriptionChannel `json:"subscription"` // subscription object
	}

	// KrakenSubscriptionChannel Msg with the channel name to be subscribed.
	KrakenSubscriptionChannel struct {
		Name string `json:"name"` // channel to be subscribed ex.: ticker
	}

	// KrakenEvent wraps the possible events from the provider.
	KrakenEvent struct {
		Event string `json:"event"` // events from kraken ex.: systemStatus | subscriptionStatus
	}

	// KrakenEventSystemStatus parse the systemStatus event message.
	KrakenEventSystemStatus struct {
		Status string `json:"status"` // online|maintenance|cancel_only|limit_only|post_only
	}

	// KrakenEventSubscriptionStatus parse the subscriptionStatus event message.
	KrakenEventSubscriptionStatus struct {
		Status       string `json:"status"`       // subscribed|unsubscribed|error
		Pair         string `json:"pair"`         // Pair symbol base/quote ex.: "XBT/USD"
		ErrorMessage string `json:"errorMessage"` // error description
	}
)

// NewKrakenProvider returns a new Kraken provider with the WS connection and msg handler.
func NewKrakenProvider(ctx context.Context, logger zerolog.Logger, pairs ...types.CurrencyPair) (*KrakenProvider, error) {
	wsURL := url.URL{
		Scheme: "wss",
		Host:   krakenHost,
	}

	wsConn, _, err := websocket.DefaultDialer.Dial(wsURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error connecting to Kraken websocket: %w", err)
	}

	provider := &KrakenProvider{
		wsURL:           wsURL,
		wsClient:        wsConn,
		logger:          logger.With().Str("module", "oracle").Logger(),
		tickers:         map[string]TickerPrice{},
		candles:         map[string]KrakenCandle{},
		subscribedPairs: map[string]types.CurrencyPair{},
	}

	if err := provider.SubscribeTickers(pairs...); err != nil {
		return nil, err
	}
	if err := provider.SubscribeCandles(pairs...); err != nil {
		return nil, err
	}

	go provider.handleWebSocketMsgs(ctx)

	return provider, nil
}

// GetTickerPrices returns the tickerPrices based on the saved map.
func (p *KrakenProvider) GetTickerPrices(pairs ...types.CurrencyPair) (map[string]TickerPrice, error) {
	tickerPrices := make(map[string]TickerPrice, len(pairs))

	for _, cp := range pairs {
		key := cp.String()
		tickerPrice, ok := p.tickers[key]
		if !ok {
			return nil, fmt.Errorf("kraken provider failed to get ticker price for %s", key)
		}
		tickerPrices[key] = tickerPrice
	}

	return tickerPrices, nil
}

// SubscribeTickers subscribe to all currency pairs and
// add the new ones into the provider subscribed pairs.
func (p *KrakenProvider) SubscribeTickers(cps ...types.CurrencyPair) error {
	pairs := make([]string, len(cps))

	for i, cp := range cps {
		pairs[i] = currencyPairToKrakenPair(cp)
	}

	if err := p.subscribePairs(pairs...); err != nil {
		return err
	}

	p.setSubscribedPairs(cps...)
	return nil
}

func (p *KrakenProvider) SubscribeCandles(cps ...types.CurrencyPair) error {
	candles := make([]string, len(cps))

	for i, cp := range cps {
		candles[i] = currencyPairToKrakenPair(cp)
	}

	if err := p.subscribeCandlePairs(candles...); err != nil {
		return err
	}

	p.setSubscribedPairs(cps...)
	return nil
}

// handleWebSocketMsgs receive all the messages from the provider
// and controls to reconnect the web socket.
func (p *KrakenProvider) handleWebSocketMsgs(ctx context.Context) {
	reconnectTicker := time.NewTicker(defaultMaxConnectionTime)
	defer reconnectTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(defaultReadNewWSMessage):
			messageType, bz, err := p.wsClient.ReadMessage()
			if err != nil {
				// if some error occurs continue to try to read the next message
				p.logger.Err(err).Msg("kraken provider could not read message")
				if err := p.ping(); err != nil {
					p.logger.Err(err).Msg("failed to send ping")
					p.keepReconnecting()
				}
				continue
			}

			if len(bz) == 0 {
				continue
			}

			p.messageReceived(messageType, bz)

		case <-reconnectTicker.C:
			if err := p.reconnect(); err != nil {
				p.logger.Err(err).Msg("kraken provider attempted to reconnect")
				p.keepReconnecting()
			}
		}
	}
}

// messageReceived handles any message sent by the provider.
func (p *KrakenProvider) messageReceived(messageType int, bz []byte) {
	if messageType != websocket.TextMessage {
		return
	}

	var krakenEvent KrakenEvent
	if err := json.Unmarshal(bz, &krakenEvent); err != nil {
		p.logger.Debug().Msg("kraken provider received a message that is not an event")
		// msg is not an event, it will try to marshal to ticker message
		if err = p.messageReceivedTickerPrice(bz); err != nil {
			if err = p.messageReceivedCandle(bz); err != nil {
				p.logger.Debug().Err(err).Msg("kraken provider was unable to unmarshall")
			}
		}
		return
	}

	switch krakenEvent.Event {
	case krakenEventSystemStatus:
		p.messageReceivedSystemStatus(bz)
		return
	case krakenEventSubscriptionStatus:
		p.messageReceivedSubscriptionStatus(bz)
		return
	}
}

// messageReceivedTickerPrice handles the ticker price msg.
func (p *KrakenProvider) messageReceivedTickerPrice(bz []byte) error {
	// the provider response is an array with different types at each index
	// kraken documentation https://docs.kraken.com/websockets/#message-ticker
	var tickerMessage []interface{}
	if err := json.Unmarshal(bz, &tickerMessage); err != nil {
		return err
	}

	if len(tickerMessage) != 4 {
		return fmt.Errorf("Kraken provider sent something different than ticker")
	}

	channelName, ok := tickerMessage[2].(string)
	if !ok || channelName != "ticker" {
		return fmt.Errorf("Kraken provider sent an unexpected channel name")
	}

	tickerBz, err := json.Marshal(tickerMessage[1])
	if err != nil {
		return err
	}

	var krakenTicker KrakenTicker
	if err := json.Unmarshal(tickerBz, &krakenTicker); err != nil {
		return err
	}

	krakenPair, ok := tickerMessage[3].(string)
	if !ok {
		return fmt.Errorf("Kraken provider sent an unexpected pair")
	}

	krakenPair = normalizeKrakenBTCPair(krakenPair)
	currencyPairSymbol := krakenPairToCurrencyPairSymbol(krakenPair)

	tickerPrice, err := krakenTicker.toTickerPrice(currencyPairSymbol)
	if err != nil {
		p.logger.Err(err).Msg("Kraken provider could not parse kraken ticker to ticker price")
		return err
	}

	p.setTickerPair(currencyPairSymbol, tickerPrice)
	return nil
}

func (kc *KrakenCandle) UnmarshalJSON(buf []byte) error {
	var tmp []interface{}
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if len(tmp) != 9 {
		return fmt.Errorf("wrong number of fields in candle")
	}

	open, ok := tmp[2].(string)
	if !ok {
		return fmt.Errorf("open field must be a string")
	}
	kc.Open = open

	high, ok := tmp[3].(string)
	if !ok {
		return fmt.Errorf("high must be a string")
	}
	kc.High = high

	low, ok := tmp[4].(string)
	if !ok {
		return fmt.Errorf("low field must be a string")
	}
	kc.Low = low

	close, ok := tmp[5].(string)
	if !ok {
		return fmt.Errorf("close field must be a string")
	}
	kc.Close = close

	volume, ok := tmp[7].(string)
	if !ok {
		return fmt.Errorf("volume field must be a string")
	}
	kc.Volume = volume

	return nil
}

// messageReceivedTickerPrice handles the ticker price msg.
func (p *KrakenProvider) messageReceivedCandle(bz []byte) error {
	// the provider response is an array with different types at each index
	// kraken documentation https://docs.kraken.com/websockets/#message-ticker
	var candleMessage []interface{}
	if err := json.Unmarshal(bz, &candleMessage); err != nil {
		return err
	}

	if len(candleMessage) != 4 {
		return fmt.Errorf("Kraken provider sent something different than candle")
	}

	channelName, ok := candleMessage[2].(string)
	if !ok || channelName != "ohlc-1" {
		return fmt.Errorf("Kraken provider sent an unexpected channel name")
	}

	tickerBz, err := json.Marshal(candleMessage[1])
	if err != nil {
		return fmt.Errorf("Kraken provider could not marshal ticker message")
	}

	var krakenCandle KrakenCandle
	if err = krakenCandle.UnmarshalJSON(tickerBz); err != nil {
		return err
	}

	krakenPair, ok := candleMessage[3].(string)
	if !ok {
		return fmt.Errorf("Kraken provider sent an unexpected pair")
	}

	krakenPair = normalizeKrakenBTCPair(krakenPair)
	currencyPairSymbol := krakenPairToCurrencyPairSymbol(krakenPair)

	p.setCandlePair(currencyPairSymbol, krakenCandle)
	return nil
}

// reconnect closes the last WS connection and create a new one.
func (p *KrakenProvider) reconnect() error {
	p.wsClient.Close()

	wsConn, _, err := websocket.DefaultDialer.Dial(p.wsURL.String(), nil)
	if err != nil {
		return fmt.Errorf("error connecting to Kraken websocket: %w", err)
	}
	p.wsClient = wsConn

	pairs := make([]string, len(p.subscribedPairs))
	iterator := 0
	for _, cp := range p.subscribedPairs {
		pairs[iterator] = currencyPairToKrakenPair(cp)
		iterator++
	}

	return p.subscribePairs(pairs...)
}

// keepReconnecting keeps trying to reconnect if an error occurs in recconnect.
func (p *KrakenProvider) keepReconnecting() {
	reconnectTicker := time.NewTicker(defaultReconnectTime)
	defer reconnectTicker.Stop()
	connectionTries := 1

	for time := range reconnectTicker.C {
		if err := p.reconnect(); err != nil {
			p.logger.Err(err).Msgf("kraken provider attempted to reconnect %d times at %s", connectionTries, time.String())
			continue
		}

		if connectionTries > maxReconnectionTries {
			p.logger.Warn().Msgf("kraken provider failed to reconnect %d times", connectionTries)
		}
		connectionTries++
		return
	}
}

// messageReceivedSubscriptionStatus handle the subscription status message
// sent by the provider.
func (p *KrakenProvider) messageReceivedSubscriptionStatus(bz []byte) {
	var subscriptionStatus KrakenEventSubscriptionStatus
	if err := json.Unmarshal(bz, &subscriptionStatus); err != nil {
		p.logger.Err(err).Msg("Kraken provider could not unmarshal KrakenEventSubscriptionStatus")
		return
	}

	switch subscriptionStatus.Status {
	case "error":
		p.logger.Error().Msg(subscriptionStatus.ErrorMessage)
		p.removeSubscribedTickers(krakenPairToCurrencyPairSymbol(subscriptionStatus.Pair))
		return
	case "unsubscribed":
		p.logger.Debug().Msgf("ticker %s was unsubscribed", subscriptionStatus.Pair)
		p.removeSubscribedTickers(krakenPairToCurrencyPairSymbol(subscriptionStatus.Pair))
		return
	}
}

// messageReceivedSystemStatus handle the system status and try to reconnect if it is not online.
func (p *KrakenProvider) messageReceivedSystemStatus(bz []byte) {
	var systemStatus KrakenEventSystemStatus
	if err := json.Unmarshal(bz, &systemStatus); err != nil {
		p.logger.Err(err).Msg("Kraken provider could not unmarshal KrakenEventSystemStatus")
		return
	}

	if strings.EqualFold(systemStatus.Status, "online") {
		return
	}

	p.keepReconnecting()
}

// setTickerPair sets an ticker to the map thread safe by the mutex.
func (p *KrakenProvider) setTickerPair(symbol string, ticker TickerPrice) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	p.tickers[symbol] = ticker
}

func (p *KrakenProvider) setCandlePair(symbol string, candle KrakenCandle) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	p.candles[symbol] = candle
}

// ping to check websocket connection.
func (p *KrakenProvider) ping() error {
	return p.wsClient.WriteMessage(websocket.PingMessage, ping)
}

// subscribePairs write the subscription msg to the provider.
func (p *KrakenProvider) subscribePairs(pairs ...string) error {
	subsMsg := newKrakenSubscriptionMsg(pairs...)
	return p.wsClient.WriteJSON(subsMsg)
}

// subscribeCandlePairs write the subscription msg to the provider.
func (p *KrakenProvider) subscribeCandlePairs(pairs ...string) error {
	subsMsg := newKrakenCandleSubscriptionMsg(pairs...)
	return p.wsClient.WriteJSON(subsMsg)
}

// setSubscribedPairs sets N currency pairs to the map of subscribed pairs.
func (p *KrakenProvider) setSubscribedPairs(cps ...types.CurrencyPair) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	for _, cp := range cps {
		p.subscribedPairs[cp.String()] = cp
	}
}

// removeSubscribedTickers delete N pairs from the subscribed map.
func (p *KrakenProvider) removeSubscribedTickers(tickerSymbols ...string) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	for _, tickerSymbol := range tickerSymbols {
		delete(p.subscribedPairs, tickerSymbol)
	}
}

// toTickerPrice return a TickerPrice based on the KrakenTicker.
func (ticker KrakenTicker) toTickerPrice(symbol string) (TickerPrice, error) {
	if len(ticker.C) != 2 || len(ticker.V) != 2 {
		return TickerPrice{}, fmt.Errorf("error converting KrakenTicker to TickerPrice")
	}
	// ticker.C has the Price in the first position
	// ticker.V has the totla	Value over last 24 hours in the second position
	return newTickerPrice("Kraken", symbol, ticker.C[0], ticker.V[1])
}

// newKrakenSubscriptionMsg returns a new subscription Msg.
func newKrakenSubscriptionMsg(pairs ...string) KrakenSubscriptionMsg {
	return KrakenSubscriptionMsg{
		Event: "subscribe",
		Pair:  pairs,
		Subscription: KrakenSubscriptionChannel{
			Name: "ticker",
		},
	}
}

// newKrakenSubscriptionMsg returns a new subscription Msg.
func newKrakenCandleSubscriptionMsg(pairs ...string) KrakenSubscriptionMsg {
	return KrakenSubscriptionMsg{
		Event: "subscribe",
		Pair:  pairs,
		Subscription: KrakenSubscriptionChannel{
			Name: "ohlc",
		},
	}
}

// krakenPairToCurrencyPairSymbol receives a kraken pair formated
// ex.: ATOM/USDT and return currencyPair Symbol ATOMUSDT.
func krakenPairToCurrencyPairSymbol(krakenPair string) string {
	return strings.Replace(krakenPair, "/", "", -1)
}

// currencyPairToKrakenPair receives a currency pair
// and return kraken ticker symbol ATOM/USDT.
func currencyPairToKrakenPair(cp types.CurrencyPair) string {
	return strings.ToUpper(cp.Base + "/" + cp.Quote)
}

// normalizeKrakenBTCPair changes XBT pairs to BTC,
// since other providers list bitcoin as BTC
func normalizeKrakenBTCPair(ticker string) string {
	return strings.Replace(ticker, "XBT", "BTC", 1)
}
