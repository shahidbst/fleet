// Automatically generated by mockimpl. DO NOT EDIT!

package mock

import (
	"context"
	"sync"
	"time"

	"github.com/micromdm/nanodep/client"
	"github.com/micromdm/nanodep/storage"
)

var _ storage.AllStorage = (*Storage)(nil)

type RetrieveAuthTokensFunc func(ctx context.Context, name string) (*client.OAuth1Tokens, error)

type RetrieveConfigFunc func(p0 context.Context, p1 string) (*client.Config, error)

type RetrieveAssignerProfileFunc func(ctx context.Context, name string) (profileUUID string, modTime time.Time, err error)

type RetrieveCursorFunc func(ctx context.Context, name string) (cursor string, modTime time.Time, err error)

type StoreCursorFunc func(ctx context.Context, name string, cursor string) error

type StoreAuthTokensFunc func(ctx context.Context, name string, tokens *client.OAuth1Tokens) error

type StoreConfigFunc func(ctx context.Context, name string, config *client.Config) error

type StoreTokenPKIFunc func(ctx context.Context, name string, pemCert []byte, pemKey []byte) error

type RetrieveTokenPKIFunc func(ctx context.Context, name string) (pemCert []byte, pemKey []byte, err error)

type StoreAssignerProfileFunc func(ctx context.Context, name string, profileUUID string) error

type Storage struct {
	RetrieveAuthTokensFunc        RetrieveAuthTokensFunc
	RetrieveAuthTokensFuncInvoked bool

	RetrieveConfigFunc        RetrieveConfigFunc
	RetrieveConfigFuncInvoked bool

	RetrieveAssignerProfileFunc        RetrieveAssignerProfileFunc
	RetrieveAssignerProfileFuncInvoked bool

	RetrieveCursorFunc        RetrieveCursorFunc
	RetrieveCursorFuncInvoked bool

	StoreCursorFunc        StoreCursorFunc
	StoreCursorFuncInvoked bool

	StoreAuthTokensFunc        StoreAuthTokensFunc
	StoreAuthTokensFuncInvoked bool

	StoreConfigFunc        StoreConfigFunc
	StoreConfigFuncInvoked bool

	StoreTokenPKIFunc        StoreTokenPKIFunc
	StoreTokenPKIFuncInvoked bool

	RetrieveTokenPKIFunc        RetrieveTokenPKIFunc
	RetrieveTokenPKIFuncInvoked bool

	StoreAssignerProfileFunc        StoreAssignerProfileFunc
	StoreAssignerProfileFuncInvoked bool

	mu sync.Mutex
}

func (s *Storage) RetrieveAuthTokens(ctx context.Context, name string) (*client.OAuth1Tokens, error) {
	s.mu.Lock()
	s.RetrieveAuthTokensFuncInvoked = true
	s.mu.Unlock()
	return s.RetrieveAuthTokensFunc(ctx, name)
}

func (s *Storage) RetrieveConfig(p0 context.Context, p1 string) (*client.Config, error) {
	s.mu.Lock()
	s.RetrieveConfigFuncInvoked = true
	s.mu.Unlock()
	return s.RetrieveConfigFunc(p0, p1)
}

func (s *Storage) RetrieveAssignerProfile(ctx context.Context, name string) (profileUUID string, modTime time.Time, err error) {
	s.mu.Lock()
	s.RetrieveAssignerProfileFuncInvoked = true
	s.mu.Unlock()
	return s.RetrieveAssignerProfileFunc(ctx, name)
}

func (s *Storage) RetrieveCursor(ctx context.Context, name string) (cursor string, modTime time.Time, err error) {
	s.mu.Lock()
	s.RetrieveCursorFuncInvoked = true
	s.mu.Unlock()
	return s.RetrieveCursorFunc(ctx, name)
}

func (s *Storage) StoreCursor(ctx context.Context, name string, cursor string) error {
	s.mu.Lock()
	s.StoreCursorFuncInvoked = true
	s.mu.Unlock()
	return s.StoreCursorFunc(ctx, name, cursor)
}

func (s *Storage) StoreAuthTokens(ctx context.Context, name string, tokens *client.OAuth1Tokens) error {
	s.mu.Lock()
	s.StoreAuthTokensFuncInvoked = true
	s.mu.Unlock()
	return s.StoreAuthTokensFunc(ctx, name, tokens)
}

func (s *Storage) StoreConfig(ctx context.Context, name string, config *client.Config) error {
	s.mu.Lock()
	s.StoreConfigFuncInvoked = true
	s.mu.Unlock()
	return s.StoreConfigFunc(ctx, name, config)
}

func (s *Storage) StoreTokenPKI(ctx context.Context, name string, pemCert []byte, pemKey []byte) error {
	s.mu.Lock()
	s.StoreTokenPKIFuncInvoked = true
	s.mu.Unlock()
	return s.StoreTokenPKIFunc(ctx, name, pemCert, pemKey)
}

func (s *Storage) RetrieveTokenPKI(ctx context.Context, name string) (pemCert []byte, pemKey []byte, err error) {
	s.mu.Lock()
	s.RetrieveTokenPKIFuncInvoked = true
	s.mu.Unlock()
	return s.RetrieveTokenPKIFunc(ctx, name)
}

func (s *Storage) StoreAssignerProfile(ctx context.Context, name string, profileUUID string) error {
	s.mu.Lock()
	s.StoreAssignerProfileFuncInvoked = true
	s.mu.Unlock()
	return s.StoreAssignerProfileFunc(ctx, name, profileUUID)
}