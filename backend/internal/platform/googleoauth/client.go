// Package googleoauth provides a minimal Google OAuth 2.0 client:
// consent-screen URL building, authorization-code exchange, and id_token
// verification against Google's JWKS.
//
// It is shared infrastructure with no business rules — callers decide what
// to do with verification outcomes (see internal/domain/account). The only
// distinguished error is ErrNonceMismatch (replay detection); every other
// failure mode returns a generic error so callers cannot accidentally treat
// a forged token differently from one with bad claims.
//
// Logging discipline: this package logs facts and sanitized categories
// ("timeout", "http_error", status codes) only — never authorization codes,
// tokens, client secrets, or raw error strings, which may embed response
// bodies or credentials.
package googleoauth

import (
	"context"
	"crypto/rsa"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	defaultAuthURL  = "https://accounts.google.com/o/oauth2/v2/auth"
	defaultTokenURL = "https://oauth2.googleapis.com/token"
	defaultJWKSURL  = "https://www.googleapis.com/oauth2/v3/certs"

	googleIssuer = "accounts.google.com"
	oauthScope   = "openid email profile"

	jwksCacheTTL  = 15 * time.Minute
	idTokenLeeway = 60 * time.Second

	maxJWKSBodyBytes = 1 << 20
	maxTokenBodySize = 1 << 20
	httpTimeout      = 10 * time.Second
)

// ErrNonceMismatch is returned when an id_token's nonce claim does not match
// the value the caller stored before the redirect to Google. It is the only
// distinguishable verification failure: callers must map it to a replay
// error (nonce_mismatch), while every other verification error maps to a
// generic invalid-token error (google_token_invalid).
var ErrNonceMismatch = errors.New("googleoauth: id_token nonce claim does not match expected value")

