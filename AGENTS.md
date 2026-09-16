# AGENTS.md — Kencleng

Root instruction map and cross-repository hard rules for AI coding agents.

Read this file before writing code. More-specific `AGENTS.md` files apply within their directory scope and own stack-specific detail.

Kencleng is a sandbox donation/crowdfunding project for correctness-critical, secure-by-design agentic engineering. Agent-generated code does not reduce correctness, security, or evidence requirements.

## 1. Route to the authority that owns the concern

```text
domain invariants / threats
→ docs/spec/<domain-dir>/invariants.md
→ docs/spec/<domain-dir>/threat-model.md

feature behavior / acceptance criteria
→ docs/spec/<domain-dir>/features/*.md

API contract
→ api/README.md (routing/editing rules)
→ api/openapi/<domain>.yaml + referenced common.yaml components
→ api/openapi.yaml only when an aggregate/generated view is needed

backend architecture
→ docs/project/kencleng-backend-tech-stack.md

frontend architecture
→ docs/project/kencleng-frontend-tech-stack.md

frontend product/design/component detail
→ frontend/AGENTS.md

active UI/UX design authority map
→ docs/ui-ux/README.md

project status
→ docs/project/kencleng-development-tracker.md

Kencleng project orchestration
→ docs/kencleng-agentic-workflow.md

feature lifecycle / generic engineering practice
→ Harscode workflow / best-practices
```

Treat this map as routing, not as an instruction to read every target in full. Start from the concern that is active, locate the authoritative section/operation/heading, follow its referenced dependencies, and expand only when the task needs broader consistency context.

For API work, prefer the split source for the active domain plus only the referenced shared components from `api/openapi/common.yaml`. `api/openapi.yaml` is the generated bundled aggregate and remains useful for aggregate/cross-domain inspection and generated-client correspondence; do not load it in full by default for a domain-local task.

For business behavior, domain invariants/threat models and feature specs are authoritative over narrative project background; the OpenAPI source owns API shape. Do not apply that precedence to unrelated concerns owned by architecture or design documents.

For frontend design work, use `docs/ui-ux/README.md` as the routing entrypoint. The active approved direction is **Sunlit Editorial / Evidence-Led Optimism**. Removed legacy design guidelines and prototype exports are historical Git evidence, not current authority.

If authorities genuinely conflict on the same concern, surface the contradiction instead of choosing whichever interpretation makes implementation easiest.

## 2. Golden rules

- **Errors:** return/handle failures intentionally. Never silently swallow expected errors.
- **Money:** backend monetary arithmetic uses the established decimal representation; never introduce `float64` for money, including fixtures.
- **SQL:** use the established parameterized `goqu` approach; never construct SQL from user-controlled values via interpolation/concatenation.
- **Sensitive data:** never log secrets, raw tokens, or PII payloads.
- **Client errors:** never leak stack traces, raw SQL, internal filesystem paths, or other implementation internals. Follow the applicable Problem Details contract in the OpenAPI source.
- **PII:** follow the established encryption/HMAC storage pattern; do not invent a second convention.
- **Authorization:** backend authorization checks are explicit. Frontend role gates or hidden controls are UX, not security authority.

## 3. File-path fencing

The following paths are read-only to an agent unless a human explicitly authorizes work on them for that session:

- `backend/internal/domain/donation/ledger.go` and any file implementing transaction/locking logic for balance updates;
- the state-machine implementation under `backend/internal/domain/disbursement/`;
- `backend/internal/platform/crypto/` (encryption, HMAC, key handling);
- `backend/internal/platform/auth/` when changing security-critical JWT signing, TOTP, refresh-token/session, or equivalent authentication core logic.

These are Tier-0 / human-authored or human-paired areas. Agents may read them for context, critique, test ideas, or adversarial review but must not modify them without explicit authorization.

More-specific scoped `AGENTS.md` files may add protections; they must not silently weaken this root fencing.

If a task appears to require a protected write, surface the boundary and re-scope the work instead of routing around it.

## 4. Spec / test / code authority separation

An implementation agent must not change an established requirement merely to make its code pass.

Do not:

