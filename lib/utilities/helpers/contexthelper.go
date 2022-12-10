package helpers

import (
	"context"
	e "github.com/pkg/errors"
)

type ContextHelper struct{}

func (c ContextHelper) Get(ctx context.Context, key string) (string, error) {
	if key == "" || ctx == nil {
		return "", e.New("ContextHelper.Get() -> Required parameters are not passing in...")
	}

	return ctx.Value(key).(string), nil
}
