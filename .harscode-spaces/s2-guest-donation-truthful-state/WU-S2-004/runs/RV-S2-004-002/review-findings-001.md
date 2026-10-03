Phase: Code Review
Author: P-S2-004-RV-002-1 (Reviewer / KC-REVIEWER)
Created: 2026-10-03
Model / Reasoning: Invocation-configured gpt-6-luna / high; active runtime values not independently exposed
Session: Not independently exposed
Target revision: bb69cd002b3f1a1056837affcd77bb2b001007b0 plus current working tree
Workflow revision: Harscode pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4 (ordinary applicable guidance current-effective)
Work Unit / Run: WU-S2-004 / RV-S2-004-002
Reviewed source set: 19 pinned frontend paths; every current per-file SHA-256 matched the Invocation list (aggregate recorded there: `5d85f4063e48c8e394fca3a5b0d9c2c519ee1457477e6d6742fe555b86dd2977`).

## 1. Safety

No findings.

The status fragment is removed from the visible URL before the credential is sent in `X-Donation-Status-Credential`; the client requests status with `cache: "no-store"`, renders only the status projection, and uses a generic unavailable message. Submission ambiguity retains the original payload and idempotency key for retry. The credential and email/runtime controls that cannot be established by this frontend diff remain explicitly deferred under Techplan O3–O5 and Testing R8.

No newly discovered specialized concurrency, performance, or security test obligation is absent from the Techplan Test Focus Pointer; no Techplan drift is raised.

## 2. Quality

No findings.

## 3. Stack-Specific Best Practices

### RV-S2-004-002-F01 — Email validation error is not programmatically associated with its field

- **Location:** `frontend/app/campaigns/[campaignId]/donate/donation-client.tsx:190-198`
- **Severity / disposition:** Medium — blocking; included in Request changes.
- **Problem:** The email input sets `aria-invalid` when validation fails, but the rendered field error has no ID and the input has no conditional `aria-describedby`. The message is visually adjacent but is not programmatically associated with the field.
- **Why it matters:** Screen-reader users can be focused on an invalid email field without receiving the validation explanation needed to correct it. This breaks the email opt-in form’s accessible recovery path.
- **Suggested resolution:** Give the email error a stable ID and set `aria-describedby` on the input while that error is present, preserving the existing `aria-invalid` behavior.
- **Best-practice source:** `../harscode-workspace/best-practices/react/accessibility-fundamentals.md`, checklist item on programmatic labels and field-error association with `aria-describedby` + `aria-invalid`.

### RV-S2-004-002-F02 — Status recheck result is not announced after loading

- **Location:** `frontend/app/donations/[donationId]/status/status-client.tsx:76-86, 91-106`
- **Severity / disposition:** Medium — blocking; included in Request changes.
- **Problem:** The initial loading state is a live `role="status"`, but it is replaced by a plain section when the request resolves. Later manual checks also update the status heading inside a section without live-region semantics. A screen-reader user can hear “Memeriksa…” and receive no announcement of pending, terminal, or unavailable results.
- **Why it matters:** Rechecking status is the only available progress action on a pending donation. Without an announced result, assistive-technology users may not know whether a check completed or whether the donation reached a terminal state.
- **Suggested resolution:** Keep a stable polite status/live region for asynchronous check results (including unavailable results), or otherwise announce the updated result without moving focus away from the recheck control.
- **Best-practice source:** `../harscode-workspace/best-practices/react/accessibility-fundamentals.md`, checklist item requiring explicit handling of async content appearing and accessible state communication.

## 4. Consistency

No additional findings. F01 and F02 also conflict with the approved Techplan R10 accessibility requirement and `docs/project/kencleng-frontend-tech-stack.md` §20, which requires semantic HTML, keyboard/focus behavior, and accessible labels/names. The target repo’s form guidance keeps field errors near the field; F01 concerns their missing programmatic association, covered specifically by the Stack-Specific finding above.

### Integration-map follow-up

`docs/project/kencleng-integration-map.md` §4 says to populate the map lazily when a frontend surface enters planning or implementation; its active mappings currently list Public Campaign Detail only. The missing Donation-flow mapping is therefore a valid cross-stack coordination-document follow-up for Orchestration before integrated delivery. It does not indicate a code/API contract inconsistency, does not change the frontend review scope, and is not a Review blocker. The map was read-only in this Run and was not edited.

## Verification executed during Review

- Source integrity only: `sha256sum` was run over the 19 Invocation-pinned frontend paths before review; all 19 file hashes matched the recorded per-file SHA-256 list. The reviewed working diff consisted of those 19 frontend paths.
- Runtime/tests/build/browser verification: none. Static inspection was sufficient to establish F01 and F02; no test/build matrix was repeated.

## Verdict

**Request changes.** Blocking findings: RV-S2-004-002-F01 and RV-S2-004-002-F02. Both are localized frontend accessibility fixes; no contract or product decision is needed. No `patch-plan` item concerns the integration-map follow-up.

## Phase handoff

- **Outcome:** COMPLETED — this Review Run completed all four passes and recorded its verdict.
- **Result refs:** `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/RV-S2-004-002/review-findings-001.md`; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/RV-S2-004-002/patch-plan-001.md`; approved Techplan `TP-S2-004-003/techplan.md` SHA-256 `e6a93bc65a760084955f63a469b238fd70a80c20ff38cc1d22ba1be000ad99eb`.
- **Findings:** Blocking F01 and F02 in this artifact. Integration-map Donation-flow row is a non-blocking Orchestration follow-up before integrated delivery.
- **Decision requests:** None; the fixes do not require a product/API decision.
- **Blockers:** WU-S2-004 frontend progression to Testing is blocked on F01 and F02 being addressed through a new Build/Patch Run. The integration-map follow-up does not block this scoped frontend patch.
- **Open / unverified:** Testing-owned R8 real-browser fragment/history/header verification; Human-owned R10 rendered acceptance; backend/API/Security/PII runtime and integration evidence; integration-map update; no runtime/browser/tests/build matrix was run in Review.
- **Recommended continuation:** Start a fresh Build/Patch Run/Participant/Session against `patch-plan-001.md` to fix F01 and F02. Then Orchestration should reconcile the new Build evidence and route independent Testing; preserve R8 and R10 as separate gates.
- **Context refs:** Approved `TP-S2-004-003/techplan.md` §3, §4, §12; this artifact F01/F02; the two cited source locations; `docs/project/kencleng-integration-map.md` §§4–5.
