Phase: Testing
Author: P-S2-003-TST-001-1 (Verifier, KC-VERIFIER)
Created: 2026-10-03
Model / Reasoning / Session: Invocation configured `gpt-6-luna` / `high`; active runtime values and Session were not exposed.
Target revision: Kencleng HEAD `19d53315ac2847a03405339c0e86df8a97850761` plus current durable working-tree changes; all eight reviewed slice file hashes matched.
Workflow revision: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; ordinary applicable guidance current-effective.
Work Unit / Run: `WU-S2-003` / `TST-S2-003-001`.
Approved Techplan: `TP-S2-003-006`, SHA-256 `b71951c86b2f74f6c9f1a164bcf43447c58ee655874deaf09275e6ff3b2316f9`.
Prior evidence: Build `BLD-S2-003-001` report SHA-256 `6800f6433b32339cc9829aa97b2987e19d6bc2882cda16dcec51dd9237a4b32a`; Review `RV-S2-003-004` findings SHA-256 `4ecd9310a4d2c98acb157b5ddc4ab395796b5c076ad1c5ca27f145d198e20eef`.

## 0. Sweep Summary

- **Confirmed:** R2 cap representation checks → `TestMapMaxDonationAmount` in `backend/internal/domain/campaign/service_test.go`, independently invoked by the focused package command; command passed. Covers accepted bounds, whole-IDR normalization, fractional/missing/invalid/out-of-range persisted values failing closed. This is projection validation, not Donation submission validation.
- **Confirmed:** R3 public wire shape → `TestPublicCampaignDetailHandler_ExactClosedWireShape` asserts ten top-level keys and exactly `amount`/`currency_code` in `max_donation_amount`; passed with the focused HTTP package command.
- **Closed from prior gap:** funding-unavailable projection → `TestPublicCampaignDetailHandler_IncludesCapWhenFundingUnavailable`; passed.
- **Still requires fresh Testing:** PostgreSQL execution/backfill/round-trip and migration reversibility; the migration is unapplied and was only inspected read-only as assigned.
- **Still requires whole-WU Testing:** Donation submit/status, D1 ordering, capacity reservation, exact-once settlement, Organization eligibility writes/handlers, O3/O4/O5 security/runtime, Campaign create/PATCH, and full WU behavior. These are unimplemented, gated, or outside this Build slice; no pass is assigned to them.

## 0a. Test Focus Pointer Execution

| Area | Evidence anchor opened | Specialized verification | Result |
|---|---|---|---|
| Cap projection slice | No specialized Test Focus row applies to this public read-only projection; exact focused unit/HTTP coverage is specified by the Invocation. | Focused `go test` only. Concurrency/security rows concern unimplemented Donation, D1, status credential, and email paths. | Passed for bounded projection; no specialized run needed. |
| Campaign draft write authorization source | Techplan §12 pointer notes no Exploration anchor and identifies E9/E10/Open Item 7. | Not exercised; no create/PATCH handlers or eligibility source are in the reviewed slice. | Deferred; no newly inferred authority or write path tested. |

## 1. Test Coverage

| Rule / scenario | Category | Observable verification | Result |
|---|---|---|---|
| R1 — slice boundary | Scope | Compared Build report and eight-file reviewed scope with this Run's assignment; this Run evaluates only the public cap projection and migration source. | In scope is bounded; whole WU remains incomplete. |
| R2 — exact whole-IDR cap projection | Boundary / invalid persisted data | `TestMapMaxDonationAmount`: Rp5.000, Rp5.001, Rp1.000.000.000, fractional, missing, invalid and below/above range. | Passed. Does not verify Donation request acceptance or database round-trip. |
| R3 — required closed cap wire shape | Compatibility / observable HTTP | `TestPublicCampaignDetailHandler_ExactClosedWireShape`: ten top-level keys, two nested cap members, decimal-string amount and `IDR`; compared with accepted API schema and Campaign detail feature. | Passed. |
| R3 — Funding unavailable | Edge | `TestPublicCampaignDetailHandler_IncludesCapWhenFundingUnavailable`. | Passed. |
| R14 — migration identity/schema collision | Migration inspection | Read migration listing and `000011` schema; `000012` has a unique version and adds a column absent from `000011`; up/down touch only the new cap column/constraint. No database state was changed. | No numbering or existing-column collision observed. Execution/reversibility not established. |

## 2. Error Verification

| Error case | Expected behavior/category | Actual | Actionable/propagated correctly? |
|---|---|---|---|
| Missing, malformed, fractional, or out-of-range stored cap | Fail closed; do not emit an invalid public cap. | Domain mapping tests assert an error for each invalid stored value; HTTP mapping is implemented through the existing public unavailable boundary. | Domain failure propagation is covered; no separate live database/error-path execution in this Run. |
| Donation request over-cap / capacity no-fit | Shared generic 422 on `amount`; capacity details hidden. | Donation submission path is not implemented in this Build slice. | Not tested; deferred to whole-WU Testing. |

