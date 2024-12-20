package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/charmingruby/ogoth/config"
	"github.com/charmingruby/ogoth/internal/auth"
	"github.com/charmingruby/ogoth/internal/auth/transport/rest/client"
	"github.com/charmingruby/ogoth/internal/health"
	"github.com/charmingruby/ogoth/internal/shared/transport/rest"
	"github.com/charmingruby/ogoth/internal/shared/transport/rest/session"
	"github.com/go-chi/chi/v5"
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

	router := chi.NewRouter()
	server := rest.NewServer(cfg.ServerConfig.Port, router)
	serverAddr := fmt.Sprintf("%s:%s", cfg.ServerConfig.Host, cfg.ServerConfig.Port)

	sessionStore := session.New(cfg.ServerConfig.CookieSecretKey, cfg.Environment)

	healthRestHandler := health.NewRestHandler(router)
	healthRestHandler.Register()

	googleOAuth2Client := client.NewGoogleOAuth2(serverAddr, cfg.GoogleConfig.ClientID, cfg.GoogleConfig.ClientSecretID)
	authRestHandler := auth.NewRestHandler(router, *googleOAuth2Client, sessionStore)
	authRestHandler.Register()

	slog.Info(fmt.Sprintf("SERVER: Running on port %s", cfg.ServerConfig.Port))

	if err := server.Run(); err != nil {
		slog.Error(fmt.Sprintf("SERVER: %v", err))
		os.Exit(1)
	}
}
