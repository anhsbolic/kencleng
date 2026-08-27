# Stage 2 — Area 5: `components/ui/` primitives inventory

## Current state

Read `button.tsx`, `input.tsx`, `label.tsx`, `banner.tsx` in full
(plus already knew `spinner.tsx`/`badge.tsx`/`progress-bar.tsx`'s
existence from the directory listing):

- **`Button`** already has a `destructive` variant
  (`bg-error-500 ... hover:bg-error-700`, matching
  `design-guidelines.md`'s Destructive token exactly) — the correct
  variant for "Lepas Tautan Google." Also has `loading` (disables +
  inline `Spinner`, `aria-busy`) and `sm`/`md`/`lg` sizes — everything
  Pattern 3's Idle/Submitting states need, already used identically by
  every existing form in this domain.
- **`Input`** already supports field-level `error` (wired via
  `aria-invalid`/`aria-describedby`, not just visual proximity) and
  passes through arbitrary native props (`type="password"`,
  `autoComplete`, etc.) — no new prop needed for this task's fields.
- **`Label`** is a bare styled wrapper, nothing feature-specific
  needed.
- **`Banner`** already covers all four semantic variants
  (success/error/warning/info) with the correct `role="alert"` vs
  `role="status"` split — sufficient for both the anti-enumeration
  `202` success banner and the two distinct `409` error banners this
  task needs.
- **No modal/dialog primitive exists** in `components/ui/` (confirmed
  by directory listing — only the five Phase-0 primitives plus
  `badge`/`progress-bar` added later). The Auth Shell's own modal
  (`(auth)/_components/auth-shell-client.tsx`) is a route-shell-level
  overlay, not a reusable `components/ui/` dialog component.

## Requirement

Per the feature spec, unlink requires only an inline `password` field
in the same request body — no separate confirmation dialog/step is
specified anywhere in `05-account-linking.md`. Pattern 3 (Form Page)
doesn't call for a modal either.

## Gap

**None found.** Every primitive this task's forms need
(destructive button, password-capable input with field errors, label,
four-variant banner, loading spinner via `Button`) already exists and
is already the established convention across every sibling account
form. No new `components/ui/` file is needed for this task — this
mirrors Task #4's equivalent finding (no primitive gap there either).

## Page-consolidation check

N/A — no primitive-level ownership question; nothing here overlaps
with Task #6 (MFA)'s eventual needs in a way that would require
coordinating who builds what (MFA's QR-code rendering, if it needs
anything beyond these primitives, is Task #6's own concern to surface
when it starts, per the Incremental Growth Rule in
`phase0-shared-infra.md`).

## Sniffing

- **Risk**: none specific to this area — the primitive set is already
  proven across four prior forms in this exact domain, so risk of a
  primitive-level surprise here is low.
- **Edge case**: none — nothing about this task's two endpoints needs
  behavior (file upload, multi-select, rich text) outside what these
  primitives already do.
- No miscontext, misleading signal, or inconsistency found in this
  area — it's the one area in this exploration that is a clean,
  boring "already sufficient" result.

Proceeding to Area 6 (visual precedent — `design-reference/`).
