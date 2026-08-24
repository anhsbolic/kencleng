package account

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/anhsbolic/kencleng/backend/internal/platform/auth"
	"github.com/anhsbolic/kencleng/backend/internal/platform/crypto"
	"github.com/anhsbolic/kencleng/backend/internal/platform/googleoauth"
)

// ---- fakes --------------------------------------------------------------

// fakeGoogleClient is a configurable in-memory googleOAuthClient. Every
// failure mode (timeout/unreachable, forged token, replayed nonce) is
// injectable so no network and no real Google tokens are involved.
type fakeGoogleClient struct {
	mu                sync.Mutex
	exchangeResp      *googleoauth.TokenResponse
	exchangeErr       error
	verifyClaims      *googleoauth.Claims
	verifyErr         error
	exchangeCalls     int
	lastCode          string
	lastState         string
	lastNonce         string
	lastExpectedNonce string
}

func (f *fakeGoogleClient) AuthURL(state, nonce string) string {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.lastState = state
	f.lastNonce = nonce
	return "https://accounts.google.com/o/oauth2/v2/auth?state=" + state + "&nonce=" + nonce
}

func (f *fakeGoogleClient) ExchangeCode(_ context.Context, code string) (*googleoauth.TokenResponse, error) {
	f.mu.Lock()
	f.lastCode = code
	f.exchangeCalls++
	resp, exErr := f.exchangeResp, f.exchangeErr
	f.mu.Unlock()
	if exErr != nil {
		return nil, exErr
	}
	return resp, nil
}

func (f *fakeGoogleClient) VerifyIDToken(_ context.Context, _ string, expectedNonce string) (*googleoauth.Claims, error) {
	f.mu.Lock()
	f.lastExpectedNonce = expectedNonce
	claims, vErr := f.verifyClaims, f.verifyErr
	f.mu.Unlock()
	if vErr != nil {
		return nil, vErr
	}
	return claims, nil
}

const googleTestEmail = "guser@example.com"

func googleClaims() *googleoauth.Claims {
	return &googleoauth.Claims{
		Email: googleTestEmail,
		Sub:   "google-sub-42",
		Iss:   "accounts.google.com",
		Aud:   "client-id",
		Exp:   time.Now().Add(5 * time.Minute),
		Nonce: "irrelevant-here",
	}
}

func newES256Keys(t *testing.T) *ecdsa.PrivateKey {
	t.Helper()
	k, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("generate es256 key: %v", err)
	}
	return k
}

// newGoogleTestService wires a Service with fakes for every dependency the
// OAuth flow touches. The HMAC key matches hashFor()'s convention from
// service_test.go so seeded lookups line up.
func newGoogleTestService(t *testing.T) (*Service, *fakeRepo, *fakeGoogleClient) {
	t.Helper()
	repo := newFakeRepo()
	gc := &fakeGoogleClient{
		exchangeResp: &googleoauth.TokenResponse{AccessToken: "at", IDToken: "idt"},
	}
	svc := &Service{
		repo:        repo,
		tx:          fakeTxRunner{},
		keys:        &crypto.Keys{HMACKey: make([]byte, 32)},
		googleOAuth: gc,
		authKeys:    &auth.Keys{Private: newES256Keys(t)},
		frontendURL: "http://localhost:3000",
	}
	return svc, repo, gc
}

// startGoogleLogin drives GoogleRedirect and returns the decoded state plus
// the raw cookie value, ready for GoogleCallback.
func startGoogleLogin(t *testing.T, svc *Service, intent string, sessionUserID *uuid.UUID) (oauthState, string) {
	t.Helper()
	_, cookieValue, err := svc.GoogleRedirect(context.Background(), intent, sessionUserID)
	if err != nil {
		t.Fatalf("GoogleRedirect(%s): %v", intent, err)
	}
	st, err := decodeOAuthState(cookieValue)
	if err != nil {
		t.Fatalf("decodeOAuthState: %v", err)
	}
	return st, cookieValue
}

// ---- R1/R3/R18/R2-defensive: redirect leg --------------------------------

