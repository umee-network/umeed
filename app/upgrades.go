package app

import (
	"github.com/cometbft/cometbft/libs/log"
	ica "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts"
	icagenesis "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/genesis/types"
	icahosttypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"

	"github.com/cosmos/cosmos-sdk/baseapp"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/group"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/umee-network/umee/v6/app/upgradev3x3"
	"github.com/umee-network/umee/v6/util"
	"github.com/umee-network/umee/v6/x/incentive"
	leveragekeeper "github.com/umee-network/umee/v6/x/leverage/keeper"
	leveragetypes "github.com/umee-network/umee/v6/x/leverage/types"
	"github.com/umee-network/umee/v6/x/metoken"

	oraclemigrator "github.com/umee-network/umee/v6/x/oracle/migrations"
	oracletypes "github.com/umee-network/umee/v6/x/oracle/types"
	"github.com/umee-network/umee/v6/x/ugov"
	"github.com/umee-network/umee/v6/x/uibc"
)

// RegisterUpgradeHandlersregisters upgrade handlers.
func (app UmeeApp) RegisterUpgradeHandlers() {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	app.registerUpgrade3_0(upgradeInfo)
	app.registerUpgrade("v3.1.0", upgradeInfo)
	app.registerUpgrade3_1to3_3(upgradeInfo)
	app.registerUpgrade3_2to3_3(upgradeInfo)
	app.registerUpgrade3_3to4_0(upgradeInfo)
	app.registerUpgrade("v4.0.1", upgradeInfo)
	app.registerUpgrade4_1(upgradeInfo)
	app.registerUpgrade("v4.2", upgradeInfo, uibc.ModuleName)
	app.registerUpgrade4_3(upgradeInfo)
	app.registerUpgrade("v4.4", upgradeInfo)
	app.registerUpgrade("v5.0", upgradeInfo, ugov.ModuleName, wasmtypes.ModuleName)
	app.registerUpgrade5_1(upgradeInfo)
	app.registerUpgrade("v5.2", upgradeInfo) // v5.2 migration is not compatible with v6, so leaving default here.
	app.registerUpgrade6(upgradeInfo)
	app.registerUpgrade6_1("v6.1", upgradeInfo)
	app.registerUpgrade6_2(upgradeInfo)
}

func (app *UmeeApp) registerUpgrade6_2(upgradeInfo upgradetypes.Plan) {
	planName := "v6.2"

	// Set param key table for params module migration
	for _, subspace := range app.ParamsKeeper.GetSubspaces() {
		subspace := subspace
		found := true
		var keyTable paramstypes.KeyTable
		switch subspace.Name() {
		case authtypes.ModuleName:
			keyTable = authtypes.ParamKeyTable() //nolint: staticcheck // deprecated but required for upgrade
		case banktypes.ModuleName:
			keyTable = banktypes.ParamKeyTable() //nolint: staticcheck // deprecated but required for upgrade
		case stakingtypes.ModuleName:
			keyTable = stakingtypes.ParamKeyTable()
		case minttypes.ModuleName:
			keyTable = minttypes.ParamKeyTable() //nolint: staticcheck // deprecated but required for upgrade
		case distrtypes.ModuleName:
			keyTable = distrtypes.ParamKeyTable() //nolint: staticcheck // deprecated but required for upgrade
		case slashingtypes.ModuleName:
			keyTable = slashingtypes.ParamKeyTable() //nolint: staticcheck // deprecated but required for upgrade
		case govtypes.ModuleName:
			keyTable = govv1.ParamKeyTable() //nolint: staticcheck // deprecated but required for upgrade
		case crisistypes.ModuleName:
			keyTable = crisistypes.ParamKeyTable() //nolint: staticcheck // deprecated but required for upgrade
		case wasmtypes.ModuleName:
			keyTable = wasmtypes.ParamKeyTable() //nolint: staticcheck // deprecated but required for upgrade
		default:
			// subspace not handled
			found = false
		}
		if found && !subspace.HasKeyTable() {
			subspace.WithKeyTable(keyTable)
		}
	}

	baseAppLegacySS := app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable())

	app.UpgradeKeeper.SetUpgradeHandler(planName,
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			// Migrate CometBFT consensus parameters from x/params module to a dedicated x/consensus module.
			baseapp.MigrateParams(ctx, baseAppLegacySS, &app.ConsensusParamsKeeper)

			// explicitly update the IBC 02-client params, adding the localhost client type
			params := app.IBCKeeper.ClientKeeper.GetParams(ctx)
			params.AllowedClients = append(params.AllowedClients, ibcexported.Localhost)
			app.IBCKeeper.ClientKeeper.SetParams(ctx, params)

			fromVM, err := app.mm.RunMigrations(ctx, app.configurator, fromVM)
			if err != nil {
				return fromVM, err
			}
			govParams := app.GovKeeper.GetParams(ctx)
			govParams.MinInitialDepositRatio = sdk.NewDecWithPrec(1, 1).String()
			err = app.GovKeeper.SetParams(ctx, govParams)
			return fromVM, err
		},
	)

	app.storeUpgrade(planName, upgradeInfo, storetypes.StoreUpgrades{
		Added: []string{
			consensustypes.ModuleName,
			crisistypes.ModuleName,
		},
	})
	// app.registerNewTokenEmissionUpgrade(upgradeInfo)
}

