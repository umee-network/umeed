package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/umee-network/umee/v2/x/oracle/types"
)

// GetTxCmd returns the CLI transaction commands for the x/oracle module.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Transaction commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdDelegateFeedConsent(),
	)

	return cmd
}

// GetCmdDelegateFeedConsent creates a Cobra command to generate or
// broadcast a transaction with a MsgDelegateFeedConsent message.
func GetCmdDelegateFeedConsent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegate-feed-consent [operator] [feeder]",
		Args:  cobra.ExactArgs(2),
		Short: "Delegate oracle feed consent from an operator to another feeder address",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Set(flags.FlagFrom, args[0]); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			feederAddr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgDelegateFeedConsent(sdk.ValAddress(clientCtx.GetFromAddress()), feederAddr)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdAggregateExchangeRatePrevote creates a Cobra command to generate or
// broadcast a transaction with a MsgAggregateExchangeRatePrevote message.
func GetCmdAggregateExchangeRatePrevote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exchange-rate-prevote [hash] [feeder]",
		Args:  cobra.ExactArgs(2),
		Short: "Submit an exchange rate prevote msg.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Set(flags.FlagFrom, args[0]); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			feederAddr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgAggregateExchangeRatePrevote(
				[]byte(args[0]),
				feederAddr,
				sdk.ValAddress(clientCtx.GetFromAddress()),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdAggregateExchangeRateVote creates a Cobra command to generate or
// broadcast a transaction with a NewMsgAggregateExchangeRateVote message.
func GetCmdAggregateExchangeRateVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exchange-rate-prevote [salt] [exchangeRates] [feeder]",
		Args:  cobra.ExactArgs(2),
		Short: "Submit an exchange rate vote msg.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Set(flags.FlagFrom, args[0]); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			feederAddr, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgAggregateExchangeRateVote(
				args[0],
				args[1],
				feederAddr,
				sdk.ValAddress(clientCtx.GetFromAddress()),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