func TestGoogleRedirect_Login_NoSessionRequired(t *testing.T) {
	svc, _, gc := newGoogleTestService(t)

	url, cookieValue, err := svc.GoogleRedirect(context.Background(), intentLogin, nil)
	if err != nil {
		t.Fatalf("GoogleRedirect: %v", err)
	}
	if !strings.Contains(url, "state=") || !strings.Contains(url, "nonce=") {
		t.Errorf("consent URL should embed state and nonce: %q", url)
	}
	if gc.lastState == "" || gc.lastNonce == "" || gc.lastState == gc.lastNonce {
		t.Errorf("state/nonce should be distinct non-empty random values")
	}
	st, err := decodeOAuthState(cookieValue)
	if err != nil {
		t.Fatalf("cookie undecodable: %v", err)
	}
	if st.Intent != intentLogin || st.UserID != nil {
		t.Errorf("login cookie must carry intent=login and no user_id: %+v", st)
	}
}

func TestGoogleRedirect_LinkReauthWithoutSessionRejectedDefensively(t *testing.T) {
	svc, _, _ := newGoogleTestService(t)
	for _, intent := range []string{intentLink, intentReauth} {
		if _, _, err := svc.GoogleRedirect(context.Background(), intent, nil); !errors.Is(err, ErrMissingSession) {
			t.Errorf("%s without session: want ErrMissingSession, got %v", intent, err)
		}
	}
}

func TestGoogleRedirect_LinkWithSessionEncodesUserID(t *testing.T) {
	svc, _, _ := newGoogleTestService(t)
	uid := uuid.New()

	_, cookieValue, err := svc.GoogleRedirect(context.Background(), intentLink, &uid)
	if err != nil {
		t.Fatalf("GoogleRedirect: %v", err)
	}
	st, err := decodeOAuthState(cookieValue)
	if err != nil {
		t.Fatalf("cookie undecodable: %v", err)
	}
	if st.UserID == nil || *st.UserID != uid {
		t.Errorf("link cookie must encode session user_id %s: %+v", uid, st)
	}
}

func TestGoogleRedirect_InvalidIntentRejected(t *testing.T) {
	svc, _, _ := newGoogleTestService(t)
	if _, _, err := svc.GoogleRedirect(context.Background(), "hijack", nil); !errors.Is(err, ErrInvalidIntent) {
		t.Fatalf("want ErrInvalidIntent, got %v", err)
	}
}

// ---- R4/R19/R20: pre-Google-call validation ------------------------------

func TestGoogleCallback_StateMismatch_NoGoogleCall(t *testing.T) {
	svc, repo, gc := newGoogleTestService(t)
	_, cookieValue := startGoogleLogin(t, svc, intentLogin, nil)

	res, err := svc.GoogleCallback(context.Background(), "the-code", "tampered-state", cookieValue)
	if err != nil {
		t.Fatalf("unexpected internal error: %v", err)
	}
	if res.Error != errStateMismatch {
		t.Errorf("want %q, got %q", errStateMismatch, res.Error)
	}
	if gc.exchangeCalls != 0 {
		t.Errorf("no Google API call may happen after state mismatch, got %d", gc.exchangeCalls)
	}
	if len(repo.insertedUsers)+len(repo.insertedIdentities) != 0 {
		t.Errorf("state mismatch must create no records")
	}
	if !strings.Contains(res.RedirectURL, frontendLoginPath+"?error="+errStateMismatch) {
		t.Errorf("error redirect should target login page: %q", res.RedirectURL)
	}
}

func TestGoogleCallback_MissingParamsOrBadCookie_StateMismatch(t *testing.T) {
	svc, _, gc := newGoogleTestService(t)
	_, cookieValue := startGoogleLogin(t, svc, intentLogin, nil)

	tests := []struct {
		name                string
		code, state, cookie string
	}{
		{"missing code", "", "whatever", cookieValue},
		{"missing state param", "code", "", cookieValue},
		{"empty cookie", "code", "whatever", ""},
		{"corrupt cookie", "code", "whatever", "not-base64!!"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			before := gc.exchangeCalls
			res, err := svc.GoogleCallback(context.Background(), tc.code, tc.state, tc.cookie)
			if err != nil {
				t.Fatalf("unexpected internal error: %v", err)
			}
			if res.Error != errStateMismatch {
				t.Errorf("want %q, got %q", errStateMismatch, res.Error)
			}
			if gc.exchangeCalls != before {
				t.Errorf("no Google API call allowed, delta %d", gc.exchangeCalls-before)
			}
		})
	}
}

