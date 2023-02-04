package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/yimsoijoi/drop/internal/delivery/restful"
	"github.com/yimsoijoi/drop/internal/repository"
	"github.com/yimsoijoi/drop/internal/usecase"
)

func main() {

	redisClient := initRedis()

	ginEngine := initGin()

	redisRepo := repository.NewRedisRepo(redisClient)

	usecase := usecase.NewUsecase(redisRepo)

	restful.NewRestfulDelivery(ginEngine, usecase)

	if err := startServer(ginEngine); err != nil {
		log.Fatal(err)
	}

}
func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
func initGin() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	return engine
}

func startServer(engine *gin.Engine) error {
	//if err := engine.Run(viper.GetString(env.GinPort)); err != nil {
	//	return fmt.Errorf("failed listen and serve at port%s", viper.GetString(env.GinPort))
	//}
	localHost := "localhost:8000"
	if err := engine.Run(localHost); err != nil {
		return fmt.Errorf("%s:%w", localHost, err)
	}
	return nil
}
