# `.agents/docs` — Index

> File: `backend/.agents/docs/README.md`

One-off operational playbooks for tasks that aren't a "feature" — no
acceptance criteria, no `docs/spec/<domain>/features/` entry, but still
need a consistent, repeatable procedure. Read this index first, then open
only the matching file — don't scan the whole folder.

This is separate from `docs/spec/`, which is the contract for WHAT to
build (invariants, threat model, acceptance criteria). Files here are
HOW-to-execute playbooks for setup/operational work.

| Trigger | File | When to use |
|---|---|---|
| Empty `backend/` (no `cmd/`, `internal/`, `migrations/` yet) | [`scaffold-backend.md`](scaffold-backend.md) | One-time, before Task #1 of any domain — bootstraps the Go project skeleton and `main.go` startup wiring |
| Setting up or revisiting `gosec`/`gitleaks`/`govulncheck` config | [`security-tooling.md`](security-tooling.md) | One-time initial setup, or when a Security Layer A gate (`make verify`) needs its config adjusted |

## Rules for playbooks in this folder

1. **Not a substitute for `docs/spec/`.** If a task has acceptance
   criteria, invariants, or a threat breakdown, it belongs in
   `docs/spec/<domain>/features/`, not here — this folder is only for
   work with no such shape (infra bootstrap, tooling setup).
2. **One-time by nature.** Once a playbook here has been executed and its
   result committed, re-running it should be a no-op or an explicit,
   deliberate re-run — not part of every session's context. Do not add
   these files to `opencode.jsonc`'s `instructions` field; read them
   on-demand via the pointer in `backend/AGENTS.md` instead.
3. **Flag gaps, don't invent.** Same rule as the rest of this repo (root
   `AGENTS.md` §1) — if a playbook is ambiguous or out of date with the
   actual repo state, stop and say so rather than guessing.