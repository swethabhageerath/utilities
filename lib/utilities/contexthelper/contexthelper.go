package contexthelper

import "context"

type ContextHelper struct {
}

func New() ContextHelper {
	return ContextHelper{}
}

func (c ContextHelper) Get(ctx context.Context, key string) string {
	if key == "" || ctx == nil {
		panic("ContextHelper.Get() -> Required parameters are not passing in...")
	}

	return ctx.Value(key).(string)
}
