package googleoauth

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// newTestClient wires a client at the mock endpoints with a generous timeout.
func newTestClient(t *testing.T, tokenSrvURL, jwksSrvURL string) *Client {
	t.Helper()
	return newClientWithEndpoints(testClientID, testClientSecret, testRedirectURI,
		"https://accounts.google.com/o/oauth2/v2/auth", tokenSrvURL, jwksSrvURL, 10*time.Second)
}

func TestAuthURL_ContainsStateNonceAndScope(t *testing.T) {
	c := NewClient(testClientID, testClientSecret, testRedirectURI)
	got := c.AuthURL("state-xyz", "nonce-xyz")
	for _, want := range []string{
		"client_id=" + testClientID,
		"redirect_uri=" + "http%3A%2F%2Flocalhost%3A8090%2Fauth%2Fgoogle%2Fcallback",
		"response_type=code",
		"scope=openid+email+profile",
		"state=state-xyz",
		"nonce=nonce-xyz",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("AuthURL missing %q in %q", want, got)
		}
	}
}

func TestExchangeCode_SendsConfiguredParamsAndParsesResponse(t *testing.T) {
	var gotBody url.Values
	tokenSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		gotBody, _ = url.ParseQuery(string(body))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"access_token":"at","id_token":"tok123","token_type":"Bearer","expires_in":3599}`))
	}))
	defer tokenSrv.Close()

	c := newTestClient(t, tokenSrv.URL, "http://unused-jwks")
	tr, err := c.ExchangeCode(context.Background(), "the-code")
	if err != nil {
		t.Fatalf("ExchangeCode: %v", err)
	}
	if tr.IDToken != "tok123" || tr.AccessToken != "at" {
		t.Errorf("parsed response wrong: %+v", tr)
	}
	want := map[string]string{
		"grant_type":    "authorization_code",
		"code":          "the-code",
		"client_id":     testClientID,
		"client_secret": testClientSecret,
		// R13: redirect_uri is the configured one — never from a request.
		"redirect_uri": testRedirectURI,
	}
	for k, v := range want {
		if got := gotBody.Get(k); got != v {
			t.Errorf("form[%q] = %q, want %q", k, got, v)
		}
	}
}

func TestExchangeCode_TimeoutReturnsCleanError(t *testing.T) {
	slow := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		time.Sleep(300 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer slow.Close()

	c := newClientWithEndpoints(testClientID, testClientSecret, testRedirectURI,
		"https://accounts.google.com/o/oauth2/v2/auth", slow.URL, "http://unused-jwks", 50*time.Millisecond)

	readLog := captureStdlog(t)
	_, err := c.ExchangeCode(context.Background(), "code")
	logged := readLog()

	if err == nil {
		t.Fatal("expected timeout error, got nil")
	}
	if !strings.Contains(err.Error(), "token exchange failed") {
		t.Errorf("error should be generic clean failure, got: %v", err)
	}
	if !strings.Contains(logged, "category=timeout") {
		t.Errorf("log should carry sanitized category=timeout, got: %s", logged)
	}
}

func TestExchangeCode_Non200_SanitizesErrorAndLog(t *testing.T) {
	const secretMarker = "SECRET-RESPONSE-BODY-MARKER"
	srv := tokenServer(t, http.StatusInternalServerError, `{"error":"`+secretMarker+`"}`)
	c := newTestClient(t, srv.URL, "http://unused-jwks")

	readLog := captureStdlog(t)
	_, err := c.ExchangeCode(context.Background(), "code")
	logged := readLog()

	if err == nil {
		t.Fatal("expected error on 500, got nil")
	}
	if strings.Contains(err.Error(), secretMarker) {
		t.Errorf("returned error leaked response body: %v", err)
	}
	if strings.Contains(logged, secretMarker) {
		t.Errorf("log leaked response body: %s", logged)
	}
	if !strings.Contains(err.Error(), "status 500") {
		t.Errorf("error should carry sanitized status only, got: %v", err)
	}
}

func TestExchangeCode_MissingIDTokenRejected(t *testing.T) {
	srv := tokenServer(t, http.StatusOK, `{"access_token":"at","token_type":"Bearer"}`)
	c := newTestClient(t, srv.URL, "http://unused-jwks")
	if _, err := c.ExchangeCode(context.Background(), "code"); err == nil {
		t.Fatal("expected error when id_token absent, got nil")
	}
}

func TestVerifyIDToken_ValidReturnsClaims(t *testing.T) {
	key := newTestKey(t)
	jwks, _ := jwksServer(t, key, nil, "kid-1", "", false)
	idTok := signRS256(t, key, "kid-1", validClaims(testNonce))
	tokenSrv := tokenServer(t, http.StatusOK, "")
	_ = tokenSrv

	c := newTestClient(t, "http://unused-token", jwks.URL)
	claims, err := c.VerifyIDToken(context.Background(), idTok, testNonce)
	if err != nil {
		t.Fatalf("VerifyIDToken: %v", err)
	}
	if claims.Email != testEmail || claims.Sub != testSub || claims.Nonce != testNonce {
		t.Errorf("claims mismatch: %+v", claims)
	}
	if claims.Iss != googleIssuer || claims.Aud != testClientID {
		t.Errorf("iss/aud mismatch: %+v", claims)
	}
	if claims.Exp.IsZero() {
		t.Errorf("exp should be populated: %+v", claims)
	}
}

func TestVerifyIDToken_NonceMismatchIsDistinguishable(t *testing.T) {
	key := newTestKey(t)
	jwks, _ := jwksServer(t, key, nil, "kid-1", "", false)
	idTok := signRS256(t, key, "kid-1", validClaims("a-different-nonce"))
	c := newTestClient(t, "http://unused-token", jwks.URL)

	err := verifyErr(c, idTok, testNonce)
	if err == nil {
		t.Fatal("expected nonce mismatch error, got nil")
	}
	if !errors.Is(err, ErrNonceMismatch) {
		t.Fatalf("expected ErrNonceMismatch, got: %v", err)
	}
}

// TestVerifyIDToken_GenericFailuresAreNotNonceMismatch — every failure other
// than replay must return a generic error (→ google_token_invalid), never
// ErrNonceMismatch (R26). Table-driven per backend/AGENTS.md §2.
func TestVerifyIDToken_GenericFailuresAreNotNonceMismatch(t *testing.T) {
	goodKey := newTestKey(t)
	forgedKey := newTestKey(t)

	tests := []struct {
		name   string
		makeID func(t *testing.T) string
	}{
		{
			name: "forged signature",
			makeID: func(t *testing.T) string {
				return signRS256(t, forgedKey, "kid-1", validClaims(testNonce))
			},
		},
		{
			name: "wrong issuer",
			makeID: func(t *testing.T) string {
				cl := validClaims(testNonce)
				cl["iss"] = "https://evil.example.com"
				return signRS256(t, goodKey, "kid-1", cl)
			},
		},
		{
			name: "wrong audience",
			makeID: func(t *testing.T) string {
				cl := validClaims(testNonce)
				cl["aud"] = "someone-elses-client-id"
				return signRS256(t, goodKey, "kid-1", cl)
			},
		},
		{
			name: "expired beyond leeway",
			makeID: func(t *testing.T) string {
				cl := validClaims(testNonce)
				cl["exp"] = time.Now().Add(-2 * time.Minute).Unix()
				return signRS256(t, goodKey, "kid-1", cl)
			},
		},
		{
			name: "missing exp",
			makeID: func(t *testing.T) string {
				cl := validClaims(testNonce)
				delete(cl, "exp")
				return signRS256(t, goodKey, "kid-1", cl)
			},
		},
		{
			name: "HS256 algorithm confusion attempt",
			makeID: func(t *testing.T) string {
				tok := jwt.NewWithClaims(jwt.SigningMethodHS256, validClaims(testNonce))
				s, err := tok.SignedString([]byte("attacker-chosen-secret"))
				if err != nil {
					t.Fatalf("sign hs256: %v", err)
				}
				return s
			},
		},
		{
			name: "missing kid header",
			makeID: func(t *testing.T) string {
				tok := jwt.NewWithClaims(jwt.SigningMethodRS256, validClaims(testNonce))
				s, err := tok.SignedString(goodKey)
				if err != nil {
					t.Fatalf("sign without kid: %v", err)
				}
				return s
			},
		},
		{
			name:   "garbage token",
			makeID: func(_ *testing.T) string { return "not.a.jwt" },
		},
	}

	for i, tc := range tests {
		kid := "kid-1"
		_ = i
		jwks, _ := jwksServer(t, goodKey, nil, kid, "", false)
		c := newTestClient(t, "http://unused-token", jwks.URL)
		t.Run(tc.name, func(t *testing.T) {
			err := verifyErr(c, tc.makeID(t), testNonce)
			if err == nil {
				t.Fatal("expected rejection, got nil error")
			}
			if errors.Is(err, ErrNonceMismatch) {
				t.Fatalf("generic failure misclassified as nonce replay: %v", err)
			}
		})
	}
}

func TestVerifyIDToken_ExpiryWithinLeewayAccepted(t *testing.T) {
	key := newTestKey(t)
	jwks, _ := jwksServer(t, key, nil, "kid-1", "", false)
	cl := validClaims(testNonce)
	cl["exp"] = time.Now().Add(-30 * time.Second).Unix() // inside 60s leeway
	idTok := signRS256(t, key, "kid-1", cl)
	c := newTestClient(t, "http://unused-token", jwks.URL)

	if err := verifyErr(c, idTok, testNonce); err != nil {
		t.Fatalf("token within clock-skew leeway should verify, got: %v", err)
	}
}

func TestVerifyIDToken_JWKSRefreshOnMiss(t *testing.T) {
	keyA := newTestKey(t) // known to Google initially
	keyB := newTestKey(t) // rotated-in key unknown to the initial cache

	// First fetch serves only keyA; later fetches include keyB.
	jwks, fetches := jwksServer(t, keyA, keyB, "kid-a", "kid-b", true)
	c := newTestClient(t, "http://unused-token", jwks.URL)

	// Warm the cache with fetch #1 (set = {kid-a}).
	idTokOld := signRS256(t, keyA, "kid-a", validClaims(testNonce))
	if err := verifyErr(c, idTokOld, testNonce); err != nil {
		t.Fatalf("warm-up verification failed: %v", err)
	}
	if got := fetches.Load(); got != 1 {
		t.Fatalf("warm-up should trigger exactly one JWKS fetch, got %d", got)
	}

	// Token signed by the rotated-in key B: cache miss must trigger exactly
	// one refetch and then succeed — never a permanent failure (R21).
	idTokNew := signRS256(t, keyB, "kid-b", validClaims(testNonce))
	if err := verifyErr(c, idTokNew, testNonce); err != nil {
		t.Fatalf("refresh-on-miss verification failed: %v", err)
	}
	if got := fetches.Load(); got != 2 {
		t.Fatalf("expected exactly one refresh after miss (total 2 fetches), got %d", got)
	}
}

func TestVerifyIDToken_JWKSServerDownReturnsGenericError(t *testing.T) {
	key := newTestKey(t)
	jwks, _ := jwksServer(t, key, nil, "kid-1", "", false)
	url := jwks.URL
	jwks.Close() // unreachable from here on
	c := newTestClient(t, "http://unused-token", url)

	idTok := signRS256(t, key, "kid-1", validClaims(testNonce))
	err := verifyErr(c, idTok, testNonce)
	if err == nil {
		t.Fatal("expected error when JWKS unreachable")
	}
	if errors.Is(err, ErrNonceMismatch) {
		t.Fatalf("JWKS outage misclassified as nonce replay: %v", err)
	}
}

func TestExchangeCode_DecodesFullResponse(t *testing.T) {
	resp := map[string]any{
		"access_token": "at-1",
		"id_token":     "id-1",
		"token_type":   "Bearer",
		"expires_in":   float64(3599),
	}
	b, _ := json.Marshal(resp)
	srv := tokenServer(t, http.StatusOK, string(b))
	c := newTestClient(t, srv.URL, "http://unused-jwks")

	tr, err := c.ExchangeCode(context.Background(), "code")
	if err != nil {
		t.Fatalf("ExchangeCode: %v", err)
	}
	if tr.ExpiresIn != 3599 || tr.TokenType != "Bearer" {
		t.Errorf("unexpected parse result: %+v", tr)
	}
}

func verifyErr(c *Client, idToken, nonce string) error {
	_, err := c.VerifyIDToken(context.Background(), idToken, nonce)
	return err
}
