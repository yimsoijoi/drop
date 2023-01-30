package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"github.com/yimsoijoi/drop/model"
)

func main() {
	fmt.Println(`this will be "drop" soon`)

}
func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
func postHandler(c *gin.Context) {
	model := new(model.CreateText)
	c.BindJSON(model)
}

func CreateText(m model.CreateText) (code *string, err error) {

	if m.Text != "" {
		var code string
		_ = m.Text[0:4]

		return &code, nil
	}

	return nil, errors.New("empty string")
}

type IRedisRepo interface {
	SetText(string, time.Time)
}

type redisRepo struct {
	client *redis.Client
}

func NewRedis(client *redis.Client) IRedisRepo {
	return redisRepo{
		client: client,
	}
}

func (rd redisRepo) SetText(text string, expTime time.Time) {
}
