package wasm

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/umee-network/umee/v6/app/wasm/msg"
	"github.com/umee-network/umee/v6/app/wasm/query"
	inckeeper "github.com/umee-network/umee/v6/x/incentive/keeper"
	leveragekeeper "github.com/umee-network/umee/v6/x/leverage/keeper"
	metokenkeeper "github.com/umee-network/umee/v6/x/metoken/keeper"
	oraclekeeper "github.com/umee-network/umee/v6/x/oracle/keeper"
)

// RegisterCustomPlugins expose the queries and msgs of native modules to wasm.
func RegisterCustomPlugins(
	leverageKeeper leveragekeeper.Keeper,
	oracleKeeper oraclekeeper.Keeper,
	incentiveKeeper inckeeper.Keeper,
	metokenBuilder metokenkeeper.Builder,
) []wasmkeeper.Option {
	wasmQueryPlugin := query.NewQueryPlugin(leverageKeeper, oracleKeeper, incentiveKeeper, metokenBuilder)
	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: wasmQueryPlugin.CustomQuerier(),
	})

	messagePluginOpt := wasmkeeper.WithMessageHandlerDecorator(msg.NewMessagePlugin(leverageKeeper))

	return []wasmkeeper.Option{
		queryPluginOpt,
		messagePluginOpt,
	}
}

// RegisterStargateQueries expose the stargate queries
func RegisterStargateQueries(queryRouter baseapp.GRPCQueryRouter, codec codec.Codec) []wasmkeeper.Option {
	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Stargate: query.StargateQuerier(queryRouter, codec),
	})

	return []wasmkeeper.Option{
		queryPluginOpt,
	}
}
