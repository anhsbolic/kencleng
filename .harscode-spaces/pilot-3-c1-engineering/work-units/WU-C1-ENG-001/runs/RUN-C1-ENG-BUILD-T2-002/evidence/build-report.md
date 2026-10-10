# Build report — RUN-C1-ENG-BUILD-T2-002

> Phase: Build; Author / Role: `PARTICIPANT-C1-ENG-T2-IMPLEMENTER-002` / Implementer; Created/Updated: 2026-10-10. Work Unit / Run: `WU-C1-ENG-001` / `RUN-C1-ENG-BUILD-T2-002`. Session: `SESSION-C1-ENG-T2-IMPLEMENTER-002` (fresh collaboration Session). Invocation selected `gpt-6-sol` / medium; runtime model identity/effort is not independently exposed. Target revision / observed HEAD: `59f2b25940fa4baf243b4f1d2ca2ba04c9b17fd4`, branch `pilot/3-c1-engineering`. Workflow revision / observed HEAD: `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`, workflow tree clean.

## What changed

**T2/G1 is incomplete and this Run stops at a new dependency advisory blocker. No dependency or implementation commit was made.** The prior T2-001 blocker was reconciled against the Human-approved successor: actual Go 1.26.9 was acquired and run, and the exact OIDC v3.16.0 / OAuth2 v0.34.0 / go-jose v4.1.4 pins resolved. The successor fixes the earlier Go 1.24.8 / jose v4.1.3 findings in this source scan. Its actual backend graph retains pgx v5.8.0 and x/text v0.29.0 with five new scanner-reported advisory IDs. The approved Techplan expressly leaves pgx outside its candidate reconciliation and requires any new affected dependency/path to be reconciled before dependency commit. No version substitution, risk acceptance, or waiver was made here.

Uncommitted, partial backend work exists and must be reviewed before continuation or reuse:

- `backend/migrations/000001_identity_session.{up,down}.sql`: G1-only person, pending-login, and local-session tables. No Organization/Owner/guard schema.
- `backend/internal/platform/auth/store.go`: one-time pending-login consumption, case-sensitive `(issuer,subject)` upsert, digest-only opaque session storage, rotation/logout, and exported `LockSession` under a caller transaction for T3's session-first lock order.
- `backend/internal/platform/auth/oidc.go`: Google provider discovery, OAuth2 S256 PKCE, RS256 verifier, exact issuer/audience/applicable azp/nonce/expiry checks. Provider-token persistence or browser authority was not added.
- `backend/internal/platform/auth/http.go`: auth start/callback, `/me`, logout, fixed redirect, environment-specific cookies, no-store, Origin/JSON/CSRF boundary, bounded start attempts.
- `backend/cmd/server/main.go`: G1 composition/config; health route retained.
- `backend/go.mod` / `go.sum`: uncommitted exact proposed auth pins and Go 1.26.9 directive. The `go-jose` pin appears in the explicit indirect require block because application code imports it through OIDC; the selected version is v4.1.4.

Exact changed-file SHA-256 identities are in `changed-file-identities.sha256`. The above is partial Build work, **not validated authentication**: no focused authored protocol, person-upsert, or session-lock tests exist yet, and migration/runtime behavior has not run. No API, frontend, Product, Techplan, task, manifest, T2-001 artifact, or G2/G3 implementation was changed by this Run. Unrelated working-tree changes were preserved.

Before edits, root/backend AGENTS, invocation, Approved parent/T2 task/manifest, T1 dependency evidence, active contract, backend architecture and live anchors were read. SHA-256 matches invocation for parent `bae2470f9038521b5d2d51fe4bc8de192fe5a9cf91bc0975831a3046fbf7a8ad`, T2 `ba34963b725b25f78d908b2e72fceb66b26aa207701110fb0031ffc34affcd8d`, manifest `2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d`, API source `ddefffbe588459cb0cba9f68dfedaccba3da8a1313f1956511cf628df7199965`, bundle `38bb6c531804b54e9d5ce56d623c397a590722c56dcb7a3c12e646bb8280fa09`, types `d21e8aa1eec92a95cc1fed87d976f726f59b48bd334aa8e463844d48e16cb80c`, and prior T2-001 report `130b888011a5c5286b818078410d2b022faf60d0033f20d0e89d628098c0ad50`. T1 contract dependency remains satisfied only for handler implementation, not runtime correctness. Approved five-task topology and G1–G3 boundaries remain unchanged.

## Tests run

