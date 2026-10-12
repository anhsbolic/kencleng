# Scope — RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-001

Status: PREPARED. Exact model approval and new Run-owned disposable fixture authorization are pending. No source/test/fixture operation has occurred in this Run.

## Bounded outcome and allowed writes

Make the existing opt-in auth test harness safely reusable by a later independent Testing Run with fresh fixture ownership. The current runner and Go test marker hardcode `RUN-C1-ENG-BUILD-T2-LOGOUT-PATCH-001`, and the runner selects only two logout tests. That prevents a new Testing-owned container and the existing broader real PostgreSQL test suite from sharing the same explicit Run/resource identity.

Allowed write surface only:

- `backend/scripts/auth_harness.py` — accept a required, validated active Run ID and a bounded suite selector; apply the same exact ID to resource labels, child environment and cleanup ownership checks before and after provisioning.
- `backend/internal/platform/auth/harness_test.go` — accept only the supplied valid harness Run identity rather than the closed Build ID; preserve existing assertions.
- `backend/README.md` — document the new Run-owned invocation and evidence limits.

Do not edit production auth/store/db/server code, API/OpenAPI, frontend, migrations, dependencies, Compose topology, Makefile, scanner policy/allowances, Product/Techplan/task authority, or another Run's evidence. Do not add the Review's optional `Set-Cookie` assertion in this assignment; it remains a nonblocking comment and is not needed for reusable ownership. Preserve all existing tests and terminal reports.

## Focused suite and outcome criteria

Create an explicit allowlisted `t2` suite selection that executes all six existing `TestHarness...` top-level cases under the new Run-owned fixture:

1. `TestHarnessRuntimePrivileges` — actual TCP/SCRAM with restricted runtime role, exact grants/denials and wrong-password rejection.
2. `TestHarnessHTTPMeMutationLogout` — real restricted Store with registered httptest handler, session/Origin/CSRF/body/logout assertions including valid `{}`.
3. `TestHarnessHTTPSimulatedCallback` — local handler + simulated Protocol; report provider simulation distinctly.
4. `TestHarnessConcurrentPerson` — real PostgreSQL upsert contention and stable person mapping.
5. `TestHarnessSessionLockContention` — real PostgreSQL expiry/revocation/cancellation/logout/rotation lock behavior.
6. `TestHarnessLogoutRejectsNull` — null -> 400 and session remains valid.

Retain the existing narrow logout subset as a separate allowlisted selector if needed; selectors must not accept arbitrary test-regex input. The required `t2` invocation for this Build Run is:

```text
python3 scripts/auth_harness.py --run-id RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-001 --suite t2
```

Build execution proves only that authored/reused tests and runner are executable against the new disposable fixture. It is not independent Testing acceptance. Existing old Build passes are claims/history to preserve; this Run must actually select the full six-case suite under its own fixture.

## Exact disposable fixture boundary — new approval required

Request fresh Human authorization for this exact Run's one PostgreSQL fixture. Reuse the previously approved bounded fixture design as a reference, not as transferred Run ownership:

- one unique, ephemeral Run-owned `postgres:16` container; use the configured Podman route and existing local image only; loopback ephemeral port, no host bind mounts, no shared named volumes or existing Compose network;
- fixture-only current identity/session schema/migration, with separate provisioner and runtime login; runtime remains NOSUPERUSER/NOCREATEDB/NOCREATEROLE/NOREPLICATION/NOBYPASSRLS with only the nine existing T2 table privileges and no grant option/owner membership;
- generated ephemeral test passwords, private child environment/stdin only; never read/load `.env`, `DATABASE_URL`, provider credentials, or print/write DSNs, passwords, tokens or payloads;
- verify the exact resource Run label and unique owner before any deletion; clean this Run's container/anonymous data volume on success or failure with bounded timeouts; never clean by broad name prefix or affect unrelated resources.

If isolation, restricted grants, or owned cleanup cannot be established, stop and report the exact limitation. Do not fall back to shared/persistent DB or elevated application identity. No real Google/provider login, browser, HTTPS deployment or app server.

## Execution envelope and next applicability

- **PREAUTHORIZED after both exact approvals:** edit the three listed test-harness files, execute the declared `t2` Build suite against this Run-owned fixture, and write this Run's report/evidence. Focused formatting/static syntax checks for authored files are allowed.
- **ORCHESTRATOR_DECISION:** after terminal report, decide exact-delta Review applicability and prepare a fresh independent Testing Run only from actual changed files, results, current runtime readiness and required command-gate freshness.
- **HUMAN_REQUIRED:** exact `gpt-6-sol` / `medium` model approval, this Run's fixture provisioning authorization, Human mechanical dispatch, any broader privilege/resource/scope change, or a new product/runtime semantic choice.

After this test-harness delta, the current-tree `RUN-C1-ENG-TEST-T2-BACKEND-GATE-002` PASS is no longer automatically current for the changed `backend/` files. A later independent Testing Run must bind its own fixture owner and re-ground/rerun applicable final command gates; it must not use either terminal Testing Run ID. Review applicability will be reconsidered after the exact delta: independent inspection of run-ID validation, privilege separation, label ownership and cleanup can materially protect evidence trust. Do not create a Review Run in advance or claim broader T2/T3 acceptance.

## Report requirements

Write `evidence/build-report.md` with canonical Build provenance, exact file delta, actual suite selection, real PostgreSQL versus simulated-provider evidence, restricted runtime identity/grants, cleanup evidence/limits, focused verification, contract check, deferred gaps and exactly one `## Phase handoff`. No secrets or full private process environment in evidence.
