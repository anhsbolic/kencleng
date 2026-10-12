//go:build authharness

package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Provisioner is only a fixture seed/metadata/lock witness. Every behavior query
// and competing session operation uses the separate restricted runtime pool.
type harnessFixture struct {
	ctx                  context.Context
	runtime, provisioner *pgxpool.Pool
	store                Store
}

func harness(t *testing.T) harnessFixture {
	t.Helper()
	if os.Getenv("C1_HARNESS_OWNED_RUN") != "RUN-C1-ENG-BUILD-T2-LOGOUT-PATCH-001" {
		t.Fatal("use scripts/auth_harness.py: owned fixture binding required")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	t.Cleanup(cancel)
	connect := func(key string) *pgxpool.Pool {
		config, err := pgxpool.ParseConfig(os.Getenv(key))
		if err != nil || config == nil || config.ConnConfig.Host != "127.0.0.1" || config.ConnConfig.Database != "c1_fixture" || config.ConnConfig.RuntimeParams["search_path"] != "c1_identity" {
			t.Fatal("invalid owned fixture connection")
		}
		config.MaxConns = 16
		pool, err := pgxpool.NewWithConfig(ctx, config)
		if err != nil {
			t.Fatal("fixture connection failed")
		}
		t.Cleanup(pool.Close)
		if err := pool.Ping(ctx); err != nil {
			t.Fatal("fixture ping failed")
		}
		return pool
	}
	runtime, provisioner := connect("C1_HARNESS_RUNTIME_DSN"), connect("C1_HARNESS_PROVISIONER_DSN")
	return harnessFixture{ctx, runtime, provisioner, Store{runtime}}
}

func (f harnessFixture) session(t *testing.T) (string, Session) {
	t.Helper()
	subject, err := newUUID()
	if err != nil {
		t.Fatal("fixture subject generation")
	}
	person, err := f.store.UpsertPerson(f.ctx, googleIssuer, subject)
	if err != nil {
		t.Fatal("runtime person query failed")
	}
	token, err := f.store.RotateSession(f.ctx, person, "", time.Now())
	if err != nil {
		t.Fatal("runtime rotation failed")
	}
	session, err := f.store.ReadSession(f.ctx, token, time.Now())
	if err != nil {
		t.Fatal("runtime session read failed")
	}
	return token, session
}

func TestHarnessRuntimePrivileges(t *testing.T) {
	f := harness(t)
	var identity, version string
	var restricted, tcp, scram, hba bool
	if err := f.runtime.QueryRow(f.ctx, `SELECT current_user,version(),inet_client_addr() IS NOT NULL,
	 NOT (rolsuper OR rolcreatedb OR rolcreaterole OR rolreplication OR rolbypassrls)
	 FROM pg_roles WHERE rolname=current_user`).Scan(&identity, &version, &tcp, &restricted); err != nil {
		t.Fatal("identity metadata query failed")
	}
	if identity != "c1_runtime" || !restricted || !tcp {
		t.Fatal("runtime identity or TCP restriction failed")
	}
	if err := f.provisioner.QueryRow(f.ctx, `SELECT rolpassword LIKE 'SCRAM-SHA-256$%' FROM pg_authid WHERE rolname='c1_runtime'`).Scan(&scram); err != nil || !scram {
		t.Fatal("fixture SCRAM storage metadata failed")
	}
	if err := f.provisioner.QueryRow(f.ctx, `SELECT count(*)>0 AND bool_and(auth_method='scram-sha-256') FROM pg_hba_file_rules WHERE type LIKE 'host%'`).Scan(&hba); err != nil || !hba {
		t.Fatal("TCP HBA SCRAM oracle failed")
	}
	config := f.runtime.Config().ConnConfig.Copy()
	config.Password = "deliberately-wrong-test-password"
	bad, err := pgx.ConnectConfig(f.ctx, config)
	if bad != nil {
		bad.Close(f.ctx)
	}
	var pgerr *pgconn.PgError
	if !errors.As(err, &pgerr) || pgerr.Code != "28P01" {
		t.Fatal("wrong-password TCP connection did not fail authentication")
	}
	for _, sql := range []string{"CREATE TABLE c1_identity.forbidden(id int)", "CREATE SCHEMA forbidden", "CREATE ROLE forbidden", "SET ROLE postgres", "ALTER TABLE persons ADD COLUMN forbidden int", "DELETE FROM persons", "UPDATE pending_logins SET nonce='forbidden'", "DELETE FROM sessions"} {
		_, err := f.runtime.Exec(f.ctx, sql)
		if !errors.As(err, &pgerr) || pgerr.Code != "42501" {
			t.Fatal("runtime forbidden operation did not return insufficient_privilege")
		}
	}
	// PostgreSQL can report a no-op GRANT as a warning; the oracle is unchanged
	// privileges/grant options, not an assumed error for that SQL statement.
	if _, err := f.runtime.Exec(f.ctx, "GRANT ALL ON persons TO c1_runtime"); err != nil && (!errors.As(err, &pgerr) || pgerr.Code != "42501") {
		t.Fatal("self-grant returned unexpected error")
	}
	var memberships int
	if err := f.provisioner.QueryRow(f.ctx, `SELECT count(*) FROM pg_auth_members WHERE member=(SELECT oid FROM pg_roles WHERE rolname='c1_runtime')`).Scan(&memberships); err != nil || memberships != 0 {
		t.Fatal("runtime membership oracle failed")
	}
	rows, err := f.provisioner.Query(f.ctx, `SELECT table_name,privilege_type,is_grantable FROM information_schema.role_table_grants WHERE grantee='c1_runtime' AND table_schema='c1_identity' ORDER BY 1,2`)
	if err != nil {
		t.Fatal("grant metadata failed")
	}
	defer rows.Close()
	var grants []string
	for rows.Next() {
		var table, priv, grantable string
		if rows.Scan(&table, &priv, &grantable) != nil || grantable != "NO" {
			t.Fatal("grant authority mismatch")
		}
		grants = append(grants, table+":"+priv)
	}
	if rows.Err() != nil || strings.Join(grants, ",") != "pending_logins:DELETE,pending_logins:INSERT,pending_logins:SELECT,persons:INSERT,persons:SELECT,persons:UPDATE,sessions:INSERT,sessions:SELECT,sessions:UPDATE" {
		t.Fatal("exact runtime grants mismatch")
	}
	t.Log("runtime=c1_runtime; TCP SCRAM success + 28P01 negative; no memberships; exact nine non-grantable table privileges; DDL/role/owner/delete denials verified")
	t.Log("server", version)
}

type harnessProtocol struct {
	state, nonce, verifier string
	calls                  int
	fail                   bool
}

func (p *harnessProtocol) AuthorizeURL(state, nonce, verifier string) string {
	p.state, p.nonce, p.verifier = state, nonce, verifier
	return "https://provider.invalid/authorize?state=" + url.QueryEscape(state)
}
func (p *harnessProtocol) ExchangeAndVerify(_ context.Context, code, nonce, verifier string) (Identity, error) {
	p.calls++
	if p.fail || code != "simulated-code" || nonce != p.nonce || verifier != p.verifier {
		return Identity{}, ErrUntrustedIdentity
	}
	return Identity{Issuer: googleIssuer, Subject: "simulated-harness-subject"}, nil
}

func harnessHandler(t *testing.T, f harnessFixture) (*Handler, *http.ServeMux, *harnessProtocol) {
	t.Helper()
	p := &harnessProtocol{}
	h, err := NewHandler(f.store, p, BrowserConfig{Origin: "http://localhost:8080", Development: true})
	if err != nil {
		t.Fatal("construct fixture handler failed")
	}
	mux := http.NewServeMux()
	h.Register(mux)
	return h, mux, p
}

func harnessRequest(mux http.Handler, method, path, body, token, csrf, origin, media string) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, path, strings.NewReader(body))
	if token != "" {
		r.AddCookie(&http.Cookie{Name: "kencleng_session", Value: token})
	}
	r.Header.Set("Origin", origin)
	r.Header.Set("Content-Type", media)
	r.Header.Set("X-CSRF-Token", csrf)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	return w
}

