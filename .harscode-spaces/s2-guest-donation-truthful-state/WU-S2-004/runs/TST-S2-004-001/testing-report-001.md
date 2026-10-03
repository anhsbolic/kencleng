Phase: Testing
Author: P-S2-004-TST-001-1 (Verifier / KC-VERIFIER)
Created: 2026-10-03
Model / Reasoning: Invocation-configured gpt-6-luna / medium; active runtime values not independently exposed
Session: Not independently exposed
Target revision: Kencleng HEAD bb69cd002b3f1a1056837affcd77bb2b001007b0 plus current working tree; pinned 19-file frontend set
Workflow revision: Harscode pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4 (ordinary guidance current-effective)
Work Unit / Run: WU-S2-004 / TST-S2-004-001

# Independent Testing Report — TST-S2-004-001

## 0. Sweep Summary

- **Source integrity:** all 19 SHA-256 values in this Run's Invocation matched current files before execution. Testing proceeded against that exact set. The worktree has unrelated changes outside those pinned paths; this report makes no claim about them.
- **Confirmed:** the two Build reports' focused assertions were rerun as part of the final `npm run verify`: 5 test files and 31 tests passed, including the current Campaign Detail and guest Donation flow assertions. The patch report's 15-test claim was not treated as independent evidence until this run.
- **Closed from prior deferred list:** current patched-state lint/unit verification, current production compilation/route generation, and the approved R8 Chromium fragment-to-header check all ran successfully.
- **Still requires fresh or external evidence:** R10 Human rendered acceptance; Integration Map Donation-flow mapping; real backend/security/runtime and residual-risk evidence. R3, R5 and some R8 negative-edge assertions are noted below as non-blocking coverage follow-ups. No code defect was observed, so no patch plan was created.

## 0a. Test Focus Pointer Execution

| Area | Evidence anchor opened | Specialized verification | Result |
|---|---|---|---|
| Guest bearer status credential | `EXP-S2-004-001/evidence/stage-2-gap-analysis.md` §Area 2 and §Area 4; `stage-3-solutioning.md` “Guest status credential and proof boundary” | Techplan-approved Playwright Chromium test against the real frontend route; synthetic credential handoff, URL cleanup, status request header, status-only rendered view | Passed for frontend handoff. Does not establish infrastructure exposure, credential strength/expiry, or server parity. |
| Optional email and PII boundary | `stage-2-gap-analysis.md` §Area 1; `stage-3-solutioning.md` “Optional email” | `npm run verify` reran opt-in payload/disclosure and email error association assertions; inspected current UI mapping | Passed for frontend opt-in/presentation. Email verification, delivery, retention and deletion controls remain deferred. |
| Idempotency, stale Campaign action, cap and capacity no-fit | `stage-2-gap-analysis.md` §Area 1 and §Area 4 | `npm run verify` reran observable stale action, cap disclosure, 422/409 distinction, and ambiguous same-key/same-payload retry cases | Covered frontend behavior. Backend ordering, persistence, concurrency and exact-wire behavior are not established. See R3/R5 coverage follow-up. |

No specialized concurrency/load/security-control execution was warranted or authorized for frontend-only mock evidence; backend and Security owners retain those obligations. The pointer contains anchors for each still-relevant sensitive area; no missing-pointer drift was found.

## 1. Test Coverage