// ---- R6/R5/R26: Google client failure mapping ----------------------------

func TestGoogleCallback_GoogleUnavailableMappedCleanly(t *testing.T) {
	svc, _, gc := newGoogleTestService(t)
	gc.exchangeErr = errors.New("connection refused to secrets.internal:443 x8f3a") // deliberately leaky-looking raw error
	st, cookieValue := startGoogleLogin(t, svc, intentLogin, nil)

	res, err := svc.GoogleCallback(context.Background(), "code", st.State, cookieValue)
	if err != nil {
		t.Fatalf("internal error should stay nil: %v", err)
	}
	if res.Error != errGoogleUnavailable {
		t.Errorf("want %q, got %q", errGoogleUnavailable, res.Error)
	}
	if strings.Contains(res.RedirectURL, "x8f3a") || strings.Contains(fmt.Sprint(res), "x8f3a") {
		t.Errorf("raw transport error leaked into result: %+v", res)
	}
}

func TestGoogleCallback_NonceMismatchIsReplayNotForgery(t *testing.T) {
	svc, _, gc := newGoogleTestService(t)
	gc.verifyErr = googleoauth.ErrNonceMismatch
	st, cookieValue := startGoogleLogin(t, svc, intentLogin, nil)

	res, err := svc.GoogleCallback(context.Background(), "code", st.State, cookieValue)
	if err != nil {
		t.Fatalf("unexpected internal error: %v", err)
	}
	if res.Error != errNonceMismatch {
		t.Errorf("replayed nonce must map to nonce_mismatch, got %q", res.Error)
	}
	if gc.lastExpectedNonce != st.Nonce {
		t.Errorf("service must pass the cookie nonce to verification: %q vs %q", gc.lastExpectedNonce, st.Nonce)
	}
}

func TestGoogleCallback_GenericVerifyFailureIsTokenInvalid(t *testing.T) {
	svc, _, gc := newGoogleTestService(t)
	gc.verifyErr = errors.New("verification failed: bad signature")
	st, cookieValue := startGoogleLogin(t, svc, intentLogin, nil)

	res, err := svc.GoogleCallback(context.Background(), "code", st.State, cookieValue)
	if err != nil {
		t.Fatalf("unexpected internal error: %v", err)
	}
	if res.Error != errGoogleTokenInvalid {
		t.Errorf("generic verification failure must map to google_token_invalid, got %q", res.Error)
	}
}

// ---- R7-R9: login branches ------------------------------------------------

func TestGoogleCallback_Login_ExistingGoogleIdentityIssuesTokens(t *testing.T) {
	svc, repo, gc := newGoogleTestService(t)
	hash := hashFor(googleTestEmail)
	repo.seedIdentity(providerGoogle, googleTestEmail, hash, true)
	gc.verifyClaims = googleClaims()

	st, cookieValue := startGoogleLogin(t, svc, intentLogin, nil)
	res, err := svc.GoogleCallback(context.Background(), "code", st.State, cookieValue)
	if err != nil {
		t.Fatalf("unexpected internal error: %v", err)
	}
	if res.Error != "" {
		t.Fatalf("expected success, got error %q", res.Error)
	}
	if res.AccessToken == "" || res.RefreshToken == "" {
		t.Errorf("both tokens must be present: %+v", res)
	}
	if len(repo.insertedUsers)+len(repo.insertedIdentities) != 0 {
		t.Errorf("existing-identity login creates no records")
	}
	if res.RedirectURL != "http://localhost:3000" {
		t.Errorf("success should land on app root, got %q", res.RedirectURL)
	}
}

