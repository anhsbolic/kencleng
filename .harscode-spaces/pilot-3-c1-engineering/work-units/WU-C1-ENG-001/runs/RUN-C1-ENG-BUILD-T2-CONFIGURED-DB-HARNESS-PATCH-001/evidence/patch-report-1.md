# Build/Patch Report — RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-PATCH-001

- **Phase:** Build/Patch
- **Author:** `PARTICIPANT-C1-ENG-T2-CONFIGURED-DB-HARNESS-PATCH-IMPLEMENTER-001`
- **Created / Updated:** 2026-10-13
- **Model / Reasoning / Session:** `gpt-6-luna` / `medium` / `SESSION-C1-ENG-T2-CONFIGURED-DB-HARNESS-PATCH-IMPLEMENTER-001` (Invocation assignment; runtime metadata not independently exposed)
- **Target revision / Workflow revision:** `79ec63d57674d1e4976ff80b2dfee8c50bf17934` / `651ce57f044bff23fba99082322cee6c066d25a6`
- **Bindings:** Every assignment-defining SHA-256 in `input-bindings.json` matched before edits, including scope, Approved Techplan, T2 task/manifest, local runtime config, three allowed source files, production `db.Open`, Store/session-lock anchors, migration, environment guidance, and bound Build/Testing reports.
- **Fixture authorization / dispatch:** Invocation records Human approval for this Run's exact disposable PostgreSQL fixture. This user instruction dispatched the Run. Initial sandboxed attempt failed before fixture creation because Podman could not write `/run/user/1000/libpod`; the same scoped command then ran with approved host Podman access.

## What changed

- `backend/internal/platform/auth/harness_test.go`: validates the child `DATABASE_URL` against loopback, `c1_runtime`, `c1_fixture`, and `c1_identity`, then creates and pings the restricted runtime pool through production `db.Open`. Provisioner and lock-witness connections remain fixture-only. All six existing top-level `t2` selections and assertions remain unchanged, including the expiry case's observed blocker edge, wait past stored expiry, and under-lock unauthenticated outcome.
- `backend/scripts/auth_harness.py`: continues removing inherited application/provider/test DSNs before fixture setup and passes the generated restricted fixture URL only as `DATABASE_URL` in the Go test child environment. It does not print the URL or credentials.
- `backend/README.md`: documents the child-only configured pool path and its evidence limits.

Changed-file SHA-256 (before → after):

| File | Before (bound) | After |
|---|---|---|
| `backend/internal/platform/auth/harness_test.go` | `eedb6e8007158a0900301702a1896b48abb6e494211ced91f72fa30e0037f274` | `72c63063aa6709216d110d3f6a41cf81d981b510c9e583305b56d43399567ba8` |
| `backend/scripts/auth_harness.py` | `932a427e7917f73f2917ea316dbe62d29e3ffca649ebc5dc361db646319fe79a` | `fa5ebc1442b18cce82d5b5b1270855baa661ad4fbadb82456ffbf3b488d8948a` |
| `backend/README.md` | `91f9d1c084e90f2b862bb15b67dc782458b61052d820fabc79b8d2a88b5c8046` | `aa48bbeca8f8d186cc806d53769806da6684a1f3fac9cc80ea16de3b66fb0ac4` |

The only Run evidence written is this report. Pre-existing dirty and untracked workspace content was preserved.

## Tests run

From `backend/`, the only scoped execution command was:

```text
python3 scripts/auth_harness.py --run-id RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-PATCH-001 --suite t2
```

First attempt: blocked before fixture creation by the sandbox's read-only Podman runtime path. Authorized retry: exit 0. Runner reported Go `go1.26.9 linux/amd64`, PostgreSQL `16.14`, ownership verified before and after provisioning, and PASS for all six selected cases:

1. `TestHarnessRuntimePrivileges`
2. `TestHarnessHTTPMeMutationLogout`
3. `TestHarnessHTTPSimulatedCallback`
4. `TestHarnessConcurrentPerson`
5. `TestHarnessSessionLockContention` (expiry, revocation, cancellation, logout, rotation subcases)
6. `TestHarnessLogoutRejectsNull`

The expiry subcase observed a real PostgreSQL blocking edge, waited beyond the stored expiry deadline, and passed the assertion that lock revalidation returns no session authority. Runtime verification used the restricted `c1_runtime` role over loopback TCP/SCRAM, checked the exact nine non-grantable table privileges and privilege denials, and observed PostgreSQL 16.14. HTTP behavior used `httptest`; callback behavior used the simulated Protocol implementation. No Google/provider exchange or browser flow was performed.

Cleanup evidence: runner printed `cleanup verified: owned container and anonymous data volumes removed` after ownership-checked cleanup of this Run's container. Fixture ownership was verified before and after provisioning. Generated credentials and DSNs were not emitted.

## Verification scope confirmation

No race detector, performance/load sweep, scanner, dependency verification, gate, or independent Testing command ran. The single required t2 command includes targeted concurrency/lock and runtime-privilege assertions because those six established cases are the explicit Run scope; these were not additional broad sweeps. No separate focused test command was run.

## Contract check

- [x] Bounded configured-pool Build target satisfied: the restricted harness pool uses child fixture `DATABASE_URL` through production `db.Open`.
- [x] Live-code re-grounding confirmed the six selected tests and expiry-under-lock assertion were preserved.
- [x] Writes stayed within the three allowed source files plus this Run's required evidence report.

## Deferred / not tested here

Independent Testing must use a fresh Run and its own fixture to establish configured-pool/session-query evidence. Real Google, browser, HTTPS, deployed application configuration, scanners, dependency gates, `make verify`, and T2 acceptance were not tested or established here. B-T2-002 still requires its owning acceptance/commit decision; B-T2-003 remains separately open for reconciliation. No readiness, acceptance, Human hold release, or T3 dispatch is claimed.

## Flagged for Techplan / Testing

None. The required child fixture path was executable through `db.Open`; simulated provider and `httptest` evidence remain explicitly bounded.

## Phase handoff

- **Completed:** Configured-pool harness Build/Patch; six `t2` cases passed on this Run's disposable PostgreSQL fixture.
- **Artifacts:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-PATCH-001/evidence/patch-report-1.md`; three changed source files listed above.
- **Human decision:** None required by this completed Run; no acceptance or residual-risk decision is implied.
- **Open / deferred:** Independent Testing owns fresh fixture verification of configured pool/session queries and its bounded R14 claim. B-T2-002 owning acceptance/commit and B-T2-003 reconciliation remain open.
- **Recommended next step:** Orchestrator reconciles this exact delta and decides whether targeted Review materially helps, then prepares a fresh independent Testing Run if appropriate.
- **Session transition:** This Run is terminal after this handoff; later Review or Testing uses its own Run identity and fresh Session context.
- **Context pointers:** Approved Techplan and T2 task; this Run's scope; changed harness, runner, and README; bound prior Build/Testing reports.
