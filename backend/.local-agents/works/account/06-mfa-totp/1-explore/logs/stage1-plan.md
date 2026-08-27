# Stage 1 — Plan Announcement

> Feature: 06-mfa-totp
> Date: 2026-08-27

## Areas to Explore

6 areas, in this order:

### Area 1: MFA Verifier Interface & Stub (`mfa_verifier.go`)
**Why first:** This is the Tier 0 fenced seam — the interface contract that task #6 must implement. Understanding its shape (method signatures, tx dependency, error semantics) determines what the rest of the code needs to look like. The stub fails closed today; the real implementation is the core of this task.

### Area 2: Database Layer — `mfa_totp_secrets` & `mfa_backup_codes` tables
**Why second:** Migrations already exist (pre-created by task #3). Need to confirm exact column shapes, indexes, and constraints. No new migrations needed — this area is about what repository methods must be added to `repository.go` / `repository_db.go` to support enroll/confirm/disable.

### Area 3: Service Layer — Enroll / Confirm / Disable business logic
**Why third:** The service is where the three endpoints' acceptance criteria get implemented. Depends on findings from Areas 1 & 2 (what the verifier can do, what the repo can do). Also need to understand the existing `LoginMfa` flow in `login.go` to ensure the verifier implementation is consistent with how it's consumed.

### Area 4: Transport Layer — HTTP handlers & routing
**Why fourth:** Handlers are thin wrappers over service calls. Depends on Area 3's service method signatures. Also need to check what's already wired in `account_security.go` / `main.go` and what needs adding.

### Area 5: Crypto Platform — TOTP secret encryption at rest
**Why fifth:** The feature spec requires `secret_encrypted` using the same AES-GCM pattern as other PII. Need to confirm `platform/crypto` can handle this without changes (it should — `Encrypt`/`Decrypt` are generic). This is also the Tier 0 boundary — TOTP generation/verification logic (`pquerna/otp` or equivalent) is the fenced sub-area.

### Area 6: Reauth Marker Infrastructure (`auth_google.go`)
**Why last:** The Google-only disable path consumes a reauth marker set by `GET /auth/google/redirect?intent=reauth`. This infrastructure already exists (built in task #02). Need to confirm the `CheckReauthMarker` function's contract and how task #6 should consume it — but this is a read-only dependency, not something task #6 builds.

## Order Rationale

Dependency-based. Area 1 (verifier interface) is the contract everything else plugs into. Areas 2-3 build outward from it. Area 4 is the thin transport wrapper. Areas 5-6 are platform dependencies consumed by the earlier areas. The Tier 0 sub-area (TOTP crypto) is intentionally explored late (Area 5) because its shape is determined by what Areas 1-3 need, not the other way around.