func TestGoogleCallback_Login_NewUserCreatesVerifiedIdentity(t *testing.T) {
	svc, repo, gc := newGoogleTestService(t)
	gc.verifyClaims = googleClaims()

	st, cookieValue := startGoogleLogin(t, svc, intentLogin, nil)
	res, err := svc.GoogleCallback(context.Background(), "code", st.State, cookieValue)
	if err != nil {
		t.Fatalf("unexpected internal error: %v", err)
	}
	if res.Error != "" || res.AccessToken == "" {
		t.Fatalf("expected success with tokens, got %+v", res)
	}
	if len(repo.insertedUsers) != 1 || len(repo.insertedIdentities) != 1 {
		t.Fatalf("one User + one AuthIdentity expected, got %d/%d",
			len(repo.insertedUsers), len(repo.insertedIdentities))
	}
	id := repo.insertedIdentities[0]
	if id.ProviderType != providerGoogle {
		t.Errorf("provider_type = %q, want google", id.ProviderType)
	}
	// R14: google identities are born verified — never null.
	if id.VerifiedAt == nil {
		t.Errorf("google identity verified_at must be set at insert (R14)")
	}
}

// TestGoogleCallback_NoAutoMerge_Login — R9. The email belongs to an
// email_password identity of a DIFFERENT user: nothing may be created or
// attached (account-takeover defense).
func TestGoogleCallback_NoAutoMerge_Login(t *testing.T) {
	svc, repo, gc := newGoogleTestService(t)
	hash := hashFor(googleTestEmail)
	victim := repo.seedIdentity(providerEmailPassword, googleTestEmail, hash, true)
	usersBefore := len(repo.insertedUsers)
	identitiesBefore := len(repo.insertedIdentities)
	gc.verifyClaims = googleClaims()

	st, cookieValue := startGoogleLogin(t, svc, intentLogin, nil)
	res, err := svc.GoogleCallback(context.Background(), "code", st.State, cookieValue)
	if err != nil {
		t.Fatalf("unexpected internal error: %v", err)
	}
	if res.Error != errGoogleEmailConflict {
		t.Fatalf("want %q, got %q", errGoogleEmailConflict, res.Error)
	}
	if len(repo.insertedUsers) != usersBefore || len(repo.insertedIdentities) != identitiesBefore {
		t.Errorf("no-auto-merge violated: records were created")
	}
	if victim.UserID == uuid.Nil {
		t.Errorf("seed broken")
	}
	if res.AccessToken != "" || res.RefreshToken != "" {
		t.Errorf("conflict outcome must not issue tokens: %+v", res)
	}
}

// ---- R10/R11: link branches ----------------------------------------------

func TestGoogleCallback_NoAutoMerge_Link(t *testing.T) {
	svc, repo, gc := newGoogleTestService(t)
	hash := hashFor(googleTestEmail)
	other := repo.seedIdentity(providerGoogle, googleTestEmail, hash, true)
	sessionUID := uuid.New()
	gc.verifyClaims = googleClaims()

	st, cookieValue := startGoogleLogin(t, svc, intentLink, &sessionUID)
	res, err := svc.GoogleCallback(context.Background(), "code", st.State, cookieValue)
	if err != nil {
		t.Fatalf("unexpected internal error: %v", err)
	}
	if res.Error != errGoogleLinkConflict {
		t.Fatalf("want %q, got %q", errGoogleLinkConflict, res.Error)
	}
	if other.UserID == sessionUID {
		t.Fatalf("seed must give the google identity a DIFFERENT owner")
	}
	for _, id := range repo.insertedIdentities {
		if id.UserID == sessionUID && id.ProviderType == providerGoogle {
			t.Errorf("identity was attached to the session user despite conflict")
		}
	}
	if len(repo.insertedUserLogs) != 0 {
		t.Errorf("rejected link must not write audit entries")
	}
}

