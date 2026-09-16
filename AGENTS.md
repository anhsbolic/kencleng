# AGENTS.md — Kencleng

This is the root instruction map for AI coding agents working in the Kencleng repository.

Read this before writing code. More specific `AGENTS.md` files apply within their directory scope.

Kencleng is a sandbox donation/crowdfunding project used to practice correctness-critical, secure-by-design agentic engineering. Agent-generated code does not reduce correctness or security requirements.

---

## 1. Use the Authority That Owns the Concern

Do not treat this repository as having one flat document precedence for every question.

```text
domain invariants / threats
→ docs/spec/<domain>/invariants.md
→ docs/spec/<domain>/threat-model.md

feature behavior
→ docs/spec/<domain>/features/*.md

API contract
→ api/openapi.yaml

backend architecture
→ docs/project/kencleng-backend-tech-stack.md

frontend architecture
→ docs/project/kencleng-frontend-tech-stack.md

product-design principles
→ docs/ui-ux/product-design-principles.md

selected brand + product-UI direction
→ docs/ui-ux/brand-product-ui-brief.md

UX patterns
→ docs/ui-ux/patterns.md

asset governance
→ docs/ui-ux/asset-governance.md

route/persona mapping
→ docs/ui-ux/page-map.md

selected visual-direction evidence
→ docs/ui-ux/visual-references/selected-direction/

frontend component contracts
→ frontend/components/README.md

Kencleng project orchestration
→ docs/kencleng-agentic-workflow.md

feature lifecycle
→ Harscode workflow
```

Concrete production visual-system authority is intentionally absent until it is deliberately derived from the approved brand/product direction. Do not revive removed legacy design guidelines or prototypes as current authority.

For business behavior, domain invariants/threat models and feature specs outrank narrative project background; OpenAPI owns the API contract. If authorities genuinely conflict on the same concern, surface the contradiction instead of choosing whichever interpretation makes implementation easiest.

---

## 2. Golden Rules

- **Errors:** return or handle expected failures intentionally. Never silently swallow them.
- **Money:** backend monetary arithmetic uses the established decimal representation; never introduce `float64` for money, including fixtures.
- **SQL:** use the established parameterized query approach; never construct SQL from user-controlled values through interpolation/concatenation.
- **Sensitive data:** never log secrets, raw tokens, or PII payloads.
- **Client-visible errors:** never leak stack traces, raw SQL, filesystem paths, or implementation internals. Follow the Problem Details contract in `api/openapi.yaml`.
- **PII:** follow the established encryption/HMAC storage pattern; do not invent a second convention.
- **Authorization:** backend authorization is explicit. Frontend role gates or hidden controls are UX, not security authority.

---

## 3. File-Path Fencing

Tier 0 / explicitly protected paths are read-only to an agent unless the human explicitly authorizes work on them for that session.

Protected paths include:

- `backend/internal/domain/donation/ledger.go` and any file implementing transaction/locking logic for balance updates;
- state-machine implementation under `backend/internal/domain/disbursement/`;
- `backend/internal/platform/crypto/`;
- `backend/internal/platform/auth/` when changing security-critical authentication/session logic.

More-specific scoped `AGENTS.md` files may add protections; they must not silently weaken this root fencing.

If a task requires a protected write, surface the boundary instead of routing around it.

---

## 4. Spec / Test / Code Authority Separation

An implementation agent must not change an established requirement merely to make code pass.

Do not weaken an acceptance criterion, alter a domain invariant, loosen an approved test, or lower a verification threshold silently. If a spec or test appears wrong, report the contradiction and route the correction through the appropriate human/project authority.

---

## 5. Workflow Authority

Generic feature lifecycle comes from Harscode:

```text
Exploration + Techplan
→ Build / Patch
→ Code Review
→ Testing
→ Pull Request
```

Use Harscode for phase responsibilities, session boundaries, phase-specific reports, patch routing, and generic engineering practices.

Use `docs/kencleng-agentic-workflow.md` only for Kencleng-specific orchestration such as domain sequencing, risk tiers, human authority, backend/frontend coordination, design readiness, mock-first integration, and domain finalization.

Do not create a second generic feature lifecycle inside project instructions.

---

## 6. Scope Discipline

