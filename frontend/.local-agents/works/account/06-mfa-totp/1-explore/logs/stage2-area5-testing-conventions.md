# Stage 2 — Area 5: Testing conventions

## Current state

Read `unlink-google-form.test.tsx`, `login-methods-section.test.tsx`,
and `use-set-password.test.ts` in full — three distinct but consistent
layers of test convention already established by Task 05:

1. **Component tests** (`vitest` + `@testing-library/react` + MSW):
   - A local `withQueryClient(children)` helper wrapping the component
     under test in a fresh `QueryClient` (mutations/queries
     `retry: false`, so a forced-failure test doesn't hang on retries).
   - Default handlers come from the shared `mocks/handlers.ts`; any
     test needing a non-default branch calls `server.use(http.post(...))`
     inline, right in the `it(...)` block, returning the exact Problem/
     response shape being tested (not a shared alternate-branch
     fixture).
   - Assertions target rendered text/roles (`screen.getByRole`,
     `screen.findByText`, `toHaveTextContent`) — never snapshot testing.
   - Every `it(...)` title references the rule ID it proves (e.g. "(R16)",
     "(R18)") — traceability back to the techplan's requirement list.
   - `LoginMethodsSection`'s test file has its own `mockMe(authProviders,
     emailVerified)` helper that overrides `GET /account/me` per test —
     the parent-section test controls the identity shape driving which
     child branch renders, rather than each child test re-deriving it.
2. **Hook/pure-function tests**: `use-set-password.test.ts` tests
   `applySetPasswordSuccess` directly (not via `renderHook` or a
   mounted component) — a plain function call against a real
   `QueryClient` instance and the real `useAuthStore`, asserting cache/
   store side effects. Only possible because `useSetPassword`'s
   branching logic was extracted into a standalone, exported function
   in the first place (Area 3's finding) — this is the payoff for that
   extraction choice.
3. Zustand store reset via `beforeEach` (`useAuthStore.setState({
   accessToken: null })`) is a stated, explicit convention to prevent
   cross-test state leakage — noted directly in a comment, not just
   observed behavior.

## Requirement

`frontend/AGENTS.md` §4: component/unit tests via `vitest`; anything
rendering user-controlled Markdown/HTML needs a script-stripping test.

## Gap

- No test files exist yet for `MfaSection` or its children, or for the
  3 new hooks, or for the 3 new `lib/api/account.ts` wrapper functions
  at the unit level (existing precedent doesn't show dedicated
  `account.test.ts` cases per new endpoint beyond what's already there
  — `lib/api/account.test.ts` exists per the Stage 1 file listing but
  wasn't opened in this area; worth a quick check in Stage 3 for
  whether wrapper functions get their own unit tests or rely entirely
  on component-level MSW tests, since `unlink-google-form.test.tsx`
  above exercises `unlinkGoogle()` only indirectly through the
  component).
- Markdown/script-stripping requirement: **N/A for this feature** — no
  field in the MFA enroll/confirm/disable flow renders user-controlled
  Markdown or arbitrary HTML (the `otpauth_uri` is rendered as a QR
  code/plain text, backup codes are plain fixed-format strings) —
  explicitly noting this as N/A rather than silently skipping it, per
  `docs/spec/README.md`'s rule to record "N/A — reason" rather than
  leave a checklist item invisibly unaddressed.

## Page-consolidation check

N/A for this area.

## Sniffing

- **Edge case worth flagging for Stage 3**: testing the QR code itself
  is unusual territory for this codebase — depending on which library
  Stage 3 picks (Area 2's flagged gap), the rendered output might be an
  `<svg>`, a `<canvas>`, or a data-URI `<img>`. None of the existing
  test-assertion idioms here (`getByRole`, `getByText`) obviously cover
  "assert a QR code encoding the right `otpauth_uri` was rendered" —
  likely the more testable approach is asserting the library was called
  with the correct `otpauth_uri` value, not pixel-inspecting output, but
  that's a Stage 3 design decision, not something existing precedent
  answers.
- **Edge case**: testing that backup codes are shown "once" (per spec)
  is a temporal/state assertion, not a static-render one — e.g. a test
  confirming that after `enrollConfirm` succeeds and the section
  re-renders (say, after a follow-up interaction), the codes are no
  longer visible/retrievable in that same component tree. No existing
  precedent test does anything shaped like this (Task 05's flows are
  all form-resets-or-redirects, not "renders sensitive data exactly
  once then it's structurally gone"). Worth designing deliberately in
  Stage 3 rather than assuming a generic render test covers it.
- **Consistency**: aside from the two points above, nothing about this
  feature's testing needs contradicts the established conventions —
  MSW override-per-test, rule-ID-labeled test names, and the extracted-
  pure-function pattern (if `MfaDisable`'s branching ends up complex
  enough to warrant it, per Area 3's open question) all transfer
  directly.
</content>