func (app *UmeeApp) registerUpgrade6_1(planName string, upgradeInfo upgradetypes.Plan) {
	app.UpgradeKeeper.SetUpgradeHandler(planName,
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			printPlanName(planName, ctx.Logger())
			err := app.OracleKeeper.SetHistoricAvgCounterParams(ctx, oracletypes.DefaultAvgCounterParams())
			if err != nil {
				return fromVM, err
			}

			oracleUpgrader := oraclemigrator.NewMigrator(&app.OracleKeeper)
			oracleUpgrader.ExgRatesWithTimestamp(ctx)

			// reference: (today) avg block size in Ethereum is 170kb
			// Cosmwasm binaries can be few hundreds KB up to 3MB
			// our blocks with oracle txs as JSON are about 80kB, so protobuf should be less than 40kB.
			var newMaxBytes int64 = 4_000_000 // 4 MB
			p := app.GetConsensusParams(ctx)
			ctx.Logger().Info("Changing consensus params", "prev", p.Block.MaxBytes, "new", newMaxBytes)
			p.Block.MaxBytes = newMaxBytes
			app.StoreConsensusParams(ctx, p)

			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	app.storeUpgrade(planName, upgradeInfo, storetypes.StoreUpgrades{
		Added: []string{metoken.ModuleName},
	})
}

func (app *UmeeApp) registerUpgrade6(upgradeInfo upgradetypes.Plan) {
	planName := "v6.0"
	gravityModuleName := "gravity" // hardcoded to avoid dependency on GB module
	emergencyGroup, err := sdk.AccAddressFromBech32("umee1gy3c8n2xysawysq2xf2hxn253srx4ehduevq6c")
	util.Panic(err)

	app.UpgradeKeeper.SetUpgradeHandler(planName,
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			printPlanName(planName, ctx.Logger())
			if err := app.LeverageKeeper.SetParams(ctx, leveragetypes.DefaultParams()); err != nil {
				return fromVM, err
			}
			app.UGovKeeperB.Keeper(&ctx).SetEmergencyGroup(emergencyGroup)

			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	app.storeUpgrade(planName, upgradeInfo, storetypes.StoreUpgrades{
		Deleted: []string{gravityModuleName},
	})
}

func (app *UmeeApp) registerUpgrade5_1(upgradeInfo upgradetypes.Plan) {
	planName := "v5.1"
	app.UpgradeKeeper.SetUpgradeHandler(planName,
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			// GravityBridge is deleted after v5.1
			// if err := app.GravityKeeper.MigrateFundsToDrainAccount(
			// 	ctx,
			// 	sdk.MustAccAddressFromBech32("umee1gx9svenfs6ktvajje2wgqau3gk5mznwnyghq4l"),
			// ); err != nil {
			// 	return nil, err
			// }
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	app.storeUpgrade(planName, upgradeInfo, storetypes.StoreUpgrades{
		Added: []string{incentive.ModuleName},
	})
}

// performs upgrade from v4.2 to v4.3
func (app *UmeeApp) registerUpgrade4_3(upgradeInfo upgradetypes.Plan) {
	const planName = "v4.3"
	app.UpgradeKeeper.SetUpgradeHandler(planName, onlyModuleMigrations(app, planName))
	app.UpgradeKeeper.SetUpgradeHandler(planName,
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			ctx.Logger().Info("Upgrade handler execution", "name", planName)

			// set the ICS27 consensus version so InitGenesis is not run
			oldIcaVersion := fromVM[icatypes.ModuleName]
			g := icagenesis.GenesisState{HostGenesisState: icagenesis.DefaultHostGenesis()}
			g.HostGenesisState.Params.AllowMessages = []string{
				sdk.MsgTypeURL(&banktypes.MsgSend{}),
				sdk.MsgTypeURL(&stakingtypes.MsgDelegate{}),
				sdk.MsgTypeURL(&stakingtypes.MsgBeginRedelegate{}),
				sdk.MsgTypeURL(&stakingtypes.MsgUndelegate{}),
				sdk.MsgTypeURL(&stakingtypes.MsgCancelUnbondingDelegation{}),
				sdk.MsgTypeURL(&stakingtypes.MsgCreateValidator{}),
				sdk.MsgTypeURL(&stakingtypes.MsgEditValidator{}),
				sdk.MsgTypeURL(&distrtypes.MsgWithdrawDelegatorReward{}),
				sdk.MsgTypeURL(&distrtypes.MsgSetWithdrawAddress{}),
				sdk.MsgTypeURL(&distrtypes.MsgWithdrawValidatorCommission{}),
				sdk.MsgTypeURL(&distrtypes.MsgFundCommunityPool{}),
				sdk.MsgTypeURL(&govv1.MsgVote{}),
				sdk.MsgTypeURL(&govv1beta1.MsgVote{}),

				sdk.MsgTypeURL(&ibctransfertypes.MsgTransfer{}),
			}
			// initialize ICS27 module
			icamodule, ok := app.mm.Modules[icatypes.ModuleName].(ica.AppModule)
			if !ok {
				panic("Modules[icatypes.ModuleName] is not of type ica.AppModule")
			}
			// skip InitModule in upgrade tests after the upgrade has gone through.
			if oldIcaVersion != fromVM[icatypes.ModuleName] {
				icamodule.InitModule(ctx, g.ControllerGenesisState.Params, g.HostGenesisState.Params)
			}

			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	app.storeUpgrade(planName, upgradeInfo, storetypes.StoreUpgrades{
		Added: []string{
			icahosttypes.StoreKey,
		},
	})
}

// performs upgrade from v4.0 to v4.1
func (app *UmeeApp) registerUpgrade4_1(_ upgradetypes.Plan) {
	const planName = "v4.1.0"
	app.UpgradeKeeper.SetUpgradeHandler(planName,
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			ctx.Logger().Info("Upgrade handler execution", "name", planName)
			leverageUpgrader := leveragekeeper.NewMigrator(&app.LeverageKeeper)
			migrated, err := leverageUpgrader.MigrateBNB(ctx)
			if err != nil {
				ctx.Logger().Error("Error in v4.1 leverage Migration!", "err", err)
				return fromVM, err
			}
			if migrated {
				// If leverage BNB migration was skipped, also skip oracle so they stay in sync
				oracleUpgrader := oraclemigrator.NewMigrator(&app.OracleKeeper)
				oracleUpgrader.MigrateBNB(ctx)
			}
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)
}

// performs upgrade from v3.3 -> v4
func (app *UmeeApp) registerUpgrade3_3to4_0(_ upgradetypes.Plan) {
	const planName = "v4.0"
	app.UpgradeKeeper.SetUpgradeHandler(planName,
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			ctx.Logger().Info("Upgrade handler execution", "name", planName)
			ctx.Logger().Info("Run v4.0 migration")
			upgrader := oraclemigrator.NewMigrator(&app.OracleKeeper)
			err := upgrader.HistoracleParams3x4(ctx)
			if err != nil {
				ctx.Logger().Error("Unable to run v4.0 Migration!", "err", err)
				return fromVM, err
			}
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)
}

// performs upgrade from v3.1 -> v3.3 (including the v3.2 changes)
func (app *UmeeApp) registerUpgrade3_1to3_3(_ upgradetypes.Plan) {
	const planName = "v3.1-v3.3"
	app.UpgradeKeeper.SetUpgradeHandler(
		planName,
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			ctx.Logger().Info("Upgrade handler execution", "name", planName)
			ctx.Logger().Info("Run v3.3 migrator")
			err := upgradev3x3.Migrator(*app.GovKeeper, app.interfaceRegistry)(ctx)
			if err != nil {
				return fromVM, err
			}
			ctx.Logger().Info("Run x/bank v0.46.5 migration")
			// err = bankkeeper.NewMigrator(app.BankKeeper, app.GetSubspace(banktypes.ModuleName)).Migrate3_V046_4_To_V046_5(ctx)
			if err != nil {
				return fromVM, err
			}
			ctx.Logger().Info("Run module migrations")
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		})
}

// performs upgrade from v3.2 -> v3.3
func (app *UmeeApp) registerUpgrade3_2to3_3(_ upgradetypes.Plan) {
	const planName = "v3.2-v3.3"
	app.UpgradeKeeper.SetUpgradeHandler(
		planName,
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			ctx.Logger().Info("Upgrade handler execution", "name", planName)
			ctx.Logger().Info("Run v3.3 migrator")
			err := upgradev3x3.Migrator(*app.GovKeeper, app.interfaceRegistry)(ctx)
			if err != nil {
				return fromVM, err
			}
			ctx.Logger().Info("Run module migrations")
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		})
}

// performs upgrade from v1->v3
func (app *UmeeApp) registerUpgrade3_0(upgradeInfo upgradetypes.Plan) {
	const planName = "v1.1-v3.0"
	app.UpgradeKeeper.SetUpgradeHandler(
		planName,
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			ctx.Logger().Info("Upgrade handler execution", "name", planName)
			// ctx.Logger().Info("Running setupBech32ibcKeeper")
			// err := upgradev3.SetupBech32ibcKeeper(&app.bech32IbcKeeper, ctx)
			// if err != nil {
			// 	return nil, errors.Wrapf(
			// 		err, "%q Upgrade: Unable to upgrade, bech32ibc module not initialized", planName)
			// }

			ctx.Logger().Info("Running module migrations")
			vm, err := app.mm.RunMigrations(ctx, app.configurator, fromVM)
			if err != nil {
				return vm, err
			}

			// ctx.Logger().Info("Updating validator minimum commission rate param of staking module")
			// minCommissionRate, err := upgradev3.UpdateMinimumCommissionRateParam(ctx, app.StakingKeeper)
			// if err != nil {
			// 	return vm, errors.Wrapf(
			// 		err, "%q Upgrade: failed to update minimum commission rate param of staking module",
			// 		planName)
			// }

			// ctx.Logger().Info("Upgrade handler execution finished, updating minimum commission rate of all validators",
			// 	"name", planName)
			// err = upgradev3.SetMinimumCommissionRateToValidators(ctx, app.StakingKeeper, minCommissionRate)
			// if err != nil {
			// 	return vm, errors.Wrapf(
			// 		err, "%q Upgrade: failed to update minimum commission rate for validators",
			// 		planName)
			// }

			return vm, err
		})

	app.storeUpgrade(planName, upgradeInfo, storetypes.StoreUpgrades{
		Added: []string{
			group.ModuleName,
			nft.ModuleName,
			// bech32ibctypes.ModuleName, // removed dependency
			oracletypes.ModuleName,
			leveragetypes.ModuleName,
		},
	})
}

func onlyModuleMigrations(app *UmeeApp, planName string) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		printPlanName(planName, ctx.Logger())
		ctx.Logger().Info("-----------------------------\n-----------------------------")
		ctx.Logger().Info("Upgrade handler execution", "name", planName)
		return app.mm.RunMigrations(ctx, app.configurator, fromVM)
	}
}

// helper function to check if the store loader should be upgraded
func (app *UmeeApp) storeUpgrade(planName string, ui upgradetypes.Plan, stores storetypes.StoreUpgrades) {
	if ui.Name == planName && !app.UpgradeKeeper.IsSkipHeight(ui.Height) {
		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(
			upgradetypes.UpgradeStoreLoader(ui.Height, &stores))
	}
}

// registerUpgrade sets an upgrade handler which only runs module migrations
// and adds new storages storages
func (app *UmeeApp) registerUpgrade(planName string, upgradeInfo upgradetypes.Plan, newStores ...string) {
	app.UpgradeKeeper.SetUpgradeHandler(planName, onlyModuleMigrations(app, planName))

	if len(newStores) > 0 {
		app.storeUpgrade(planName, upgradeInfo, storetypes.StoreUpgrades{
			Added: newStores,
		})
	}
}

func printPlanName(planName string, logger log.Logger) {
	logger.Info("-----------------------------\n-----------------------------")
	logger.Info("Upgrade handler execution", "name", planName)
}
