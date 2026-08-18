# Playbook — Security Tooling Setup (Layer A)

> File: `backend/.agents/docs/security-tooling.md`
> Scope: one-time setup of config for the tools `backend/Makefile`'s
> `security:` target already calls (`gitleaks`, `govulncheck`) and the
> `lint:` target's `gosec` call. See
> `docs/kencleng-agentic-workflow.md` §8.1 for what each tool is
> responsible for and why.
> Status: starting draft — exact CLI flags/config schema below should be
> checked against each tool's own current documentation before relying on
> them, per root `AGENTS.md` §1 ("if none of the source-of-truth docs
> answer the question, stop and surface the gap explicitly"). Do not
> assume the flags below are still current without checking.

## Why this exists as a separate playbook

`docs/kencleng-roadmap-next-steps.md` deferred this explicitly: *"Full
configuration for security gate layer A per tool... deferred to when
each tool is actually invoked, not blocking the start of domain work."*
That moment is now — Task #1 (`account`) will hit `make verify`'s
`security:` target on its first gate run, and it should hit an actual
config, not tool defaults nobody decided on.

## `gosec` (part of `lint:` target)

Current `backend/Makefile` runs `gosec ./...` with no config. Before
Task #1 lands:

- Decide whether any rule needs a project-wide exclusion, and if so, add
  a config file and point `gosec ./...` at it (flag name — **TBD,
  verify against gosec's current CLI docs**, historically `-conf`).
- Cases that plausibly need an explicit, *documented* exclusion rather
  than a silent one:
  - G108 (pprof endpoints) — if a debug/pprof route is ever added, decide
    if it's dev-only and gate it, don't just suppress the finding.
  - Any SQL-related rule (G201/G202) — these should almost never be
    suppressed given the project's golden rule (`goqu` only, no raw SQL)
    — a finding here is more likely a real violation than a false
    positive.
- Prefer **inline** `#nosec` comments with a reason, over a broad
  project-wide rule exclusion — keeps the suppression visible next to
  the code it applies to, and each one is individually reviewable in a
  PR diff.
- Don't pre-emptively exclude rules "just in case" — start with gosec's
  defaults, add an exclusion only when a specific finding is confirmed
  to be a false positive for this codebase.

## `gitleaks` (`security:` target)

- Add `backend/.gitleaks.toml` (or repo-root, if a single config should
  cover both `backend/` and `frontend/` — confirm which `gitleaks
  detect` invocation this needs to match; the current `Makefile` runs it
  from `backend/`).
- Needed allowlist entry: `.env.example` — it contains variable names
  that look like secrets (`ENCRYPTION_KEY=`, `HMAC_KEY=`,
  `GOOGLE_CLIENT_SECRET=`) but with empty values. Confirm whether
  gitleaks' default rules already skip empty-value assignments, or
  whether `.env.example` needs an explicit path allowlist entry to avoid
  a permanent false-positive finding on every scan.
- Exact TOML syntax for the allowlist — **TBD, verify against gitleaks'
  current config docs** before writing the file.

## `govulncheck` ("baseline")

- `docs/kencleng-roadmap-next-steps.md` calls for a "baseline" here —
  worth confirming what that concretely means for `govulncheck`
  specifically. Unlike some scanners, govulncheck doesn't ship with a
  general-purpose baseline/suppression-file feature as far as documented
  usage goes — **TBD, verify against govulncheck's current
  documentation** whether such a mechanism exists before assuming one
  and building a workflow around it.
- If no such mechanism exists: the practical equivalent is running
  `govulncheck ./...` once the module has actual dependencies (after
  Task #1 pulls in real packages — an empty scaffold with no
  dependencies won't produce a meaningful baseline) and recording that
  first clean/non-clean output in a short note here, as the reference
  point for "was this finding already present before this PR."

## What this playbook does NOT cover

- Security Layer B (LLM adversarial review) — separate from this;
  covered by `@workspace` (harscode-workspace) guidelines plus
  `docs/kencleng-agentic-workflow.md` §8.2, not by these three CLI
  tools.
- `staticcheck` — not a security tool, already covered by the `lint:`
  target with no config needed beyond defaults.

## Before marking this playbook done

- [ ] Every "TBD — verify" item above has actually been checked against
      the tool's current docs, not left as an assumption.
- [ ] `make verify` (specifically the `lint:` and `security:` targets)
      runs clean against the empty scaffold from `scaffold-backend.md`,
      or fails with findings that are genuinely worth fixing before
      Task #1 — not tool-default noise nobody's looked at.