| Rule / scenario | Category | Observable verification | Result |
|---|---|---|---|
| R1 | Entry action / stale state | `npm run verify`: Campaign available action links to the flow; unavailable action exposes no link; stale donation route blocks form | Pass |
| R2 | Cap / currency truth | `npm run verify`: IDR cap appears on regular and Funding-unavailable Campaign details before entry; cap explanation appears on donation form | Pass |
| R3 | Amount boundaries | `npm run verify`: Rp4.999 is rejected; Rp5.001 is accepted and sent as whole-IDR string. UI schema/code inspected for integer-only rule and absence of a client cap. | Partial: fractional/malformed inputs and exact Rp5.000 are not directly asserted in current source tests despite the checklist's table-driven plan. No implementation defect observed. |
| R4 | Server errors | `npm run verify`: generic amount 422 is field-level and server detail is not shown; closed/ineligible 409 is request-level; code preserves input on recoverable paths | Pass for asserted cases |
| R5 | Idempotency / duplicate suppression | `npm run verify`: ambiguous retry preserves exact key and payload. Current implementation inspected: submit guard and locked controls suppress concurrent activation; accepted mock detects same-key changed-payload conflict; a new submit creates a UUID. | Partial: duplicate activation, changed-payload conflict, and new-key-after-known-failure lack explicit assertions in current source tests. Backend durability/concurrency remains unproven. |
| R6 | Method and simulation truth | `npm run verify`: QRIS submission, other three methods unavailable, pending and terminal labels, no automatic polling path in current component | Pass for covered frontend states; Human judges rendered comprehension separately |
| R7 | Optional name/email | `npm run verify`: opt-in payload and exact helper, optional fields omitted by default, invalid email linked to field and stale association cleared | Pass for frontend behavior; no claim of verification or delivery |
| R8 | Credential handoff / status-only | Playwright Chromium 1 test: navigated to real Next route with fragment; observed fragment removed, exact synthetic credential on `X-Donation-Status-Credential`, pending status-only heading and no credential/Donation ID in rendered text. RTL also covers status-only and missing-credential generic state. | Pass for frontend happy handoff. Wrong/expired/absent lookups are not each separately asserted in source tests; all non-OK responses map to the same generic UI in implementation. No backend parity inferred. |
| R9 | Safe errors / recovery | `npm run verify`: 422, 409, transport ambiguity, retry, missing status credential, and unavailable result render safe distinct or generic text as applicable | Pass for asserted categories; no internal error data observed in UI assertions |
| R10 | Accessibility / responsive / keyboard | `npm run verify`: accessible names, email `aria-invalid`/`aria-describedby`, persistent status live region, headings/actions and existing focus behavior assertions. | Automated semantic subset passed. **Human rendered desktop/mobile and keyboard acceptance remains pending and is not passed here.** |
| R11 | MSW boundary / generated contract | `npm run verify`: handlers and flow assertions exercise generated request/response-shaped fixtures; inspected production API adapter still uses `/api/...` real endpoint paths and has no environment-selected mock response branch | Pass as frontend mock evidence only |

## 2. Error Verification

| Error case | Expected behavior/category | Actual | Actionable/propagated correctly? |
|---|---|---|---|
| Local amount below minimum | Field validation; do not submit | Rp4.999 receives minimum field message; request not sent | Yes |
| POST 422 amount | Safe amount-level feedback, no hidden capacity/close detail | Generic field message; server-supplied detail omitted | Yes |
| POST 409 closed/ineligible | Request-level outcome distinct from amount error | Safe Campaign request alert; no amount message | Yes |
| Ambiguous submission transport error | Preserve intent; retry same key and payload | Retry action offered; assertions confirm key/body stability | Yes, within frontend boundary |
| Missing status credential / unavailable status response | Generic status-link failure | Generic link message; no status details | Yes for exercised missing/unavailable cases |
| Malformed API JSON, credential parity/body/header/cache/timing | Backend/public error contract | Not exercised end-to-end; client maps non-OK/invalid status to generic unavailable; server parity is deferred | No parity conclusion |

## 3. Final Verification

