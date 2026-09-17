# AGENTS.md — frontend/

Frontend-specific rules on top of root `AGENTS.md`.

Read root `AGENTS.md` first. This file is the always-applicable frontend rule/router surface; detailed rationale remains in the canonical frontend architecture/design/component documents it points to.

Scope: `frontend/`.

Do not modify backend production code from a frontend-scoped Build.

Use progressive disclosure for project authorities: identify the active concern, open the authoritative section(s) that govern it, follow referenced dependencies, and expand only when a cross-cutting decision/check requires broader context. A canonical file being important does not make every section mandatory on every frontend task.

## 1. Project map

The following paths are **semantic ownership destinations**, not a requirement that every directory already exists. The clean-start frontend intentionally creates a layer only when real implementation needs it.

```text
frontend/
├── app/                    # App Router; route-local composition may live here
├── components/
│   ├── ui/                 # generic design-system primitives, when established
│   ├── features/<domain>/  # domain-semantic components, when needed
│   └── shared/             # genuinely cross-domain semantic components, when proven
├── lib/
│   ├── api/                # focused typed API functions / generated types when needed
│   ├── hooks/              # TanStack Query hooks when needed
│   └── stores/             # genuinely shared client-owned Zustand state only
├── mocks/                  # MSW handlers when a task needs mocks
├── tests/browser/          # optional committed Playwright browser regressions
└── public/                 # justified static assets only
```

Component placement is **semantic-owner-first**, not reusable-first.

Canonical component governance: `components/README.md`.

The previous frontend product implementation was intentionally retired. Git history is archive/evidence, not current implementation precedent. During the reboot/readiness window, follow `../docs/project/frontend-reboot-plan.md`; after the clean baseline is frozen, normal Harscode feature development resumes.

## 2. Product, business, and API authority

Frontend is not product/business authority.

Start from the active product slice:

```text
../docs/product/product-overview.md
+ ../docs/product/mvp-scope.md
+ ../docs/product/mvp-delivery-slices.md
+ relevant ../docs/ui-ux authority
        ↓
reconciled delivery/domain spec
        ↓
reconciled API contract
        ↓
frontend implementation
```

Do not recreate backend decisions for money, eligibility, permissions, verification, lifecycle transitions, or financial validity.

Do not let an old domain spec or OpenAPI shape silently override canonical Product/MVP truth. Existing contracts/code are evidence until reconciled for the active slice.

Client validation, presentation logic, formatting, interaction state, and derived presentation values are legitimate frontend responsibilities.

If required **product meaning** is missing, surface the Product Authority gap. If product meaning is clear but API/spec data is missing, surface the delivery/contract gap instead of inventing business behavior in React.

API request/response shape comes from generated OpenAPI types based on the reconciled canonical OpenAPI sources. Do not maintain parallel handwritten API models for shapes the contract already owns.

For contract inspection, start with `../api/README.md`, then prefer the active domain source (`../api/openapi/<domain>.yaml`) plus only the referenced components from `../api/openapi/common.yaml`. Use the bundled `../api/openapi.yaml` when an aggregate/cross-domain view or generated-client correspondence is actually needed; do not load the whole bundle by default for a domain-local task.

For contract-parallel frontend work, keep production data access pointed at the real API contract. Use MSW at the network boundary for contract-faithful mock responses while the backend implementation is unavailable; do not add production service branches that return mock JSON based on environment/mode.

## 3. State ownership

Before creating state, use this order:

```text
derivable
→ derive it

server/API authoritative
→ Server Component / TanStack Query owner as appropriate

navigation/share/bookmark/back-forward state
→ URL when appropriate and non-sensitive

form lifecycle
→ React Hook Form

ephemeral UI interaction
→ narrowest local React owner

genuinely shared client-owned state
→ Zustand only when justified
```

Hard rules:

- do not mirror server data into Zustand;
- do not create a store merely because a domain exists;
- do not synchronize deterministic projections through effects;
- do not put sensitive values in the URL for convenience;
- do not recreate retired store/provider architecture merely because it existed before the reboot.

Detailed architecture: `../docs/project/kencleng-frontend-tech-stack.md`. Open the sections governing the active architecture/state/API concern rather than treating the whole architecture document as mandatory startup prose.

## 4. Forms and user-controlled content

Forms use React Hook Form + Zod for UX validation when a task actually needs a form. Server validation remains authoritative.

