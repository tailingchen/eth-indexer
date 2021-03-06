// Code generated by mockery v1.0.0
package mocks

import big "math/big"
import client "github.com/getamis/eth-indexer/client"
import common "github.com/ethereum/go-ethereum/common"
import context "context"
import mock "github.com/stretchr/testify/mock"
import model "github.com/getamis/eth-indexer/model"
import state "github.com/ethereum/go-ethereum/core/state"

import types "github.com/ethereum/go-ethereum/core/types"

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// FindERC20 provides a mock function with given fields: address
func (_m *Manager) FindERC20(address common.Address) (*model.ERC20, error) {
	ret := _m.Called(address)

	var r0 *model.ERC20
	if rf, ok := ret.Get(0).(func(common.Address) *model.ERC20); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ERC20)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHeaderByNumber provides a mock function with given fields: number
func (_m *Manager) GetHeaderByNumber(number int64) (*model.Header, error) {
	ret := _m.Called(number)

	var r0 *model.Header
	if rf, ok := ret.Get(0).(func(int64) *model.Header); ok {
		r0 = rf(number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTd provides a mock function with given fields: hash
func (_m *Manager) GetTd(hash []byte) (*model.TotalDifficulty, error) {
	ret := _m.Called(hash)

	var r0 *model.TotalDifficulty
	if rf, ok := ret.Get(0).(func([]byte) *model.TotalDifficulty); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TotalDifficulty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields: balancer
func (_m *Manager) Init(balancer client.Balancer) error {
	ret := _m.Called(balancer)

	var r0 error
	if rf, ok := ret.Get(0).(func(client.Balancer) error); ok {
		r0 = rf(balancer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertERC20 provides a mock function with given fields: code
func (_m *Manager) InsertERC20(code *model.ERC20) error {
	ret := _m.Called(code)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.ERC20) error); ok {
		r0 = rf(code)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertTd provides a mock function with given fields: block, td
func (_m *Manager) InsertTd(block *types.Block, td *big.Int) error {
	ret := _m.Called(block, td)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Block, *big.Int) error); ok {
		r0 = rf(block, td)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LatestHeader provides a mock function with given fields:
func (_m *Manager) LatestHeader() (*model.Header, error) {
	ret := _m.Called()

	var r0 *model.Header
	if rf, ok := ret.Get(0).(func() *model.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBlocks provides a mock function with given fields: ctx, blocks, receipts, dumps, events, mode
func (_m *Manager) UpdateBlocks(ctx context.Context, blocks []*types.Block, receipts [][]*types.Receipt, dumps []*state.DirtyDump, events [][]*types.TransferLog, mode int) error {
	ret := _m.Called(ctx, blocks, receipts, dumps, events, mode)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*types.Block, [][]*types.Receipt, []*state.DirtyDump, [][]*types.TransferLog, int) error); ok {
		r0 = rf(ctx, blocks, receipts, dumps, events, mode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
