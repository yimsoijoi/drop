package repository

import (
	"github.com/redis/go-redis/v9"
	"github.com/yimsoijoi/drop/internal/domain"
)

type redisRepo struct {
	client *redis.Client
}

func NewRedisRepo(client *redis.Client) domain.RedisRepo {
	return redisRepo{
		client: client,
	}
}
