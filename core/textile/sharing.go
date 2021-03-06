package textile

import (
	"context"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/FleekHQ/space-daemon/core/space/domain"
	"github.com/FleekHQ/space-daemon/core/util/address"
	"github.com/FleekHQ/space-daemon/log"
	crypto "github.com/libp2p/go-libp2p-crypto"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/textile/v2/buckets"
)

func (tc *textileClient) ShareFilesViaPublicKey(ctx context.Context, paths []domain.FullPath, pubkeys []crypto.PubKey, keys [][]byte) error {
	var err error
	ctx, err = tc.getHubCtx(ctx)
	if err != nil {
		return err
	}

	for i, pth := range paths {
		ctx, _, err = tc.getBucketContext(ctx, pth.DbId, pth.Bucket, true, keys[i])
		if err != nil {
			return err
		}

		log.Info("Adding roles for pth: " + pth.Path)
		roles := make(map[string]buckets.Role)
		for _, pk := range pubkeys {
			tpk := thread.NewLibp2pPubKey(pk)
			// NOTE: setting to admin because receiving user
			// should be able to see members and reshare
			// as well
			roles[tpk.String()] = buckets.Admin
		}

		sbc := NewSecureBucketsClient(tc.hb, pth.Bucket)

		err := sbc.PushPathAccessRoles(ctx, pth.BucketKey, pth.Path, roles)
		if err != nil {
			return err
		}
	}

	return nil
}

var errInvitationNotPending = errors.New("invitation is no more pending")
var errInvitationAlreadyAccepted = errors.New("invitation is already accepted")
var errInvitationAlreadyRejected = errors.New("invitation is already rejected")

func (tc *textileClient) AcceptSharedFilesInvitation(
	ctx context.Context,
	invitation domain.Invitation,
) (domain.Invitation, error) {
	if invitation.Status == domain.ACCEPTED {
		return domain.Invitation{}, errInvitationAlreadyAccepted
	}

	if invitation.Status != domain.PENDING {
		return domain.Invitation{}, errInvitationNotPending
	}

	err := tc.createReceivedFiles(ctx, invitation, true)
	if err != nil {
		return domain.Invitation{}, err
	}
	invitation.Status = domain.ACCEPTED

	return invitation, nil
}

func (tc *textileClient) RejectSharedFilesInvitation(
	ctx context.Context,
	invitation domain.Invitation,
) (domain.Invitation, error) {
	if invitation.Status == domain.REJECTED {
		return domain.Invitation{}, errInvitationAlreadyRejected
	}

	if invitation.Status != domain.PENDING {
		return domain.Invitation{}, errInvitationNotPending
	}

	err := tc.createReceivedFiles(ctx, invitation, false)
	if err != nil {
		return domain.Invitation{}, err
	}
	invitation.Status = domain.REJECTED

	return invitation, nil
}

func (tc *textileClient) createReceivedFiles(
	ctx context.Context,
	invitation domain.Invitation,
	accepted bool,
) error {
	if len(invitation.ItemPaths) != len(invitation.Keys) {
		return errors.New("size of encryption keys does not match all items shared")
	}

	// TODO: Make this is call a transaction on threads so any failure can be easily reverted

	var allErr error
	for i, path := range invitation.ItemPaths {
		encryptionKeys := []byte("")
		if accepted {
			encryptionKeys = invitation.Keys[i]
		}
		_, err := tc.GetModel().CreateReceivedFile(ctx, path, invitation.InvitationID, accepted, encryptionKeys)

		// compose each create error
		if err != nil {
			if allErr == nil {
				allErr = errors.Wrap(err, "Failed to accept some invitations")
			}
			allErr = errors.Wrap(err, allErr.Error())
		}
	}

	return allErr
}

