# Scope — RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-CREDENTIAL-PATCH-001

Status: PREPARED — Human mechanical dispatch pending. This patch requires no fixture authorization; no source edit or test has occurred.

## Re-entry and bounded outcome

Address only blocking Review finding F1 from `RUN-C1-ENG-CODEREVIEW-T2-CONFIGURED-DB-HARNESS-PATCH-001`: prevent fixture credentials from reaching the `go version` metadata subprocess while preserving the configured runtime URL in the intended Go test child. Keep the fix inside the existing test harness.

## Allowed writes

- `backend/scripts/auth_harness.py` — run toolchain metadata lookup with a sanitized child environment that excludes fixture secrets and DSNs; leave the intended Go test child's fixture environment intact.
- `backend/scripts/test_auth_harness.py` — add a focused mocked regression proving the `go version` subprocess receives no fixture credential variables while preserving required nonsecret toolchain configuration and leaving the original test-child environment unchanged.

Do not change production code, auth behavior, DB/Store, migrations, README, dependency pins, Makefile, scanner policy, Product/Techplan/task authority, prior Run evidence, shared DB, `.env`, or any other file. Preserve all dirty/untracked workspace content.

## Verification

Run only the focused mocked Python unit suite from `backend/`:

```text
python3 -m unittest discover -s scripts -p test_auth_harness.py
```

No PostgreSQL fixture, Podman operation, database, Provider, browser, runtime, scanner, gate, `make verify`, dependency command, or real Google login is authorized. If F1 cannot be fixed and tested within the two allowed files, stop and report the exact gap rather than broadening scope.

## Handoff and limits

Write only this Run's `evidence/patch-report-1.md`, with exact bindings, delta/hashes, the focused command/result, deferred claims, and exactly one `## Phase handoff`. This patch does not establish R14 or B-T2-003 and cannot close either blocker or B-T2-002. After terminal Build, Orchestrator will decide targeted Review confirmation from the actual delta; Testing remains held until Review reconciliation and must use its own separately authorized disposable fixture.
