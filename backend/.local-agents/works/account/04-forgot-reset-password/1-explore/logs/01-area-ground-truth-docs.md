# Area 1 — Domain ground-truth docs

> Stage 2 gap analysis. Files: `docs/spec/1-account/invariants.md`,
> `docs/spec/1-account/threat-model.md`, `docs/spec/1-account/tasks.md`.

## Current state

- **INV-account-05** (invariants.md:77–87): after reset, every
  `refresh_tokens` row for that `user_id` that was un-revoked at request
  start must be revoked by request end — *in the same transaction as the
  credential update*, explicitly "not a separate best-effort follow-up
  step". Verification named: test with 2+ active refresh tokens (two
  devices).
- **INV-account-08** (invariants.md:131–144): redemption requires
  `used_at IS NULL AND revoked_at IS NULL AND expires_at > now()` at the
  UPDATE; once redeemed, no further redemption. Verification named:
  concurrent double-submit of the same reset link, exactly one succeeds.
- **Threat model component 3** (threat-model.md:52–63): covers
  spoofing/tampering/enumeration/DoS for exactly these two endpoints;
  anti-enumeration marked "Resolved"; rate-limit on `/auth/*` cited as
  flood mitigation.
- **tasks.md**: task #4 = this feature, Tier 1, serial group S1, status
  **not started**. KPI table demands: named invariant test traceable to
  each referenced `INV-account-NN`, `-race` clean + ≥100-goroutine stress
  harness for race-sensitive invariants (08 listed), ≥80% coverage on new
  lines, security layer A/B. Tasks #1, #2 merged; #3 in progress at doc
  write time (see Inconsistency below).

## Requirement

Feature spec (`04-forgot-reset-password.md`) references these docs as
ground truth: INV-05 atomicity, INV-08 single-use, generic 202 across
three branches, validate-before-consume ordering (Assumption B), no
proactive revocation of prior tokens (Assumption A).

## Gap

No doc-level gap: the domain docs were written anticipating this exact
feature (INV-05/08 name Fitur 2B directly). This area's role is to supply
the correctness bar, and it does.

## Sniffing findings

1. **Risk** — INV-05 is the sharpest constraint: mass session-revoke must
   be transactional with the credential update. Two-step implementations
   fail the invariant even when behavior "looks right".
2. **Edge cases** — INV-08's verification demands a concurrency race
   test; tasks.md KPI escalates to ≥100 goroutines stress, stricter than
   the feature spec's plain double-submit criterion.
3. **Miscontext** — feature spec AC phrases the guarded update as
   `WHERE used_at IS NULL AND expires_at > now()`, omitting `revoked_at`;
   INV-account-08's canonical Statement includes it. An implementation
   matching only the AC text would technically violate INV-08.
4. **Misleading signals** — INV-08 mentions "any resend flow that sets
   `revoked_at` on a superseded token" as an operation it holds after;
   forgot-password deliberately has **no** such flow (Assumption A). Do
   not let INV-08's phrasing imply forgot-password must revoke outstanding
   tokens.
5. **Inconsistency** — minor: the `revoked_at` wording mismatch above
   (flagged, not silently resolved); plus trivial path drift — doc headers
   say `docs/spec/account/...` but the actual directory is
   `docs/spec/1-account/`.
6. **Process observation** — tasks.md places #4 strictly after #3 in
   serial group S1 (shared tables), yet #3's build status was unclear at
   explore time (see Area 3's inconsistency finding re tracker staleness).
