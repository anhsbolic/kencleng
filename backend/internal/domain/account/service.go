package account

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/anhsbolic/kencleng/backend/internal/platform/breachcheck"
	"github.com/anhsbolic/kencleng/backend/internal/platform/crypto"
	"github.com/anhsbolic/kencleng/backend/internal/platform/notification"
	"github.com/anhsbolic/kencleng/backend/internal/platform/secrets"
)

// Sentinel errors — mapped to HTTP status by transport/http (Task 05).
// Defined here so the service owns its error vocabulary; the transport
// layer only translates.
var (
	// ErrValidation indicates a request failed validation (password
	// policy, breach-list hit). Maps to 422.
	ErrValidation = errors.New("validation error")
	// ErrTokenExpired indicates a verification token exists but has
	// expired. Maps to 410.
	ErrTokenExpired = errors.New("token expired")
	// ErrTokenNotFound indicates a verification token was not found,
	// already used, or revoked (the 3-clause guard rejected it for any
	// reason other than expiry). Maps to 404.
	ErrTokenNotFound = errors.New("token not found")
)

// tokenTTL is the validity window for an email_verification token.
const tokenTTL = 24 * time.Hour

// emailVerification is the auth_identities provider_type / auth_tokens
// purpose for the email/password verification flow.
const (
	providerEmailPassword = "email_password"
	providerGoogle        = "google"
	purposeEmailVerify    = "email_verification"
)

// TxRunner abstracts db.BeginTx so the service is unit-testable without a
// real pgxpool. The real implementation, poolRunner, wraps *pgxpool.Pool.
type TxRunner interface {
	BeginTx(ctx context.Context) (pgx.Tx, error)
}

// breachChecker is the subset of breachcheck.Client the service depends
// on. *breachcheck.Client satisfies it; tests inject a fake so no network
// is hit and the fail-open path is exercisable deterministically.
type breachChecker interface {
	IsBreached(ctx context.Context, password string) (bool, error)
}

// poolRunner adapts *pgxpool.Pool to TxRunner.
type poolRunner struct{ pool *pgxpool.Pool }

// BeginTx begins a transaction on the underlying pool.
func (p poolRunner) BeginTx(ctx context.Context) (pgx.Tx, error) {
	return p.pool.BeginTx(ctx, pgx.TxOptions{})
}

// Service implements the account domain's registration, verification,
// and resend flows. It is safe for concurrent use: it holds no mutable
// state of its own, and every injected dependency is goroutine-safe.
type Service struct {
	repo        Repository
	tx          TxRunner
	breachCheck breachChecker
	email       notification.Sender
	keys        *crypto.Keys
}

// NewService constructs an account Service. db is used only to begin
// transactions; reads/standalone writes go through repo.
func NewService(repo Repository, db *pgxpool.Pool, bc *breachcheck.Client, sender notification.Sender, keys *crypto.Keys) *Service {
	return &Service{
		repo:        repo,
		tx:          poolRunner{pool: db},
		breachCheck: bc,
		email:       sender,
		keys:        keys,
	}
}

// Register orchestrates R1-R7, R16-R19. The four internal branches
// (new / unverified-existing / verified-existing / Google-only-conflict)
// all return nil so the handler (Task 05) writes an identical 202
// generic response. Branch wall-clock time is shaped to be equivalent
// (R7): bcrypt runs on every branch, and every branch performs at least
// one DB round-trip.
//
// PII handling: email is never logged. Only the fact ("register
// attempt"/"register completed") and, after creation, user_id are
// logged.
func (s *Service) Register(ctx context.Context, name, email, password string) error {
	// R5/R18: validate password policy BEFORE any enumeration-sensitive
	// branch lookup so a validation failure cannot leak whether the
	// email is registered.
	if err := s.validatePassword(ctx, password); err != nil {
		return err
	}

	// R7 (CPU time): always run bcrypt. The new-user branch stores the
	// result as credential_secret; the other three branches discard it.
	// Never skip this on no-op branches, never replace with a sleep.
	passwordHash, err := secrets.HashPassword(password)
	if err != nil {
		return fmt.Errorf("account: hash password: %w", err)
	}

	// Compute the lookup hash for the email_password identity.
	identifierHash := crypto.HMAC([]byte(email), s.keys)

	// R1-R4 branch dispatch on the existing email_password identity.
	identity, err := s.repo.FindAuthIdentityByIdentifierHash(ctx, providerEmailPassword, identifierHash)
	if err != nil {
		return fmt.Errorf("account: lookup email_password identity: %w", err)
	}

	switch {
	case identity == nil:
		// No email_password identity. Check for a Google-only conflict
		// (R4, R17) before creating a new user.
		google, err := s.repo.FindAuthIdentityByIdentifierHash(ctx, providerGoogle, identifierHash)
		if err != nil {
			return fmt.Errorf("account: lookup google identity: %w", err)
		}
		if google != nil {
			// R4 / R17: Google-only conflict. No new user, no token.
			// Send nudge after the (read-only) DB round-trip above.
			s.sendNudge(ctx, email, notification.NudgeGoogleOnly)
			return nil
		}
		// R1: new user. Insert user + identity + token in one tx.
		return s.registerNewUser(ctx, name, email, passwordHash)

	case identity.VerifiedAt == nil:
		// R2: unverified existing. Resend-verification flow: revoke old
		// tokens, issue a new one, send a resend-verification nudge. No
		// new user/identity.
		if _, err := s.issueNewVerificationToken(ctx, identity.UserID); err != nil {
			return err
		}
		s.sendNudge(ctx, email, notification.NudgeResendVerification)
		return nil

	default:
		// R3: verified existing. No new record. The lookup above is the
		// DB round-trip; send a password-reset nudge.
		s.sendNudge(ctx, email, notification.NudgePasswordReset)
		return nil
	}
}

