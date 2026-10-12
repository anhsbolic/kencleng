# C1 Auth Verification Environment

> Scope: Kencleng Pilot #3 C1 authentication/session verification.
> Status: project-specific engineering/runtime guidance.
> Authority: this document operationalizes current C1 verification obligations; it does not create Product meaning, change the Approved Techplan, or add a new verification requirement.

Use this guide when preparing real C1 auth/runtime verification. Keep simulated/provider-fixture evidence, real PostgreSQL evidence, real Google/browser evidence, and production-HTTPS evidence explicitly distinct.

## Environment boundaries

### Application runtime

The auth-enabled backend uses:

- `APP_ENV` / `APP_ORIGIN`;
- `DATABASE_URL` for the application runtime;
- `GOOGLE_CLIENT_ID` / `GOOGLE_CLIENT_SECRET` for Google OIDC.

For C1 verification, `DATABASE_URL` must represent the project-authorized restricted application runtime identity. Do not substitute the Compose bootstrap/owner identity merely to make auth work. The required schema must already have been applied through an authorized migration/provisioning path; this guide does not grant DDL or role-management authority to the runtime identity.

Development browser origin is exactly `http://localhost:8080`. The corresponding Google redirect URI is exactly:

```text
http://localhost:8080/api/auth/google/callback
```

The backend listens on its configured backend port (current local default `:8090`), while Caddy exposes the browser-facing `/api/*` path at `localhost:8080`.

### Automated integration fixture

`C1_AUTH_TEST_DATABASE_URL` is a separate disposable PostgreSQL integration-test configuration.

It may use fixture/provisioning capability required by the integration tests and must never point to production or shared application data. Passing this integration suite does not prove that the application runtime credential is restricted, that SCRAM/runtime configuration is correct, or that real Google/provider interaction works.

Tests do not load `backend/.env` automatically; export this variable explicitly for the integration-test process.

### External Google provider

Real-provider verification requires all of the following before claiming that scenario is runnable:

- a Google OAuth client configured for this local C1 flow;
- the exact localhost callback registered;
- consent/test-user access appropriate to the configured Google project;
- network access to the provider;
- a Human-controlled Google account able to complete the interactive sign-in.

Presence/readiness may be checked without exposing secret values. Do not persist Google passwords, MFA secrets/codes, client-secret values, authorization codes, ID/access tokens, or local session tokens as verification evidence.

### Human browser interaction

Real Google browser verification remains Testing-owned but is Human-performed for the private provider-account interaction. The Human controls account selection, credential entry, consent, and MFA. The agent prepares the bounded scenario and may independently observe application-side evidence it can legitimately access.

Record the Human-observed portion as `Human-attested`; do not relabel it as independently executed by the agent.

## Readiness preflight

Before dispatching or executing a Testing scope that claims real Google/runtime evidence, establish only the material prerequisites for that scope:

- backend can be started with the intended C1 configuration;
- the required schema exists and `DATABASE_URL` is the authorized restricted runtime identity;
- `APP_ORIGIN=http://localhost:8080` for local development;
- Google client configuration is present without recording its secret value;
- the exact callback is registered;
- Caddy/browser route is available when browser verification is in scope;
- the Human confirms an eligible Google account/test-user path is available;
- required network access is available.

A known missing item is a verification prerequisite gap, not an implementation failure. Resolve it first or explicitly narrow the Testing claim. Bounded verification that does not need the missing prerequisite may still proceed.

## Local preparation

Create local backend configuration from the tracked template:

```bash
cd backend
cp .env.example .env
# Fill local values without committing the file.
```

Keep `.env` / `.env.local` ignored. The repository secret gate intentionally scans Git-visible backend content rather than legitimate ignored local credentials.

Start only the services needed by the active verification scope. The repository provides Podman/Compose PostgreSQL and Caddy routes, but service availability alone does not establish restricted database identity or Google readiness.

When real-browser auth is in scope:

1. ensure the database schema is already migrated by an authorized path;
2. start the backend with the restricted runtime `DATABASE_URL` and local Google configuration;
3. make the browser-facing `http://localhost:8080/api/*` route available through the current Caddy topology;
4. verify readiness without printing secrets;
5. only then ask the Human to execute the bounded browser scenario.

Do not invent a privileged database role, weaken auth/security behavior, or change provider configuration merely to make Testing pass.

## Real Google browser scenario

This scenario establishes provider interoperability and usable local authenticated state; it does not replace the protocol-negative, database-concurrency, restricted-credential, or production-HTTPS evidence assigned elsewhere.

Before Human execution, Testing states the exact observable PASS criteria. A minimal T2 real-provider scenario is:

1. Human opens `http://localhost:8080/api/auth/google/start` in the intended browser context.
2. Human completes Google account selection/sign-in/consent/MFA privately.
3. Browser returns through the fixed Kencleng callback flow.
4. Human observes that `http://localhost:8080/api/me` returns the authenticated local state in that browser context.
5. Where stable-person-across-sessions is being verified, repeat the authorized login/session scenario and combine the Human-observed provider flow with application/database evidence sufficient to establish that the same Google issuer/subject maps to the same local person.
6. When logout is in the bounded scope, exercise the current logout contract and report only the observable result required by that test.

The Human should report only the requested PASS/FAIL observation and relevant symptom. Do not ask the Human to paste passwords, MFA values, OAuth codes, provider tokens, client secrets, cookie/session values, or unnecessary person/account identifiers.

Example evidence classification:

```text
Scenario: real Google OIDC browser flow
Execution: Human-performed
Result: PASS — Human-attested
Observed: provider flow completed; browser returned through Kencleng; authenticated /api/me state observed
Agent-independent provider execution: not performed
Sensitive values captured: none
```

Any additional application-side observation performed by the agent must be recorded separately as agent-executed evidence rather than silently folded into the Human attestation.

## Evidence boundaries

A successful localhost Google flow may establish, within the exact exercised scenario:

- real Google authorization interoperability;
- registered redirect/callback compatibility;
- successful application handling of a real provider response;
- issuance/use of a local authenticated session;
- stable local-person mapping only when the required repeat-login plus application/database evidence is actually established.

It does not by itself establish:

- every issuer/audience/azp/signature/nonce/PKCE negative;
- code/state replay resistance not explicitly exercised;
- PostgreSQL concurrency/locking behavior;
- restricted runtime privilege proof;
- all CSRF/Origin/logout negative behavior;
- production `Secure` / `__Host-` cookie behavior.

Local development uses the explicit HTTP localhost cookie exception. Production-cookie evidence requires a separate HTTPS environment/scenario; do not relabel localhost HTTP evidence as production HTTPS verification.

## Failure and cleanup

Classify failures before routing:

- provider/account/config/network prerequisite missing → environment/external prerequisite gap;
- Human reports an unexpected observable result → Human-attested FAIL for that exact scenario;
- backend/application-side deterministic failure reproduced by the agent → agent-executed verification failure;
- conflicting Human and application-side evidence → verification not established until reconciled.

After bounded verification, stop disposable processes/services and clean disposable integration state according to the active Run's authorization. Do not delete shared/project state or credentials merely for cleanup.
