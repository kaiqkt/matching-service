package di

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	var envs map[string]string
	envs, err := godotenv.Read("../../.env")

	if err != nil {
		log.Fatalf("Erro loading .env file, error: %s", err.Error())
	}

	return envs[key]
}
