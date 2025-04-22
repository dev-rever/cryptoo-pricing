package config

import (
	"log"

	"github.com/joho/godotenv"
)

const ENV_FILE_PATH = "./config/.env"

func LoadEnv() {
	if err := godotenv.Load(ENV_FILE_PATH); err != nil {
		log.Fatalln("Error loading .env file.")
	}
}
