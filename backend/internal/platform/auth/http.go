package auth

import (
	"crypto/subtle"
	"encoding/json"
	"errors"
	"io"
	"mime"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const loginCookieName = "kencleng_login"
const callbackPath = "/api/auth/google/callback"
const formPath = "/organization-establishments/new"

type BrowserConfig struct {
	Origin      string
	Development bool
}

func (c BrowserConfig) Validate() error {
	u, err := url.Parse(c.Origin)
	if err != nil || u.Host == "" || u.Path != "" || u.RawQuery != "" || u.Fragment != "" || u.User != nil {
		return errors.New("APP_ORIGIN must be an origin")
	}
	if c.Development {
		if c.Origin != "http://localhost:8080" {
			return errors.New("development auth origin must be http://localhost:8080")
		}
	} else if u.Scheme != "https" {
		return errors.New("production auth origin must use HTTPS")
	}
	return nil
}

func (c BrowserConfig) sessionCookieName() string {
	if c.Development {
		return "kencleng_session"
	}
	return "__Host-kencleng_session"
}

func (c BrowserConfig) cookie(name, value, path string, ttl time.Duration) *http.Cookie {
	return &http.Cookie{Name: name, Value: value, Path: path, MaxAge: int(ttl.Seconds()),
		Expires: time.Now().Add(ttl), Secure: !c.Development, HttpOnly: true, SameSite: http.SameSiteLaxMode}
}

type Handler struct {
	Store    Store
	Protocol Protocol
	Browser  BrowserConfig
	Now      func() time.Time
	limitMu  sync.Mutex
	limit    map[string]loginWindow
}

type loginWindow struct {
	At    time.Time
	Count int
}

func NewHandler(store Store, protocol Protocol, browser BrowserConfig) (*Handler, error) {
	if err := browser.Validate(); err != nil {
		return nil, err
	}
	if store.Pool == nil || protocol == nil {
		return nil, errors.New("auth dependencies are missing")
	}
	return &Handler{Store: store, Protocol: protocol, Browser: browser, Now: time.Now, limit: make(map[string]loginWindow)}, nil
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /auth/google/start", h.start)
	mux.HandleFunc("GET /auth/google/callback", h.callback)
	mux.HandleFunc("GET /me", h.me)
	mux.HandleFunc("POST /auth/logout", h.logout)
}

func noStore(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "private, no-store")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Referrer-Policy", "no-referrer")
	w.Header().Set("X-Content-Type-Options", "nosniff")
}

func respondError(w http.ResponseWriter, status int, code, message string) {
	noStore(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	requestID, err := newUUID()
	if err != nil {
		requestID = "unavailable"
	}
	_ = json.NewEncoder(w).Encode(map[string]any{"error": map[string]string{
		"code": code, "message": message, "request_id": requestID}})
}

func (h *Handler) allowLogin(remote string, now time.Time) bool {
	host, _, err := net.SplitHostPort(remote)
	if err != nil {
		host = remote
	}
	h.limitMu.Lock()
	defer h.limitMu.Unlock()
	if len(h.limit) > 10000 {
		for key, window := range h.limit {
			if now.Sub(window.At) >= time.Minute {
				delete(h.limit, key)
			}
		}
	}
	w := h.limit[host]
	if now.Sub(w.At) >= time.Minute {
		w = loginWindow{At: now}
	}
	w.Count++
	h.limit[host] = w
	return w.Count <= 20
}

func (h *Handler) start(w http.ResponseWriter, r *http.Request) {
	noStore(w)
	now := h.Now()
	if !h.allowLogin(r.RemoteAddr, now) {
		respondError(w, http.StatusTooManyRequests, "too_many_requests", "Please try again later.")
		return
	}
	nonce, err := randomToken()
	if err != nil {
		respondError(w, 503, "processing_unavailable", "Sign-in is unavailable.")
		return
	}
	verifier, err := randomToken()
	if err != nil {
		respondError(w, 503, "processing_unavailable", "Sign-in is unavailable.")
		return
	}
	state, browser, err := h.Store.StartLogin(r.Context(), now, nonce, verifier)
	if err != nil {
		respondError(w, 503, "processing_unavailable", "Sign-in is unavailable.")
		return
	}
	h.clearLoginCookie(w)
	h.setLoginCookie(w, browser)
	http.Redirect(w, r, h.Protocol.AuthorizeURL(state, nonce, verifier), http.StatusFound)
}

func (h *Handler) setLoginCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, h.Browser.cookie(loginCookieName, token, callbackPath, loginLifetime))
}

func (h *Handler) clearLoginCookie(w http.ResponseWriter) {
	c := h.Browser.cookie(loginCookieName, "", callbackPath, 0)
	c.MaxAge = -1
	http.SetCookie(w, c)
}

