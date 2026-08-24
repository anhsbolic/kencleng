package http

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/anhsbolic/kencleng/backend/internal/domain/account"
)

// ---- stub service ----------------------------------------------------------

// stubGoogleService records what handlers hand to the service and returns
// canned results, so the transport contract is tested in isolation.
type stubGoogleService struct {
	consentURL     string
	stateCookieVal string
	redirectErr    error

	callbackRes account.CallbackResult
	callbackErr error

	lastIntent        string
	lastSessionUserID *uuid.UUID
	lastCode          string
	lastState         string
	lastCookieValue   string
}

func (s *stubGoogleService) GoogleRedirect(_ context.Context, intent string, sessionUserID *uuid.UUID) (string, string, error) {
	s.lastIntent = intent
	s.lastSessionUserID = sessionUserID
	if s.redirectErr != nil {
		return "", "", s.redirectErr
	}
	return s.consentURL, s.stateCookieVal, nil
}

func (s *stubGoogleService) GoogleCallback(_ context.Context, code, state, cookieValue string) (account.CallbackResult, error) {
	s.lastCode, s.lastState, s.lastCookieValue = code, state, cookieValue
	return s.callbackRes, s.callbackErr
}

func newStubService() *stubGoogleService {
	return &stubGoogleService{
		consentURL:     "https://accounts.google.com/o/oauth2/v2/auth?state=xyz",
		stateCookieVal: "encoded-state-opaque",
	}
}

func newES256Signer(t *testing.T) (*ecdsa.PrivateKey, func(sub string) string) {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("generate key: %v", err)
	}
	return key, func(sub string) string {
		now := time.Now()
		tok := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.RegisteredClaims{
			Subject:   sub,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(10 * time.Minute)),
		})
		signed, err := tok.SignedString(key)
		if err != nil {
			t.Fatalf("sign: %v", err)
		}
		return signed
	}
}

// findSetCookie picks the Set-Cookie header for a given cookie name.
func findSetCookie(t *testing.T, resp *http.Response, name string) string {
	t.Helper()
	for _, c := range resp.Cookies() {
		if c.Name == name {
			return c.Value
		}
	}
	return ""
}

// ---- redirect handler ------------------------------------------------------

// TestGoogleRedirect_LinkReauthRequireAuth — R2's named regression test:
// link/reauth without a verifiable session must be rejected with 401 BEFORE
// any Google redirect is generated or contacted.
func TestGoogleRedirect_LinkReauthRequireAuth(t *testing.T) {
	svc := newStubService()
	h := GoogleRedirectHandler(svc, GoogleTokenVerifier(nil), false)

	for _, intent := range []string{"link", "reauth"} {
		req := httptest.NewRequest(http.MethodGet, "/auth/google/redirect?intent="+intent, nil)
		rec := httptest.NewRecorder()
		h(rec, req)

		if rec.Code != http.StatusUnauthorized {
			t.Errorf("intent=%s: want 401, got %d", intent, rec.Code)
		}
		if loc := rec.Header().Get("Location"); loc != "" {
			t.Errorf("intent=%s: no redirect may happen on rejection, got %q", intent, loc)
		}
		if svc.lastIntent != "" {
			t.Errorf("intent=%s: service must not be reached before authz passes", intent)
		}
	}
}

func TestGoogleRedirect_LoginNeedsNoAuth(t *testing.T) {
	svc := newStubService()
	h := GoogleRedirectHandler(svc, GoogleTokenVerifier(nil), true)

	req := httptest.NewRequest(http.MethodGet, "/auth/google/redirect?intent=login", nil)
	rec := httptest.NewRecorder()
	h(rec, req)

	if rec.Code != http.StatusFound {
		t.Fatalf("want 302, got %d", rec.Code)
	}
	if loc := rec.Header().Get("Location"); loc != svc.consentURL {
		t.Errorf("Location = %q, want consent URL", loc)
	}
	val := findSetCookie(t, rec.Result(), oauthStateCookieName)
	if val != svc.stateCookieVal {
		t.Errorf("state cookie value mismatch: %q", val)
	}
	for _, c := range rec.Result().Cookies() {
		if c.Name != oauthStateCookieName {
			continue
		}
		// R24 attribute contract.
		if !c.HttpOnly || !c.Secure {
			t.Errorf("state cookie must be HttpOnly+Secure (secure mode): %+v", c)
		}
		if c.SameSite != http.SameSiteLaxMode {
			t.Errorf("state cookie must be Lax (Strict breaks the cross-origin return): %v", c.SameSite)
		}
		if c.Path != "/auth/google" {
			t.Errorf("state cookie Path = %q, want /auth/google", c.Path)
		}
		if c.MaxAge != int(stateCookieMaxAge.Seconds()) {
			t.Errorf("state cookie MaxAge = %d, want %d", c.MaxAge, int(stateCookieMaxAge.Seconds()))
		}
	}
}

