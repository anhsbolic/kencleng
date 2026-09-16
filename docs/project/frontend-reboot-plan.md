# Kencleng — Frontend Reboot Plan

> Status: Ready for merge
> Decision owner: Anhar Solehudin
> Prepared with: ChatGPT — GPT-5.6 Sol
> Prepared: 2026-09-16
> Last updated: 2026-09-16
> Scope: project preparation before the next frontend Harscode development run
> Base design authority: `docs/ui-ux/README.md`
> Harscode candidate baseline: `workflow-v2@4199c6db1b26ef1920ba670f222aff0c6d0f9e59`

## 1. Purpose

Kencleng intentionally retires the previous frontend product implementation before the next frontend development cycle.

The goal is not to delete the frontend repository or discard proven engineering tooling. The goal is to return the active frontend tree to a clean engineering scaffold so the next product UI implementation is derived from current canonical product/design authorities rather than inherited accidentally from the superseded frontend generation.

This reboot is project preparation/maintenance, not a Harscode feature-development run. Exploration, Techplan, Build, Code Review, and Testing for the new frontend begin only from the merged reboot baseline.

## 2. Human decisions

- The approved frontend design direction is **Sunlit Editorial / Evidence-Led Optimism**.
- The next frontend implementation starts from a clean presentation foundation rather than preserving the old UI by default.
- Existing product/UI implementation has no preservation privilege merely because it already exists.
- Git history remains the archive for retired implementation and removed design generations; repository history is not rewritten.
- `frontend/.local-agents/` remains intentionally tracked because Kencleng is a learning-by-doing project and process evidence should be inspectable by readers.
- Old `.local-agents/works/**` artifacts tied to the retired frontend generation are removed from the active tree; their history remains in Git. New development runs will create fresh committed process history.

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

The reboot avoids both failure modes:

1. preserving stale UI/component decisions because they already exist; and
2. deleting engineering capability that remains valuable independently of the retired product implementation.

## 4. Retained capability

Retained project capabilities include:

- Next.js App Router / React / TypeScript scaffold;
- package/build/lint configuration where still valid;
- Tailwind CSS v4 capability, but not the retired visual tokens;
- Vitest + React Testing Library capability;
- Playwright capability;
- OpenAPI TypeScript generation/tooling;
- MSW capability, but not retired feature fixtures/handlers;
- Harscode/Kencleng agent guidance;
- canonical product/domain specs and API contracts;
- canonical `docs/ui-ux/` authorities and selected-direction references;
- repository-level tooling that does not encode retired frontend product behavior.

Retention is about capability, not preservation of old product code.

## 5. Retired active implementation

The reboot removes from the active tree the previous generation's:

- public/authenticated route UI and layouts;
- product components and old reusable UI/shared implementations;
- old visual tokens, fonts, colors, radius/elevation values, and presentation CSS;
- Lucide-based retired presentation usage;
- feature hooks/stores/API wrappers tied to retired frontend implementation;
- feature mocks/fixtures;
- component/UI/browser tests whose subjects were retired;
- presentation-specific PWA/service-worker assets that were no longer justified;
- `frontend/.local-agents/works/**` history from the retired generation.

Git history remains the archive for all retired files.

## 6. Clean-start component rule

The reboot does not pre-create abstractions merely to reproduce the old directory shape.

These locations may remain absent until real implementation creates a truthful owner:

```text
components/ui/
components/shared/
components/features/
lib/hooks/
lib/stores/
mocks/
tests/browser/
```

The architecture defines placement rules, not a requirement to keep speculative folders. The living component registry begins with no established production `ui`/`shared` contracts.

## 7. Local-agent process-history policy

`frontend/.local-agents/` remains tracked intentionally to make the engineering learning process inspectable.

Rules after reboot:

- Git history archives the retired work directories;
- new runs use fresh task directories;
- generated workflow artifacts remain task evidence/history, not project-wide authority merely because they are committed;
- reusable project truth must still be promoted into the owning spec/architecture/design document.

## 8. Documentation reconciliation

The reboot reconciles:

