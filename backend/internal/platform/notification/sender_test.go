package notification

import (
	"bytes"
	"context"
	"log"
	"strings"
	"testing"
)

// TestFakeSender_SendVerificationEmail_ReturnsNilNoPIIInLog verifies the
// call succeeds and the log line contains neither the recipient address
// nor the verification token.
func TestFakeSender_SendVerificationEmail_ReturnsNilNoPIIInLog(t *testing.T) {
	var logBuf bytes.Buffer
	origOut := log.Writer()
	log.SetOutput(&logBuf)
	defer log.SetOutput(origOut)

	s := NewFakeSender()
	recipient := "user@example.com"
	token := "super-secret-token-value"

	if err := s.SendVerificationEmail(context.Background(), recipient, token); err != nil {
		t.Fatalf("SendVerificationEmail: unexpected error: %v", err)
	}

	logged := logBuf.String()
	if strings.Contains(logged, recipient) {
		t.Fatalf("log must not contain recipient email; got: %q", logged)
	}
	if strings.Contains(logged, token) {
		t.Fatalf("log must not contain the verification token; got: %q", logged)
	}
}

// TestFakeSender_SendNudgeEmail_ReturnsNilNoPIIInLog verifies the call
// succeeds, the recipient address is absent from the log, and the nudge
// type (non-PII) is present.
func TestFakeSender_SendNudgeEmail_ReturnsNilNoPIIInLog(t *testing.T) {
	var logBuf bytes.Buffer
	origOut := log.Writer()
	log.SetOutput(&logBuf)
	defer log.SetOutput(origOut)

	s := NewFakeSender()
	recipient := "user@example.com"

	if err := s.SendNudgeEmail(context.Background(), recipient, NudgePasswordReset); err != nil {
		t.Fatalf("SendNudgeEmail: unexpected error: %v", err)
	}

	logged := logBuf.String()
	if strings.Contains(logged, recipient) {
		t.Fatalf("log must not contain recipient email; got: %q", logged)
	}
	if !strings.Contains(logged, NudgePasswordReset) {
		t.Fatalf("log should contain nudge type %q; got: %q", NudgePasswordReset, logged)
	}
}