func (tc *textileClient) GetReceivedFiles(ctx context.Context, accepted bool, seek string, limit int) ([]*domain.SharedDirEntry, string, error) {
	files, err := tc.GetModel().ListReceivedFiles(ctx, accepted, seek, limit)
	if err != nil {
		return nil, "", err
	}

	items := []*domain.SharedDirEntry{}

	if len(files) == 0 {
		return items, "", nil
	}

	for _, file := range files {
		ctx, _, err = tc.getBucketContext(ctx, file.DbID, file.Bucket, true, file.EncryptionKey)
		if err != nil {
			return nil, "", err
		}

		sbc := NewSecureBucketsClient(tc.hb, file.Bucket)

		f, err := sbc.ListPath(ctx, file.BucketKey, file.Path)
		if err != nil {
			return nil, "", err
		}

		ipfsHash := f.Item.Cid
		name := f.Item.Name
		isDir := false
		size := f.GetItem().Size
		ext := strings.Replace(filepath.Ext(name), ".", "", -1)

		rs, err := sbc.PullPathAccessRoles(ctx, file.BucketKey, file.Path)
		if err != nil {
			// TEMP: returning empty members list until we
			// fix it on textile side
			//return nil, "", err
			rs = make(map[string]buckets.Role)
		}

		members := make([]domain.Member, 0)
		for pubk, _ := range rs {
			key := &thread.Libp2pPubKey{}

			err = key.UnmarshalString(pubk)
			if err != nil {
				log.Error(fmt.Sprintf("key.UnmarshalString(pubk=%+v)", pubk), err)
				return nil, "", err
			}

			pk := key.PubKey

			b, err := pk.Raw()
			if err != nil {
				return nil, "", err
			}

			members = append(members, domain.Member{
				Address:   address.DeriveAddress(pk),
				PublicKey: hex.EncodeToString(b),
			})
		}

		res := &domain.SharedDirEntry{
			Bucket: file.Bucket,
			DbID:   file.DbID,
			FileInfo: domain.FileInfo{
				IpfsHash:         ipfsHash,
				LocallyAvailable: false,
				BackedUp:         true,

				// TODO: Reflect correct state when we add local updates syncing to remote
				BackupInProgress: false,

				DirEntry: domain.DirEntry{
					Path:          file.Path,
					IsDir:         isDir,
					Name:          name,
					SizeInBytes:   strconv.FormatInt(size, 10),
					FileExtension: ext,
					Created:       strconv.FormatInt(time.Unix(0, file.CreatedAt).Unix(), 10),
					Updated:       strconv.FormatInt(time.Unix(0, file.CreatedAt).Unix(), 10), // NOTE: there is no modified yet so using same as create
				},
			},
			Members: members,
		}

		items = append(items, res)
	}

	offset := files[len(files)-1].ID.String()

	return items, offset, nil
}

func (tc *textileClient) GetPathAccessRoles(ctx context.Context, b Bucket, path string) ([]domain.Member, error) {
	var err error
	var bucketSlug, bucketKey string

	bucketSlug = b.Slug()

	bucket, err := tc.GetModel().FindBucket(ctx, bucketSlug)
	if err != nil {
		return nil, err
	}

	bucketKey = bucket.RemoteBucketKey

	hubCtx, _, err := tc.getBucketContext(ctx, bucket.RemoteDbID, bucketSlug, true, bucket.EncryptionKey)
	if err != nil {
		return nil, err
	}

	sbc := NewSecureBucketsClient(tc.hb, bucketSlug)

	rs, err := sbc.PullPathAccessRoles(hubCtx, bucketKey, path)
	if err != nil {
		// log.Error(fmt.Sprintf("PullPathAccessRoles not resolved (bucketKey=%s bucketSlug=%s path=%s)", bucketKey, bucketSlug, path), err)
		return []domain.Member{}, nil
	}

	// log.Debug(fmt.Sprintf("PullPathAccessRoles roles=%+v", rs))

	members := make([]domain.Member, 0)
	for pubk, _ := range rs {
		key := &thread.Libp2pPubKey{}

		err = key.UnmarshalString(pubk)
		if err != nil {
			log.Error(fmt.Sprintf("key.UnmarshalString(pubk=%+v)", pubk), err)
			return nil, err
		}

		pk := key.PubKey

		b, err := pk.Raw()
		if err != nil {
			return nil, err
		}

		members = append(members, domain.Member{
			Address:   address.DeriveAddress(pk),
			PublicKey: hex.EncodeToString(b),
		})
	}

	return members, nil
}

// return true if file was shared
// XXX: export this func?
func (tc *textileClient) isSharedFile(ctx context.Context, bucket Bucket, path string) bool {
	sbc := NewSecureBucketsClient(tc.hb, bucket.Slug())

	roles, err := sbc.PullPathAccessRoles(ctx, bucket.Key(), path)
	if err != nil {
		return false
	}

	pk, err := tc.kc.GetStoredPublicKey()
	if err != nil {
		return false
	}

	tpk := thread.NewLibp2pPubKey(pk)

	// shared means other roles than the user
	delete(roles, tpk.String())

	return len(roles) > 0
}