| Command / scope | Result and evidence limit |
|---|---|
| `GOTOOLCHAIN=go1.26.9 go version` / backend | Initial sandbox run failed writing read-only module cache; approved escalated rerun exit 0, actual `go version go1.26.9 linux/amd64`. Local default was Go 1.24.8; no silent toolchain substitution. |
| Exact `go get` three approved pins / isolated `/tmp/c1-t2-002-deps` | Initial sandbox DNS/socket denial; approved escalated rerun exit 0. Probe kept separate from backend until pins resolved. |
| `go list -m all` / probe then actual backend | Exit 0. Actual graph recorded in `dependency-graph.txt` SHA-256 `6095e0b6d7d80d83c10411602246c2cf0c02b1d78888f5579f0350ad57823b9f`: exact OIDC/OAuth2/jose pins; pgx v5.8.0; x/text v0.29.0; no x/net. |
| `go mod download -json` for exact auth pins / backend | Exit 0. Module cache source/checksum locations recorded in `auth-module-source.json` SHA-256 `d3135377285e0c13d6657de0b419a488f7ef79717255181d4f61c9011230625d`. Tagged OIDC verifier and OAuth2 PKCE source were inspected for RS256/issuer/audience/expiry verification and S256 helpers; application checks supplement defaults. This source review does not settle pgx/x/text advisory risk. |
| `GOTOOLCHAIN=go1.26.9 go mod tidy` / backend | Sandbox build-cache denial; approved escalated rerun exit 0. Actual go.mod uses Go 1.26.9 and the three approved auth pin versions. |
| `GOTOOLCHAIN=go1.26.9 go mod verify` / probe and actual backend | Both exit 0, `all modules verified`; actual output in `dependency-integrity.txt` SHA-256 `b4537ed75f533f993f371954de47e42a793b8e5b0587577de7e27fb3e50696bd`. Integrity does not clear advisories. |
| `GOTOOLCHAIN=go1.26.9 go test ./internal/platform/auth ./cmd/server` / backend | Approved escalated focused compile exit 0 before and after gofmt; both packages report `[no test files]`. This establishes compile compatibility only, no auth correctness. |
| `GOTOOLCHAIN=go1.26.9 govulncheck -json ./... > RUN_PATH/evidence/dependency-vuln.json` / actual backend | First sandbox attempt failed loading Go cache; approved escalated run and post-gofmt rerun exit 0 as **JSON emission**, not clean security result. Actual scanner v1.7.0, source/symbol mode, DB modified 2026-10-08, Go 1.26.9; 15 finding records / 5 unique IDs remain. Full final JSON SHA-256 `15ea30ec5eb124d75299be0109f632cd4884648464525b06d9d5bfa02d4bf4fe`. |
| `gofmt -w` touched Go files; `git diff --check` | Formatting completed; diff check exit 0. No race/concurrency/performance/security-class final sweep or broad Testing-owned suite ran. |

### Remaining advisory disposition (actual graph)

All five IDs are **open** for bounded owning reconciliation, not accepted risk. The primary Go advisory pages and actual scanner traces were checked on 2026-10-10:

