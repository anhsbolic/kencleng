package http

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/anhsbolic/kencleng/backend/internal/domain/account"
)

// Cookie lifetimes. The state cookie must match the service's stateTTL so a
// browser-enforced expiry and the service's validation agree (~10 min).
// Auth-cookie lifetimes mirror the token TTLs issued by IssueTokens.
const (
	stateCookieMaxAge     = 10 * time.Minute
	accessTokenCookieTTL  = 15 * time.Minute
	refreshTokenCookieTTL = 30 * 24 * time.Hour

	reauthMarkerTTL = 5 * time.Minute // matches the state TTL convention (spec Assumption A)
)

// googleOAuthService is the slice of the account service the two Google
// OAuth handlers depend on. *account.Service satisfies it; tests inject a
// stub so the transport contract (status codes, cookies, redirects) is
// exercisable without wiring the full domain — same seam philosophy as the
// domain's own googleOAuthClient/breachChecker interfaces.
type googleOAuthService interface {
	GoogleRedirect(ctx context.Context, intent string, sessionUserID *uuid.UUID) (consentURL, cookieValue string, err error)
	GoogleCallback(ctx context.Context, code, state, cookieValue string) (account.CallbackResult, error)
}

// ---- inline ES256 session verification (R25) ------------------------------

// GoogleTokenVerifier builds the token-verification function used by
// GoogleRedirectHandler to authenticate link/reauth intents. It verifies the
// application's own ES256 access tokens with golang-jwt/jwt/v5 and the
// public key passed in as a dependency.
//
// platform/auth/ is deliberately NOT modified: it is a Tier 0 fenced path
// (root AGENTS.md §3). This verification lives inline at the handler
// boundary where the conditional-auth decision is visible and testable on
// its own; the login/session task (#3) may later extract a shared helper,
// possibly as a human-paired change (techplan §9 step 5).
func GoogleTokenVerifier(publicKey *ecdsa.PublicKey) func(token string) (uuid.UUID, error) {
	return func(token string) (uuid.UUID, error) {
		if token == "" {
			return uuid.Nil, fmt.Errorf("no token presented")
		}
		rc := &jwt.RegisteredClaims{}
		_, err := jwt.ParseWithClaims(token, rc, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
			}
			return publicKey, nil
		},
			jwt.WithValidMethods([]string{"ES256"}),
			jwt.WithExpirationRequired(),
			jwt.WithLeeway(time.Minute), // same clock-skew tolerance as id_token checks
		)
		if err != nil {
			return uuid.Nil, fmt.Errorf("invalid session token")
		}
		userID, err := uuid.Parse(rc.Subject)
		if err != nil {
			return uuid.Nil, fmt.Errorf("session token subject is not a user id")
		}
		return userID, nil
	}
}

// sessionToken extracts the presented session token: the HttpOnly
// access-token cookie first (the natural carrier for browser-initiated
// navigations, since our own setAuthCookies delivers it there), falling back
// to an Authorization: Bearer header for non-browser clients and tests.
func sessionToken(r *http.Request) string {
	if c, err := r.Cookie(accessTokenCookieName); err == nil && c.Value != "" {
		return c.Value
	}
	const prefix = "Bearer "
	if auth := r.Header.Get("Authorization"); len(auth) > len(prefix) && auth[:len(prefix)] == prefix {
		return auth[len(prefix):]
	}
	return ""
}

// ---- reauth marker store (R12; consumed by task #06) ----------------------

// reauthMarkers holds short-lived "recently re-authenticated" markers keyed
// by user id. In-memory sync.Map per techplan §5 Decision (Redis reverted —
// not in docker-compose, no client in go.mod): markers are lost on restart,
// accepted because the TTL is only 5 minutes. Expired entries are deleted
// lazily on read and swept periodically by a background goroutine — the
// same pattern as the rate limiter in middleware.go.
var reauthMarkers sync.Map // key: uuid.UUID → value: time.Time (expiry)

// SetReauthMarker records that userID completed re-authentication; the
// marker is valid until expiry (5 minutes).
func SetReauthMarker(userID uuid.UUID, expiry time.Time) {
	reauthMarkers.Store(userID, expiry)
}

// CheckReauthMarker reports whether userID has a currently-valid reauth
// marker, deleting it if expired (lazy eviction).
func CheckReauthMarker(userID uuid.UUID) bool {
	v, ok := reauthMarkers.Load(userID)
	if !ok {
		return false
	}
	expiry := v.(time.Time)
	if time.Now().After(expiry) {
		reauthMarkers.Delete(userID)
		return false
	}
	return true
}

// ConsumeReauthMarker atomically checks AND invalidates the reauth marker
// for userID in one step (LoadAndDelete): it returns whether a currently-
// valid marker was present, and in all cases removes it so it can never be
// replayed for a second sensitive action. This fulfills feature-06's
// consume-on-use clause for the Google-only MFA-disable path: a second
// disable call finds the marker gone and is rejected. Existing
// CheckReauthMarker (read-only + lazy-expiry sweep) and the background
// sweeper are left untouched.
func ConsumeReauthMarker(userID uuid.UUID) bool {
	v, ok := reauthMarkers.LoadAndDelete(userID)
	if !ok {
		return false
	}
	expiry := v.(time.Time)
	if time.Now().After(expiry) {
		return false // expired — consumed and invalid
	}
	return true
}

