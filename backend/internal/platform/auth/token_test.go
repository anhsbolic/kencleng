package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// testKeys generates a deterministic-per-test ECDSA keypair. Each call
// makes a fresh pair so "wrong key" scenarios are trivially constructible.
func testKeys(t *testing.T) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	t.Helper()
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("generate ecdsa key: %v", err)
	}
	return priv, &priv.PublicKey
}

func testSecret(t *testing.T) []byte {
	t.Helper()
	secret := make([]byte, 32)
	if _, err := rand.Read(secret); err != nil {
		t.Fatalf("read secret: %v", err)
	}
	return secret
}

var baseNow = time.Date(2026, 8, 26, 12, 0, 0, 0, time.UTC)

// ---- access token ----------------------------------------------------------

func TestAccessToken_RoundTrip(t *testing.T) {
	priv, pub := testKeys(t)
	userID := uuid.MustParse("01930f2e-7f1a-7c3a-9b1e-2f6a1c8e4d10")

	token, err := MintAccessToken(priv, userID, baseNow)
	if err != nil {
		t.Fatalf("mint: %v", err)
	}

	got, err := VerifyAccessToken(pub, token, baseNow.Add(time.Minute))
	if err != nil {
		t.Fatalf("verify: %v", err)
	}
	if got != userID {
		t.Errorf("subject = %s, want %s", got, userID)
	}
}

func TestAccessToken_ExpiresAfterTTL(t *testing.T) {
	priv, pub := testKeys(t)

	token, err := MintAccessToken(priv, uuid.New(), baseNow)
	if err != nil {
		t.Fatalf("mint: %v", err)
	}

	// Within TTL + leeway: valid.
	if _, err := VerifyAccessToken(pub, token, baseNow.Add(AccessTokenTTL+accessTokenLeeway-time.Second)); err != nil {
		t.Errorf("verify just inside leeway window failed: %v", err)
	}
	// Past TTL + leeway: invalid.
	if _, err := VerifyAccessToken(pub, token, baseNow.Add(AccessTokenTTL+accessTokenLeeway+time.Second)); !errors.Is(err, ErrInvalidToken) {
		t.Errorf("expired token: err = %v, want ErrInvalidToken", err)
	}
}

// TestAuthMiddleware_RejectsWrongSigningKey proves the cryptographic layer
// of the token-confusion mitigation: a token that does not verify under the
// expected key fails outright, regardless of its claims.
//
// Covers R6 (pending-token under access verifier) and R17 layer (a).
func TestAuthMiddleware_RejectsWrongSigningKey(t *testing.T) {
	_, pub := testKeys(t)
	wrongPriv, _ := testKeys(t)
	mfaSecret := testSecret(t)

	userID := uuid.New()

	// Access token signed by a DIFFERENT ECDSA key.
	foreignAccess, err := MintAccessToken(wrongPriv, userID, baseNow)
	if err != nil {
		t.Fatalf("mint foreign access: %v", err)
	}
	if _, err := VerifyAccessToken(pub, foreignAccess, baseNow); !errors.Is(err, ErrInvalidToken) {
		t.Errorf("access verifier accepted foreign-key ES256 token: err = %v", err)
	}

	// MFA-pending token (HS256, dedicated secret) presented to the ACCESS
	// verifier: must fail signature verification outright — this is the
	// cryptographic guarantee from Assumption A/B.
	pending, err := MintMFAPending(mfaSecret, userID, baseNow)
	if err != nil {
		t.Fatalf("mint pending: %v", err)
	}
	if _, err := VerifyAccessToken(pub, pending, baseNow); !errors.Is(err, ErrInvalidToken) {
		t.Errorf("access verifier accepted HS256 mfa_pending token (token confusion!): err = %v", err)
	}
}

