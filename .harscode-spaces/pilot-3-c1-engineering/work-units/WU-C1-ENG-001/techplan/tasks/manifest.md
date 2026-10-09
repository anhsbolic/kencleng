# C1 execution task manifest

> Phase: Techplan decomposition; Author: `PARTICIPANT-C1-ENG-DECOMPOSER-001`; Created/Updated: 2026-10-09. Model: `gpt-6-luna`; Reasoning: medium; Session: `SESSION-C1-ENG-DECOMPOSER-001` (Run binding). Target revision: `524ef600c7f71246af6b71671d89c3c040fd9d44`; workflow revision: `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`.

## Parent

`../techplan.md` — SHA-256 `bd732748b486f0327f6cba2cdf14cab32f1ece3c5be8ebb24f9659444947fa47`, Approved. Parent remains the sole cross-task source for scope, rules, decisions, risks, contracts, verification and Open Items.

## Gate and axis

Step 0: **YES**. Independent review/context boundaries span shared contract, identity/session security, atomic aggregate/guard, operator documentation, and rendered frontend. The cross-layer/auth invariants and hard dependencies benefit from separate task context. This is not a file-count split.

Chosen axis: **dependency/sequence with concern ownership**. T1 establishes the shared contract; T2 and T5 are separate backend/frontend concerns after that contract. T3 consumes the authenticated session interface from T2. T4 documents the actual guard controls delivered by T3. Review and verification obligations remain owned by parent §11/§12 across tasks, not a new implementation task.

## Tasks and hard dependency graph

| Task | Purpose | Hard dependencies |
|---|---|---|
| [T1-shared-contract.md](T1-shared-contract.md) | Coordinated OpenAPI source/bundle and generated types; no consumers. | None |
| [T2-backend-identity-session.md](T2-backend-identity-session.md) | Backend OIDC identity and local session boundary (G1). | T1: published contract source/bundle define auth shapes; durable source/bundle revision and validation/bundle result. |
| [T3-backend-aggregate-guard.md](T3-backend-aggregate-guard.md) | Atomic Organization/Owner establishment, guard and scoped reads (G2/G3). | T1: published contract source/bundle define establishment/read/error shapes; durable source/bundle revision and validation/bundle result. T2: server-side verified person/session context, same-person binding, mutation Origin/CSRF, session-row revalidation; T2 code/interface plus authored focused protocol/person-upsert/session-lock checks. |
| [T4-guard-operations-docs.md](T4-guard-operations-docs.md) | Guard runbook and scoped backend operating examples. | T3: implemented schema/control/privilege/audit/rollback behavior; durable schema/control/config refs and actual privilege-check evidence. |
| [T5-frontend-c1.md](T5-frontend-c1.md) | Frontend disclosure/confirmation/uncertainty/list/detail experience. | T1: published contract source, validated bundle and generated types; exact revisions and validation/bundle result. |

Graph: `T1 → T2 → T3 → T4`; `T1 → T3`; `T1 → T5`. No hard edge between T5 and T2/T3: frontend can implement against the contract and use contract-faithful fixtures; real integration is a later verification concern. Each edge is satisfied by its stated observable result and durable evidence, not broad predecessor completion.

## Shared coordination and invariant anchors

- Shared contract: T1 coordinates with backend/frontend concern owners; consumers do not rewrite API source/types.
- Backend/frontend production writes stay separate (T2/T3 vs T5).
- T3 owns server-side identity-to-Owner grant, session/guard/preparation lock order and atomic persistence; frontend cannot assert authority.
- T4 is operational documentation coordinated with T3; docs do not create independent authority or claim external verification.
- Preserve parent R1–R14, D1–D10, RISK-1–RISK-8 and §8–§12 verification/evidence ownership. G1–G3 authorization is bounded by WU events and parent Resolved 9–11. Real provider/DB/operator evidence remains owned by later execution/Testing/Human as parent specifies.
- T6 review/verification obligations in parent §9 are satisfied through the later canonical Review/Testing phases and evidence owners in §11; this manifest does not create a new Work Unit or implementation task.
