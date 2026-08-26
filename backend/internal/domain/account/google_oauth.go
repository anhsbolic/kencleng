package account

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/anhsbolic/kencleng/backend/internal/platform/auth"
	"github.com/anhsbolic/kencleng/backend/internal/platform/crypto"
	"github.com/anhsbolic/kencleng/backend/internal/platform/googleoauth"
)

// Intent values for the shared Google OAuth endpoints.
const (
	intentLogin  = "login"
	intentLink   = "link"
	intentReauth = "reauth"
)

// Callback error codes surfaced to the frontend via ?error={code}
// (techplan §8; frontend sign-off on the final set is pending — §14).
const (
	errStateMismatch       = "state_mismatch"
	errNonceMismatch       = "nonce_mismatch"
	errGoogleTokenInvalid  = "google_token_invalid"
	errGoogleUnavailable   = "google_unavailable"
	errGoogleEmailConflict = "google_email_conflict"
	errGoogleLinkConflict  = "google_link_conflict"
)

// actionAccountLinking is the canonical user_logs.action_type literal for a
// successful link intent (Fitur 9's "account linking baru"). The full
// vocabulary and the DB-level immutability constraint are owned by task #08.
const actionAccountLinking = "account_linking"

// TTLs for the OAuth flow. stateCookieTTL matches the transport cookie
// MaxAge (R24); refreshTokenTTL is the first-generation session lifetime
// (rotation arrives with the login/session task). accessTokenTTL lives in
// platform/auth (AccessTokenTTL) — the single source of truth shared with
// the login/session slice.
const (
	stateCookieTTL  = 10 * time.Minute
	refreshTokenTTL = 30 * 24 * time.Hour
)

// Frontend destinations relative to the configured FRONTEND_URL. The login
// page hosts login-flow errors; the security page hosts link/reauth results.
// The security path mirrors the account-security endpoint namespace
// (/account/security/*); flagged as an assumption in the ticket report.
const (
	frontendLoginPath    = "/login"
	frontendSecurityPath = "/account/security"
)

// Sentinel errors for the redirect leg — mapped by the transport handler:
// ErrInvalidIntent → 400 (R18), ErrMissingSession → 401 (R2). The callback
// leg reports outcomes via CallbackResult.Error instead (302 contract).
var (
	ErrInvalidIntent  = errors.New("invalid intent")
	ErrMissingSession = errors.New("session required")
)

// googleOAuthClient is the subset of googleoauth.Client the service depends
// on. *googleoauth.Client satisfies it; tests inject a fake so no network
// is hit and every failure mode (timeout, forged token, replayed nonce) is
// exercisable deterministically — same seam pattern as breachChecker.
type googleOAuthClient interface {
	AuthURL(state, nonce string) string
	ExchangeCode(ctx context.Context, code string) (*googleoauth.TokenResponse, error)
	VerifyIDToken(ctx context.Context, idToken, expectedNonce string) (*googleoauth.Claims, error)
}

// oauthState is the JSON payload carried in the short-TTL HttpOnly cookie
// between redirect and callback: the CSRF state, the replay nonce, which
// flow initiated the round-trip, and (for link/reauth) the authenticated
// user the round-trip belongs to.
type oauthState struct {
	State  string     `json:"state"`
	Nonce  string     `json:"nonce"`
	Intent string     `json:"intent"`
	UserID *uuid.UUID `json:"user_id,omitempty"`
}

// CallbackResult is the outcome of GoogleCallback. On success Error is empty
// and Tokens carry the issued session tokens; on failure Error carries one
// of the ?error={code} literals and Tokens are zero. RedirectURL is always
// populated (success or error destination). Reauth is true only when the
// reauth intent completed successfully — the handler uses it to set the
// short-lived reauth marker, keyed by UserID (carried from the cookie's
// session binding; zero when no user is bound).
type CallbackResult struct {
	RedirectURL  string
	Error        string
	AccessToken  string
	RefreshToken string
	Reauth       bool
	UserID       uuid.UUID
}

// randomHex returns n random bytes hex-encoded (64 chars for n=32).
func randomHex(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("read rand: %w", err)
	}
	return hex.EncodeToString(b), nil
}

