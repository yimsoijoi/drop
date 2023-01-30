package env

import (
	"github.com/spf13/viper"
	"log"
)

const (
	GinPort       = "GIN_PORT"
	AesCipherKey  = "AES_KEY"
	RedisPort     = "REDIS_PORT"
	RedisPassword = "REDIS_PASSWORD"
)

var commonEnvList = []string{
	GinPort,
	AesCipherKey,
	RedisPort,
	RedisPassword,
}

func GetCommonEnvList() []string {
	return commonEnvList
}

func BindingEnvironment(envList []string) {
	for i := 0; i < len(envList); i++ {
		if err := viper.BindEnv(envList[i]); err != nil {
			log.Fatalf("%s: %v", "failed to bind env", err)
		}
	}
}
