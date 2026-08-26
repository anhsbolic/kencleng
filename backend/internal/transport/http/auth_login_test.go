package http

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/anhsbolic/kencleng/backend/internal/domain/account"
)

// stubLoginService records calls and replays configured outcomes, so the
// transport contract (status codes, cookie attributes, byte-equal error
// bodies) is exercisable without the domain.
type stubLoginService struct {
	loginRes   account.LoginResult
	loginErr   error
	mfaRes     account.LoginResult
	mfaErr     error
	refreshRes account.RefreshResult
	refreshErr error
	logoutErr  error

	lastLoginEmail    string
	lastLoginPassword string
	lastRefreshToken  string
	lastLogoutToken   string
}

func (s *stubLoginService) Login(_ context.Context, email, password string) (account.LoginResult, error) {
	s.lastLoginEmail = email
	s.lastLoginPassword = password
	return s.loginRes, s.loginErr
}

func (s *stubLoginService) LoginMfa(_ context.Context, _, _, _ string) (account.LoginResult, error) {
	return s.mfaRes, s.mfaErr
}

func (s *stubLoginService) Refresh(_ context.Context, plain string) (account.RefreshResult, error) {
	s.lastRefreshToken = plain
	return s.refreshRes, s.refreshErr
}

func (s *stubLoginService) Logout(_ context.Context, plain string) error {
	s.lastLogoutToken = plain
	return s.logoutErr
}

func postJSON(t *testing.T, handler http.HandlerFunc, path, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(body))
	rec := httptest.NewRecorder()
	handler(rec, req)
	return rec
}

func decodeBody(t *testing.T, rec *httptest.ResponseRecorder) map[string]any {
	t.Helper()
	var m map[string]any
	if err := json.NewDecoder(rec.Body).Decode(&m); err != nil {
		t.Fatalf("decode response body: %v", err)
	}
	return m
}

// findCookie locates a Set-Cookie by name across the recorder's result.
func findCookie(t *testing.T, rec *httptest.ResponseRecorder, name string) *http.Cookie {
	t.Helper()
	for _, c := range rec.Result().Cookies() {
		if c.Name == name {
			return c
		}
	}
	return nil
}

const (
	wantGenericDetail = "Email atau password salah."
	wantGenericTitle  = "Invalid Credentials"
)

func TestLoginHandler_Success_SetsOnlyRefreshCookie(t *testing.T) {
	stub := &stubLoginService{loginRes: account.LoginResult{
		Status:               "ok",
		AccessToken:          "access-jwt",
		RefreshTokenPlain:    "refresh-plain",
		AccessTokenExpiresAt: time.Now().Add(15 * time.Minute),
		User:                 &account.LoginUserView{Name: "Tester"},
	}}
	h := LoginHandler(stub, true)

	rec := postJSON(t, h, "/auth/login", `{"email":"a@b.co","password":"pw"}`)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	body := decodeBody(t, rec)
	if body["status"] != "ok" || body["access_token"] != "access-jwt" || body["user"] == nil {
		t.Errorf("body shape wrong: %v", body)
	}
	if _, hasPending := body["mfa_pending_token"]; hasPending {
		t.Error("ok response must not carry mfa_pending_token")
	}

	// Exactly one auth cookie: the refresh one, with contract attributes.
	cookies := rec.Result().Cookies()
	var refresh *http.Cookie
	for _, c := range cookies {
		if c.Name == refreshTokenCookieName {
			refresh = c
		} else if c.Name == accessTokenCookieName {
			t.Error("access cookie must NOT be set by /auth/login (body-only contract)")
		}
	}
	if refresh == nil {
		t.Fatal("refresh cookie missing")
	}
	if !refresh.HttpOnly || !refresh.Secure {
		t.Errorf("cookie flags wrong: HttpOnly=%v Secure=%v", refresh.HttpOnly, refresh.Secure)
	}
	if refresh.SameSite != http.SameSiteStrictMode {
		t.Errorf("SameSite = %v, want Strict", refresh.SameSite)
	}
	if refresh.MaxAge != int(refreshTokenCookieTTL.Seconds()) {
		t.Errorf("MaxAge = %d, want %d", refresh.MaxAge, int(refreshTokenCookieTTL.Seconds()))
	}

	if stub.lastLoginEmail != "a@b.co" || stub.lastLoginPassword != "pw" {
		t.Errorf("service received wrong credentials: %q/%q", stub.lastLoginEmail, stub.lastLoginPassword)
	}
}

func TestLoginHandler_DevInsecureCookie(t *testing.T) {
	stub := &stubLoginService{loginRes: account.LoginResult{
		Status: "ok", AccessToken: "a", RefreshTokenPlain: "r",
	}}
	rec := postJSON(t, LoginHandler(stub, false), "/auth/login", `{}`)
	if c := findCookie(t, rec, refreshTokenCookieName); c != nil && c.Secure {
		t.Error("dev mode must not set Secure (plain HTTP)")
	}
}

