# ⚠️ This is NOT the real Kencleng frontend

This directory contains **Claude Design-exported prototype code** —
a frozen, point-in-time visual/structural reference, generated to
validate `docs/ui-ux/design-guidelines.md` and `docs/ui-ux/
patterns.md` before real implementation started.

**It is not wired to any real backend, has no real state management,
no real API integration, and does not follow the architecture
decisions in `docs/project/kencleng-frontend-tech-stack.md`** (no
`components/features/` vs `components/ui/` vs `components/shared/`
split, no TanStack Query, no Zustand, no `openapi-typescript`
generated types). The real frontend lives in `frontend/` at repo
root — that is the actual application.

## How this should be used

- **As a human or agent building `frontend/`**: use this as a visual/
  structural precedent for what a page should look like — spacing,
  component composition, responsive behavior. Do not copy code from
  here wholesale into `frontend/`.
- **See `docs/ui-ux/prototype-reference.md`** for which routes have a
  prototype here (and which don't — most routes don't, by design; see
  that doc for the closest precedent to use instead), plus known
  issues in this export that should NOT be carried into the real
  implementation.
- **Do not modify files in this directory** as part of implementation
  work — it's frozen reference output, not a working codebase. See
  `AGENTS.md` §3.