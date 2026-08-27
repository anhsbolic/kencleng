# Task 2 — MFA enroll flow UI (QR + confirm + manual-entry fallback)

> Derived from: `../techplan.md` ("Tech Plan: MFA TOTP (Frontend)",
> account/06-mfa-totp). This task file redistributes §8-13 detail
> relevant to its own scope, in full — it does not summarize. For the
> Summary, §1-7 rationale, and §14 Open Items, read the source techplan
> directly.
> Splitting axis: dependency/sequence chain + component boundary (see
> `manifest.md`).
> Dependencies: **Task 1** (API layer, hooks, mocks & manual-entry
> parsing utility) — this task imports and calls `useMfaEnroll()`,
> `useMfaEnrollConfirm()`, and `parseOtpauthSecret()`, all of which must
> exist first. Do not start this task by stubbing fake hooks "to
> unblock" — this task's own tests exercise the real hook + MSW mock
> path (`component-test-mocking-discipline.md`'s network-layer-mocking
> principle — see Task 1's mocks).
> Parallel-eligible with: **Task 3** (Disable flow UI) — no shared
> files, no shared state, both consume Task 1 independently.
> Feeds into: **Task 4** (`MfaSection` composition + page wiring)
> imports `MfaEnrollFlow` from this task and lifts its `onEnrolled`
> callback.
> Recommended model: **GLM 5.2 (max)** — per `best-practices/
> model-routing.md`'s Complex-tier "Coding/build" row and its
> tie-breaker ("GLM when the work leans on diagrams, state-transitions,
> or multi-step reasoning") — this component owns a genuine 3-state
> machine (`idle` → `confirming` → `done`) with a subtle non-obvious
> requirement (must NOT remount/re-fetch on a `422`, per R10) and is the
> first QR-rendering component anywhere in this codebase — zero existing
> visual or code precedent to copy, unlike Task 3's near-direct mirror of
> `UnlinkGoogleForm`.

## Scope

Build `MfaEnrollFlow` (the not-enrolled-state UI: "Aktifkan MFA" trigger
→ QR code + manual-entry secret + TOTP confirm form → hand-off to the
parent for the one-time backup-codes view), its `QrCode` wrapper
component, its `zod` confirm-code schema, and add the `qrcode.react`
dependency.

**Rules owned by this task** (full text, copied from techplan §4 — this
task owns the *component-layer* behavior for every rule below, including
R6/R7/R9/R10/R11/R24 whose *wrapper-function/parsing-utility-layer*
behavior was already built in Task 1):

- **R4** (no auto-fire): Given the not-enrolled state, When
  `MfaEnrollFlow` first mounts, Then `useMfaEnroll` does NOT fire
  automatically — only on an explicit "Aktifkan MFA" click. No
  precedent anywhere in this codebase fires a mutation as a render side
  effect; do not introduce the first one here.
- **R5** (enroll success): Given the user clicks "Aktifkan MFA", When
  `POST /account/security/mfa/enroll` resolves `200` (via Task 1's
  `useMfaEnroll()`), Then render `QrCode` (`value={otpauth_uri}`) plus
  the manual-entry secret line (R24) plus a `totp_code` input and a
  "Konfirmasi" submit button. Local step state transitions from `"idle"`
  to `"confirming"`.
- **R6** (enroll `409`, defensive) — **component-layer portion**: Given
  `useMfaEnroll()`'s mutation rejects with `ApiError(409, ...)` (Task 1's
  wrapper-layer throw), When received, Then show a generic error banner
  (`.detail` if present, else a frontend-owned fallback) — expected
  unreachable in normal single-tab use since this trigger only renders
  in the not-enrolled branch (R2, owned by Task 4).
- **R7** (enroll network/5xx) — **component-layer portion**: Given a
  network failure or unexpected `5xx` (Task 1's wrapper throws
  `ApiError`), When received, Then show a generic error banner;
  "Aktifkan MFA" re-enabled for retry (step state stays `"idle"`).
- **R8** (confirm submit): Given the QR+code view (`"confirming"`
  step), When the user submits `totp_code`, Then call Task 1's
  `useMfaEnrollConfirm()` mutation with `{ totp_code }`.
- **R9** (confirm success) — **component-layer portion**: Given confirm
  resolves `200` with `{ backup_codes }` (exactly 10 strings, Task 1's
  hook already invalidated `accountKeys.me()`), When received, Then call
  the `onEnrolled(backup_codes)` prop (lifted to `MfaSection`, Task 4)
  and transition local step state to `"done"`, unmounting this
  component's own QR/form UI — the parent (Task 4) takes over rendering
  the backup-codes-once view. This component does **not** render the
  codes itself.
- **R10** (confirm `422`): Given confirm rejects with `ApiError` (Task
  1's wrapper throws on `422`), When received, Then show one fixed
  frontend-owned message ("Kode tidak valid, coba lagi.") tied to the
  `totp_code` field via `aria-describedby`/`aria-invalid`; the QR and
  form **remain mounted/interactive** — no remount, no re-fetch of
  `enroll` (pending secret is not discarded, per the spec — this is the
  single easiest-to-get-wrong rule in this task, see Common Mistakes
  below).
- **R11** (confirm network/5xx): Given a network failure or unexpected
  `5xx` on confirm (Task 1's wrapper throws), When received, Then show a
  generic request-level error banner; form stays interactive
  (identical treatment to R10, minus the field-level tie-in).
- **R21** (a11y — banner focus, this component's own banners): Given any
  error banner in this component renders (R6/R7/R10/R11), When it
  renders, Then focus moves into it (`bannerRef` + `useEffect`, matching
  `UnlinkGoogleForm`/`SetPasswordForm`'s existing convention).
- **R24** (a11y — manual-entry secret fallback) — **component-layer
  portion**: Given the QR+confirm-code view renders (R5), When
  `parseOtpauthSecret()` (Task 1) returns a secret string, Then render it
  as selectable/copyable monospace text alongside the QR (e.g. "Tidak
  bisa scan? Masukkan kode ini secara manual: ..."); if it returns
  `null`, hide the manual-entry line entirely without breaking the QR
  view.

## Interface Contract (relevant subset of techplan §8)

**This task consumes from Task 1:**

```typescript
useMfaEnroll(): UseMutationResult<MfaEnrollResponse, ApiError, void>;
useMfaEnrollConfirm(): UseMutationResult<MfaEnrollConfirmResponse, ApiError, MfaEnrollConfirmRequest>;
parseOtpauthSecret(otpauthUri: string): string | null;
```

**This task's exports:**

```typescript
// components/features/account/qr-code.tsx (new)
function QrCode(props: { value: string }): JSX.Element;

// components/features/account/mfa-enroll-flow.tsx (new)
function MfaEnrollFlow(props: {
  onEnrolled: (backupCodes: string[]) => void; // R9 — lifted to MfaSection (Task 4)
}): JSX.Element;
```

**Business logic flow (this task's slice, verbatim from §8):**

```
MfaEnrollFlow (local step: "idle" | "confirming" | "done")
  "idle": "Aktifkan MFA" button
    click -> useMfaEnroll().mutate()
      -> 200  => step="confirming", render QrCode + manual-entry secret + totp_code form (R5/R24)
      -> 409/network/5xx => banner, step stays "idle" (R6/R7)
  "confirming": QrCode + manual-entry secret + totp_code form
    submit -> useMfaEnrollConfirm().mutate({ totp_code })
      -> 200  => onEnrolled(backup_codes), step="done", this component unmounts its own UI (R9)
      -> 422  => fixed field-level message, step stays "confirming", QR/form untouched (R10)
      -> network/5xx => generic banner, step stays "confirming" (R11)
```

## Architecture (relevant note from §9)

`qr-code.tsx` is a thin wrapper around `qrcode.react`'s `QRCodeSVG`,
`value`/`size` props only — SVG output needs no jsdom canvas polyfill
(D3). `mfa-enroll-flow.tsx` owns the local step state described above;
`"done"` hands off to the parent rather than rendering the codes view
itself (D1's "independent section components" philosophy, applied one
level down — this component's job ends at enrollment, not at showing
codes). `mfa-enroll-confirm-schema.ts` is a minimal `zod` schema
(required-field-only for `totp_code`, no invented length/format policy
the spec doesn't state — per `form-validation-boundary.md`'s "don't
invent client-side rules the backend doesn't authoritatively define").

## Implementation Details (verbatim from §10)

**File**: `components/features/account/qr-code.tsx` (new)
- `QrCode({ value }: { value: string })` — renders `<QRCodeSVG
  value={value} size={200} />` (size: `TBD — verify` final value against
  `design-guidelines.md` spacing once built; 200px is a placeholder
  starting point, not confirmed).

**File**: `components/features/account/mfa-enroll-flow.tsx` (new)
- Local step state: `"idle" | "confirming" | "done"`. `"idle"` →
  "Aktifkan MFA" button (R4). `"confirming"` → `QrCode` + the
  manual-entry secret line (R24, hidden if `parseOtpauthSecret` returns
  `null`) + `totp_code` form (R5/R10/R11). `"done"` → calls
  `onEnrolled(backup_codes)` prop (R9) and unmounts itself.

**File**: `components/features/account/mfa-enroll-confirm-schema.ts`
(new)
- `zod` schema for `totp_code`: required, non-empty (no digit-count/
  format regex — not stated anywhere in the feature spec, per
  `form-validation-boundary.md`'s checklist item on not inventing
  unstated client-side policy).

**File**: `package.json` / lockfile
- Add `qrcode.react` dependency (D3, confirmed by Anhar) — pin exact
  version.

## Files Changed (this task's rows from §11)

| File | Change Type | Description |
|---|---|---|
| `components/features/account/qr-code.tsx` | Add | Thin `qrcode.react` wrapper |
| `components/features/account/mfa-enroll-flow.tsx` | Add | Enroll → confirm → hand-off flow |
| `components/features/account/mfa-enroll-confirm-schema.ts` | Add | `zod` schema for `totp_code` |
| `package.json` / lockfile | Modify | Add `qrcode.react` dependency (no other task touches this dependency) |
| `components/features/account/mfa-enroll-flow.test.tsx` | Add | Component tests (also covers `QrCode` via a mocked `qrcode.react` module — see Common Mistakes) |

**Reason untouched** (relevant row from §11): `components/features/account/mfa-disable-form.tsx` and its schema — Task 3's independent scope, no shared file.

## Testing Checklist (this task's items from §12, verbatim)

- [ ] R4: `useMfaEnroll` does not fire on mount, only on click
- [ ] R5: enroll success renders QR + manual-entry line + `totp_code` form
- [ ] R6 (component-layer): a mocked `409` (via MSW, Task 1's endpoint)
  shows a generic error banner
- [ ] R7 (component-layer): a mocked network failure/`5xx` on enroll
  shows a generic error banner, button re-enabled
- [ ] R8: confirm submit calls the endpoint with `{ totp_code }`
- [ ] R9 (component-layer): confirm success calls `onEnrolled` with the
  10-item `backup_codes` array and unmounts this component's own QR/form
  UI
- [ ] R10: confirm `422` (via MSW) shows the fixed message, QR/form
  remain mounted and interactive (not remounted, no re-fetch of enroll)
- [ ] R11: confirm network/5xx (via MSW) shows a generic banner, form
  stays interactive
- [ ] R21: focus moves into any of this component's banners on render
- [ ] R24 (component-layer): manual-entry secret rendered as selectable
  text next to the QR when `parseOtpauthSecret()` returns a value;
  hidden entirely when it returns `null`

**Count-check** (this task's slice): 10 checklist items above, covering
R4, R5, R6 (component-layer), R7 (component-layer), R8, R9
(component-layer), R10, R11, R21, R24 (component-layer) — the
*wrapper-function/parsing-utility-layer* halves of R6/R7/R9/R10/R11/R24
live in Task 1's own checklist, not duplicated here.

## Testing Examples & Common Mistakes (relevant rows from §13)

| Mistake | Error/Behavior | Fix |
|---|---|---|
| Unmounting/remounting the QR view on a `422` | Loses the pending secret's on-screen QR, user has to re-scan even though the backend never discarded the secret (R10 explicitly forbids this) | Keep the `"confirming"` step state unchanged on `422` — only the inline field error updates |
| Testing the QR code by inspecting rendered SVG path data | Brittle, tests an implementation detail of `qrcode.react`, not this app's own logic | Mock the `qrcode.react` module (narrow, justified exception per `component-test-mocking-discipline.md` — a network-layer mock can't reach "what a third-party rendering library drew") and assert the wrapper received the correct `value` prop |
| Rendering the backup-codes-once view directly inside `MfaEnrollFlow`'s `"done"` step instead of handing off via `onEnrolled` | Duplicates the codes-view rendering logic that Task 4's `MfaSection` also needs to own (the persist-across-refetch requirement, R12, lives at the parent level) — two divergent implementations | `"done"` only calls `onEnrolled(backup_codes)` and unmounts; the parent (`MfaSection`, Task 4) renders the codes view |
</content>
