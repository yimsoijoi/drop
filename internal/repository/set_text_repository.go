package repository

import (
	"context"
	"fmt"
	"github.com/yimsoijoi/drop/internal/common"
	error2 "github.com/yimsoijoi/drop/lib/error"
)

func (rd redisRepo) SetText(ctx context.Context, key string, val string) error {
	if err := rd.client.Set(ctx, key, val, common.RedisDefaultExpireTime).Err(); err != nil {
		return fmt.Errorf("%s%w", error2.ErrRedisGetText, err)
	}
	return nil
}
