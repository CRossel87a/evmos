// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	client "github.com/cosmos/cosmos-sdk/client"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// AccountRetriever is an autogenerated mock type for the AccountRetriever type
type AccountRetriever struct {
	mock.Mock
}

// EnsureExists provides a mock function with given fields: clientCtx, addr
func (_m *AccountRetriever) EnsureExists(clientCtx client.Context, addr types.AccAddress) error {
	ret := _m.Called(clientCtx, addr)

	var r0 error
	if rf, ok := ret.Get(0).(func(client.Context, types.AccAddress) error); ok {
		r0 = rf(clientCtx, addr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAccount provides a mock function with given fields: clientCtx, addr
func (_m *AccountRetriever) GetAccount(clientCtx client.Context, addr types.AccAddress) (client.Account, error) {
	ret := _m.Called(clientCtx, addr)

	var r0 client.Account
	if rf, ok := ret.Get(0).(func(client.Context, types.AccAddress) client.Account); ok {
		r0 = rf(clientCtx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(client.Context, types.AccAddress) error); ok {
		r1 = rf(clientCtx, addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountNumberSequence provides a mock function with given fields: clientCtx, addr
func (_m *AccountRetriever) GetAccountNumberSequence(clientCtx client.Context, addr types.AccAddress) (uint64, uint64, error) {
	ret := _m.Called(clientCtx, addr)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(client.Context, types.AccAddress) uint64); ok {
		r0 = rf(clientCtx, addr)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(client.Context, types.AccAddress) uint64); ok {
		r1 = rf(clientCtx, addr)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(client.Context, types.AccAddress) error); ok {
		r2 = rf(clientCtx, addr)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAccountWithHeight provides a mock function with given fields: clientCtx, addr
func (_m *AccountRetriever) GetAccountWithHeight(clientCtx client.Context, addr types.AccAddress) (client.Account, int64, error) {
	ret := _m.Called(clientCtx, addr)

	var r0 client.Account
	if rf, ok := ret.Get(0).(func(client.Context, types.AccAddress) client.Account); ok {
		r0 = rf(clientCtx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Account)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(client.Context, types.AccAddress) int64); ok {
		r1 = rf(clientCtx, addr)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(client.Context, types.AccAddress) error); ok {
		r2 = rf(clientCtx, addr)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewAccountRetriever interface {
	mock.TestingT
	Cleanup(func())
}

// NewAccountRetriever creates a new instance of AccountRetriever. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAccountRetriever(t mockConstructorTestingTNewAccountRetriever) *AccountRetriever {
	mock := &AccountRetriever{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
