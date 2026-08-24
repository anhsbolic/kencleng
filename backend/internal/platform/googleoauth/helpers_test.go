package googleoauth

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	testClientID     = "test-client-id"
	testClientSecret = "test-client-secret"
	testRedirectURI  = "http://localhost:8090/auth/google/callback"
	testNonce        = "nonce-abc123"
	testEmail        = "user@example.com"
	testSub          = "google-sub-42"
)

// newTestKey generates an RSA-2048 private key for signing mock id_tokens.
func newTestKey(t *testing.T) *rsa.PrivateKey {
	t.Helper()
	k, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("generate rsa key: %v", err)
	}
	return k
}

// jwksServer serves a JWKS document built from the given keys and counts
// fetches so refresh-on-miss behavior can be asserted. If evolving is true,
// the FIRST request returns only primaryKey's entry and every later request
// returns both entries — simulating Google rotating in a new key mid-flow.
func jwksServer(t *testing.T, primaryKey, secondaryKey *rsa.PrivateKey, primaryKid, secondaryKid string, evolving bool) (*httptest.Server, *atomic.Int64) {
	t.Helper()
	var fetches atomic.Int64
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		n := fetches.Add(1)
		keys := map[string]*rsa.PublicKey{primaryKid: &primaryKey.PublicKey}
		if secondaryKey != nil && (!evolving || n > 1) {
			keys[secondaryKid] = &secondaryKey.PublicKey
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(jwksBody(keys)))
	}))
	t.Cleanup(srv.Close)
	return srv, &fetches
}

func jwksBody(keys map[string]*rsa.PublicKey) string {
	var b strings.Builder
	b.WriteString(`{"keys":[`)
	first := true
	for kid, pub := range keys {
		if !first {
			b.WriteString(",")
		}
		first = false
		eb := big.NewInt(int64(pub.E)).Bytes()
		fmt.Fprintf(&b, `{"kty":"RSA","kid":%q,"use":"sig","alg":"RS256","n":%q,"e":%q}`,
			kid,
			base64.RawURLEncoding.EncodeToString(pub.N.Bytes()),
			base64.RawURLEncoding.EncodeToString(eb))
	}
	b.WriteString(`]}`)
	return b.String()
}

func validClaims(nonce string) jwt.MapClaims {
	return jwt.MapClaims{
		"iss":            googleIssuer,
		"aud":            testClientID,
		"sub":            testSub,
		"email":          testEmail,
		"email_verified": true,
		"exp":            time.Now().Add(5 * time.Minute).Unix(),
		"iat":            time.Now().Unix(),
		"nonce":          nonce,
	}
}

func signRS256(t *testing.T, key *rsa.PrivateKey, kid string, claims jwt.MapClaims) string {
	t.Helper()
	tok := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tok.Header["kid"] = kid
	s, err := tok.SignedString(key)
	if err != nil {
		t.Fatalf("sign rs256: %v", err)
	}
	return s
}

// tokenServer returns a stub Google token endpoint responding with the given
// status and body.
func tokenServer(t *testing.T, status int, body string) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(srv.Close)
	return srv
}

const goodTokenResponseFmt = `{"access_token":"at","id_token":%q,"token_type":"Bearer","expires_in":3599}`

// captureStdlog redirects the standard logger into a thread-safe buffer for
// the duration of the test and returns a function reading everything logged
// so far.
func captureStdlog(t *testing.T) func() string {
	t.Helper()
	lw := &lockedWriter{b: &strings.Builder{}}
	old := log.Writer()
	log.SetOutput(lw)
	t.Cleanup(func() { log.SetOutput(old) })
	return func() string {
		lw.mu.Lock()
		defer lw.mu.Unlock()
		return lw.b.String()
	}
}

// lockedWriter serializes concurrent writes into a strings.Builder.
type lockedWriter struct {
	mu sync.Mutex
	b  *strings.Builder
}

func (lw *lockedWriter) Write(p []byte) (int, error) {
	lw.mu.Lock()
	defer lw.mu.Unlock()
	return lw.b.Write(p)
}
