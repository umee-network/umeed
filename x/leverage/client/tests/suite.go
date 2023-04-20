package tests

import (
	"github.com/cosmos/cosmos-sdk/testutil/network"
	itestsuite "github.com/umee-network/umee/v4/tests/integration_suite"
)

type IntegrationTests struct {
	*itestsuite.IntegrationTestSuite
}

func NewIntegrationTestSuite(cfg network.Config) *IntegrationTests {
	return &IntegrationTests{
		&itestsuite.IntegrationTestSuite{Cfg: cfg},
	}
}
