# Stage 3 — Solutioning

> Status: decisions D1–D7 resolved; open questions Q1–Q4 answered by
> Anhar 2026-08-26 (all recommendations accepted).
> Precedes the 2-plan phase (techplan / task breakdown).

## D1. Sequencing: MFA + user_roles dependencies — **Option C (approved)**

Schema-pre-settle + verifier seam:

- This task creates `login_attempts` **plus schema-only**
  `mfa_totp_secrets`, `mfa_backup_codes`, `user_roles` migrations (ERD DDL
  verbatim). Rationale: tasks.md's own Group-B phrasing ("once their
  independent tables' migrations are settled") anticipates tables being
  settled ahead of feature work. Empty-table semantics ≡ today's reality:
  nobody has MFA enrolled, nobody has roles — no behavior silently weakened;
  `enabled_at`/`roles` queries are real queries, not hardcoded `[]`/`false`.
- All TOTP/backup-code **verification** sits behind a `mfaVerifier` interface
  seam (house pattern: `breachChecker`, `googleOAuthClient`), backed by a
  stub that fails closed until #6 supplies the real implementation.
- Consequence to carry into planning: #6/#8 owners must be told their table
  migrations are taken (numbering coordinated); `/auth/login/mfa`
  success-path is fake-tested until #6 lands.

Rejected: A (full build now — scope explosion, ownership collision);
B (defer MFA endpoints — contradicts the spec folding MFA into this slice).

## D2. Tier 0 fenced sub-area authorship — **Option C (approved)**

Agent drafts full implementation + tests; the specific Tier 0 files are named
in the build report and go through mandatory human pair/rewrite before merge
(satisfies tasks.md KPI boolean "human-authored or human-rewritten, not
agent-generated wholesale").

Tier 0 file set (proposed):
- JWT primitives home: `platform/auth/token.go` extension via paired session —
  HS256 mint/verify for mfa_pending token; `purpose:"access"` private claim on
  ES256 access tokens; shared verifier/middleware checking purpose.
- Refresh rotation/reuse-detection repo/service core (wherever it lands, those
  files get flagged no-agent-write post-pairing).

Design content for the pairing session:
- Access token gains `purpose:"access"`; verifier rejects absent/wrong purpose.
  Sandbox-only exposure for pre-claim tokens; Google's verifier stays lenient.
- `mfa_pending_token`: HS256 `{sub, purpose:"mfa_pending", exp}`, 5-min TTL,
  key from `MFA_PENDING_TOKEN_SECRET` (32-byte base64, startup-validated like
  `crypto.New`). Key separation makes wrong-purpose tokens fail signature
  verification outright; `purpose` claim kept as belt-and-suspenders.

## D3. Rotation / reuse-detection mechanics

- **Rotate:** single tx = guarded parent UPDATE (`SET replaced_by_id = :child
  WHERE token_hash = :h AND replaced_by_id IS NULL AND revoked_at IS NULL AND
  expires_at > now() RETURNING user_id, family_id`) + child INSERT — one tx so
  a child-insert failure rolls back the parent mark (otherwise a transient
  error bricks the family via reuse detection).
- **Reuse/race-loser:** guard lost or rotated-token presented ⇒ revoke whole
  family (`SET revoked_at WHERE family_id = :fid AND revoked_at IS NULL`) ⇒
  401. Per spec Assumption D, race-loser treated identically to attack — no
  disambiguation branch needed; expired/revoked/rotated collapse to one 401.
- **Logout:** `SET revoked_at WHERE token_hash AND revoked_at IS NULL`;
  handler always clears cookie, always 204.
- Mirrors `RedeemToken`'s guarded-UPDATE-RETURNING pattern; `-race` +
  ≥100-goroutine stress harness per tasks.md KPI.

## D4. Lockout (Fitur 2C)

- Check BEFORE credential verify; rejected attempts write nothing.
- Password stage: `COUNT(*) … identifier_hash=$1 AND stage='password' AND
  success=false AND attempted_at > now() - interval '15 minutes'` ≥ 5 → 429.
- MFA stage: same threshold/window keyed `(user_id, stage='mfa')`, checked
  before code verification.
- Attempt rows written post-verify: `success=true` even when MFA branch
  follows; `success=false` on bad credentials/code.
- Anti-timing: wrong-email path burns a dummy bcrypt compare (extends
  register's R7 discipline).

## D5. 429 body conflict — **spec rule wins; amendment requested (approved)**

Implement per feature spec (outranks openapi in AGENTS.md §1 order): detail
byte-identical to the 401 (`"Email atau password salah."`), only status
differs; type URI `…/errors/too-many-requests`. Because shared
`TooManyRequests` is `$ref`'d by register/resend/forgot-password too, the fix
is a **login-specific 429 response definition in openapi**, not editing the
shared component — to be proposed to Anhar as an explicit spec amendment,
never silently applied. Transport rate-limiter's English 429 stays separate
(abuse-throttle ≠ credential-lockout; distinct problem type is honest).

## D6. Transport assembly

- Routes on `authMux` (inherit limiter): `POST /auth/login`, `/auth/login/mfa`,
  `/auth/refresh`, `/auth/logout`.
- New helpers `writeRefreshCookie` / `clearRefreshCookie` (Strict+HttpOnly+
  Secure-cond, Path=/, 30 d). `/auth/login` sets refresh-only cookie — do NOT
  reuse `writeAuthCookies` (would over-deliver access cookie); access token
  goes in body per contract. OAuth callback path unchanged.
- Sentinels `ErrInvalidCredentials`(401), `ErrLockedOut`(429),
  `ErrMfaPendingInvalid`(401) + `MapServiceError` branches.
- Env: `MFA_PENDING_TOKEN_SECRET` added to `requireEnv` + startup validation.

## D7. Testing map

Spec-named tests map onto the seams: lockout tests need an injectable clock
(15-min window math anyway); reuse/concurrency via real Postgres
(`integration` build tag) + `-race`; middleware rejection tests via forged
cross-key/cross-purpose tokens; generic-message test asserting byte-equal
bodies across wrong-email/wrong-password/lockout-shape cases; INV-named tests
per tasks.md KPI convention (`TestInvariant_Account03_*`, `_Account04_*`,
`_Account06_*` — the latter fake-verifier-scoped until #6).

## Approved open questions (2026-08-26)

1. D1 Option C approved (schema-pre-settle of mfa/user_roles tables here).
2. D2 Option C approved (agent drafts; human pair/rewrites Tier 0 files).
3. D5 approved (implement spec-rule 429; log openapi amendment request).
4. Rate-limiter X-Forwarded-For flaw: **deferred** as a flagged follow-up
   (scope discipline), not absorbed into this task.

## Risk-note seeds (for the eventual PR, root AGENTS.md §5)

- Assumptions: empty-table semantics stand in for unenrolled-MFA/no-roles
  reality until #6/#8 ship; stub `mfaVerifier` fails closed.
- Edge case intentionally NOT handled: multi-tab refresh race server-side
  (spec Assumption D — frontend BroadcastChannel concern).
- Concurrency assumptions: guarded UPDATE guarantees exactly-one-winner;
  rotate+insert atomicity in one tx prevents family-bricking on insert error.
- Not tested until #6: real TOTP/backup-code verification success paths.
