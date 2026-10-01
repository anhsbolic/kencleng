# Task 02 — Rekonsiliasi authored Donation OpenAPI

> Phase: Techplan decomposition task — post-approval snapshot reconciliation  
> Author: Codex Planner  
> Created / Updated: 2026-10-01  
> Model / reasoning: Invocation configured `gpt-6-luna` / `high`; active runtime not independently exposed  
> Session: Fresh Planner Session for Run `TPD-S2-002-002`; Session ID not exposed  
> Target revision: `650e73c5d646c29c0ddf1931618f02685d15f7b7` plus current durable decisions/projections; reopen live sources  
> Workflow revision: `33b03a3f62cc3aacba6534b8a011465613c64b09`  
> Work Unit: `WU-S2-002`  
> Run: `TPD-S2-002-002`  
> Participant: `P-S2-002-PLD-002-1` (ephemeral for this Run)

## Purpose and scoped outcome

Reconcile the authored Donation OpenAPI contract for Slice 2 guest submission and status lookup against Product/MVP, Design, the reconciled Donation domain specifications, current live authorities, and Approved TP-015. Scope includes only the paths, fields, encoding, error/response, headers/cache, simulator-control contract, and shared components justified by those authorities. Preserve unresolved details and downstream evidence gates; the task does not claim `CONTRACT_READY`.

## Parent Techplan