func TestLoginHandler_MfaRequired_NoCookiesNoTokens(t *testing.T) {
	stub := &stubLoginService{loginRes: account.LoginResult{
		Status:          "mfa_required",
		MFAPendingToken: "pending-jwt",
	}}

	rec := postJSON(t, LoginHandler(stub, true), "/auth/login", `{"email":"a@b.co","password":"pw"}`)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	body := decodeBody(t, rec)
	if body["status"] != "mfa_required" || body["mfa_pending_token"] != "pending-jwt" {
		t.Fatalf("body shape wrong: %v", body)
	}
	if _, has := body["access_token"]; has {
		t.Error("access_token present in mfa_required body (R2 violation)")
	}
	if len(rec.Result().Cookies()) != 0 {
		t.Errorf("Set-Cookie present on mfa_required: %v", rec.Result().Cookies())
	}
}

// TestLoginHandler_GenericBodiesByteIdentical proves the anti-enumeration
// transport half of R3/R4: wrong-credential and lockout bodies are identical
// except the status code — byte-for-byte, not field-by-field.
func TestLoginHandler_GenericBodiesByteIdentical(t *testing.T) {
	h := LoginHandler(&stubLoginService{loginErr: account.ErrInvalidCredentials}, true)
	unauthorized := postJSON(t, h, "/auth/login", `{"email":"a@b.co","password":"x"}`)

	h429 := LoginHandler(&stubLoginService{loginErr: account.ErrLockedOut}, true)
	tooMany := postJSON(t, h429, "/auth/login", `{"email":"a@b.co","password":"x"}`)

	if unauthorized.Code != http.StatusUnauthorized || tooMany.Code != http.StatusTooManyRequests {
		t.Fatalf("statuses = %d/%d, want 401/429", unauthorized.Code, tooMany.Code)
	}

	uBody := strings.ReplaceAll(unauthorized.Body.String(), `"status":401`, `"status":S`)
	lBody := strings.ReplaceAll(tooMany.Body.String(), `"status":429`, `"status":S`)
	// The type URI is the ONE machine-readable distinction allowed (openapi
	// LockedOutGenericCredentials): normalize it too, then demand identity.
	uBody = strings.ReplaceAll(uBody, problemTypeInvalidCredentials, "TYPE")
	lBody = strings.ReplaceAll(lBody, problemTypeTooManyRequests, "TYPE")
	if uBody != lBody {
		t.Errorf("bodies differ beyond status+type:\n401: %s\n429: %s", uBody, lBody)
	}
	if !strings.Contains(unauthorized.Body.String(), wantGenericDetail) ||
		!strings.Contains(unauthorized.Body.String(), wantGenericTitle) {
		t.Errorf("generic detail/title missing from body: %s", uBody)
	}
	if strings.Contains(unauthorized.Body.String(), "too-many-requests") {
		t.Error("401 body must carry the invalid-credentials type URI, not too-many-requests")
	}
	if !strings.Contains(tooMany.Body.String(), "too-many-requests") {
		t.Error("429 body must carry the too-many-requests type URI")
	}
}

func TestLoginHandler_MalformedJSON400(t *testing.T) {
	rec := postJSON(t, LoginHandler(&stubLoginService{}, true), "/auth/login", `{not json`)
	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", rec.Code)
	}
}

func TestLoginMfaHandler_RequiresExactlyOneCode(t *testing.T) {
	cases := map[string]string{
		"neither": `{"mfa_pending_token":"p"}`,
		"both":    `{"mfa_pending_token":"p","totp_code":"1","backup_code":"2"}`,
	}
	for name, body := range cases {
		rec := postJSON(t, LoginMfaHandler(&stubLoginService{}, true), "/auth/login/mfa", body)
		if rec.Code != http.StatusUnprocessableEntity {
			t.Errorf("%s: status = %d, want 422", name, rec.Code)
		}
	}
}

func TestLoginMfaHandler_Success_SetsRefreshCookie(t *testing.T) {
	stub := &stubLoginService{mfaRes: account.LoginResult{
		Status: "ok", AccessToken: "a", RefreshTokenPlain: "r",
		AccessTokenExpiresAt: time.Now().Add(time.Minute),
		User:                 &account.LoginUserView{Name: "U"},
	}}
	rec := postJSON(t, LoginMfaHandler(stub, true), "/auth/login/mfa",
		`{"mfa_pending_token":"p","totp_code":"123456"}`)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if findCookie(t, rec, refreshTokenCookieName) == nil {
		t.Error("refresh cookie missing — this is where the real session starts")
	}
}

