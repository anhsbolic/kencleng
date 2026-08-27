package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"

	"github.com/anhsbolic/kencleng/backend/internal/domain/account"
)

// sessionUserIDKey is the context key under which RequireSession stores
// the authenticated user's UUID. Handlers read it via UserIDFromContext.
type contextKey string

const sessionUserIDKey contextKey = "sessionUserID"

// UserIDFromContext extracts the authenticated user's ID from the request
// context. Returns ok=false if the context has no session binding (the
// handler was reached without passing through RequireSession).
func UserIDFromContext(ctx context.Context) (uuid.UUID, bool) {
	uid, ok := ctx.Value(sessionUserIDKey).(uuid.UUID)
	return uid, ok
}

// RequireSession returns middleware that verifies the presented access
// token (HttpOnly cookie first, Authorization: Bearer fallback) via the
// provided verifier, and injects the authenticated user_id into the
// request context. Requests without a valid token receive a 401 Problem
// Details response and never reach the wrapped handler.
//
// The verifier is expected to be ES256SessionVerifier (or GoogleTokenVerifier,
// which is the same function — both verify ES256 JWTs against the app's
// public key and return the subject as a user UUID). platform/auth is
// Tier 0 fenced and is NOT modified; verification lives inline at the
// transport boundary (task-02 precedent).
func RequireSession(verifier func(string) (uuid.UUID, error)) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := sessionToken(r)
			if token == "" {
				WriteProblem(w, http.StatusUnauthorized,
					"https://kencleng.dev/errors/unauthorized",
					"Unauthorized", "Authentication required.")
				return
			}
			userID, err := verifier(token)
			if err != nil {
				WriteProblem(w, http.StatusUnauthorized,
					"https://kencleng.dev/errors/unauthorized",
					"Unauthorized", "Authentication required.")
				return
			}
			ctx := context.WithValue(r.Context(), sessionUserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// securityService is the subset of *account.Service the security handlers
// depend on. *account.Service satisfies it; tests inject a stub so the
// transport contract (status codes, problem types) is exercisable without
// wiring the full domain — same seam philosophy as googleOAuthService.
type securityService interface {
	SetPassword(ctx context.Context, userID uuid.UUID, email, currentPassword, newPassword string) (bool, error)
	UnlinkGoogle(ctx context.Context, userID uuid.UUID, password string) error
	MfaEnroll(ctx context.Context, userID uuid.UUID) (string, error)
	MfaEnrollConfirm(ctx context.Context, userID uuid.UUID, totpCode string) ([]string, error)
	MfaDisable(ctx context.Context, userID uuid.UUID, password string) error
	MfaDisableReauthRequired(ctx context.Context, userID uuid.UUID) (bool, error)
}

type setPasswordRequest struct {
	Email           string `json:"email,omitempty"`
	CurrentPassword string `json:"current_password,omitempty"`
	Password        string `json:"password"`
}

type unlinkGoogleRequest struct {
	Password string `json:"password"`
}

// SetPasswordHandler handles POST /account/security/set-password.
// The server selects Branch 1 (add identity) vs Branch 2 (change password)
// based on the caller's existing identities — never a client-supplied
// flag. Branch 1 always returns a generic 202 (anti-enumeration); Branch 2
// returns 200 on success or 401 on wrong current_password.
func SetPasswordHandler(svc securityService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := UserIDFromContext(r.Context())
		if !ok {
			WriteProblem(w, http.StatusUnauthorized,
				"https://kencleng.dev/errors/unauthorized",
				"Unauthorized", "Authentication required.")
			return
		}

		var req setPasswordRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			write400InvalidJSON(w)
			return
		}

		// Boundary validation — defense-in-depth; the service also
		// checks (R4). Field names are not sensitive; values are never
		// echoed.
		var fieldErrs []fieldError
		if len(req.Password) < 8 {
			fieldErrs = append(fieldErrs, fieldError{Field: "password", Message: "must be at least 8 characters"})
		}
		if len(fieldErrs) > 0 {
			WriteValidationError(w, fieldErrs)
			return
		}

		changed, err := svc.SetPassword(r.Context(), userID, req.Email, req.CurrentPassword, req.Password)
		if err != nil {
			if isErrValidation(err) {
				WriteValidationError(w, []fieldError{
					{Field: "password", Message: "password is not allowed"},
				})
				return
			}
			// Everything else goes through the shared mapping:
			// ErrInvalidCredentials → 401 (Branch 2 wrong
			// current_password), wrapped internal errors → 500 generic.
			// A bare `default → 401` here would misclassify a DB/tx
			// failure as bad credentials (code-review S1/BP1/C1).
			MapServiceError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if changed {
			// Branch 2: password changed immediately, all sessions revoked.
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(map[string]string{
				"message": "Password berhasil diganti. Semua sesi lain telah keluar otomatis.",
			})
			return
		}
		// Branch 1: generic 202 (anti-enumeration — identical whether or
		// not the email was available).
		w.WriteHeader(http.StatusAccepted)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"message": "Kalau email tersedia, cek inbox untuk verifikasi.",
		})
	}
}

