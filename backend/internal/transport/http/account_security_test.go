package http

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/anhsbolic/kencleng/backend/internal/domain/account"
)

// stubSecurityService is a configurable securityService for handler
// tests — exercises the transport contract (status codes, problem types,
// response bodies) without touching the domain logic (already covered
// by domain tests).
type stubSecurityService struct {
	setPasswordResult bool
	setPasswordErr    error
	setPasswordCalled bool
	unlinkErr         error
	unlinkCalled      bool

	enrollURI      string
	enrollErr      error
	enrollCalled   bool
	confirmCodes   []string
	confirmErr     error
	confirmCalled  bool
	disableErr     error
	disableCalled  bool
	reauthRequired bool
	reauthErr      error
	reauthCalled   bool
}

func (s *stubSecurityService) SetPassword(_ context.Context, _ uuid.UUID, _, _, _ string) (bool, error) {
	s.setPasswordCalled = true
	return s.setPasswordResult, s.setPasswordErr
}

func (s *stubSecurityService) UnlinkGoogle(_ context.Context, _ uuid.UUID, _ string) error {
	s.unlinkCalled = true
	return s.unlinkErr
}

func (s *stubSecurityService) MfaEnroll(_ context.Context, _ uuid.UUID) (string, error) {
	s.enrollCalled = true
	return s.enrollURI, s.enrollErr
}

func (s *stubSecurityService) MfaEnrollConfirm(_ context.Context, _ uuid.UUID, _ string) ([]string, error) {
	s.confirmCalled = true
	return s.confirmCodes, s.confirmErr
}

func (s *stubSecurityService) MfaDisable(_ context.Context, _ uuid.UUID, _ string) error {
	s.disableCalled = true
	return s.disableErr
}

func (s *stubSecurityService) MfaDisableReauthRequired(_ context.Context, _ uuid.UUID) (bool, error) {
	s.reauthCalled = true
	if s.reauthErr != nil {
		return false, s.reauthErr
	}
	return s.reauthRequired, nil
}

// ---- R15: session enforcement ---------------------------------------

func TestRequireSession_MissingToken_401(t *testing.T) {
	key, _ := newES256Signer(t)
	verifier := GoogleTokenVerifier(&key.PublicKey)

	called := false
	h := RequireSession(verifier)(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		called = true
	}))

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/account/security/set-password", nil)
	h.ServeHTTP(rec, req)

	if called {
		t.Errorf("handler must not be called when no token is presented")
	}
	if rec.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", rec.Code)
	}
}

func TestRequireSession_ExpiredOrGarbageToken_401(t *testing.T) {
	key, _ := newES256Signer(t)
	verifier := GoogleTokenVerifier(&key.PublicKey)

	cases := []struct {
		name  string
		token string
	}{
		{"garbage", "not-a-jwt"},
		{"expired", signExpiredToken(t, key)},
		{"wrong-key", signWithDifferentKey(t)},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			called := false
			h := RequireSession(verifier)(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
				called = true
			}))

			rec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/account/security/set-password", nil)
			req.Header.Set("Authorization", "Bearer "+tc.token)
			h.ServeHTTP(rec, req)

			if called {
				t.Errorf("handler must not be called for %s token", tc.name)
			}
			if rec.Code != http.StatusUnauthorized {
				t.Errorf("expected 401 for %s, got %d", tc.name, rec.Code)
			}
		})
	}
}

func TestRequireSession_BearerFallback_Accepted(t *testing.T) {
	key, signer := newES256Signer(t)
	verifier := GoogleTokenVerifier(&key.PublicKey)
	userID := uuid.New()

	var gotUserID uuid.UUID
	h := RequireSession(verifier)(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		gotUserID, _ = UserIDFromContext(r.Context())
	}))

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/account/security/set-password", nil)
	req.Header.Set("Authorization", "Bearer "+signer(userID.String()))
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}
	if gotUserID != userID {
		t.Errorf("userID mismatch: got %s, want %s", gotUserID, userID)
	}
}

