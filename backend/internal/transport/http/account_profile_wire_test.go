package http

import (
	"crypto/ecdsa"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Wire-level tests (testing phase, task #07): the route registration in
// cmd/server/main.go is composition that no handler-level unit test can
// prove — these tests rebuild the exact mux pattern + middleware chain
// (RateLimit ∘ RequireSession ∘ AccountMeHandler) behind a real
// httptest.Server and exercise it with real ES256-verified tokens over
// actual HTTP requests. Handler-internal contract (shapes, statuses) is
// covered by account_profile_test.go; this file proves the wiring and the
// middleware boundary behave the same on the wire.

// newWireServer builds a server with the same registration shape as
// cmd/server/main.go:199-202, returning the server and a token minting
// closure signed by the verifier's key pair.
func newWireServer(t *testing.T, svc profileService) (*httptest.Server, func(sub string) string, *ecdsa.PrivateKey) {
	t.Helper()
	key, mint := newES256Signer(t)
	verifier := GoogleTokenVerifier(&key.PublicKey)

	mux := http.NewServeMux()
	mux.Handle("GET /account/me",
		RateLimit(100, 200)(
			RequireSession(verifier)(
				AccountMeHandler(svc))))
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return srv, mint, key
}

// getAccountMeWire issues a real GET /account/me against the wire server
// with the given Authorization header value ("" = no header).
func getAccountMeWire(t *testing.T, srv *httptest.Server, authHeader string) *http.Response {
	t.Helper()
	req, err := http.NewRequest(http.MethodGet, srv.URL+"/account/me", nil)
	if err != nil {
		t.Fatalf("build request: %v", err)
	}
	if authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("GET %s: %v", srv.URL+"/account/me", err)
	}
	t.Cleanup(func() { resp.Body.Close() })
	return resp
}

// decodeWireBody decodes a wire response body to a JSON object.
func decodeWireBody(t *testing.T, resp *http.Response) map[string]any {
	t.Helper()
	var m map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		t.Fatalf("decode response body: %v", err)
	}
	return m
}

// readWireBody slurps the full body for byte-equality assertions.
func readWireBody(t *testing.T, resp *http.Response) string {
	t.Helper()
	var buf [4096]byte
	n, _ := resp.Body.Read(buf[:])
	return string(buf[:n])
}

// Wire happy path: a real ES256 bearer token round-trips through the
// middleware chain to a 200 with the exact User-schema shape.
func TestAccountMeWire_Success_BearerToken(t *testing.T) {
	userID := uuid.New()
	svc := &stubProfileService{view: fullView("budi@example.com")}
	srv, mint, _ := newWireServer(t, svc)

	resp := getAccountMeWire(t, srv, "Bearer "+mint(userID.String()))
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", resp.StatusCode)
	}
	if ct := resp.Header.Get("Content-Type"); ct != "application/json" {
		t.Errorf("content-type = %q, want application/json", ct)
	}
	body := decodeWireBody(t, resp)
	assertUserShape(t, body)
	if body["email"] != "budi@example.com" {
		t.Errorf("email = %v, want budi@example.com", body["email"])
	}
	if !svc.called || svc.gotUserID != userID {
		t.Errorf("service received %s (called=%v), want session %s", svc.gotUserID, svc.called, userID)
	}
}

// Wire negative: no token is rejected by the middleware with 401 before
// the handler — over the wire, through the real chain.
func TestAccountMeWire_NoToken_Middleware401(t *testing.T) {
	svc := &stubProfileService{view: fullView("x@example.com")}
	srv, _, _ := newWireServer(t, svc)

	resp := getAccountMeWire(t, srv, "")
	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", resp.StatusCode)
	}
	if svc.called {
		t.Error("service must not be called without a token")
	}
	if body := decodeWireBody(t, resp); body["type"] != "https://kencleng.dev/errors/unauthorized" {
		t.Errorf("problem type = %v, want unauthorized", body["type"])
	}
}

