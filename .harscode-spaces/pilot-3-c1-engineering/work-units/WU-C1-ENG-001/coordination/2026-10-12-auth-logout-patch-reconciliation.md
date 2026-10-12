# Logout patch terminal reconciliation and claim-scoped route

Date: 2026-10-12. Author: Orchestrator / Codex. Read-only verification of terminal Build evidence and live identities; no tests, scanner, DB, Podman, or provider action by Orchestrator.

## Terminal Run reconciliation

`RUN-C1-ENG-BUILD-T2-LOGOUT-PATCH-001/evidence/patch-report-1.md` SHA-256 `68db135f66a8755c900734bad86fd6dc4e94db9d3bcf0b82cc50a40d40594bc2` matches the Human-submitted identity. The report contains exactly one `## Phase handoff`, outcome `COMPLETED` for its bounded Build claim. Its Participant, fresh Session, `gpt-6-luna` / `medium`, branch/target revision, and workflow revision match the prepared Invocation and current registry routing (`approval_required: false`). This is report/Invocation identity reconciliation, not independent runtime attestation of the model.

All 20 assignment-defining pre-execution hashes in `input-bindings.json` were reported matched before edit/setup. The current manifest itself matches the Invocation SHA-256 `077c4b755a7b452f4280375f1884140d049b94fe0f274d88eb44b78fee5c2053`. Read-only live comparison confirms 15 bound inputs remain identical; the five permitted target deltas match the report's after hashes:

| Path | Before | After |
| --- | --- | --- |
| `backend/internal/platform/auth/http.go` | `bd6a7706c4e5b9ad9af69a33a537b9af0e28fa74a852672692a96f3d034b5009` | `83b575b0e43d17d85303ae7dfeb0273fa2516bb18952a9660da7b80a2d039d63` |
| `backend/internal/platform/auth/harness_test.go` | `bdbfe0b59267125de9572ff534eb6839650b7ace1a3ecacd5965d22219c28552` | `c9997611413446641da20b73da2e251459a9ccbb33fbb7c88429f18206d0f45a` |
| `backend/scripts/auth_harness.py` | `01a2e72822f990528a9d3831430c42aa4685a9581b95ac1b5374bd5116d7a6f6` | `04e8173381bb978e9ab52898969418b893672729f82b741645267aaef1e41e53` |
| `backend/README.md` | `38b5623e84b9b2fef704978b8f3100fd68e51d2b35f14fe939e02893ba85d99f` | `38d6dbb59f19f7f936ac43932c8d3a2028e9b07585e34e5baf9844b94b41574e` |
| `backend/scripts/scanner_gate.py` | `6c0f06e3ce2f48fa47218f68d4d0ee0405b5d6094b62400815401542d7c615dc` | `ea1aa24035c3532e93b0faa3419883e08306e4bb624d6debae845121d4746cd7` |

The production delta is one condition: reject a nil decoded JSON object before `Store.RevokeSession`. Contract and live regression source agree that `null` must return safe 400 and preserve the session; valid `{}` still follows successful revocation and deletion-cookie behavior. Reported focused HTTP/PostgreSQL regression passed for both paths and the selected prior invalid-body/Origin/CSRF checks. This is Builder-reported Build evidence, not independent runtime acceptance. Provider Protocol simulation was not real Google interaction.

G124 source re-triage was narrowly performed after the `http.go` hash changed. The report records one same-command retry allowed by the active bounded recovery envelope; effective result PASS with the same three already-accepted G124 cookie candidates, no new candidate or allowance. Orchestrator did not rerun Gosec. Fixture identity/ownership markers matched the new Run; runner reported owned container and anonymous volumes removed. The subsequent sandbox Podman listing was unavailable, so cleanup is established by runner checks only, not an independent Podman observation.

The previous `AUTH-HARNESS-001` terminal FAILED report and its evidence remain unchanged. Its five passing groups remain Build-only historical evidence; its null 204 failure remains preserved. No historical report was rewritten and no prior approval was transferred.

