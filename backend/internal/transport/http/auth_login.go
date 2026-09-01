package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/anhsbolic/kencleng/backend/internal/domain/account"
)

// loginSessionService is the slice of the account service the four session
// endpoints depend on. *account.Service satisfies it; tests inject a stub
// so transport contract (status codes, cookie attributes, byte-equal error
// bodies) is exercisable without the full domain — same seam philosophy as
// googleOAuthService.
type loginSessionService interface {
	Login(ctx context.Context, email, password string) (account.LoginResult, error)
	LoginMfa(ctx context.Context, pendingToken, totpCode, backupCode string) (account.LoginResult, error)
	Refresh(ctx context.Context, refreshTokenPlain string) (account.RefreshResult, error)
	Logout(ctx context.Context, refreshTokenPlain string) error
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginMfaRequest struct {
	MFAPendingToken string `json:"mfa_pending_token"`
	TotpCode        string `json:"totp_code"`
	BackupCode      string `json:"backup_code"`
}

// loginOKResponse mirrors openapi LoginResponse: status "ok", access token
// in the BODY (never a cookie on these endpoints), plus the assembled user.
type loginOKResponse struct {
	Status               string        `json:"status"`
	AccessToken          string        `json:"access_token"`
	AccessTokenExpiresAt time.Time     `json:"access_token_expires_at,omitempty"`
	User                 *userResponse `json:"user,omitempty"`
}

// loginMfaRequiredResponse mirrors openapi LoginMfaRequiredResponse. The
// INFERRED marker in openapi was settled by feature-spec Assumptions A/B.
type loginMfaRequiredResponse struct {
	Status          string `json:"status"`
	MFAPendingToken string `json:"mfa_pending_token"`
}

// refreshResponse mirrors openapi RefreshResponse; the new refresh token is
// delivered implicitly via Set-Cookie, not in the body.
type refreshResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

// LoginHandler handles POST /auth/login. Anti-enumeration discipline: no
// field-level rejection that could differ by account state — malformed JSON
// gets a 400 (state-independent), everything else flows into the service
// whose sentinel vocabulary renders identical bodies for wrong-email,
// wrong-password, and lockout (only 401 vs 429 differs).
//
// Cookie rule: sets ONLY the refresh cookie (HttpOnly + Secure-conditional +
// SameSite=Strict); the access token travels in the body per contract. The
// mfa_required branch sets NO cookie and issues NO tokens (R2).
func LoginHandler(svc loginSessionService, cookieSecure bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req loginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			write400InvalidJSON(w)
			return
		}

		res, err := svc.Login(r.Context(), req.Email, req.Password)
		if err != nil {
			MapServiceError(w, err)
			return
		}

		if res.Status == "mfa_required" {
			writeJSON(w, http.StatusOK, loginMfaRequiredResponse{
				Status:          "mfa_required",
				MFAPendingToken: res.MFAPendingToken,
			})
			return
		}

		writeRefreshCookie(w, cookieSecure, res.RefreshTokenPlain)
		writeJSON(w, http.StatusOK, loginOKResponse{
			Status:               "ok",
			AccessToken:          res.AccessToken,
			AccessTokenExpiresAt: res.AccessTokenExpiresAt,
			User:                 toUserResponse(res.User),
		})
	}
}

// LoginMfaHandler handles POST /auth/login/mfa: completes a login after the
// password step, gated by the mfa_pending_token + one second factor.
// Boundary check enforces exactly-one-code before the service (the service
// defends too). Success sets the refresh cookie — this is where the real
// session starts (R2's deferred issuance lands here).
func LoginMfaHandler(svc loginSessionService, cookieSecure bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req loginMfaRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			write400InvalidJSON(w)
			return
		}
		if (req.TotpCode == "") == (req.BackupCode == "") {
			// Neither present, or both present — exactly one is required
			// (openapi LoginMfaRequest description).
			WriteValidationError(w, []fieldError{
				{Field: "totp_code", Message: "exactly one of totp_code or backup_code must be present"},
				{Field: "backup_code", Message: "exactly one of totp_code or backup_code must be present"},
			})
			return
		}

		res, err := svc.LoginMfa(r.Context(), req.MFAPendingToken, req.TotpCode, req.BackupCode)
		if err != nil {
			MapServiceError(w, err)
			return
		}

		writeRefreshCookie(w, cookieSecure, res.RefreshTokenPlain)
		writeJSON(w, http.StatusOK, loginOKResponse{
			Status:               "ok",
			AccessToken:          res.AccessToken,
			AccessTokenExpiresAt: res.AccessTokenExpiresAt,
			User:                 toUserResponse(res.User),
		})
	}
}

// RefreshHandler handles POST /auth/refresh: reads the refresh token from
// the HttpOnly cookie (never body/bearer), rotates it, replaces the cookie.
// A missing cookie flows through as empty plain text — the service rejects
// it with the SAME generic 401 body as every other refresh-rejection class
// (missing/expired/revoked/reuse-detected are indistinguishable by contract).
func RefreshHandler(svc loginSessionService, cookieSecure bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := svc.Refresh(r.Context(), readRefreshCookie(r))
		if err != nil {
			MapServiceError(w, err)
			return
		}

		writeRefreshCookie(w, cookieSecure, res.RefreshTokenPlain)
		writeJSON(w, http.StatusOK, refreshResponse{
			AccessToken:          res.AccessToken,
			AccessTokenExpiresAt: res.AccessTokenExpiresAt,
		})
	}
}

// LogoutHandler handles POST /auth/logout idempotently: revokes the
// presented refresh token when present, ALWAYS clears the cookie, and
// answers 204 for every idempotent case — an absent or already-dead
// cookie is not an error condition (R16). The sole exception is a
// genuine infrastructure failure (e.g. the revoke UPDATE fails at the DB
// level): that surfaces as a 500 Problem Details response rather than a
// masked 204, so real outages are not hidden behind a success code.
func LogoutHandler(svc loginSessionService, cookieSecure bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := svc.Logout(r.Context(), readRefreshCookie(r))
		clearRefreshCookie(w, cookieSecure)
		if err != nil {
			MapServiceError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
