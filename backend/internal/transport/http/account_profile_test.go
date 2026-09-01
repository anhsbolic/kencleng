package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/anhsbolic/kencleng/backend/internal/domain/account"
)

// stubProfileService replays a scripted GetProfile result and records the
// userID it was called with, so the transport contract (status codes,
// response shape) is exercisable without the domain.
type stubProfileService struct {
	view      *account.LoginUserView
	err       error
	gotUserID uuid.UUID
	called    bool
}

func (s *stubProfileService) GetProfile(_ context.Context, userID uuid.UUID) (*account.LoginUserView, error) {
	s.called = true
	s.gotUserID = userID
	return s.view, s.err
}

// fullView builds a complete LoginUserView for shape assertions. Roles is
// deliberately empty ([]), AuthProviders has two entries — both cases the
// contract must serialize as arrays, never null.
func fullView(email string) *account.LoginUserView {
	return &account.LoginUserView{
		ID:            uuid.New(),
		Name:          "Budi",
		Email:         email,
		EmailVerified: true,
		Roles:         []string{},
		AuthProviders: []string{"email_password", "google"},
		MFAEnabled:    true,
		CreatedAt:     time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC),
	}
}

// assertUserShape asserts the exact User-schema key set and that the two
// array fields serialize as arrays (not null). Shared by the /account/me
// and login/mfa-login shape tests (R1, R8) so the endpoints can never
// silently diverge.
func assertUserShape(t *testing.T, userObj any) {
	t.Helper()
	user, ok := userObj.(map[string]any)
	if !ok {
		t.Fatalf("user object is not a JSON object: %#v", userObj)
	}
	want := map[string]bool{
		"id": true, "name": true, "email": true, "email_verified": true,
		"roles": true, "auth_providers": true, "mfa_enabled": true, "created_at": true,
	}
	for k := range user {
		if !want[k] {
			t.Errorf("unexpected key %q in user object", k)
		}
	}
	for k := range want {
		if _, ok := user[k]; !ok {
			t.Errorf("missing key %q in user object", k)
		}
	}
	if _, ok := user["roles"].([]any); !ok {
		t.Errorf("roles must serialize as an array, got %#v", user["roles"])
	}
	if _, ok := user["auth_providers"].([]any); !ok {
		t.Errorf("auth_providers must serialize as an array, got %#v", user["auth_providers"])
	}
}

// getAccountMe runs AccountMeHandler against target, injecting sessionID
// into the request context when non-nil (uuid.Nil means "no session").
func getAccountMe(t *testing.T, svc profileService, sessionID uuid.UUID, target string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, target, nil)
	if sessionID != uuid.Nil {
		injectSessionUserID(req, sessionID)
	}
	rec := httptest.NewRecorder()
	AccountMeHandler(svc)(rec, req)
	return rec
}

// R1 — success returns the exact User-schema shape with correct values.
func TestAccountMe_Success_UserShapeSnakeCase(t *testing.T) {
	userID := uuid.New()
	v := fullView("budi@example.com")
	rec := getAccountMe(t, &stubProfileService{view: v}, userID, "/account/me")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	body := decodeBody(t, rec)
	assertUserShape(t, body)
	if body["email"] != "budi@example.com" {
		t.Errorf("email = %v, want budi@example.com", body["email"])
	}
	if body["name"] != "Budi" {
		t.Errorf("name = %v, want Budi", body["name"])
	}
	if body["email_verified"] != true {
		t.Errorf("email_verified = %v, want true", body["email_verified"])
	}
	if body["mfa_enabled"] != true {
		t.Errorf("mfa_enabled = %v, want true", body["mfa_enabled"])
	}
	if body["id"] != v.ID.String() {
		t.Errorf("id = %v, want %s", body["id"], v.ID)
	}
}

