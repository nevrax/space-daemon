// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	client "github.com/textileio/textile/v2/api/users/client"

	mock "github.com/stretchr/testify/mock"

	thread "github.com/textileio/go-threads/core/thread"
)

// UsersClient is an autogenerated mock type for the UsersClient type
type UsersClient struct {
	mock.Mock
}

// ListInboxMessages provides a mock function with given fields: ctx, opts
func (_m *UsersClient) ListInboxMessages(ctx context.Context, opts ...client.ListOption) ([]client.Message, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []client.Message
	if rf, ok := ret.Get(0).(func(context.Context, ...client.ListOption) []client.Message); ok {
		r0 = rf(ctx, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...client.ListOption) error); ok {
		r1 = rf(ctx, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendMessage provides a mock function with given fields: ctx, from, to, body
func (_m *UsersClient) SendMessage(ctx context.Context, from thread.Identity, to thread.PubKey, body []byte) (client.Message, error) {
	ret := _m.Called(ctx, from, to, body)

	var r0 client.Message
	if rf, ok := ret.Get(0).(func(context.Context, thread.Identity, thread.PubKey, []byte) client.Message); ok {
		r0 = rf(ctx, from, to, body)
	} else {
		r0 = ret.Get(0).(client.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, thread.Identity, thread.PubKey, []byte) error); ok {
		r1 = rf(ctx, from, to, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetupMailbox provides a mock function with given fields: ctx
func (_m *UsersClient) SetupMailbox(ctx context.Context) (thread.ID, error) {
	ret := _m.Called(ctx)

	var r0 thread.ID
	if rf, ok := ret.Get(0).(func(context.Context) thread.ID); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(thread.ID)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
