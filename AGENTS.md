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

selected brand + Product UI direction
→ docs/ui-ux/brand-product-ui-brief.md

UX patterns
→ docs/ui-ux/patterns.md

asset governance
→ docs/ui-ux/asset-governance.md

route/persona mapping
→ docs/ui-ux/page-map.md

selected direction visual evidence
→ docs/ui-ux/visual-references/selected-direction/

frontend component contracts
→ frontend/components/README.md

Kencleng project orchestration
→ docs/kencleng-agentic-workflow.md

feature lifecycle
→ Harscode workflow
```

There is intentionally no current canonical concrete visual-system document defining exact production colors, fonts, radii, spacing tokens, icon family, or motion values. Those must be deliberately derived from `brand-product-ui-brief.md`; existing frontend values are not design authority merely because they exist.

For business behavior:

```text
domain invariant / threat model
→ feature spec
→ OpenAPI
```

has higher authority than narrative project background or visual artifacts.

If authorities genuinely conflict on the same concern, surface the contradiction. Do not silently choose whichever interpretation makes implementation easiest.

---

## 2. Golden Rules

### Errors
Errors are returned or handled intentionally. Never silently swallow expected failures.

### Money
Backend monetary arithmetic uses the established decimal representation. Never introduce `float64` for money, including fixtures.

### SQL
Use the established parameterized query approach. Never construct SQL from user-controlled values through string interpolation/concatenation.

### Sensitive data
Never log secrets, raw tokens, or PII payloads. Log safe operational facts instead.

### Error responses
Client-visible errors must not leak stack traces, raw SQL, internal filesystem paths, or other implementation internals. Follow the Problem Details contract in `api/openapi.yaml`.

### PII
Follow the established encryption/HMAC storage pattern from backend architecture. Do not invent a second PII-storage convention.

### Authorization
Authorization is explicit at the backend boundary. A frontend role gate or hidden control is UX, not security authority.

---

## 3. File-Path Fencing

Tier 0 / explicitly protected paths are read-only to an agent unless the human explicitly authorizes work on them for that session.

Follow the current protected-path list and local scoped `AGENTS.md`.

The former `docs/design-reference/` prototype generation is no longer active design authority. Historical versions belong to Git history, not as a current implementation reference.

If a task requires writing an explicitly protected path, surface the boundary instead of routing around it.

---

## 4. Spec / Test / Code Authority Separation

An implementation agent must not change an established requirement merely to make its code pass.

Do not weaken feature acceptance criteria, domain invariants, approved tests, or verification thresholds solely to accommodate implementation.

If a spec or test appears wrong:

```text
report contradiction
→ explain evidence
→ route correction through appropriate human/project authority
```

---

## 5. Workflow Authority

Generic feature lifecycle comes from Harscode.

Default lifecycle:

```text
Exploration + Techplan
→ Build / Patch
→ Code Review
→ Testing
→ Pull Request
```

Use Harscode for phase responsibilities, session boundaries, phase-specific reports, patch routing, and generic engineering practices.

Use `docs/kencleng-agentic-workflow.md` only for Kencleng-specific orchestration such as domain sequencing, risk tiers, human authority, backend/frontend coordination, design readiness, mock-first integration, and domain finalization.

Do not create a second feature lifecycle inside project instructions.

---

## 6. Scope Discipline

Work on one coherent Harscode feature/task work unit at a time.

Do not combine unrelated changes merely because they are nearby. A small task unexpectedly touching many unrelated files is a risk signal; investigate before expanding scope.

---

## 7. Directory Boundary

Backend and frontend production writes remain separate by default.

A Build scoped to `backend/` must not casually modify `frontend/`, and vice versa. Cross-stack API changes require explicit coordination.

Reading shared `docs/` and `api/` for context is allowed.

---

## 8. Backend Scope

When working under `backend/`, read `backend/AGENTS.md` when present and only the relevant domain/project material.

Backend remains authority for financial correctness, authorization, domain lifecycle, transaction semantics, and persistence invariants.

Frontend behavior must not compensate for missing backend correctness.

---

## 9. Frontend Scope

When working under `frontend/`, read `frontend/AGENTS.md` before implementation.

Frontend must follow current:

```text
product-design principles
brand + Product UI brief
UX patterns
asset governance
component governance
```

Material frontend UI work must classify design readiness:

```text
READY
PARTIAL
OPEN
```

and must not silently invent consequential product/design intent.

The approved direction is **Sunlit Editorial / Evidence-Led Optimism**. Selected visual references are direction evidence, not screenshots to clone.

Rendered UI changes require rendered verification according to frontend/Harscode Testing guidance.

---

## 10. Shared Component Changes

Changes to `frontend/components/ui/` or `frontend/components/shared/` require impact analysis according to `frontend/components/README.md`.

At minimum:

```text
classify change
→ discover consumers
→ identify representative risks
→ verify component
→ verify representative downstream consumers
```

Compilation alone does not prove compatibility.

---

## 11. Required Verification

Use the target repository's actual commands/configuration. Do not claim a check ran unless it ran.

Root/project verification remains the baseline required by affected scope. Additional evidence is risk-driven.

Examples include invariant/property evidence, concurrency verification, threat-focused security verification, hostile-content rendering coverage, real-browser verification for material frontend changes, and real backend/frontend integration before claiming integrated completion.

Do not indiscriminately run expensive test classes during Build when Harscode assigns them to Testing.

---

## 12. Browser Verification

The project has chosen Playwright as the target standard browser-automation capability.

Until Playwright is actually wired into the repository, do not invent a command or claim Playwright verification ran.

Once available, use it for appropriate rendered interaction, responsive verification, real integration checks, and high-value regression coverage.

---

## 13. Required Pull-Request Evidence

A behavior-changing PR must identify:

### Scope
Which feature/spec/task it fulfills.

### Verification
What commands/tests/checks actually ran.

### Risk note

```markdown
## Risk note
- Assumptions made: ...
- Edge cases intentionally NOT handled (and why): ...
- Concurrency assumptions: ...
- What is not tested, and why: ...
```

### Tier-specific evidence
Follow `docs/kencleng-agentic-workflow.md` and use evidence appropriate to the responsibility being changed.

### Frontend design evidence
For material frontend UI changes, identify:

- design basis / relevant canonical authority;
- material design assumptions;
- rendered verification performed;
- provisional visual assets, if any.

Do not present provisional brand/design work as canonical.

---

## 14. Risk Reporting

Reporting an actual limitation, assumption, or unresolved risk is useful output.

Distinguish:

```text
verified
assumed
deferred
not tested
```

A known material risk silently omitted is worse than a clearly documented limitation.

---

## 15. Human Authority

Human approval is required where defined for Tier 0 implementation, Tier 1 merge/review, product/domain contract changes, protected spec/test changes, brand-defining visual assets, material OPEN product/design decisions, and manual DB/index application.

An agent must not approve its own work where independent human authority is required.

---

## 16. One-Off Playbooks

Files under `.agents/docs/` are on-demand project-specific playbooks.

Read them only when relevant. They do not override canonical specs, architecture docs, UI/UX sources of truth, or Harscode workflow.

---

## 17. Generated Task Artifacts

Generated exploration, techplan, build, review, testing, and PR artifacts belong in the project's established work-artifact location.

They are task evidence/history and are not automatically project-wide precedent.

When a decision becomes reusable project truth, promote it into the source that owns that concern.

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
