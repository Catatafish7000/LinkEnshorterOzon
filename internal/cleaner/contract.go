package cleaner

import "context"

type repo interface {
	Clear(ctx context.Context)
}
