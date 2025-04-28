package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	LocalMode  = "local"
	DockerMode = "docker"
)

const (
	LocalEnvPath  = "./config/.env"
	DockerEnvPath = "./.env"
)

type EnvMode string

func LoadEnv() {
	envMode := getEnvMode()
	envFilePath := LocalEnvPath

	if envMode == DockerMode {
		envFilePath = DockerEnvPath
	}

	if err := godotenv.Load(envFilePath); err != nil {
		log.Println("No env file found at", envFilePath, ", relying on system environment variables")
	}
}

func getEnvMode() EnvMode {
	mode := os.Getenv("ENVIRONMENT")
	return EnvMode(mode)
}

func GetDBUrl() (url string) {
	envMode := getEnvMode()
	if envMode == DockerMode {
		url = os.Getenv("DATABASE_URL")
	} else if envMode == LocalMode {
		url = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PROT"),
			os.Getenv("POSTGRES_DB"),
		)
	} else {
		log.Fatalln("Env DATABASE_URL invalid")
	}
	return
}

func GetRedisAddr() (addr string) {
	envMode := getEnvMode()
	if envMode == DockerMode {
		addr = os.Getenv("REDIS_ADDR")
	} else if envMode == LocalMode {
		addr = "localhost:6379"
	} else {
		log.Fatalln("Env Redis address invalid")
	}
	return
}
