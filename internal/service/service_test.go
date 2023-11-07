package service_test

import (
	. "LinkEnshorter/internal/service"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_ShowLink(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		in       string
		prepare  func(mock *Mockrepo)
		expected func(t assert.TestingT, res string, err error)
	}{
		{
			name: "Success",
			in:   "1234567890",
			prepare: func(mock *Mockrepo) {
				mock.EXPECT().GetURL(gomock.Any(), "1234567890").Return("link", nil)
			},
			expected: func(t assert.TestingT, res string, err error) {
				assert.Equal(t, res, "link")
				assert.NoError(t, err)
			},
		},
		{
			name: "Failed",
			in:   "zfas",
			prepare: func(mock *Mockrepo) {
				mock.EXPECT().GetURL(gomock.Any(), "zfas").Return("", errors.New("afgasfasa"))
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

			repo := NewMockrepo(ctrl)
			gen := NewMockgenerator(ctrl)
			tc.prepare(repo)

			obj := NewService(repo, gen)
			ans, err := obj.ShowLink(context.Background(), tc.in)
			tc.expected(t, ans, err)
		})
	}
}
