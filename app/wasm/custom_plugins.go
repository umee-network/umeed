package wasm

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/umee-network/umee/v4/app/wasm/msg"
	"github.com/umee-network/umee/v4/app/wasm/query"
	stargate "github.com/umee-network/umee/v4/app/wasm/stargate"
	leveragekeeper "github.com/umee-network/umee/v4/x/leverage/keeper"
	oraclekeeper "github.com/umee-network/umee/v4/x/oracle/keeper"
)

// RegisterCustomPlugins expose the queries and msgs of native modules to wasm.
func RegisterCustomPlugins(
	leverageKeeper leveragekeeper.Keeper,
	oracleKeeper oraclekeeper.Keeper,
) []wasmkeeper.Option {
	wasmQueryPlugin := query.NewQueryPlugin(leverageKeeper, oracleKeeper)
	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: wasmQueryPlugin.CustomQuerier(),
	})

	messagePluginOpt := wasmkeeper.WithMessageHandlerDecorator(msg.NewMessagePlugin(leverageKeeper))

	return []wasm.Option{
		queryPluginOpt,
		messagePluginOpt,
	}
}

// RegisterStargateQueries expose the stargate queries
func RegisterStargateQueries(queryRouter baseapp.GRPCQueryRouter, codec codec.Codec) []wasmkeeper.Option {
	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Stargate: stargate.StargateQuerier(queryRouter, codec),
	})

	return []wasmkeeper.Option{
		queryPluginOpt,
	}
}