| ID / selected version / first fixed | Actual path and applicability | Disposition |
|---|---|---|
| [GO-2026-4771](https://pkg.go.dev/vuln/GO-2026-4771), pgx v5.8.0 / v5.9.0 | Memory safety in `pgproto3.Backend.Receive` / `Bind.Decode`. Scanner reports module/package level, **no application symbol trace**. The backend uses pgx as a client; affected parser invocation/exploitability was not established. | Affected module retained; no path-based clearance or version change. |
| [GO-2026-4772](https://pkg.go.dev/vuln/GO-2026-4772), pgx v5.8.0 / v5.9.0 | Memory safety in `pgproto3.Backend.Receive` / `FunctionCall.Decode`. Module/package only; no application symbol trace. | Affected module retained; no path-based clearance or version change. |
| [GO-2026-5004](https://pkg.go.dev/vuln/GO-2026-5004), pgx v5.8.0 / v5.9.2 | Placeholder confusion when using **non-default simple protocol** with a dollar-quoted SQL literal. Scanner traces `SanitizeSQL` → pgx QueryRow → new auth `ReadSession`; actual queries are parameterized and do not contain dollar-quoted literals. Runtime `DATABASE_URL`/query mode has not been constrained or verified, and T3 queries do not exist yet. | Current new query shape does not show trigger conditions; graph remains affected and future/configured path is open. No blanket non-applicability claim. |
| [GO-2026-5970](https://pkg.go.dev/vuln/GO-2026-5970), x/text v0.29.0 / v0.39.0 | Invalid UTF-8 can cause normalization infinite loop. Scanner symbol trace runs through x/text → pgx SCRAM client → pgxpool.New → existing `db.Open`; invalid input/provider path not reproduced. | Potential startup/DB-auth use path, affected graph. No clearance. |
| [GO-2026-6629](https://pkg.go.dev/vuln/GO-2026-6629), x/text v0.29.0 / v0.41.0 | Crafted input can panic in `secure/precis` Nickname profile. Scanner traces `precis.String` → pgx SCRAM client → pgxpool.New → `db.Open`; crafted input and configured DB-auth conditions not reproduced. | Potential startup/DB-auth use path, affected graph. No clearance. |

No standard-library, OIDC, OAuth2, or go-jose advisory appears among the actual scan's finding records. That statement is limited to this scanner snapshot and implemented source. The old 50-ID T2-001 matrix does not transfer as a clean bill; the new five IDs above are actual-graph evidence. Primary advisories state affected/fixed ranges, while source traces express possible reachability rather than exploit reproduction. The proposed fix versions are **advisory facts, not an Implementer-selected dependency change**; owning Techplan/security reconciliation must assess pgx/x/text compatibility, graph, and disposition before Build may commit dependencies or claim the gate passed.

## Verification scope confirmation

The required focused `govulncheck` supply-chain scan ran; it is not a final auth/security-class sweep. No race/concurrency, performance/load, security-class final sweep, or broad Testing-owned suite ran. Focused authored protocol-negative/person-upsert/session-lock tests were **not written or run** because the new graph blocker stopped the Build. Compile-only output must not be read as those tests passing.

## Contract check

- [ ] Current T2/G1 build target satisfied — partial code only; advisory gate and focused tests outstanding.
- [x] Bound Techplan/T2/manifest/API/T1 inputs matched the fresh invocation; no material contract assumption was silently reinterpreted.
- [x] Actual Go 1.26.9, exact auth pins, graph, module integrity, source scan, and focused compile were observed before any dependency commit.
- [ ] Surviving advisory disposition accepted by owning authority — five IDs remain open; no dependency commit.
- [x] G1 permission respected; no G2/G3, API/frontend/Product/Techplan/task/manifest write by this Implementer.

## Deferred / not tested here

Protocol-negative tests, person-upsert uniqueness/concurrency tests, session lock/revocation revalidation tests, migration apply/rollback, and restricted runtime DB role remain outstanding. Real Google login/client/secret/callback/consent, configured PostgreSQL concurrency and SCRAM input path, HTTPS/browser cookie behavior, provider failure handling, CSRF/browser integration, independent Testing, and Human evidence were not run. The partial code may need correction after independent review; it has no runtime acceptance. T3's durable authenticated person/session prerequisite is **not satisfied**.

## Flagged for Techplan / Testing

**F-T2-002 / B-T2-002 — new actual-graph pgx/x/text advisories block dependency acceptance.** Owner: Orchestrator routes a bounded dependency/security reconciliation to the relevant Techplan concern and Anhar if approval changes. The approved D8 candidate reconciled prior Go/jose findings, but approved parent §10 explicitly requires checking and reconciling any newly affected dependency/path before dependency commit. The current graph identifies five remaining IDs, including source traces through DB startup/auth and new session QueryRow. This Run cannot silently upgrade pgx/x/text (pgx was outside the approved candidate), accept residual risk, or claim T2 completion from compilation. Preserve partial backend files and this Run evidence for a fresh owning continuation; T2-001 remains terminal BLOCKED.

## Phase handoff

- **Outcome:** BLOCKED — actual dependency graph scan exposed new affected pgx/x/text selections; T2/G1 partial backend writes are uncommitted and not behavior-verified.
- **Result refs:** This `build-report.md`; `dependency-vuln.json` SHA-256 `15ea30ec5eb124d75299be0109f632cd4884648464525b06d9d5bfa02d4bf4fe`; `dependency-graph.txt`, `dependency-integrity.txt`, `auth-module-source.json`, and `changed-file-identities.sha256` in this Run evidence.
- **Findings:** F-T2-002 — five actual-graph advisory IDs above; old Go/jose finding set not reported in current scan.
- **Decision requests:** Bounded pgx/x/text version/compatibility/applicability disposition by owning Techplan/security concern, with required Human approval for material candidate change if applicable.
- **Blockers:** B-T2-002 affects T2 dependency commit and completion only. No T3 auth prerequisite can be declared ready; unrelated work is outside this Run.
- **Open / unverified:** Focused authored tests, migration and real provider/DB/browser behavior, independent Testing, and partial-code review remain open.
- **Recommended continuation:** Orchestrator reconcile the smallest dependency/security decision; after approval, fresh Build Run rechecks actual graph/source/integrity/scan/compile, finishes T2 tests/implementation, then routes T3 prerequisite assessment. Do not resume T2-001 or dispatch T3 from this report.
- **Context refs:** Fresh invocation, Approved Techplan §10/§12, refreshed T2 task, accepted manifest, bound T1 contract/review evidence, prior T2-001 blocker, and this Run's source scan/changed-file identities.
