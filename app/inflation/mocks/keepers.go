// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/inflation/expected_keepers.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	math "cosmossdk.io/math"
	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/mint/types"
	gomock "github.com/golang/mock/gomock"
	ugov "github.com/umee-network/umee/v5/x/ugov"
	keeper "github.com/umee-network/umee/v5/x/ugov/keeper"
)

// MockMintKeeper is a mock of MintKeeper interface.
type MockMintKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockMintKeeperMockRecorder
}

// MockMintKeeperMockRecorder is the mock recorder for MockMintKeeper.
type MockMintKeeperMockRecorder struct {
	mock *MockMintKeeper
}

// NewMockMintKeeper creates a new mock instance.
func NewMockMintKeeper(ctrl *gomock.Controller) *MockMintKeeper {
	mock := &MockMintKeeper{ctrl: ctrl}
	mock.recorder = &MockMintKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMintKeeper) EXPECT() *MockMintKeeperMockRecorder {
	return m.recorder
}

// SetParams mocks base method.
func (m *MockMintKeeper) SetParams(ctx types.Context, params types0.Params) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetParams", ctx, params)
}

// SetParams indicates an expected call of SetParams.
func (mr *MockMintKeeperMockRecorder) SetParams(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetParams", reflect.TypeOf((*MockMintKeeper)(nil).SetParams), ctx, params)
}

// StakingTokenSupply mocks base method.
func (m *MockMintKeeper) StakingTokenSupply(ctx types.Context) math.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StakingTokenSupply", ctx)
	ret0, _ := ret[0].(math.Int)
	return ret0
}

// StakingTokenSupply indicates an expected call of StakingTokenSupply.
func (mr *MockMintKeeperMockRecorder) StakingTokenSupply(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StakingTokenSupply", reflect.TypeOf((*MockMintKeeper)(nil).StakingTokenSupply), ctx)
}

// MockUGovKeeper is a mock of UGovKeeper interface.
type MockUGovKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockUGovKeeperMockRecorder
}

// MockUGovKeeperMockRecorder is the mock recorder for MockUGovKeeper.
type MockUGovKeeperMockRecorder struct {
	mock *MockUGovKeeper
}

// NewMockUGovKeeper creates a new mock instance.
func NewMockUGovKeeper(ctrl *gomock.Controller) *MockUGovKeeper {
	mock := &MockUGovKeeper{ctrl: ctrl}
	mock.recorder = &MockUGovKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUGovKeeper) EXPECT() *MockUGovKeeperMockRecorder {
	return m.recorder
}

// EmergencyGroup mocks base method.
func (m *MockUGovKeeper) EmergencyGroup() types.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EmergencyGroup")
	ret0, _ := ret[0].(types.AccAddress)
	return ret0
}

// EmergencyGroup indicates an expected call of EmergencyGroup.
func (mr *MockUGovKeeperMockRecorder) EmergencyGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmergencyGroup", reflect.TypeOf((*MockUGovKeeper)(nil).EmergencyGroup))
}

// ExportGenesis mocks base method.
func (m *MockUGovKeeper) ExportGenesis() *ugov.GenesisState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportGenesis")
	ret0, _ := ret[0].(*ugov.GenesisState)
	return ret0
}

// ExportGenesis indicates an expected call of ExportGenesis.
func (mr *MockUGovKeeperMockRecorder) ExportGenesis() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportGenesis", reflect.TypeOf((*MockUGovKeeper)(nil).ExportGenesis))
}

// GetInflationCycleEnd mocks base method.
func (m *MockUGovKeeper) GetInflationCycleEnd() (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInflationCycleEnd")
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInflationCycleEnd indicates an expected call of GetInflationCycleEnd.
func (mr *MockUGovKeeperMockRecorder) GetInflationCycleEnd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInflationCycleEnd", reflect.TypeOf((*MockUGovKeeper)(nil).GetInflationCycleEnd))
}

// InflationParams mocks base method.
func (m *MockUGovKeeper) InflationParams() ugov.InflationParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InflationParams")
	ret0, _ := ret[0].(ugov.InflationParams)
	return ret0
}