func init() {
	// Periodic sweep bounding memory from abandoned markers (same shape as
	// middleware.go's rate-limiter sweeper). A package-level goroutine is
	// acceptable here: the process is a long-lived server, matching how the
	// rate limiter already behaves.
	go func() {
		t := time.NewTicker(time.Minute)
		defer t.Stop()
		for range t.C {
			now := time.Now()
			reauthMarkers.Range(func(key, value any) bool {
				if now.After(value.(time.Time)) {
					reauthMarkers.Delete(key)
				}
				return true
			})
		}
	}()
}

// ---- handlers --------------------------------------------------------------

// GoogleRedirectHandler handles GET /auth/google/redirect?intent={login|link|reauth}.
//
// R18: unknown intent → 400 Problem Details.
// R1: intent=login requires no authentication.
// R2: intent=link/reauth without a verifiable session → 401 BEFORE any
// Google redirect — the explicit authz check lives right here, not hidden
// behind a query filter (AGENTS.md golden rule).
// R3: with a session, the verified user_id rides inside the state cookie.
//
// On success the response sets the state cookie (R24 attributes) and 302s
// to the Google consent screen.
func GoogleRedirectHandler(svc googleOAuthService, verifyToken func(string) (uuid.UUID, error), cookieSecure bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		intent := r.URL.Query().Get("intent")

		var sessionUserID *uuid.UUID
		if intent == "link" || intent == "reauth" {
			userID, err := verifyToken(sessionToken(r))
			if err != nil {
				// R2: reject before generating anything or contacting Google.
				log.Printf("transport: google redirect rejected (intent=%s, reason=unauthenticated)", intent)
				WriteProblem(w, http.StatusUnauthorized,
					"https://kencleng.dev/problems/unauthenticated",
					"Authentication Required", "Sign in before continuing.")
				return
			}
			sessionUserID = &userID
		} else if intent != "login" {
			WriteProblem(w, http.StatusBadRequest,
				"https://kencleng.dev/problems/invalid-intent",
				"Invalid Intent", "The intent parameter must be one of: login, link, reauth.")
			return
		}

		consentURL, cookieValue, err := svc.GoogleRedirect(r.Context(), intent, sessionUserID)
		if err != nil {
			switch {
			case err == account.ErrInvalidIntent:
				WriteProblem(w, http.StatusBadRequest,
					"https://kencleng.dev/problems/invalid-intent",
					"Invalid Intent", "The intent parameter must be one of: login, link, reauth.")
			case err == account.ErrMissingSession:
				WriteProblem(w, http.StatusUnauthorized,
					"https://kencleng.dev/problems/unauthenticated",
					"Authentication Required", "Sign in before continuing.")
			default:
				log.Printf("transport: google redirect failed: %v", err)
				WriteProblem(w, http.StatusInternalServerError,
					"https://kencleng.dev/problems/internal",
					"Internal Error", "An unexpected error occurred.")
			}
			return
		}

		writeOAuthStateCookie(w, cookieSecure, cookieValue)
		http.Redirect(w, r, consentURL, http.StatusFound)
	}
}

// GoogleCallbackHandler handles GET /auth/google/callback?code=...&state=....
//
// All validation logic (state comparison, code exchange, id_token
// verification, intent branching) lives in the service; this handler
// translates its CallbackResult into HTTP: auth cookies on success (or the
// reauth marker), a cleared state cookie always, and a 302 to the frontend
// carrying ?error={code} on failure. No internal detail ever reaches the
// redirect target — only the stable error-code vocabulary.
func GoogleCallbackHandler(svc googleOAuthService, cookieSecure bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		state := r.URL.Query().Get("state")
		cookieValue := readOAuthStateCookie(r)

		res, err := svc.GoogleCallback(r.Context(), code, state, cookieValue)

		// Consume the state cookie regardless of outcome — a replayed
		// callback must never reuse it (techplan §13 final row).
		clearOAuthStateCookie(w, cookieSecure)

		if err != nil {
			log.Printf("transport: google callback failed: %v", err)
			WriteProblem(w, http.StatusInternalServerError,
				"https://kencleng.dev/problems/internal",
				"Internal Error", "An unexpected error occurred.")
			return
		}

		if res.Error != "" || res.RedirectURL == "" {
			// Failure leg: RedirectURL already carries ?error={code}.
			target := res.RedirectURL
			if target == "" {
				target = "/login" // defensive: never 302 to an empty Location
			}
			http.Redirect(w, r, target, http.StatusFound)
			return
		}

		if res.Reauth {
			SetReauthMarker(res.UserID, time.Now().Add(reauthMarkerTTL))
			log.Printf("transport: reauth marker set user_id=%s", res.UserID)
		} else {
			writeAuthCookies(w, cookieSecure, res.AccessToken, res.RefreshToken)
		}
		http.Redirect(w, r, res.RedirectURL, http.StatusFound)
	}
}
