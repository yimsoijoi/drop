package repository

import (
	"context"

	errorLibrary "github.com/yimsoijoi/drop/lib/error"
)

func (repo redisRepo) GetText(ctx context.Context, key *string) (*string, error) {
	if key != nil {
		val := repo.client.Get(ctx, *key)
		if val.Err() != nil {
			return returnErrGetText(val.Err())
		}
		valStr := val.Val()
		return &valStr, nil
	}
	return returnErrGetText(errorLibrary.ErrRedisNotFound)
}

func returnErrGetText(err error) (*string, error) {
	return nil, errorLibrary.NewError(errorLibrary.ErrRedisGetText.Error(), err)
}