// InflationParams indicates an expected call of InflationParams.
func (mr *MockUGovKeeperMockRecorder) InflationParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InflationParams", reflect.TypeOf((*MockUGovKeeper)(nil).InflationParams))
}

// InitGenesis mocks base method.
func (m *MockUGovKeeper) InitGenesis(gs *ugov.GenesisState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitGenesis", gs)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitGenesis indicates an expected call of InitGenesis.
func (mr *MockUGovKeeperMockRecorder) InitGenesis(gs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitGenesis", reflect.TypeOf((*MockUGovKeeper)(nil).InitGenesis), gs)
}

// MinGasPrice mocks base method.
func (m *MockUGovKeeper) MinGasPrice() types.DecCoin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MinGasPrice")
	ret0, _ := ret[0].(types.DecCoin)
	return ret0
}

// MinGasPrice indicates an expected call of MinGasPrice.
func (mr *MockUGovKeeperMockRecorder) MinGasPrice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MinGasPrice", reflect.TypeOf((*MockUGovKeeper)(nil).MinGasPrice))
}

// SetEmergencyGroup mocks base method.
func (m *MockUGovKeeper) SetEmergencyGroup(p types.AccAddress) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetEmergencyGroup", p)
}

// SetEmergencyGroup indicates an expected call of SetEmergencyGroup.
func (mr *MockUGovKeeperMockRecorder) SetEmergencyGroup(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEmergencyGroup", reflect.TypeOf((*MockUGovKeeper)(nil).SetEmergencyGroup), p)
}

// SetInflationCycleEnd mocks base method.
func (m *MockUGovKeeper) SetInflationCycleEnd(startTime time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetInflationCycleEnd", startTime)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetInflationCycleEnd indicates an expected call of SetInflationCycleEnd.
func (mr *MockUGovKeeperMockRecorder) SetInflationCycleEnd(startTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInflationCycleEnd", reflect.TypeOf((*MockUGovKeeper)(nil).SetInflationCycleEnd), startTime)
}

// SetInflationParams mocks base method.
func (m *MockUGovKeeper) SetInflationParams(lp ugov.InflationParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetInflationParams", lp)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetInflationParams indicates an expected call of SetInflationParams.
func (mr *MockUGovKeeperMockRecorder) SetInflationParams(lp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInflationParams", reflect.TypeOf((*MockUGovKeeper)(nil).SetInflationParams), lp)
}

// SetMinGasPrice mocks base method.
func (m *MockUGovKeeper) SetMinGasPrice(p types.DecCoin) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMinGasPrice", p)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMinGasPrice indicates an expected call of SetMinGasPrice.
func (mr *MockUGovKeeperMockRecorder) SetMinGasPrice(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMinGasPrice", reflect.TypeOf((*MockUGovKeeper)(nil).SetMinGasPrice), p)
}

// MockUGovBKeeperI is a mock of UGovBKeeperI interface.
type MockUGovBKeeperI struct {
	ctrl     *gomock.Controller
	recorder *MockUGovBKeeperIMockRecorder
}

// MockUGovBKeeperIMockRecorder is the mock recorder for MockUGovBKeeperI.
type MockUGovBKeeperIMockRecorder struct {
	mock *MockUGovBKeeperI
}

// NewMockUGovBKeeperI creates a new mock instance.
func NewMockUGovBKeeperI(ctrl *gomock.Controller) *MockUGovBKeeperI {
	mock := &MockUGovBKeeperI{ctrl: ctrl}
	mock.recorder = &MockUGovBKeeperIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUGovBKeeperI) EXPECT() *MockUGovBKeeperIMockRecorder {
	return m.recorder
}

// Keeper mocks base method.
func (m *MockUGovBKeeperI) Keeper(ctx *types.Context) keeper.IKeeper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keeper", ctx)
	ret0, _ := ret[0].(keeper.IKeeper)
	return ret0
}

// Keeper indicates an expected call of Keeper.
func (mr *MockUGovBKeeperIMockRecorder) Keeper(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keeper", reflect.TypeOf((*MockUGovBKeeperI)(nil).Keeper), ctx)
}