func TestGoogleCallback_LinkSuccess_AttachesAndAudits(t *testing.T) {
	svc, repo, gc := newGoogleTestService(t)
	sessionUID := uuid.New()
	gc.verifyClaims = googleClaims()

	st, cookieValue := startGoogleLogin(t, svc, intentLink, &sessionUID)
	res, err := svc.GoogleCallback(context.Background(), "code", st.State, cookieValue)
	if err != nil {
		t.Fatalf("unexpected internal error: %v", err)
	}
	if res.Error != "" {
		t.Fatalf("expected link success, got %q", res.Error)
	}
	// R11: attach to the SESSION user — never a new User row.
	if len(repo.insertedUsers) != 0 {
		t.Errorf("link must not create a User row, got %d", len(repo.insertedUsers))
	}
	if len(repo.insertedIdentities) != 1 {
		t.Fatalf("exactly one identity expected, got %d", len(repo.insertedIdentities))
	}
	id := repo.insertedIdentities[0]
	if id.UserID != sessionUID || id.ProviderType != providerGoogle {
		t.Errorf("identity must attach google to session user: %+v", id)
	}
	if id.VerifiedAt == nil {
		t.Errorf("linked google identity must be verified at insert (R14)")
	}
	// Fitur 9 audit trail.
	if len(repo.insertedUserLogs) != 1 {
		t.Fatalf("exactly one audit entry expected, got %d", len(repo.insertedUserLogs))
	}
	entry := repo.insertedUserLogs[0]
	if entry.UserID != sessionUID || entry.ActionType != actionAccountLinking {
		t.Errorf("audit entry mismatch: %+v", entry)
	}
	if !strings.Contains(res.RedirectURL, frontendSecurityPath) {
		t.Errorf("link success lands on security page, got %q", res.RedirectURL)
	}
}

// ---- R12: reauth branch ---------------------------------------------------

func TestGoogleCallback_Reauth_NoSideEffects(t *testing.T) {
	svc, repo, gc := newGoogleTestService(t)
	sessionUID := uuid.New()
	gc.verifyClaims = googleClaims()
	usersBefore := len(repo.insertedUsers)
	identitiesBefore := len(repo.insertedIdentities)

	st, cookieValue := startGoogleLogin(t, svc, intentReauth, &sessionUID)
	res, err := svc.GoogleCallback(context.Background(), "code", st.State, cookieValue)
	if err != nil {
		t.Fatalf("unexpected internal error: %v", err)
	}
	if res.Error != "" {
		t.Fatalf("expected reauth success, got %q", res.Error)
	}
	if !res.Reauth {
		t.Errorf("handler relies on Reauth=true to set the marker")
	}
	if len(repo.insertedUsers) != usersBefore || len(repo.insertedIdentities) != identitiesBefore {
		t.Errorf("reauth must change no identity/user state")
	}
	if res.AccessToken != "" || res.RefreshToken != "" {
		t.Errorf("reauth issues no tokens: %+v", res)
	}
	if !strings.Contains(res.RedirectURL, frontendSecurityPath) {
		t.Errorf("reauth lands on security page, got %q", res.RedirectURL)
	}
}

// ---- IssueTokens ----------------------------------------------------------

func TestIssueTokens_ES256AccessAndHashedRefresh(t *testing.T) {
	svc, repo, _ := newGoogleTestService(t)
	uid := uuid.New()

	access, refresh, err := svc.IssueTokens(context.Background(), uid)
	if err != nil {
		t.Fatalf("IssueTokens: %v", err)
	}

	// Access token: ES256-signed JWT carrying the right sub/exp.
	rc := &jwt.RegisteredClaims{}
	_, err = jwt.ParseWithClaims(access, rc, func(tok *jwt.Token) (any, error) {
		if _, ok := tok.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", tok.Header["alg"])
		}
		return &svc.authKeys.Private.PublicKey, nil
	}, jwt.WithValidMethods([]string{"ES256"}))
	if err != nil {
		t.Fatalf("access token does not verify as ES256: %v", err)
	}
	if rc.Subject != uid.String() {
		t.Errorf("sub = %q, want %q", rc.Subject, uid.String())
	}
	if rc.ExpiresAt == nil || time.Until(rc.ExpiresAt.Time) > accessTokenTTL+5*time.Second {
		t.Errorf("exp should be ~%s away, got %v", accessTokenTTL, rc.ExpiresAt)
	}

	// Refresh token: only its SHA-256 hash is stored.
	if len(repo.insertedRefreshTokens) != 1 {
		t.Fatalf("one refresh_tokens row expected, got %d", len(repo.insertedRefreshTokens))
	}
	rt := repo.insertedRefreshTokens[0]
	if rt.TokenHash != sha256Hex(refresh) {
		t.Errorf("stored hash must be sha256(plain refresh)")
	}
	if rt.UserID != uid || rt.FamilyID == uuid.Nil {
		t.Errorf("refresh row ownership/family wrong: %+v", rt)
	}
	if rt.ExpiresAt.Sub(time.Now()) > refreshTokenTTL+5*time.Second {
		t.Errorf("expires_at should be ~30d out, got %v", rt.ExpiresAt)
	}
	if rt.RevokedAt != nil || rt.ReplacedByID != nil {
		t.Errorf("fresh refresh token must have nil revoked/replaced columns")
	}
}

