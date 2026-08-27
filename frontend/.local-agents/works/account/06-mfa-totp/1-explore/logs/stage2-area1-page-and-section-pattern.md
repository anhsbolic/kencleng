# Stage 2 — Area 1: `/dashboard/security` page + Task 05's sibling-section pattern

## Current state

`app/(dashboard)/dashboard/security/page.tsx` (15 lines) renders only
`<LoginMethodsSection />`, with an explicit comment:

> "Account Task #6 (MFA) adds its own `<MfaSection />` here as a
> sibling ... techplan account/05-account-linking D1 — independent
> section components, not a monolithic form."

This confirms **D1 from Task 05's techplan**: the page composes
independent, self-contained section components rather than one
monolithic form. `MfaSection` is expected to be added as a sibling
`<section>` inside the same `flex flex-col gap-6` wrapper, not a new
route.

### The established section-component shape (from Task 05)

- `LoginMethodsSection` (`components/features/account/login-methods-
  section.tsx`): the top-level section component. Pattern: `useAccountMe()`
  read once at the top, a loading skeleton (`if (!user) return <skeleton>`,
  no `isError` branch — waits rather than treats "not yet loaded" as
  failure), then derives presentational flags (`hasEmailPassword`,
  `hasGoogle`, `verified`, `mode`, `canUnlink`, `blockedReason`) from
  `user.auth_providers`/`user.email_verified` **client-side**, mirroring
  the backend's own guard conditions to avoid a guaranteed-4xx round
  trip for the common blocked cases. Renders an `<section className="...
  rounded-lg border ... bg-white p-6">` wrapper with an `<h2>` title,
  then composes two independent child components inside.
- `SetPasswordForm` (mode-driven: `"add" | "change"`, prop from parent —
  branch selection is server-derived state passed down, not internal
  component state) and `GoogleIdentityControl` (three-way branch: no
  Google → `GoogleAuthButton intent="link"`; Google + blocked → info
  `Banner`; Google + unlinkable → `UnlinkGoogleForm`) are each their own
  file, each independently testable.
- Every form component follows the same internal shape: `react-hook-form`
  + `zodResolver`, a `bannerRef` + `useEffect` that moves focus into
  whichever banner (`success`/`error`) is currently shown, a
  `GENERIC_ERROR_MESSAGE` constant marked `// TBD` (placeholder-pending-
  copy convention used everywhere in this codebase), and `ApiError.detail`
  shown verbatim for documented 401/409 cases (never for network/5xx).

### Re-auth precedent for MFA disable

`UnlinkGoogleForm` is the closest existing precedent for a re-auth-gated
destructive action: single `password` field, `variant="destructive"`
submit button, shows the backend's own `.detail` verbatim on 401/409.
MFA disable's `email_password` branch (password confirm) can follow this
shape almost exactly.

For the **Google-only branch** of MFA disable, the spec requires
`GET /auth/google/redirect?intent=reauth` first. Crucially,
`GoogleAuthButton`'s own doc comment (written during Task #2, before
this task existed) **already anticipates this exact use case**:

