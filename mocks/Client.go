// Code generated by mockery v2.0.0. DO NOT EDIT.

package mocks

import (
	config "github.com/FleekHQ/space-daemon/config"
	cid "github.com/ipfs/go-cid"
	client "github.com/textileio/go-threads/api/client"

	context "context"

	crypto "github.com/libp2p/go-libp2p-core/crypto"

	db "github.com/textileio/go-threads/db"

	domain "github.com/FleekHQ/space-daemon/core/space/domain"

	io "io"

	mock "github.com/stretchr/testify/mock"

	model "github.com/FleekHQ/space-daemon/core/textile/model"

	textile "github.com/FleekHQ/space-daemon/core/textile"

	usersclient "github.com/textileio/textile/v2/api/users/client"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// AcceptSharedFilesInvitation provides a mock function with given fields: ctx, invitation
func (_m *Client) AcceptSharedFilesInvitation(ctx context.Context, invitation domain.Invitation) (domain.Invitation, error) {
	ret := _m.Called(ctx, invitation)

	var r0 domain.Invitation
	if rf, ok := ret.Get(0).(func(context.Context, domain.Invitation) domain.Invitation); ok {
		r0 = rf(ctx, invitation)
	} else {
		r0 = ret.Get(0).(domain.Invitation)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Invitation) error); ok {
		r1 = rf(ctx, invitation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AttachMailboxNotifier provides a mock function with given fields: notif
func (_m *Client) AttachMailboxNotifier(notif textile.GrpcMailboxNotifier) {
	_m.Called(notif)
}

// CreateBucket provides a mock function with given fields: ctx, bucketSlug
func (_m *Client) CreateBucket(ctx context.Context, bucketSlug string) (textile.Bucket, error) {
	ret := _m.Called(ctx, bucketSlug)

	var r0 textile.Bucket
	if rf, ok := ret.Get(0).(func(context.Context, string) textile.Bucket); ok {
		r0 = rf(ctx, bucketSlug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(textile.Bucket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, bucketSlug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DownloadPublicGatewayItem provides a mock function with given fields: ctx, _a1
func (_m *Client) DownloadPublicGatewayItem(ctx context.Context, _a1 cid.Cid) (io.ReadCloser, error) {
	ret := _m.Called(ctx, _a1)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, cid.Cid) io.ReadCloser); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, cid.Cid) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucket provides a mock function with given fields: ctx, slug, remoteFile
func (_m *Client) GetBucket(ctx context.Context, slug string, remoteFile *textile.GetBucketForRemoteFileInput) (textile.Bucket, error) {
	ret := _m.Called(ctx, slug, remoteFile)

	var r0 textile.Bucket
	if rf, ok := ret.Get(0).(func(context.Context, string, *textile.GetBucketForRemoteFileInput) textile.Bucket); ok {
		r0 = rf(ctx, slug, remoteFile)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(textile.Bucket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *textile.GetBucketForRemoteFileInput) error); ok {
		r1 = rf(ctx, slug, remoteFile)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDefaultBucket provides a mock function with given fields: ctx
func (_m *Client) GetDefaultBucket(ctx context.Context) (textile.Bucket, error) {
	ret := _m.Called(ctx)

	var r0 textile.Bucket
	if rf, ok := ret.Get(0).(func(context.Context) textile.Bucket); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(textile.Bucket)
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

// GetFailedHealthchecks provides a mock function with given fields:
func (_m *Client) GetFailedHealthchecks() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetMailAsNotifications provides a mock function with given fields: ctx, seek, limit
func (_m *Client) GetMailAsNotifications(ctx context.Context, seek string, limit int) ([]*domain.Notification, error) {
	ret := _m.Called(ctx, seek, limit)

	var r0 []*domain.Notification
	if rf, ok := ret.Get(0).(func(context.Context, string, int) []*domain.Notification); ok {
		r0 = rf(ctx, seek, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Notification)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int) error); ok {
		r1 = rf(ctx, seek, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetModel provides a mock function with given fields:
func (_m *Client) GetModel() model.Model {
	ret := _m.Called()

	var r0 model.Model
	if rf, ok := ret.Get(0).(func() model.Model); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Model)
		}
	}

	return r0
}

// GetPathAccessRoles provides a mock function with given fields: ctx, b, path
func (_m *Client) GetPathAccessRoles(ctx context.Context, b textile.Bucket, path string) ([]domain.Member, error) {
	ret := _m.Called(ctx, b, path)

	var r0 []domain.Member
	if rf, ok := ret.Get(0).(func(context.Context, textile.Bucket, string) []domain.Member); ok {
		r0 = rf(ctx, b, path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Member)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, textile.Bucket, string) error); ok {
		r1 = rf(ctx, b, path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPublicShareBucket provides a mock function with given fields: ctx
func (_m *Client) GetPublicShareBucket(ctx context.Context) (textile.Bucket, error) {
	ret := _m.Called(ctx)

	var r0 textile.Bucket
	if rf, ok := ret.Get(0).(func(context.Context) textile.Bucket); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(textile.Bucket)
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

// GetReceivedFiles provides a mock function with given fields: ctx, accepted, seek, limit
func (_m *Client) GetReceivedFiles(ctx context.Context, accepted bool, seek string, limit int) ([]*domain.SharedDirEntry, string, error) {
	ret := _m.Called(ctx, accepted, seek, limit)

	var r0 []*domain.SharedDirEntry
	if rf, ok := ret.Get(0).(func(context.Context, bool, string, int) []*domain.SharedDirEntry); ok {
		r0 = rf(ctx, accepted, seek, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.SharedDirEntry)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, bool, string, int) string); ok {
		r1 = rf(ctx, accepted, seek, limit)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, bool, string, int) error); ok {
		r2 = rf(ctx, accepted, seek, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetThreadsConnection provides a mock function with given fields:
func (_m *Client) GetThreadsConnection() (*client.Client, error) {
	ret := _m.Called()

	var r0 *client.Client
	if rf, ok := ret.Get(0).(func() *client.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Client)
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

// IsHealthy provides a mock function with given fields:
func (_m *Client) IsHealthy() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsInitialized provides a mock function with given fields:
func (_m *Client) IsInitialized() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsRunning provides a mock function with given fields:
func (_m *Client) IsRunning() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// JoinBucket provides a mock function with given fields: ctx, slug, ti
func (_m *Client) JoinBucket(ctx context.Context, slug string, ti *domain.ThreadInfo) (bool, error) {
	ret := _m.Called(ctx, slug, ti)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, *domain.ThreadInfo) bool); ok {
		r0 = rf(ctx, slug, ti)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *domain.ThreadInfo) error); ok {
		r1 = rf(ctx, slug, ti)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBuckets provides a mock function with given fields: ctx
func (_m *Client) ListBuckets(ctx context.Context) ([]textile.Bucket, error) {
	ret := _m.Called(ctx)

	var r0 []textile.Bucket
	if rf, ok := ret.Get(0).(func(context.Context) []textile.Bucket); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]textile.Bucket)
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

// RejectSharedFilesInvitation provides a mock function with given fields: ctx, invitation
func (_m *Client) RejectSharedFilesInvitation(ctx context.Context, invitation domain.Invitation) (domain.Invitation, error) {
	ret := _m.Called(ctx, invitation)

	var r0 domain.Invitation
	if rf, ok := ret.Get(0).(func(context.Context, domain.Invitation) domain.Invitation); ok {
		r0 = rf(ctx, invitation)
	} else {
		r0 = ret.Get(0).(domain.Invitation)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Invitation) error); ok {
		r1 = rf(ctx, invitation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveKeys provides a mock function with given fields:
func (_m *Client) RemoveKeys() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendMessage provides a mock function with given fields: ctx, recipient, body
func (_m *Client) SendMessage(ctx context.Context, recipient crypto.PubKey, body []byte) (*usersclient.Message, error) {
	ret := _m.Called(ctx, recipient, body)

	var r0 *usersclient.Message
	if rf, ok := ret.Get(0).(func(context.Context, crypto.PubKey, []byte) *usersclient.Message); ok {
		r0 = rf(ctx, recipient, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usersclient.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, crypto.PubKey, []byte) error); ok {
		r1 = rf(ctx, recipient, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShareBucket provides a mock function with given fields: ctx, bucketSlug
func (_m *Client) ShareBucket(ctx context.Context, bucketSlug string) (*db.Info, error) {
	ret := _m.Called(ctx, bucketSlug)

	var r0 *db.Info
	if rf, ok := ret.Get(0).(func(context.Context, string) *db.Info); ok {
		r0 = rf(ctx, bucketSlug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, bucketSlug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShareFilesViaPublicKey provides a mock function with given fields: ctx, paths, pubkeys, keys
func (_m *Client) ShareFilesViaPublicKey(ctx context.Context, paths []domain.FullPath, pubkeys []crypto.PubKey, keys [][]byte) error {
	ret := _m.Called(ctx, paths, pubkeys, keys)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []domain.FullPath, []crypto.PubKey, [][]byte) error); ok {
		r0 = rf(ctx, paths, pubkeys, keys)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Shutdown provides a mock function with given fields:
func (_m *Client) Shutdown() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: ctx, cfg
func (_m *Client) Start(ctx context.Context, cfg config.Config) error {
	ret := _m.Called(ctx, cfg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, config.Config) error); ok {
		r0 = rf(ctx, cfg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ToggleBucketBackup provides a mock function with given fields: ctx, bucketSlug, bucketBackup
func (_m *Client) ToggleBucketBackup(ctx context.Context, bucketSlug string, bucketBackup bool) (bool, error) {
	ret := _m.Called(ctx, bucketSlug, bucketBackup)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) bool); ok {
		r0 = rf(ctx, bucketSlug, bucketBackup)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, bucketSlug, bucketBackup)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitForHealthy provides a mock function with given fields:
func (_m *Client) WaitForHealthy() chan error {
	ret := _m.Called()

	var r0 chan error
	if rf, ok := ret.Get(0).(func() chan error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan error)
		}
	}

	return r0
}

// WaitForInitialized provides a mock function with given fields:
func (_m *Client) WaitForInitialized() chan bool {
	ret := _m.Called()

	var r0 chan bool
	if rf, ok := ret.Get(0).(func() chan bool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan bool)
		}
	}

	return r0
}

// WaitForReady provides a mock function with given fields:
func (_m *Client) WaitForReady() chan bool {
	ret := _m.Called()

	var r0 chan bool
	if rf, ok := ret.Get(0).(func() chan bool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan bool)
		}
	}

	return r0
}