func (h *Handler) sessionToken(r *http.Request) string {
	cookie, err := r.Cookie(h.Browser.sessionCookieName())
	if err != nil {
		return ""
	}
	return cookie.Value
}

func (h *Handler) callback(w http.ResponseWriter, r *http.Request) {
	noStore(w)
	h.clearLoginCookie(w)
	// Callback output is fixed regardless of provider error or attacker input.
	defer http.Redirect(w, r, formPath, http.StatusFound)
	query := r.URL.Query()
	stateValues, codeValues, errorValues := query["state"], query["code"], query["error"]
	if len(stateValues) != 1 || len(codeValues) > 1 || len(errorValues) > 1 {
		return
	}
	browser, err := r.Cookie(loginCookieName)
	if err != nil {
		return
	}
	pending, err := h.Store.ConsumeLogin(r.Context(), stateValues[0], browser.Value, h.Now())
	if err != nil {
		return
	}
	if len(errorValues) != 0 || len(codeValues) != 1 || codeValues[0] == "" {
		return
	}
	identity, err := h.Protocol.ExchangeAndVerify(r.Context(), codeValues[0], pending.Nonce, pending.Verifier)
	if err != nil || identity.Issuer != googleIssuer || identity.Subject == "" {
		return
	}
	personID, err := h.Store.UpsertPerson(r.Context(), identity.Issuer, identity.Subject)
	if err != nil {
		return
	}
	token, err := h.Store.RotateSession(r.Context(), personID, h.sessionToken(r), h.Now())
	if err != nil {
		return
	}
	http.SetCookie(w, h.Browser.cookie(h.Browser.sessionCookieName(), token, "/", sessionLifetime))
}

func (h *Handler) me(w http.ResponseWriter, r *http.Request) {
	noStore(w)
	session, err := h.Store.ReadSession(r.Context(), h.sessionToken(r), h.Now())
	if errors.Is(err, ErrUnauthenticated) {
		respondError(w, 401, "authentication_required", "Sign-in is required.")
		return
	}
	if err != nil {
		respondError(w, 503, "processing_unavailable", "Please try again later.")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"person_id": session.PersonID, "csrf_token": session.CSRFToken})
}

// CheckMutation applies the same browser/CSRF boundary to T3 handlers.
// It returns the current unlocked session; confirmation must still call
// LockSession in its transaction before grant decisions.
func (h *Handler) CheckMutation(w http.ResponseWriter, r *http.Request) (Session, bool) {
	if r.Header.Get("Origin") != h.Browser.Origin || strings.Contains(r.Header.Get("Origin"), ",") {
		respondError(w, 403, "request_not_allowed", "Request is not allowed.")
		return Session{}, false
	}
	mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil || mediaType != "application/json" {
		respondError(w, 403, "request_not_allowed", "Request is not allowed.")
		return Session{}, false
	}
	session, err := h.Store.ReadSession(r.Context(), h.sessionToken(r), h.Now())
	if errors.Is(err, ErrUnauthenticated) {
		respondError(w, 401, "authentication_required", "Sign-in is required.")
		return Session{}, false
	}
	if err != nil {
		respondError(w, 503, "processing_unavailable", "Please try again later.")
		return Session{}, false
	}
	provided := r.Header.Get("X-CSRF-Token")
	if provided == "" || subtle.ConstantTimeCompare([]byte(provided), []byte(session.CSRFToken)) != 1 {
		respondError(w, 403, "request_not_allowed", "Request is not allowed.")
		return Session{}, false
	}
	return session, true
}

func (h *Handler) logout(w http.ResponseWriter, r *http.Request) {
	noStore(w)
	if _, ok := h.CheckMutation(w, r); !ok {
		return
	}
	defer r.Body.Close()
	decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, 16))
	var payload map[string]json.RawMessage
	if err := decoder.Decode(&payload); err != nil || len(payload) != 0 {
		respondError(w, 400, "invalid_request", "Invalid request.")
		return
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		respondError(w, 400, "invalid_request", "Invalid request.")
		return
	}
	if err := h.Store.RevokeSession(r.Context(), h.sessionToken(r), h.Now()); err != nil {
		if errors.Is(err, ErrUnauthenticated) {
			respondError(w, 401, "authentication_required", "Sign-in is required.")
		} else {
			respondError(w, 503, "processing_unavailable", "Please try again later.")
		}
		return
	}
	c := h.Browser.cookie(h.Browser.sessionCookieName(), "", "/", 0)
	c.MaxAge = -1
	http.SetCookie(w, c)
	w.WriteHeader(http.StatusNoContent)
}