// validIntent reports whether intent is one of the three supported values.
func validIntent(intent string) bool {
	switch intent {
	case intentLogin, intentLink, intentReauth:
		return true
	}
	return false
}

// encodeOAuthState serializes the state payload into the opaque cookie value:
// JSON → base64. Decoding failures at callback time map to state_mismatch.
func encodeOAuthState(st oauthState) (string, error) {
	raw, err := json.Marshal(st)
	if err != nil {
		return "", fmt.Errorf("marshal oauth state: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(raw), nil
}

func decodeOAuthState(cookieValue string) (oauthState, error) {
	var st oauthState
	raw, err := base64.RawURLEncoding.DecodeString(cookieValue)
	if err != nil {
		return st, fmt.Errorf("decode oauth state cookie: %w", err)
	}
	if err := json.Unmarshal(raw, &st); err != nil {
		return st, fmt.Errorf("unmarshal oauth state cookie: %w", err)
	}
	if !validIntent(st.Intent) || st.State == "" || st.Nonce == "" {
		return st, errors.New("oauth state cookie incomplete")
	}
	return st, nil
}

// GoogleRedirect validates the requested intent, generates the CSRF state
// and replay nonce, encodes both into the opaque cookie value, and builds
// the Google consent-screen URL (R1). The session check itself lives in the
// handler (explicit authz boundary, R25) — this method only enforces the
// invariant defensively: link/reauth without a session user is rejected
// before anything is generated.
//
// Returned values: consentURL for the 302 Location, cookieValue for the
// Set-Cookie header (transport sets attributes per R24).
func (s *Service) GoogleRedirect(ctx context.Context, intent string, sessionUserID *uuid.UUID) (consentURL, cookieValue string, err error) {
	if !validIntent(intent) {
		return "", "", ErrInvalidIntent
	}
	if intent != intentLogin && sessionUserID == nil {
		return "", "", ErrMissingSession
	}

	state, err := randomHex(32)
	if err != nil {
		return "", "", fmt.Errorf("account: generate oauth state: %w", err)
	}
	nonce, err := randomHex(32)
	if err != nil {
		return "", "", fmt.Errorf("account: generate oauth nonce: %w", err)
	}

	st := oauthState{State: state, Nonce: nonce, Intent: intent, UserID: sessionUserID}
	cookieValue, err = encodeOAuthState(st)
	if err != nil {
		return "", "", fmt.Errorf("account: encode oauth state: %w", err)
	}

	log.Printf("account: google oauth redirect intent=%s session=%t", intent, sessionUserID != nil)
	return s.googleOAuth.AuthURL(state, nonce), cookieValue, nil
}

// failResult builds an error CallbackResult targeting the right frontend
// page: login-intent failures land on /login, link/reauth failures on the
// security page. An undecodable cookie means the intent is unknown — the
// safest landing is the login entry point.
func (s *Service) failResult(intentKnown bool, intent string, code string) CallbackResult {
	base := s.frontendURL + frontendLoginPath
	if intentKnown && intent == intentLink || intentKnown && intent == intentReauth {
		base = s.frontendURL + frontendSecurityPath
	}
	return CallbackResult{
		Error:       code,
		RedirectURL: base + "?error=" + code,
	}
}

// successResult wraps issued tokens with the post-login app URL.
func (s *Service) successResult(access, refresh string) CallbackResult {
	return CallbackResult{
		RedirectURL:  s.frontendURL,
		AccessToken:  access,
		RefreshToken: refresh,
	}
}

// GoogleCallback validates state (constant-time, R23) BEFORE any Google API
// call (R4/R19/R20), exchanges the code, verifies the id_token (replay vs
// forgery distinguished per R5/R26), then branches on the cookie's intent:
//
//   - login: existing google identity → issue tokens (R7); email claimed by
//     an email_password identity of another user → NO auto-merge, clean
//     error (R9, top-severity takeover threat); otherwise create User +
//     verified AuthIdentity in one tx and issue tokens (R8, R14).
//   - link: email claimed by another user → reject (R10); otherwise attach
//     the identity to the SESSION user (never a new User row) plus an
//     audit entry, atomically (R11, Fitur 9).
//   - reauth: no identity/token writes at all; result flags the handler to
//     set the short-lived marker (R12).
//
// Every failure path returns a non-nil result with Error set and nil error —
// the error channel is reserved for unexpected internal failures.
func (s *Service) GoogleCallback(ctx context.Context, code, state, cookieValue string) (CallbackResult, error) {
	// R19: missing params → state_mismatch before anything else.
	if code == "" || state == "" || cookieValue == "" {
		return s.failResult(false, "", errStateMismatch), nil
	}

	st, err := decodeOAuthState(cookieValue)
	intentKnown := err == nil
	if err != nil {
		// R20: expired/corrupt/missing cookie.
		log.Printf("account: google oauth state cookie rejected (reason=undecodable)")
		return s.failResult(intentKnown, st.Intent, errStateMismatch), nil
	}

	// R4/R23: constant-time comparison, and strictly before any Google API
	// call — a wrong or replayed state must never trigger outbound traffic.
	if subtle.ConstantTimeCompare([]byte(state), []byte(st.State)) != 1 {
		log.Printf("account: google oauth state mismatch intent=%s", st.Intent)
		return s.failResult(true, st.Intent, errStateMismatch), nil
	}

	tokens, err := s.googleOAuth.ExchangeCode(ctx, code)
	if err != nil {
		// R6: client already logged a sanitized category; raw error stays
		// internal — the caller sees a clean error redirect.
		log.Printf("account: google token exchange failed category=sanitized")
		return s.failResult(true, st.Intent, errGoogleUnavailable), nil
	}

	claims, err := s.googleOAuth.VerifyIDToken(ctx, tokens.IDToken, st.Nonce)
	if err != nil {
		if errors.Is(err, googleoauth.ErrNonceMismatch) {
			// R5: replay — distinguishable from every other verification
			// failure (R26 keeps those generic).
			log.Printf("account: google oauth nonce mismatch intent=%s", st.Intent)
			return s.failResult(true, st.Intent, errNonceMismatch), nil
		}
		log.Printf("account: google id_token verification failed intent=%s", st.Intent)
		return s.failResult(true, st.Intent, errGoogleTokenInvalid), nil
	}

	identifierHash := crypto.HMAC([]byte(claims.Email), s.keys)

	switch st.Intent {
	case intentLogin:
		return s.callbackLogin(ctx, claims.Email, identifierHash)
	case intentLink:
		return s.callbackLink(ctx, st, claims.Email, identifierHash)
	case intentReauth:
		// R12: no AuthIdentity or token changes. The handler sets the
		// 5-minute marker keyed on the cookie's user_id upon seeing
		// Reauth=true.
		if st.UserID == nil {
			return s.failResult(true, st.Intent, errStateMismatch), nil
		}
		log.Printf("account: google reauth succeeded user_id=%s", *st.UserID)
		return CallbackResult{
			RedirectURL: s.frontendURL + frontendSecurityPath,
			Reauth:      true,
			UserID:      *st.UserID,
		}, nil
	default:
		return s.failResult(true, st.Intent, errStateMismatch), nil
	}
}

// callbackLogin implements the three login branches (R7-R9).
func (s *Service) callbackLogin(ctx context.Context, email, identifierHash string) (CallbackResult, error) {
	googleIdentity, err := s.repo.FindAuthIdentityByIdentifierHash(ctx, providerGoogle, identifierHash)
	if err != nil {
		return CallbackResult{}, fmt.Errorf("account: lookup google identity: %w", err)
	}
	if googleIdentity != nil {
		access, refresh, err := s.IssueTokens(ctx, googleIdentity.UserID)
		if err != nil {
			return CallbackResult{}, err
		}
		log.Printf("account: google login completed user_id=%s (existing identity)", googleIdentity.UserID)
		return s.successResult(access, refresh), nil
	}

	epIdentity, err := s.repo.FindAuthIdentityByIdentifierHash(ctx, providerEmailPassword, identifierHash)
	if err != nil {
		return CallbackResult{}, fmt.Errorf("account: lookup email_password identity: %w", err)
	}
	if epIdentity != nil {
		// R9 — NO AUTO-MERGE. The email belongs to an email_password
		// account of a different user. Creating or attaching anything here
		// would enable account takeover via an unverified provider email
		// claim. No new records; clean error redirect.
		log.Printf("account: google login blocked by no-auto-merge rule (email_password claim)")
		return s.failResult(true, intentLogin, errGoogleEmailConflict), nil
	}

	userID, conflictAfterRace, err := s.registerGoogleUser(ctx, email, identifierHash)
	if err != nil {
		return CallbackResult{}, err
	}
	if conflictAfterRace {
		// R15 fallback: unique index fired and the winning identity is not
		// visible to re-lookup yet — surface the closest recoverable signal
		// rather than crashing or auto-merging. Flagged assumption in the
		// ticket report.
		return s.failResult(true, intentLogin, errGoogleEmailConflict), nil
	}

	access, refresh, err := s.IssueTokens(ctx, userID)
	if err != nil {
		return CallbackResult{}, err
	}
	log.Printf("account: google login completed user_id=%s (new user)", userID)
	return s.successResult(access, refresh), nil
}

// registerGoogleUser creates User + google AuthIdentity in one transaction
// (R8). The identity is created already verified — verified_at = now(), it
// never passes through null (R14). On a concurrent duplicate (unique index
// on (provider_type, identifier_hash), INV-account-01) the tx rolls back
// cleanly and the winner's identity is looked up again: if visible, the
// caller proceeds with a normal login for that user; if not yet visible,
// conflictAfterRace=true gives a clean, non-crashing failure (R15).
//
// Display name: the Google email doubles as the initial display name — the
// id_token claims consumed here carry no separate name field. Flagged
// assumption in the ticket report.
func (s *Service) registerGoogleUser(ctx context.Context, email, identifierHash string) (userID uuid.UUID, conflictAfterRace bool, err error) {
	tx, err := s.tx.BeginTx(ctx)
	if err != nil {
		return uuid.UUID{}, false, fmt.Errorf("account: begin tx: %w", err)
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback(ctx)
		}
	}()

	now := time.Now()
	newUserID := uuid.New()
	user := &User{ID: newUserID, Name: email, PrimaryEmail: email}
	if err := s.repo.InsertUser(ctx, tx, user); err != nil {
		return uuid.UUID{}, false, fmt.Errorf("account: insert user: %w", err)
	}

	identity := &AuthIdentity{
		ID:           uuid.New(),
		UserID:       newUserID,
		ProviderType: providerGoogle,
		Identifier:   email,
		VerifiedAt:   &now, // R14: google identities start verified
	}
	insertErr := s.repo.InsertAuthIdentity(ctx, tx, identity)
	if insertErr != nil && isUniqueViolation(insertErr) {
		// R15: concurrent duplicate registration. Roll back our partial
		// user row and re-lookup the winner's identity.
		if lookErr := tx.Rollback(ctx); lookErr != nil {
			return uuid.UUID{}, false, fmt.Errorf("account: rollback after duplicate: %w", lookErr)
		}
		winner, findErr := s.repo.FindAuthIdentityByIdentifierHash(ctx, providerGoogle, identifierHash)
		if findErr != nil {
			return uuid.UUID{}, false, fmt.Errorf("account: lookup winner identity after duplicate: %w", findErr)
		}
		if winner == nil {
			log.Printf("account: concurrent google registration detected, winner not yet visible")
			return uuid.UUID{}, true, nil
		}
		return winner.UserID, false, nil
	}
	if insertErr != nil {
		return uuid.UUID{}, false, fmt.Errorf("account: insert google auth_identity: %w", insertErr)
	}

	if err := tx.Commit(ctx); err != nil {
		return uuid.UUID{}, false, fmt.Errorf("account: commit google register: %w", err)
	}
	committed = true

	log.Printf("account: google user registered user_id=%s (provider=google)", newUserID)
	return newUserID, false, nil
}

// callbackLink implements the two link branches (R10/R11). The identity
// attaches to the SESSION user from the cookie — never a new User row — and
// the audit entry commits atomically with it.
func (s *Service) callbackLink(ctx context.Context, st oauthState, email, identifierHash string) (CallbackResult, error) {
	if st.UserID == nil {
		// Defensive: the handler must have enforced a session for link
		// intent. Treat a missing binding as an invalid round-trip.
		log.Printf("account: google link callback without session binding")
		return s.failResult(true, st.Intent, errStateMismatch), nil
	}
	sessionUserID := *st.UserID

	googleIdentity, err := s.repo.FindAuthIdentityByIdentifierHash(ctx, providerGoogle, identifierHash)
	if err != nil {
		return CallbackResult{}, fmt.Errorf("account: lookup google identity: %w", err)
	}
	if googleIdentity != nil && googleIdentity.UserID != sessionUserID {
		// R10: claimed by a different user — no attach, clean rejection.
		log.Printf("account: google link blocked by ownership conflict")
		return s.failResult(true, st.Intent, errGoogleLinkConflict), nil
	}

	if googleIdentity == nil {
		tx, err := s.tx.BeginTx(ctx)
		if err != nil {
			return CallbackResult{}, fmt.Errorf("account: begin tx: %w", err)
		}
		committed := false
		defer func() {
			if !committed {
				_ = tx.Rollback(ctx)
			}
		}()

		now := time.Now()
		identity := &AuthIdentity{
			ID:           uuid.New(),
			UserID:       sessionUserID,
			ProviderType: providerGoogle,
			Identifier:   email,
			VerifiedAt:   &now, // R14
		}
		if err := s.repo.InsertAuthIdentity(ctx, tx, identity); err != nil {
			return CallbackResult{}, fmt.Errorf("account: insert linked identity: %w", err)
		}
		entry := &UserLog{
			ID:         uuid.New(),
			UserID:     sessionUserID,
			ActionType: actionAccountLinking,
			CreatedAt:  now,
		}
		if err := s.repo.InsertUserLog(ctx, tx, entry); err != nil {
			return CallbackResult{}, fmt.Errorf("account: insert link audit log: %w", err)
		}
		if err := tx.Commit(ctx); err != nil {
			return CallbackResult{}, fmt.Errorf("account: commit link: %w", err)
		}
		committed = true
	} else {
		// Already linked to this very user — idempotent success, nothing
		// to write (a second insert would violate the unique index).
		log.Printf("account: google link no-op, already attached user_id=%s", sessionUserID)
	}

	log.Printf("account: google link completed user_id=%s", sessionUserID)
	return CallbackResult{
		RedirectURL: s.frontendURL + frontendSecurityPath,
	}, nil
}

// IssueTokens creates a first-generation session for userID: an ES256
// access token (15 min) and a refresh token (30 days) whose plain value is
// returned once and whose SHA-256 hash is persisted (family_id starts a new
// lineage; rotation/reuse detection is built on these primitives by the
// login/session task).
func (s *Service) IssueTokens(ctx context.Context, userID uuid.UUID) (accessToken, refreshToken string, err error) {
	if s.authKeys == nil || s.authKeys.Private == nil {
		return "", "", errors.New("account: auth keys not configured for token issuance")
	}

	now := time.Now()
	accessClaims := jwt.RegisteredClaims{
		Subject:   userID.String(),
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(auth.AccessTokenTTL)),
	}
	access, err := jwt.NewWithClaims(jwt.SigningMethodES256, accessClaims).SignedString(s.authKeys.Private)
	if err != nil {
		return "", "", fmt.Errorf("account: sign access token: %w", err)
	}

	plainRefresh, err := randomHex(32)
	if err != nil {
		return "", "", fmt.Errorf("account: generate refresh token: %w", err)
	}

	rt := &RefreshToken{
		ID:        uuid.New(),
		UserID:    userID,
		FamilyID:  uuid.New(),
		TokenHash: sha256Hex(plainRefresh),
		ExpiresAt: now.Add(refreshTokenTTL),
		CreatedAt: now,
	}
	if err := s.insertRefreshToken(ctx, rt); err != nil {
		return "", "", err
	}

	log.Printf("account: session tokens issued user_id=%s", userID)
	return access, plainRefresh, nil
}

// insertRefreshToken persists the refresh token row in its own transaction
// (single-row write wrapped for consistency with the repo's tx-taking
// interface).
func (s *Service) insertRefreshToken(ctx context.Context, rt *RefreshToken) error {
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
	if err := s.repo.InsertRefreshToken(ctx, tx, rt); err != nil {
		return fmt.Errorf("account: insert refresh_token: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("account: commit refresh_token: %w", err)
	}
	committed = true
	return nil
}