## 3. Final Verification

- **Target repo required final build/lint/test commands:** Invocation-pinned `go test ./internal/domain/campaign ./internal/transport/http` from `backend/` → exit 0; both packages reported `ok (cached)`. Then `go test -count=1 ./internal/domain/campaign ./internal/transport/http` → exit 0; both packages passed with test cache disabled. The first uncached attempt in the sandbox was blocked before HTTP tests ran because loopback listener creation was denied; the same command passed after the requested escalation was approved. No broader command was required for this bounded read-projection slice.
- **Broad checks intentionally not rerun:** `make verify`, `go test ./...`, race, contract, security and PostgreSQL integration checks; the Invocation explicitly bounds verification to changed Campaign/HTTP packages, and this slice adds no concurrency-sensitive behavior. Whole-WU owners/gates remain responsible for their assigned evidence.
- **Migration/schema collision:** Version `000012` is unique after `000011`; the cap column is absent from `000011`; the down SQL removes the new constraint and column. Migration not applied or executed. Per routed `postgresql/migrations-safety.md`, the current single-step `NOT NULL DEFAULT` migration's locking impact and tested two-way reversibility are not established. Human-controlled migration lifecycle remains required; this report does not approve application.
- **Backward compatibility:** Source-level shape/default checked against accepted Campaign API and feature; test asserts the new required field and exact closed shape. Existing database row backfill is specified by the SQL default but not verified against PostgreSQL because the migration was not applied.
- **Broader-suite requirement for cross-cutting change:** N/A for this scoped read projection; no cross-domain transaction, concurrency, or public contract authority was changed in this Build slice.
- **Fresh Techplan consistency read:** Full current Techplan read completed. Bounded cap requirements, tests, and Build/Review scope align. Documentation drift: §13 Open Item 5 still calls the refreshed Techplan `Draft/In Review` and says independent Techplan Review remains active, while the current artifact header is `Approved` and this Run's Invocation names it as the approved spine. This does not change the cap-slice test result; reconcile the stale status text in the Techplan-owning workflow. The whole Techplan also remains incomplete, with Tier-0, Open Item 7, O3/O4/O5, migration, PostgreSQL, and Human gates active.

## 4. New Recurring Bug Patterns

None. The Techplan status mismatch is a task-specific documentation drift, not a reusable bug pattern.

## Verdict

**Pass with flagged follow-ups — public Campaign cap projection slice only.** The focused domain and HTTP evidence passes, including exact cap boundaries/fail-closed mapping, required closed wire shape, and Funding-unavailable presence. The unexecuted migration and PostgreSQL evidence are explicit follow-ups and are not covered by this verdict. This is not a verdict on WU-S2-003 or the whole approved Techplan; no `BACKEND_VERIFIED` claim is made.

Flagged, non-blocking for this bounded code-slice verdict:

1. Migration execution, existing-row backfill/round-trip, locking impact and up/down reversibility remain unverified; migration application stays Human-triggered.
2. Reconcile the stale `Draft/In Review` / independent Techplan Review wording in §13 against the current Approved Techplan state.
3. Whole-WU requirements and gates listed above remain open for a later whole-spine Testing Run.

## Phase handoff

- **Outcome:** `COMPLETED` — independent focused verification completed for the public Campaign cap projection slice.
- **Result refs:** `testing-report-1.md`; `launch-record.md`.
- **Findings:** Non-blocking Techplan status-text drift at §13 Open Item 5; see Final Verification. Migration execution and reversibility evidence remain unverified.
- **Decision requests:** None for this bounded slice. Migration application remains a Human-owned lifecycle action outside this Run.
- **Blockers:** None for the tested public projection. This result does not unblock whole-WU gates.
- **Open / unverified:** Migration/PostgreSQL evidence; whole-Techplan Donation/D1, create/PATCH and Organization-source behavior; O3/O4/O5 runtime/security evidence; Human acceptance and remaining whole-WU gates.
- **Recommended continuation:** Reconcile the stale Techplan status text through its owning workflow. Preserve the migration as unapplied pending its Human-controlled lifecycle and required migration evidence. Route remaining implementation and a fresh whole-WU Testing Run only after the approved spine's prerequisites are met.
- **Context refs:** Approved `runs/TP-S2-003-006/techplan.md` §§12–13; `runs/BLD-S2-003-001/report.md`; `runs/RV-S2-003-004/review-findings-1.md`; `docs/spec/4-campaign/features/02-campaign-detail-listing.md`; `api/openapi/campaign.yaml`; backend migration `000012` up/down.
