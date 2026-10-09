# T2 — Backend identity and session

> Phase: Techplan decomposition; Author: `PARTICIPANT-C1-ENG-DECOMPOSER-001`; Created/Updated: 2026-10-09. Model: `gpt-6-luna`; Reasoning: medium; Session: `SESSION-C1-ENG-DECOMPOSER-001` (Run binding). Target revision: `524ef600c7f71246af6b71671d89c3c040fd9d44`; workflow revision: `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`.

## Purpose and outcome

Implement Google OIDC identity mapping and local opaque session/authentication boundary in backend, on the coordinated contract. Provide the authenticated person and session/CSRF behavior required by T3 and the approved interface. This is the bounded G1 surface.

## Authority

- Parent: `../techplan.md`, SHA-256 `bd732748b486f0327f6cba2cdf14cab32f1ece3c5be8ebb24f9659444947fa47`, status Approved.
- Governing items: Q1, Q3, Q6; R1, R8–R10, R14; D3, D6; RISK-2, RISK-3; §8 auth operations and cross-layer boundary; §9 lock/revocation semantics; §10 backend anchors; §11 verification; §12 G1 authorization; §13 Active item 5.
- G1 implementation authorization is recorded in WU `events.md` and parent Resolved item 9, bounded exactly as stated there. Do not request it again or expand it.

## Hard dependency: T1

**Condition before contract-dependent implementation:** T1's coordinated OpenAPI source and validated bundle define the auth/session operations and error/authority shapes this backend implements. Durable evidence: the committed/reviewable API source and generated bundle revision plus recorded validation/bundle result. T2 may prepare package/config structure beforehand only if it does not establish a conflicting external contract; the dependent handlers must follow the published contract.

## Scoped implementation

Use live anchors `backend/cmd/server/main.go` (composition/config/health), `backend/internal/platform/db/db.go` (pool), `backend/go.mod` (Go 1.24.8/pgx baseline), and `backend/migrations/`. Additive schema and backend auth/config/dependencies as needed. Implement fixed Google OIDC trust (issuer/client, RS256 verified keys, exact issuer/audience/applicable azp, expiry/nonce, browser-bound one-time state and S256 PKCE); stable case-sensitive `(issuer, subject)` person mapping; provider-token discard; opaque random 32-byte local session with digest-only persistence, CSRF, expiry/rotation/revocation; cookie, Origin, callback redirect, no-store and secrecy rules from parent R9–R10. Preserve health route and use restricted application DB identity where applicable.

Apply the exact explicit localhost cookie exception only in its stated development context. Pin selected compatible dependencies; inspect resolved source/graph and preserve parent supply-chain checks.

## Verification and handoff

Author focused protocol-negative/person-upsert/session-lock checks; record dependency source/graph review, `go mod verify`, vulnerability scan and available compile/auth checks per parent §10/§12. These do not substitute for real configured Google login, real PostgreSQL concurrency, or independent Testing. Provide T3 the durable server-side authenticated person/session interface, including the session-row lock/revalidation behavior used during confirmation, and its code/test pointers.

## Explicit boundaries

No custom cryptographic implementation, local JWT/signing-key system, refresh token, password/MFA/recovery, email/domain/body-person authority, role engine, or broad auth features. No Organization aggregate/guard implementation (T3 owns that); no frontend writes. Surface missing material interface meaning for parent reconciliation.
