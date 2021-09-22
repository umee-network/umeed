package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "umee" // TODO: "leverage"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability" // Question: Should this be "mem-leverage" instead?

	// Prefixes to be used when storing info about individual asset or uToken types
	// Note: The full list will likely be in ADR 003
	AssetAssociatedUtokenPrefix = 0x01 // Stores each asset denom's accepted uToken
	UtokenAssociatedAssetPrefix = 0x02 // Stores each uToken denom's accepted asset
	// TODO: Add more prefixes, like AssetAssociatedOvercollatRequirement
)

func KeyPrefix(p string) []byte {
	return []byte(p) // Question: Is this currently used, and should we forceASCII?
}

// returns store key used to store a specific asset coin's associated utoken denom
func assetAssociatedUtokenKey(coin sdk.Coin) []byte {
	return prefixDenomStoreKey(AssetAssociatedUtokenPrefix, coin)
	// intent: store["leverage_asset_utoken_uatom"] = "u/uatom"
}

// returns store key used to store a specific utoken coin's associated asset denom
func utokenAssociatedAssetKey(coin sdk.Coin) []byte {
	return prefixDenomStoreKey(UtokenAssociatedAssetPrefix, coin)
	// intent: store["leverage_utoken_asset_u/uatom"] = "uatom"
}

// prefixDenomStoreKey turns a coin to the key used to store specific info by appending a prefix to the denom.
// Returns empty bytes on invalid denom. Modeled after x/auth/types.AccountStoreKey
func prefixDenomStoreKey(prefix byte, coin sdk.Coin) []byte {
	if sdk.ValidateDenom(coin.Denom) != nil {
		// Denom did not match ^[a-z][a-z0-9/]{2,63}$
		return []byte{}
	}
	// example: byte(0x01) + []byte("uatom")
	key := []byte{prefix}
	key = append(key, forceASCII(coin.Denom)...)
	return key
	// Note: After IBC enable, want a reliable way to convert token denominations to bytes, such that
	//	a) Each token type (e.g. Atom) has only one prefix/key, regardless of the path it took via IBC
	//		to get to umee
	//	b) Tokens cannot be spoofed (e.g. EvilChain cannot name a token 'atom', mint their own, and deposit)
}

// TODO: Review if this is necessary
// internal (string normalization): reliably convert denoms to ascii bytes, and strip non-ascii
func forceASCII(s string) []byte {
	// forceASCII("Hello, 世界!") // => []byte("Hello, !")
	b := []byte{}
	for _, r := range s {
		// Iterates over 'runes', which are often UTF-8
		if r <= 127 {
			b = append(b, byte(r))
		}
	}
	return b
}
