//go:build integration

package auth

import (
	"context"
	"errors"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Only use a disposable test database. This does not load application .env.
func disposableIdentityPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	dsn := os.Getenv("C1_AUTH_TEST_DATABASE_URL")
	if dsn == "" {
		t.Skip("C1_AUTH_TEST_DATABASE_URL must identify a disposable fixture")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	admin, err := pgxpool.New(ctx, dsn)
	if err != nil {
		t.Fatal("create fixture pool")
	}
	name, err := newUUID()
	if err != nil {
		admin.Close()
		t.Fatal(err)
	}
	schema := "c1_auth_" + strings.ReplaceAll(name, "-", "")
	identifier := pgx.Identifier{schema}.Sanitize()
	if _, err := admin.Exec(ctx, "CREATE SCHEMA "+identifier); err != nil {
		admin.Close()
		t.Fatal("create disposable identity schema")
	}
	t.Cleanup(func() {
		cleanupCtx, stop := context.WithTimeout(context.Background(), 5*time.Second)
		defer stop()
		if _, err := admin.Exec(cleanupCtx, "DROP SCHEMA "+identifier+" CASCADE"); err != nil {
			t.Error("cleanup disposable identity schema failed")
		}
		admin.Close()
	})
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		t.Fatal("parse fixture configuration")
	}
	config.ConnConfig.RuntimeParams["search_path"] = schema
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		t.Fatal("create schema-bound fixture pool")
	}
	t.Cleanup(pool.Close)
	migration, err := os.ReadFile("../../../migrations/000001_identity_session.up.sql")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, string(migration)); err != nil {
		t.Fatal("apply identity migration to disposable schema")
	}
	return pool
}

func TestPersonUpsertPostgreSQL(t *testing.T) {
	pool := disposableIdentityPool(t)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	store := Store{Pool: pool}
	id, err := store.UpsertPerson(ctx, googleIssuer, "CaseSensitiveSubject")
	if err != nil || id == "" {
		t.Fatal("first verified-person upsert failed")
	}
	repeated, err := (Store{Pool: pool}).UpsertPerson(ctx, googleIssuer, "CaseSensitiveSubject")
	if err != nil || repeated != id {
		t.Fatal("same immutable issuer/subject did not retain person ID")
	}
	for _, pair := range [][2]string{
		{googleIssuer, "casesensitivesubject"},
		{"https://other.example", "CaseSensitiveSubject"},
		{googleIssuer, "subject'); DROP TABLE persons; --"},
	} {
		other, err := store.UpsertPerson(ctx, pair[0], pair[1])
		if err != nil || other == "" || other == id {
			t.Fatal("distinct issuer/subject lost case-sensitive identity or parameter safety")
		}
		var issuer, subject string
		if err := pool.QueryRow(ctx, "SELECT issuer,subject FROM persons WHERE id=$1", other).Scan(&issuer, &subject); err != nil || issuer != pair[0] || subject != pair[1] {
			t.Fatal("identity input was changed or interpreted as SQL")
		}
	}
	for _, pair := range [][2]string{{"", "subject"}, {googleIssuer, ""}} {
		if _, err := store.UpsertPerson(ctx, pair[0], pair[1]); err == nil {
			t.Fatal("empty identity accepted")
		}
	}
	var count int
	if err := pool.QueryRow(ctx, "SELECT count(*) FROM persons").Scan(&count); err != nil || count != 4 {
		t.Fatal("unexpected identity rows after upserts")
	}
}

func TestPendingLoginBindingAndReplayPostgreSQL(t *testing.T) {
	store := Store{Pool: disposableIdentityPool(t)}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	now := time.Now()
	state, browser, err := store.StartLogin(ctx, now, "fixture-nonce", "fixture-verifier")
	if err != nil {
		t.Fatal("create fixture pending login")
	}
	wrongBrowser, err := randomToken()
	if err != nil {
		t.Fatal(err)
	}
	wrongState, err := randomToken()
	if err != nil {
		t.Fatal(err)
	}
	for _, pair := range [][2]string{{state, wrongBrowser}, {wrongState, browser}} {
		if _, err := store.ConsumeLogin(ctx, pair[0], pair[1], now); !errors.Is(err, ErrInvalidLogin) {
			t.Fatal("mismatched state/browser accepted")
		}
	}
	pending, err := store.ConsumeLogin(ctx, state, browser, now)
	if err != nil || pending.Nonce != "fixture-nonce" || pending.Verifier != "fixture-verifier" {
		t.Fatal("correct binding lost nonce/PKCE verifier")
	}
	if _, err := store.ConsumeLogin(ctx, state, browser, now); !errors.Is(err, ErrInvalidLogin) {
		t.Fatal("consumed authorization flow accepted a replay")
	}
	state, browser, err = store.StartLogin(ctx, now.Add(-loginLifetime), "fixture-nonce", "fixture-verifier")
	if err != nil {
		t.Fatal("create expired fixture login")
	}
	if _, err := store.ConsumeLogin(ctx, state, browser, now); !errors.Is(err, ErrInvalidLogin) {
		t.Fatal("expired login accepted")
	}
}
