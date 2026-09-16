# Kencleng — Frontend Reboot Plan

> Status: Active preparation
> Decision owner: Anhar Solehudin
> Prepared with: ChatGPT — GPT-5.6 Sol
> Prepared: 2026-09-16
> Last updated: 2026-09-16
> Scope: project preparation before the next frontend Harscode development run
> Base design authority: `docs/ui-ux/README.md`

## 1. Purpose

Kencleng is intentionally retiring the current frontend product implementation before the next frontend development cycle.

The goal is **not** to delete the frontend repository or discard proven engineering tooling. The goal is to return the active frontend tree to a clean engineering scaffold so the next product UI implementation is derived from the current canonical product/design authorities rather than inherited accidentally from the superseded frontend generation.

This preparation work is **not a Harscode feature-development run**. Exploration, Techplan, Build, Code Review, and Testing for the new frontend begin only after the reboot baseline is frozen and declared ready for development.

## 2. Human decisions already made

- The approved frontend design direction is **Sunlit Editorial / Evidence-Led Optimism**.
- The next frontend implementation should be allowed to start from a clean presentation foundation rather than preserving the old UI by default.
- Existing product/UI implementation has no preservation privilege merely because it already exists.
- Git history remains the archive for the retired implementation and removed design generations; repository history will not be rewritten.
- `frontend/.local-agents/` remains intentionally tracked because Kencleng is a learning-by-doing project and process evidence should be inspectable by readers.
- Existing `.local-agents` work artifacts tied to the retired frontend generation may be removed from the active tree during the reboot; their history remains available in Git. New development runs will create a fresh committed process history.

## 3. Reboot principle

```text
canonical product/domain truth
+
canonical Sunlit Editorial design authority
+
clean engineering scaffold
        ↓
new frontend development
```

The reboot must avoid both failure modes:

1. preserving stale UI/component decisions because they already exist; and
2. deleting infrastructure that is valuable independently of the retired product implementation.

## 4. Retain by default

Retain project capabilities that are implementation-independent and remain intentionally selected for the next frontend generation:

- Next.js App Router / React / TypeScript project scaffold;
- package manager lockfile and build/lint configuration where still valid;
- Tailwind CSS v4 capability, but not the old visual tokens;
- Vitest + React Testing Library capability;
- Playwright capability;
- OpenAPI TypeScript generation/tooling;
- MSW capability, but not old feature fixtures/handlers merely because they exist;
- Harscode/Kencleng agent guidance;
- canonical product/domain specs and API contracts;
- canonical `docs/ui-ux/` authorities and approved selected-direction references;
- repository-level development tooling that does not encode the retired frontend product implementation.

Retention is about capability, not preservation of old product code.

## 5. Retire/reset by default

The active tree should not preserve implementation whose primary value comes from the retired frontend generation. The deletion manifest should evaluate and normally retire:

- existing public/authenticated route UI and layouts;
- existing product components;
- current `components/ui/` and `components/shared/` implementations and their old contract registry;
- old visual tokens, colors, fonts, radius/elevation values, and presentation CSS;
- Lucide-specific product/icon presentation where it represents the retired visual system;
- feature hooks/stores/API wrappers created only for the retired frontend implementation, unless a separate project authority requires retention;
- old frontend feature mocks/fixtures;
- old component/UI tests whose subject is being retired;
- old browser scenarios whose subject is being retired;
- legacy frontend assets and presentation-specific PWA shell artifacts when no longer justified;
- tracked `frontend/.local-agents/works/**` artifacts from the retired frontend generation.

Git history remains the archive for all retired files.

## 6. Clean-start component rule

The reboot must not pre-create abstractions simply to reproduce the old directory shape.

After reset, it is acceptable for these directories not to exist until a real implementation creates a truthful owner:

```text
components/ui/
components/shared/
components/features/
lib/hooks/
lib/stores/
mocks/
tests/browser/
```

The architecture describes placement rules, not a requirement to keep empty or speculative layers.

The living component registry should start with no established production `ui`/`shared` contracts. New reusable contracts are added only when the new implementation creates and validates them.

## 7. Local-agent process-history policy

`frontend/.local-agents/` remains tracked intentionally.

Purpose:

- make the engineering learning process inspectable;
- preserve representative Harscode artifacts for education and retrospective analysis;
- allow readers to see how requirements became plans, code, review findings, patches, and verification evidence.

Rules after reboot:

