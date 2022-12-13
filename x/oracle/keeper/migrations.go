package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/umee-network/umee/v3/x/oracle/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper *Keeper
}

// NewMigrator creates a Migrator.
func NewMigrator(keeper *Keeper) Migrator {
	return Migrator{keeper: keeper}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	m.keeper.SetHistoricStampPeriod(ctx, 1)
	m.keeper.SetMedianStampPeriod(ctx, 1)
	m.keeper.SetMaximumPriceStamps(ctx, 1)
	m.keeper.SetMaximumMedianStamps(ctx, 1)
	return nil
}

// Migrate2to3 migrates from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	m.keeper.SetHistoricStampPeriod(ctx, types.DefaultHistoricStampPeriod)
	m.keeper.SetMedianStampPeriod(ctx, types.DefaultMedianStampPeriod)
	m.keeper.SetMaximumPriceStamps(ctx, types.DefaultMaximumPriceStamps)
	m.keeper.SetMaximumMedianStamps(ctx, types.DefaultMaximumMedianStamps)
	return nil
}