// signExpiredToken signs an ES256 JWT that expired 1 hour ago.
func signExpiredToken(t *testing.T, key *ecdsa.PrivateKey) string {
	t.Helper()
	now := time.Now().Add(-1 * time.Hour)
	tok := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.RegisteredClaims{
		Subject:   uuid.New().String(),
		ExpiresAt: jwt.NewNumericDate(now),
	})
	signed, err := tok.SignedString(key)
	if err != nil {
		t.Fatalf("sign: %v", err)
	}
	return signed
}

// signWithDifferentKey generates a fresh key pair, signs a valid-shaped
// token with it, and returns the signed string — the verifier uses a
// different key, so this token must be rejected.
func signWithDifferentKey(t *testing.T) string {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("generate key: %v", err)
	}
	now := time.Now()
	tok := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.RegisteredClaims{
		Subject:   uuid.New().String(),
		ExpiresAt: jwt.NewNumericDate(now.Add(10 * time.Minute)),
	})
	signed, err := tok.SignedString(key)
	if err != nil {
		t.Fatalf("sign: %v", err)
	}
	return signed
}

// ---- SetPassword handler tests ---------------------------------------

func TestSetPasswordHandler_Branch1_202(t *testing.T) {
	stub := &stubSecurityService{setPasswordResult: false}
	h := SetPasswordHandler(stub)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/account/security/set-password",
		strings.NewReader(`{"email":"new@company.com","password":"strong-pw-123"}`))
	injectSessionUserID(req, uuid.New())

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusAccepted {
		t.Errorf("expected 202, got %d", rec.Code)
	}
	var resp map[string]string
	_ = json.NewDecoder(rec.Body).Decode(&resp)
	if resp["message"] != "Kalau email tersedia, cek inbox untuk verifikasi." {
		t.Errorf("unexpected message: %q", resp["message"])
	}
}

func TestSetPasswordHandler_Branch2_200(t *testing.T) {
	stub := &stubSecurityService{setPasswordResult: true}
	h := SetPasswordHandler(stub)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/account/security/set-password",
		strings.NewReader(`{"current_password":"old-pw","password":"new-strong-pw"}`))
	injectSessionUserID(req, uuid.New())

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}
}

func TestSetPasswordHandler_PolicyFail_422(t *testing.T) {
	stub := &stubSecurityService{setPasswordErr: account.ErrValidation}
	h := SetPasswordHandler(stub)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/account/security/set-password",
		strings.NewReader(`{"email":"new@company.com","password":"short"}`))
	injectSessionUserID(req, uuid.New())

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected 422, got %d", rec.Code)
	}
}

func TestSetPasswordHandler_WrongPassword_401(t *testing.T) {
	stub := &stubSecurityService{setPasswordErr: account.ErrInvalidCredentials}
	h := SetPasswordHandler(stub)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/account/security/set-password",
		strings.NewReader(`{"current_password":"WRONG","password":"new-strong-pw"}`))
	injectSessionUserID(req, uuid.New())

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", rec.Code)
	}
}

// Internal errors (DB outage, tx begin/commit failure — anything that is
// a wrapped fmt.Errorf rather than a service sentinel) must surface as
// the generic 500 Problem via MapServiceError's default, NOT as
// 401 invalid-credentials. Proves the dispatch category, not just that
// an error fired (code-review S1/BP1/C1).
func TestSetPasswordHandler_InternalError_500(t *testing.T) {
	stub := &stubSecurityService{
		setPasswordErr: fmt.Errorf("account: lookup identities for set-password: %w", errors.New("db down")),
	}
	h := SetPasswordHandler(stub)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/account/security/set-password",
		strings.NewReader(`{"email":"new@company.com","password":"strong-pw-123"}`))
	injectSessionUserID(req, uuid.New())

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500 for a wrapped internal error, got %d", rec.Code)
	}
	var p problem
	_ = json.NewDecoder(rec.Body).Decode(&p)
	if p.Type != "https://kencleng.dev/problems/internal" {
		t.Errorf("expected internal problem type, got %s", p.Type)
	}
	if p.Detail == "" || p.Detail == problemDetailGenericCredential {
		t.Errorf("detail must be the generic internal message, got %q", p.Detail)
	}
}

