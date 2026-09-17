# AGENTS.md — Kencleng

Root instruction map and cross-repository hard rules for AI coding agents.

Read this file before writing code. More-specific `AGENTS.md` files apply within their directory scope and own stack-specific detail.

Kencleng is a sandbox donation/crowdfunding project for correctness-critical, secure-by-design agentic engineering. Agent-generated code does not reduce correctness, security, or evidence requirements.

## 1. Route to the authority that owns the concern

```text
whole-product purpose / actors / concepts / durable business truth
→ docs/product/README.md
→ docs/product/product-overview.md

current MVP inclusion / exclusion / release boundary
→ docs/product/mvp-scope.md

current MVP vertical delivery order / slice boundaries
→ docs/product/mvp-delivery-slices.md

product-design / UX / visual / asset expression
→ docs/ui-ux/README.md

reconciled domain/delivery invariants / threats
→ docs/spec/<domain-dir>/invariants.md
→ docs/spec/<domain-dir>/threat-model.md

reconciled feature behavior / acceptance criteria
→ docs/spec/<domain-dir>/features/*.md

API contract shape for the active reconciled slice
→ api/README.md
→ api/openapi/<domain>.yaml + referenced common.yaml components
→ api/openapi.yaml only when an aggregate/generated view is needed

backend architecture
→ docs/project/kencleng-backend-tech-stack.md

frontend architecture
→ docs/project/kencleng-frontend-tech-stack.md

frontend component/execution detail
→ frontend/AGENTS.md

project status
→ docs/project/kencleng-development-tracker.md

Kencleng project orchestration
→ docs/kencleng-agentic-workflow.md

feature lifecycle / generic engineering practice
→ Harscode workflow / best-practices
```

Treat this map as routing, not as an instruction to read every target in full. Start from the concern that is active, locate the authoritative section/operation/heading, follow its referenced dependencies, and expand only when the task needs broader consistency context.

### Product-first precedence

Product/MVP authority owns **what should be true and what is currently in scope**. Product Design / Brand Authority owns the concrete experience/design concerns assigned to it. Delivery specs and OpenAPI own lower-level detail **after that detail has been reconciled for the active slice**.

Use this direction:

```text
Product Authority + approved MVP scope/sequencing
        +
Product Design / Brand Authority when relevant
        ↓
active delivery slice
        ↓
domain/delivery specification
        ↓
shared API contract
        ↓
implementation
```

Existing detailed specs, OpenAPI, ERD/data-model material, migrations, tests, and code may contain valuable implementation evidence. They do **not** silently override Product/MVP authority merely because they are more detailed or already implemented.

If Product/MVP truth is clear and a historical delivery artifact conflicts, treat the lower-level artifact as needing reconciliation. If Product Authority is genuinely missing a material product decision, surface that gap rather than inventing the answer in a spec, contract, or implementation.

For API work, prefer the split source for the active slice/domain plus only the referenced shared components from `api/openapi/common.yaml`. `api/openapi.yaml` is the generated bundled aggregate and remains useful for aggregate/cross-domain inspection and generated-client correspondence; do not load it in full by default for a domain-local task.

For frontend design work, use `docs/ui-ux/README.md` as the routing entrypoint. The active approved direction is **Sunlit Editorial / Evidence-Led Optimism**. Removed legacy design guidelines and prototype exports are historical Git evidence, not current authority.

If authorities genuinely conflict on the same concern, surface the contradiction instead of choosing whichever interpretation makes implementation easiest.

## 2. Golden rules

- **Errors:** return/handle failures intentionally. Never silently swallow expected errors.
- **Money:** backend monetary arithmetic uses the established decimal representation; never introduce `float64` for money, including fixtures.
- **SQL:** use the established parameterized `goqu` approach; never construct SQL from user-controlled values via interpolation/concatenation.
- **Sensitive data:** never log secrets, raw tokens, or PII payloads.
- **Client errors:** never leak stack traces, raw SQL, internal filesystem paths, or other implementation internals. Follow the applicable Problem Details contract in the OpenAPI source.
- **PII:** follow the established encryption/HMAC storage pattern while it remains applicable; do not invent a second convention casually. If an active slice materially changes the data semantics, re-evaluate the pattern explicitly rather than weakening protection.
- **Authorization:** backend authorization checks are explicit. Frontend role gates or hidden controls are UX, not security authority.
- **Simulation:** sandbox behavior must never be presented as real external settlement, independent verification, or real-world evidence.

