package repository

import (
	"context"

	"github.com/yimsoijoi/drop/internal/common"
	errorLibrary "github.com/yimsoijoi/drop/lib/error"
)

func (rd redisRepo) SetText(ctx context.Context, key string, val string) error {
	if err := rd.client.Set(ctx, key, val, common.RedisDefaultExpireTime).Err(); err != nil {
		return errorLibrary.NewError(errorLibrary.ErrRedisSetText.Error(), err)
	}
	return nil
}
