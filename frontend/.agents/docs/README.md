# `.agents/docs` — Index

> File: `frontend/.agents/docs/README.md`

One-off operational playbooks for tasks that aren't a "feature" — no
acceptance criteria, no `docs/spec/<domain>/features/` entry, but still
need a consistent, repeatable procedure. Read this index first, then open
only the matching file — don't scan the whole folder.

This is separate from `docs/spec/` (WHAT to build — invariants, threat
model, acceptance criteria, shared with the backend track) and
`docs/ui-ux/` (WHAT it looks like/behaves like — page map, patterns,
design tokens). Files here are HOW-to-execute playbooks for
setup/operational work that neither of those covers.

| Trigger | File | When to use |
|---|---|---|
| Empty `frontend/` (only default `create-next-app` scaffold — no `components/`, no `lib/api/`, no providers wired) | [`scaffold-frontend.md`](scaffold-frontend.md) | One-time, before Task #1 of any domain's frontend track — bootstraps folder structure, state/data providers, PWA manifest, and test infra |
| `campaign` domain's frontend track starting, `(public)/layout.tsx` still the Phase 0 pass-through stub | [`scaffold-public-shell.md`](scaffold-public-shell.md) | One-time — builds the real Public Shell nav (desktop top nav, mobile hamburger+drawer) that `/`, `/campaign`, `/campaign/[id]` all sit inside |

## Rules for playbooks in this folder

1. **Not a substitute for `docs/spec/` or `docs/ui-ux/`.** If a task has
   acceptance criteria, a page-map entry, or a visual spec, it belongs in
   one of those, not here — this folder is only for work with no such
   shape (infra bootstrap, tooling setup).
2. **One-time by nature.** Once a playbook here has been executed and its
   result committed, re-running it should be a no-op or an explicit,
   deliberate re-run — not part of every session's context. Do not add
   these files to `opencode.jsonc`'s `instructions` field; read them
   on-demand via the pointer in `frontend/AGENTS.md` instead.
3. **Flag gaps, don't invent.** Same rule as the rest of this repo (root
   `AGENTS.md` §1) — if a playbook is ambiguous, out of date with the
   actual repo state, or depends on a decision that hasn't actually been
   made yet (see `scaffold-frontend.md`'s Open Items), stop and say so
   rather than guessing.