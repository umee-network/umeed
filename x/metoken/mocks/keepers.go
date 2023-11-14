// Code generated by MockGen. DO NOT EDIT.
// Source: ./../../x/metoken/expected_keepers.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	math "cosmossdk.io/math"
	types "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
	types0 "github.com/umee-network/umee/v6/x/leverage/types"
	types1 "github.com/umee-network/umee/v6/x/oracle/types"
)

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// BurnCoins mocks base method.
func (m *MockBankKeeper) BurnCoins(ctx types.Context, moduleName string, amounts types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BurnCoins", ctx, moduleName, amounts)
	ret0, _ := ret[0].(error)
	return ret0
}

// BurnCoins indicates an expected call of BurnCoins.
func (mr *MockBankKeeperMockRecorder) BurnCoins(ctx, moduleName, amounts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BurnCoins", reflect.TypeOf((*MockBankKeeper)(nil).BurnCoins), ctx, moduleName, amounts)
}

// MintCoins mocks base method.
func (m *MockBankKeeper) MintCoins(ctx types.Context, moduleName string, amounts types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintCoins", ctx, moduleName, amounts)
	ret0, _ := ret[0].(error)
	return ret0
}

// MintCoins indicates an expected call of MintCoins.
func (mr *MockBankKeeperMockRecorder) MintCoins(ctx, moduleName, amounts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintCoins", reflect.TypeOf((*MockBankKeeper)(nil).MintCoins), ctx, moduleName, amounts)
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromAccountToModule(ctx types.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx types.Context, senderModule string, recipientAddr types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// MockLeverageKeeper is a mock of LeverageKeeper interface.
type MockLeverageKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockLeverageKeeperMockRecorder
}

// MockLeverageKeeperMockRecorder is the mock recorder for MockLeverageKeeper.
type MockLeverageKeeperMockRecorder struct {
	mock *MockLeverageKeeper
}

// NewMockLeverageKeeper creates a new mock instance.
func NewMockLeverageKeeper(ctrl *gomock.Controller) *MockLeverageKeeper {
	mock := &MockLeverageKeeper{ctrl: ctrl}
	mock.recorder = &MockLeverageKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeverageKeeper) EXPECT() *MockLeverageKeeperMockRecorder {
	return m.recorder
}

// GetAllSupplied mocks base method.
func (m *MockLeverageKeeper) GetAllSupplied(ctx types.Context, supplierAddr types.AccAddress) (types.Coins, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSupplied", ctx, supplierAddr)
	ret0, _ := ret[0].(types.Coins)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSupplied indicates an expected call of GetAllSupplied.
func (mr *MockLeverageKeeperMockRecorder) GetAllSupplied(ctx, supplierAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSupplied", reflect.TypeOf((*MockLeverageKeeper)(nil).GetAllSupplied), ctx, supplierAddr)
}

// GetTokenSettings mocks base method.
func (m *MockLeverageKeeper) GetTokenSettings(ctx types.Context, denom string) (types0.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokenSettings", ctx, denom)
	ret0, _ := ret[0].(types0.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTokenSettings indicates an expected call of GetTokenSettings.
func (mr *MockLeverageKeeperMockRecorder) GetTokenSettings(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenSettings", reflect.TypeOf((*MockLeverageKeeper)(nil).GetTokenSettings), ctx, denom)
}

// GetTotalSupply mocks base method.
func (m *MockLeverageKeeper) GetTotalSupply(ctx types.Context, denom string) (types.Coin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalSupply", ctx, denom)
	ret0, _ := ret[0].(types.Coin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalSupply indicates an expected call of GetTotalSupply.
func (mr *MockLeverageKeeperMockRecorder) GetTotalSupply(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalSupply", reflect.TypeOf((*MockLeverageKeeper)(nil).GetTotalSupply), ctx, denom)
}

// ModuleMaxWithdraw mocks base method.
func (m *MockLeverageKeeper) ModuleMaxWithdraw(ctx types.Context, spendableUTokens types.Coin) (math.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModuleMaxWithdraw", ctx, spendableUTokens)
	ret0, _ := ret[0].(math.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModuleMaxWithdraw indicates an expected call of ModuleMaxWithdraw.
func (mr *MockLeverageKeeperMockRecorder) ModuleMaxWithdraw(ctx, spendableUTokens interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModuleMaxWithdraw", reflect.TypeOf((*MockLeverageKeeper)(nil).ModuleMaxWithdraw), ctx, spendableUTokens)
}

// SupplyFromModule mocks base method.
func (m *MockLeverageKeeper) SupplyFromModule(ctx types.Context, fromModule string, coin types.Coin) (types.Coin, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SupplyFromModule", ctx, fromModule, coin)
	ret0, _ := ret[0].(types.Coin)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SupplyFromModule indicates an expected call of SupplyFromModule.
func (mr *MockLeverageKeeperMockRecorder) SupplyFromModule(ctx, fromModule, coin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupplyFromModule", reflect.TypeOf((*MockLeverageKeeper)(nil).SupplyFromModule), ctx, fromModule, coin)
}

// ToToken mocks base method.
func (m *MockLeverageKeeper) ToToken(ctx types.Context, uToken types.Coin) (types.Coin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToToken", ctx, uToken)
	ret0, _ := ret[0].(types.Coin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToToken indicates an expected call of ToToken.
func (mr *MockLeverageKeeperMockRecorder) ToToken(ctx, uToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToToken", reflect.TypeOf((*MockLeverageKeeper)(nil).ToToken), ctx, uToken)
}

// ToUToken mocks base method.
func (m *MockLeverageKeeper) ToUToken(ctx types.Context, token types.Coin) (types.Coin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToUToken", ctx, token)
	ret0, _ := ret[0].(types.Coin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToUToken indicates an expected call of ToUToken.
func (mr *MockLeverageKeeperMockRecorder) ToUToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToUToken", reflect.TypeOf((*MockLeverageKeeper)(nil).ToUToken), ctx, token)
}

// WithdrawToModule mocks base method.
func (m *MockLeverageKeeper) WithdrawToModule(ctx types.Context, toModule string, uToken types.Coin) (types.Coin, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithdrawToModule", ctx, toModule, uToken)
	ret0, _ := ret[0].(types.Coin)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// WithdrawToModule indicates an expected call of WithdrawToModule.
func (mr *MockLeverageKeeperMockRecorder) WithdrawToModule(ctx, toModule, uToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawToModule", reflect.TypeOf((*MockLeverageKeeper)(nil).WithdrawToModule), ctx, toModule, uToken)
}

// MockOracleKeeper is a mock of OracleKeeper interface.
type MockOracleKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockOracleKeeperMockRecorder
}

// MockOracleKeeperMockRecorder is the mock recorder for MockOracleKeeper.
type MockOracleKeeperMockRecorder struct {
	mock *MockOracleKeeper
}

// NewMockOracleKeeper creates a new mock instance.
func NewMockOracleKeeper(ctrl *gomock.Controller) *MockOracleKeeper {
	mock := &MockOracleKeeper{ctrl: ctrl}
	mock.recorder = &MockOracleKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOracleKeeper) EXPECT() *MockOracleKeeperMockRecorder {
	return m.recorder
}

// AllMedianPrices mocks base method.
func (m *MockOracleKeeper) AllMedianPrices(ctx types.Context) types1.Prices {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllMedianPrices", ctx)
	ret0, _ := ret[0].(types1.Prices)
	return ret0
}

// AllMedianPrices indicates an expected call of AllMedianPrices.
func (mr *MockOracleKeeperMockRecorder) AllMedianPrices(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllMedianPrices", reflect.TypeOf((*MockOracleKeeper)(nil).AllMedianPrices), ctx)
}

// SetExchangeRate mocks base method.
func (m *MockOracleKeeper) SetExchangeRate(ctx types.Context, denom string, rate types.Dec) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetExchangeRate", ctx, denom, rate)
}

// SetExchangeRate indicates an expected call of SetExchangeRate.
func (mr *MockOracleKeeperMockRecorder) SetExchangeRate(ctx, denom, rate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetExchangeRate", reflect.TypeOf((*MockOracleKeeper)(nil).SetExchangeRate), ctx, denom, rate)
}
