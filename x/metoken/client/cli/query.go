package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/umee-network/umee/v5/util/cli"
	"github.com/umee-network/umee/v5/x/metoken"
)

// GetQueryCmd returns the CLI query commands for the x/metoken module.
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        metoken.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", metoken.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdQueryParams(),
		GetCmdIndexes(),
		GetCmdIndexBalances(),
		GetCmdSwapFee(),
		GetCmdRedeemFee(),
		GetCmdIndexPrice(),
	)

	return cmd
}

// GetCmdQueryParams creates a Cobra command to query for the x/metoken module parameters.
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the metoken module parameters",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := metoken.NewQueryClient(clientCtx)
			resp, err := queryClient.Params(cmd.Context(), &metoken.QueryParams{})
			return cli.PrintOrErr(resp, err, clientCtx)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdIndexes creates a Cobra command to query for the x/metoken module registered Indexes.
// metoken_denom is optional, if it isn't provided then all the indexes will be returned.
func GetCmdIndexes() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "indexes [metoken_denom]",
		Args: cobra.MaximumNArgs(1),
		Short: "Get all the registered indexes in the x/metoken module or search for a specific index with" +
			" metoken_denom.",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := metoken.NewQueryClient(clientCtx)
			queryReq := metoken.QueryIndexes{}
			if len(args) > 0 {
				queryReq.MetokenDenom = args[0]
			}

			resp, err := queryClient.Indexes(cmd.Context(), &queryReq)
			return cli.PrintOrErr(resp, err, clientCtx)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdSwapFee creates a Cobra command to query for the SwapFee
// Both arguments are required:
// coin: the coin that is taken as base for the fee calculation.
func GetCmdSwapFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap-fee [coin] [metoken_denom]",
		Args:  cobra.ExactArgs(2),
		Short: "Get the fee amount to be charged for a swap. Both args are required.",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := metoken.NewQueryClient(clientCtx)
			queryReq := metoken.QuerySwapFee{}

			asset, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}
			queryReq.Asset = asset
			queryReq.MetokenDenom = args[1]

			resp, err := queryClient.SwapFee(cmd.Context(), &queryReq)
			return cli.PrintOrErr(resp, err, clientCtx)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdRedeemFee creates a Cobra command to query for the RedeemFee
// Both arguments are required:
// metoken: the meToken coin that is taken as base for the fee calculation.
func GetCmdRedeemFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "redeem-fee [metoken] [asset_denom]",
		Args:  cobra.ExactArgs(2),
		Short: "Get the fee amount to be charged for a redeem. Both args are required.",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := metoken.NewQueryClient(clientCtx)
			queryReq := metoken.QueryRedeemFee{}

			metoken, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}
			queryReq.Metoken = metoken
			queryReq.AssetDenom = args[1]

			resp, err := queryClient.RedeemFee(cmd.Context(), &queryReq)
			return cli.PrintOrErr(resp, err, clientCtx)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdIndexBalances creates a Cobra command to query for the x/metoken module Indexes assets balances
// metoken_denom is optional, if it isn't provided then all the balances will be returned.
func GetCmdIndexBalances() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "index-balance [metoken_denom]",
		Args: cobra.MaximumNArgs(1),
		Short: "Get all the indexes' balances in the x/metoken module or search for a specific index with" +
			" metoken_denom.",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := metoken.NewQueryClient(clientCtx)
			queryReq := metoken.QueryIndexBalances{}
			if len(args) > 0 {
				queryReq.MetokenDenom = args[0]
			}
			resp, err := queryClient.IndexBalances(cmd.Context(), &queryReq)
			return cli.PrintOrErr(resp, err, clientCtx)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdIndexPrice creates a Cobra command to query for the x/metoken module Index Price.
// metoken_denom is optional, if it isn't provided then prices for all the registered indexes will be returned.
func GetCmdIndexPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "index-price [metoken_denom]",
		Args: cobra.MaximumNArgs(1),
		Short: "Get price for all the registered indexes in the x/metoken module or search for a specific price with" +
			" metoken_denom.",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := metoken.NewQueryClient(clientCtx)
			queryReq := metoken.QueryIndexPrice{}
			if len(args) > 0 {
				queryReq.MetokenDenom = args[0]
			}
			resp, err := queryClient.IndexPrice(cmd.Context(), &queryReq)
			return cli.PrintOrErr(resp, err, clientCtx)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
