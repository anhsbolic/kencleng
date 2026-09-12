# AGENTS.md — Kencleng

Root instruction map for AI coding agents working in this repository.

Read this file before writing code. More-specific `AGENTS.md` files apply within their directory scope.

Kencleng is a sandbox donation/crowdfunding project for correctness-critical, secure-by-design agentic engineering. Agent-generated code does not reduce correctness, security, or evidence requirements.

## 1. Route to the authority that owns the concern

```text
domain invariants / threats
→ docs/spec/<domain-dir>/invariants.md
→ docs/spec/<domain-dir>/threat-model.md

feature behavior / acceptance criteria
→ docs/spec/<domain-dir>/features/*.md

API contract
→ api/openapi.yaml

backend architecture
→ docs/project/kencleng-backend-tech-stack.md

frontend architecture
→ docs/project/kencleng-frontend-tech-stack.md

product-design principles
→ docs/ui-ux/product-design-principles.md

UX patterns
→ docs/ui-ux/patterns.md

visual system
→ docs/ui-ux/design-guidelines.md

brand / visual assets
→ docs/ui-ux/brand-and-visual-assets.md

route/persona mapping
→ docs/ui-ux/page-map.md

prototype authority / consumption
→ docs/ui-ux/prototype-reference.md
→ docs/ui-ux/design-reference-usage.md

frontend component contracts
→ frontend/components/README.md

project status
→ docs/project/kencleng-development-tracker.md

Kencleng project orchestration
→ docs/kencleng-agentic-workflow.md

feature lifecycle / generic engineering practice
→ Harscode workflow / best-practices
```

For business behavior, domain invariants/threat models and feature specs are authoritative over narrative project background; `api/openapi.yaml` owns the API shape. Do not apply that precedence to unrelated concerns owned by architecture or design docs.

If authorities genuinely conflict on the same concern, surface the contradiction instead of choosing whichever interpretation makes implementation easiest.

## 2. Golden rules

- **Errors:** return/handle failures intentionally. Never silently swallow expected errors.
- **Money:** backend monetary arithmetic uses the established decimal representation; never introduce `float64` for money, including fixtures.
- **SQL:** use the established parameterized `goqu` approach; never construct SQL from user-controlled values via interpolation/concatenation.
- **Sensitive data:** never log secrets, raw tokens, or PII payloads.
- **Client errors:** never leak stack traces, raw SQL, internal filesystem paths, or other implementation internals. Follow the Problem Details contract in `api/openapi.yaml`.
- **PII:** follow the established encryption/HMAC storage pattern; do not invent a second convention.
- **Authorization:** backend authorization checks are explicit. Frontend role gates or hidden controls are UX, not security authority.

## 3. File-path fencing

The following paths are read-only to an agent unless a human explicitly authorizes work on them for that session:

- `backend/internal/domain/donation/ledger.go` and any file implementing transaction/locking logic for balance updates;
- the state-machine implementation under `backend/internal/domain/disbursement/`;
- `backend/internal/platform/crypto/` (encryption, HMAC, key handling);
- `backend/internal/platform/auth/` when changing security-critical JWT signing, TOTP, refresh-token/session, or equivalent authentication core logic;
- `design-reference/` at repo root.

The first four are Tier-0 / human-authored or human-paired areas. Agents may read them for context, critique, test ideas, or adversarial review but must not modify them without explicit authorization.

