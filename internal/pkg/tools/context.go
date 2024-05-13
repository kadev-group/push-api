package tools

import (
	"context"
	"sync"
)

const (
	key             = "context-holder-key"
	keyClientIP     = "client_ip"
	refreshTokenKey = "refreshToken"
)

func ContextHolderKey() string {
	return key
}

func SetAttribute(ctx context.Context, attribute string, value interface{}) {
	if contextHolder, ok := ctx.Value(key).(*sync.Map); ok {
		contextHolder.Store(attribute, value)
	}
}

func SetClientIP(ctx context.Context, ip string) {
	SetAttribute(ctx, keyClientIP, ip)
}

func GetClientIP(ctx context.Context) string {
	value := getAttribute(ctx, keyClientIP)
	if value != nil {
		if i, ok := value.(string); ok {
			return i
		}
	}
	return ""
}

func GetAttributeInt64(ctx context.Context, attribute string) int64 {
	value := getAttribute(ctx, attribute)
	if value != nil {
		if i, ok := value.(int64); ok {
			return i
		}
	}
	return 0
}

func getAttribute(ctx context.Context, attribute string) interface{} {
	if contextHolder, ok := ctx.Value(key).(*sync.Map); ok {
		value, ok := contextHolder.Load(attribute)
		if ok {
			return value
		}
	}
	return nil
}
