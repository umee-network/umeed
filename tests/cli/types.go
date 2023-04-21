package cli_itest

import (
	"cosmossdk.io/errors"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"
)

type TestTransaction struct {
	Name        string
	Command     *cobra.Command
	Args        []string
	ExpectedErr *errors.Error
}

type TestQuery struct {
	Name             string
	Command          *cobra.Command
	Args             []string
	ResponseType     proto.Message
	ExpectedResponse proto.Message
	ErrMsg           string
}
