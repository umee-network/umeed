package client

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

func (c Client) GovClient() govtypes.QueryClient {
	return govtypes.NewQueryClient(c.Query.GrpcConn)
}

func (c Client) GovProposal(proposalID uint64) (*govtypes.Proposal, error) {
	ctx, cancel := c.NewQCtx()
	defer cancel()

	queryResponse, err := c.GovClient().Proposal(ctx, &govtypes.QueryProposalRequest{ProposalId: proposalID})
	if err != nil {
		return nil, err
	}
	return queryResponse.Proposal, nil
}
