package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Open creates a pgx connection pool for the given database URL and verifies
// connectivity with a ping before returning it, so a bad connection fails
// fast at startup rather than surfacing later as a confusing query error.
func Open(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	if databaseURL == "" {
		return nil, fmt.Errorf("database URL is empty")
	}
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("create connection pool: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}
	return pool, nil
}