func safeResponse(t *testing.T, w *httptest.ResponseRecorder, status int, code string) {
	t.Helper()
	if w.Code != status {
		t.Fatalf("HTTP status: got %d want %d", w.Code, status)
	}
	if w.Header().Get("Cache-Control") != "private, no-store" || w.Header().Get("Access-Control-Allow-Credentials") != "" || w.Header().Get("Access-Control-Allow-Origin") != "" {
		t.Fatal("cache/CORS oracle failed")
	}
	if code != "" {
		var body struct {
			Error map[string]string `json:"error"`
		}
		if json.Unmarshal(w.Body.Bytes(), &body) != nil || len(body.Error) != 3 || body.Error["code"] != code || body.Error["request_id"] == "" {
			t.Fatal("safe error shape failed")
		}
		expected := map[string]string{"authentication_required": "Sign-in is required.", "request_not_allowed": "Request is not allowed.", "invalid_request": "Invalid request.", "processing_unavailable": "Please try again later."}
		if body.Error["message"] != expected[code] {
			t.Fatal("safe public message mismatch")
		}
	}
}

func TestHarnessHTTPMeMutationLogout(t *testing.T) {
	f := harness(t)
	h, mux, _ := harnessHandler(t, f)
	token, session := f.session(t)
	for _, item := range []struct{ name, token string }{{"anonymous", ""}, {"malformed", "not-a-token"}} {
		t.Run(item.name, func(t *testing.T) {
			safeResponse(t, harnessRequest(mux, "GET", "/me", "", item.token, "", "", ""), 401, "authentication_required")
		})
	}
	for _, state := range []string{"expired", "revoked"} {
		other, _ := f.session(t)
		var err error
		if state == "expired" {
			_, err = f.provisioner.Exec(f.ctx, `UPDATE sessions SET created_at=now()-interval '9 hours',expires_at=now()-interval '1 hour' WHERE token_digest=$1`, digest(other))
		} else {
			err = f.store.RevokeSession(f.ctx, other, time.Now())
		}
		if err != nil {
			t.Fatal("fixture denied-session setup failed")
		}
		t.Run(state, func(t *testing.T) {
			safeResponse(t, harnessRequest(mux, "GET", "/me", "", other, "", "", ""), 401, "authentication_required")
		})
	}
	w := harnessRequest(mux, "GET", "/me", "", token, "", "", "")
	safeResponse(t, w, 200, "")
	var me map[string]string
	if json.Unmarshal(w.Body.Bytes(), &me) != nil || len(me) != 2 || me["person_id"] != session.PersonID || me["csrf_token"] != session.CSRFToken {
		t.Fatal("local-only me fields failed")
	}
	other, otherSession := f.session(t)
	for _, tc := range []struct {
		name, origin, media, token, csrf string
		status                           int
		code                             string
	}{
		{"missing-origin", "", "application/json", token, session.CSRFToken, 403, "request_not_allowed"},
		{"wrong-origin", "https://attacker.invalid", "application/json", token, session.CSRFToken, 403, "request_not_allowed"},
		{"comma-origin", h.Browser.Origin + ", https://attacker.invalid", "application/json", token, session.CSRFToken, 403, "request_not_allowed"},
		{"prefix-origin", h.Browser.Origin + ".attacker.invalid", "application/json", token, session.CSRFToken, 403, "request_not_allowed"},
		{"wrong-type", h.Browser.Origin, "text/plain", token, session.CSRFToken, 403, "request_not_allowed"},
		{"malformed-type", h.Browser.Origin, "application/json; broken", token, session.CSRFToken, 403, "request_not_allowed"},
		{"anonymous", h.Browser.Origin, "application/json", "", session.CSRFToken, 401, "authentication_required"},
		{"missing-csrf", h.Browser.Origin, "application/json", token, "", 403, "request_not_allowed"},
		{"other-session-csrf", h.Browser.Origin, "application/json", token, otherSession.CSRFToken, 403, "request_not_allowed"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := httptest.NewRequest("POST", "/mutation", nil)
			r.Header.Set("Origin", tc.origin)
			r.Header.Set("Content-Type", tc.media)
			r.Header.Set("X-CSRF-Token", tc.csrf)
			r.AddCookie(&http.Cookie{Name: "kencleng_session", Value: tc.token})
			w := httptest.NewRecorder()
			s, ok := h.CheckMutation(w, r)
			if ok || s != (Session{}) {
				t.Fatal("denied mutation returned authority")
			}
			safeResponse(t, w, tc.status, tc.code)
		})
	}
	r := httptest.NewRequest("POST", "/mutation", nil)
	r.Header.Set("Origin", h.Browser.Origin)
	r.Header.Set("Content-Type", "application/json; charset=utf-8")
	r.Header.Set("X-CSRF-Token", session.CSRFToken)
	r.AddCookie(&http.Cookie{Name: "kencleng_session", Value: token})
	if got, ok := h.CheckMutation(httptest.NewRecorder(), r); !ok || got.PersonID != session.PersonID {
		t.Fatal("valid mutation denied")
	}
	_ = other
	for _, body := range []string{"", "{", `{"extra":true}`, "{}{}", strings.Repeat(" ", 17) + "{}", "[]", `"text"`} {
		safeResponse(t, harnessRequest(mux, "POST", "/auth/logout", body, token, session.CSRFToken, h.Browser.Origin, "application/json"), 400, "invalid_request")
		if _, err := f.store.ReadSession(f.ctx, token, time.Now()); err != nil {
			t.Fatal("invalid logout revoked session")
		}
	}
	w = harnessRequest(mux, "POST", "/auth/logout", "{}", token, session.CSRFToken, h.Browser.Origin, "application/json")
	safeResponse(t, w, 204, "")
	cookies := w.Result().Cookies()
	if len(cookies) != 1 || cookies[0].Name != "kencleng_session" || cookies[0].MaxAge != -1 || cookies[0].Value != "" || cookies[0].Path != "/" {
		t.Fatal("logout deletion cookie failed")
	}
	if _, err := f.store.ReadSession(f.ctx, token, time.Now()); !errors.Is(err, ErrUnauthenticated) {
		t.Fatal("logout did not revoke")
	}
	safeResponse(t, harnessRequest(mux, "POST", "/auth/logout", "{}", token, session.CSRFToken, h.Browser.Origin, "application/json"), 401, "authentication_required")
	// A separate closed runtime pool creates a safe storage-failure boundary.
	closed, err := pgxpool.New(f.ctx, os.Getenv("C1_HARNESS_RUNTIME_DSN"))
	if err != nil {
		t.Fatal("closed pool fixture failed")
	}
	closed.Close()
	h.Store = Store{closed}
	safeResponse(t, harnessRequest(mux, "GET", "/me", "", token, "", "", ""), 503, "processing_unavailable")
}

