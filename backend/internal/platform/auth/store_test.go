package auth

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
)

type sessionRow struct {
	scan func(...any) error
}

func (r sessionRow) Scan(dest ...any) error { return r.scan(dest...) }

type sessionTx struct {
	pgx.Tx
	query func(context.Context, string, ...any) pgx.Row
}

func (tx sessionTx) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return tx.query(ctx, sql, args...)
}

func fixtureToken() string {
	return base64.RawURLEncoding.EncodeToString(bytes.Repeat([]byte{7}, 32))
}

// The Scan barrier simulates a row-lock wait; it is not PostgreSQL contention.
func TestLockSessionExpiredDuringLockWait(t *testing.T) {
	before := time.Date(2026, 10, 11, 0, 0, 0, 0, time.UTC)
	clock := before
	locked := false
	token := fixtureToken()
	tx := sessionTx{query: func(_ context.Context, sql string, args ...any) pgx.Row {
		if !strings.Contains(sql, "FOR UPDATE") || len(args) != 1 || !bytes.Equal(args[0].([]byte), digest(token)) {
			t.Fatal("session lookup must lock the row using the token digest")
		}
		return sessionRow{scan: func(dest ...any) error {
			*(dest[0].(*string)) = "person-from-locked-row"
			*(dest[1].(*string)) = "fixture-csrf"
			*(dest[2].(*time.Time)) = before.Add(time.Minute)
			*(dest[3].(**time.Time)) = nil
			clock = before.Add(2 * time.Minute)
			locked = true
			return nil
		}}
	}}
	session, err := lockSession(context.Background(), tx, token, func() time.Time {
		if !locked {
			t.Fatal("expiry clock sampled before row lock acquisition")
		}
		return clock
	})
	if !errors.Is(err, ErrUnauthenticated) || session != (Session{}) {
		t.Fatalf("expired during simulated lock wait: session=%v error=%v", session, err)
	}
}

func TestLockSessionRevalidation(t *testing.T) {
	now := time.Date(2026, 10, 11, 0, 0, 0, 0, time.UTC)
	storageErr := errors.New("fixture database unavailable")
	for _, tc := range []struct {
		name    string
		expires time.Time
		revoked *time.Time
		rowErr  error
		wantErr error
	}{
		{"valid", now.Add(time.Minute), nil, nil, nil},
		{"expiry boundary", now, nil, nil, ErrUnauthenticated},
		{"expired", now.Add(-time.Second), nil, nil, ErrUnauthenticated},
		{"revoked", now.Add(time.Minute), &now, nil, ErrUnauthenticated},
		{"missing", time.Time{}, nil, pgx.ErrNoRows, ErrUnauthenticated},
		{"query error", time.Time{}, nil, storageErr, storageErr},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tx := sessionTx{query: func(context.Context, string, ...any) pgx.Row {
				return sessionRow{scan: func(dest ...any) error {
					if tc.rowErr != nil {
						return tc.rowErr
					}
					*(dest[0].(*string)) = "locked-person"
					*(dest[1].(*string)) = "fixture-csrf"
					*(dest[2].(*time.Time)) = tc.expires
					*(dest[3].(**time.Time)) = tc.revoked
					return nil
				}}
			}}
			s, err := lockSession(context.Background(), tx, fixtureToken(), func() time.Time { return now })
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("error=%v, want %v", err, tc.wantErr)
			}
			if tc.wantErr != nil && s != (Session{}) {
				t.Fatal("failed validation returned session authority")
			}
			if tc.wantErr == nil && (s.PersonID != "locked-person" || s.CSRFToken != "fixture-csrf" || !s.ExpiresAt.Equal(tc.expires)) {
				t.Fatal("valid locked session attribution changed")
			}
		})
	}
}

func TestLockSessionIgnoresStaleCallerTime(t *testing.T) {
	caller := time.Now().Add(-time.Hour)
	tx := sessionTx{query: func(context.Context, string, ...any) pgx.Row {
		return sessionRow{scan: func(dest ...any) error {
			*(dest[0].(*string)) = "locked-person"
			*(dest[1].(*string)) = "fixture-csrf"
			*(dest[2].(*time.Time)) = time.Now().Add(-time.Minute)
			*(dest[3].(**time.Time)) = nil
			return nil
		}}
	}}
	if _, err := LockSession(context.Background(), tx, fixtureToken(), caller); !errors.Is(err, ErrUnauthenticated) {
		t.Fatalf("public lock API accepted expiry after stale caller time: %v", err)
	}
}

func TestMalformedBrowserAuthorityRejectedBeforeStorage(t *testing.T) {
	ctx := context.Background()
	// Nil storage intentionally proves malformed tokens never reach the database.
	if _, err := LockSession(ctx, nil, "invalid", time.Now()); !errors.Is(err, ErrUnauthenticated) {
		t.Fatal("malformed session accepted")
	}
	for _, pair := range [][2]string{{"invalid", fixtureToken()}, {fixtureToken(), "invalid"}, {"", ""}} {
		if _, err := (Store{}).ConsumeLogin(ctx, pair[0], pair[1], time.Now()); !errors.Is(err, ErrInvalidLogin) {
			t.Fatal("malformed state/browser token accepted")
		}
	}
}
