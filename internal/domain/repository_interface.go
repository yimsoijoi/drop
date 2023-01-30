package domain

import "context"

type RedisRepo interface {
	SetText(ctx context.Context, key string, val string) error
	GetText(ctx context.Context, key *string) (*string, error)
}
