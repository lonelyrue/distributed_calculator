package auth

import (
	"context"
)

type contextKey string

var userIDKey = contextKey("userID")

func WithUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func GetUserID(ctx context.Context) int {
	if ctx == nil {
		return 0
	}
	if v, ok := ctx.Value(userIDKey).(int); ok {
		return v
	}
	return 0
}
