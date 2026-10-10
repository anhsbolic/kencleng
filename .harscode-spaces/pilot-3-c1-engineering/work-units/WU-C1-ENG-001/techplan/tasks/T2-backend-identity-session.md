# T2 — Backend identity and session

> Phase: Post-Approval task-snapshot reconciliation; Author: `PARTICIPANT-C1-ENG-T2-TASK-REFRESH-002`; Created/Updated: 2026-10-10. Model / Reasoning: `gpt-6-luna` / `medium` (invocation-selected dispatch configuration; runtime identity not independently exposed). Session: `SESSION-C1-ENG-T2-TASK-REFRESH-002` (fresh Run binding). Target revision: `59f2b25940fa4baf243b4f1d2ca2ba04c9b17fd4`; workflow revision: `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`.

## Purpose and outcome

Implement Google OIDC identity mapping and the local opaque session/authentication boundary in the backend against the coordinated contract. Provide the authenticated person and session/CSRF behavior required by T3 and the approved interface. This remains the bounded G1 surface and implementation outcome.

## Authority

- Parent: `../techplan.md`, SHA-256 `6d2e41e177e4adaf74f53a06d81dd568def75019c8dc4e4792d938a869896d79`, status Approved.
- Governing items: Q1, Q3, Q6; R1, R8–R10, R14; D3, D6, D8; RISK-2, RISK-3; §8 auth operations/cross-layer boundary; §9 lock/revocation semantics; §10 implementation and dependency gates; §11/§12 verification ownership; §13 Active item 5.
- G1 implementation authorization is recorded in WU `events.md` and parent Resolved item 9, bounded exactly as stated there. Do not request it again or expand it. G2 and G3 remain separately bounded as already approved; T2 does not implement them.

## Hard dependency: T1 — contract prerequisite satisfied

**Condition:** T1's coordinated OpenAPI source and published bundle define the auth/session operations and error/authority shapes implemented by T2. The condition is satisfied for dependent handler implementation by the reviewed and validated/published contract artifacts and durable records below; this does not imply backend implementation or runtime correctness.

- Source `api/openapi/index.yaml`: SHA-256 `ddefffbe588459cb0cba9f68dfedaccba3da8a1313f1956511cf628df7199965`.
- Bundle `api/openapi.yaml`: SHA-256 `38bb6c531804b54e9d5ce56d623c397a590722c56dcb7a3c12e646bb8280fa09`.
- Generated types `api/openapi.d.ts`: SHA-256 `d21e8aa1eec92a95cc1fed87d976f726f59b48bd334aa8e463844d48e16cb80c`.
- Durable evidence: `RUN-C1-ENG-CODEREVIEW-T1-PATCH-001/evidence/review-findings-1.md` (Approve; F1–F4 resolved, exact reviewed target); `RUN-C1-ENG-BUILD-T1-PATCH-001/evidence/patch-report-1.md` (Implementer-reported `npm run validate`, bundle, and types generation results).

## Scoped implementation

Use live anchors `backend/cmd/server/main.go` (composition/config/health), `backend/internal/platform/db/db.go` (pool), `backend/go.mod` (current observed `go 1.24.8`, pgx baseline), and `backend/migrations/`. The `go.mod` version is the observed current baseline anchor only; it is not the approved planned Go toolchain.

Implement the fixed Google OIDC trust (issuer/client, RS256 verified keys, exact issuer/audience/applicable azp, expiry/nonce, browser-bound one-time state and S256 PKCE); stable case-sensitive `(issuer, subject)` person mapping; provider-token discard; opaque random 32-byte local session with digest-only persistence, CSRF, expiry/rotation/revocation; cookie, Origin, callback redirect, no-store and secrecy rules from parent R9–R10. Preserve the health route and use a restricted application DB identity where applicable. Apply the exact localhost cookie exception only in its stated development context.

## Approved toolchain and dependency baseline (D8)

The approved planned Go runtime/build toolchain is **Go 1.26.9**. Use the parent's exact explicit dependency baseline:

- `github.com/coreos/go-oidc/v3` `v3.16.0`
- `golang.org/x/oauth2` `v0.34.0`
- `github.com/go-jose/go-jose/v4` `v4.1.4`
- `github.com/jackc/pgx/v5` `v5.9.2`
- `golang.org/x/text` `v0.41.0` (explicit planned baseline; it may remain annotated indirect in the module file)

These approved pins are the planned baseline only; their approval does not prove that the selected actual dependency graph, source, compatibility, or vulnerability scan is clean. Before any dependency commit, Build must verify and record the actual Go 1.26.9 toolchain/runtime; exact direct and explicit pins and resolved graph/source; `go mod verify` integrity; a fresh full source vulnerability scan (`govulncheck ./...`) with findings and traces; module/toolchain compatibility and focused compile; and applicability/disposition for every advisory in the actual graph, including surviving or newly surfaced findings. Confirm the five previously identified advisory ranges are absent from the selected graph and reconcile any new affected dependency/path or incompatibility before commit. Do not commit dependency changes until every required check passes. A zero scan exit status alone is not proof of a clean result. The earlier T2 probe's 50 findings and `B-T2-001` remain open context; they do not substitute for checks against the selected graph.

## Verification and handoff

Build authors focused protocol-negative, person-upsert, and session-lock checks; records exact toolchain, resolved dependency graph/source, integrity, source-scan, compatibility, compile, and surviving-advisory evidence under parent §10/§12 gates. These do not substitute for configured real Google login, real PostgreSQL concurrency, or independent Testing. Provide T3 the durable server-side authenticated person/session interface, including session-row lock/revalidation behavior used during confirmation, with code/test pointers.

## Explicit boundaries

No custom cryptographic implementation, local JWT/signing-key system, refresh token, password/MFA/recovery, email/domain/body-person authority, role engine, or broad auth features. No Organization aggregate/guard implementation (T3 owns that); no frontend writes. Preserve G1–G3 and all accepted manifest dependency conditions/topology. Surface any missing material interface meaning for parent reconciliation; do not alter parent or sibling tasks here.