// TestAuthMiddleware_RejectsNonAccessPurposeToken proves the logic layer of
// the mitigation on top of key separation: even a structurally valid token
// with the wrong or missing purpose claim is rejected.
//
// Covers R6 and R17 layer (b).
func TestAuthMiddleware_RejectsNonAccessPurposeToken(t *testing.T) {
	priv, pub := testKeys(t)
	userID := uuid.New()

	// Forge an ES256 token signed with the CORRECT private key but
	// purpose="mfa_pending" — signature verifies, purpose check must catch it.
	wrongPurpose := jwt.NewWithClaims(jwt.SigningMethodES256, purposeClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			IssuedAt:  jwt.NewNumericDate(baseNow),
			ExpiresAt: jwt.NewNumericDate(baseNow.Add(time.Hour)),
		},
		Purpose: purposeMFAPending,
	})
	signedWrongPurpose, err := wrongPurpose.SignedString(priv)
	if err != nil {
		t.Fatalf("forge wrong-purpose token: %v", err)
	}
	if _, err := VerifyAccessToken(pub, signedWrongPurpose, baseNow); !errors.Is(err, ErrInvalidToken) {
		t.Errorf("verifier accepted correct-key token with mfa_pending purpose: err = %v", err)
	}

	// Legacy shape: no purpose claim at all (feature-02-era tokens).
	noPurpose := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.RegisteredClaims{
		Subject:   userID.String(),
		IssuedAt:  jwt.NewNumericDate(baseNow),
		ExpiresAt: jwt.NewNumericDate(baseNow.Add(time.Hour)),
	})
	signedNoPurpose, err := noPurpose.SignedString(priv)
	if err != nil {
		t.Fatalf("sign legacy token: %v", err)
	}
	if _, err := VerifyAccessToken(pub, signedNoPurpose, baseNow); !errors.Is(err, ErrInvalidToken) {
		t.Errorf("verifier accepted legacy purpose-less token: err = %v", err)
	}
}

func TestAccessToken_MalformedInputs(t *testing.T) {
	priv, pub := testKeys(t)

	if _, err := MintAccessToken(nil, uuid.New(), baseNow); err == nil {
		t.Error("mint with nil private key should fail")
	}
	if _, err := VerifyAccessToken(pub, "", baseNow); !errors.Is(err, ErrInvalidToken) {
		t.Errorf("empty token: err = %v, want ErrInvalidToken", err)
	}
	for _, garbage := range []string{"not-a-jwt", "a.b.c", "eyJhbGciOiJIUzI1NiJ9.e30.sig"} {
		if _, err := VerifyAccessToken(pub, garbage, baseNow); !errors.Is(err, ErrInvalidToken) {
			t.Errorf("garbage token %q: err = %v, want ErrInvalidToken", garbage, err)
		}
	}
	// Unparseable subject.
	badSub := jwt.NewWithClaims(jwt.SigningMethodES256, purposeClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "not-a-uuid",
			IssuedAt:  jwt.NewNumericDate(baseNow),
			ExpiresAt: jwt.NewNumericDate(baseNow.Add(time.Hour)),
		},
		Purpose: purposeAccess,
	})
	signedBadSub, err := badSub.SignedString(priv)
	if err != nil {
		t.Fatalf("sign bad-subject: %v", err)
	}
	if _, err := VerifyAccessToken(pub, signedBadSub, baseNow); !errors.Is(err, ErrInvalidToken) {
		t.Errorf("non-UUID subject: err = %v, want ErrInvalidToken", err)
	}
	// Nil public key at verify time.
	token, _ := MintAccessToken(priv, uuid.New(), baseNow)
	if _, err := VerifyAccessToken(nil, token, baseNow); !errors.Is(err, ErrInvalidToken) {
		t.Errorf("nil public key: err = %v, want ErrInvalidToken", err)
	}
}

func TestAccessToken_AlgorithmConfusionRejected(t *testing.T) {
	_, pub := testKeys(t)
	secret := testSecret(t)

	// HS256 token whose keyid tricks nothing — algorithm pinned to ES256.
	hs := jwt.NewWithClaims(jwt.SigningMethodHS256, purposeClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   uuid.New().String(),
			IssuedAt:  jwt.NewNumericDate(baseNow),
			ExpiresAt: jwt.NewNumericDate(baseNow.Add(time.Hour)),
		},
		Purpose: purposeAccess,
	})
	signedHS, err := hs.SignedString(secret)
	if err != nil {
		t.Fatalf("sign hs256: %v", err)
	}
	if _, err := VerifyAccessToken(pub, signedHS, baseNow); !errors.Is(err, ErrInvalidToken) {
		t.Errorf("HS256 token accepted by access verifier (algorithm confusion): err = %v", err)
	}
}

// ---- mfa_pending token -----------------------------------------------------

func TestMFAPending_RoundTrip(t *testing.T) {
	secret := testSecret(t)
	userID := uuid.New()

	token, err := MintMFAPending(secret, userID, baseNow)
	if err != nil {
		t.Fatalf("mint: %v", err)
	}

	got, err := VerifyMFAPending(secret, token, baseNow.Add(4*time.Minute))
	if err != nil {
		t.Fatalf("verify within TTL: %v", err)
	}
	if got != userID {
		t.Errorf("subject = %s, want %s", got, userID)
	}
}

