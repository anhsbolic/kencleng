package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/anhsbolic/kencleng/backend/internal/domain/account"
)

// newMFAReq builds an authenticated-ish request (session user injected via
// context) targeting the given path with an optional JSON body.
func newMFAReq(t *testing.T, path, body string) *http.Request {
	t.Helper()
	var rdr *bytes.Reader
	if body == "" {
		rdr = bytes.NewReader(nil)
	} else {
		rdr = bytes.NewReader([]byte(body))
	}
	req := httptest.NewRequest(http.MethodPost, path, rdr)
	req.Header.Set("Content-Type", "application/json")
	return req
}

// ---- R3 wire: enroll 409 problem type + happy 200 -------------------------

func TestMfaEnrollHandler_Returns200WithUri(t *testing.T) {
	w := httptest.NewRecorder()
	stub := &stubSecurityService{enrollURI: "otpauth://totp/Kencleng:u?secret=JBS"}
	req := newMFAReq(t, "/account/security/mfa/enroll", "")
	req = req.WithContext(context.WithValue(req.Context(), sessionUserIDKey, uuid.New()))
	MfaEnrollHandler(stub)(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}
	var body map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &body)
	if body["otpauth_uri"] != "otpauth://totp/Kencleng:u?secret=JBS" {
		t.Errorf("otpauth_uri = %q", body["otpauth_uri"])
	}
}

func TestMfaEnrollHandler_AlreadyEnabled_409(t *testing.T) {
	w := httptest.NewRecorder()
	stub := &stubSecurityService{enrollErr: account.ErrMfaAlreadyEnabled}
	req := newMFAReq(t, "/account/security/mfa/enroll", "")
	req = req.WithContext(context.WithValue(req.Context(), sessionUserIDKey, uuid.New()))
	MfaEnrollHandler(stub)(w, req)

	if w.Code != http.StatusConflict {
		t.Fatalf("status = %d, want 409", w.Code)
	}
	if !strings.Contains(w.Body.String(), "mfa-already-enabled") {
		t.Errorf("body missing problem type mfa-already-enabled: %s", w.Body.String())
	}
}

// ---- R6/R7 wire: confirm 422 identical shape for wrong-code + no-pending ---

func TestMfaEnrollConfirmHandler_ReturnsBackupCodes(t *testing.T) {
	w := httptest.NewRecorder()
	uid := uuid.New()
	stub := &stubSecurityService{confirmCodes: []string{"aaa", "bbb"}}
	req := newMFAReq(t, "/account/security/mfa/enroll/confirm", `{"totp_code":"123456"}`)
	req = req.WithContext(context.WithValue(req.Context(), sessionUserIDKey, uid))
	MfaEnrollConfirmHandler(stub)(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}
	var body map[string][]string
	_ = json.Unmarshal(w.Body.Bytes(), &body)
	if len(body["backup_codes"]) != 2 {
		t.Errorf("backup_codes = %v", body["backup_codes"])
	}
}

func TestMfaEnrollConfirmHandler_WrongCode_And_NoPending_Identical422(t *testing.T) {
	uid := uuid.New()

	wrong := httptest.NewRecorder()
	stubWrong := &stubSecurityService{confirmErr: account.ErrInvalidTOTPCode}
	reqWrong := newMFAReq(t, "/account/security/mfa/enroll/confirm", `{"totp_code":"000000"}`)
	reqWrong = reqWrong.WithContext(context.WithValue(reqWrong.Context(), sessionUserIDKey, uid))
	MfaEnrollConfirmHandler(stubWrong)(wrong, reqWrong)

	noPending := httptest.NewRecorder()
	stubNoPending := &stubSecurityService{confirmErr: account.ErrMfaNotPending}
	reqNoPending := newMFAReq(t, "/account/security/mfa/enroll/confirm", `{"totp_code":"000000"}`)
	reqNoPending = reqNoPending.WithContext(context.WithValue(reqNoPending.Context(), sessionUserIDKey, uid))
	MfaEnrollConfirmHandler(stubNoPending)(noPending, reqNoPending)

	if wrong.Code != http.StatusUnprocessableEntity || noPending.Code != http.StatusUnprocessableEntity {
		t.Fatalf("statuses = %d,%d want 422,422", wrong.Code, noPending.Code)
	}
	if wrong.Body.String() != noPending.Body.String() {
		t.Errorf("wrong-code body != no-pending body (R7 byte-identical violated):\n%s\nvs\n%s",
			wrong.Body.String(), noPending.Body.String())
	}
}

func TestMfaEnrollConfirmHandler_MissingCode_422(t *testing.T) {
	w := httptest.NewRecorder()
	stub := &stubSecurityService{}
	req := newMFAReq(t, "/account/security/mfa/enroll/confirm", `{}`)
	req = req.WithContext(context.WithValue(req.Context(), sessionUserIDKey, uuid.New()))
	MfaEnrollConfirmHandler(stub)(w, req)
	if w.Code != http.StatusUnprocessableEntity {
		t.Fatalf("status = %d, want 422", w.Code)
	}
	if !strings.Contains(w.Body.String(), "totp_code") {
		t.Errorf("missing totp_code field error: %s", w.Body.String())
	}
}

// ---- R12/R14 wire: disable email_password vs google-only marker ------------

