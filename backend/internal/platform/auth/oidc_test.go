package auth

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	jose "github.com/go-jose/go-jose/v4"
	"golang.org/x/oauth2"
)

type fixtureTransport func(*http.Request) (*http.Response, error)

func (f fixtureTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

// These are signed-token and simulated-provider fixtures, not real Google login.
func TestGoogleProtocolNegatives(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	otherKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name     string
		edit     func(map[string]any)
		alg      jose.SignatureAlgorithm
		badKey   bool
		noKeys   bool
		noToken  bool
		provider bool
		valid    bool
	}{
		{name: "valid signed fixture", valid: true},
		{name: "wrong issuer", edit: func(c map[string]any) { c["iss"] = "https://other.example" }},
		{name: "Google issuer alias", edit: func(c map[string]any) { c["iss"] = "accounts.google.com" }},
		{name: "wrong audience", edit: func(c map[string]any) { c["aud"] = "other-client" }},
		{name: "multiple audiences", edit: func(c map[string]any) { c["aud"] = []string{"fixture-client", "other"} }},
		{name: "wrong azp", edit: func(c map[string]any) { c["azp"] = "other-client" }},
		{name: "wrong nonce", edit: func(c map[string]any) { c["nonce"] = "other-nonce" }},
		{name: "missing nonce", edit: func(c map[string]any) { delete(c, "nonce") }},
		{name: "expired", edit: func(c map[string]any) { c["exp"] = time.Now().Add(-time.Hour).Unix() }},
		{name: "empty subject", edit: func(c map[string]any) { c["sub"] = "" }},
		{name: "forged signature", badKey: true},
		{name: "unapproved algorithm", alg: jose.RS384},
		{name: "missing verification key", noKeys: true},
		{name: "missing ID token", noToken: true},
		{name: "provider failure", provider: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			claims := map[string]any{"iss": googleIssuer, "sub": "CaseSensitiveSubject", "aud": "fixture-client",
				"exp": time.Now().Add(time.Hour).Unix(), "iat": time.Now().Unix(), "nonce": "fixture-nonce", "azp": "fixture-client",
				"email": "spoof@example.invalid", "hd": "spoof.example.invalid", "person_id": "attacker-body-person"}
			if tc.edit != nil {
				tc.edit(claims)
			}
			alg := tc.alg
			if alg == "" {
				alg = jose.RS256
			}
			signingKey := key
			if tc.badKey {
				signingKey = otherKey
			}
			signer, err := jose.NewSigner(jose.SigningKey{Algorithm: alg, Key: signingKey}, nil)
			if err != nil {
				t.Fatal(err)
			}
			payload, err := json.Marshal(claims)
			if err != nil {
				t.Fatal(err)
			}
			signed, err := signer.Sign(payload)
			if err != nil {
				t.Fatal(err)
			}
			raw, err := signed.CompactSerialize()
			if err != nil {
				t.Fatal(err)
			}
			client := &http.Client{Transport: fixtureTransport(func(r *http.Request) (*http.Response, error) {
				fixtureResponse := func(value any) (*http.Response, error) {
					body, err := json.Marshal(value)
					if err != nil {
						return nil, err
					}
					return &http.Response{StatusCode: 200, Header: http.Header{"Content-Type": {"application/json"}}, Body: io.NopCloser(strings.NewReader(string(body)))}, nil
				}
				switch r.URL.Path {
				case "/.well-known/openid-configuration":
					return fixtureResponse(map[string]any{"issuer": googleIssuer,
						"authorization_endpoint": "https://provider.example/authorize", "token_endpoint": "https://provider.example/token",
						"jwks_uri": "https://provider.example/jwks", "id_token_signing_alg_values_supported": []string{"RS256"}})
				case "/jwks":
					keys := []jose.JSONWebKey{{Key: &key.PublicKey, Use: "sig", Algorithm: "RS256"}}
					if tc.noKeys {
						keys = nil
					}
					return fixtureResponse(jose.JSONWebKeySet{Keys: keys})
				case "/token":
				default:
					return nil, errors.New("unexpected fixture provider request")
				}
				if err := r.ParseForm(); err != nil {
					t.Fatal(err)
				}
				if r.Form.Get("code_verifier") != "fixture-verifier" || r.Form.Get("grant_type") != "authorization_code" {
					t.Fatal("exchange omitted PKCE or authorization-code grant")
				}
				if tc.provider {
					return nil, errors.New("simulated provider unavailable")
				}
				response := map[string]any{"access_token": "discarded-fixture", "token_type": "Bearer"}
				if !tc.noToken {
					response["id_token"] = raw
				}
				body, err := json.Marshal(response)
				if err != nil {
					t.Fatal(err)
				}
				return &http.Response{StatusCode: 200, Header: http.Header{"Content-Type": {"application/json"}}, Body: io.NopCloser(strings.NewReader(string(body)))}, nil
			})}
			ctx := oidc.ClientContext(context.Background(), client)
			p, err := NewGoogleProtocol(ctx, "fixture-client", "fixture-client-secret", "https://app.example/api/auth/google/callback")
			if err != nil {
				t.Fatal("initialize simulated provider through production constructor")
			}
			identity, err := p.ExchangeAndVerify(ctx, "fixture-code", "fixture-nonce", "fixture-verifier")
			if tc.valid {
				if err != nil || identity != (Identity{Issuer: googleIssuer, Subject: "CaseSensitiveSubject"}) {
					t.Fatal("valid fixture did not map solely to issuer/subject")
				}
			} else if err == nil || identity != (Identity{}) {
				t.Fatal("negative fixture returned identity authority")
			}
		})
	}
}

func TestGoogleProtocolPKCEAndMissingExchangeInputs(t *testing.T) {
	p := &GoogleProtocol{oauth: oauth2.Config{ClientID: "fixture-client",
		Endpoint: oauth2.Endpoint{AuthURL: "https://provider.example/authorize"}, Scopes: []string{oidc.ScopeOpenID}}}
	u, err := url.Parse(p.AuthorizeURL("fixture-state", "fixture-nonce", "fixture-verifier"))
	if err != nil {
		t.Fatal(err)
	}
	q := u.Query()
	if q.Get("state") != "fixture-state" || q.Get("nonce") != "fixture-nonce" || q.Get("code_challenge_method") != "S256" || q.Get("code_challenge") != oauth2.S256ChallengeFromVerifier("fixture-verifier") {
		t.Fatal("authorization URL lost browser/state/nonce/PKCE binding")
	}
	for _, inputs := range [][3]string{{"", "nonce", "verifier"}, {"code", "", "verifier"}, {"code", "nonce", ""}} {
		if _, err := p.ExchangeAndVerify(context.Background(), inputs[0], inputs[1], inputs[2]); !errors.Is(err, ErrUntrustedIdentity) {
			t.Fatal("incomplete exchange inputs accepted")
		}
	}
}