// Wire negative: garbage and wrong-key tokens both 401 at the middleware.
func TestAccountMeWire_InvalidTokens_401(t *testing.T) {
	userID := uuid.New()
	svc := &stubProfileService{view: fullView("x@example.com")}
	srv, mint, realKey := newWireServer(t, svc)

	// A token signed by a different key parses fine cryptographically
	// elsewhere but must fail this verifier.
	otherKey, _ := newES256Signer(t)
	otherToken := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.RegisteredClaims{
		Subject:   userID.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
	})
	otherSigned, err := otherToken.SignedString(otherKey)
	if err != nil {
		t.Fatalf("sign foreign token: %v", err)
	}

	// Expired token from the REAL key (15 min past — beyond the 1 min leeway).
	expired := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.RegisteredClaims{
		Subject:   userID.String(),
		IssuedAt:  jwt.NewNumericDate(time.Now().Add(-time.Hour)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(-15 * time.Minute)),
	})
	expiredSigned, err := expired.SignedString(realKey)
	if err != nil {
		t.Fatalf("sign expired token: %v", err)
	}

	cases := []struct {
		name string
		auth string
	}{
		{"garbage", "Bearer not-a-jwt"},
		{"wrong-key", "Bearer " + otherSigned},
		{"expired", "Bearer " + expiredSigned},
		{"no-bearer-prefix", "Basic " + mint(userID.String())},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resp := getAccountMeWire(t, srv, tc.auth)
			if resp.StatusCode != http.StatusUnauthorized {
				t.Errorf("status = %d, want 401", resp.StatusCode)
			}
		})
	}
}

// Wire negative: the gone-user 401 (handler-emitted, past the middleware)
// must be byte-identical to the middleware's no-token 401 — the contract
// treats both as "session not usable" (R4/D5, now proven on the wire).
func TestAccountMeWire_GoneUser_401ByteIdenticalToNoToken(t *testing.T) {
	userID := uuid.New()
	svc := &stubProfileService{view: nil} // (nil, nil) = gone user
	srv, mint, _ := newWireServer(t, svc)

	noToken := getAccountMeWire(t, srv, "")
	noTokenBody := readWireBody(t, noToken)

	gone := getAccountMeWire(t, srv, "Bearer "+mint(userID.String()))
	if gone.StatusCode != noToken.StatusCode {
		t.Fatalf("status = %d, want %d (middleware's no-token status)", gone.StatusCode, noToken.StatusCode)
	}
	if goneBody := readWireBody(t, gone); goneBody != noTokenBody {
		t.Errorf("gone-user 401 body differs from no-token 401 body:\ngone:    %s\nno-token: %s", goneBody, noTokenBody)
	}
}

// Wire boundary: the mux pattern registers GET only — a wrong method gets
// the mux's 405, never the handler.
func TestAccountMeWire_WrongMethod_405(t *testing.T) {
	svc := &stubProfileService{view: fullView("x@example.com")}
	srv, _, _ := newWireServer(t, svc)

	req, err := http.NewRequest(http.MethodPost, srv.URL+"/account/me", nil)
	if err != nil {
		t.Fatalf("build request: %v", err)
	}
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("POST %s: %v", srv.URL+"/account/me", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("status = %d, want 405", resp.StatusCode)
	}
	if svc.called {
		t.Error("service must not be called on a method mismatch")
	}
}

// Wire edge case: a view whose Roles/AuthProviders are nil (the value
// toUserResponse's normalization branch exists for — the repository
// initializes both non-nil today, so no other test reaches this branch)
// must still serialize both arrays as [], never null, on the wire.
func TestAccountMeWire_NilSlices_NormalizedToEmptyArrays(t *testing.T) {
	userID := uuid.New()
	v := fullView("nil-slices@example.com")
	v.Roles = nil
	v.AuthProviders = nil
	svc := &stubProfileService{view: v}
	srv, mint, _ := newWireServer(t, svc)

	resp := getAccountMeWire(t, srv, "Bearer "+mint(userID.String()))
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", resp.StatusCode)
	}
	body := decodeWireBody(t, resp)
	assertUserShape(t, body)
	if roles, ok := body["roles"].([]any); !ok || len(roles) != 0 {
		t.Errorf("roles = %#v, want empty non-null array", body["roles"])
	}
	if ap, ok := body["auth_providers"].([]any); !ok || len(ap) != 0 {
		t.Errorf("auth_providers = %#v, want empty non-null array", body["auth_providers"])
	}
}