Work on one coherent Harscode feature/task work unit at a time. Do not combine unrelated changes merely because they are nearby. A small task unexpectedly touching many unrelated files is a risk signal.

---

## 7. Directory Boundary

Backend and frontend production writes remain separate by default. A backend-scoped Build must not casually modify `frontend/`, and vice versa. Cross-stack API changes require explicit coordination.

Reading shared `docs/` and `api/` for context is allowed. Integration Testing may exercise both sides together without granting uncontrolled cross-directory write scope.

---

## 8. Backend Scope

When working under `backend/`, read `backend/AGENTS.md` when present. Backend remains authority for financial correctness, authorization, domain lifecycle, transaction semantics, and persistence invariants.

---

## 9. Frontend Scope

When working under `frontend/`, read `frontend/AGENTS.md` before implementation.

Frontend must follow the current product-design principles, approved Brand + Product UI Brief, UX patterns, asset governance, component governance, and applicable product/domain truth.

Material UI work must classify design readiness as `READY`, `PARTIAL`, or `OPEN` and must not silently invent consequential product/design intent.

The removed legacy prototype/design-reference system is not current precedent. Selected visual references demonstrate direction only and are not pixel-perfect route specifications.

---

## 10. Shared Component Changes

Changes under `frontend/components/ui/` or `frontend/components/shared/` require impact analysis according to `frontend/components/README.md`: classify the change, discover consumers, identify representative risks, verify the component, and verify representative downstream consumers.

Compilation alone does not prove compatibility.

---

## 11. Required Verification

Use the target repository's actual commands/configuration. Do not claim a check ran unless it ran. Additional evidence is risk-driven.

Examples include invariant/property evidence, race/concurrency verification, threat-focused security verification, hostile-content rendering coverage, real-browser verification for material rendered frontend changes, and real backend/frontend integration before claiming integrated completion.

Do not indiscriminately run expensive test classes during Build when the Harscode phase boundary assigns them to Testing.

---

## 12. Browser Verification

The project has chosen Playwright as the target browser-automation capability. Until it is actually wired into the repository, do not invent a command or claim Playwright verification ran.

---

## 13. Required Pull-Request Evidence

A behavior-changing PR must identify scope, verification actually performed, and a risk note:

```markdown
## Risk note
- Assumptions made: ...
- Edge cases intentionally NOT handled (and why): ...
- Concurrency assumptions: ...
- What is not tested, and why: ...
```

Tier-specific evidence follows `docs/kencleng-agentic-workflow.md`.

For material frontend UI changes, identify the canonical design basis, material assumptions, rendered verification performed, and any provisional visual assets.

---

## 14. Risk Reporting

Distinguish `verified`, `assumed`, `deferred`, and `not tested`. A known material risk silently omitted is worse than a clearly documented limitation.

---

## 15. Human Authority

Human approval is required where defined for Tier-0 implementation, Tier-1 merge/review, product/domain contract changes, protected spec/test changes, brand-defining visual assets, material `OPEN` product/design decisions, and manual DB/index application.

An agent must not approve its own work where independent human authority is required.

---

## 16. One-Off Playbooks

Files under `.agents/docs/` are on-demand project-specific playbooks. They do not override canonical specs, architecture docs, UI/UX sources of truth, or Harscode workflow.

---

## 17. Generated Task Artifacts

Generated exploration, techplan, build, review, testing, and PR artifacts are task evidence/history. They are not automatically project-wide precedent. Promote reusable truth into the source that owns that concern.

---

## 18. Related Documents

Project:

- `docs/kencleng-agentic-workflow.md`
- `docs/kencleng-repo-setup.md`
- `docs/spec/README.md`
- `docs/project/kencleng-backend-tech-stack.md`
- `docs/project/kencleng-frontend-tech-stack.md`

Frontend design:

- `docs/ui-ux/README.md`
- `docs/ui-ux/product-design-principles.md`
- `docs/ui-ux/brand-product-ui-brief.md`
- `docs/ui-ux/patterns.md`
- `docs/ui-ux/asset-governance.md`
- `docs/ui-ux/page-map.md`
- `docs/ui-ux/visual-references/selected-direction/`
- `frontend/components/README.md`

Portable engineering workflow:

- Harscode `workflow/`
- Harscode `best-practices/`
