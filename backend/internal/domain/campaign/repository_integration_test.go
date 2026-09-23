//go:build integration

package campaign

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestCampaignMigrations_ApplyConstraintsDownAndReUp(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	pool, databaseURL := isolatedCampaignPostgres(t, ctx)
	defer pool.Close()

	migrationsURL := "file://" + filepath.ToSlash(filepath.Join(repositoryRoot(t), "migrations"))
	migrationDatabaseURL := "pgx5" + databaseURL[len("postgres"):]
	m, err := migrate.New(migrationsURL, migrationDatabaseURL)
	if err != nil {
		t.Fatalf("create isolated migration runner: %v", err)
	}
	t.Cleanup(func() { _, _ = m.Close() })
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		t.Fatalf("fresh migration apply: %v", err)
	}

	assertCampaignConstraints(t, ctx, pool)
	if err := m.Steps(-1); err != nil {
		t.Fatalf("down Campaign migration: %v", err)
	}
	var exists bool
	if err := pool.QueryRow(ctx, "SELECT to_regclass('public.campaign_media') IS NOT NULL").Scan(&exists); err != nil {
		t.Fatalf("check child removal after down: %v", err)
	}
	if exists {
		t.Fatal("campaign_media remains after down migration")
	}
	if err := m.Up(); err != nil {
		t.Fatalf("re-apply Campaign migration: %v", err)
	}
}

func isolatedCampaignPostgres(t *testing.T, ctx context.Context) (*pgxpool.Pool, string) {
	t.Helper()
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:16-alpine",
			ExposedPorts: []string{"5432/tcp"},
			Env:          map[string]string{"POSTGRES_USER": "kencleng", "POSTGRES_PASSWORD": "kencleng", "POSTGRES_DB": "kencleng"},
			WaitingFor:   wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(time.Minute),
		},
		Started: true,
	})
	if err != nil {
		t.Fatalf("start isolated Postgres through the configured Docker-compatible runtime: %v", err)
	}
	t.Cleanup(func() {
		if err := container.Terminate(context.Background()); err != nil {
			t.Errorf("terminate isolated Postgres: %v", err)
		}
	})
	host, err := container.Host(ctx)
	if err != nil {
		t.Fatalf("resolve isolated Postgres host: %v", err)
	}
	port, err := container.MappedPort(ctx, "5432/tcp")
	if err != nil {
		t.Fatalf("resolve isolated Postgres port: %v", err)
	}
	databaseURL := fmt.Sprintf("postgres://kencleng:kencleng@%s:%s/kencleng?sslmode=disable", host, port.Port())
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		t.Fatalf("open isolated Postgres pool: %v", err)
	}
	return pool, databaseURL
}

func assertCampaignConstraints(t *testing.T, ctx context.Context, pool *pgxpool.Pool) {
	t.Helper()
	orgID, campaignID, mediaID := uuid.New(), uuid.New(), uuid.New()
	if _, err := pool.Exec(ctx, "INSERT INTO organizations (id, name) VALUES ($1, $2)", orgID, "Steward"); err != nil {
		t.Fatalf("insert organization: %v", err)
	}
	if _, err := pool.Exec(ctx, "INSERT INTO campaigns (id, organization_id, title, purpose, story, status, published_at, fundraising_ends_at, target_amount, collected_amount, media_state) VALUES ($1, $2, 'title', 'purpose', 'story', 'published', now(), now() + interval '1 day', 0, 0, 'available')", campaignID, orgID); err != nil {
		t.Fatalf("insert valid campaign: %v", err)
	}
	if _, err := pool.Exec(ctx, "INSERT INTO campaign_media (id, campaign_id, object_key, display_order, content_type, alt_text) VALUES ($1, $2, 'opaque-key', 0, 'image/jpeg', 'alt')", mediaID, campaignID); err != nil {
		t.Fatalf("insert valid media: %v", err)
	}
	constraints := []struct {
		name string
		args []any
	}{
		{name: "foreign key", args: []any{uuid.New(), uuid.Nil, "1", "1"}},
		{name: "non-negative amount", args: []any{uuid.New(), orgID, "-1", "1"}},
		{name: "nullable funding pair", args: []any{uuid.New(), orgID, "1", nil}},
	}
	const invalidCampaignSQL = "INSERT INTO campaigns (id, organization_id, title, purpose, story, status, fundraising_ends_at, target_amount, collected_amount, media_state) VALUES ($1, $2, 'title', 'purpose', 'story', 'draft', now(), $3, $4, 'absent')"
	for _, constraint := range constraints {
		if _, err := pool.Exec(ctx, invalidCampaignSQL, constraint.args...); err == nil {
			t.Fatalf("%s constraint unexpectedly allowed insert", constraint.name)
		}
	}
}

func repositoryRoot(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("locate repository root")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(file), "../../.."))
}
