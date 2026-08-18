# AGENTS.md — Kencleng

This file is the entry point for any AI coding agent working on this
repository. Read this **before** reading any spec file or writing any
code. It applies to both `backend/` and `frontend/`.

Kencleng is a sandbox donation-platform project. Its purpose is
learning — Go, PWA development, concurrency-safe backend patterns, and
secure-by-design coding — through a realistic domain. Code quality
(correctness under concurrency, security) is a hard requirement, not a
nice-to-have, even though most code here is agent-generated.

## 1. Source of truth, in order

1. `docs/spec/<domain>/invariants.md` and
   `docs/spec/<domain>/threat-model.md` — domain-level ground truth.
2. `docs/spec/<domain>/features/*.md` — feature-level acceptance
   criteria, which reference (not redefine) the domain docs above.
3. `api/openapi.yaml` — the API contract. Handler behavior must match
   it; it is hand-authored, not generated from code.
4. `docs/project/*.md` — narrative background docs (ERD, tech stack,
   phase details). Useful for context, but if a `docs/spec/*` document
   contradicts one of these for the same domain/feature, `docs/spec/*`
   wins — and the contradiction itself must be flagged, not silently
   resolved by picking one side.

If none of the above answer the question you're facing: **stop and
surface the gap explicitly** (as an assumption note in the feature
spec's "Assumption / open questions" section, or as a question to the
human). Do not silently pick the interpretation that seems most
reasonable to you and proceed. A wrong but flagged assumption is
recoverable; a wrong and silent one usually isn't caught until much
later.

## 2. Golden rules (non-negotiable)

These apply everywhere in the codebase, regardless of which feature
you're working on.

- **Errors are always returned, never swallowed.** No empty `catch`,
  no ignored `err`. `panic` is reserved for truly unrecoverable
  programmer errors, never for expected error paths (validation
  failures, not-found, etc.).
- **Money is always `decimal` (`shopspring/decimal`), never
  `float64`.** No exceptions, including in test fixtures.
- **SQL is always parameterized via `goqu`.** Never build a query
  string with `fmt.Sprintf` or string concatenation, even for
  read-only queries.
- **No secrets, PII, or tokens in logs.** Log the fact that an
  operation happened and its outcome, not the payload.
- **Error responses returned to clients never leak internals** — no
  stack traces, no raw SQL errors, no internal file paths. Wrap
  everything into the Problem Details format defined in
  `api/openapi.yaml`.
- **PII fields follow the established encryption pattern**
  (`{field}` ciphertext + `{field}_hash` HMAC) exactly as specified in
  `docs/project/kencleng-backend-tech-stack.md` — do not invent a
  different pattern for a new PII field.
- **Every authorization check is explicit at the handler/service
  boundary.** Don't rely on a query filter alone to enforce ownership
  — write the check so it's visible and testable on its own.

## 3. File-path fencing (Tier 0 — no agent write)

The following paths are human-authored or human-paired only. An agent
may **read** these for context but must not modify them without a
human explicitly asking for it in that specific session:

- `backend/internal/domain/donation/ledger.go` and any file
  implementing transaction/locking logic for balance updates
- `backend/internal/domain/disbursement/` state machine implementation
- `backend/internal/platform/crypto/` (encryption, HMAC, key handling)
- `backend/internal/platform/auth/` (JWT signing, TOTP, session logic)

If a task seems to require touching one of these, stop and say so
instead of proceeding — this is a signal that the task should be
re-scoped as Tier 0 work, not routed around.

## 4. Spec vs test vs code — authority separation

An agent implementing a feature **must not**:

- Edit a `docs/spec/*.md` file to make it match code that was already
  written, or to remove/loosen an acceptance criterion the code
  doesn't satisfy.
- Edit an already-approved test to make it pass (loosen an assertion,
  delete a failing case, lower a threshold) instead of fixing the
  code.

If you believe a spec or test is genuinely wrong, say so explicitly
and explain why — don't just change it and move on. Changing spec/test
files is a decision that goes through a human, not something folded
into an implementation PR.

## 5. Required per pull request

Every PR that introduces or changes behavior must include:

1. **Passing `make verify`** (see `kencleng-repo-setup.md` §5.5 for
   what this runs — lint, unit, race, contract, security-layer-A,
   integration, in that order).
2. **A risk note**, following this exact structure (see
   `kencleng-agentic-workflow.md` §9 for the reasoning behind this):

   ```markdown
   ## Risk note

   - Assumptions made: ...
   - Edge cases intentionally NOT handled (and why): ...
   - Concurrency assumptions: ...
   - What is not tested, and why: ...
   ```

   Every claim of "this edge case is handled" in this note must name
   the specific test that proves it. A claim with no named test is
   treated as unverified.
3. **For Tier 1 features**: a property-based or invariant test that
   exercises the relevant `docs/spec/<domain>/invariants.md` entry,
   not just a happy-path unit test.
4. **A commit/PR description that references which feature spec file
   this fulfills** (`docs/spec/<domain>/features/<fitur>.md`).

## 6. Reporting framing

Reporting a risk, limitation, or uncertainty in your risk note counts
as success. Finding out later that a risk existed but wasn't reported
counts as failure — regardless of whether the code otherwise works.
When in doubt about whether something is a risk worth flagging, flag
it.

## 7. Scope discipline

One vertical slice / endpoint per session. If a change starts pulling
in files unrelated to the stated task, stop and flag it rather than
continuing — a small stated task touching many unrelated files is a
signal something went wrong, not a sign of thoroughness.

**Directory boundary.** A session started from `backend/` must not
modify anything under `frontend/`, and vice versa — each side has its
own `opencode.jsonc` with `references` into `docs/` and `api/` for
shared contract material, but that's read access for context, not
license to edit across the boundary. If a task genuinely needs changes
on both sides (e.g. an API contract change), treat it as two separate
sessions coordinated by a human, not one session crossing the
boundary.

**One-off playbook sessions** (`.agents/docs/*.md` — scaffold,
tooling setup) are scoped to exactly what the playbook says, nothing
more. In particular, a scaffold session must not touch `docs/spec/`
even though it can read it — scaffolding is infrastructure wiring, not
a feature, and has no business editing the feature contract.

## 8. Related docs

- `docs/kencleng-agentic-workflow.md` — full process this file is
  extracted from (tiering rationale, testing stages, human checkpoint
  scope).
- `docs/kencleng-repo-setup.md` — repo structure, local dev setup.
- `docs/spec/README.md` — how to write/read domain invariants, threat
  models, and feature specs.