func TestMFAPending_StrictExpiry_NoLeeway(t *testing.T) {
	secret := testSecret(t)
	userID := uuid.New()

	token, err := MintMFAPending(secret, userID, baseNow)
	if err != nil {
		t.Fatalf("mint: %v", err)
	}

	// Exactly at expiry boundary + 1 second: rejected — no leeway on this
	// token class (tightness is deliberate; see MFAPendingTokenTTL).
	if _, err := VerifyMFAPending(secret, token, baseNow.Add(MFAPendingTokenTTL+time.Second)); !errors.Is(err, ErrInvalidToken) {
		t.Errorf("expired pending token: err = %v, want ErrInvalidToken", err)
	}
	// Just inside: accepted.
	if _, err := VerifyMFAPending(secret, token, baseNow.Add(MFAPendingTokenTTL-time.Second)); err != nil {
		t.Errorf("valid pending token rejected inside TTL: %v", err)
	}
}

func TestMFAPending_WrongSecretRejected(t *testing.T) {
	secret := testSecret(t)
	otherSecret := testSecret(t)

	token, err := MintMFAPending(secret, uuid.New(), baseNow)
	if err != nil {
		t.Fatalf("mint: %v", err)
	}
	if _, err := VerifyMFAPending(otherSecret, token, baseNow); !errors.Is(err, ErrInvalidToken) {
		t.Errorf("pending token verified under a different secret: err = %v", err)
	}
}

// TestMFAPending_CrossPurposeMatrix proves BOTH directions of the
// cross-purpose rejection: neither token type passes the other's verifier,
// even when every other property is legitimate.
func TestMFAPending_CrossPurposeMatrix(t *testing.T) {
	priv, pub := testKeys(t)
	secret := testSecret(t)
	userID := uuid.New()

	access, err := MintAccessToken(priv, userID, baseNow)
	if err != nil {
		t.Fatalf("mint access: %v", err)
	}
	pending, err := MintMFAPending(secret, userID, baseNow)
	if err != nil {
		t.Fatalf("mint pending: %v", err)
	}

	if _, err := VerifyMFAPending(secret, access, baseNow); !errors.Is(err, ErrInvalidToken) {
		t.Errorf("pending verifier accepted an access token: err = %v", err)
	}
	if _, err := VerifyAccessToken(pub, pending, baseNow); !errors.Is(err, ErrInvalidToken) {
		t.Errorf("access verifier accepted a pending token: err = %v", err)
	}
}

func TestMFAPending_MalformedAndValidation(t *testing.T) {
	secret := testSecret(t)

	if _, err := VerifyMFAPending(secret, "", baseNow); !errors.Is(err, ErrInvalidToken) {
		t.Errorf("empty token: err = %v, want ErrInvalidToken", err)
	}
	if _, err := VerifyMFAPending([]byte("short"), "x.y.z", baseNow); err == nil {
		t.Error("short secret should fail validation before parsing")
	}
	if _, err := MintMFAPending(nil, uuid.New(), baseNow); err == nil {
		t.Error("mint with nil secret should fail")
	}

	// Garbage inputs collapse to ErrInvalidToken.
	for _, garbage := range []string{"not-a-jwt", "a.b.c"} {
		if _, err := VerifyMFAPending(secret, garbage, baseNow); !errors.Is(err, ErrInvalidToken) {
			t.Errorf("garbage %q: err = %v, want ErrInvalidToken", garbage, err)
		}
	}
}

// ---- secret validation -----------------------------------------------------

func TestValidateMFAPendingSecret(t *testing.T) {
	valid32 := base64.StdEncoding.EncodeToString(make([]byte, 32))

	got, err := ValidateMFAPendingSecret(valid32)
	if err != nil {
		t.Fatalf("valid 32-byte secret rejected: %v", err)
	}
	if len(got) != 32 {
		t.Errorf("decoded length = %d, want 32", len(got))
	}

	cases := map[string]string{
		"empty":         "",
		"bad base64":    "!!!not-base64!!!",
		"too short":     base64.StdEncoding.EncodeToString(make([]byte, 16)),
		"too long":      base64.StdEncoding.EncodeToString(make([]byte, 64)),
		"empty decoded": base64.StdEncoding.EncodeToString([]byte{}),
	}
	for name, in := range cases {
		if _, err := ValidateMFAPendingSecret(in); err == nil {
			t.Errorf("%s: expected error, got none", name)
		}
	}
}
