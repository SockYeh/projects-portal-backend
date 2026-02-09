package logger

import (
	"os"

	"go.uber.org/zap"
)

func Init() (*zap.Logger, error) {
	env := os.Getenv("APP_ENV")
	if env == "production" {
		return zap.NewProduction()
	}

	return zap.NewDevelopment()
}