Distinguish field validation from request/business failure.

User-controlled Markdown/HTML must use an established safe rendering/sanitization path. Do not manually convert user content into unsafe `dangerouslySetInnerHTML`.

Any such rendering requires hostile-content test coverage.

## 5. Component ownership and shared-change impact

Use the narrowest truthful owner:

```text
route-specific composition → app/<route>/
domain concept             → components/features/<domain>/
cross-domain semantic      → components/shared/
generic primitive          → components/ui/
```

Do not extract by line count or promote by usage count alone. Repetition is evidence to evaluate an abstraction, not proof that one is needed.

Before materially changing anything in `components/ui/` or `components/shared/`:

```text
classify change
→ discover actual consumers/wrappers
→ identify representative risk cases
→ implement
→ verify component contract
→ verify representative downstream consumers
```

Compilation alone does not prove visual, behavioral, semantic, responsive, or accessibility compatibility.

Read `components/README.md` when introducing or materially changing a broad `ui/` or `shared/` contract, moving a component between ownership layers, or deciding a non-obvious abstraction boundary. Ordinary route-local work does not require reading the full component-governance document merely because it uses an existing primitive.

At the clean-start baseline the reusable registry is intentionally empty. Do not recreate historical Button/Badge/Input/etc. contracts until real new-generation usage establishes them.

## 6. Product design readiness and authority routing

Frontend development is **design-authority-driven, not design-file-driven**. A high-fidelity UI artifact is not a universal prerequisite for Build.

For material UI work classify design readiness according to `../docs/ui-ux/product-design-principles.md`:

```text
READY   → implement established intent directly
PARTIAL → resolve ordinary presentation gaps autonomously from established authority
OPEN    → distinguish material design ambiguity from missing product truth
```

Use this decision path:

```text
clear intent / READY
→ implement

ordinary presentation ambiguity
→ derive from principles + patterns + visual system + production precedent
→ implement

material design ambiguity
→ propose a small set of low-fidelity alternatives
→ recommend a default with trade-offs
→ obtain human decision when material
→ implement

missing durable product/MVP truth
→ surface the Product Authority gap
→ resolve in ../docs/product/
→ do not invent it in frontend code

product truth clear but delivery/contract data missing
→ surface owning spec/API gap
→ do not invent backend semantics in frontend code
```

Low-fidelity alternatives may be textual, diagrammatic, or wireframe-like. They exist to resolve meaningful hierarchy/interaction choices, not to create approval ceremony. Do not stop for routine micro-decisions already governed by current design authority.

Do not silently invent consequential product or interaction intent while coding.

Use `../docs/ui-ux/README.md` as the active UI/UX routing entrypoint. Open only the concern owner needed for the task:

| Active concern | Authority to open |
|---|---|
| design readiness, product/design decision authority, trust/clarity principles | `../docs/ui-ux/product-design-principles.md` |
| approved brand/product UI thesis, public-vs-product expression, trust language, visual invariants | `../docs/ui-ux/brand-product-ui-brief.md` |
| concrete color, typography, spacing, surfaces, radius, iconography, provenance grammar, progress grammar | `../docs/ui-ux/design-guidelines.md` |
| established recurring interaction/feedback/form/review pattern | matching section(s) of `../docs/ui-ux/patterns.md` |
| expressive/product-semantic/brand asset need, asset status or generation/handoff | `../docs/ui-ux/asset-governance.md` |
| route, persona, navigation/IA relationship | matching route/persona section of `../docs/ui-ux/page-map.md` |
| approved visual-direction evidence | `../docs/ui-ux/visual-references/selected-direction/` |
| broad component contract/placement/change blast radius | `components/README.md` |

The approved direction is **Sunlit Editorial / Evidence-Led Optimism** and the concrete visual system is owned by `../docs/ui-ux/design-guidelines.md`.

Do not resurrect removed legacy visual guidelines, prototype exports, green identity rules, old fonts, old tokens, or retired component decisions from Git history as current precedent.

Material UI normally needs product-design readiness plus only the specific behavior/visual/asset authorities relevant to the feature. It does not require reading every UI/UX document or producing a high-fidelity design artifact first.

## 7. Visual assets

Standard library icons are appropriate for ordinary utility actions. The current utility-icon baseline is Phosphor Icons as defined by `../docs/ui-ux/design-guidelines.md`.

Do not install/use a competing icon family merely because one existed in the retired frontend or remains in historical package-lock evidence.

