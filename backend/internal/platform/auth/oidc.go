package auth

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

const googleIssuer = "https://accounts.google.com"

var ErrUntrustedIdentity = errors.New("untrusted provider identity")

type Identity struct {
	Issuer  string
	Subject string
}

// Protocol is the provider boundary. Implementations used in production must
// verify the ID-token signature and keys before returning an Identity.
type Protocol interface {
	AuthorizeURL(state, nonce, verifier string) string
	ExchangeAndVerify(ctx context.Context, code, nonce, verifier string) (Identity, error)
}

type GoogleProtocol struct {
	oauth    oauth2.Config
	verifier *oidc.IDTokenVerifier
	clientID string
}

func NewGoogleProtocol(ctx context.Context, clientID, clientSecret, redirectURL string) (*GoogleProtocol, error) {
	if clientID == "" || clientSecret == "" || redirectURL == "" {
		return nil, errors.New("Google client configuration is incomplete")
	}
	callback, err := url.Parse(redirectURL)
	if err != nil || callback.Host == "" || callback.Scheme == "" {
		return nil, errors.New("Google redirect URL is invalid")
	}
	provider, err := oidc.NewProvider(ctx, googleIssuer)
	if err != nil {
		return nil, fmt.Errorf("discover Google OIDC provider: %w", err)
	}
	return &GoogleProtocol{
		oauth: oauth2.Config{ClientID: clientID, ClientSecret: clientSecret, RedirectURL: redirectURL,
			Endpoint: provider.Endpoint(), Scopes: []string{oidc.ScopeOpenID}},
		verifier: provider.Verifier(&oidc.Config{ClientID: clientID, SupportedSigningAlgs: []string{"RS256"}}),
		clientID: clientID,
	}, nil
}

func (p *GoogleProtocol) AuthorizeURL(state, nonce, verifier string) string {
	return p.oauth.AuthCodeURL(state, oidc.Nonce(nonce), oauth2.S256ChallengeOption(verifier))
}

func (p *GoogleProtocol) ExchangeAndVerify(ctx context.Context, code, nonce, verifier string) (Identity, error) {
	if code == "" || nonce == "" || verifier == "" {
		return Identity{}, ErrUntrustedIdentity
	}
	token, err := p.oauth.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		return Identity{}, fmt.Errorf("exchange authorization code: %w", err)
	}
	raw, ok := token.Extra("id_token").(string)
	if !ok || raw == "" {
		return Identity{}, ErrUntrustedIdentity
	}
	idToken, err := p.verifier.Verify(ctx, raw)
	if err != nil {
		return Identity{}, fmt.Errorf("verify ID token: %w", err)
	}
	var claims struct {
		Issuer          string `json:"iss"`
		Subject         string `json:"sub"`
		AuthorizedParty string `json:"azp"`
		Nonce           string `json:"nonce"`
	}
	if err := idToken.Claims(&claims); err != nil {
		return Identity{}, fmt.Errorf("decode verified claims: %w", err)
	}
	if err := validateIdentity(idToken.Issuer, idToken.Subject, idToken.Audience, idToken.Expiry,
		claims.Issuer, claims.Subject, claims.AuthorizedParty, claims.Nonce, p.clientID, nonce, time.Now()); err != nil {
		return Identity{}, err
	}
	return Identity{Issuer: claims.Issuer, Subject: claims.Subject}, nil
}

func validateIdentity(verifiedIssuer, verifiedSubject string, audience []string, expiry time.Time,
	claimIssuer, claimSubject, azp, claimNonce, clientID, expectedNonce string, now time.Time) error {
	if verifiedIssuer != googleIssuer || claimIssuer != googleIssuer || verifiedSubject == "" || verifiedSubject != claimSubject ||
		len(audience) != 1 || audience[0] != clientID || (azp != "" && azp != clientID) ||
		claimNonce == "" || claimNonce != expectedNonce || !now.Before(expiry) {
		return ErrUntrustedIdentity
	}
	return nil
}