func TestGoogleRedirect_LinkWithBearerTokenPassesUserID(t *testing.T) {
	key, sign := newES256Signer(t)
	uid := uuid.New()
	svc := newStubService()
	h := GoogleRedirectHandler(svc, GoogleTokenVerifier(&key.PublicKey), false)

	req := httptest.NewRequest(http.MethodGet, "/auth/google/redirect?intent=link", nil)
	req.Header.Set("Authorization", "Bearer "+sign(uid.String()))
	rec := httptest.NewRecorder()
	h(rec, req)

	if rec.Code != http.StatusFound {
		t.Fatalf("want 302, got %d", rec.Code)
	}
	if svc.lastSessionUserID == nil || *svc.lastSessionUserID != uid {
		t.Errorf("verified user id must reach the service: %+v", svc.lastSessionUserID)
	}
}

func TestGoogleRedirect_LinkWithAccessTokenCookie(t *testing.T) {
	key, sign := newES256Signer(t)
	uid := uuid.New()
	svc := newStubService()
	h := GoogleRedirectHandler(svc, GoogleTokenVerifier(&key.PublicKey), false)

	req := httptest.NewRequest(http.MethodGet, "/auth/google/redirect?intent=reauth", nil)
	req.AddCookie(&http.Cookie{Name: accessTokenCookieName, Value: sign(uid.String())})
	rec := httptest.NewRecorder()
	h(rec, req)

	if rec.Code != http.StatusFound {
		t.Fatalf("cookie-carried token should authenticate, got %d", rec.Code)
	}
	if svc.lastSessionUserID == nil || *svc.lastSessionUserID != uid {
		t.Errorf("user id from cookie token wrong: %+v", svc.lastSessionUserID)
	}
}

func TestGoogleRedirect_TamperedTokenRejected(t *testing.T) {
	wrongKey, _ := newES256Signer(t)
	rightKey, sign := newES256Signer(t)
	_ = wrongKey
	svc := newStubService()
	h := GoogleRedirectHandler(svc, GoogleTokenVerifier(&rightKey.PublicKey), false)

	// Token signed by a different key — signature verification must fail.
	otherKey, otherSign := newES256Signer(t)
	_ = otherKey
	req := httptest.NewRequest(http.MethodGet, "/auth/google/redirect?intent=link", nil)
	req.Header.Set("Authorization", "Bearer "+otherSign(uuid.New().String()))
	rec := httptest.NewRecorder()
	h(rec, req)
	if rec.Code != http.StatusUnauthorized {
		t.Errorf("foreign-key token must be rejected, got %d", rec.Code)
	}
	_ = sign
}

