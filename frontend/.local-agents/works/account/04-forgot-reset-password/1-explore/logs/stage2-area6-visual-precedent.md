# Stage 2 — Area 6: Visual/prototype precedent

## Current state

- Extracted `docs/design-reference/login-register.html` per
  `design-reference-usage.md`'s Step 1 (bundler-template unpack, not
  raw `cat`/grep). Confirmed no forgot/reset-password prototype exists
  anywhere in this export — grepping the extracted JSX for
  `forgot|lupa|password` finds only:
  - Line 22: `<a href="#" ...>Lupa password?</a>` — a bare, dead link
    inside the login form's `PasswordField` component. No destination
    page, no separate `/forgot-password` or `/reset-password` mockup
    exists in this or any other `design-reference/` file (already
    confirmed by the Stage 1 directory listing — no
    `forgot-password*`/`reset-password*` file present).
  - Line 50: `error={err ? "Email atau password salah." : undefined}`
    attached directly to the login form's email `<Input>` — this is
    the literal source of the confirmed-not-fixed-in-export **Known
    Issue #1** from `prototype-reference.md` (generic auth failure
    rendered as a field-level error instead of a banner). Directly
    observed in the actual export file, not just taken on the doc's
    word.
- The real `login-form.tsx` (Area 4) has already structurally fixed
  this exact issue (banner-first-child, never attached to the email
  input) — confirms the fix described in `AuthShellClient`'s own JSDoc
  (Area 1) was actually carried through into the real component, not
  just documented as an intent.
- No new visual tokens, spacing values, or component compositions
  found in the extracted JSX beyond what `design-guidelines.md` and
  the already-implemented `login-form.tsx`/`register-form.tsx` already
  encode (same `PasswordField` show/hide-toggle composition already
  mirrored in the real `LoginForm`).

## Requirement

- `prototype-reference.md`: Tier 2, closest precedent `/login` — take
  it as structural/token precedent, not literal, per its own framing.
  Also: don't carry over Known Issue #1.

## Gap

- No prototype exists for either page's actual layout — Stage 3 must
  derive the form's visual shape entirely from `patterns.md` Pattern 3
  + `design-guidelines.md` tokens + the already-implemented sibling
  Auth Shell forms (Areas 1 and 4), since there's nothing more specific
  to extract from `design-reference/` for this feature. This isn't a
  missing deliverable — it's the expected Tier 2 situation the docs
  already describe, confirmed by direct inspection rather than assumed.

## Sniffing

- **Misleading signal**: the "Lupa password?" link inside
  `login-register.html` could look like there's a mocked destination
  page nearby in the same export — there isn't; it's a dead `href="#"`
  anchor with no corresponding page markup anywhere in the file.
- **Miscontext**: none — `prototype-reference.md`'s own Tier 2 table
  already correctly predicts this outcome; the doc did not overclaim
  coverage that turned out to be wrong.
- **Risk**: none new — this area is purely confirmatory.
- **Edge case**: n/a — not applicable to a visual-precedent check.
- **Inconsistency**: none — the real `login-form.tsx` and
  `AuthShellClient`'s documented convention already agree with each
  other and with `prototype-reference.md`'s Known Issue #1 write-up;
  all three point the same direction.

## Cross-references note

Extraction was done to a scratchpad temp file per
`design-reference-usage.md` and deleted after reading (not committed —
regenerable from the frozen source export at any time), consistent
with that doc's own "don't commit extracted files" guidance.

---

All six areas of Stage 2 complete. Awaiting confirmation before Stage 3
(Solutioning).