`design-reference/` is protected for a different reason: it is frozen prototype/reference output. Agents may inspect it for visual/structural precedent, but must never modify it, wholesale-copy it into production, or treat prototype state/data/component architecture as production authority.

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
Exploration + Techplan
→ Build / Patch
→ Code Review
→ Testing
→ Pull Request
```

Use Harscode for phase responsibilities, session boundaries, patch routing, reports, and generic engineering best-practices.

Use `docs/kencleng-agentic-workflow.md` only for Kencleng-specific orchestration such as domain sequencing, risk tiers, human authority, backend/frontend coordination, design readiness, mock-first integration, and domain finalization.

Do not create a second feature lifecycle in project instructions.

## 6. Scope and directory boundaries

Work on one coherent Harscode work unit at a time. Do not combine unrelated changes merely because they are nearby. Session boundaries follow the active Harscode workflow rather than an unconditional one-endpoint-per-session rule.

Backend and frontend production writes remain separate by default:

- backend-scoped Build → do not casually modify `frontend/`;
- frontend-scoped Build → do not casually modify `backend/`;
- shared `docs/` and `api/` may be read for context;
- cross-stack contract changes require explicit coordination.

A small task unexpectedly touching many unrelated files is a risk signal. Investigate before expanding scope.

## 7. Backend scope

When working under `backend/`, read `backend/AGENTS.md` plus only the domain/project material relevant to the task.

Backend remains authoritative for financial correctness, authorization, domain lifecycle, transaction semantics, and persistence invariants. Frontend behavior must not compensate for missing backend correctness.

## 8. Frontend scope

When working under `frontend/`, read `frontend/AGENTS.md` before implementation.

Material frontend UI work must use the current product-design, UX, visual, asset, prototype, and component authorities and classify design readiness as `READY`, `PARTIAL`, or `OPEN` according to `docs/ui-ux/product-design-principles.md`.

Changes to `frontend/components/ui/` or `frontend/components/shared/` require consumer/blast-radius analysis according to `frontend/components/README.md`.

Material rendered UI changes require rendered verification according to frontend guidance and Harscode Testing.

## 9. Verification and evidence

Use the repository's actual executable commands/configuration. Never claim a check ran unless it ran.

Evidence is risk-driven. Examples include:

- invariant/property evidence for applicable correctness-critical backend rules;
- race/concurrency evidence for relevant concurrent code;
- threat-focused security verification;
- hostile-content rendering coverage;
- real-browser verification for material frontend UI changes;
- real backend/frontend integration before claiming integrated completion.

Do not indiscriminately run expensive test classes during Build when Harscode assigns them to Testing.

The project has selected Playwright as the target browser-automation capability, but until it is actually wired into the repository do not invent a command or claim Playwright verification ran. Once available, use it for appropriate rendered interaction, responsive checks, integration verification, and high-value browser regressions—not as a mandate for a giant E2E suite.

## 10. Pull-request evidence

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

Tier-specific requirements follow `docs/kencleng-agentic-workflow.md`. Use evidence appropriate to the responsibility being changed.

For material frontend UI changes also identify the design basis/precedent, material design assumptions, rendered verification performed, and any provisional visual assets.

Distinguish `verified`, `assumed`, `deferred`, and `not tested`. A known material risk silently omitted is worse than a clearly documented limitation.

## 11. Human authority

Human approval is required where defined for:

- Tier-0 implementation;
- Tier-1 merge/review;
- product/domain contract changes;
- protected spec/test changes;
- brand-defining visual assets;
- material `OPEN` product/design decisions;
- manual DB/index application.

An agent must not approve its own work where independent human authority is required.

## 12. One-off docs and generated artifacts

One-off setup/playbook documents, when present, are on-demand context only. They do not override canonical specs, architecture docs, UI/UX authorities, root/scoped `AGENTS.md`, or Harscode workflow.

Generated exploration, techplan, build, review, testing, and PR artifacts are task evidence/history, not automatically project-wide precedent. Promote reusable truth into the source that owns that concern.

## 13. Related documents

Project:

- `docs/kencleng-agentic-workflow.md` — Kencleng orchestration overlay
- `docs/project/kencleng-development-tracker.md` — current cross-domain/project status
- `docs/project/kencleng-repo-setup.md`
- `docs/spec/README.md`
- `docs/project/kencleng-backend-tech-stack.md`
- `docs/project/kencleng-frontend-tech-stack.md`

Frontend design:

- `docs/ui-ux/product-design-principles.md`
- `docs/ui-ux/patterns.md`
- `docs/ui-ux/design-guidelines.md`
- `docs/ui-ux/brand-and-visual-assets.md`
- `docs/ui-ux/page-map.md`
- `docs/ui-ux/prototype-reference.md`
- `docs/ui-ux/design-reference-usage.md`
- `frontend/components/README.md`

Portable engineering workflow:

- Harscode `workflow/`
- Harscode `best-practices/`
