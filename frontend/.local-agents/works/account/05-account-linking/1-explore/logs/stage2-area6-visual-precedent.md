# Stage 2 — Area 6: Visual precedent (`design-reference/`)

> Extracted per `docs/ui-ux/design-reference-usage.md`'s documented
> method (not raw `cat`/grep — the exports are Claude Design's JSX-in-
> HTML format). Both files are Tier 2 precedent only (per
> `prototype-reference.md`), informing visual polish, not behavior/data
> shape — consistent with why this area was explored last.

## Current state

**`campaign-new.html`** (closest Tier 1 precedent for a dashboard
multi-section form):
- Structured as a two-column grid (large form card left, two small
  independent cards right — Status + Tips), **not** stacked full-width
  sections with their own headers/dividers. All fields inside the left
  card are simply stacked (`gap: 20`) with no section headers or
  dividers between field groups — the only divider in the file is a
  dashed rule directly above the Cancel/Submit button row.
- Idle/error/submitting states are all driven by one `state` prop on a
  *fixed* field set: error state adds an icon + `error-600` caption per
  field (and border/background/focus-ring swap on `CurrencyInput`);
  submitting state fades the whole card to `opacity: 0.7`, disables
  every field, and swaps the submit button to a spinner + "Menyimpan
  draf" label.
- Primary/secondary button pairing: `outline` Cancel next to a filled
  primary Submit — the only button-variant contrast this file
  demonstrates.

**`login-register.html`** (password field precedent):
- `PasswordField`'s show/hide toggle is an `IconButton` (`variant=
  "brand"`, `size="sm"`) placed inside the same bordered box as the
  input, flush right — matches `LoginForm`'s already-implemented
  real-code pattern (Area 4) closely enough that no new visual idea is
  needed here beyond what's already built.
- The component declares an `error` prop but never renders it — no
  password-field error-state styling exists in this reference at all
  (only the sibling email `Input` demonstrates the shared error
  pattern).
- Primary action is a full-width filled `Button`; secondary
  (Google) is full-width `outline` with an icon, separated by a
  `Divider` + centered "atau" caption — matches `LoginForm`'s already-
  built divider treatment exactly.
- **No destructive button appears anywhere in either file.** A
  `--action-destructive` / `--action-destructive-hover` CSS token pair
  is defined in the shared token block used by both exports, but no
  `Button`/`IconButton` instance anywhere actually uses it — it's a
  defined-but-unused token, not a worked example.

## Requirement

Per `prototype-reference.md`, Tier 2 pages derive from `patterns.md` +
the closest Tier 1 precedent, not a literal copy. This task's page
needs: a destructive "Lepas Tautan Google" button, a re-auth password
field with error state, and a conditional two-branch form (Branch 1
vs. Branch 2 field sets) — none of which the prototype layer actually
demonstrates.

## Gap

- **No destructive-button visual precedent exists anywhere in
  `design-reference/`** — the real codebase's own `Button` component
  (Area 5) already has a working `destructive` variant matching
  `design-guidelines.md`'s tokens exactly, so this isn't a blocking
  gap, just confirmation that Stage 3 can't lean on a prototype example
  here and should treat `design-guidelines.md`'s token table as the
  sole source of truth for this specific element (consistent with
  `prototype-reference.md`'s own instruction to prefer the spec docs
  wherever the prototype is silent, not just wherever it's wrong).
- **No conditional/branching form precedent exists** — `campaign-new
  .html`'s only field-set variation is idle vs. error vs. submitting on
  one *fixed* set of fields, never an alternate field set switched on
  by state (which is exactly what Branch 1 vs. Branch 2 requires). The
  real codebase's own `LoginForm` (Area 4) already demonstrates this
  exact shape (password step vs. MFA step) — that's the precedent to
  actually build from, not the prototype.
- **No multi-independent-section-on-one-page precedent exists** —
  `campaign-new.html`'s two-column card+sidebar layout is a different
  shape from a Settings-style page with several stacked, independently
  titled sections (which `/dashboard/security` will need once both
  this task's section(s) and Task #6's MFA section eventually coexist,
  per Area 3's page-consolidation finding). This reference doesn't
  resolve that layout question either way.

## Page-consolidation check

No new finding — visual precedent has no bearing on the cross-task
(#5/#6) ownership question raised in Area 3; noting only that the
prototype layer offers no layout precedent for the "multiple stacked
sections" shape that question will eventually need, so Stage 3 can't
lean on it there either.

## Sniffing

- **Misleading signal**: the defined-but-unused `--action-destructive`
  CSS variable in the shared prototype token block could look like "a
  destructive button was designed and validated here" on a shallow
  grep of the token block alone — it wasn't; no component instance
  uses it anywhere in either file.
- **Risk**: none elevated by this area specifically — since the real
  codebase's own `Button`/`LoginForm` precedents (Areas 4-5) already
  cover everything this task visually needs, the prototype layer's
  gaps here don't block Stage 3, they just mean Stage 3 shouldn't
  spend time searching `design-reference/` further for these three
  patterns.
- **Inconsistency**: none found between this area and `design-
  guidelines.md`/the real component library — the prototype simply
  doesn't cover these cases, which is expected and already anticipated
  by `prototype-reference.md`'s own Tier 2 framing ("derive from
  patterns.md," not "the prototype has everything").

---

Stage 2 complete — all six areas explored. Summary of cross-area
findings ready for review before Stage 3 solutioning begins.