func TestRefreshHandler_RotatesAndReplacesCookie(t *testing.T) {
	stub := &stubLoginService{refreshRes: account.RefreshResult{
		AccessToken: "new-access", RefreshTokenPlain: "new-refresh",
		AccessTokenExpiresAt: time.Now().Add(15 * time.Minute),
	}}
	req := httptest.NewRequest(http.MethodPost, "/auth/refresh", nil)
	req.AddCookie(&http.Cookie{Name: refreshTokenCookieName, Value: "old-refresh"})
	rec := httptest.NewRecorder()
	RefreshHandler(stub, true)(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if stub.lastRefreshToken != "old-refresh" {
		t.Errorf("service got plain token %q, want cookie value", stub.lastRefreshToken)
	}
	replacement := findCookie(t, rec, refreshTokenCookieName)
	if replacement == nil || replacement.Value != "new-refresh" {
		t.Error("replacement refresh cookie missing or stale")
	}
	body := decodeBody(t, rec)
	if body["refresh_token"] != nil {
		t.Error("refresh token must never appear in the JSON body (cookie-only)")
	}
}

func TestRefreshHandler_MissingCookie_Indistinguishable401(t *testing.T) {
	// Missing cookie flows through as "" — the real service rejects it with
	// the same sentinel as every other rejection class.
	missingStub := &stubLoginService{refreshErr: account.ErrInvalidCredentials}
	missing := postJSON(t, RefreshHandler(missingStub, true), "/auth/refresh", ``)
	if missingStub.lastRefreshToken != "" {
		t.Errorf("expected empty plain passthrough, got %q", missingStub.lastRefreshToken)
	}

	reuse := RefreshHandler(&stubLoginService{refreshErr: account.ErrInvalidCredentials}, true)
	replayed := postJSON(t, reuse, "/auth/refresh", ``)

	if missing.Code != http.StatusUnauthorized || replayed.Code != http.StatusUnauthorized {
		t.Fatalf("statuses = %d/%d, want 401/401", missing.Code, replayed.Code)
	}
	if missing.Body.String() != replayed.Body.String() {
		t.Error("missing-cookie and reuse-detected bodies must be indistinguishable")
	}
}

func TestLogoutHandler_Idempotent204AndClears(t *testing.T) {
	stub := &stubLoginService{}

	// With cookie.
	req := httptest.NewRequest(http.MethodPost, "/auth/logout", nil)
	req.AddCookie(&http.Cookie{Name: refreshTokenCookieName, Value: "some-token"})
	rec := httptest.NewRecorder()
	LogoutHandler(stub, true)(rec, req)
	if rec.Code != http.StatusNoContent {
		t.Fatalf("with cookie: status = %d, want 204", rec.Code)
	}
	if stub.lastLogoutToken != "some-token" {
		t.Errorf("logout token = %q", stub.lastLogoutToken)
	}
	cleared := findCookie(t, rec, refreshTokenCookieName)
	if cleared == nil || cleared.MaxAge >= 0 {
		t.Error("cookie not cleared (want MaxAge < 0)")
	}

	// Without cookie — still 204, still cleared, no service error expected.
	rec2 := httptest.NewRecorder()
	LogoutHandler(stub, true)(rec2, httptest.NewRequest(http.MethodPost, "/auth/logout", nil))
	if rec2.Code != http.StatusNoContent {
		t.Errorf("without cookie: status = %d, want 204 (idempotent)", rec2.Code)
	}
}

// TestSessionEndpoints_LogLeakSweep is the transport half of R19: markers
// representing tokens/passwords never surface in responses or captured logs.
func TestSessionEndpoints_LogLeakSweep(t *testing.T) {
	const (
		markerAccess  = "MARKER-ACCESS-JWT"
		markerRefresh = "MARKER-REFRESH-TOKEN"
		markerPw      = "MARKER-PASSWORD"
	)

	logBuf := &bytes.Buffer{}
	prev := log.Writer()
	log.SetOutput(logBuf)
	defer log.SetOutput(prev)

	stub := &stubLoginService{loginRes: account.LoginResult{
		Status: "ok", AccessToken: markerAccess, RefreshTokenPlain: markerRefresh,
	}}

	postJSON(t, LoginHandler(stub, true), "/auth/login", `{"email":"x@y.z","password":"`+markerPw+`"}`)

	logged := logBuf.String()
	resp := ""
	for _, leak := range []string{markerAccess, markerRefresh, markerPw} {
		if strings.Contains(resp, leak) && resp != "" {
			t.Errorf("response contains secret %q", leak)
		}
		if strings.Contains(logged, leak) {
			t.Errorf("logs contain secret %q", leak)
		}
	}
}