// UnlinkGoogleHandler handles POST /account/security/google/unlink.
// Requires password re-authentication. Returns 200 on success, 401 on
// wrong password, or 409 in two distinct cases (INV-account-02 /
// INV-account-12) with different problem types and Indonesian details.
func UnlinkGoogleHandler(svc securityService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := UserIDFromContext(r.Context())
		if !ok {
			WriteProblem(w, http.StatusUnauthorized,
				"https://kencleng.dev/errors/unauthorized",
				"Unauthorized", "Authentication required.")
			return
		}

		var req unlinkGoogleRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			write400InvalidJSON(w)
			return
		}

		// Boundary validation: password is required.
		if len(req.Password) < 1 {
			WriteValidationError(w, []fieldError{
				{Field: "password", Message: "required"},
			})
			return
		}

		err := svc.UnlinkGoogle(r.Context(), userID, req.Password)
		if err != nil {
			// Shared mapping: ErrOnlyIdentity / ErrRemainingUnverified →
			// distinct 409s, ErrInvalidCredentials → 401 (wrong
			// password), wrapped internal errors → 500 generic. A bare
			// `default → 401` would misclassify a DB/tx failure as bad
			// credentials (code-review S1/BP1/C1).
			MapServiceError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"message": "Akun Google berhasil dilepas.",
		})
	}
}

// MfaEnrollHandler handles POST /account/security/mfa/enroll. Returns 200
// with the otpauth:// URI on success, 409 (ErrMfaAlreadyEnabled) when MFA is
// already active. The URI embeds the TOTP secret and is never logged.
func MfaEnrollHandler(svc securityService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := UserIDFromContext(r.Context())
		if !ok {
			WriteProblem(w, http.StatusUnauthorized,
				"https://kencleng.dev/problems/unauthenticated",
				"Authentication Required", "Sign in before continuing.")
			return
		}

		uri, err := svc.MfaEnroll(r.Context(), userID)
		if err != nil {
			MapServiceError(w, err) // ErrMfaAlreadyEnabled → 409
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"otpauth_uri": uri})
	}
}

// mfaEnrollConfirmRequest mirrors openapi MfaEnrollConfirmRequest.
type mfaEnrollConfirmRequest struct {
	TOTPCode string `json:"totp_code"`
}

// MfaEnrollConfirmHandler handles POST /account/security/mfa/enroll/confirm.
// Returns 200 with the once-shown backup codes, or 422 with an
// indistinguishable per-field ValidationError for both "wrong code" and "no
// pending enrollment" (R7 — no enumeration signal on a self-targeting
// endpoint).
func MfaEnrollConfirmHandler(svc securityService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := UserIDFromContext(r.Context())
		if !ok {
			WriteProblem(w, http.StatusUnauthorized,
				"https://kencleng.dev/problems/unauthenticated",
				"Authentication Required", "Sign in before continuing.")
			return
		}

		var req mfaEnrollConfirmRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			write400InvalidJSON(w)
			return
		}
		if req.TOTPCode == "" {
			WriteValidationError(w, []fieldError{{Field: "totp_code", Message: "required"}})
			return
		}

		codes, err := svc.MfaEnrollConfirm(r.Context(), userID, req.TOTPCode)
		if err != nil {
			// R7: byte-identical 422 for wrong-code and no-pending.
			if errors.Is(err, account.ErrInvalidTOTPCode) || errors.Is(err, account.ErrMfaNotPending) {
				WriteValidationError(w, []fieldError{{
					Field:   "totp_code",
					Message: "TOTP code salah, atau tidak ada sesi pendaftaran aktif.",
				}})
				return
			}
			MapServiceError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string][]string{"backup_codes": codes})
	}
}

// mfaDisableRequest mirrors openapi MfaDisableRequest (optional password).
type mfaDisableRequest struct {
	Password string `json:"password"`
}

// MfaDisableHandler handles POST /account/security/mfa/disable.
//
// Re-authentication, server-side (R14):
//   - email_password caller: password is required (422 if absent; the service
//     verifies it, wrong → 401).
//   - Google-only caller: a currently-valid reauth marker must be present; the
//     handler atomically CONSUMES it (consume-on-use — a second call finds it
//     gone and gets 401). Any submitted password is ignored for Google-only
//     (R14: a password does not bypass the marker requirement).
//
// The marker never crosses into the domain service (D6).
func MfaDisableHandler(svc securityService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := UserIDFromContext(r.Context())
		if !ok {
			WriteProblem(w, http.StatusUnauthorized,
				"https://kencleng.dev/problems/unauthenticated",
				"Authentication Required", "Sign in before continuing.")
			return
		}

		// Tolerant decode: an empty/absent body leaves password "".
		var req mfaDisableRequest
		if r.Body != nil {
			_ = json.NewDecoder(r.Body).Decode(&req)
		}

		googleOnly, err := svc.MfaDisableReauthRequired(r.Context(), userID)
		if err != nil {
			MapServiceError(w, err)
			return
		}
		if googleOnly {
			if !ConsumeReauthMarker(userID) {
				WriteProblem(w, http.StatusUnauthorized,
					"https://kencleng.dev/errors/unauthorized",
					"Unauthorized",
					"Perlu autentikasi ulang melalui Google sebelum menonaktifkan MFA.")
				return
			}
			req.Password = "" // ignore any submitted password (R14)
		} else if req.Password == "" {
			WriteValidationError(w, []fieldError{{Field: "password", Message: "required"}})
			return
		}

		if err := svc.MfaDisable(r.Context(), userID, req.Password); err != nil {
			MapServiceError(w, err) // ErrInvalidCredentials → 401, ErrValidation → 422
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"message": "MFA berhasil dinonaktifkan."})
	}
}
