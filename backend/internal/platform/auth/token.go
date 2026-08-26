package auth

import (
	"crypto/ecdsa"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Token primitives for the login/session slice (account task #3).
//
// ⚠️ TIER 0 — FENCED SUB-AREA (root AGENTS.md §3; tasks.md KPI table):
// this file is agent-DRAFTED and MUST go through a dedicated human paired
// rewrite/review pass BEFORE `make verify` sign-off and commit. See
// .local-agents/works/account/03-login-session-management/2-plan/tasks/
// task-03-tier0-token-helpers.md and techplan Resolved #13.
//
// Design core — two token purposes, cryptographically separated:
//
//	Access token   ES256 (asymmetric), signed with Keys.Private, verified
//	               with Keys.Public. Designed to be verifiable by other
//	               future services, hence asymmetric.
//	MFA-pending    HS256 (symmetric), signed with a DEDICATED 32-byte
//	token          secret (MFA_PENDING_TOKEN_SECRET) that shares key
//	               material with nothing else. Nothing outside this one
//	               backend process ever verifies it.
//
// Why a separate KEY rather than only a `purpose` claim (spec Assumption
// A/B): a claim check is application logic that can be buggy or omitted on
// some future code path. A separate key turns wrong-purpose acceptance into
// outright signature-verification failure — a cryptographic guarantee, not
// a logic guarantee. The `purpose` claim is kept anyway as defense-in-depth
// on top of the key separation (belt and suspenders, not either/or).

const (
	// AccessTokenTTL is the single source of truth for the access-token
	// session lifetime (15 min), shared by feature 02's OAuth issuance
	// and the login/session slice's credential/MFA issuance paths.
	AccessTokenTTL = 15 * time.Minute

	// MFAPendingTokenTTL bounds the exposure window of a proof-of-password
	// carrier that is never persisted: long enough to type a 6-digit code,
	// short enough that leakage (browser history, proxy log) expires fast.
	// There is deliberately NO verification leeway on this token — unlike
	// the access token, mint and verify happen inside one process seconds
	// apart, and tightness is a security property here (R6).
	MFAPendingTokenTTL = 5 * time.Minute

	// accessTokenLeeway tolerates small clock skew between issuing and
	// verifying processes for the longer-lived access token, matching the
	// 1-minute convention used for Google id_token checks (auth_google.go).
	accessTokenLeeway = time.Minute
)

// Token purpose claim values. A verifier accepts ONLY its own purpose;
// anything else (including a missing claim, i.e. legacy OAuth-era tokens)
// is rejected.
const (
	purposeAccess     = "access"
	purposeMFAPending = "mfa_pending"
)

// purposeClaims extends the standard registered claims with the explicit
// purpose marker checked as defense-in-depth on top of key separation.
type purposeClaims struct {
	jwt.RegisteredClaims
	Purpose string `json:"purpose"`
}

// ErrInvalidToken is the generic failure returned by every verifier for any
// rejection reason (bad signature, wrong algorithm, wrong purpose, missing
// or expired claims, malformed input). Callers must not — and cannot —
// distinguish further: leaking WHY a token failed would aid probing.
var ErrInvalidToken = errors.New("invalid token")

// MintAccessToken signs an ES256 access JWT for userID with claims
// {sub, iat, exp=now+15m, purpose:"access"}.
func MintAccessToken(private *ecdsa.PrivateKey, userID uuid.UUID, now time.Time) (string, error) {
	if private == nil {
		return "", fmt.Errorf("auth: access token signing key not configured")
	}
	claims := purposeClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(AccessTokenTTL)),
		},
		Purpose: purposeAccess,
	}
	signed, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString(private)
	if err != nil {
		return "", fmt.Errorf("auth: sign access token: %w", err)
	}
	return signed, nil
}

