package distributed_work

import (
	"context"
)

type Producer interface {
	RunOnce(ctx context.Context, testMode bool) error
}