- old active work directories tied to the retired frontend generation are removed from the active tree;
- Git history is the archive for those retired artifacts;
- new runs use fresh task directories;
- generated workflow artifacts remain task evidence/history and do not become project-wide authority merely because they are committed;
- project truth must still be promoted into the owning spec/architecture/design document.

## 8. Documentation reconciliation required before deletion

Before deleting product implementation, reconcile at least:

- `docs/spec/0-foundations/tasks.md`;
- `docs/spec/0-foundations/features/01-frontend-experience-foundation.md`;
- `docs/project/kencleng-frontend-tech-stack.md`;
- `docs/project/kencleng-development-tracker.md`;
- `frontend/AGENTS.md` where current-runtime assumptions need removal;
- `frontend/README.md`;
- `frontend/components/README.md`;
- `docs/ui-ux/asset-governance.md`;
- `docs/ui-ux/visual-references/selected-direction/README.md`.

The active documentation must not point new agents back toward removed prototype authority, green-brand rules, retired frontend components, or old workflow artifacts as current precedent.

## 9. Deletion-manifest requirement

Do not perform a broad `rm` based only on top-level directory names.

Before deletion, produce an explicit manifest classifying each active frontend area as:

```text
RETAIN
RESET / REPLACE WITH MINIMAL SCAFFOLD
DELETE
REVIEW MANUALLY
```

The manifest must explain the reason for every retained implementation-level area so retention does not happen by inertia.

## 10. Minimal reboot baseline

The target active frontend tree should contain only enough code/configuration to prove the engineering scaffold remains healthy.

A representative target shape is:

```text
frontend/
├── app/
│   ├── layout.tsx          # minimal technical root
│   ├── page.tsx            # minimal bootstrap surface, not product precedent
│   └── globals.css         # minimal bootstrap only; no legacy visual system
├── components/
│   └── README.md           # governance; no established component registry entries
├── public/                 # only still-justified technical/static essentials
├── .local-agents/          # committed process evidence; fresh works begin later
├── AGENTS.md
├── README.md
├── package.json
├── package-lock.json
├── tsconfig.json
├── next.config.*
├── postcss.config.*
├── eslint.config.*
├── vitest.config.*
└── playwright.config.ts
```

Exact minimal files follow live scaffold/tooling requirements; this diagram is a target posture, not permission to invent unused folders.

## 11. Reboot verification

The reboot itself verifies **scaffold health**, not product behavior.

Minimum checks after deletion/reset should prove, as applicable:

- dependencies/configuration are internally consistent;
- lint/static configuration loads;
- the minimal Next.js application builds;
- the minimal application can start/render;
- test runners/configuration can initialize even if no product tests remain;
- Playwright capability remains configured if intentionally retained;
- generated OpenAPI tooling still works if part of the retained scaffold.

Do not carry old product tests merely to make a test count non-zero.

## 12. Ready-for-development gate

The frontend reboot baseline is ready only when all of the following are true:

- [ ] Sunlit Editorial authorities are canonical and internally reconciled.
- [ ] Frontend architecture docs describe the new clean-start posture.
- [ ] Foundation spec assumes a clean frontend implementation baseline.
- [ ] Development tracker records the old frontend generation as intentionally retired.
- [ ] No active docs route agents to removed prototype/design authority.
- [ ] No legacy product UI implementation remains in the active frontend tree unless explicitly retained by the deletion manifest.
- [ ] No legacy reusable UI/shared contract is treated as established by default.
- [ ] Old feature-specific frontend tests/mocks/browser scenarios are retired with their subjects.
- [ ] Old `.local-agents/works/**` implementation history is removed from the active tree; Git remains the archive.
- [ ] `.local-agents/` remains available for fresh committed learning/process evidence.
- [ ] Retained engineering scaffold builds/boots at the agreed minimum level.
- [ ] The first new frontend development task/spec is clear and points to current authorities.
- [ ] The Harscode candidate baseline for the first new run is frozen.

Only after this gate passes should the next real frontend Harscode run begin.

## 13. Intended next development sequence

After the reboot baseline is frozen:

```text
clean frontend scaffold
→ Frontend Experience Foundation on representative `/` slice
→ human brand/experience acceptance
→ subsequent dependency-driven page/flow delivery
```

The representative `/` slice remains deliberately smaller than a complete landing page. It exists to establish the first production expression of the approved design system, not to recreate the retired landing implementation.