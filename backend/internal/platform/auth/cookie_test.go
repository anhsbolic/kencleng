package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// These assertions cover emitted headers in local unit fixtures, not browser
// enforcement, a real provider callback, or database-backed logout.
func TestBrowserConfigCookieBoundary(t *testing.T) {
	for _, tc := range []struct {
		name   string
		config BrowserConfig
		valid  bool
	}{
		{"production HTTPS", BrowserConfig{Origin: "https://app.example"}, true},
		{"production HTTP", BrowserConfig{Origin: "http://app.example"}, false},
		{"production localhost HTTP", BrowserConfig{Origin: "http://localhost:8080"}, false},
		{"exact development", BrowserConfig{Origin: "http://localhost:8080", Development: true}, true},
		{"development different host", BrowserConfig{Origin: "http://127.0.0.1:8080", Development: true}, false},
		{"development different port", BrowserConfig{Origin: "http://localhost:8090", Development: true}, false},
		{"development HTTPS", BrowserConfig{Origin: "https://localhost:8080", Development: true}, false},
		{"origin path", BrowserConfig{Origin: "https://app.example/"}, false},
		{"origin query", BrowserConfig{Origin: "https://app.example?query=value"}, false},
		{"origin fragment", BrowserConfig{Origin: "https://app.example#fragment"}, false},
		{"origin userinfo", BrowserConfig{Origin: "https://user@app.example"}, false},
		{"empty origin", BrowserConfig{}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.config.Validate(); (err == nil) != tc.valid {
				t.Fatal("origin validation disagrees with the cookie transport boundary")
			}
			if !tc.valid {
				return
			}
			wantSessionName := "__Host-kencleng_session"
			if tc.config.Development {
				wantSessionName = "kencleng_session"
			}
			if tc.config.sessionCookieName() != wantSessionName {
				t.Fatal("session name does not match the environment")
			}
			for _, spec := range []struct {
				name string
				path string
				ttl  time.Duration
			}{
				{loginCookieName, callbackPath, loginLifetime},
				{wantSessionName, "/", sessionLifetime},
			} {
				recorder := httptest.NewRecorder()
				http.SetCookie(recorder, tc.config.cookie(spec.name, "unit-fixture", spec.path, spec.ttl))
				response := recorder.Result()
				cookies := response.Cookies()
				response.Body.Close()
				if len(cookies) != 1 {
					t.Fatal("expected one emitted cookie")
				}
				c := cookies[0]
				if c.Name != spec.name || c.Path != spec.path || c.Domain != "" ||
					c.Secure != !tc.config.Development || !c.HttpOnly || c.SameSite != http.SameSiteLaxMode ||
					c.MaxAge != int(spec.ttl.Seconds()) || !c.Expires.After(time.Now()) {
					t.Fatal("emitted cookie lost its transport, host, path, lifetime, or script boundary")
				}
			}
		})
	}
}

func TestClearLoginCookieAttributes(t *testing.T) {
	for _, config := range []BrowserConfig{
		{Origin: "https://app.example"},
		{Origin: "http://localhost:8080", Development: true},
	} {
		if err := config.Validate(); err != nil {
			t.Fatal("invalid unit fixture")
		}
		recorder := httptest.NewRecorder()
		h := Handler{Browser: config}
		h.clearLoginCookie(recorder)
		response := recorder.Result()
		cookies := response.Cookies()
		response.Body.Close()
		if len(cookies) != 1 {
			t.Fatal("expected one deletion cookie")
		}
		c := cookies[0]
		if c.Name != loginCookieName || c.Value != "" || c.Path != callbackPath || c.Domain != "" ||
			c.MaxAge != -1 || c.Expires.After(time.Now()) || c.Secure != !config.Development ||
			!c.HttpOnly || c.SameSite != http.SameSiteLaxMode {
			t.Fatal("login deletion cookie lost its target or security attributes")
		}
	}
}