func TestMfaDisableHandler_EmailPassword_SuccessAnd401(t *testing.T) {
	uid := uuid.New()

	// Success (correct password).
	w := httptest.NewRecorder()
	stub := &stubSecurityService{reauthRequired: false}
	req := newMFAReq(t, "/account/security/mfa/disable", `{"password":"correct-horse"}`)
	req = req.WithContext(context.WithValue(req.Context(), sessionUserIDKey, uid))
	MfaDisableHandler(stub)(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("success status = %d, want 200", w.Code)
	}

	// Wrong password → service returns ErrInvalidCredentials → 401.
	w2 := httptest.NewRecorder()
	stub2 := &stubSecurityService{reauthRequired: false, disableErr: account.ErrInvalidCredentials}
	req2 := newMFAReq(t, "/account/security/mfa/disable", `{"password":"wrong"}`)
	req2 = req2.WithContext(context.WithValue(req2.Context(), sessionUserIDKey, uid))
	MfaDisableHandler(stub2)(w2, req2)
	if w2.Code != http.StatusUnauthorized {
		t.Fatalf("wrong-password status = %d, want 401", w2.Code)
	}
}

func TestMfaDisableHandler_EmailPassword_MissingPassword_422(t *testing.T) {
	w := httptest.NewRecorder()
	stub := &stubSecurityService{reauthRequired: false}
	req := newMFAReq(t, "/account/security/mfa/disable", `{}`)
	req = req.WithContext(context.WithValue(req.Context(), sessionUserIDKey, uuid.New()))
	MfaDisableHandler(stub)(w, req)
	if w.Code != http.StatusUnprocessableEntity {
		t.Fatalf("status = %d, want 422", w.Code)
	}
	if !strings.Contains(w.Body.String(), "password") {
		t.Errorf("missing password field error: %s", w.Body.String())
	}
	if stub.disableCalled {
		t.Error("service must not be called when password missing")
	}
}

func TestMfaDisableHandler_GoogleOnly_ConsumesMarker(t *testing.T) {
	uid := uuid.New()

	// No marker → 401, service not reached.
	w := httptest.NewRecorder()
	stub := &stubSecurityService{reauthRequired: true}
	req := newMFAReq(t, "/account/security/mfa/disable", `{}`)
	req = req.WithContext(context.WithValue(req.Context(), sessionUserIDKey, uid))
	MfaDisableHandler(stub)(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("no-marker status = %d, want 401", w.Code)
	}
	if stub.disableCalled {
		t.Error("service must not be reached without a valid marker")
	}

	// Set a valid marker → 200, marker consumed.
	SetReauthMarker(uid, time.Now().Add(time.Minute))
	stub2 := &stubSecurityService{reauthRequired: true}
	w2 := httptest.NewRecorder()
	req2 := newMFAReq(t, "/account/security/mfa/disable", `{}`)
	req2 = req2.WithContext(context.WithValue(req2.Context(), sessionUserIDKey, uid))
	MfaDisableHandler(stub2)(w2, req2)
	if w2.Code != http.StatusOK {
		t.Fatalf("with-marker status = %d, want 200 (%s)", w2.Code, w2.Body.String())
	}
	// Marker must be consumed (one-shot).
	if ConsumeReauthMarker(uid) {
		t.Error("marker must be consumed on use, not replayable")
	}

	// Second disable → marker gone → 401.
	w3 := httptest.NewRecorder()
	stub3 := &stubSecurityService{reauthRequired: true}
	req3 := newMFAReq(t, "/account/security/mfa/disable", `{}`)
	req3 = req3.WithContext(context.WithValue(req3.Context(), sessionUserIDKey, uid))
	MfaDisableHandler(stub3)(w3, req3)
	if w3.Code != http.StatusUnauthorized {
		t.Fatalf("replay status = %d, want 401", w3.Code)
	}
}

func TestMfaDisableHandler_GoogleOnly_PasswordDoesNotBypass(t *testing.T) {
	// R14: a Google-only caller submitting a password still needs a marker.
	uid := uuid.New()
	w := httptest.NewRecorder()
	stub := &stubSecurityService{reauthRequired: true}
	req := newMFAReq(t, "/account/security/mfa/disable", `{"password":"attacker-guess"}`)
	req = req.WithContext(context.WithValue(req.Context(), sessionUserIDKey, uid))
	MfaDisableHandler(stub)(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("google-only + password (no marker) status = %d, want 401", w.Code)
	}
	if stub.disableCalled {
		t.Error("password must not bypass the marker requirement for google-only users")
	}
}

// ---- R16: handlers are session-guarded (401 without session user) ----------

func TestMfaHandlers_RequireSession(t *testing.T) {
	// No session user in context (as if RequireSession never ran) → each
	// handler 401s before touching the service.
	for _, tc := range []struct {
		path  string
		body  string
		build func() securityService
	}{
		{"/account/security/mfa/enroll", "", func() securityService { return &stubSecurityService{} }},
		{"/account/security/mfa/enroll/confirm", `{"totp_code":"123456"}`, func() securityService { return &stubSecurityService{} }},
		{"/account/security/mfa/disable", `{}`, func() securityService { return &stubSecurityService{} }},
	} {
		w := httptest.NewRecorder()
		var h http.HandlerFunc
		switch tc.path {
		case "/account/security/mfa/enroll":
			h = MfaEnrollHandler(tc.build())
		case "/account/security/mfa/enroll/confirm":
			h = MfaEnrollConfirmHandler(tc.build())
		default:
			h = MfaDisableHandler(tc.build())
		}
		req := newMFAReq(t, tc.path, tc.body) // no session user injected
		h(w, req)
		if w.Code != http.StatusUnauthorized {
			t.Errorf("%s without session => %d, want 401", tc.path, w.Code)
		}
	}
}