func TestGoogleRedirect_InvalidIntent400(t *testing.T) {
	svc := newStubService()
	h := GoogleRedirectHandler(svc, GoogleTokenVerifier(nil), false)

	req := httptest.NewRequest(http.MethodGet, "/auth/google/redirect?intent=hijack", nil)
	rec := httptest.NewRecorder()
	h(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("R18: want 400, got %d", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "application/problem+json" {
		t.Errorf("errors use Problem Details, got %q", ct)
	}
}

// ---- callback handler ------------------------------------------------------

func TestGoogleCallback_ClearsStateCookieOnEveryOutcome(t *testing.T) {
	for _, tc := range []struct {
		name string
		res  account.CallbackResult
	}{
		{"error outcome", account.CallbackResult{Error: errStateMismatchLabel, RedirectURL: "http://f/login?error=" + errStateMismatchLabel}},
		{"success outcome", account.CallbackResult{RedirectURL: "http://f", AccessToken: "a", RefreshToken: "r"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			svc := &stubGoogleService{callbackRes: tc.res}
			h := GoogleCallbackHandler(svc, false)

			req := httptest.NewRequest(http.MethodGet, "/auth/google/callback?code=c&state=s", nil)
			req.AddCookie(&http.Cookie{Name: oauthStateCookieName, Value: "opaque"})
			rec := httptest.NewRecorder()
			h(rec, req)

			var cleared bool
			for _, c := range rec.Result().Cookies() {
				if c.Name == oauthStateCookieName && c.MaxAge < 0 {
					cleared = true
				}
			}
			if !cleared {
				t.Errorf("state cookie must be cleared (MaxAge<0) after consumption")
			}
		})
	}
}

const errStateMismatchLabel = "state_mismatch"

func TestGoogleCallback_MissingInputsStillHandled(t *testing.T) {
	// R19/R20 at the boundary: even with no params/cookie at all, the
	// response is a clean 302 error redirect — never a 500 or a bare body.
	svc := &stubGoogleService{callbackRes: account.CallbackResult{
		Error:       errStateMismatchLabel,
		RedirectURL: "http://localhost:3000/login?error=" + errStateMismatchLabel,
	}}
	h := GoogleCallbackHandler(svc, false)

	req := httptest.NewRequest(http.MethodGet, "/auth/google/callback", nil)
	rec := httptest.NewRecorder()
	h(rec, req)

	if rec.Code != http.StatusFound {
		t.Fatalf("want 302, got %d", rec.Code)
	}
	if loc := rec.Header().Get("Location"); !strings.Contains(loc, "?error=state_mismatch") {
		t.Errorf("Location should carry the error code: %q", loc)
	}
	if svc.lastCode != "" || svc.lastState != "" || svc.lastCookieValue != "" {
		t.Errorf("empty inputs must be forwarded as empty, got %+v", svc)
	}
}

func TestGoogleCallback_SuccessWritesAuthCookies(t *testing.T) {
	svc := &stubGoogleService{callbackRes: account.CallbackResult{
		RedirectURL:  "http://localhost:3000",
		AccessToken:  "access-jwt",
		RefreshToken: "refresh-plain",
	}}
	h := GoogleCallbackHandler(svc, true) // secure mode

	req := httptest.NewRequest(http.MethodGet, "/auth/google/callback?code=c&state=s", nil)
	req.AddCookie(&http.Cookie{Name: oauthStateCookieName, Value: "opaque"})
	rec := httptest.NewRecorder()
	h(rec, req)

	if rec.Code != http.StatusFound {
		t.Fatalf("want 302, got %d", rec.Code)
	}
	if loc := rec.Header().Get("Location"); loc != "http://localhost:3000" {
		t.Errorf("success lands on app root, got %q", loc)
	}

	var access, refresh *http.Cookie
	for _, c := range rec.Result().Cookies() {
		switch c.Name {
		case accessTokenCookieName:
			access = c
		case refreshTokenCookieName:
			refresh = c
		}
	}
	if access == nil || access.Value != "access-jwt" {
		t.Fatalf("access cookie missing/wrong: %+v", access)
	}
	if refresh == nil || refresh.Value != "refresh-plain" {
		t.Fatalf("refresh cookie missing/wrong: %+v", refresh)
	}
	if !access.HttpOnly || !access.Secure {
		t.Errorf("access cookie must be HttpOnly+Secure: %+v", access)
	}
	if !refresh.HttpOnly || !refresh.Secure {
		t.Errorf("refresh cookie must be HttpOnly+Secure: %+v", refresh)
	}
	if refresh.SameSite != http.SameSiteStrictMode {
		t.Errorf("refresh cookie must be Strict (never crosses sites): %v", refresh.SameSite)
	}
	if refresh.MaxAge != int(refreshTokenCookieTTL.Seconds()) {
		t.Errorf("refresh MaxAge = %d, want ~30d (%d)", refresh.MaxAge, int(refreshTokenCookieTTL.Seconds()))
	}
	if access.MaxAge != int(accessTokenCookieTTL.Seconds()) {
		t.Errorf("access MaxAge = %d, want %d", access.MaxAge, int(accessTokenCookieTTL.Seconds()))
	}
}

// TestGoogleCallback_ErrorCodesSurfaceVerbatim — all six error codes from
// the techplan §8 vocabulary must appear unmodified in the redirect target.
func TestGoogleCallback_ErrorCodesSurfaceVerbatim(t *testing.T) {
	codes := []string{
		"state_mismatch", "nonce_mismatch", "google_token_invalid",
		"google_unavailable", "google_email_conflict", "google_link_conflict",
	}
	for _, code := range codes {
		t.Run(code, func(t *testing.T) {
			target := "http://localhost:3000/login?error=" + code
			svc := &stubGoogleService{callbackRes: account.CallbackResult{
				Error: code, RedirectURL: target,
			}}
			h := GoogleCallbackHandler(svc, false)

			req := httptest.NewRequest(http.MethodGet, "/auth/google/callback?code=c&state=s", nil)
			rec := httptest.NewRecorder()
			h(rec, req)

			if loc := rec.Header().Get("Location"); !strings.HasSuffix(loc, "?error="+code) {
				t.Errorf("code must surface verbatim, got %q", loc)
			}
		})
	}
}

func TestGoogleCallback_ReauthSetsMarker(t *testing.T) {
	uid := uuid.New()
	svc := &stubGoogleService{callbackRes: account.CallbackResult{
		RedirectURL: "http://localhost:3000/account/security",
		Reauth:      true,
		UserID:      uid,
	}}
	h := GoogleCallbackHandler(svc, false)

	if CheckReauthMarker(uid) {
		t.Fatal("marker must not exist before the flow")
	}
	req := httptest.NewRequest(http.MethodGet, "/auth/google/callback?code=c&state=s", nil)
	rec := httptest.NewRecorder()
	h(rec, req)

	if rec.Code != http.StatusFound {
		t.Fatalf("want 302, got %d", rec.Code)
	}
	if !CheckReauthMarker(uid) {
		t.Error("reauth success must set a valid marker for the user")
	}

	// Auth cookies must NOT be written for reauth (no new session).
	for _, c := range rec.Result().Cookies() {
		if c.Name == accessTokenCookieName || c.Name == refreshTokenCookieName {
			t.Errorf("reauth must not set %s", c.Name)
		}
	}
}

func TestGoogleCallback_ServiceErrorIsGeneric500(t *testing.T) {
	svc := &stubGoogleService{callbackErr: context.DeadlineExceeded}
	h := GoogleCallbackHandler(svc, false)

	req := httptest.NewRequest(http.MethodGet, "/auth/google/callback?code=c&state=s", nil)
	rec := httptest.NewRecorder()
	h(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("internal service errors map to generic 500, got %d", rec.Code)
	}
	if strings.Contains(rec.Body.String(), "context deadline") {
		t.Errorf("internal error text leaked into response")
	}
}

func TestGoogleCallback_LogsNeverCarryTokens(t *testing.T) {
	logBuf := &strings.Builder{}
	old := log.Writer()
	log.SetOutput(logBuf)
	defer log.SetOutput(old)

	svc := &stubGoogleService{callbackRes: account.CallbackResult{
		RedirectURL:  "http://localhost:3000",
		AccessToken:  "SECRET-ACCESS-JWT-MARKER",
		RefreshToken: "SECRET-REFRESH-MARKER",
	}}
	h := GoogleCallbackHandler(svc, false)

	req := httptest.NewRequest(http.MethodGet, "/auth/google/callback?code=c&state=s", nil)
	req.AddCookie(&http.Cookie{Name: oauthStateCookieName, Value: "opaque-state"})
	rec := httptest.NewRecorder()
	h(rec, req)

	for _, secret := range []string{"SECRET-ACCESS-JWT-MARKER", "SECRET-REFRESH-MARKER", "opaque-state"} {
		if strings.Contains(logBuf.String(), secret) {
			t.Errorf("secret material leaked into logs:\n%s", logBuf.String())
		}
	}
}
