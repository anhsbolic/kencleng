// Command server is the Kencleng backend entry point.
//
// It is wiring only: it loads configuration, opens the shared platform
// dependencies (database, object storage, crypto keys, JWT key pair), and
// starts the HTTP server. No domain logic lives in this package.
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/anhsbolic/kencleng/backend/internal/platform/auth"
	platformcrypto "github.com/anhsbolic/kencleng/backend/internal/platform/crypto"
	"github.com/anhsbolic/kencleng/backend/internal/platform/db"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("server: %v", err)
	}
}

// run performs the startup sequence in order, failing fast on the first
// stage that cannot complete rather than continuing in a half-initialized
// state. See backend/.agents/docs/scaffold-backend.md Step 2.
func run() error {
	ctx := context.Background()

	// 1. Configuration: load .env from the working directory. Fatal when
	// missing in development; non-fatal otherwise (env vars are injected
	// externally in other environments).
	loadErr := godotenv.Load()
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}
	if loadErr != nil && appEnv == "development" {
		return fmt.Errorf("load .env: %w", loadErr)
	}

	// 2. Validate required env vars are present and well-formed.
	if err := requireEnv(
		"APP_PORT",
		"DATABASE_URL",
		"MINIO_ENDPOINT",
		"MINIO_ACCESS_KEY",
		"MINIO_SECRET_KEY",
		"MINIO_BUCKET_PUBLIC",
		"MINIO_BUCKET_PRIVATE",
		"JWT_PRIVATE_KEY_PATH",
		"JWT_PUBLIC_KEY_PATH",
	); err != nil {
		return err
	}
	keys, err := platformcrypto.New(os.Getenv("ENCRYPTION_KEY"), os.Getenv("HMAC_KEY"))
	if err != nil {
		return err
	}

	// 3. Database: open the pgx pool and ping before proceeding.
	pool, err := db.Open(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer pool.Close()

	// 4. Object storage: initialize MinIO and verify both buckets exist.
	if err := initMinIO(ctx); err != nil {
		return err
	}

	// 5. Auth: load the ES256 key pair.
	authKeys, err := auth.Load(os.Getenv("JWT_PRIVATE_KEY_PATH"), os.Getenv("JWT_PUBLIC_KEY_PATH"))
	if err != nil {
		return err
	}

	// The wired dependencies are intentionally unused until the first domain
	// task; keeping them referenced documents that they are part of startup.
	_, _ = keys, authKeys

	// 6. Router: a single health check until the first domain task adds real
	// endpoints.
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthz)

	srv := &http.Server{
		Addr:    ":" + os.Getenv("APP_PORT"),
		Handler: mux,
	}

	// 7. Graceful shutdown on SIGINT/SIGTERM.
	notifyCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	log.Printf("server listening on %s", srv.Addr)

	select {
	case err := <-errCh:
		return fmt.Errorf("listen: %w", err)
	case <-notifyCtx.Done():
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("graceful shutdown: %w", err)
	}
	return nil
}

// healthz reports that the process is up and listening.
func healthz(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte(`{"status":"ok"}`)); err != nil {
		log.Printf("healthz: write response: %v", err)
	}
}

// requireEnv returns an error naming the first of the given variables that
// is missing or empty.
func requireEnv(names ...string) error {
	for _, name := range names {
		if os.Getenv(name) == "" {
			return fmt.Errorf("required env var %s is empty", name)
		}
	}
	return nil
}

// initMinIO connects to the MinIO endpoint and verifies both buckets exist.
func initMinIO(ctx context.Context) error {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	useSSL := os.Getenv("MINIO_USE_SSL") == "true"

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), ""),
		Secure: useSSL,
	})
	if err != nil {
		return fmt.Errorf("create minio client: %w", err)
	}

	for _, bucket := range []string{os.Getenv("MINIO_BUCKET_PUBLIC"), os.Getenv("MINIO_BUCKET_PRIVATE")} {
		ok, err := client.BucketExists(ctx, bucket)
		if err != nil {
			return fmt.Errorf("check minio bucket %q: %w", bucket, err)
		}
		if !ok {
			return fmt.Errorf("minio bucket %q does not exist", bucket)
		}
	}
	return nil
}
