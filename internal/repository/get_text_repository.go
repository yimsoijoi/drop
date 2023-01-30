package repository

import (
	"context"
	error2 "github.com/yimsoijoi/drop/lib/error"
)

func (repo redisRepo) GetText(ctx context.Context, key *string) (*string, error) {
	if key != nil {
		val := repo.client.Get(ctx, *key)
		valStr := val.Val()
		return &valStr, nil
	}
	return nil, error2.NewError(error2.ErrNotFoundInput.Error(), error2.ErrNotFoundInput)
}
