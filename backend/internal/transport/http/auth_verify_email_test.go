package http

import (
	"bytes"
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/anhsbolic/kencleng/backend/internal/domain/account"
	"github.com/anhsbolic/kencleng/backend/internal/platform/crypto"
)

// fakeResendRepo is a minimal account.Repository for handler tests. It
// embeds the interface (nil — unneeded methods panic only if called) and
// overrides FindAuthIdentityByIdentifierHash to inject a configurable
// error, so ResendVerification returns an error without needing a real DB.
type fakeResendRepo struct {
	account.Repository // nil; only the overridden method is called
	findErr            error
}

func (f *fakeResendRepo) FindAuthIdentityByIdentifierHash(_ context.Context, _, _ string) (*account.AuthIdentity, error) {
	return nil, f.findErr
}

// TestResendVerificationHandler_ServiceError_Still202_ButLogs proves the
// S4 fix: when the service returns an error from ResendVerification, the
// handler still writes 202 (anti-enumeration — a 500 would distinguish
// match from no-match) AND logs the error server-side so ops can see the
// failure. The log must not contain the recipient email (PII).
func TestResendVerificationHandler_ServiceError_Still202_ButLogs(t *testing.T) {
	const recipient = "leak@example.com"

	// Wire a real account.Service with a fake repo whose
	// FindAuthIdentityByIdentifierHash returns an error, so
	// ResendVerification returns a wrapped error. db/breachcheck/sender
	// are nil — they're not reached on this error path.
	keys := &crypto.Keys{
		EncryptionKey: make([]byte, 32),
		HMACKey:       make([]byte, 32),
	}
	repo := &fakeResendRepo{findErr: errors.New("db connection lost")}
	svc := account.NewService(repo, nil, nil, nil, keys, nil, nil, "http://localhost:3000", nil, nil, nil, nil, nil, nil)

	// Capture log output.
	var logBuf bytes.Buffer
	origOut := log.Writer()
	log.SetOutput(&logBuf)
	defer log.SetOutput(origOut)

	body := `{"email":"` + recipient + `"}`
	req := httptest.NewRequest(http.MethodPost, "/auth/verify-email/resend", strings.NewReader(body))
	rr := httptest.NewRecorder()

	ResendVerificationHandler(svc).ServeHTTP(rr, req)

	// Anti-enumeration: response must still be 202, not 500.
	if rr.Code != http.StatusAccepted {
		t.Errorf("status = %d, want 202 (anti-enumeration: a 500 would leak match vs no-match)", rr.Code)
	}

	// The failure must be logged so ops can see it.
	logOut := logBuf.String()
	if !strings.Contains(logOut, "resend verification failed") {
		t.Errorf("expected failure log line, got: %q", logOut)
	}

	// The log must not leak the recipient email (PII).
	if strings.Contains(logOut, recipient) {
		t.Errorf("log leaked recipient email %q: %q", recipient, logOut)
	}
}

// Compile-time assertion that fakeResendRepo satisfies account.Repository.
// The embedding makes the compiler check the interface even though the
// embedded value is nil at runtime; the methods that are actually called
// are overridden.
var _ account.Repository = (*fakeResendRepo)(nil)

// Silence unused imports that the compiler checks for in the test build
// (pgx, uuid, time are needed by the full interface signature even though
// only FindAuthIdentityByIdentifierHash is overridden in this test).
var (
	_ pgx.Tx
	_ uuid.UUID
	_ time.Time
)
