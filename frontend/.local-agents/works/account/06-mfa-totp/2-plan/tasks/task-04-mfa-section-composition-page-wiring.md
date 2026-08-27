# Task 4 — `MfaSection` composition & page wiring

> Derived from: `../techplan.md` ("Tech Plan: MFA TOTP (Frontend)",
> account/06-mfa-totp). This task file redistributes §8-13 detail
> relevant to its own scope, in full — it does not summarize. For the
> Summary, §1-7 rationale, and §14 Open Items, read the source techplan
> directly.
> Splitting axis: dependency/sequence chain + component boundary (see
> `manifest.md`).
> Dependencies: **Task 2** (Enroll flow UI) and **Task 3** (Disable flow
> UI) — this task imports `MfaEnrollFlow` and `MfaDisableForm` from both
> and composes them. Do not stub fake versions of either — this task's
> own tests exercise the real components + MSW mock path.
> Recommended model: **GLM 5.2 (max)** — per `best-practices/
> model-routing.md`'s Complex-tier "Coding/build" row and its
> tie-breaker ("GLM when the work leans on diagrams, state-transitions,
> or multi-step reasoning") — this task owns the single subtlest
> correctness risk in the whole plan (R12: the backup-codes-once view
> must survive an `account.me` cache refetch without being swapped away
> prematurely), explicitly flagged in the source techplan's §13 Common
> Mistakes as the easiest thing to get wrong across the entire feature.

## Scope

Build `MfaSection` (reads `useAccountMe()`, holds the "just enrolled,
codes not yet acknowledged" local state that overrides the `mfa_enabled`
branch, composes `MfaEnrollFlow`/`MfaDisableForm`) and wire it into
`app/(dashboard)/dashboard/security/page.tsx`.

**Rules owned by this task** (full text, copied from techplan §4):

- **R1** (loading): Given `useAccountMe()` hasn't resolved, When
  `MfaSection` renders, Then show a skeleton (same shape convention as
  `LoginMethodsSection`'s existing loading branch).
- **R2** (branch: not enrolled): Given `user.mfa_enabled === false` and
  no in-progress "just enrolled" local state, When `MfaSection` renders,
  Then render `MfaEnrollFlow` (Task 2), passing the `onEnrolled` callback.
- **R3** (branch: enrolled): Given `user.mfa_enabled === true` and no
  in-progress "just enrolled" local state, When `MfaSection` renders,
  Then render `MfaDisableForm` (Task 3), passing `hasEmailPassword`
  derived from `user.auth_providers`.
- **R12** (codes persist past refetch) — **the core risk this task
  owns**: Given confirm succeeded (Task 2's `onEnrolled(backup_codes)`
  fired) and `accountKeys.me()` has refetched (`mfa_enabled` now `true`),
  When `MfaSection` re-renders, Then it continues showing the
  backup-codes-once view — driven by the lifted `justEnrolledCodes`
  local state in `MfaSection`, **not** the `mfa_enabled` branch — until
  explicitly acknowledged. Getting this wrong means the codes view
  disappears within milliseconds of the cache invalidation Task 1's hook
  already performs — see Common Mistakes below.
- **R13** (codes acknowledged): Given the backup-codes-once view, When
  the user clicks "Saya sudah menyimpan kode ini", Then clear
  `justEnrolledCodes` and render `MfaDisableForm` (no additional API
  call — `mfa_enabled` is already `true` by this point, from R12's
  earlier invalidation).
- **R22** (a11y — codes are text): Given the backup-codes-once view,
  When rendered (by this component, not by Task 2's `MfaEnrollFlow` —
  see R9/Task 2's hand-off), Then the 10 codes are plain, selectable/
  copyable text (not an image/canvas snapshot) — usable via screen
  reader and copy-paste.

## Interface Contract (relevant subset of techplan §8)

**This task consumes from Task 2 and Task 3:**

```typescript
function MfaEnrollFlow(props: { onEnrolled: (backupCodes: string[]) => void }): JSX.Element; // Task 2
function MfaDisableForm(props: { hasEmailPassword: boolean }): JSX.Element; // Task 3
```

**This task's exports:**

```typescript
// components/features/account/mfa-section.tsx (new)
function MfaSection(): JSX.Element;
```

**Business logic flow (this task's slice, verbatim from §8 — the full
cross-component state machine):**

```
MfaSection:
  useAccountMe() not resolved           -> skeleton (R1)
  justEnrolledCodes !== null            -> backup-codes-once view (R12/R22),
                                            "Saya sudah menyimpan" button clears it (R13)
  justEnrolledCodes === null:
    mfa_enabled === false               -> <MfaEnrollFlow onEnrolled={setJustEnrolledCodes} /> (R2)
    mfa_enabled === true                -> <MfaDisableForm hasEmailPassword={...} /> (R3)
```

This is the same diagram as the source techplan's Summary decision-flow
diagram, at this task's own resolution — cross-check both if revising
either.

## Architecture (relevant note from §9)

`MfaSection` holds `justEnrolledCodes: string[] | null` as local
`useState`, lifted from `MfaEnrollFlow`'s `onEnrolled` callback (Task 2).
This state — not the `mfa_enabled` query value — is the actual
render-branch discriminant whenever it's non-null, which is what makes
R12 hold: `accountKeys.me()`'s refetch (triggered by Task 1's
`useMfaEnrollConfirm` hook) changes `mfa_enabled` to `true`
independently of `justEnrolledCodes`, and `MfaSection` must not let that
refetch short-circuit the codes view. Page wiring: replace the existing
placeholder comment in `app/(dashboard)/dashboard/security/page.tsx`
with `<MfaSection />`, below the already-shipped
`<LoginMethodsSection />` (Task 05's own component, untouched).

## Implementation Details (verbatim from §10)

**File**: `components/features/account/mfa-section.tsx` (new)
- Reads `useAccountMe()`. Holds `justEnrolledCodes: string[] | null`
  local state. Renders: skeleton (R1) → codes-once view (if
  `justEnrolledCodes` set, R12/R22, with the acknowledgment button, R13)
  → `MfaEnrollFlow`/`MfaDisableForm` (R2/R3) based on `mfa_enabled`.

**File**: `app/(dashboard)/dashboard/security/page.tsx`
- Replace the `{/* Account Task #6 (MFA) adds <MfaSection /> here */}`
  comment with `<MfaSection />`.

## Files Changed (this task's rows from §11)

| File | Change Type | Description |
|---|---|---|
| `components/features/account/mfa-section.tsx` | Add | Top-level composition + codes-once view + state-lifting logic (R12/R13) |
| `app/(dashboard)/dashboard/security/page.tsx` | Modify | Add `<MfaSection />`, remove placeholder comment |
| `components/features/account/mfa-section.test.tsx` | Add | Component tests |

**Reason untouched** (relevant row from §11): `components/features/account/login-methods-section.tsx` + children — `MfaSection` is added as an independent sibling only, per Task 05's own D1. `app/(dashboard)/_components/{dashboard-shell-client.tsx,nav-items.ts}` — "Keamanan" nav entry already exists from Task 05, no nav change needed.

## Testing Checklist (this task's items from §12, verbatim)

- [ ] R1: skeleton shown while `useAccountMe()` is loading
- [ ] R2: not-enrolled renders `MfaEnrollFlow`
- [ ] R3: enrolled renders `MfaDisableForm`
- [ ] R12: codes-once view persists across an `account.me` refetch
  (`mfa_enabled: true`) until acknowledged — **the critical regression
  test for this task**: simulate `MfaEnrollFlow`'s `onEnrolled` firing,
  then simulate/await the `account.me` query refetching with
  `mfa_enabled: true`, and assert the codes view is still shown, not
  `MfaDisableForm`
- [ ] R13: acknowledging codes transitions to `MfaDisableForm` with no
  extra API call
- [ ] R22: backup codes render as plain text, not an image

**Count-check** (this task's slice): 6 checklist items above, covering
R1, R2, R3, R12, R13, R22 — R2/R3's *branch-selection* behavior is
tested here; `MfaEnrollFlow`'s/`MfaDisableForm`'s own internal behavior
is Task 2's/Task 3's checklist responsibility, not re-tested here beyond
confirming the correct child renders.

## Testing Examples & Common Mistakes (relevant rows from §13)

| Mistake | Error/Behavior | Fix |
|---|---|---|
| `MfaSection` branching purely on `mfa_enabled` with no lifted "just enrolled" state | As soon as `useMfaEnrollConfirm`'s `onSuccess` (Task 1) invalidates `account.me` and it refetches `mfa_enabled: true`, the backup-codes-once view instantly disappears, replaced by `MfaDisableForm` — codes shown for a few hundred ms at most, effectively unusable | Hold `justEnrolledCodes` in `MfaSection` (R12) so the codes view overrides the `mfa_enabled` branch until explicitly acknowledged (R13) |
| Re-fetching or re-deriving enrollment state instead of trusting the lifted `onEnrolled` callback value | Risks a race between the callback firing and the cache refetch completing, or silently dropping the one-time `backup_codes` if not captured immediately | Capture `backup_codes` synchronously from the `onEnrolled(backupCodes)` callback argument into local state — never re-fetch it, it is never retrievable a second time (spec) |
</content>