Do not silently replace a materially important expressive/brand asset need with generic iconography, random gradients, stock-like imagery, synthetic documentary people, or generic AI decoration.

When a required asset is missing, follow `../docs/ui-ux/asset-governance.md`:

- reuse a canonical asset when one exists;
- use the approved utility-icon library when that is semantically sufficient;
- when a custom/expressive asset is justified, recommend materially distinct directions when a design choice is needed;
- generate a candidate when the current harness can do so adequately;
- otherwise produce an asset brief + ready-to-use generation prompt for human/tool handoff;
- keep generated/temporary assets explicitly provisional until their lifecycle state changes;
- require human approval for brand-defining assets and other reusable precedent where governance requires it.

Illustration must not masquerade as real campaign evidence. Photography or imagery must not imply beneficiary identity, distribution, verification, or impact without supporting product truth. When evidence is unavailable, preserve a truthful missing/provisional state instead of generating synthetic evidence.

Tool limitation must not silently become design limitation.

## 8. Visual system

Concrete production visual-system authority lives in:

```text
../docs/ui-ux/design-guidelines.md
```

Use it for the approved:

- warm-neutral palette and exact core color values;
- Newsreader / Instrument Sans role separation and type scales;
- spacing and density posture;
- border-first / spacing-first grouping;
- radius and elevation family;
- Phosphor utility-icon baseline;
- action hierarchy;
- provenance/truth presentation grammar;
- funding vs operational vs reported-outcome visual grammar;
- campaign imagery / placeholder treatment;
- public-expressive vs product-disciplined relationship.

Bootstrap `app/globals.css`, package availability, Git history, or incidental implementation choices are not higher visual authority. New production tokens/components are derived deliberately from current design authority and real usage; do not rebuild the retired token/component taxonomy for compatibility.

Visual-system semantics are not permission to invent business behavior. Product/MVP truth plus the reconciled domain/API contract remain authoritative for what a status, report, verification state, amount, or outcome actually means.

## 9. Selected visual references

Approved direction evidence lives under:

```text
../docs/ui-ux/visual-references/selected-direction/
```

Use it to understand visual character, hierarchy, public-vs-product expressive intensity, restrained editorial-yellow usage, and the Evidence Journal concept.

Do not treat those references as product/domain truth, component architecture, route contracts, or pixel-perfect screenshots to clone.

When a concern is concretely owned by `design-guidelines.md`, the guideline owns reusable system behavior and the images remain supporting visual evidence.

If visual evidence conflicts with canonical Product/MVP truth or a reconciled domain/API fact, the owning semantic authority wins.

## 10. Rendered iteration and human acceptance

Material rendered UI or spatial-interaction changes need real rendered feedback; compilation, Vitest, RTL, and MSW do not prove hierarchy, responsiveness, clipping/overflow, or whether an interaction feels understandable in context.

During Build, agent-driven `edit → render → inspect → fix` is valid implementation feedback when the selected client/environment supports it. Use the smallest representative set of realistic states, content conditions, and viewports capable of exposing likely failures.

Agent rendered feedback is **not final product acceptance**.

Before merge/delivery of material frontend UI, a human must manually exercise the rendered result in a real browser at a representative scope appropriate to the change. Human acceptance should focus on:

- visual hierarchy and clarity;
- interaction comprehension;
- responsive usability;
- product/design intent;
- asset appropriateness;
- obvious loading/empty/error/success behavior in context.

Keep the human acceptance scope proportional to the change. Screenshots may support discussion/evidence, but they are not a ritual and do not replace actual interaction where interaction matters.

## 11. Testing and Playwright boundary

Automated baseline capabilities:

```text
Vitest
React Testing Library
MSW
```

Tests should verify observable behavior rather than internal structure. Absence of tests at the clean scaffold baseline is valid; new tests are created with behavior worth protecting.

Current local commands:

```bash
npm run dev
npm run build
npm run lint
npm run test
npm run verify
```

`npm run verify` intentionally remains the fast lint + unit/component baseline.

Playwright is a **separate, on-demand browser automation capability**. It is not automatically required because a task reached Build, Code Review, Testing, or PR.

Available commands:

```bash
npm run browser:install
npm run test:browser
npm run test:browser:headed
npm run test:browser:ui
```

For a non-trivial proposed browser check/regression, the planning/verification contract should state:

