package uibc

import (
	"github.com/umee-network/umee/v4/util"
)

const (
	// ModuleName defines the module name
	ModuleName = "uibc"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for uibc
	RouterKey = ModuleName
)

var (
	KeyPrefixDenomOutflows = []byte{0x01}
	KeyPrefixTotalOutflows = []byte{0x02}
	// KeyPrefixParams is the key to query all gov params
	KeyPrefixParams = []byte{0x03}
	// KeyPrefixQuotaExpires is the key to store the next expire time of ibc-transfer quota
	KeyPrefixQuotaExpires = []byte{0x04}
)

func KeyTotalOutflows(ibcDenom string) []byte {
	//  KeyPrefixDenomQuota| denom
	return util.ConcatBytes(0, KeyPrefixDenomOutflows, []byte(ibcDenom))
}

// UIBCInflowQuotaEnabled
func UIBCInflowQuotaEnabled(status IBCTransferQuotaStatus) bool {
	// outflow disabled means inflow check enabled
	if status == IBCTransferQuotaStatus_QUOTA_ENABLED || status == IBCTransferQuotaStatus_QUOTA_OUT_DISABLED {
		return true
	}
	return false
}

// UIBCOutflowQuotaEnabled
func UIBCOutflowQuotaEnabled(status IBCTransferQuotaStatus) bool {
	// inflow disabled means outflow check enabled
	if status == IBCTransferQuotaStatus_QUOTA_ENABLED || status == IBCTransferQuotaStatus_QUOTA_IN_DISABLED {
		return true
	}
	return false
}
