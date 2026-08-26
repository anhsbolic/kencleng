# Task 05 — Contract completion: 429 on reset-password + bundle regen

> Back-reference (contract): `../techplan.md` — sections 1–8 are the source of truth. Techplan wins over this file on any apparent conflict.
> Splitting axis: dependency/sequence chain (see `manifest.md`). NO dependency on Tasks 01–04 — independently executable and reviewable at any point (this independence is one of the two justifications for decomposing at all).

## Scope

**In scope:**
- One response addition to `/auth/reset-password` in `api/openapi/account.yaml`
- Regenerating `api/openapi.yaml` via the redocly bundler
- Validation (`npm run validate`) and bundle-diff review

**Out of scope (this task):**
- ANY other spec edit — notably NOT adding a documented 422 to `/auth/forgot-password` (explicitly rejected in Q3: behavior addition beyond precedent needs human), and NOT unifying problem-type URI prefixes (`errors/*` vs `problems/*`, deferred repo-wide per techplan §14 Active #1)
- Handler code (Task 04 owns keeping code matching spec)

## Dependencies

None. Can run first, last, or parallel with everything.

## Binding decision (techplan §5, Q2 — resolved by Anhar)

> **Q2 missing 429 on reset-password → ADD to spec.** The contract under-documents reachable middleware behavior: the mount-time `RateLimit` wrapper (main.go:172) wraps ALL `/auth/*` routes, so reset-password can absolutely return 429, but only forgot-password documents it (account.yaml:259–260). This is a spec COMPLETION documenting existing behavior — not a loosening, not a wire change.

Authority note (root AGENTS.md §4): this edit was approved by the human during the explore session — it is not an agent-side spec change. Record it as such in the commit message.

## Implementation details

**File**: `api/openapi/account.yaml`
- Inside `/auth/reset-password` → `post` → `responses` (currently lines 274–299), add:
```yaml
        "429":
          $ref: "./common.yaml#/components/responses/TooManyRequests"
```
matching exactly how forgot-password declares its 429 (account.yaml:259–260). Placement order should mirror the existing file's convention (forgot-password lists it after "202").

**File**: `api/openapi.yaml` (generated — NEVER hand-edit)
- Regenerate: `cd api && npm install && npm run bundle`
- Per api/README.md workflow: commit BOTH the split-source edit and the regenerated bundle in the same change.

## Spec-first drift discipline (best-practices lens: openapi-spec-first-drift)

Checklist before committing:
1. Spec updated alongside the endpoint work (this task may land before/with Task 04's handlers — either order is fine as long as both are in the merge window)
2. Bundle regenerated and committed, not stale
3. Error-response shape verified to match: the shared `TooManyRequests` response already exists in common.yaml:105–115 — no new component authored
4. Review the bundle diff for UNEXPECTED changes elsewhere — unexpected churn signals pre-existing drift; if the diff touches anything beyond reset-password's responses block, STOP and flag (guardrail §6 territory)

## Testing checklist (this task's items from techplan §12)

- [ ] `npm run validate` passes against the split source
- [ ] Bundle diff contains ONLY the reset-password 429 addition (+ any deterministic reflow from the bundler)
- [ ] Task 04's R6 rate-limit tests double as the behavioral proof that the documented status is real

## Common mistakes that apply here

| Mistake | Fix |
|---|---|
| Hand-editing `api/openapi.yaml` | Generated artifact — edit `openapi/account.yaml`, run bundler |
| Editing the bundle without the source, or committing one without the other | Both in the same commit (api/README discipline) |
| "While we're in there" spec edits (422 on forgot, URI prefix unification) | Explicitly out of scope; each needs its own human decision |
