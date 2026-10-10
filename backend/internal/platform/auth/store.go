package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrUnauthenticated = errors.New("session is missing, expired, or revoked")
var ErrInvalidLogin = errors.New("pending login is missing, expired, or browser-mismatched")

const loginLifetime = 10 * time.Minute
const sessionLifetime = 8 * time.Hour

type Store struct{ Pool *pgxpool.Pool }

type PendingLogin struct {
	Nonce    string
	Verifier string
}

// Session is local authority only. PersonID is the server-owned UUID, not a
// provider subject. A caller must use LockSession inside a grant transaction.
type Session struct {
	PersonID  string
	CSRFToken string
	ExpiresAt time.Time
}

func randomToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("random token: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func digest(token string) []byte {
	sum := sha256.Sum256([]byte(token))
	return sum[:]
}

func validToken(token string) bool {
	b, err := base64.RawURLEncoding.DecodeString(token)
	return err == nil && len(b) == 32
}

func newUUID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("random UUID: %w", err)
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

func (s Store) StartLogin(ctx context.Context, now time.Time, nonce, verifier string) (state, browser string, err error) {
	state, err = randomToken()
	if err != nil {
		return "", "", err
	}
	browser, err = randomToken()
	if err != nil {
		return "", "", err
	}
	_, err = s.Pool.Exec(ctx, `INSERT INTO pending_logins (state_digest,browser_digest,nonce,pkce_verifier,expires_at)
		VALUES ($1,$2,$3,$4,$5)`, digest(state), digest(browser), nonce, verifier, now.Add(loginLifetime))
	if err != nil {
		return "", "", fmt.Errorf("insert pending login: %w", err)
	}
	return state, browser, nil
}

// ConsumeLogin atomically burns the state even for a provider-error callback.
func (s Store) ConsumeLogin(ctx context.Context, state, browser string, now time.Time) (PendingLogin, error) {
	var p PendingLogin
	if !validToken(state) || !validToken(browser) {
		return p, ErrInvalidLogin
	}
	err := s.Pool.QueryRow(ctx, `DELETE FROM pending_logins
		WHERE state_digest=$1 AND browser_digest=$2 AND expires_at>$3
		RETURNING nonce,pkce_verifier`, digest(state), digest(browser), now).Scan(&p.Nonce, &p.Verifier)
	if errors.Is(err, pgx.ErrNoRows) {
		return p, ErrInvalidLogin
	}
	if err != nil {
		return p, fmt.Errorf("consume pending login: %w", err)
	}
	return p, nil
}

func (s Store) UpsertPerson(ctx context.Context, issuer, subject string) (string, error) {
	if issuer == "" || subject == "" {
		return "", errors.New("verified issuer and subject are required")
	}
	id, err := newUUID()
	if err != nil {
		return "", err
	}
	var personID string
	err = s.Pool.QueryRow(ctx, `INSERT INTO persons(id,issuer,subject) VALUES ($1,$2,$3)
		ON CONFLICT (issuer,subject) DO UPDATE SET subject=EXCLUDED.subject
		RETURNING id`, id, issuer, subject).Scan(&personID)
	if err != nil {
		return "", fmt.Errorf("upsert verified person: %w", err)
	}
	return personID, nil
}

// LockSession rechecks revocation and expiry after acquiring the row lock.
// T3 must call it as its first lock in the confirmation transaction.
func LockSession(ctx context.Context, tx pgx.Tx, token string, now time.Time) (Session, error) {
	var session Session
	if !validToken(token) {
		return session, ErrUnauthenticated
	}
	var revokedAt *time.Time
	err := tx.QueryRow(ctx, `SELECT person_id::text,csrf_token,expires_at,revoked_at
		FROM sessions WHERE token_digest=$1 FOR UPDATE`, digest(token)).Scan(
		&session.PersonID, &session.CSRFToken, &session.ExpiresAt, &revokedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return Session{}, ErrUnauthenticated
	}
	if err != nil {
		return Session{}, fmt.Errorf("lock session: %w", err)
	}
	if revokedAt != nil || !now.Before(session.ExpiresAt) {
		return Session{}, ErrUnauthenticated
	}
	return session, nil
}

func (s Store) ReadSession(ctx context.Context, token string, now time.Time) (Session, error) {
	var session Session
	if !validToken(token) {
		return session, ErrUnauthenticated
	}
	err := s.Pool.QueryRow(ctx, `SELECT person_id::text,csrf_token,expires_at FROM sessions
		WHERE token_digest=$1 AND revoked_at IS NULL AND expires_at>$2`, digest(token), now).Scan(
		&session.PersonID, &session.CSRFToken, &session.ExpiresAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return Session{}, ErrUnauthenticated
	}
	if err != nil {
		return Session{}, fmt.Errorf("read session: %w", err)
	}
	return session, nil
}

// RotateSession serializes with logout/confirmation on an existing session.
// A missing/invalid previous cookie has no authority but does not block login.
func (s Store) RotateSession(ctx context.Context, personID, oldToken string, now time.Time) (string, error) {
	token, err := randomToken()
	if err != nil {
		return "", err
	}
	csrf, err := randomToken()
	if err != nil {
		return "", err
	}
	tx, err := s.Pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return "", fmt.Errorf("begin session rotation: %w", err)
	}
	defer tx.Rollback(ctx)
	if validToken(oldToken) {
		// This lock takes the same place in the order as a T3 confirmation.
		var priorDigest []byte
		err = tx.QueryRow(ctx, `SELECT token_digest FROM sessions WHERE token_digest=$1 FOR UPDATE`, digest(oldToken)).Scan(&priorDigest)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return "", fmt.Errorf("lock prior session: %w", err)
		}
		if err == nil {
			if _, err = tx.Exec(ctx, `UPDATE sessions SET revoked_at=$2 WHERE token_digest=$1 AND revoked_at IS NULL`, priorDigest, now); err != nil {
				return "", fmt.Errorf("revoke prior session: %w", err)
			}
		}
	}
	_, err = tx.Exec(ctx, `INSERT INTO sessions(token_digest,person_id,csrf_token,created_at,expires_at)
		VALUES($1,$2,$3,$4,$5)`, digest(token), personID, csrf, now, now.Add(sessionLifetime))
	if err != nil {
		return "", fmt.Errorf("insert session: %w", err)
	}
	if err = tx.Commit(ctx); err != nil {
		return "", fmt.Errorf("commit session rotation: %w", err)
	}
	return token, nil
}

func (s Store) RevokeSession(ctx context.Context, token string, now time.Time) error {
	tx, err := s.Pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return fmt.Errorf("begin logout: %w", err)
	}
	defer tx.Rollback(ctx)
	if _, err = LockSession(ctx, tx, token, now); err != nil {
		return err
	}
	_, err = tx.Exec(ctx, `UPDATE sessions SET revoked_at=$2 WHERE token_digest=$1`, digest(token), now)
	if err != nil {
		return fmt.Errorf("revoke session: %w", err)
	}
	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit logout: %w", err)
	}
	return nil
}
