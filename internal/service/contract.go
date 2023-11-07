//go:generate mockgen -source ${GOFILE} -destination mocks_test.go -package ${GOPACKAGE}_test
package service

import "context"

type repo interface {
	GetURL(ctx context.Context, hash string) (string, error)
	SaveHashByURL(ctx context.Context, url, hash string) error
	Clear(ctx context.Context)
}

type generator interface {
	GenerateHash() string
}
