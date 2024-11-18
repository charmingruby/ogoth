package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/charmingruby/ogoth/config"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Warn("CONFIGURATION: .env file not found")
	}

	cfg, err := config.New()
	if err != nil {
		slog.Error(fmt.Sprintf("CONFIGURATION: %v", err))
		os.Exit(1)
	}

	slog.Info(fmt.Sprintf("CONFIGURATION: %+v", cfg))
}
