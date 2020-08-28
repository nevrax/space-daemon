// Code generated by mockery v2.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	hub "github.com/FleekHQ/space-daemon/core/textile/hub"
	mock "github.com/stretchr/testify/mock"
)

// HubAuth is an autogenerated mock type for the HubAuth type
type HubAuth struct {
	mock.Mock
}

// GetHubContext provides a mock function with given fields: ctx
func (_m *HubAuth) GetHubContext(ctx context.Context) (context.Context, error) {
	ret := _m.Called(ctx)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context) context.Context); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokensWithCache provides a mock function with given fields: ctx
func (_m *HubAuth) GetTokensWithCache(ctx context.Context) (*hub.AuthTokens, error) {
	ret := _m.Called(ctx)

	var r0 *hub.AuthTokens
	if rf, ok := ret.Get(0).(func(context.Context) *hub.AuthTokens); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*hub.AuthTokens)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
