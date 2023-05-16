package ugov

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/umee-network/umee/v4/util/checkers"
)

var (
	_ sdk.Msg = &MsgGovUpdateMinGasPrice{}
)

// ValidateBasic implements Msg
func (msg *MsgGovUpdateMinGasPrice) ValidateBasic() error {
	if err := checkers.ValidateAddr(msg.Authority, "authority"); err != nil {
		return err
	}

	return msg.MinGasPrice.Validate()
}

// GetSignBytes implements Msg
func (msg *MsgGovUpdateMinGasPrice) GetSigners() []sdk.AccAddress {
	return checkers.Signers(msg.Authority)
}

// GetSignBytes implements Stringer method
func (msg *MsgGovUpdateMinGasPrice) String() string {
	return fmt.Sprintf("<authority: %s, min_gas_price: %s>", msg.Authority, msg.MinGasPrice.String())
}