- Authority: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md`
- Version/status: current-effective, Human-approved `Approved` Techplan (2026-10-01).
- TP-015 remains the complete authoritative spine; this task does not reinterpret it or make material decisions.

## Governing parent references

- Requirements: Q1–Q13, especially Q2/Q12 (major-unit decimal string plus explicit currency code; exact decimal end-to-end), Q6/Q13 (terminal notice lifecycle), and Q7 (status credential/failure contract).
- Rules: R1–R14, especially R3/R6–R9 and R12–R14.
- Decisions: D1–D19 and rejected alternative D16-alt. Preserve D1/D15 threshold ordering; D7 O4/O5 directions; D17 monetary representation; D18 bounded O8 clearance; and D19 terminal-notice direction.
- Risks: RISK-1–RISK-10, especially amount precision, credential exposure/parity, bounded O8 compatibility scope, and atomic funding correctness.
- Verification: applicable TP-015 §12 obligations. OpenAPI validation/generation does not prove runtime, concurrency, security/PII controls, response/timing parity, or rendered acceptance.
- Active Open Items: O1 only for concrete currency/range/fraction/scale parameters; O2 simulator timing/mechanism; O3 security and email-lifecycle controls/evidence; O4 security/API implementation controls, evidence, and residual-risk decision (the Human-selected credential direction is settled); O5 empirical parity/timing/abuse evidence (uniform `404` contract direction is settled); and independent Human acceptance of the current Donation specs. Resolved history: O1 shared representation, O8 replacement clearance within its bounded scope, O11/D19 terminal-notice direction, O4 fragment/handoff/HMAC direction, and O5 uniform `404`/cache direction. Do not call these settled directions open alternatives.

## Scope and live contract/spec anchors

- Hard prerequisite: current Task 01 output `01-donation-domain-spec-reconciliation.md`. Task 01 Build history includes `BLD-S2-002-001/report.md`, `BLD-S2-002-002/patch-report-1.md`; reviews are `RV-S2-002-007/review-findings.md` and `RV-S2-002-008/review-confirmation.md`. Keep this history visible as context, but do not treat it as Human acceptance of the current spec drafts. Human acceptance remains a separate parallel item and is not claimed complete.
- `api/openapi/donation.yaml` — primary authored Donation paths and schemas.
- `api/openapi/common.yaml` — modify only if a required shared component is justified by the reconciled contract.
- `api/openapi/index.yaml` — update mechanically only if a path is added or removed.
- `api/README.md` — split-source validation, bundle, and generated-client workflow.
- `docs/spec/5-donation/` plus any approved Campaign cross-reference — current domain behavior after Task 01; continue to read TP-015 and live Product/API authority independently.
- Authorities: `docs/product/mvp-scope.md`, `docs/product/mvp-delivery-slices.md` §§5–6, `docs/ui-ux/README.md`, `docs/project/kencleng-monetary-data-standard.md`, `api/README.md`, `AGENTS.md` §§1–3, and TP-015 §§3–13.

## Execution detail and boundaries

- Express every Donation monetary API/wire value as a major-unit decimal string with explicit currency code, with exact-decimal/no-float meaning. Input remains whole-IDR Rupiah (minimum 5000; one-Rupiah increments). Do not invent currency support, range, fraction rules, universal precision/scale, migration detail, or derived-value rounding.
- Preserve same-key/same-payload retry, same-key/different-payload rejection, deliberate new-key behavior, and separation of request idempotency from settlement replay. Translate retry/serialization details only to the extent supported by TP-015 and live authorities.
- Keep QRIS the only active sandbox method among the four display labels. No real payment rail, client-callable settlement, or donor-controlled terminal outcome. Backend simulator owns status.
- Preserve D1/D15: submit eligibility is ordered atomically against Campaign close; an accepted-while-eligible donation remains settleable at full amount; close-first rejects a new submission; successful settlement and funding increment commit atomically and exactly once; later settlement does not reopen Campaign or change its winning close reason, and funding may exceed the threshold. Do not select a DB/locking mechanism or import Slice 3.
- Carry settled O4 credential contract: fragment URL with frontend handoff and URL cleanup; one-way HMAC verifier; hard 24-hour expiry from issuance without extension; status-only projection. Authored API must express the selected flow and resolve any remaining carrier/header expression consistently with fragment handoff. Do not reopen these choices.
- Carry settled O5 contract: absent Donation, missing/wrong credential, and expired credential all produce one uniform public `404` with identical Problem Details body, headers, and cache behavior, including `Cache-Control: private, no-store`. Authored OpenAPI must specify exact Problem Details coordinates and the cache directive. Treat exact authored contract expression as distinct from downstream empirical evidence.
- Downstream only: Security/PII and API owners define/evidence browser history/referrer/log/cache protections, key handling/comparison, expiry/lifecycle enforcement, and abuse controls; Testing demonstrates actual body/header/cache equality, material timing behavior, and abuse controls. No implementation control or empirical parity is established here, and no residual risk is accepted.
- Carry D19/O11 terminal-notice eligibility through fulfillment and bounded/recoverable terminalization direction where represented by the API; do not select bound, architecture, timeout-as-`failed`, or deletion/race policy. O2 simulator specifics and O3 verification/delivery controls remain with owners.
- O8/D18 is clear only to replace historical submit/status operations in this repository/current Slice-2 scope because the Human/API owner confirmed no external distribution. Do not generalize to all possible consumers; keep the evidence and ordinary API review visible.
- Update authored split source only. Follow `api/README.md` for path registry, validation, bundle, and frontend types as relevant; never hand-edit aggregate/generated outputs.

## Hard dependency and external gates

- Hard dependency: Task 01 must be completed so OpenAPI follows reconciled delivery-domain behavior. Its current spec Human acceptance remains a separate independent gate and must not be claimed complete. Keep Task 01 Build/review history identified above; no Task 02 Build is implied by this snapshot reconciliation.
- O1 concrete amount parameters, O2 timing/mechanism, O3 lifecycle/security controls, O4 implementation controls/evidence and residual-risk gate, O5 empirical evidence, and any owner review remain with their assigned authorities. They block only dependent contract detail or later proof as TP-015 specifies; they do not reopen settled directions.
- O8 replacement permission is boundedly resolved; maintain evidence scope and do not add a new audit gate absent contradictory evidence.
- Do not claim `CONTRACT_READY` until TP-015 R14 and §9 requirements are met: relevant contract-time authority/owner acceptance, reconciled specs and authored API, and source validation. Runtime-only proof stays downstream.

## Verification before handoff

- Owner/Human review checks scope and Slice 3 boundary (R1); monetary representation and unresolved parameter boundaries (R2/R12); atomic state/funding (R3); simulator semantics (R4); email direction and O3/O11 lifecycle gates (R5/R13); exact O4/O5 credential and `404` contract/cache direction plus separate downstream proof (R6); D1/D15 translation (R7); request idempotency (R8); bounded O8 evidence (R9); exact O7 wording/guidance (R10–R11); and the `CONTRACT_READY` evidence boundary (R14).
- Run `cd api && npm run validate`; if authored sources change, follow `api/README.md` to regenerate the aggregate bundle and frontend types, then inspect authored source/index/bundle/generated consistency. Do not hand-edit generated output.
- Keep R3/R4/R5/R6/R7/R8 runtime proof for downstream Build/Testing. Schema validation is not proof of atomicity, concurrency, simulator behavior, PII lifecycle, or anti-enumeration.
- If TP-015 lacks a material contract/security/verification decision needed for scope, stop and report the exact gap; do not amend the spine or decide it in this task.

## NOT in this task

No Product/MVP or Design authority, TP-015, runtime code, tests, migrations, manual DB/index application, or orchestration projections. No hand edits to `api/openapi.yaml` or generated frontend types; no Tier-0 changes. Do not resolve Active O1–O5, accept residual risk, claim current spec acceptance, or declare `CONTRACT_READY` without the required gates/evidence.