> "Always 'login' for this feature — 'link'/'reauth' belong to a
> different, session-authenticated flow (account linking / MFA
> re-auth, out of this component's scope), typed here only so the prop
> can't silently drift from the backend's own accepted set."

`GoogleAuthIntent` (derived from the generated `schema.d.ts`, not
hand-written) already includes `"reauth"` as a valid value — no new type
needed, `<GoogleAuthButton intent="reauth" label="..." />` is directly
usable.

### `mfa_enabled` already flows through

`schema.d.ts`'s `User` schema already has `mfa_enabled?: boolean`, and
`lib/api/account.ts`'s `getMe(): Promise<User>` returns it as-is —
`useAccountMe()` (used by `LoginMethodsSection` already) is sufficient
for `MfaSection` to know enrolled/not-enrolled state without a new
endpoint or hook. Same `accountKeys.me()` cache — enroll/confirm/disable
mutations should invalidate this key on success, same convention as
`useSetPassword`/`useUnlinkGoogle`.

## Requirement

Feature spec (`06-mfa-totp.md`) + `page-map.md`'s Donatur row: enable
MFA (QR scan + confirm code), view/regenerate backup codes, disable
MFA — all on `/dashboard/security`, as a section alongside login
methods. `patterns.md`'s Form pattern (multi-section) applies; no
dedicated Tier 1 prototype exists (confirmed by listing
`design-reference/` — 10 files, none security/MFA-shaped), so Tier 2
applies: derive from `patterns.md` + nearest Tier 1 precedent
(`/dashboard/campaign/new` or `/login`).

## Gap

- No `MfaSection` component exists yet — page has an explicit slot
  comment for it but nothing implements it.
- No QR code rendering precedent exists anywhere in this codebase yet
  (no QR library in `package.json` observed so far — to be confirmed
  in a later area/Stage 3, not re-derived from this area alone).
- No backup-codes display precedent exists (nothing analogous shown
  once/non-retrievable in the app yet — closest conceptually is
  nothing at all; this is a new UI shape for the codebase).
- **The Google-reauth round trip's redirect-back handling is an open
  gap, not just for this task but inherited from Task 05.** See
  Sniffing below — this is the biggest concrete finding of this area.

## Page-consolidation check

`/dashboard/security` already exists from Task 05 — confirmed via
direct file read, not inferred. This task is purely additive (one new
sibling component), not a new route. No `page-map.md` action here
lacks a backing endpoint — `docs/spec/account/features/06-mfa-totp.md`'s
3 endpoints (`enroll`, `enroll/confirm`, `disable`) map exactly onto
this row's "Enable/disable MFA... view/regenerate backup codes" text
(backup codes are returned by `enroll/confirm`, not a separate
endpoint — "regenerate" in `page-map.md`'s wording is actually
"disable then re-enroll," per the feature spec's explicit statement
that regeneration is "an already-resolved rule," not a distinct
endpoint). No orphaned endpoint or orphaned UI action found.

## Sniffing

- **Miscontext** (the significant one): `page-map.md` describes the
  action as "view/regenerate backup codes" as if regeneration were a
  standalone, repeatable action. The feature spec is explicit that
  there is no regenerate endpoint — the only way to get a fresh set of
  backup codes is disable → enroll → confirm. `MfaSection`'s copy/UX
  must not imply a one-click "regenerate" action exists; it needs to
  guide the user through disable-then-re-enable instead. Worth
  surfacing at Stage 3, not silently building a UI that promises a
  capability the backend doesn't have.
- **Misleading signal / real gap**: Task 05's "link" flow for Google
  never actually built any callback-redirect success/error handling —
  `GoogleIdentityControl` just re-derives its branch from
  `useAccountMe()` on remount, and there is no `useSearchParams`/query-
  param handling anywhere in `app/(dashboard)/dashboard/security/` or
  its components (confirmed via grep — zero matches). This "worked" for
  `link` because linking is idempotent/inspectable purely from account
  state after the fact. **MFA disable's `reauth` intent is different in
  kind**: the backend sets a short-lived (~5 min), single-use,
  session-tied server-side marker — there is no user-visible account
  state to re-derive from (the marker isn't part of `User`). The
  frontend has no existing mechanism to know "the reauth marker is now
  valid, the disable button should proceed" versus "the user just
  landed back on the page with no marker yet." This means either (a)
  the disable action must optimistically attempt the call and handle
  the `401` if the marker is missing, or (b) the backend's callback
  redirect needs to carry a query param this task must newly handle —
  neither exists as precedent, this is a genuine open design point for
  Stage 3, not a solved problem being reused.
- **Risk**: backup codes are shown exactly once, never retrievable
  again (spec, `INV-account-06`/07 area). If `MfaSection`'s enroll-
  confirm success view unmounts or the user navigates away before
  copying them, they're gone for good (by design) — this is a real UX
  risk (no "I saved them" confirmation gate observed as a requirement
  anywhere) worth flagging for Stage 3, not a defect in existing code.
- **Edge case**: the spec's enroll endpoint allows silently overwriting
  a pending unconfirmed secret if `enroll` is called again before
  confirming (abandoned-and-restarted case). If a user starts
  enrollment, leaves the page, and comes back, `MfaSection` needs to
  decide whether to auto-resume (call `enroll` again, get a fresh QR)
  or show some "you have an unfinished enrollment" state — no
  precedent for this either, since `mfa_enabled` (from `useAccountMe`)
  can't distinguish "never started" from "pending, unconfirmed" (both
  are `mfa_enabled: false`). Flagged for Stage 3.
- **Consistency**: the `GENERIC_ERROR_MESSAGE`/banner-focus/`ApiError.
  detail`-verbatim conventions are consistent across every existing
  section component (`SetPasswordForm`, `UnlinkGoogleForm`) — no
  contradiction found here; `MfaSection`'s children should follow the
  same shape for consistency, not invent a new error-handling idiom.
</content>