// TokenResponse is the subset of Google's token-endpoint response that this
// application consumes.
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	IDToken     string `json:"id_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

// Claims carries the id_token claims consumed by the account domain after
// successful verification.
type Claims struct {
	Email string
	Sub   string
	Iss   string
	Aud   string
	Exp   time.Time
	Nonce string
}

// idTokenClaims mirrors the registered claims plus the Google-specific and
// nonce claims needed for verification. It embeds jwt.RegisteredClaims so
// golang-jwt validates iss/aud/exp/nbf automatically per parser options.
type idTokenClaims struct {
	Email string `json:"email"`
	Nonce string `json:"nonce"`
	jwt.RegisteredClaims
}

// jwkSet matches the JSON shape of https://www.googleapis.com/oauth2/v3/certs.
type jwkSet struct {
	Keys []jwk `json:"keys"`
}

type jwk struct {
	Kid string `json:"kid"`
	N   string `json:"n"`
	E   string `json:"e"`
}

// Client exchanges Google authorization codes for tokens and verifies the
// resulting id_tokens against Google's JWKS. It is safe for concurrent use
// and should be constructed once at startup and reused.
type Client struct {
	httpClient   *http.Client
	clientID     string
	clientSecret string
	redirectURI  string
	scope        string
	authURL      string
	tokenURL     string
	jwksURL      string

	mu            sync.Mutex
	jwksCache     map[string]*rsa.PublicKey
	jwksFetchedAt time.Time
}

// NewClient constructs a client with Google's production endpoints and an
// HTTP client with an explicit timeout. The returned Client is safe for
// concurrent use; construct it once and reuse it — never build one per call.
func NewClient(clientID, clientSecret, redirectURI string) *Client {
	return newClientWithEndpoints(clientID, clientSecret, redirectURI,
		defaultAuthURL, defaultTokenURL, defaultJWKSURL, httpTimeout)
}

// newClientWithEndpoints is the test seam: injectable endpoints and timeout
// so tests can point at httptest servers without touching production paths.
func newClientWithEndpoints(clientID, clientSecret, redirectURI, authURL, tokenURL, jwksURL string, timeout time.Duration) *Client {
	return &Client{
		httpClient:   &http.Client{Timeout: timeout},
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURI:  redirectURI,
		scope:        oauthScope,
		authURL:      authURL,
		tokenURL:     tokenURL,
		jwksURL:      jwksURL,
	}
}

// AuthURL builds the Google consent-screen URL for a redirect response. The
// state parameter round-trips through Google untouched (CSRF protection);
// the nonce parameter makes Google embed the value into the returned
// id_token (replay protection).
func (c *Client) AuthURL(state, nonce string) string {
	q := url.Values{}
	q.Set("client_id", c.clientID)
	q.Set("redirect_uri", c.redirectURI)
	q.Set("response_type", "code")
	q.Set("scope", c.scope)
	q.Set("state", state)
	q.Set("nonce", nonce)
	return c.authURL + "?" + q.Encode()
}

// ExchangeCode posts the authorization code to Google's token endpoint and
// returns the parsed response. The redirect_uri sent here is always the one
// configured at construction time — never taken from a request (open-redirect
// defense). On transport failure the returned error carries no response body;
// callers map any error from this method to a generic "Google unavailable"
// outcome.
func (c *Client) ExchangeCode(ctx context.Context, code string) (*TokenResponse, error) {
	form := url.Values{}
	form.Set("grant_type", "authorization_code")
	form.Set("code", code)
	form.Set("client_id", c.clientID)
	form.Set("client_secret", c.clientSecret)
	form.Set("redirect_uri", c.redirectURI)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.tokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("googleoauth: build token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("googleoauth: token exchange failed category=%s", categorizeError(err))
		return nil, fmt.Errorf("googleoauth: token exchange failed")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, maxTokenBodySize))
		log.Printf("googleoauth: token exchange returned non-200 status=%d", resp.StatusCode)
		return nil, fmt.Errorf("googleoauth: token exchange returned status %d", resp.StatusCode)
	}

	var tr TokenResponse
	if err := json.NewDecoder(io.LimitReader(resp.Body, maxTokenBodySize)).Decode(&tr); err != nil {
		return nil, fmt.Errorf("googleoauth: decode token response: %w", err)
	}
	if tr.IDToken == "" {
		return nil, errors.New("googleoauth: token response contains no id_token")
	}
	return &tr, nil
}

// VerifyIDToken parses and verifies an id_token issued by Google: RS256
// signature against Google's JWKS, issuer accounts.google.com, audience the
// configured client ID, expiry with small clock-skew leeway, and finally a
// constant-time nonce comparison against expectedNonce.
//
// A replayed nonce yields ErrNonceMismatch; every other verification failure
// (signature, issuer, audience, expiry, malformed token, unknown signing key)
// returns a wrapped generic error.
func (c *Client) VerifyIDToken(ctx context.Context, idToken, expectedNonce string) (*Claims, error) {
	parser := jwt.NewParser(
		jwt.WithValidMethods([]string{"RS256"}),
		jwt.WithIssuer(googleIssuer),
		jwt.WithAudience(c.clientID),
		jwt.WithExpirationRequired(),
		jwt.WithLeeway(idTokenLeeway),
	)

	claims := &idTokenClaims{}
	_, err := parser.ParseWithClaims(idToken, claims, func(t *jwt.Token) (any, error) {
		kid, _ := t.Header["kid"].(string)
		if kid == "" {
			return nil, errors.New("googleoauth: id_token header has no kid")
		}
		return c.signingKey(ctx, kid)
	})
	if err != nil {
		return nil, fmt.Errorf("googleoauth: verify id_token: %w", err)
	}

	if subtle.ConstantTimeCompare([]byte(claims.Nonce), []byte(expectedNonce)) != 1 {
		return nil, ErrNonceMismatch
	}

	aud := ""
	if len(claims.Audience) > 0 {
		aud = claims.Audience[0]
	}
	var exp time.Time
	if claims.ExpiresAt != nil {
		exp = claims.ExpiresAt.Time
	}
	return &Claims{
		Email: claims.Email,
		Sub:   claims.Subject,
		Iss:   claims.Issuer,
		Aud:   aud,
		Exp:   exp,
		Nonce: claims.Nonce,
	}, nil
}

// signingKey resolves the RSA public key for a key ID. It refreshes the JWKS
// cache when it is stale OR when the kid is absent (refresh-on-miss), then
// retries once — a cache miss never becomes a permanent failure while Google
// serves the key.
func (c *Client) signingKey(ctx context.Context, kid string) (*rsa.PublicKey, error) {
	if key := c.cachedKey(kid); key != nil && !c.cacheStale() {
		return key, nil
	}
	if err := c.refreshJWKS(ctx); err != nil {
		return nil, err
	}
	if key := c.cachedKey(kid); key != nil {
		return key, nil
	}
	return nil, fmt.Errorf("googleoauth: no JWKS signing key for kid %q after refresh", kid)
}

func (c *Client) cachedKey(kid string) *rsa.PublicKey {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.jwksCache[kid]
}

func (c *Client) cacheStale() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return time.Since(c.jwksFetchedAt) > jwksCacheTTL
}

// refreshJWKS fetches Google's current JWKS document and rebuilds the cache.
func (c *Client) refreshJWKS(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.jwksURL, nil)
	if err != nil {
		return fmt.Errorf("googleoauth: build JWKS request: %w", err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("googleoauth: JWKS fetch failed category=%s", categorizeError(err))
		return errors.New("googleoauth: JWKS fetch failed")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, maxJWKSBodyBytes))
		log.Printf("googleoauth: JWKS fetch returned non-200 status=%d", resp.StatusCode)
		return fmt.Errorf("googleoauth: JWKS fetch returned status %d", resp.StatusCode)
	}

	var set jwkSet
	if err := json.NewDecoder(io.LimitReader(resp.Body, maxJWKSBodyBytes)).Decode(&set); err != nil {
		return fmt.Errorf("googleoauth: decode JWKS: %w", err)
	}

	keys := make(map[string]*rsa.PublicKey, len(set.Keys))
	for _, k := range set.Keys {
		pub, err := rsaKeyFromJWK(k)
		if err != nil {
			continue // skip malformed entries; a healthy JWKS has usable keys
		}
		keys[k.Kid] = pub
	}

	c.mu.Lock()
	c.jwksCache = keys
	c.jwksFetchedAt = time.Now()
	c.mu.Unlock()
	return nil
}

// rsaKeyFromJWK converts a base64url-encoded RSA JWK into a public key.
func rsaKeyFromJWK(k jwk) (*rsa.PublicKey, error) {
	if k.N == "" || k.E == "" {
		return nil, errors.New("googleoauth: JWK missing n or e")
	}
	nb, err := base64.RawURLEncoding.DecodeString(k.N)
	if err != nil {
		return nil, fmt.Errorf("googleoauth: decode JWK modulus: %w", err)
	}
	eb, err := base64.RawURLEncoding.DecodeString(k.E)
	if err != nil {
		return nil, fmt.Errorf("googleoauth: decode JWK exponent: %w", err)
	}
	e := 0
	for _, b := range eb {
		e = e<<8 | int(b)
	}
	return &rsa.PublicKey{N: new(big.Int).SetBytes(nb), E: e}, nil
}

// categorizeError reduces an outbound-call error to a log-safe category.
// Raw transport errors can embed URLs or response fragments; only the fact
// of a timeout vs other failure is worth logging.
func categorizeError(err error) string {
	var ne net.Error
	if errors.As(err, &ne) && ne.Timeout() {
		return "timeout"
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return "timeout"
	}
	return "http_error"
}
