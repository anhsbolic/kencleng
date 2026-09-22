# TST-001 — Testing Report 1

> Phase             : Testing
> Work Unit         : `WU-S1-002`
> Run               : `TST-001`
> Role              : Verifier
> Specialization    : Contract / security-boundary verification
> Participant       : Codex CLI agent
> Session           : Fresh independent session
> Created/Updated   : 2026-09-22
> Model             : `gpt-5.6-terra`
> Reasoning         : `high`
> Target revision   : `9f6df246a5ec99a22a68e462a8b46b08b0ff30c8`
> Workflow revision : not exposed by the TST-001 invocation

## 0. Sweep Summary

- **Confirmed:** `BLD-002` TypeScript-globals correction remains effective: the unchanged `./node_modules/.bin/tsc --noEmit` now passes, and `npm run lint` also passes.
- **Confirmed:** `CR-001-F01` remains closed. The dereferenced `content_url` pattern accepts the controlled same-origin path and rejects a direct absolute URL, missing `/api`, query, fragment, malformed UUID, and trailing path.
- **Confirmed:** split validation, bundle generation, frontend type generation, and independent bundle/type reproduction all pass. OpenAPI validation reports zero errors and 126 historical warnings, matching the post-reconciliation Build/Review baseline; no warning is emitted at a touched Slice-1 public-contract coordinate.
- **Closed from prior gap/deferred list:** no active Build/Test gap remains. `BLD-001`'s only flagged TypeScript verification gap was closed by `TP-002` / `BLD-002`.
- **Still requires fresh Testing:** runtime-only public eligibility, response/timing parity, controlled storage recheck/retraction, proxy cache-header preservation, storage `503`, decimal calculation, and safe frontend rendering. These are deliberately downstream evidence, not claims made by this reconciliation Work Unit.

## 0a. Test Focus Pointer Execution

| Area | Evidence anchor opened | Specialized verification | Result |
|---|---|---|---|
| Public projection regression | `EXP-001/evidence/solutioning.md#risks-and-required-carry-forward-evidence` item 1 | Parsed the dereferenced bundle and walked the public-schema graph from `PublicCampaignDetail`; asserted standalone exact allowlist, all 12 exact objects closed, no `Campaign`/`Organization` inheritance, and no permissive object. Checked generated declarations. | Pass at artifact boundary; runtime forbidden-field mapping is deferred. |
| Existence disclosure / optional auth variance | `EXP-001/evidence/solutioning.md#risks-and-required-carry-forward-evidence` item 2; `gap-analysis.md#area-4--existing-shared-api-contract` | Inspected both public operations for `security: []`, exactly `200`/`404`/`503`, and the shared `PublicCampaignNotFound` response; no public-operation `403`. | Pass at contract boundary; response-body/header/timing parity is deferred to backend Testing. |
| Media revocation / cache staleness | `EXP-001/evidence/solutioning.md#risks-and-required-carry-forward-evidence` items 3–4 | Asserted JPEG/PNG-only controlled content operation, no attachment-list GET, anchored same-origin `content_url`, and `Cache-Control: private, no-store` on each success/error contract response. | Pass at contract boundary; origin recheck, private bucket, retraction, and topology preservation are deferred. |
| Money truth | `EXP-001/evidence/solutioning.md#risks-and-required-carry-forward-evidence` item 5 | Checked decimal-string amount/percentage schemas, nullable non-computable percentage, IDR, uncapped `above_target`, and generated TypeScript string types. Value examples cover zero collected amount and over-target percentage; the decimal pattern accepts the documented large-value range. | Pass at contract boundary; decimal calculation/serialization is deferred. |
| Organizer-controlled content | `EXP-001/evidence/solutioning.md#risks-and-required-carry-forward-evidence` item 6 | Checked plain unformatted strings, `source: organizer`, exact closed text object, and absence of HTML/Markdown contract escape hatches. | Pass at contract boundary; hostile-string rendered behavior is deferred to frontend Testing. |

