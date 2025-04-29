package config

import (
	"errors"
	"fmt"
	"os"

	logger "github.com/dev-rever/cryptoo-pricing/utils/logutils"
	"github.com/joho/godotenv"
)

type path string
type mode string

const (
	local  mode = "local"
	docker mode = "docker"
)

const (
	localEnv  path = "./config/.env"
	dockerEnv path = "./.env"
)

func LoadEnv() {
	var p path
	if envMode() == docker {
		p = dockerEnv
	} else {
		p = localEnv
	}

	if err := godotenv.Load(string(p)); err != nil {
		logger.LogError(errors.New(fmt.Sprintf("No env file found at %v, relying on system environment variables", p)))
		return
	}
}

func envMode() mode {
	env := os.Getenv("ENVIRONMENT")
	return mode(env)
}

func GetDBUrl() (url string) {
	envMode := envMode()
	if envMode == docker {
		url = os.Getenv("DATABASE_URL")
	} else {
		url = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PROT"),
			os.Getenv("POSTGRES_DB"),
		)
	}
	return
}

func GetRedisAddr() (addr string) {
	envMode := envMode()
	if envMode == docker {
		addr = os.Getenv("REDIS_ADDR")
	} else {
		addr = "localhost:6379"
	}
	return
}
