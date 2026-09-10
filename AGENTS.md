# AGENTS.md — Kencleng

This is the root instruction map for AI coding agents working in the Kencleng repository.

Read this before writing code.

More specific `AGENTS.md` files apply within their directory scope.

Kencleng is a sandbox donation/crowdfunding project used to practice correctness-critical, secure-by-design agentic engineering.

Agent-generated code does not reduce correctness or security requirements.

---

## 1. Use the Authority That Owns the Concern

Do not treat this repository as having one flat document precedence for every question.

Use the source that owns the concern.

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

UX patterns
→ docs/ui-ux/patterns.md

visual system
→ docs/ui-ux/design-guidelines.md

brand / visual assets
→ docs/ui-ux/brand-and-visual-assets.md

route/persona mapping
→ docs/ui-ux/page-map.md

prototype authority
→ docs/ui-ux/prototype-reference.md

frontend component contracts
→ frontend/components/README.md

Kencleng project orchestration
→ docs/kencleng-agentic-workflow.md

feature lifecycle
→ Harscode workflow
```

For business behavior:

```text
domain invariant / threat model
→ feature spec
→ OpenAPI
```

has higher authority than narrative project background.

If authorities genuinely conflict on the same concern, surface the contradiction.

Do not silently pick whichever interpretation makes implementation easiest.

---

## 2. Golden Rules

These apply across the repository.

### Errors

Errors are returned or handled intentionally.

Never silently swallow expected failures.

### Money

Backend monetary arithmetic uses the established decimal representation.

Never introduce `float64` for money, including fixtures.

### SQL

Use the established parameterized query approach.

Never construct SQL from user-controlled values through string interpolation/concatenation.

### Sensitive data

Never log:

* secrets;
* raw tokens;
* PII payloads.

Log safe operational facts instead.

### Error responses

Client-visible errors must not leak:

* stack traces;
* raw SQL;
* internal filesystem paths;
* other implementation internals.

Follow the Problem Details contract in `api/openapi.yaml`.

### PII

Follow the established encryption/HMAC storage pattern from the backend architecture.

Do not invent a second PII-storage convention.

### Authorization

Authorization is explicit at the backend boundary.

A frontend role gate or hidden control is UX, not security authority.

---

## 3. File-Path Fencing

Tier 0 / explicitly protected paths are read-only to an agent unless the human explicitly authorizes work on them for that session.

Follow the current protected-path list and local scoped `AGENTS.md`.

### `design-reference/`

`design-reference/` is also read-only, for a different reason.

It is frozen prototype/reference output.

Agents may inspect it for visual/structural intent.

Agents must not:

* modify it;
* wholesale-copy its source into production;
* treat prototype state/data/component architecture as production authority.

Use:

```text
docs/ui-ux/prototype-reference.md
docs/ui-ux/design-reference-usage.md
```

when consuming it.

If a task requires writing an explicitly protected path, surface the boundary instead of routing around it.

---

## 4. Spec / Test / Code Authority Separation

An implementation agent must not change an established requirement merely to make its code pass.

Do not:

* weaken a feature acceptance criterion to match implementation;
* alter a domain invariant because code violates it;
* loosen an already-approved test solely to obtain green output;
* lower a verification threshold silently.

If a spec or test appears wrong:

```text
report the contradiction
→ explain evidence
→ route the correction through appropriate human/project authority
```

Do not hide requirement changes inside an implementation patch.

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

Use Harscode for:

* phase responsibilities;
* session boundaries;
* phase-specific reports;
* patch routing;
* generic engineering best practices.

Use:

```text
docs/kencleng-agentic-workflow.md
```

only for Kencleng-specific orchestration such as:

* domain sequencing;
* risk tiers;
* human authority;
* backend/frontend coordination;
* design readiness;
* mock-first integration;
* domain finalization.

Do not create a second feature lifecycle inside project instructions.

---

## 6. Scope Discipline

Work on one coherent Harscode feature/task work unit at a time.

Do not combine unrelated changes merely because they are nearby.

Session boundaries follow the active Harscode workflow rather than an unconditional "one endpoint = one session" rule.

A small task unexpectedly touching many unrelated files is a risk signal.

Investigate before expanding scope.

---

## 7. Directory Boundary

Backend and frontend production writes remain separate by default.

A Build scoped to:

```text
backend/
```

must not casually modify:

```text
frontend/
```

and vice versa.

Cross-stack API changes require explicit coordination.

Reading shared:

```text
docs/
api/
```

for context is allowed.

Integration Testing may exercise both sides together without granting uncontrolled cross-directory write scope.

Use nested `AGENTS.md` for directory-specific rules.

---

## 8. Backend Scope

When working under:

```text
backend/
```

read:

```text
backend/AGENTS.md
```

when present.

Also read only the domain/project material relevant to the task.

Backend remains authority for:

* financial correctness;
* authorization;
* domain lifecycle;
* transaction semantics;
* persistence invariants.

Frontend behavior must not compensate for missing backend correctness.

---

## 9. Frontend Scope

When working under:

```text
frontend/
```

read:

```text
frontend/AGENTS.md
```

before implementation.

Frontend must follow current:

```text
product-design principles
UX patterns
visual system
brand/asset system
component governance
```

Material frontend UI work must classify design readiness:

```text
READY
PARTIAL
OPEN
```

and must not silently invent consequential product/design intent.

Rendered UI changes require rendered verification according to frontend/Harscode Testing guidance.

---

## 10. Shared Component Changes

Changes to:

```text
frontend/components/ui/
frontend/components/shared/
```

require impact analysis according to:

```text
frontend/components/README.md
```

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

Use the target repository's actual commands/configuration.

Do not claim a check ran unless it ran.

Root/project verification remains the baseline required by the affected scope.

Additional evidence is risk-driven.

Examples:

* invariant/property evidence for applicable correctness-critical backend rules;
* race/concurrency verification for relevant concurrent code;
* threat-focused security verification;
* hostile-content rendering coverage;
* real-browser verification for material rendered frontend changes;
* real backend/frontend integration before claiming integrated completion.

Do not indiscriminately run expensive test classes during the Build loop when the Harscode phase boundary assigns them to Testing.

---

## 12. Browser Verification

The project has chosen Playwright as the target standard browser-automation capability.

Until Playwright is actually wired into the repository:

* do not invent a command;
* do not claim Playwright verification ran.

Once available, use it for appropriate:

* rendered interaction;
* responsive verification;
* real integration checks;
* high-value browser regression tests.

Playwright availability does not mean every feature requires a permanent comprehensive E2E test.

---

## 13. Required Pull-Request Evidence

A behavior-changing PR must identify:

### Scope

Which feature/spec/task it fulfills.

### Verification

What commands/tests/checks actually ran.

### Risk note

Use:

```markdown
## Risk note
- Assumptions made: ...
- Edge cases intentionally NOT handled (and why): ...
- Concurrency assumptions: ...
- What is not tested, and why: ...
```

Claims of handled risk should point to concrete evidence when such evidence exists.

### Tier-specific evidence

Tier 0/1 requirements follow:

```text
docs/kencleng-agentic-workflow.md
```

Do not force a frontend-only UI PR to create a meaningless property test merely because the backing backend feature is Tier 1.

Use evidence appropriate to the responsibility being changed.

### Frontend design evidence

For material frontend UI changes, identify:

* design basis / relevant precedent;
* material design assumptions;
* rendered verification performed;
* provisional visual assets, if any.

Do not present provisional brand/design work as canonical.

---

## 14. Risk Reporting

Reporting an actual limitation, assumption, or unresolved risk is useful output.

Do not optimize reports toward:

```text
everything looks done
```

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

Follow project risk/design authority.

Human approval is required where defined for:

* Tier 0 implementation;
* Tier 1 merge/review;
* product/domain contract changes;
* protected spec/test changes;
* brand-defining visual assets;
* material OPEN product/design decisions;
* manual DB/index application.

An agent must not approve its own work where independent human authority is required.

---

## 16. One-Off Playbooks

Files under:

```text
.agents/docs/
```

are on-demand project-specific playbooks.

Read them only when relevant.

They do not override canonical:

* specs;
* architecture docs;
* UI/UX sources of truth;
* Harscode workflow.

A one-off scaffold/setup playbook must not become the permanent owner of living project policy or status.

---

## 17. Generated Task Artifacts

Generated exploration, techplan, build, review, testing, and PR artifacts belong in the target project's established work-artifact location.

They are task evidence/history.

They are not automatically project-wide precedent.

When a decision becomes reusable project truth, promote it into the source that owns that concern.

---

## 18. Related Documents

Project:

* `docs/kencleng-agentic-workflow.md` — Kencleng orchestration overlay
* `docs/kencleng-repo-setup.md`
* `docs/spec/README.md`
* `docs/project/kencleng-backend-tech-stack.md`
* `docs/project/kencleng-frontend-tech-stack.md`

Frontend design:

* `docs/ui-ux/product-design-principles.md`
* `docs/ui-ux/patterns.md`
* `docs/ui-ux/design-guidelines.md`
* `docs/ui-ux/brand-and-visual-assets.md`
* `docs/ui-ux/page-map.md`
* `docs/ui-ux/prototype-reference.md`
* `docs/ui-ux/design-reference-usage.md`
* `frontend/components/README.md`

Portable engineering workflow:

* Harscode `workflow/`
* Harscode `best-practices/`
