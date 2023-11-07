package strategy

import "context"

type repo interface {
	GetURL(ctx context.Context, url string) (string, error)
	SaveHashByURL(ctx context.Context, url, hash string) error
	Clear(ctx context.Context)
}
