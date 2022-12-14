package helpers

import "context"

type IContextHelper interface {
	Get(ctx context.Context, key string) (string, error)
}