- **Target repo required final commands:** `cd frontend && npm run verify` — exit 0; ESLint 0 errors and 1 existing warning for unused disable directive in generated `frontend/public/mockServiceWorker.js`; Vitest 5 files / 31 tests passed. `cd frontend && npm run build` — exit 0; Next.js compilation and TypeScript passed, static pages generated, routes include `/campaigns/[campaignId]/donate` and `/donations/[donationId]/status`.
- **Playwright R8:** `npm run test:browser -- --config ../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/TST-S2-004-001/playwright.config.ts` — 1 Chromium test passed. The first sandbox launch could not bind the local server; after authorized local server startup outside sandbox, the Chromium run passed. Temporary config/spec are in this Run's `playwright.config.ts` and `browser/`; output is in `browser-output/`.
- **Broad checks intentionally not rerun:** no backend, race/concurrency, load/performance, security-control, or real integration checks; these belong to WU-S2-003 and API/Security/PII/runtime owners and cannot be established by frontend mocks. No additional whole-repository suite beyond the target `verify` command was required for this frontend-scoped change.
- **Migration/schema collision:** N/A — frontend-only change; no schema or migration authored.
- **Backward compatibility:** frontend route and API compatibility evidenced by current generated types, MSW assertions, build route validation, and request-shaped tests. No live backend consumer/wire compatibility claim; exact wire remains WU-S2-003.
- **Broader-suite requirement for cross-cutting change:** N/A — implementation is route-local frontend scope; final frontend baseline ran.
- **Fresh Techplan consistency read:** reread TP-S2-004-003 §§1–13, including all R1–R11, backward compatibility, risk/open items, interface, checklist and Test Focus Pointer. No contradiction was found. The known planned coverage gaps for edge-specific R3/R5 and separate Human R10 are surfaced above rather than hidden.

## 4. New Recurring Bug Patterns

None. The test coverage follow-ups are specific to this Work Unit and do not establish a reusable defect category.

## Verdict

**Pass with flagged follow-ups.** Independent target verification and the planned R8 frontend browser handoff passed; no code finding was identified. Follow up the missing explicit R3/R5 edge assertions and R8 wrong/expired/absent negative cases in the owning verification/test work before treating those exact scenarios as independently asserted. These are coverage gaps, not regressions of the passing tests. Human R10 acceptance and all cross-stack gates remain open; this verdict is not `FRONTEND_MOCK_VERIFIED`, runtime readiness, residual-risk acceptance, or Work Unit completion.

## Phase handoff

- **Outcome:** COMPLETED — independent frontend Testing execution ended with Pass with flagged follow-ups; occurrence outcome only.
- **Result refs:** `testing-report-001.md`; `launch-record.md`; approved `TP-S2-004-003/techplan.md` SHA-256 `e6a93bc65a760084955f63a469b238fd70a80c20ff38cc1d22ba1be000ad99eb`; this Run's temporary R8 harness and browser output.
- **Findings:** no code finding; R3/R5/R8 assertion coverage follow-ups are recorded in §1.
- **Decision requests:** none required to close this Run. Human rendered R10 acceptance remains a delivery gate.
- **Blockers:** none for this Run's bounded frontend verification. Downstream gates remain independently required.
- **Open / unverified:** R10 Human representative desktop/mobile rendered acceptance and keyboard exercise; Integration Map Donation-flow row; WU-S2-003 exact Campaign response DTO/wire; backend eligibility/capacity enforcement; persisted idempotency and concurrency; simulator/runtime behavior; credential generation/key/strength/expiry and browser/infrastructure exposure controls; anti-enumeration response/body/header/cache/timing parity and abuse controls; email verification/delivery/retention/deletion-race controls; real backend integration and residual-risk acceptance. No settlement or real payment rail evidence.
- **Recommended continuation:** Orchestrator to reconcile the Integration Map Donation-flow mapping and route Human R10 acceptance plus backend/Security/API/runtime gates to their owners. Keep flagged R3/R5/R8 assertions visible for the appropriate test maintenance/re-entry; do not route to Build for a code patch absent a code finding.
- **Context refs:** this Run's Invocation, approved Techplan §§4, 12–13, Build/Patch reports, latest Review confirmation, `docs/project/kencleng-integration-map.md`, and WU-S2-003 / Security/PII/API ownership.