## 1. Test Coverage

| Rule / scenario | Category | Observable verification | Result |
|---|---|---|---|
| R1 | Public operation shape | Dereferenced `getPublicCampaignDetail` and `getPublicCampaignMediaContent` have `security: []`; detail has the sole `PublicCampaignDetail` success projection. Feature, invariant, and threat-model text state optional-auth invariance. | Pass — artifact/contract. |
| R2 | Anti-enumeration errors | Both operations expose only `200`/`404`/`503`; their `404` dereferences to the same local `PublicCampaignNotFound` Problem response and neither has `403`. | Pass — artifact/contract. |
| R3 | Closed allowlist | Independent parser assertions passed for all 12 required exact objects: `additionalProperties: false`, every declared property required, no `allOf`, no internal `Campaign`/`Organization` graph member, and no forbidden top-level field. Corresponding generated declarations have no permissive index signature. | Pass. |
| R4–R5 | Provenance and lifecycle | `PublicCampaignOrganizerText` is a plain string plus `organizer`; `PublicCampaignLifecycle` is closed with only `fundraising`, `published_at`, and `fundraising_ends_at`. | Pass. |
| R6 | Funding truth | `target_amount`, `collected_amount`, and nullable `percentage` are strings; no public funding number/float; relationship includes `above_target` and `not_computable`; zero/over-target examples and exact decimal patterns are present. | Pass. |
| R7 | No fake action | Generated `PublicCampaignDonationAction` has exactly unavailable `availability` and `donation_flow_not_available` `reason`, with no URL/action property. | Pass. |
| R8 | Media union truth | Discriminated union branches are closed; available has `minItems: 1`; absent/unavailable have `maxItems: 0`; the item allowlist is exact and includes nullable `caption`. | Pass. |
| R9–R10 | Controlled byte delivery and cache contract | Controlled operation has only JPEG/PNG binary success content, public-safe errors, anchored same-origin path validation, no attachment-list GET, and `Cache-Control: private, no-store` on `200`, `404`, and `503`. | Pass — artifact/contract. |
| R11 | Split-contract execution | `cd api && npm run validate` passed with zero errors / 126 historical warnings. `cd api && npm run bundle` passed and left the committed bundle reproducible. | Pass. |
| R12 | Generated correspondence | `cd frontend && npm run generate:api-types`, `./node_modules/.bin/tsc --noEmit`, and `npm run lint` passed. Independent `redocly bundle` and `openapi-typescript` generation to `/tmp`, followed by `cmp`, matched both committed artifacts; both operationIds occur in generated types. | Pass. |
| R13 | Delivery-authority consistency | Task 02/03, both active features, `INV-campaign-14`, threat model, backend architecture, OpenAPI, and integration map agree on public-only closed detail, controlled media, `DEFER` historical breadth, and no-store. The development tracker remains intentionally pre-Human-promotion, but its obsolete PR #27/Product Authority wording must be replaced during the Human-owned status update. | Pass with Human-status follow-up. |
| R14 | Milestone honesty | No backend/frontend/topology/integration runtime claim was found. The manifest/control surface remains in the dispatched Testing state; this Run does not mutate it. | Evidence is sufficient for the Human decision; Testing does not accept the milestone. |

## 2. Error Verification

| Error case | Expected behavior/category | Actual | Actionable/propagated correctly? |
|---|---|---|---|
| Absent, invalid, non-public Campaign; absent/non-member media | Identical `404` `PublicCampaignNotFound` Problem Details with no-store | Same reusable response contract is referenced by both public operations. | Yes at contract boundary; runtime byte/body/header and timing parity deferred. |
| Eligible detail/funding or media-byte dependency unavailable | `503` `PublicCampaignUnavailable` Problem Details with no-store | Same reusable response contract is referenced by both public operations. | Yes at contract boundary; dependency propagation deferred. |
| Direct-object or malformed `content_url` | Contract rejects non-controlled URL shapes | Anchored pattern negative assertions passed for all required variants. | Yes at schema boundary; runtime redirect prevention deferred. |