## Finding and phase applicability

Close `F-AUTH-HARNESS-001` only for the specific object-validation mismatch: the live handler now rejects JSON null before revocation, and the bound Builder regression reports the retained null/non-revocation and `{}`/successful-logout oracles passing. This does not assert absence of a CSRF/privilege bypass, independent runtime acceptance, or whole-auth/T2 acceptance.

An independent four-pass Code Review is applicable to the exact five-file patch delta. Although the production edit is one condition, it changes authenticated logout behavior and its ordering relative to session revocation. Fresh reasoning can add material value on Safety/contract placement, the regression oracle and fixture ownership, and the one guarded G124 identity update. Scope is strictly the actual patch delta; canonical Review requires all four passes after invocation.

An additional Testing Run is not applicable to closing this scoped Finding: the current bounded claim is already supported by the targeted fixture regression and narrow sourceguard evidence, and this reconciliation makes no independent-runtime or final backend acceptance claim. Repeating the same Build-authored regression solely because Build ended would add no distinct evidence. Independent runtime/security/dependency acceptance remains required under the owning T2 conditions and open blockers; the report does not satisfy it. Real Google and integrated FE/UI verification remain separately deferred by Human-approved holds.

## Next route

`RUN-C1-ENG-CODEREVIEW-T2-LOGOUT-PATCH-001` terminal report SHA-256 `f518be882dc2c60a9417b9a927a916731a467b5d693070256f98ae48ea430d0f` matches the Human report identity and prepared Run. It contains exactly one Phase handoff; 12/12 current assignment bindings match; Participant/Session/model/approval match Invocation and the Human-owned registry. Review completed all four passes, with **Approve with minor comments, 0 blocking, 1 nonblocking R1**. Reviewer performed no source edit or runtime verification. Orchestrator independently repeated no tests, scanners, DB, Podman, or provider actions.

R1 is an explicit-oracle gap: `TestHarnessLogoutRejectsNull` does not assert the absence of `Set-Cookie`. The live handler returns on decode error/nil before the only logout cookie emission, and the Reviewer confirms that ordering from source. R1 is therefore recorded as nonblocking; no Build patch plan or extra test Run is warranted for this exact closure claim. The byte preimage of README was not independently reconstructed, and the Review did not independently verify fixture cleanup/runtime; those limits remain explicit. This does not invalidate the scoped null-before-revocation fix, but it also does not establish independent runtime or T2 acceptance.

Claim-scoped Testing applicability remains split: another Testing Run for this precise Finding closure would duplicate bounded Build evidence without material independent value (N/A); independent T2 runtime/security/dependency Testing remains required under Techplan and open blockers. The durable `RUN-C1-ENG-TEST-T2-001/evidence/testing-report-1.md` is terminal STALLED (SHA-256 `b41db9398c5a68be240b458be0a07712645eae3304fe54219e46a69ec9b6ff94`), but a later `invocation.md` at the same Run path claims READY_FOR_HUMAN_DISPATCH. Monotonic terminal closure controls: preserve the terminal report and do not dispatch/reopen/reuse that Run ID; its later prepared Invocation/model approval is not a new occurrence. No new Testing Run was created here. Any broader T2 Testing re-entry requires a fresh Run ID/Participant/Session, exact Run-specific model approval, current claim/source/evidence bindings, and preflight/re-grounding of the executable DB/runtime fixture. The current `auth_harness.py` remains bound to LOGOUT-PATCH-001 and selects only its two logout HTTP tests; it is not a ready independent T2 runtime harness. Google remains Human-deferred; DB/SCRAM/session contention and current final gate remain unestablished.

`B-T2-002`, `B-T2-003`, historical `B-T2-001`, final runtime/security/dependency conditions, and both verification holds remain unchanged. No T2/T3 readiness, acceptance, dependency commit, or T3 dispatch is established.
