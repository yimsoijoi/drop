package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/yimsoijoi/drop/internal/delivery"
	"github.com/yimsoijoi/drop/internal/repository"
	"github.com/yimsoijoi/drop/internal/usecase"
	"log"
)

func main() {

	redisClient := initRedis()

	ginEngine := initGin()

	redisRepo := repository.NewRedisRepo(redisClient)

	usecase := usecase.NewUsecase(redisRepo)

	delivery.NewRestfulDelivery(ginEngine, usecase)

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
	if err := engine.Run("localhost:8000"); err != nil {
		return fmt.Errorf("failed listen and serve at port%s", "localhost:8000")
	}
	return nil
}
