package http

import "net/http"

// Cookie names for the Google OAuth flow and session delivery. Single
// definition point so tests and handlers cannot drift apart.
const (
	oauthStateCookieName   = "kencleng_oauth_state"
	accessTokenCookieName  = "kencleng_access"
	refreshTokenCookieName = "kencleng_refresh"
)

// writeOAuthStateCookie sets the OAuth state cookie carrying the
// service-encoded {state, nonce, intent, user_id} payload.
//
// Attributes per R24 (techplan §3/§13): HttpOnly (no JS access), Secure in
// every non-dev environment (dev serves plain HTTP), SameSite=Lax — NOT
// Strict: the redirect back from Google's domain is cross-site navigation,
// and a Strict cookie would be dropped, breaking every flow. MaxAge matches
// the service's state TTL (~10 min); Path is scoped to the two Google
// endpoints so the browser sends the value nowhere else.
func writeOAuthStateCookie(w http.ResponseWriter, cookieSecure bool, value string) {
	http.SetCookie(w, &http.Cookie{
		Name:     oauthStateCookieName,
		Value:    value,
		Path:     "/auth/google",
		MaxAge:   int(stateCookieMaxAge.Seconds()),
		HttpOnly: true,
		Secure:   cookieSecure,
		SameSite: http.SameSiteLaxMode,
	})
}

// readOAuthStateCookie reads the raw state cookie value from the request.
// An absent cookie yields "" — the service maps that to state_mismatch
// (R20); no error plumbing needed at this layer.
func readOAuthStateCookie(r *http.Request) string {
	c, err := r.Cookie(oauthStateCookieName)
	if err != nil {
		return ""
	}
	return c.Value
}

// clearOAuthStateCookie invalidates the state cookie after the callback has
// consumed it (techplan §13 final row: without this, a replayed callback
// could reuse the still-valid state). MaxAge < 0 instructs the browser to
// delete it immediately.
func clearOAuthStateCookie(w http.ResponseWriter, cookieSecure bool) {
	http.SetCookie(w, &http.Cookie{
		Name:     oauthStateCookieName,
		Value:    "",
		Path:     "/auth/google",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   cookieSecure,
		SameSite: http.SameSiteLaxMode,
	})
}

// readRefreshCookie returns the raw refresh-cookie value; absent/empty
// cookie yields "" (the service treats that as unauthenticated).
func readRefreshCookie(r *http.Request) string {
	c, err := r.Cookie(refreshTokenCookieName)
	if err != nil {
		return ""
	}
	return c.Value
}

// writeRefreshCookie delivers ONLY the rotated refresh token — the contract
// for /auth/login, /auth/login/mfa, and /auth/refresh (openapi: the access
// token travels in the JSON body for these endpoints, never as a cookie;
// only the OAuth 302 callback needs writeAuthCookies' both-cookie shape).
//
// Attributes per index.yaml conventions + LockedOutGenericCredentials-era
// spec: HttpOnly (no JS access), Secure in every non-dev environment,
// SameSite=Strict (never crosses sites — only same-site refresh requests
// use it), Path="/" so /auth/refresh and /auth/logout see it.
func writeRefreshCookie(w http.ResponseWriter, cookieSecure bool, value string) {
	// #nosec G124 -- Secure is deliberately environment-conditional (dev
	// serves plain HTTP); HttpOnly + SameSite=Strict are always set.
	http.SetCookie(w, &http.Cookie{
		Name:     refreshTokenCookieName,
		Value:    value,
		Path:     "/",
		MaxAge:   int(refreshTokenCookieTTL.Seconds()),
		HttpOnly: true,
		Secure:   cookieSecure,
		SameSite: http.SameSiteStrictMode,
	})
}

// clearRefreshCookie instructs the browser to delete the refresh cookie
// immediately (MaxAge < 0). Called on every logout regardless of whether a
// valid cookie was presented — idempotent by contract.
func clearRefreshCookie(w http.ResponseWriter, cookieSecure bool) {
	// #nosec G124 -- same environment-conditional Secure as writeRefreshCookie.
	http.SetCookie(w, &http.Cookie{
		Name:     refreshTokenCookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   cookieSecure,
		SameSite: http.SameSiteStrictMode,
	})
}

// writeAuthCookies delivers issued session tokens as cookies for the Google
// OAuth 302-redirect contract (techplan §14 Resolved item 7 — cookie
// delivery is the only option compatible with a redirect, where there is no
// JSON body to carry the access token):
//
//   - access token: short TTL, HttpOnly. Readable path-wide so the SPA's
//     navigations carry it; it expires quickly on its own.
//   - refresh token: HttpOnly + Secure + SameSite=Strict (never crosses
//     sites — only same-site refresh requests will use it), 30-day lifetime.
//
// The access cookie carries the ES256 JWT; neither cookie ever contains the
// raw refresh token in any log output (R16). NOT used by /auth/login,
// /auth/login/mfa, or /auth/refresh — those deliver the access token in the
// JSON body and set only the refresh cookie (writeRefreshCookie).
func writeAuthCookies(w http.ResponseWriter, cookieSecure bool, accessToken, refreshToken string) {
	http.SetCookie(w, &http.Cookie{
		Name:     accessTokenCookieName,
		Value:    accessToken,
		Path:     "/",
		MaxAge:   int(accessTokenCookieTTL.Seconds()),
		HttpOnly: true,
		Secure:   cookieSecure,
		SameSite: http.SameSiteLaxMode,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     refreshTokenCookieName,
		Value:    refreshToken,
		Path:     "/",
		MaxAge:   int(refreshTokenCookieTTL.Seconds()),
		HttpOnly: true,
		Secure:   cookieSecure,
		SameSite: http.SameSiteStrictMode,
	})
}
