# Auth harness Build scope / disposable fixture authorization request

Status: PREPARED; requires exact Human approval for disposable provisioning and model before execution. Scope based on Approved Techplan/T2/API and current backend gaps, not a new behavior contract.

## Outcome and allowed writes

Author meaningful tests for current T2 HTTP auth boundary and real PostgreSQL person/session contention, plus smallest reusable test-owned isolated setup. Production auth/store/db/server/migration/module/scanner-policy source remains read-only. Allowed: backend/internal/platform/auth/*_test.go; minimal backend/scripts auth-test fixture runner/init support and focused test fixtures; backend README test instructions; backend Makefile focused harness target if needed. No .env/user credential updates, API/frontend/Compose/root infrastructure, owning Product/Techplan/task or migration semantics writes. Tests may be tagged to keep DB checks opt-in; tags do not substitute for assertions.

## Exact requested provisioning boundary

- Create one unique Run-owned disposable PostgreSQL container using existing configured postgres:16 major route; record actual image identity/version rather than claim immutable digest not observed. No version/dependency change or image build.
- Only generated test data. No mount of repository .env, host existing PGDATA, shared named volumes or connection to existing Compose/dev/historical databases. Loopback-only ephemeral host port or isolated private container network; Run-specific container/resource labels and collision check, ephemeral storage.
- Separate fixture bootstrap/provisioner from restricted LOGIN runtime. Bootstrap/provisioner may create/drop this fixture schema/roles and apply exact current identity/session migration in this container only. Runtime NOSUPERUSER/NOCREATEDB/NOCREATEROLE/NOREPLICATION/NOBYPASSRLS; no owner membership, DDL/schema CREATE/role management or ability to grant itself authority.
- Runtime receives only CONNECT/schema USAGE and T2 operations: persons SELECT/INSERT/UPDATE; pending_logins SELECT/INSERT/DELETE; sessions SELECT/INSERT/UPDATE, preserving active code lock/revoke requirements. No historical/Organization/guard tables introduced. Revoke unnecessary PUBLIC privileges in fixture. Confirm actual runtime identity/grants/denials; provisioner seeding/expiry setup is fixture orchestration, never runtime authorization evidence.
- Generate ephemeral test-only passwords privately, SCRAM authentication over actual TCP and inspect applicable HBA/password-storage evidence without values. No host .env credentials read or user passwords requested; DSNs/tokens private in-memory/child environment or private temporary file only if needed, never command/evidence/log output. Use non-system env names for explicit fixture provisioning/runtime DSNs if needed; do not reuse user's DATABASE_URL for fixture.
- Test-owned migration setup only; no migrate-up/down on populated/shared DB, no volume reset or modifying persistent PostgreSQL roles/grants/config. Cleanup all Run-owned container/network/temp fixture credentials/data on success/abort with bounded timeouts; never delete resources by broad prefix without verified ownership. Inspect leftovers without printing private values.

This authorization, if granted, applies solely to ephemeral test fixture; no new production elevated rule, credential rotation, auth/key mechanism or privileged user path. If environment cannot enforce isolation or needs wider privilege/resources/new dependency, stop for owner decision.

## Meaningful checks to author and execute

HTTP registered handler at httptest/local loopback boundary, actual restricted Store pool where valid session required:
- /me anonymous/malformed/expired/revoked deny and valid session response/no-store; JSON fields contain local person/CSRF only, no provider payload; storage failure safe envelope.
- CheckMutation exact Origin/JSON/session-bound CSRF negatives and valid controls; logout body invalid/extra/size cases, successful revocation plus cookie deletion, repeat revoked session denial; no implicit credentialed CORS or raw internals.
- /auth/google/start and callback with **simulated Protocol**: fixed local redirects, browser-bound pending state/one-time consumption and provider failure/duplicate/malformed input safe no-session paths; success maps local identity/rotates session under fixture. No real Google network/login, code replay claims only simulated/local pending-state boundary.
- Preserve reviewed cookie/config oracle; add actual handler checks rather than duplicate all existing verifier/cookie tests. Production-cookie serialized headers may be checked, but no actual HTTPS/browser enforcement claim.
Real PostgreSQL/runtime checks:
- Actual runtime pool connection/SCRAM/grant restrictions; lack of DDL/role/owner powers, operation compatibility for T2 queries.
- Controlled concurrent same issuer/subject first-login upsert yields one stable person (case/issuer boundaries), no duplicate identity authority.
- Controlled row-lock wait: session expires before acquisition => ErrUnauthenticated/zero authority after lock; revoked state under lock denied; logout/revoke/rotation serialize on actual row and old session invalidates as approved. Include cancellation/timeouts/resource cleanup and bounded release order.
- Deterministic barriers/observable lock waiting, not merely sleep or Go race detector as proof. Use current real LockSession path; fake clock/provider portions explicitly labeled. No T3 confirmation/Organization/guard operation simulation labeled real.

Build executes newly authored tests enough to establish executable meaningful oracles, explicitly lists justified specialized real DB/concurrency/security checks in report; final Testing owns independent result. If tests expose source/contract defect or new semantic decision, retain failed evidence and report smallest patch plan/owner route; do not change expected assertions to match bug or modify production source silently.
