// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	log "github.com/smartcontractkit/chainlink/core/services/log"
	mock "github.com/stretchr/testify/mock"
)

// Broadcaster is an autogenerated mock type for the Broadcaster type
type Broadcaster struct {
	mock.Mock
}

// AddDependents provides a mock function with given fields: n
func (_m *Broadcaster) AddDependents(n int) {
	_m.Called(n)
}

// AwaitDependents provides a mock function with given fields:
func (_m *Broadcaster) AwaitDependents() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// DependentReady provides a mock function with given fields:
func (_m *Broadcaster) DependentReady() {
	_m.Called()
}

// Register provides a mock function with given fields: address, listener
func (_m *Broadcaster) Register(address common.Address, listener log.Listener) bool {
	ret := _m.Called(address, listener)

	var r0 bool
	if rf, ok := ret.Get(0).(func(common.Address, log.Listener) bool); ok {
		r0 = rf(address, listener)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *Broadcaster) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *Broadcaster) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unregister provides a mock function with given fields: address, listener
func (_m *Broadcaster) Unregister(address common.Address, listener log.Listener) {
	_m.Called(address, listener)
}