## 3. File-path fencing

The following paths are read-only to an agent unless a human explicitly authorizes work on them for that session:

- `backend/internal/domain/donation/ledger.go` and any file implementing transaction/locking logic for balance updates;
- the state-machine implementation under `backend/internal/domain/disbursement/`;
- `backend/internal/platform/crypto/` (encryption, HMAC, key handling);
- `backend/internal/platform/auth/` when changing security-critical JWT signing, TOTP, refresh-token/session, or equivalent authentication core logic.

These are Tier-0 / human-authored or human-paired areas. Agents may read them for context, critique, test ideas, or adversarial review but must not modify them without explicit authorization.

More-specific scoped `AGENTS.md` files may add protections; they must not silently weaken this root fencing.

If a task appears to require a protected write, surface the boundary and re-scope the work instead of routing around it.

## 4. Product / spec / contract / test / code separation

An implementation agent must not change an established higher-level requirement merely to make its code pass.

Do not:

- use a historical domain spec to override canonical Product/MVP truth;
- weaken an active reconciled feature acceptance criterion to match implementation;
- alter an applicable invariant because code violates it;
- loosen an already-approved test solely to obtain green output;
- lower a verification/security threshold silently;
- expand MVP scope because a historical feature is already implemented.

If a Product/MVP, design, spec, contract, test, or implementation artifact appears wrong, report the contradiction, explain the evidence, and route the correction through the authority that owns that concern.

For an active slice, old specs/contracts/code may be classified `KEEP`, `ADAPT`, `REPLACE`, or `DEFER`. Existing work is evidence, not sunk-cost authority and not disposable by default.

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

When manually invoking a Harscode phase, start from that phase's current canonical `workflow/*-prompt.md` entrypoint.

For CRTV and normal new work, use the canonical kickoff with its normal variables/context. Do **not** add a task-specific solution-steering prompt that tells the agent which authorities, gaps, or conclusions it is supposed to discover. Repository authority and the task reference should be sufficient; inability to discover derivable information is validation evidence.

Use `docs/kencleng-agentic-workflow.md` for Kencleng-specific orchestration such as MVP slice sequencing, risk tiers, human authority, backend/frontend coordination, design readiness, mock-first integration, and delivery/finalization. Read the relevant section(s); it is not mandatory startup prose for every phase.

Older feature/task docs may contain numeric section references to previous workflow revisions. Follow the current named rule/source owner rather than reviving superseded policy from an old `§NN` reference.

## 6. Scope and directory boundaries

Work on one coherent Harscode work unit at a time. Do not combine unrelated changes merely because they are nearby.

For the current MVP, the approved scope and sequencing live in:

- `docs/product/mvp-scope.md`;
- `docs/product/mvp-delivery-slices.md`.

Historical numbered domain order is not the current MVP delivery order.

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

- **Scope:** active product slice/spec/task fulfilled.
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
- durable Product Authority changes;
- current MVP scope/sequencing changes;
- product/domain contract changes where the owning authority requires it;
- protected spec/test changes;
- brand-defining visual assets;
- material `OPEN` product/design decisions;
- manual DB/index application;
- any additional human gate explicitly owned by the applicable scoped/project authority.

An agent must not approve its own work where independent human authority is required.

## 10. One-off docs and generated artifacts

One-off setup/playbook documents, when present, are on-demand context only. They do not override canonical Product/MVP authority, canonical design authority, reconciled delivery specs/contracts, architecture, root/scoped `AGENTS.md`, or Harscode workflow.

Generated exploration, techplan, build, review, testing, and PR artifacts are task evidence/history, not automatically project-wide precedent. Promote reusable truth into the source that owns that concern.