- `docs/spec/0-foundations/tasks.md`;
- `docs/spec/0-foundations/features/01-frontend-experience-foundation.md`;
- `docs/project/kencleng-frontend-tech-stack.md`;
- `docs/project/kencleng-development-tracker.md`;
- `frontend/AGENTS.md`;
- `frontend/README.md`;
- `frontend/components/README.md`;
- `docs/ui-ux/asset-governance.md`;
- `docs/ui-ux/visual-references/selected-direction/README.md`;
- frontend harness/config references that still pointed to removed prototype-era authority.

Active documentation must not route new agents back toward removed prototype authority, green-brand rules, retired components, or old workflow artifacts as current precedent.

## 9. Deletion manifest

The executed reset boundary is documented in:

```text
docs/project/frontend-reboot-deletion-manifest.md
```

The manifest classifies active frontend areas as `RETAIN`, `RESET / REPLACE WITH MINIMAL SCAFFOLD`, `DELETE`, or `REVIEW MANUALLY` so retention does not happen by inertia.

## 10. Resulting minimal baseline

The active frontend tree now intentionally contains only a minimal technical application shell plus retained tooling/governance.

Representative shape:

```text
frontend/
├── app/
│   ├── layout.tsx
│   ├── page.tsx
│   └── globals.css
├── components/
│   └── README.md
├── .local-agents/
│   └── README.md
├── AGENTS.md
├── README.md
├── package.json
├── package-lock.json
├── tsconfig.json
├── next.config.ts
├── postcss.config.mjs
├── eslint.config.mjs
├── vitest.config.ts
└── playwright.config.ts
```

The bootstrap `/` surface is intentionally plain and is not product/design precedent.

## 11. Verification evidence

Reboot verification proves scaffold health, not product behavior.

On 2026-09-16, Anhar reported successful local execution on `frontend-reboot-preparation` of the requested reboot checks, including dependency installation/synchronization, `npm run verify`, `npm run build`, and boot/render inspection of the minimal frontend. ChatGPT did not execute those commands and records them as **operator-reported verification**.

The resulting `package-lock.json` was pushed in commit:

```text
6fd4663c4fe3e8d9d5e89f4f6a9a71c8d18a2ba2
```

GitHub inspection confirms that commit only synchronized `frontend/package-lock.json`, and the lockfile root dependency set now matches the cleaned `package.json`.

## 12. Ready-for-development gate

- [x] Sunlit Editorial authorities are canonical and internally reconciled.
- [x] Frontend architecture docs describe the new clean-start posture.
- [x] Foundation spec assumes a clean frontend implementation baseline.
- [x] Development tracker records the old frontend generation as intentionally retired.
- [x] No active docs intentionally route agents to removed prototype/design authority.
- [x] No legacy product UI implementation remains in the active frontend tree unless explicitly retained by the deletion manifest.
- [x] No legacy reusable UI/shared contract is treated as established by default.
- [x] Old feature-specific frontend tests/mocks/browser scenarios are retired with their subjects.
- [x] Old `.local-agents/works/**` implementation history is removed from the active tree; Git remains the archive.
- [x] `.local-agents/` remains available for fresh committed learning/process evidence.
- [x] Retained engineering scaffold was operator-verified at the agreed minimum level.
- [x] The first new frontend development task/spec is clear and points to current authorities.
- [x] Harscode candidate baseline for the first new run is frozen at `4199c6db1b26ef1920ba670f222aff0c6d0f9e59`.

**Gate result: READY FOR MERGE.**

The clean frontend becomes the new frozen project baseline only after this reboot branch is merged to `main`; the exact merged `main` commit then becomes the Kencleng baseline for the first new frontend run.

## 13. Intended next development sequence

After merge/freeze:

```text
clean frontend scaffold
→ Frontend Experience Foundation on representative `/` slice
→ human brand/experience acceptance
→ subsequent dependency-driven page/flow delivery
```

The representative `/` slice remains deliberately smaller than a complete landing page. It establishes the first production expression of the approved design system; it does not recreate the retired landing implementation.

CRTV measurement for the next run begins before its first Exploration session. The reboot preparation itself is not part of that feature benchmark.