## 3. Final Verification

- Target-repo required checks passed:
  - `cd api && npm run validate` — zero errors; 126 documented historical warnings.
  - `cd api && npm run bundle` — pass.
  - `cd frontend && npm run generate:api-types` — pass.
  - `cd frontend && ./node_modules/.bin/tsc --noEmit` — pass.
  - `cd frontend && npm run lint` — pass.
  - Independent temporary `redocly bundle` + `openapi-typescript` + `cmp` against `api/openapi.yaml` and `frontend/lib/api/generated/openapi.ts` — both match.
  - `git diff --check` — pass; pre-report worktree was clean.
- Broad checks intentionally not rerun: browser/UI, backend, database, MinIO/proxy, and integration suites. No runtime implementation exists in this Work Unit, so these checks cannot exercise introduced behavior and are explicitly assigned downstream.
- Migration/schema collision: N/A — no migration, persistence schema, or runtime route was introduced.
- Backward compatibility: pass at contract authority boundary. The intentional paper-breaking public-detail/attachment-list reconciliation is explicitly accepted because no live Campaign route or consumer exists; historical listing/upload remains `DEFER`.
- Broader-suite requirement for cross-cutting change: N/A — this is generated-contract/document reconciliation only; the independent regeneration and TypeScript/lint checks are the target-repo mechanism selected by the Techplan.
- Fresh Techplan consistency read: no new material contract contradiction. `CR-001-C01` stale threat-model reference labels remain a known non-blocking documentation follow-up. The development tracker is deliberately not yet promoted and contains obsolete pre-promotion wording; the Human milestone action must correct it rather than treating it as current authority.

## 4. New Recurring Bug Patterns

None. The closed-object and `content_url`-constraint issues were already recorded and resolved through the existing planning/review loop; no new reusable defect category was found.

## Verdict

**Pass with flagged follow-ups.**

There is no Build/Patch finding. The reconciliation artifacts satisfy R1–R13 at the observable contract/document/generated-artifact boundary. Runtime-only evidence is intentionally deferred to downstream backend, frontend, topology, and integration Work Units; this is not a failure of the reconciliation Work Unit.

The evidence is sufficient for a Human to accept the `CONTRACT_READY` milestone and, in the same Human-owned status action, update the development tracker from its stale pre-promotion / `NEEDS_RECONCILIATION` wording while keeping `BACKEND_VERIFIED`, `FRONTEND_MOCK_VERIFIED`, and `INTEGRATED_VERIFIED` unclaimed. No `patch-plan-1.md` is written because no Build-owned change is required.

## Phase handoff

- **Completed:** independent contract, security-boundary, generated-artifact, and final target-repo verification for `WU-S1-002`.
- **Artifacts:** `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/TST-001/testing-report-1.md`.
- **Human decision:** accept or reject `CONTRACT_READY`; if accepted, perform the tracker/control-surface status transition and correct obsolete tracker Product Authority / PR #27 wording.
- **Open / deferred:** runtime projection/forbidden-field tests; authorization/body/header/timing parity; private-storage recheck/retraction; no-store proxy behavior; storage `503`; decimal computation; hostile plain-text rendering; `CR-001-C01` labels.
- **Recommended next step:** Human milestone decision. If accepted, route separate downstream backend, frontend, and topology Work Units; do not represent this contract evidence as runtime or integrated verification.
- **Session transition:** a fresh downstream planning/build session is appropriate for each separately scoped Work Unit after the Human gate; no return to Build/Patch is needed for this reconciliation Run.
- **Context pointers:** current-effective `TP-001/techplan.md`, `TP-002/amendment.md`, `BLD-003/patch-report-1.md`, `CR-002/review-findings-1.md`, and this report.