```text
why browser automation is useful
risk if omitted
which phase owns the authoritative run
```

An agent may recommend Playwright based on concrete browser/interaction/regression risk. It becomes part of the task when the human explicitly requests it **or approves a task/spec/Techplan that includes that verification contract**. Do not require a second redundant permission prompt after the human has already approved the plan containing it.

Code Review may run a targeted browser reproduction when needed to prove/disprove a suspected finding; it should not rerun broad final verification by default. Build should use focused implementation feedback. Broad/final independent verification belongs to Testing when the workflow uses a Testing phase.

Committed tests under `tests/browser/` are automation assets scoped by the behavior they protect: smoke, a feature flow, a known regression, or a deliberate broader journey. Their scope is not derived from a Harscode task folder.

Playwright does not replace required human rendered acceptance.

Do not claim any verification that was not actually run.

## 12. Security presentation

- frontend role-aware rendering does not provide authorization security;
- backend authorization remains authoritative;
- public-facing fields must come from deliberate public-safe contracts, not raw internal models;
- never expose secrets, raw tokens, or PII through logs/debug UI;
- follow root `AGENTS.md` for security/fencing rules.

## 13. Workflow and Codex execution authority

Harscode owns the generic feature lifecycle and generic frontend engineering practice.

During frontend reboot preparation, do **not** manufacture Exploration/Techplan/Build artifacts for the reset itself. The reboot is project preparation/maintenance governed by `../docs/project/frontend-reboot-plan.md`.

After the clean reboot baseline is verified/frozen, when manually invoking a Harscode phase, start from the current canonical Harscode `workflow/*-prompt.md` entrypoint. Fill its normal variables/context. Do not create a parallel frontend copy of the Harscode lifecycle prompts or add solution-steering instructions that tell the agent what repository conclusions to discover.

`../docs/kencleng-agentic-workflow.md` owns Kencleng-specific MVP slice sequencing, risk/human authority, and integration states. Read the specific project-precondition/risk/integration section needed for the current decision; the full orchestration overlay is not a default per-phase read.

For Codex frontend work, `../docs/project/codex-frontend-execution-profile.md` owns project-specific model/reasoning/client/capability routing. Read the routing section(s) needed to choose or reconsider the current execution capability; do not repeatedly reload unrelated rendered-acceptance/Playwright rationale that is already enforced here.

Historical feature/task/local-agent docs may contain older workflow/design/domain-order assumptions. Treat the **current named rule/source owner** as authoritative rather than inferring policy from historical artifacts.

## 14. Source routing

```text
whole-product truth        → ../docs/product/README.md → ../docs/product/product-overview.md
current MVP boundary       → ../docs/product/mvp-scope.md
current MVP slice          → ../docs/product/mvp-delivery-slices.md
reconciled domain behavior → ../docs/spec/<domain-dir>/
API shape                  → ../api/README.md → ../api/openapi/<domain>.yaml + referenced common.yaml components
aggregate API view         → ../api/openapi.yaml only when cross-domain/generated-bundle context is needed
cross-stack dependency map → ../docs/project/kencleng-integration-map.md when a surface crosses/depends on backend capabilities
frontend architecture      → ../docs/project/kencleng-frontend-tech-stack.md (active concern sections)
frontend reboot gate       → ../docs/project/frontend-reboot-plan.md (pre-development only)
Codex execution profile    → ../docs/project/codex-frontend-execution-profile.md (current routing concern)
UI/UX authority map        → ../docs/ui-ux/README.md
product-design authority   → ../docs/ui-ux/product-design-principles.md
brand/product UI direction → ../docs/ui-ux/brand-product-ui-brief.md
visual system              → ../docs/ui-ux/design-guidelines.md
UX behavior                → ../docs/ui-ux/patterns.md (matching pattern)
asset governance           → ../docs/ui-ux/asset-governance.md
route/persona inventory    → ../docs/ui-ux/page-map.md (matching route/persona)
selected visual evidence   → ../docs/ui-ux/visual-references/selected-direction/
component contracts        → components/README.md, when broad contracts exist
project status             → ../docs/project/kencleng-development-tracker.md
```

This is a clue map, not a checklist of documents to preload. If authorities genuinely conflict on the same concern, surface the contradiction.

## 15. Output style

Default narration is terse. Final deliverables remain complete where Harscode requires complete risk notes, review findings, testing/build reports, or PR descriptions.