// R2 (spec-named) — no session user in context → 401 before touching the service.
func TestAccountMe_RequiresAuth(t *testing.T) {
	stub := &stubProfileService{view: fullView("x@example.com")}
	rec := getAccountMe(t, stub, uuid.Nil, "/account/me")
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", rec.Code)
	}
	if stub.called {
		t.Error("service must not be called without a session user")
	}
	body := decodeBody(t, rec)
	if body["type"] != "https://kencleng.dev/errors/unauthorized" {
		t.Errorf("problem type = %v, want unauthorized", body["type"])
	}
}

// R3 (spec-named) — foreign identifiers in the request never influence the
// resource; the service always receives the session userID.
func TestAccountMe_NoIDParameter_SessionScoped(t *testing.T) {
	sessionID := uuid.New()
	foreign := uuid.New()

	cases := []struct {
		name   string
		target string
	}{
		{"query-param", "/account/me?user_id=" + foreign.String()},
		{"path-suffix", "/account/me/" + foreign.String()},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			stub := &stubProfileService{view: fullView("self@example.com")}
			rec := getAccountMe(t, stub, sessionID, tc.target)
			if rec.Code != http.StatusOK {
				t.Fatalf("status = %d, want 200", rec.Code)
			}
			if stub.gotUserID != sessionID {
				t.Errorf("service received %s, want session %s", stub.gotUserID, sessionID)
			}
			if body := decodeBody(t, rec); body["email"] != "self@example.com" {
				t.Errorf("email = %v, want the session user's own", body["email"])
			}
		})
	}

	// Junk body — the handler never reads r.Body, so a foreign identifier in
	// a body is equally inert.
	t.Run("junk-body", func(t *testing.T) {
		stub := &stubProfileService{view: fullView("self@example.com")}
		req := httptest.NewRequest(http.MethodGet, "/account/me",
			strings.NewReader(`{"user_id":"`+foreign.String()+`"}`))
		injectSessionUserID(req, sessionID)
		rec := httptest.NewRecorder()
		AccountMeHandler(stub)(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want 200", rec.Code)
		}
		if stub.gotUserID != sessionID {
			t.Errorf("service received %s, want session %s", stub.gotUserID, sessionID)
		}
	})
}

// R4 — the session references a gone user: (nil, nil) → 401 with the same
// generic problem shape as a missing token.
func TestAccountMe_UserGone_SessionInvalidated(t *testing.T) {
	rec := getAccountMe(t, &stubProfileService{view: nil}, uuid.New(), "/account/me")
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", rec.Code)
	}
	body := decodeBody(t, rec)
	if body["type"] != "https://kencleng.dev/errors/unauthorized" {
		t.Errorf("problem type = %v, want unauthorized", body["type"])
	}
	if body["title"] != "Unauthorized" {
		t.Errorf("title = %v, want Unauthorized", body["title"])
	}
}

// R5 — internal error maps to a generic 500, no error detail leaked.
func TestAccountMe_InternalError_Generic500(t *testing.T) {
	rec := getAccountMe(t, &stubProfileService{err: errors.New("db exploded")}, uuid.New(), "/account/me")
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500", rec.Code)
	}
	body := decodeBody(t, rec)
	if body["type"] != "https://kencleng.dev/problems/internal" {
		t.Errorf("problem type = %v, want problems/internal", body["type"])
	}
	if strings.Contains(strings.ToLower(asString(body["detail"])), "exploded") {
		t.Error("detail leaked internal error text")
	}
}

// R6 — the decrypted email never appears in any log line.
func TestAccountMe_LogsFreeOfPII(t *testing.T) {
	logs := captureLogOutput(t)
	stub := &stubProfileService{view: fullView("owner@example.com")}
	rec := getAccountMe(t, stub, uuid.New(), "/account/me")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if strings.Contains(logs(), "owner@example.com") {
		t.Errorf("log leaked decrypted email: %q", logs())
	}
}

// asString coerces a decoded JSON value to a string for the detail-leak
// assertion (nil-safe).
func asString(v any) string {
	s, _ := v.(string)
	return s
}
