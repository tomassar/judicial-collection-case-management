package initializers

import (
	"log/slog"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		slog.Error("Error loading .env file", "error", err)
	}
}
