// Command server is the Kencleng backend entry point.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"github.com/anhsbolic/kencleng/backend/internal/platform/auth"
	"github.com/anhsbolic/kencleng/backend/internal/platform/db"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("server: %v", err)
	}
}

func run() error {
	if err := loadEnvironment(); err != nil {
		return err
	}

	ctx := context.Background()
	pool, err := db.Open(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer pool.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthz)
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}
	browser := auth.BrowserConfig{Origin: os.Getenv("APP_ORIGIN"), Development: appEnv == "development"}
	if err := browser.Validate(); err != nil {
		return fmt.Errorf("configure authentication: %w", err)
	}
	providerCtx, cancelProvider := context.WithTimeout(ctx, 10*time.Second)
	protocol, err := auth.NewGoogleProtocol(providerCtx, os.Getenv("GOOGLE_CLIENT_ID"),
		os.Getenv("GOOGLE_CLIENT_SECRET"), browser.Origin+"/api/auth/google/callback")
	cancelProvider()
	if err != nil {
		return fmt.Errorf("initialize Google sign-in: %w", err)
	}
	authHandler, err := auth.NewHandler(auth.Store{Pool: pool}, protocol, browser)
	if err != nil {
		return fmt.Errorf("configure authentication: %w", err)
	}
	authHandler.Register(mux)

	srv := &http.Server{
		Addr:              ":" + os.Getenv("APP_PORT"),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	notifyCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-notifyCtx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("shutdown server: %w", err)
		}
		return nil
	}
}

func loadEnvironment() error {
	loadErr := godotenv.Load()
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}
	if loadErr != nil && appEnv == "development" {
		return fmt.Errorf("load .env: %w", loadErr)
	}
	if os.Getenv("APP_PORT") == "" {
		return fmt.Errorf("APP_PORT is required")
	}
	if os.Getenv("DATABASE_URL") == "" {
		return fmt.Errorf("DATABASE_URL is required")
	}
	if appEnv != "development" && appEnv != "production" {
		return fmt.Errorf("APP_ENV must be development or production")
	}
	if os.Getenv("APP_ORIGIN") == "" || os.Getenv("GOOGLE_CLIENT_ID") == "" || os.Getenv("GOOGLE_CLIENT_SECRET") == "" {
		return fmt.Errorf("APP_ORIGIN and Google OIDC client configuration are required")
	}
	return nil
}

func healthz(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
