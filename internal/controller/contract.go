//go:generate mockgen -source ${GOFILE} -destination mocks_test.go -package ${GOPACKAGE}_test
package controller

import "context"

type service interface {
	ShowLink(ctx context.Context, hash string) (string, error)
	SaveShortURL(ctx context.Context, url string) (string, error)
}
