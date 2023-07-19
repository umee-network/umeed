package grpc

import (
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	gtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	proposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"

	"github.com/umee-network/umee/v5/client"
	"github.com/umee-network/umee/v5/x/uibc"
)

var govDeposit sdk.Coins

func init() {
	var err error
	govDeposit, err = sdk.ParseCoinsNormalized("10000000uumee")
	if err != nil {
		panic(err)
	}
}

func SubmitAndPassProposal(umee client.Client, changes []proposal.ParamChange) error {
	resp, err := umee.Tx.GovSubmitProposal(changes, govDeposit)
	if err != nil {
		return err
	}

	// retry
	for i := 0; i < 5; i++ {
		newResp, err := umee.QueryTxHash(resp.TxHash)
		if err != nil && i == 4 {
			return err
		}
		if err == nil {
			resp = newResp
			break
		}

		time.Sleep(time.Second * (1 + time.Duration(i)))
	}

	return MakeVoteAndCheckProposal(umee, *resp)
}

func OracleParamChanges(
	historicStampPeriod uint64,
	maximumPriceStamps uint64,
	medianStampPeriod uint64,
) []proposal.ParamChange {
	return []proposal.ParamChange{
		{
			Subspace: "oracle",
			Key:      "HistoricStampPeriod",
			Value:    fmt.Sprintf("\"%d\"", historicStampPeriod),
		},
		{
			Subspace: "oracle",
			Key:      "MaximumPriceStamps",
			Value:    fmt.Sprintf("\"%d\"", maximumPriceStamps),
		},
		{
			Subspace: "oracle",
			Key:      "MedianStampPeriod",
			Value:    fmt.Sprintf("\"%d\"", medianStampPeriod),
		},
	}
}

func UIBCIBCTransferSatusUpdate(umeeClient client.Client, status uibc.IBCTransferStatus) error {
	msg := uibc.MsgGovSetIBCStatus{
		Authority:   authtypes.NewModuleAddress(gtypes.ModuleName).String(),
		Title:       "Update the ibc transfer status",
		Description: "Update the ibc transfer status",
		IbcStatus:   status,
	}

	resp, err := umeeClient.Tx.TxSubmitProposalWithMsg([]sdk.Msg{&msg})
	if err != nil {
		return err
	}

	// retry
	for i := 0; i < 5; i++ {
		newResp, err := umeeClient.QueryTxHash(resp.TxHash)
		if err != nil && i == 4 {
			return err
		}
		if err == nil {
			resp = newResp
			break
		}

		time.Sleep(time.Second * (1 + time.Duration(i)))
	}

	if len(resp.Events) == 0 {
		return fmt.Errorf("no events in response")
	}

	return MakeVoteAndCheckProposal(umeeClient, *resp)
}

func MakeVoteAndCheckProposal(umeeClient client.Client, resp sdk.TxResponse) error {
	var proposalID string
	for _, event := range resp.Events {
		if event.Type == "submit_proposal" {
			for _, attribute := range event.Attributes {
				if attribute.Key == "proposal_id" {
					proposalID = attribute.Value
				}
			}
		}
	}

	if proposalID == "" {
		return fmt.Errorf("failed to parse proposalID from %s", resp)
	}

	proposalIDInt, err := strconv.ParseUint(proposalID, 10, 64)
	if err != nil {
		return err
	}

	err = umeeClient.Tx.TxGovVoteYesAll(proposalIDInt)
	if err != nil {
		return err
	}

	prop, err := umeeClient.GovProposal(proposalIDInt)
	if err != nil {
		return err
	}

	now := time.Now()
	sleepDuration := prop.VotingEndTime.Sub(now) + 1*time.Second
	fmt.Printf("sleeping %s until end of voting period + 1 block\n", sleepDuration)
	time.Sleep(sleepDuration)

	var propStatus string
	for retry := 0; retry < 5; retry++ {
		prop, err = umeeClient.GovProposal(proposalIDInt)
		if err != nil {
			return err
		}

		propStatus = prop.Status.String()
		if propStatus == "PROPOSAL_STATUS_PASSED" {
			return nil
		}
		time.Sleep(time.Second * (1 + time.Duration(retry)))
	}

	return fmt.Errorf("proposal %d failed to pass with status: %s", proposalIDInt, propStatus)
}
