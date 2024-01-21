package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Host string
	Port string
	PostgresHost string
	PostgresPort string
	PostgresPassword string
	PostgresDb string
	PostgresUser string
}

func Load() Config {
	var cfg Config
	err := godotenv.Load("./.env")

	if err != nil {
		fmt.Println("This .env file was not found")
	}

	cfg.Host = cast.ToString(getEnvArg("HTTP_HOST", "localhost"))
	cfg.Port = cast.ToString(getEnvArg("HTTP_PORT", "5000"))

	cfg.PostgresHost = cast.ToString(getEnvArg("POSTGRES_HOST", "localhost"))
	cfg.PostgresPort = cast.ToString(getEnvArg("POSTGRES_PORT", "5434"))
	cfg.PostgresUser = cast.ToString(getEnvArg("POSTGRES_USER", "postgres"))
	cfg.PostgresDb = cast.ToString(getEnvArg("POSTGRES_DATABASE", "golang_crud"))
	cfg.PostgresPassword = cast.ToString(getEnvArg("POSTGRES_PASSWORD", "postgres"))

	return cfg
}

func getEnvArg(key string, defaultValue interface{}) interface{} {

	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}

	return defaultValue
}
