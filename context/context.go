package context

import (
	"context"

	"s16-tech-test/config"
)

type ContextKey string

const (
	APIKey ContextKey = "api_key"
)

func IsAuthenticated(ctx context.Context) bool {
	key, ok := ctx.Value(APIKey).(string)
	if !ok {
		return false
	}
	return config.ConfigInUse.Server.APIKey == key
}