// registerNewUser inserts a new User + AuthIdentity + AuthToken in a
// single transaction (R16: on a concurrent duplicate, the
// auth_identities unique index fails and the whole tx rolls back
// cleanly — no orphaned users row). Sends the verification email after
// commit. R1.
func (s *Service) registerNewUser(ctx context.Context, name, email, passwordHash string) error {
	userID := uuid.New()
	identityID := uuid.New()
	plainToken, tokenHash, err := generateToken()
	if err != nil {
		return fmt.Errorf("account: generate token: %w", err)
	}

	tx, err := s.tx.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("account: begin tx: %w", err)
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback(ctx)
		}
	}()

	user := &User{
		ID:           userID,
		Name:         name,
		PrimaryEmail: email,
	}
	if err := s.repo.InsertUser(ctx, tx, user); err != nil {
		return fmt.Errorf("account: insert user: %w", err)
	}

	identity := &AuthIdentity{
		ID:               identityID,
		UserID:           userID,
		ProviderType:     providerEmailPassword,
		Identifier:       email,
		CredentialSecret: &passwordHash,
	}
	if err := s.repo.InsertAuthIdentity(ctx, tx, identity); err != nil {
		// R16: concurrent duplicate registration. The unique index on
		// (provider_type, identifier_hash) failed — map to a clean
		// no-op (return nil → 202 generic), indistinguishable from a
		// normal R1 to the caller. The tx rolls back via the deferred
		// Rollback, leaving no orphaned users row.
		if isUniqueViolation(err) {
			log.Printf("account: concurrent duplicate registration suppressed (provider=%s)", providerEmailPassword)
			return nil
		}
		return fmt.Errorf("account: insert auth_identity: %w", err)
	}

	token := &AuthToken{
		ID:        uuid.New(),
		UserID:    userID,
		Purpose:   purposeEmailVerify,
		TokenHash: tokenHash,
		ExpiresAt: time.Now().Add(tokenTTL),
		CreatedAt: time.Now(),
	}
	if err := s.repo.InsertAuthToken(ctx, tx, token); err != nil {
		return fmt.Errorf("account: insert auth_token: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("account: commit register: %w", err)
	}
	committed = true

	// Send the verification email AFTER commit — never inside the tx.
	// If the send fails, the DB state is correct and the user can resend.
	s.sendVerification(ctx, email, plainToken)
	log.Printf("account: registration completed user_id=%s", userID)
	return nil
}

// issueNewVerificationToken revokes the user's old email_verification
// tokens and issues a new one in a single transaction (R13), returning
// the plain token so the caller can send the appropriate email after
// commit. Shared by R2 (register unverified existing, sends a nudge)
// and R13 (resend endpoint, sends the verification email with the
// token).
func (s *Service) issueNewVerificationToken(ctx context.Context, userID uuid.UUID) (string, error) {
	plainToken, tokenHash, err := generateToken()
	if err != nil {
		return "", fmt.Errorf("account: generate token: %w", err)
	}

	tx, err := s.tx.BeginTx(ctx)
	if err != nil {
		return "", fmt.Errorf("account: begin tx: %w", err)
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback(ctx)
		}
	}()

	if err := s.repo.RevokeTokens(ctx, tx, userID, purposeEmailVerify); err != nil {
		return "", fmt.Errorf("account: revoke tokens: %w", err)
	}
	token := &AuthToken{
		ID:        uuid.New(),
		UserID:    userID,
		Purpose:   purposeEmailVerify,
		TokenHash: tokenHash,
		ExpiresAt: time.Now().Add(tokenTTL),
		CreatedAt: time.Now(),
	}
	if err := s.repo.InsertAuthToken(ctx, tx, token); err != nil {
		return "", fmt.Errorf("account: insert auth_token: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return "", fmt.Errorf("account: commit resend: %w", err)
	}
	committed = true

	return plainToken, nil
}

