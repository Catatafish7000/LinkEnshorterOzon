package controller_test

import (
	. "LinkEnshorter/internal/controller"
	"LinkEnshorter/pb"
	"context"
	"database/sql"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestServer_ShowURL(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		in       pb.Hash
		prepare  func(mock *Mockservice)
		expected func(t assert.TestingT, res string, err error)
	}{
		{
			name: "Success",
			in:   pb.Hash{Hash: "1234567890"},
			prepare: func(mock *Mockservice) {
				mock.EXPECT().ShowLink(gomock.Any(), "1234567890").Return("link", nil)
			},
			expected: func(t assert.TestingT, res string, err error) {
				assert.Equal(t, res, "link")
				assert.NoError(t, err)
			},
		},
		{
			name: "SuccessError",
			in:   pb.Hash{Hash: "1234567890"},
			prepare: func(mock *Mockservice) {
				mock.EXPECT().ShowLink(gomock.Any(), "1234567890").Return("", sql.ErrNoRows)
			},
			expected: func(t assert.TestingT, res string, err error) {
				assert.Equal(t, res, "")
				assert.Equal(t, err.Error(), "rpc error: code = InvalidArgument desc = sql: no rows in result set")
			},
		},
		{
			name: "Failed",
			in:   pb.Hash{Hash: "zfas"},
			prepare: func(mock *Mockservice) {
				mock.EXPECT().ShowLink(gomock.Any(), "zfas").Return("", errors.New("afgasfasa"))
			},
			expected: func(t assert.TestingT, res string, err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)

			service := NewMockservice(ctrl)
			tc.prepare(service)

			obj := NewServer(service)
			ans, err := obj.ShowURL(context.Background(), &tc.in)
			tc.expected(t, ans.Url, err)
		})
	}
}

func TestServer_SaveURL(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		in       pb.Url
		prepare  func(mock *Mockservice)
		expected func(t assert.TestingT, res string, err error)
	}{
		{
			name: "Success",
			in:   pb.Url{Url: "lol"},
			prepare: func(mock *Mockservice) {
				mock.EXPECT().SaveShortURL(gomock.Any(), "lol").Return("hash", nil)
			},
			expected: func(t assert.TestingT, res string, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "Failed",
			in:   pb.Url{Url: "zfas"},
			prepare: func(mock *Mockservice) {
				mock.EXPECT().SaveShortURL(gomock.Any(), "zfas").Return("", status.Error(codes.Internal, "wrong"))
			},
			expected: func(t assert.TestingT, res string, err error) {
				assert.Equal(t, err.Error(), "rpc error: code = Internal desc = rpc error: code = Internal desc = wrong")
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)

			service := NewMockservice(ctrl)
			tc.prepare(service)

			obj := NewServer(service)
			ans, err := obj.SaveURL(context.Background(), &tc.in)
			tc.expected(t, ans.Hash, err)
		})
	}
}
