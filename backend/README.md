# Kencleng Backend

> Status: clean delivery baseline

Neutral Go backend engineering substrate for Kencleng Pilot #3. No product domain, authentication model, API operation, database migration, or background job is considered delivered at baseline.

Product meaning is owned by `../docs/product/product-intent.md`. Backend architecture is owned by `../docs/project/kencleng-backend-tech-stack.md`.

## Local commands

Run these commands from `backend/`. Create the local configuration with
`cp .env.example .env`, then fill `DATABASE_URL`, `GOOGLE_CLIENT_ID`, and
`GOOGLE_CLIENT_SECRET`. Use a restricted application DB role. The development
browser origin is `http://localhost:8080`; register the exact Google callback
`http://localhost:8080/api/auth/google/callback`. Production requires an HTTPS
origin. The server reads `.env` from its working directory; keep the file here
and do not commit it.

Backend secret verification distinguishes repository exposure from legitimate
ignored local configuration. `make verify` scans backend Git history, staged
and unstaged tracked changes, and untracked files that Git does not ignore.
Ignored local `.env` / `.env.local` files are therefore outside the repository
leak gate unless they are force-tracked; a force-tracked secret remains visible
to the staged scan. Do not weaken Gitleaks rules or commit local credentials to
make the gate pass.

The scanner targets use Python 3 (`scripts/scanner_gate.py`) to apply the
2026-10-12 Human-approved C1 dispositions. `make gosec` runs every applicable
production Go file/rule (including findings behind `nosec` comments); it accepts
only G124 at the three reviewed cookie locations while the reviewed production
source set and hashes are unchanged. New or changed production Go source,
changed analysis, loading errors, or other findings require re-triage. The
supported Gosec build is module v2.28.0, whose report version is `dev`.

`make gitleaks-history` retains all/full-history/backend coverage and accepts
only the four reviewed commit/path/rule/line fingerprints, with exact historical
blob guards. `make gitleaks-worktree` has no dispositions and attempts both
tracked-change scans and every nonignored untracked backend file even after a
finding. Before any Gitleaks invocation, root/backend `.gitleaksignore` and
`.gitleaks.toml` filesystem entries are rejected, including directories and
symlinks (even dangling links); filesystem inspection errors also fail closed.
This is a configuration failure with scans not executed, not a scan PASS.
Gitleaks v8.30.1 can auto-load target ignore files despite the explicit ignore
path; the preflight prevents that bypass without deleting local artifacts.
Inline allowance comments remain disabled. Structured reports, execution status,
coverage and source identities fail closed on errors or drift. Output separates
findings (configured Gitleaks exit 10) from scanner/partial-scan errors (exit 1),
so even a valid partial report containing approved candidates fails. It separates
raw findings/exit from the effective accepted-disposition result and never
prints report payloads or scanner stderr. Run `make test-scanner-gate` for the
focused behavioral tests. An effective PASS is not raw zero-findings, runtime
authentication acceptance, or final backend verification.

`C1_AUTH_TEST_DATABASE_URL` is separate disposable integration-test configuration.
Tests do not load `.env` automatically, so export that variable explicitly when
running integration tests. A configured URL alone does not establish restricted
runtime/SCRAM or real-provider verification.

The test-only `scripts/auth_harness.py` runs the opt-in `authharness` tests in a
new owned `postgres:16` container (the image must already exist locally). It
applies the current identity migration through the fixture provisioner and gives
the separate TCP/SCRAM runtime only the nine required table privileges. It uses
an ephemeral loopback port, generated private child-environment credentials,
anonymous disposable storage, and ownership-checked container/volume cleanup.
It never loads `.env` or uses the application `DATABASE_URL`.

```bash
python3 scripts/auth_harness.py
```

Run only under an active approved disposable-fixture invocation. The runner and
test marker are currently bound to
`RUN-C1-ENG-BUILD-T2-LOGOUT-PATCH-001`; a later occurrence must rebind resource
provenance before executing these reusable assertions. The runner selects the
HTTP mutation/logout and null-logout tests with `-failfast`, a 90-second test
timeout, and preserved assertions. Callback Protocol is simulated;
HTTP checks use httptest. Neither proves real Google login, browser/HTTPS
enforcement, deployed runtime credentials, or independent T2 acceptance.

For real C1 authentication/session verification, including the restricted runtime DB,
Google callback/test-user prerequisites, and Human-performed browser flow, follow
[`docs/project/c1-auth-verification-environment.md`](../docs/project/c1-auth-verification-environment.md).

```bash
go run ./cmd/server
go test ./...
go test -race ./...
make verify
make migrate-up
make migrate-down
```

The baseline server exposes only `GET /healthz` and connects to PostgreSQL. Add product behavior only after the relevant delivery requirement is sufficiently clear.
