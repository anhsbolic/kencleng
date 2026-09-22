# TP-002 — Corrective Techplan Amendment for Frontend TypeScript Verification

> Work Unit         : `WU-S1-002`
> Run               : `TP-002`
> Phase             : Techplan corrective re-entry
> Role              : Planner
> Participant       : Codex CLI agent
> Session           : Fresh session
> Created           : 2026-09-22
> Status            : Current-effective mechanical/non-material amendment
> Source Techplan   : `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-001/techplan.md`
> Build Finding     : `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/BLD-001/report.md#flagged-for-techplan--testing`

## Finding and live evidence

`BLD-001` completed the approved reconciliation edits, but the approved Build check `cd frontend && ./node_modules/.bin/tsc --noEmit` failed only in the pre-existing `frontend/app/page.test.tsx`: `describe`, `it`, and `expect` were unresolved.

Direct checks in this Run establish that:

- `frontend/tsconfig.json` includes `**/*.ts` and `**/*.tsx`, therefore includes `app/page.test.tsx`.
- `frontend/vitest.config.ts` already has `test.globals: true`.
- Installed `vitest@4.1.11` exposes `vitest/globals`; its declaration file supplies `describe`, `it`, and `expect`.
- The unchanged approved command fails; the same TypeScript program succeeds with `./node_modules/.bin/tsc --noEmit --types vitest/globals`.
- `git log` shows `frontend/tsconfig.json`, `frontend/vitest.config.ts`, and `frontend/app/page.test.tsx` last changed together before this Work Unit, on 2026-09-16. `BLD-001` did not change them.

## Classification

**Mechanical/non-material amendment.** The corrective change makes the pre-approved verification executable. It does not change product/domain scope, authority or security boundaries, architecture/ownership, API/interface/data semantics, risk acceptance, production behavior, test behavior, or verification strategy. The approved `tsc --noEmit` command remains mandatory.

No independent Techplan re-review or renewed Human approval is required under the materiality rule in this Run invocation. The approved `TP-001` history remains intact; its amendment notice, implementation anchor, affected-file boundary, Build checklist entry, and resolved-item history now point to this amendment.

## Current-effective authorization

For a narrow new Build/Patch Run only:

1. Modify `frontend/tsconfig.json` at `compilerOptions` to add exactly `"types": ["vitest/globals"]`.
2. Do not modify `frontend/app/page.test.tsx`, `frontend/vitest.config.ts`, production frontend code, API artifacts, tests' runtime semantics, or any other Work Unit boundary to resolve this Finding.
3. Re-run the unchanged approved Build checks relevant to the correction:
   - `cd frontend && ./node_modules/.bin/tsc --noEmit`
   - `cd frontend && npm run lint`
   - `git diff --check`
4. Preserve `BLD-001` evidence. Do not claim `CONTRACT_READY` or update the tracker until the complete current-effective Techplan evidence is satisfied and the existing Human-owned milestone check is performed.

The source Techplan at `TP-001/techplan.md`, including its dated `TP-002` amendment notice and targeted section updates, is the current-effective executable spine. This file is the traceable corrective rationale; it does not replace the Techplan.

## Handoff

- Completed: finding re-grounded; amendment classified and recorded; current-effective Techplan narrowly updated.
- Artifacts: this amendment and `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TP-001/techplan.md`.
- Human decision: none required for this mechanical/non-material amendment.
- Recommended next step: dispatch a fresh, narrow Build/Patch Run limited to `frontend/tsconfig.json` and the checks above.
- Independent Techplan review: Skip — no material contract or verification-strategy change.
- Decomposition: Skip — one configuration-line correction with direct executable evidence.