func TestHarnessHTTPSimulatedCallback(t *testing.T) {
	f := harness(t)
	for _, scenario := range []string{"success", "browser-mismatch", "provider-error", "protocol-failure", "duplicate-state", "duplicate-code", "malformed-state", "missing-code"} {
		t.Run(scenario, func(t *testing.T) {
			_, mux, p := harnessHandler(t, f)
			start := harnessRequest(mux, "GET", "/auth/google/start?return_to=https://attacker.invalid", "", "", "", "", "")
			safeResponse(t, start, 302, "")
			if !strings.HasPrefix(start.Header().Get("Location"), "https://provider.invalid/authorize?") {
				t.Fatal("simulated authorize redirect failed")
			}
			var browser *http.Cookie
			for _, c := range start.Result().Cookies() {
				if c.Name == loginCookieName && c.Value != "" {
					browser = c
				}
			}
			if browser == nil || browser.Path != callbackPath || !browser.HttpOnly || browser.SameSite != http.SameSiteLaxMode {
				t.Fatal("pending browser cookie failed")
			}
			old, oldSession := f.session(t)
			q := url.Values{"state": {p.state}, "code": {"simulated-code"}, "return_to": {"https://attacker.invalid"}}
			switch scenario {
			case "browser-mismatch":
				browser.Value = old
			case "provider-error":
				q.Set("error", "access_denied")
				q.Del("code")
			case "protocol-failure":
				p.fail = true
			case "duplicate-state":
				q.Add("state", p.state)
			case "duplicate-code":
				q.Add("code", "second")
			case "malformed-state":
				q.Set("state", "malformed")
			case "missing-code":
				q.Del("code")
			}
			callback := func(q url.Values) *httptest.ResponseRecorder {
				r := httptest.NewRequest("GET", "/auth/google/callback?"+q.Encode(), nil)
				r.AddCookie(browser)
				r.AddCookie(&http.Cookie{Name: "kencleng_session", Value: old})
				w := httptest.NewRecorder()
				mux.ServeHTTP(w, r)
				return w
			}
			w := callback(q)
			safeResponse(t, w, 302, "")
			if w.Header().Get("Location") != "/organization-establishments/new" {
				t.Fatal("callback redirect is not fixed local path")
			}
			var issued string
			for _, c := range w.Result().Cookies() {
				if c.Name == "kencleng_session" {
					issued = c.Value
					if c.Path != "/" || !c.HttpOnly || c.SameSite != http.SameSiteLaxMode || c.MaxAge != 28800 {
						t.Fatal("actual session cookie flags failed")
					}
				}
			}
			if scenario == "success" {
				if issued == "" || issued == old || p.calls != 1 {
					t.Fatal("simulated callback did not rotate")
				}
				if _, err := f.store.ReadSession(f.ctx, old, time.Now()); !errors.Is(err, ErrUnauthenticated) {
					t.Fatal("previous session retained authority")
				}
				current, err := f.store.ReadSession(f.ctx, issued, time.Now())
				if err != nil || current.PersonID == oldSession.PersonID {
					t.Fatal("simulated local mapping failed")
				}
			} else {
				if issued != "" {
					t.Fatal("negative callback minted a session")
				}
				if _, err := f.store.ReadSession(f.ctx, old, time.Now()); err != nil {
					t.Fatal("failed callback altered previous authority")
				}
			}
			if scenario == "success" || scenario == "provider-error" || scenario == "protocol-failure" || scenario == "missing-code" {
				before := p.calls
				replay := callback(url.Values{"state": {p.state}, "code": {"simulated-code"}})
				for _, c := range replay.Result().Cookies() {
					if c.Name == "kencleng_session" {
						t.Fatal("consumed pending state replay minted session")
					}
				}
				if p.calls != before {
					t.Fatal("consumed state reached protocol again")
				}
			}
		})
	}
}