// VerifyAccessToken parses and validates an access JWT: ES256 signature
// under public, exp required (1-minute leeway), subject parseable as a
// UUID, and purpose == "access". Any failure collapses to ErrInvalidToken.
//
// Note: tokens minted before the purpose claim existed (feature-02 OAuth
// flow) are rejected here BY DESIGN — accepted breaking edge for a sandbox
// with no deployed clients. The legacy GoogleTokenVerifier stays lenient
// for link/reauth gating; this verifier serves new/future protected
// endpoints.
func VerifyAccessToken(public *ecdsa.PublicKey, token string, now time.Time) (uuid.UUID, error) {
	if public == nil || token == "" {
		return uuid.Nil, ErrInvalidToken
	}
	claims := &purposeClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return public, nil
	},
		jwt.WithValidMethods([]string{"ES256"}),
		jwt.WithExpirationRequired(),
		jwt.WithLeeway(accessTokenLeeway),
		jwt.WithTimeFunc(func() time.Time { return now }),
	)
	if err != nil || claims.Purpose != purposeAccess {
		return uuid.Nil, ErrInvalidToken
	}
	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return uuid.Nil, ErrInvalidToken
	}
	return userID, nil
}

// MintMFAPending signs the short-lived step-up token presented at
// /auth/login/mfa after a successful password check: HS256 under the
// dedicated secret, claims {sub, iat, exp=now+5m, purpose:"mfa_pending"}.
// The token is NEVER persisted — its signature and expiry are the only
// server-side state; every use still requires a correct TOTP/backup code
// on top, which is why non-single-use is acceptable.
func MintMFAPending(secret32 []byte, userID uuid.UUID, now time.Time) (string, error) {
	if err := validateSecretBytes(secret32); err != nil {
		return "", err
	}
	claims := purposeClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(MFAPendingTokenTTL)),
		},
		Purpose: purposeMFAPending,
	}
	signed, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret32)
	if err != nil {
		return "", fmt.Errorf("auth: sign mfa_pending token: %w", err)
	}
	return signed, nil
}

// VerifyMFAPending enforces: HS256 signature under secret32, exp required
// with NO leeway, subject parseable as a UUID, purpose == "mfa_pending".
// Wrong key material makes an ES256 access token fail signature
// verification outright — the cryptographic layer doing its job even if a
// caller forgets the purpose check somewhere.
func VerifyMFAPending(secret32 []byte, token string, now time.Time) (uuid.UUID, error) {
	if err := validateSecretBytes(secret32); err != nil {
		return uuid.Nil, err
	}
	if token == "" {
		return uuid.Nil, ErrInvalidToken
	}
	claims := &purposeClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return secret32, nil
	},
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithExpirationRequired(),
		// No leeway: see MFAPendingTokenTTL above.
		jwt.WithTimeFunc(func() time.Time { return now }),
	)
	if err != nil || claims.Purpose != purposeMFAPending {
		return uuid.Nil, ErrInvalidToken
	}
	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return uuid.Nil, ErrInvalidToken
	}
	return userID, nil
}

// ValidateMFAPendingSecret parses the MFA_PENDING_TOKEN_SECRET env value:
// base64 (std encoding), decoding to EXACTLY 32 bytes — the same startup
// discipline as platform/crypto New(), so misconfiguration fails fast at
// boot instead of at the first login attempt.
func ValidateMFAPendingSecret(b64 string) ([]byte, error) {
	const envName = "MFA_PENDING_TOKEN_SECRET"
	if b64 == "" {
		return nil, fmt.Errorf("%s is empty", envName)
	}
	key, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, fmt.Errorf("%s is not valid base64: %w", envName, err)
	}
	if err := validateSecretBytes(key); err != nil {
		return nil, fmt.Errorf("%s invalid: %w", envName, err)
	}
	return key, nil
}

func validateSecretBytes(key []byte) error {
	if len(key) != 32 {
		return fmt.Errorf("auth: mfa_pending secret must decode to 32 bytes, got %d", len(key))
	}
	return nil
}
