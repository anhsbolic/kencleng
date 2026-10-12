Phase: Build/Patch. Author: Codex / Implementer, backend narrow auth-patch.
Created: 2026-10-12. Model / Reasoning: gpt-6-luna / medium.
Session: SESSION-C1-ENG-T2-LOGOUT-PATCH-IMPLEMENTER-001 (fresh Run session).
Target revision: `dae55ac44beb0e967b94c2f791d8a0b8d8e72372`, `pilot/3-c1-engineering`.
Workflow revision: `651ce57f044bff23fba99082322cee6c066d25a6`.
WORK_UNIT_ID: WU-C1-ENG-001. RUN_ID: RUN-C1-ENG-BUILD-T2-LOGOUT-PATCH-001.
PARTICIPANT: PARTICIPANT-C1-ENG-T2-LOGOUT-PATCH-IMPLEMENTER-001.

## What changed

`backend/internal/platform/auth/http.go`: logout now rejects a nil decoded object (`null`) through the existing safe HTTP 400 `invalid_request` path before `Store.RevokeSession`. Empty object `{}` remains accepted. The change is one condition only; it does not alter cookie, Origin/CSRF, body limit, trailing-token, or revocation semantics.

Supporting scope changes rebind `backend/internal/platform/auth/harness_test.go` and `backend/scripts/auth_harness.py` to this fresh Run, and the runner selects only the affected HTTP/logout test pair. `backend/README.md` now describes the passing null regression and fresh-run binding. `backend/scripts/scanner_gate.py` refreshes only the guarded `http.go` source hash after targeted G124 re-triage. No allowances or policy changed.

All 20 bound input hashes matched before edit/setup; `input-bindings.json` SHA-256 is `077c4b755a7b452f4280375f1884140d049b94fe0f274d88eb44b78fee5c2053`. Before → after SHA-256 identities:

- `backend/internal/platform/auth/http.go`: `bd6a7706c4e5b9ad9af69a33a537b9af0e28fa74a852672692a96f3d034b5009` → `83b575b0e43d17d85303ae7dfeb0273fa2516bb18952a9660da7b80a2d039d63`.
- `backend/internal/platform/auth/harness_test.go`: `bdbfe0b59267125de9572ff534eb6839650b7ace1a3ecacd5965d22219c28552` → `c9997611413446641da20b73da2e251459a9ccbb33fbb7c88429f18206d0f45a`.
- `backend/scripts/auth_harness.py`: `01a2e72822f990528a9d3831430c42aa4685a9581b95ac1b5374bd5116d7a6f6` → `04e8173381bb978e9ab52898969418b893672729f82b741645267aaef1e41e53`.
- `backend/README.md`: `38b5623e84b9b2fef704978b8f3100fd68e51d2b35f14fe939e02893ba85d99f` → `38d6dbb59f19f7f936ac43932c8d3a2028e9b07585e34e5baf9844b94b41574e`.
- `backend/scripts/scanner_gate.py`: `6c0f06e3ce2f48fa47218f68d4d0ee0405b5d6094b62400815401542d7c615dc` → `ea1aa24035c3532e93b0faa3419883e08306e4bb624d6debae845121d4746cd7`.

## Tests run

- `python3 scripts/auth_harness.py` from `backend/` → focused actual HTTP regressions on the approved disposable restricted-PostgreSQL fixture; PASS. Runner selected Go `go1.26.9`, `TestHarnessHTTPMeMutationLogout`, and `TestHarnessLogoutRejectsNull` only. The former passed its safe error, Origin/CSRF, invalid/non-object/extra/trailing/body-limit checks and valid `{}` response/cookie/revocation assertions. The latter observed JSON `null` → 400 and confirmed the session remained valid through the restricted runtime identity.
- `make gosec` → targeted G124 sourceguard re-triage; first sandbox attempt failed closed because two compiled Go cache entries were unavailable. One same-command retry under the invocation's bounded environment-recovery allowance passed the effective gate: `raw_exit=1 raw_findings=3 accepted=3 effective=PASS`. Three raw findings remain the same accepted G124 cookie candidates. The source diff was reviewed and does not change the cookie constructions/composition at lines 49, 158, and 277 or their callers/configuration. No new candidate or allowance was added.
- `git diff --check` → PASS.

Fixture identity: fresh container `c1-auth-harness-8b045f4720a9`, image identity `88a36c64c1003dad93f56daa12d1f8916ec66d1fa3e5fb1fb0ae7cb77efd56d1`; Run label and `C1_HARNESS_OWNED_RUN` marker matched `RUN-C1-ENG-BUILD-T2-LOGOUT-PATCH-001`, and owner label matched the exact container name. Fixture used generated synthetic credentials, loopback TCP, separate provisioner/restricted runtime identities, and the current migration. Credentials/DSNs stayed in child process environment and are not recorded.

Cleanup: runner verified ownership before removal and reported `cleanup verified: owned container and anonymous data volumes removed`; its post-removal container and individual anonymous-volume checks passed. A later sandbox-only Podman listing could not access the user runtime directory, so independent post-run listing was unavailable. This cleanup evidence is limited to the newly owned disposable container and its anonymous storage; it does not establish cleanup of unrelated resources.

## Verification scope confirmation

No race/concurrency suite, performance/load test, or security-class sweep was executed. The restricted PostgreSQL HTTP regressions were the expressly authorized focused Build exception needed to execute the retained null-before-revocation and valid-object oracles. Gosec was the specifically authorized narrow sourceguard re-triage, not a broad security gate. The result is Build evidence, not independent runtime acceptance.

## Contract check

- [x] Scoped patch target satisfied: JSON `null` receives safe 400 before revocation; `{}` remains valid.
- [x] Live-code re-grounding preserves the material contract assumptions and existing HTTP/cookie/CSRF boundaries.

## Deferred / not tested here

Independent Review/Testing confirmation, deployed application runtime credential configuration, real Google/provider interoperability, browser or production HTTPS behavior, T3 behavior, final backend verification, and whole-T2 acceptance remain unverified. No shared DB, `.env`, real Google login, provider/browser check, race suite, vulnerability/provenance scan, or full security gate was accessed or run.

## Flagged for Techplan / Testing

No new Techplan contradiction identified. `F-AUTH-HARNESS-001` remains open until evidence reconciliation. `B-T2-002`, `B-T2-003`, and integration/Google holds remain in force.

## Phase handoff

- Outcome: COMPLETED (bounded Build patch and focused verification passed).
- Result refs: this report; source delta in `backend/internal/platform/auth/http.go`; focused runner/test provenance in `backend/scripts/auth_harness.py` and `backend/internal/platform/auth/harness_test.go`.
- Findings: `F-AUTH-HARNESS-001` observed behavior is corrected by this patch; status remains open pending reconciliation.
- Decision requests: none for this bounded patch.
- Blockers: none introduced by this Run.
- Open / unverified: independent confirmation and all deferred items above; `B-T2-002`, `B-T2-003`, and integration/Google holds remain open.
- Recommended continuation: Orchestrator reconcile this report and hash, then route the smallest claim-scoped independent confirmation appropriate to the evidence gap. Do not infer blocker closure or T2/T3 readiness from this Build.
- Context refs: bound patch scope and input bindings; prior terminal AUTH-HARNESS report/finding; current source and focused regression evidence named above.
