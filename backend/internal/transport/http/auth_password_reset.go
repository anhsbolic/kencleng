package http

import (
	"encoding/json"
	"log"
	"net/http"

	"context"
)

// Narrow service ports for these handlers (same convention as
// loginSessionService): the transport layer depends on the two methods it
// calls, not on the whole account.Service surface.
type forgotPasswordService interface {
	ForgotPassword(ctx context.Context, email string) error
}

type resetPasswordService interface {
	ResetPassword(ctx context.Context, token, newPassword string) error
}

type forgotPasswordRequest struct {
	Email string `json:"email"`
}

type resetPasswordRequest struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

// ForgotPasswordHandler handles POST /auth/forgot-password.
// Anti-enumeration: the response is the identical 202 generic regardless
// of whether the email is registered, Google-only, or unknown (R1-R3).
// An internal failure must NOT surface as a 500 — that would distinguish
// the registered branch from the no-op branches — so it is swallowed into
// the same 202 and logged server-side with a sanitized chain (the leaf is
// a DB driver error: SQLSTATE/constraint name, parameterized SQL, no PII
// values). Same pattern as ResendVerificationHandler.
func ForgotPasswordHandler(svc forgotPasswordService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req forgotPasswordRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			write400InvalidJSON(w)
			return
		}

		if !looksLikeEmail(req.Email) {
			WriteValidationError(w, []fieldError{
				{Field: "email", Message: "must be a valid email"},
			})
			return
		}

		if err := svc.ForgotPassword(r.Context(), req.Email); err != nil {
			log.Printf("transport: forgot password failed (recipient redacted): %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"message": "Kalau email terdaftar, instruksi sudah dikirim.",
		})
	}
}

// ResetPasswordHandler handles POST /auth/reset-password.
// R7 valid → 200. R8 policy failure → 422 with the token left unconsumed
// (validation runs before redemption in the service — spec Assumption B).
// R9 expired → 410. R10 not-found/used/wrong-purpose → 404.
func ResetPasswordHandler(svc resetPasswordService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req resetPasswordRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			write400InvalidJSON(w)
			return
		}

		// Reject an empty token at the boundary — saves a DB round-trip
		// and avoids any timing distinction between "empty" and "not
		// found" (same discipline as VerifyEmailHandler).
		if req.Token == "" {
			WriteProblem(w, http.StatusNotFound,
				"https://kencleng.dev/problems/token-not-found",
				"Token Not Found", "The verification token was not found.")
			return
		}

		if err := svc.ResetPassword(r.Context(), req.Token, req.NewPassword); err != nil {
			MapServiceError(w, err) // 422 / 410 / 404 per sentinel mapping
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{
			"message": "Password berhasil diubah. Silakan login ulang.",
		})
	}
}