// Wait for an actual PostgreSQL blocker edge, not an elapsed sleep or Go race.
func waitBlocked(t *testing.T, f harnessFixture, blocker uint32, waiter uint32) {
	t.Helper()
	deadline := time.NewTimer(3 * time.Second)
	defer deadline.Stop()
	tick := time.NewTicker(10 * time.Millisecond)
	defer tick.Stop()
	for {
		var blocked bool
		var err error
		if waiter != 0 {
			err = f.provisioner.QueryRow(f.ctx, `SELECT $1::int=ANY(pg_blocking_pids($2::int))`, blocker, waiter).Scan(&blocked)
		} else {
			err = f.provisioner.QueryRow(f.ctx, `SELECT EXISTS(SELECT 1 FROM pg_stat_activity WHERE usename='c1_runtime' AND $1::int=ANY(pg_blocking_pids(pid)))`, blocker).Scan(&blocked)
		}
		if err != nil {
			t.Fatal("lock witness query failed")
		}
		if blocked {
			return
		}
		select {
		case <-deadline.C:
			t.Fatal("expected PostgreSQL blocker edge not observed")
		case <-tick.C:
		case <-f.ctx.Done():
			t.Fatal("lock witness context expired")
		}
	}
}

func TestHarnessConcurrentPerson(t *testing.T) {
	f := harness(t)
	subject, err := newUUID()
	if err != nil {
		t.Fatal("subject fixture failed")
	}
	subject = "Case-" + subject
	holder, err := f.runtime.Begin(f.ctx)
	if err != nil {
		t.Fatal("runtime begin failed")
	}
	defer holder.Rollback(context.Background())
	id, err := newUUID()
	if err != nil {
		t.Fatal("person fixture failed")
	}
	if _, err = holder.Exec(f.ctx, `INSERT INTO persons(id,issuer,subject) VALUES($1,$2,$3)`, id, googleIssuer, subject); err != nil {
		t.Fatal("runtime first uncommitted person failed")
	}
	type result struct {
		id  string
		err error
	}
	results := make(chan result, 8)
	start := make(chan struct{})
	for i := 0; i < 8; i++ {
		go func() {
			<-start
			id, err := f.store.UpsertPerson(f.ctx, googleIssuer, subject)
			results <- result{id, err}
		}()
	}
	close(start)
	waitBlocked(t, f, holder.Conn().PgConn().PID(), 0)
	if err = holder.Commit(f.ctx); err != nil {
		t.Fatal("runtime first-login commit failed")
	}
	for i := 0; i < 8; i++ {
		select {
		case r := <-results:
			if r.err != nil || r.id != id {
				t.Fatal("concurrent identity not stable")
			}
		case <-f.ctx.Done():
			t.Fatal("concurrent upsert timeout")
		}
	}
	var count int
	if err = f.runtime.QueryRow(f.ctx, `SELECT count(*) FROM persons WHERE issuer=$1 AND subject=$2`, googleIssuer, subject).Scan(&count); err != nil || count != 1 {
		t.Fatal("duplicate identity invariant failed")
	}
	for _, pair := range [][2]string{{googleIssuer, strings.ToUpper(subject)}, {"https://other.invalid", subject}} {
		other, err := f.store.UpsertPerson(f.ctx, pair[0], pair[1])
		if err != nil || other == id {
			t.Fatal("issuer/case boundary collapsed")
		}
	}
	t.Log("eight runtime upserts; actual first-login blocker edge witnessed; one stable person; issuer/case distinct")
}