// ---- UnlinkGoogle handler tests -------------------------------------

func TestUnlinkGoogleHandler_Success_200(t *testing.T) {
	stub := &stubSecurityService{}
	h := UnlinkGoogleHandler(stub)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/account/security/google/unlink",
		strings.NewReader(`{"password":"correct-pw"}`))
	injectSessionUserID(req, uuid.New())

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}
	var resp map[string]string
	_ = json.NewDecoder(rec.Body).Decode(&resp)
	if resp["message"] != "Akun Google berhasil dilepas." {
		t.Errorf("unexpected message: %q", resp["message"])
	}
}

func TestUnlinkGoogleHandler_OnlyIdentity_409(t *testing.T) {
	stub := &stubSecurityService{unlinkErr: account.ErrOnlyIdentity}
	h := UnlinkGoogleHandler(stub)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/account/security/google/unlink",
		strings.NewReader(`{"password":"pw"}`))
	injectSessionUserID(req, uuid.New())

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusConflict {
		t.Errorf("expected 409, got %d", rec.Code)
	}
	var p problem
	_ = json.NewDecoder(rec.Body).Decode(&p)
	if p.Type != "https://kencleng.dev/errors/only-identity" {
		t.Errorf("expected only-identity type, got %s", p.Type)
	}
}

func TestUnlinkGoogleHandler_UnverifiedRemaining_409(t *testing.T) {
	stub := &stubSecurityService{unlinkErr: account.ErrRemainingUnverified}
	h := UnlinkGoogleHandler(stub)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/account/security/google/unlink",
		strings.NewReader(`{"password":"pw"}`))
	injectSessionUserID(req, uuid.New())

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusConflict {
		t.Errorf("expected 409, got %d", rec.Code)
	}
	var p problem
	_ = json.NewDecoder(rec.Body).Decode(&p)
	if p.Type != "https://kencleng.dev/errors/unverified-remaining-identity" {
		t.Errorf("expected unverified-remaining-identity type, got %s", p.Type)
	}
}

func TestUnlinkGoogleHandler_WrongPassword_401(t *testing.T) {
	stub := &stubSecurityService{unlinkErr: account.ErrInvalidCredentials}
	h := UnlinkGoogleHandler(stub)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/account/security/google/unlink",
		strings.NewReader(`{"password":"WRONG"}`))
	injectSessionUserID(req, uuid.New())

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", rec.Code)
	}
}

// Same dispatch contract as SetPassword: a wrapped internal error must
// surface as the generic 500 Problem, never as 401 invalid-credentials.
func TestUnlinkGoogleHandler_InternalError_500(t *testing.T) {
	stub := &stubSecurityService{
		unlinkErr: fmt.Errorf("account: delete google identities: %w", errors.New("db down")),
	}
	h := UnlinkGoogleHandler(stub)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/account/security/google/unlink",
		strings.NewReader(`{"password":"correct-pw"}`))
	injectSessionUserID(req, uuid.New())

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500 for a wrapped internal error, got %d", rec.Code)
	}
	var p problem
	_ = json.NewDecoder(rec.Body).Decode(&p)
	if p.Type != "https://kencleng.dev/problems/internal" {
		t.Errorf("expected internal problem type, got %s", p.Type)
	}
	if p.Detail == "" || p.Detail == problemDetailGenericCredential {
		t.Errorf("detail must be the generic internal message, got %q", p.Detail)
	}
}

// ---- helpers ---------------------------------------------------------

func injectSessionUserID(r *http.Request, userID uuid.UUID) {
	ctx := context.WithValue(r.Context(), sessionUserIDKey, userID)
	*r = *r.WithContext(ctx)
}