- weaken a feature acceptance criterion to match implementation;
- alter a domain invariant because code violates it;
- loosen an already-approved test solely to obtain green output;
- lower a verification threshold silently.

If a spec or test appears wrong, report the contradiction, explain the evidence, and route the correction through the appropriate human/project authority.

## 5. Workflow authority

Harscode owns the generic per-feature lifecycle:

```text
Exploration
→ Techplan
→ Build / Patch
→ Code Review
→ Testing
→ Pull Request
```

Harscode owns phase responsibilities, session/context boundaries, patch routing, reports, and generic engineering best-practices. Same-session vs fresh-session choices follow the active Harscode context/session guidance; the arrow above is lifecycle order, not a requirement that adjacent phases share a session.

When manually invoking a Harscode phase, start from that phase's current canonical `workflow/*-prompt.md` entrypoint. Fill its variables and add only narrow Kencleng/task context that is not already owned by the prompt. Do not maintain a second Kencleng-authored copy of generic phase instructions.

Use `docs/kencleng-agentic-workflow.md` for Kencleng-specific orchestration such as domain sequencing, risk tiers, human authority, backend/frontend coordination, design readiness, mock-first integration, and domain finalization. Read the relevant section(s); it is not mandatory startup prose for every phase.

Older feature/task docs may contain numeric section references to previous workflow revisions. Follow the current named rule/source owner rather than reviving superseded policy from an old `§NN` reference.

## 6. Scope and directory boundaries

Work on one coherent Harscode work unit at a time. Do not combine unrelated changes merely because they are nearby.

Backend and frontend production writes remain separate by default:

- backend-scoped Build → do not casually modify `frontend/`;
- frontend-scoped Build → do not casually modify `backend/`;
- shared `docs/` and `api/` may be read for context;
- cross-stack contract changes require explicit coordination.

A small task unexpectedly touching many unrelated files is a risk signal. Investigate before expanding scope.

Within stack scope:

- `backend/` → read `backend/AGENTS.md`; it owns backend-specific execution rules.
- `frontend/` → read `frontend/AGENTS.md`; it owns frontend architecture/design/component/rendered-verification routing.

Do not repeat those scoped rules here merely to make the root file self-contained.

## 7. Verification and evidence

Use the repository's actual executable commands/configuration. Never claim a check ran unless it ran.

Evidence is risk-driven. Applicable examples include invariant/property evidence, concurrency/race evidence, threat-focused security checks, hostile-content rendering coverage, human rendered acceptance for material UI, explicitly requested browser automation, and real backend/frontend integration before integrated completion.

Do not indiscriminately run expensive test classes during Build when the active Harscode workflow assigns them to independent Testing. Stack-specific verification boundaries live in the scoped `AGENTS.md` files.

Distinguish `verified`, `assumed`, `deferred`, and `not tested`. A known material risk silently omitted is worse than a clearly documented limitation.

## 8. Pull-request evidence

A behavior-changing PR must identify:

- **Scope:** feature/spec/task fulfilled.
- **Verification:** commands/checks actually run.
- **Risk note:**

```markdown
## Risk note
- Assumptions made: ...
- Edge cases intentionally NOT handled (and why): ...
- Concurrency assumptions: ...
- What is not tested, and why: ...
```

Tier-specific and stack-specific additions follow `docs/kencleng-agentic-workflow.md` and the applicable scoped `AGENTS.md`.

## 9. Human authority

Human approval is required where defined for:

- Tier-0 implementation;
- Tier-1 merge/review;
- product/domain contract changes;
- protected spec/test changes;
- brand-defining visual assets;
- material `OPEN` product/design decisions;
- manual DB/index application;
- any additional human gate explicitly owned by the applicable scoped/project authority.

An agent must not approve its own work where independent human authority is required.

## 10. One-off docs and generated artifacts

One-off setup/playbook documents, when present, are on-demand context only. They do not override canonical specs, architecture/design authorities, root/scoped `AGENTS.md`, or Harscode workflow.

Generated exploration, techplan, build, review, testing, and PR artifacts are task evidence/history, not automatically project-wide precedent. Promote reusable truth into the source that owns that concern.