func TestHarnessSessionLockContention(t *testing.T) {
	for _, scenario := range []string{"expiry", "revocation", "cancellation", "logout", "rotation"} {
		t.Run(scenario, func(t *testing.T) {
			f := harness(t)
			token, s := f.session(t)
			holder, err := f.runtime.Begin(f.ctx)
			if err != nil {
				t.Fatal("runtime holder begin failed")
			}
			defer holder.Rollback(context.Background())
			if _, err = LockSession(f.ctx, holder, token, time.Now()); err != nil {
				t.Fatal("runtime holder lock failed")
			}
			if scenario == "expiry" {
				if _, err = holder.Exec(f.ctx, `UPDATE sessions SET created_at=$2,expires_at=$3 WHERE token_digest=$1`, digest(token), time.Now().Add(-time.Hour), time.Now().Add(time.Second)); err != nil {
					t.Fatal("fixture expiry update failed")
				}
			}
			if scenario == "revocation" {
				if _, err = holder.Exec(f.ctx, `UPDATE sessions SET revoked_at=now() WHERE token_digest=$1`, digest(token)); err != nil {
					t.Fatal("fixture revoke update failed")
				}
			}
			type result struct {
				s     Session
				token string
				err   error
			}
			done := make(chan result, 1)
			ctx, cancel := context.WithCancel(f.ctx)
			defer cancel()
			var waiter pgx.Tx
			if scenario == "logout" {
				go func() { done <- result{err: f.store.RevokeSession(ctx, token, time.Now())} }()
			} else if scenario == "rotation" {
				go func() {
					next, err := f.store.RotateSession(ctx, s.PersonID, token, time.Now())
					done <- result{token: next, err: err}
				}()
			} else {
				waiter, err = f.runtime.Begin(ctx)
				if err != nil {
					t.Fatal("runtime waiter begin failed")
				}
				defer waiter.Rollback(context.Background())
				go func() { s, err := LockSession(ctx, waiter, token, time.Now()); done <- result{s: s, err: err} }()
			}
			var waiterPID uint32
			if waiter != nil {
				waiterPID = waiter.Conn().PgConn().PID()
			}
			waitBlocked(t, f, holder.Conn().PgConn().PID(), waiterPID)
			if scenario == "expiry" {
				// Bound release to the stored deadline, after the blocker edge exists.
				var expires time.Time
				if err = holder.QueryRow(f.ctx, `SELECT expires_at FROM sessions WHERE token_digest=$1`, digest(token)).Scan(&expires); err != nil {
					t.Fatal("fixture expiry witness failed")
				}
				timer := time.NewTimer(time.Until(expires) + 20*time.Millisecond)
				defer timer.Stop()
				select {
				case <-timer.C:
				case <-f.ctx.Done():
					t.Fatal("expiry barrier timeout")
				}
			}
			if scenario == "cancellation" {
				cancel()
			} else {
				if err = holder.Commit(f.ctx); err != nil {
					t.Fatal("runtime lock release commit failed")
				}
			}
			var got result
			select {
			case got = <-done:
			case <-f.ctx.Done():
				t.Fatal("session operation timeout")
			}
			switch scenario {
			case "expiry", "revocation":
				if !errors.Is(got.err, ErrUnauthenticated) || got.s != (Session{}) {
					t.Fatal("under-lock invalid session returned authority")
				}
			case "cancellation":
				if !errors.Is(got.err, context.Canceled) || got.s != (Session{}) {
					t.Fatal("cancelled lock did not release waiter without authority")
				}
			case "logout", "rotation":
				if got.err != nil {
					t.Fatal("serialized session operation failed")
				}
				if _, err = f.store.ReadSession(f.ctx, token, time.Now()); !errors.Is(err, ErrUnauthenticated) {
					t.Fatal("old session remains valid")
				}
				if scenario == "rotation" {
					newSession, err := f.store.ReadSession(f.ctx, got.token, time.Now())
					if err != nil || newSession.PersonID != s.PersonID {
						t.Fatal("rotated session invalid")
					}
				}
			}
			t.Log(fmt.Sprintf("%s: runtime operations, PostgreSQL blocking edge observed, bounded release, asserted authority outcome", scenario))
		})
	}
}

// OpenAPI requires an object, so JSON null must not revoke a session.
func TestHarnessLogoutRejectsNull(t *testing.T) {
	f := harness(t)
	h, mux, _ := harnessHandler(t, f)
	token, session := f.session(t)
	w := harnessRequest(mux, "POST", "/auth/logout", "null", token, session.CSRFToken, h.Browser.Origin, "application/json")
	safeResponse(t, w, 400, "invalid_request")
	if _, err := f.store.ReadSession(f.ctx, token, time.Now()); err != nil {
		t.Fatal("JSON null revoked session despite object-only contract")
	}
}