// VerifyEmail orchestrates R8-R12. Single-use correctness is delegated
// to the repository's atomic RedeemToken (the 3-clause INV-account-08
// guard) — the service never does a read-then-write on tokens.
func (s *Service) VerifyEmail(ctx context.Context, token string) error {
	tokenHash := sha256Hex(token)

	ok, err := s.repo.RedeemToken(ctx, tokenHash)
	if err != nil {
		return fmt.Errorf("account: redeem token: %w", err)
	}
	if ok {
		// R8: valid token redeemed. Set the user's email_password
		// identity verified_at. The token carries user_id; verification
		// is keyed on (user_id, provider_type=email_password).
		if err := s.repo.SetUserVerified(ctx, s.userIDForToken(ctx, tokenHash), providerEmailPassword, time.Now()); err != nil {
			return fmt.Errorf("account: set verified: %w", err)
		}
		log.Printf("account: email verified (token redacted)")
		return nil
	}

	// !ok: disambiguate expired (R9) vs not-found/used/revoked (R10/R11).
	t, err := s.repo.FindAuthTokenByHash(ctx, tokenHash)
	if err != nil {
		return fmt.Errorf("account: find token: %w", err)
	}
	if t != nil && !t.ExpiresAt.After(time.Now()) {
		return ErrTokenExpired
	}
	return ErrTokenNotFound
}

// userIDForToken returns the user_id for a just-redeemed token. After a
// successful RedeemToken the row's used_at is set; we re-read it to get
// user_id. (RedeemToken returns only a bool to keep its atomic guard
// single-purpose.)
func (s *Service) userIDForToken(ctx context.Context, tokenHash string) uuid.UUID {
	t, err := s.repo.FindAuthTokenByHash(ctx, tokenHash)
	if err != nil || t == nil {
		// Should not happen immediately after a successful redeem; if it
		// does, return a zero UUID and let SetUserVerified affect zero
		// rows rather than crash. Logged for ops.
		log.Printf("account: could not re-fetch redeemed token (token redacted)")
		return uuid.Nil
	}
	return t.UserID
}

// ResendVerification orchestrates R13-R14. Rate limiting (R15) is
// transport middleware (Task 05). Returns nil for both the match and
// no-match branches so the handler writes an identical 202 generic.
func (s *Service) ResendVerification(ctx context.Context, email string) error {
	identifierHash := crypto.HMAC([]byte(email), s.keys)
	identity, err := s.repo.FindAuthIdentityByIdentifierHash(ctx, providerEmailPassword, identifierHash)
	if err != nil {
		return fmt.Errorf("account: lookup identity for resend: %w", err)
	}
	if identity != nil && identity.VerifiedAt == nil {
		// R13: unverified match — revoke + new token, then send the
		// verification email with the new token after commit.
		plainToken, err := s.issueNewVerificationToken(ctx, identity.UserID)
		if err != nil {
			return err
		}
		s.sendVerification(ctx, email, plainToken)
		return nil
	}
	// R14: no match / verified / google-only — no token, no email, no-op.
	return nil
}

// validatePassword enforces the password policy: length >= 8 and not in
// the breach list (fail-open). R5/R6/R18/R19. Checked before any branch
// lookup so it cannot leak email state.
func (s *Service) validatePassword(ctx context.Context, password string) error {
	if len(password) < 8 {
		return ErrValidation
	}
	breached, err := s.breachCheck.IsBreached(ctx, password)
	if err != nil {
		// IsBreached fails open (returns false, nil) on API unreachable;
		// a non-nil error here is a real failure, not a fail-open.
		return fmt.Errorf("account: breach check: %w", err)
	}
	if breached {
		return ErrValidation
	}
	return nil
}

// sendVerification sends the verification email with the plain token.
// The plain token leaves the process exactly once: here, after commit.
// It is never logged.
func (s *Service) sendVerification(ctx context.Context, email, plainToken string) {
	if err := s.email.SendVerificationEmail(ctx, email, plainToken); err != nil {
		// Post-commit email failure is non-fatal: the DB state is
		// correct and the user can resend. Log the fact, not the token.
		log.Printf("account: send verification email failed (recipient redacted): %v", err)
	}
}

// sendNudge sends a nudge email. The recipient address is passed to the
// sender but never logged by it (FakeSender redacts).
func (s *Service) sendNudge(ctx context.Context, email, nudgeType string) {
	if err := s.email.SendNudgeEmail(ctx, email, nudgeType); err != nil {
		log.Printf("account: send nudge email failed type=%s: %v", nudgeType, err)
	}
}

// generateToken produces a user-facing plain token and its stored
// SHA-256 hex hash. 32 random bytes from crypto/rand, hex-encoded for
// the email link; the SHA-256 of that plain token is stored (never the
// plain token itself).
func generateToken() (plainToken, tokenHash string, err error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", "", fmt.Errorf("generate token: read rand: %w", err)
	}
	plain := hex.EncodeToString(b)
	return plain, sha256Hex(plain), nil
}

// sha256Hex returns the lowercase SHA-256 hex digest of s.
func sha256Hex(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

// isUniqueViolation reports whether err is a PostgreSQL unique-constraint
// violation (SQLSTATE 23505). Used to detect concurrent duplicate
// registration (R16) without string-matching the error message.
func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}