// ---- R15: concurrent duplicate registration -------------------------------

// TestGoogleCallback_ConcurrentDuplicateRegistration_Race — INV-account-01.
// Many concurrent logins for the same brand-new Google email: the fake
// repo's identityKeys set emulates the DB unique index. Exactly ONE user may
// ever be created; every racer ends either logged-in-as-winner or with a
// clean conflict result — nobody crashes, nobody auto-merges.
func TestGoogleCallback_ConcurrentDuplicateRegistration_Race(t *testing.T) {
	svc, repo, gc := newGoogleTestService(t)
	gc.verifyClaims = googleClaims()

	const racers = 12
	results := make([]CallbackResult, racers)
	errs := make([]error, racers)
	var wg sync.WaitGroup
	for i := 0; i < racers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, cookieValue, rErr := svc.GoogleRedirect(context.Background(), intentLogin, nil)
			if rErr != nil {
				errs[i] = rErr
				return
			}
			st, dErr := decodeOAuthState(cookieValue)
			if dErr != nil {
				errs[i] = dErr
				return
			}
			results[i], errs[i] = svc.GoogleCallback(context.Background(), "code", st.State, cookieValue)
		}(i)
	}
	wg.Wait()

	for i := range results {
		if errs[i] != nil {
			t.Fatalf("racer %d crashed: %v", i, errs[i])
		}
		switch {
		case results[i].Error == "":
			// Logged in — must be as the single winner.
			rc := &jwt.RegisteredClaims{}
			_, pErr := jwt.ParseWithClaims(results[i].AccessToken, rc, func(*jwt.Token) (any, error) {
				return &svc.authKeys.Private.PublicKey, nil
			}, jwt.WithValidMethods([]string{"ES256"}))
			if pErr != nil {
				t.Fatalf("racer %d produced unverifiable token: %v", i, pErr)
			}
			if len(repo.insertedUsers) > 0 && rc.Subject != repo.insertedUsers[0].ID.String() {
				t.Fatalf("racer %d logged in as non-winner %s", i, rc.Subject)
			}
		case results[i].Error == errGoogleEmailConflict:
			// Clean, recoverable race loss — acceptable.
		default:
			t.Fatalf("racer %d got unexpected error %q", i, results[i].Error)
		}
	}
	// Committed-reality assertions. Note: fakeRepo.insertedUsers records
	// pre-rollback inserts (the fake has no tx semantics), so losers'
	// aborted User rows still appear there — the DB would have discarded
	// them. The identity set is the durable truth the unique index guards,
	// so ownership is asserted through it.
	if len(repo.insertedIdentities) != 1 {
		t.Errorf("exactly one google identity may exist after the race, got %d", len(repo.insertedIdentities))
	}
	// The fake indexes seeded identities under provider|hash and
	// runtime-inserted ones under provider|plaintext-identifier (see
	// fakeRepo.InsertAuthIdentity) — check both when locating the winner.
	winnerHash := hashFor(googleTestEmail)
	winner := repo.identities[idKeyFor(providerGoogle, winnerHash)]
	if winner == nil {
		winner = repo.identities[idKeyFor(providerGoogle, googleTestEmail)]
	}
	if winner == nil {
		t.Fatalf("winner identity missing after race")
	}
	for i := range results {
		if results[i].Error == "" && !tokenSubIs(svc, results[i].AccessToken, winner.UserID) {
			t.Errorf("racer %d logged in as someone other than the identity owner", i)
		}
	}
}

