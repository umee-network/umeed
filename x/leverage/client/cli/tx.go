package cli

import (
	"fmt"
	"strings"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/umee-network/umee/v6/x/leverage/types"
)

// GetTxCmd returns the CLI transaction commands for the x/leverage module.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Transaction commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		Supply(),
		Withdraw(),
		MaxWithdraw(),
		Collateralize(),
		Decollateralize(),
		Borrow(),
		MaxBorrow(),
		Repay(),
		Liquidate(),
		LeveragedLiquidate(),
		SupplyCollateral(),
	)

	return cmd
}

// Supply creates a Cobra command to generate or broadcast a
// transaction with a MsgSupply message.
func Supply() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "supply [amount]",
		Args:  cobra.ExactArgs(1),
		Short: "Supply a specified amount of a supported asset",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			asset, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgSupply(clientCtx.GetFromAddress(), asset)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// Withdraw creates a Cobra command to generate or broadcast a
// transaction with a MsgWithdraw message.
func Withdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw [amount]",
		Args:  cobra.ExactArgs(1),
		Short: "Withdraw a specified amount of a supplied asset",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			asset, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdraw(clientCtx.GetFromAddress(), asset)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MaxWithdraw creates a Cobra command to generate or broadcast a
// transaction with a MsgMaxWithdraw message.
func MaxWithdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-max [denom]",
		Args:  cobra.ExactArgs(1),
		Short: "Withdraw the maximum valid amount of a supplied asset",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			denom := args[0]

			msg := types.NewMsgMaxWithdraw(clientCtx.GetFromAddress(), denom)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// Collateralize creates a Cobra command to generate or broadcast a
// transaction with a MsgCollateralize message.
func Collateralize() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collateralize [coin]",
		Args:  cobra.ExactArgs(1),
		Short: "Add uTokens as collateral",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coin, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}
			msg := types.NewMsgCollateralize(clientCtx.GetFromAddress(), coin)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// Decollateralize returns a CLI command handler to generate or broadcast a
// transaction with a MsgDecollateralize message.
func Decollateralize() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decollateralize [coin]",
		Args:  cobra.ExactArgs(1),
		Short: "Remove uTokens from collateral",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coin, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}
			msg := types.NewMsgDecollateralize(clientCtx.GetFromAddress(), coin)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// Borrow creates a Cobra command to generate or broadcast a
// transaction with a MsgBorrow message.
func Borrow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "borrow [amount]",
		Args:  cobra.ExactArgs(1),
		Short: "Borrow a specified amount of a supported asset",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			asset, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgBorrow(clientCtx.GetFromAddress(), asset)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// MaxBorrow creates a Cobra command to generate or broadcast a
// transaction with a MsgBorrow message.
func MaxBorrow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "max-borrow [denom]",
		Args:  cobra.ExactArgs(1),
		Short: "Borrow the maximum acceptable amount of a supported asset",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMaxBorrow(clientCtx.GetFromAddress(), args[0])

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// Repay creates a Cobra command to generate or broadcast a
// transaction with a MsgRepay message.
func Repay() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repay [amount]",
		Args:  cobra.ExactArgs(1),
		Short: "Repay a specified amount of a borrowed asset",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			asset, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgRepay(clientCtx.GetFromAddress(), asset)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// Liquidate creates a Cobra command to generate or broadcast a
// transaction with a MsgLiquidate message.
func Liquidate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "liquidate [borrower] [amount] [reward-denom]",
		Args:  cobra.ExactArgs(3),
		Short: "Liquidate a specified amount of a borrower's debt for a chosen reward denomination",
		Long: strings.TrimSpace(
			fmt.Sprintf(`
Liquidate up to a specified amount of a borrower's debt for a chosen reward denomination.

Example:
$ umeed tx leverage liquidate %s  50000000uumee u/uumee --from mykey`,
				"umee1qqy7cst5qm83ldupph2dcq0wypprkfpc9l3jg2",
			),
		),

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			borrowerAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			asset, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			rewardDenom := args[2]

			msg := types.NewMsgLiquidate(clientCtx.GetFromAddress(), borrowerAddr, asset, rewardDenom)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// LeveragedLiquidate creates a Cobra command to generate or broadcast a
// transaction with a MsgLeveragedLiquidate message.
func LeveragedLiquidate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lev-liquidate [borrower] [repay-denom] [reward-denom] [max-repay]",
		Args:  cobra.RangeArgs(1, 3),
		Short: "Liquidates by moving borrower debt to the liquidator and immediately collateralizes the reward.",
		Long: strings.TrimSpace(
			fmt.Sprintf(`
Borrow tokens to liquidate a borrower's debt and immediately collateralize the reward.

Will attempt to repay the max amount allowed by the targeted borrower's position unless max repay (USD) is specified.

The transaction will fail if the liquidator, with new borrow and collateral positions, would be above 0.8 borrow limit.

Example:
$ umeed tx leverage lev-liquidate %s uumee uumee 123.4 --from mykey`,
				"umee1qqy7cst5qm83ldupph2dcq0wypprkfpc9l3jg2",
			),
		),

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			borrowerAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			var repayDenom, rewardDenom string
			maxRepay := sdkmath.LegacyZeroDec()
			if len(args) > 1 {
				repayDenom = args[1]
			}

			if len(args) > 2 {
				rewardDenom = args[2]
			}

			if len(args) > 3 {
				maxRepay = sdkmath.LegacyMustNewDecFromStr(args[3])
			}

			msg := types.NewMsgLeveragedLiquidate(clientCtx.GetFromAddress(), borrowerAddr, repayDenom, rewardDenom, maxRepay)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// SupplyCollateral creates a Cobra command to generate or broadcast a
// transaction with a MsgSupply message.
func SupplyCollateral() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "supply-collateral [amount]",
		Args:  cobra.ExactArgs(1),
		Short: "Supply and collateralize a specified amount of a supported asset",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			asset, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgSupplyCollateral(clientCtx.GetFromAddress(), asset)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
