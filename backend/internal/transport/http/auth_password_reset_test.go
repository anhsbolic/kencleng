package http

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/anhsbolic/kencleng/backend/internal/domain/account"
)

// captureLogOutput redirects the standard logger into a buffer and
// returns a func reading everything logged so far.
func captureLogOutput(t *testing.T) func() string {
	t.Helper()
	var buf strings.Builder
	orig := log.Writer()
	log.SetOutput(&buf)
	t.Cleanup(func() { log.SetOutput(orig) })
	return buf.String
}

// stubForgotResetService records calls and returns scripted results.
type stubForgotResetService struct {
	forgotCalls []string
	resetCalls  [][2]string

	forgotErr error
	resetErr  error
}

func (s *stubForgotResetService) ForgotPassword(_ context.Context, email string) error {
	s.forgotCalls = append(s.forgotCalls, email)
	return s.forgotErr
}

func (s *stubForgotResetService) ResetPassword(_ context.Context, token, newPassword string) error {
	s.resetCalls = append(s.resetCalls, [2]string{token, newPassword})
	return s.resetErr
}

// Compile-time check: *account.Service must satisfy both ports.
var (
	_ forgotPasswordService = (*account.Service)(nil)
	_ resetPasswordService  = (*account.Service)(nil)
)

// R5: all three forgot branches return nil from the service, so the
// handler's output is byte-identical regardless of branch.
func TestForgotPassword_Handler_Generic202_AllBranches(t *testing.T) {
	cases := []struct {
		name string
		svc  *stubForgotResetService
	}{
		{"registered", &stubForgotResetService{}},
		{"google-only", &stubForgotResetService{}},
		{"unknown", &stubForgotResetService{}},
	}
	var first []byte
	for i, tc := range cases {
		rr := postJSON(t, ForgotPasswordHandler(tc.svc), "/auth/forgot-password", `{"email":"someone@example.com"}`)

		if rr.Code != http.StatusAccepted {
			t.Errorf("%s: status = %d, want 202", tc.name, rr.Code)
		}
		body := rr.Body.Bytes()
		var parsed map[string]string
		if err := json.Unmarshal(body, &parsed); err != nil {
			t.Fatalf("%s: body not JSON object: %v", tc.name, err)
		}
		if _, ok := parsed["message"]; !ok {
			t.Errorf("%s: body missing generic message field", tc.name)
		}
		if i == 0 {
			first = body
		} else if string(body) != string(first) {
			t.Errorf("%s: body %q differs from registered branch body %q", tc.name, body, first)
		}
	}
}

// Internal failures are swallowed into the identical 202 with a sanitized
// server-side log (no recipient in output).
func TestForgotPassword_Handler_InternalError_Still202_ButLogs(t *testing.T) {
	const recipient = "internal-error@example.com"
	svc := &stubForgotResetService{forgotErr: errors.New("db connection lost")}

	logs := captureLogOutput(t)

	rr := postJSON(t, ForgotPasswordHandler(svc), "/auth/forgot-password", `{"email":"`+recipient+`"}`)

	if rr.Code != http.StatusAccepted {
		t.Errorf("status = %d, want 202 (a 500 would leak match vs no-match)", rr.Code)
	}
	if !strings.Contains(logs(), "forgot password failed") {
		t.Errorf("expected failure log line, got %q", logs())
	}
	if strings.Contains(logs(), recipient) {
		t.Errorf("log leaked recipient email: %q", logs())
	}
	if len(svc.forgotCalls) != 1 || svc.forgotCalls[0] != recipient {
		t.Errorf("service called with wrong email: %v", svc.forgotCalls)
	}
}

func TestForgotPassword_Handler_MalformedBodyAndEmail(t *testing.T) {
	svc := &stubForgotResetService{}

	rr := postJSON(t, ForgotPasswordHandler(svc), "/auth/forgot-password", `{not-json`)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("malformed json: status = %d, want 400", rr.Code)
	}
	if len(svc.forgotCalls) != 0 {
		t.Error("malformed body must not reach the service")
	}

	rr = postJSON(t, ForgotPasswordHandler(svc), "/auth/forgot-password", `{"email":"not-an-email"}`)
	if rr.Code != http.StatusUnprocessableEntity {
		t.Errorf("malformed email: status = %d, want 422", rr.Code)
	}
	if len(svc.forgotCalls) != 0 {
		t.Error("malformed email must not reach the service")
	}
}

