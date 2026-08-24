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

// writeAuthCookies delivers issued session tokens as cookies (techplan §14
// Resolved item 7 — cookie delivery is the only option compatible with the
// 302-redirect contract):
//
//   - access token: short TTL, HttpOnly. Readable path-wide so the SPA's
//     navigations carry it; it expires quickly on its own.
//   - refresh token: HttpOnly + Secure + SameSite=Strict (never crosses
//     sites — only same-site refresh requests will use it), 30-day lifetime.
//
// The access cookie carries the ES256 JWT; neither cookie ever contains the
// raw refresh token in any log output (R16).
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