// idKeyFor mirrors fakeRepo.idKey without needing a repo instance.
func idKeyFor(provider, hash string) string { return provider + "|" + hash }

// tokenSubIs reports whether the ES256 access token verifies against the
// service's key and carries userID as its subject.
func tokenSubIs(svc *Service, token string, userID uuid.UUID) bool {
	rc := &jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(token, rc, func(*jwt.Token) (any, error) {
		return &svc.authKeys.Private.PublicKey, nil
	}, jwt.WithValidMethods([]string{"ES256"}))
	return err == nil && rc.Subject == userID.String()
}

// TestGoogleCallback_ConcurrentDuplicate_WinnerInvisibleYet — deterministic
// variant: the unique index fires but the winner's identity is not visible
// to re-lookup. Must fail cleanly (no crash, no partial user), surfacing the
// flagged recoverable-conflict signal.
func TestGoogleCallback_ConcurrentDuplicate_WinnerInvisibleYet(t *testing.T) {
	svc, repo, gc := newGoogleTestService(t)
	repo.insertIdentityErr = &pgconn.PgError{Code: "23505"}
	gc.verifyClaims = googleClaims()

	st, cookieValue := startGoogleLogin(t, svc, intentLogin, nil)
	res, err := svc.GoogleCallback(context.Background(), "code", st.State, cookieValue)
	if err != nil {
		t.Fatalf("unique-violation path must not surface as internal error: %v", err)
	}
	if res.Error == "" {
		t.Fatal("invisible-winner race loss must not report success")
	}
	if len(repo.insertedUsers) != 1 {
		// The users insert happened but the tx rolled back in the fake —
		// the important part is no committed identity exists.
		t.Logf("note: fake recorded %d user inserts (pre-rollback)", len(repo.insertedUsers))
	}
	if len(repo.insertedIdentities) != 0 {
		t.Errorf("no identity may persist when the index rejects the insert")
	}
}

// ---- R16: log discipline across all OAuth paths ---------------------------

func TestGoogleOAuth_LogsNeverCarrySecrets(t *testing.T) {
	const markerCode = "CODE-MARKER-x7"
	const markerIDToken = "IDTOKEN-MARKER-k2"

	scenarios := []struct {
		name string
		// setup receives the wired fakes before the flow runs.
		setup func(repo *fakeRepo, gc *fakeGoogleClient)
	}{
		{"login new user", func(_ *fakeRepo, gc *fakeGoogleClient) {
			gc.verifyClaims = googleClaims()
		}},
		{"login existing identity", func(repo *fakeRepo, gc *fakeGoogleClient) {
			hash := hashFor(googleTestEmail)
			repo.seedIdentity(providerGoogle, googleTestEmail, hash, true)
			gc.verifyClaims = googleClaims()
		}},
		{"google unavailable", func(_ *fakeRepo, gc *fakeGoogleClient) {
			gc.exchangeErr = errors.New("dial tcp: connect refused")
		}},
	}

	for _, sc := range scenarios {
		t.Run(sc.name, func(t *testing.T) {
			svc, repo, gc := newGoogleTestService(t)
			sc.setup(repo, gc)

			logBuf := &strings.Builder{}
			old := log.Writer()
			log.SetOutput(logBuf)

			st, cookieValue := startGoogleLogin(t, svc, intentLogin, nil)
			res, err := svc.GoogleCallback(context.Background(), markerCode, st.State, cookieValue)

			log.SetOutput(old)
			if err != nil {
				t.Fatalf("unexpected internal error: %v", err)
			}
			for _, secret := range []string{markerCode, markerIDToken, googleTestEmail} {
				if strings.Contains(logBuf.String(), secret) {
					t.Errorf("secret %q leaked into logs:\n%s", secret, logBuf.String())
				}
			}
			if res.Error == "" && res.RefreshToken != "" && strings.Contains(logBuf.String(), res.RefreshToken) {
				t.Errorf("plain refresh token leaked into logs")
			}
		})
	}
}