// R6: burst-exceeding requests through the shared limiter → 429.
func TestForgotPassword_RateLimited(t *testing.T) {
	svc := &stubForgotResetService{}
	h := RateLimit(0.001, 2)(ForgotPasswordHandler(svc)).ServeHTTP

	var last *httptest.ResponseRecorder
	for i := 0; i < 5; i++ {
		last = postJSON(t, h, "/auth/forgot-password", `{"email":"rl@example.com"}`)
	}
	if last.Code != http.StatusTooManyRequests {
		t.Fatalf("status = %d, want 429 once burst is exhausted", last.Code)
	}
}

func TestResetPassword_RateLimited(t *testing.T) {
	svc := &stubForgotResetService{}
	h := RateLimit(0.001, 2)(ResetPasswordHandler(svc)).ServeHTTP

	var last *httptest.ResponseRecorder
	for i := 0; i < 5; i++ {
		last = postJSON(t, h, "/auth/reset-password", `{"token":"t","new_password":"good-password-1"}`)
	}
	if last.Code != http.StatusTooManyRequests {
		t.Fatalf("status = %d, want 429 once burst is exhausted", last.Code)
	}
}

// R15: empty token → 404 without reaching the service.
func TestResetPassword_EmptyToken_404_NoServiceCall(t *testing.T) {
	svc := &stubForgotResetService{}

	rr := postJSON(t, ResetPasswordHandler(svc), "/auth/reset-password", `{"token":"","new_password":"good-password-1"}`)

	if rr.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404", rr.Code)
	}
	if len(svc.resetCalls) != 0 {
		t.Error("empty token must not reach the service")
	}
}

// Status mapping passthrough: sentinel errors land on their contract codes.
func TestResetPassword_Handler_StatusMapping(t *testing.T) {
	cases := []struct {
		name     string
		resetErr error
		password string
		wantCode int
	}{
		{"expired", account.ErrTokenExpired, "good-password-1", http.StatusGone},
		{"not-found", account.ErrTokenNotFound, "good-password-1", http.StatusNotFound},
		{"validation-failed", account.ErrValidation, "short", http.StatusUnprocessableEntity},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			svc := &stubForgotResetService{resetErr: tc.resetErr}
			body := `{"token":"some-token","new_password":"` + tc.password + `"}`

			rr := postJSON(t, ResetPasswordHandler(svc), "/auth/reset-password", body)
			if rr.Code != tc.wantCode {
				t.Fatalf("status = %d, want %d (body: %s)", rr.Code, tc.wantCode, rr.Body.String())
			}
			ct := rr.Header().Get("Content-Type")
			if !strings.Contains(ct, "application/problem+json") {
				t.Errorf("problem content-type = %q", ct)
			}
		})
	}
}

// R7 at the handler boundary: success maps to 200 with the contract's
// message body, and the service received the raw request fields.
func TestResetPassword_Handler_Success_200(t *testing.T) {
	svc := &stubForgotResetService{}

	rr := postJSON(t, ResetPasswordHandler(svc), "/auth/reset-password", `{"token":"tok","new_password":"good-password-1"}`)

	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200 (body: %s)", rr.Code, rr.Body.String())
	}
	var parsed map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &parsed); err != nil {
		t.Fatalf("body not JSON: %v", err)
	}
	if parsed["message"] == "" {
		t.Error("success body must carry a message field")
	}
	if len(svc.resetCalls) != 1 || svc.resetCalls[0] != [2]string{"tok", "good-password-1"} {
		t.Errorf("service args = %v", svc.resetCalls)
	}
}

// R16: malformed JSON on reset → 400.
func TestResetPassword_Handler_MalformedJSON_400(t *testing.T) {
	svc := &stubForgotResetService{}
	rr := postJSON(t, ResetPasswordHandler(svc), "/auth/reset-password", `{bad`)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", rr.Code)
	}
}
