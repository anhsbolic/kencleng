# Patch Plan — `RV-S2-002-007`

> Phase: Code Review patch plan  
> Author: Codex Reviewer  
> Created: 2026-09-30  
> Model / reasoning: `gpt-6-luna` / `high` (dispatch configuration; runtime selection not independently exposed)  
> Session: Fresh independent Reviewer Session; Session ID not exposed  
> Target revision: `6891341a050982e14174ab5af132a200f24e71d9` plus current Task 01 source diff and Run artifacts  
> Workflow revision: `b179360088fbf8a7dec6c285cd8c0b8992cbb4aa`

## Blocking finding F-01

- **Location:** `docs/spec/4-campaign/features/09-closure.md:22-26`
- **Problem:** The Summary states all three Campaign close triggers share a `WHERE status = 'published'` idempotency guard. This prescribes a concurrency mechanism, while D1, `INV-campaign-13`, the Approved Techplan, and accepted Task 01 leave the mechanism unselected. The cross-reference in this feature explicitly limits its current reconciliation to the narrow D1 boundary.
- **Impact:** The historical mechanism claim can be mistaken for approved delivery guidance and used to assert ordering/correctness without owner approval or evidence.
- **Suggested resolution:** Remove the claim that all triggers share that SQL guard. Preserve the trigger names as applicable feature context, but clarify that the mechanism for close ordering is not selected in this narrow reference and point to `INV-campaign-13` for D1. Do not introduce a replacement locking/isolation mechanism or expand this task into broader Slice 3 closure reconciliation.
- **Classification:** Blocking.

## Build scope and verification

Make only the minimal correction needed to resolve F-01. Do not change the approved D1 behavior, broader closure lifecycle, Product/MVP, API, implementation, tests, or orchestration projections. Re-read the resulting Summary against `INV-campaign-13` and the accepted Task 01; no broad tests are indicated for this documentation correction. Keep the affected spec `draft` until applicable owner/Human review is recorded.
