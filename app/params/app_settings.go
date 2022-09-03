package params

import (
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// Name defines the application name of the Umee network.
	Name = "umee"

	// BondDenom defines the native staking token denomination.
	BondDenom = "uumee"

	// DisplayDenom defines the name, symbol, and display value of the umee token.
	DisplayDenom = "UMEE"
)

var (
	// MinMinGasPrice is the minimum value a validator can set for `minimum-gas-prices` his app.toml config
	MinMinGasPrice = sdk.NewDecCoinFromDec(BondDenom, sdk.MustNewDecFromStr("0.05"))
)

func init() {
	// XXX: If other upstream or external application's depend on any of Umee's
	// CLI or command functionality, then this would require us to move the
	// SetAddressConfig call to somewhere external such as the root command
	// constructor and anywhere else we contract the app.
	SetAddressConfig()

	if AccountAddressPrefix != Name {
		log.Fatal("AccountAddresPrefix must equal Name")
	}
}
