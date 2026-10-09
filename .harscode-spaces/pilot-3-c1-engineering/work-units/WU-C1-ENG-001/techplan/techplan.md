# Tech Plan: C1 — Legitimate Organization representation

> Phase             : Techplan
> Ticket            : none
> Author            : `PARTICIPANT-C1-ENG-PLANNER-002`
> Participant ID    : `PARTICIPANT-C1-ENG-PLANNER-002`
> Role              : Planner
> Model             : `gpt-6-sol` — pilihan dispatch invocation, runtime model tidak terekspos independen
> Reasoning         : medium — konfigurasi dispatch invocation
> Session           : `SESSION-C1-ENG-PLANNER-002` — binding Run, bukan ID thread runtime
> Created           : 2026-10-09
> Updated           : 2026-10-09
> Target revision   : `524ef600c7f71246af6b71671d89c3c040fd9d44`
> Workflow revision : `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`
> Status            : Approved — Anhar, 2026-10-09; bounded G1–G3 implementation permissions authorized separately on 2026-10-09
> Approach          : Google OIDC + local session; preparation/confirmation; satu aggregate PostgreSQL, conservative guard, Owner-scoped inspection.
> Refs              : Stage 7 C1 handoff; Stage 5 confirmed behavior; Stage 6 requirements; Exploration `RUN-C1-ENG-EXPLORATION-001`; Solution Contract; WU-C1-ENG-001 / RUN-C1-ENG-TECHPLAN-002.

---

## 1. Background

C1 membentuk Organization + initial Kencleng Owner yang durably attributable tanpa membuktikan legal/external authority. Live baseline masih health endpoint, PostgreSQL connectivity, neutral frontend dan OpenAPI kosong. Solution Contract menyelesaikan pilihan material H1–H3; revisi ini menurunkan rencana eksekusi dari Stage 5/6 dan kontrak tersebut, bukan mengubah Product Authority.

Input identity cocok: Solution SHA-256 `c136e673937c9ac1a583ecfa6a98ee630e1b0f6b42b78d373876243e55b01482`; Draft sebelumnya `188a55d8a6b951bdbc8d29063689606c082969ca8dca52861d0256d84b54c878`. Preparation baseline `3e123bcece561f1d0a181b68c85e2677dcfa37eb` → observed HEAD di atas hanya mengubah koordinasi Run, bukan solusi/Draft/production anchors. Workflow revision cocok dan tree bersih. Working-tree coordination edits yang sudah ada dipertahankan. Provenance/self-check: [Run evidence](../runs/RUN-C1-ENG-TECHPLAN-002/evidence/synthesis-check.md).

## 2. Scope

**In scope:**

- Google OIDC backend, durable local person/opaque session, preparation tanpa hak, explicit confirmation, aggregate Organization dengan inline initial Owner.
- Durable conservative global guard, restricted operator input/privileges/runbook dan enforcement saat commit.
- Shared OpenAPI/generated types, same-origin client, display-name form, empat konsekuensi, truthful failure/uncertainty, my-list/durable detail.
- Additive migrations, constraints, scoped reads, real PostgreSQL/security/concurrency/provider/rendered verification obligations.

**Out of scope (explicit):**

- Organization review/status taxonomy, legal proof/adjudication, matching/duplicate detection/uniqueness, conflict resolution/merge/transfer, Campaign/donations/fund-use verification.
- Staff/general role engine, invitations/additional Owners, local passwords/MFA/account linking/recovery, refresh tokens, signed local JWT/custom signing/encryption/HMAC core.
- Saved user drafts/general recovery journey, blind retries/new create after unknown commit, extra services/outbox/object storage, multi-instance pending-login coordination.
- Run ini tidak menulis production/API executable/migration, menjalankan tests/build/runtime, memberi approval atau mengubah upstream/coordination projections.

## 3. Requirements

| ID | Requirement | Source / evidence |
|---|---|---|
| Q1 | C1 may begin without external-authority proof; successful establishment and initial Owner relationship must be attributable to a person. | Stage 5 B1; Stage 6 ER1.1–ER1.3 |
| Q2 | Before effect, communicate Organization establishment, initial Owner formation, internal-only authority meaning, and Organization-provided information status. | Stage 5 B2; Stage 6 ER2.1–ER2.2 |
| Q3 | Successful establishment means a valid Organization context plus valid attributable initial Owner relationship; neither side may appear successful alone. | Stage 5 B3; Stage 6 ER3.1–ER3.2 |
| Q4 | The resulting Organization–Owner meaning and attribution remain durably inspectable after the success moment. | Stage 5 B4; Stage 6 ER4.1–ER4.2 |
| Q5 | Establishment must not upgrade Organization-provided information/evidence into reviewed, independently verified, or Campaign-curated truth. | Stage 5 B5; Stage 6 ER5.1–ER5.2 |
| Q6 | Missing attribution or incomplete/failed formation cannot produce successful establishment meaning. Recovery is not required. | Stage 5 B6; Stage 6 ER6.1–ER6.2 |
| Q7 | If a relevant unresolved representation conflict is already known, C1 must not report normal uncontested success or grant a new initial Owner as though uncontested. C1 does not define how such conflicts are detected or resolved. | Stage 5 B7; Stage 6 ER7.1–ER7.2 |

## 4. Rules & Validation

- **R1 — Attributable success:** Given an establishment request, when it succeeds, then durable state links the establishing person, Organization context, and initial Owner relationship; without a trusted person attribution, it cannot succeed. (Q1)
- **R2 — Consequence before effect:** Given the establishment experience, when the consequential action is submitted, then the required consequences were made available before the state becomes effective; external authority is not claimed or required. (Q2)
- **R3 — Coherent formation:** Given an attempt to establish C1, when either Organization or initial Owner formation is invalid or fails, then neither is exposed as successful; success is observable only when both are valid. (Q3, Q6)
- **R4 — Durable inspection:** Given a successful establishment, when the person later inspects its result, then the Organization context, their Owner relationship, internal authority meaning, and the establishment transition’s non-review/non-verification effects remain determinable beyond transient messaging. (Q4)
- **R5 — Truth-class preservation:** Given Organization-provided information/evidence, when C1 succeeds, then its truth class is not changed to reviewed, externally verified, or curated by that transition. (Q5)
- **R6 — Incomplete is not success:** Given missing attribution or failed/incomplete Organization + Owner formation, when processing ends, then no successful establishment meaning is persisted or returned; any exposed incomplete state remains distinguishable. (Q6)
- **R7 — Known conflict boundary:** Given a relevant unresolved representation conflict already known to the system, when C1 processes an establishment, then it does not produce normal uncontested-success meaning or grant initial Owner authority as though representation were uncontested; C1 performs no new matching/adjudication. (Q7)
- **R8 — Server-owned privileged authorization:** Given a request to create an initial Owner relationship, when the backend evaluates the grant, then it enforces the explicitly authorized actor/privilege rule server-side; frontend visibility is not security authority. (Root `AGENTS.md`, High-risk write boundary)

- **R9 — OIDC trust:** Fixed issuer/client, verified RS256 signature/keys, exact issuer, audience/azp, expiry/nonce, browser-bound one-time state + S256 PKCE; pending login ≤10m, `openid` only. Verified `(issuer, subject)` unique/case-sensitive; email/domain/body principal never grants Owner. Provider/config/key failure gives no fake session. (Solution §5)
- **R10 — Session/browser:** 32 random bytes opaque bearer, SHA-256 digest-only persistence, absolute 8h expiry, login rotation/logout revocation. Production host-only `__Host-kencleng_session`, HttpOnly/Secure/Lax/Path=/; distinct insecure cookie only explicit development localhost:8080. Mutation exact Origin + JSON + session-bound CSRF header, no credentialed CORS; fixed callback redirect/no-store/secrecy. Recheck validity under session lock during confirm. (Solution §§5,9)
- **R11 — Concurrent/replay/uncertainty:** READ COMMITTED lock session → guard → same-person preparation; one committed result per preparation. Logout uses session lock, rotation deterministic session order; guard writer locks guard. No external call in transaction. Known rollback may bounded-retry same receipt; possibly committed result resolves through read/replay same receipt, no definite failure/local success/blind replacement. Committed replay survives receipt expiry/new hold and does not grant again. (Solution §9)
- **R12 — Guard control:** Durable singleton initialized closed; closed/missing/unreadable permission denies new grant; hold between prepare/confirm enforced. Runtime SELECT + UPDATE(id) only on constrained id=1, no state/audit/DDL/role bypass. Operator change + append-only audit atomic; reopening requires external owning evidence for all outstanding holds. Inspection/historical replay available on hold; no matching/adjudication. (Solution §8)
- **R13 — Input/private contract:** Trim name, 1–200 Unicode scalar values, reject controls, escaped rendering; repeated names valid. Strict authority-bearing payload rejection, stable safe errors/private-no-store, bounded scoped stable pagination; no cross-person cache/secret leakage. (Solution §§4,7,10)
- **R14 — Evidence/gates:** Exact Techplan approval and separately recorded G1–G3 before protected writes. Pin/source-review/dependency-integrity/vulnerability checks before dependency commit. Real provider/DB/operator/rendered evidence distinguished from simulations and compilation. (Solution §§11–12; root AGENTS)

## 5. Decision Log

| ID | Decision / option | Status | Rationale / consequence |
|---|---|---|---|
| D1 | One successful Organization + attributable initial Owner consequence. | Binding B1–B7 | Internal authority only; no Product meaning changed. |
| D2 | Go monolith, one PostgreSQL aggregate/transaction; inline non-null initial Owner; session/preparation/result/guard co-owned. | Chosen — Solution §§4,9 | Earlier conditional same-DB assumption settled; no second service owns success. |
| D2-alt | Staged Organization then Owner/saga/separate Owner service. | Rejected | Risks half-success; optional recovery becomes unnecessary complexity. |
| D3 | Google OIDC → immutable local person → opaque local session. | Chosen — H1, §5 | No provider token in browser authority; G1 pending separately. |
| D3-alt | Local password auth, body principal/dev headers/production fixtures/email-domain authority. | Rejected | Adds unnecessary credential ownership or spoofed privilege. |
| D4 | Real durable conservative global guard, restricted operator input. | Chosen — H2, §8 | Accept broad availability hold rather than invented relevance policy; G3 pending. |
| D4-alt | Always-clear predicate/ignore conflict until future source; name matching/legal adjudication. | Rejected | No real input/enforcement or out-of-scope policy. |
| D5 | Display name → sign-in if needed → four-consequence summary → final confirm → detail; my-list later. | Chosen — H3 | Material flow/field settled; copy/layout/route grouping ordinary freedom. |
| D5-alt | Re-raise minimum field as Product blocker; legal proof/forced checkbox; public/browser-only inspection. | Rejected | Stage6/7 grant field freedom; durable scoped inspection required. |
| D6 | Frozen same-person preparation UUID = idempotency identity, version c1-v1. | Chosen — §§7,9 | Separate disclosure opportunity/effect, resolve lost response without saved-draft product. |
| D6-alt | Single create without preparation/replay, blind new create after timeout. | Rejected | Uncertain commit/duplicate authority unsafe. |
| D7 | Organization-provided provenance + facts of establishment effects. | Chosen — §7 | No `verified`/`pending_review`/review taxonomy or broader current review-state claim. |
| D8 | Pinned OIDC/OAuth modules and application validation supplement library defaults. | Selected for approval — §10 | Metadata-compatible with Go baseline; compile/security scan pending. Never weaken trust for dependency convenience. |
| D9 | Typed centralized same-origin client; local form, private query state; no refresh/BFF auth. | Selected for approval | React client guidance applies centralization/CSRF. Its bearer/refresh example N/A: specific solution has opaque cookie, no refresh. JWT signing-key/refresh guidance conditional N/A for local session, provider verification remains mandatory. |
| D10 | Separate future guard runbook, linked from spine. | Selected for approval | Own execution/rollback/monitoring lifecycle, rules §3. No child task files or independent operational authority created now. |

## 6. Backward Compatibility

- Additive schema from baseline without C1 data; no historical migration import/destructive conversion. Preserve health route.
- OpenAPI currently `paths: {}`. §8 is proposed execution contract; coordinated source/bundle/types batch precedes consumers.
- Retain committed receipts and referenced persons; cleanup only expired uncommitted preparations/pending logins. Future consequence-version changes preserve historical replay and redisclose new intent.
- Auth-enabled runtime must use restricted DB credential, not Compose bootstrap/owner. Initial guard closed. Rollback closes new establishment and preserves durable success/receipts; never default `migrate-down -all` on populated environment.

## 7. Edge Cases & Risks

| ID | Risk / edge case | Likelihood | Severity | Mitigation / accepted exposure |
|---|---|---|---|---|
| RISK-1 | Partial persistence/concurrent confirmation/response loss. | Medium | High | R3/R6/R11; unique receipt, DB fault/race tests, authoritative replay. |
| RISK-2 | Initial Owner/IDOR privilege escalation. | Medium | High | Exact server scope/constraints; G2; bypass requests. |
| RISK-3 | New OIDC/session integration, fixation/leakage. | Medium | High | H1 settled; G1, R9/R10, protocol negatives + real login. Pending-login restart requires new login. |
| RISK-4 | Display label interpreted as legal identity/uniqueness. | Medium | High | H3 settles technical field validity; duplicates allowed; no legal badge. |
| RISK-5 | Unrecorded conflict/guard-close race/runtime privileged credential. | Medium | High | H2/R12; operator receipt/all-holds evidence, real privilege tests. Baseline bootstrap user is not suitable application runtime. Human awareness outside recorded input not detected automatically. |
| RISK-6 | Truth upgrade in data/copy/inspection. | Medium | High | R4/R5, transition facts not current review taxonomy; rendered acceptance. |
| RISK-7 | Singleton serializes confirmations/broad availability hold. | Expected | Medium | H2 accepts pilot tradeoff; short transaction/bounded timeouts. Scale/multi-instance change needs reconciliation; no broad load suite mandated. |
| RISK-8 | Private state cached across persons/stale browser role after logout. | Medium | High | private/no-store; session-scoped query keys/cache clear; two-person browser check. |

## 8. Interface Contract

### Persistence/data shape

Logical names are proposed, not existing schema; physical names/SQL mechanics remain Build freedom preserving constraints.

| Record | Material content / constraints |
|---|---|
| persons | Server UUID; case-sensitive unique `(issuer, subject)` verified mapping. No email/domain linking; retain while referenced. |
| sessions | Digest of 32-byte random opaque token, person FK, expiry/revocation, separate CSRF token. Provider tokens discarded after verification. |
| organization_establishment_preparations | Server UUID; person FK; frozen validated name, created/expires, c1-v1, nullable committed Organization reference. Not Organization/Owner/hak. |
| organizations | Server UUID/name; non-null initial_owner_person_id + established_by_person_id FKs, equality CHECK; timestamp; unique preparation FK; version; information_source=organization_provided CHECK. No ownership mutation API. |
| c1_establishment_guard + append-only changes | Singleton id=1 CHECK, permission closed initially, revision/time; change records operator/reason/evidence refs atomically with update. Runtime cannot change state/audit. |

Result linkage and unique preparation FK enforce one result. Exact FK ordering/deferrability may vary but no relaxation of attribution/atomicity. No legal identifiers/uploads/review enum/Staff table. Name trims whitespace, 1–200 Unicode scalar values, rejects controls, escaped display; duplicate names allowed, references UUID.

### HTTP / external boundary

Paths below are backend/OpenAPI paths; browser uses `/api`, stripped by Caddy. Development auth browser entrypoint http://localhost:8080, callback `/api/auth/google/callback`; direct UI :3000 not auth origin. Production HTTPS fixed origin, secret configuration backend-only.

| Operation | Input / authority | Result / effect |
|---|---|---|
| GET /auth/google/start | Fixed local destination; browser-bound one-time state/nonce/PKCE | Google redirect; no Organization/Owner. |
| GET /auth/google/callback | Code/state + initiating browser, verified identity | Map person/rotate local session; fixed C1 form redirect. Failure grants no session/Owner. |
| GET /me | Valid session | `{person_id, csrf_token}`, no provider subject/token/email. |
| POST /auth/logout | Session/Origin/CSRF | Revoke session + expire cookie; 204. |
| POST /organization-establishments | Session/Origin/CSRF; `{display_name}` | 201 receipt `{id,display_name,expires_at,consequence_version,consequences}`; preparation only. |
| GET /organization-establishments/{id} | Same-person session | 200 union `{state:prepared,...receipt}`, `{state:expired,id}`, `{state:established,organization}`; outside-person/unknown 404. |
| POST /organization-establishments/{id}/confirm | Same-person session/Origin/CSRF; `{consequence_version}` only | 200 `{state:established,organization}` after confirmed commit or same-result replay. No new name/person/role/target/override. |
| GET /me/organizations | Session | `{items:OrganizationView[],next_cursor:string|null}`; scoped stable timestamp/UUID keyset, bounded page size. |
| GET /organizations/{id} | Session + object-level Owner scope | 200 OrganizationView; outside-person/unknown 404. |

`consequences={creates_organization_context:true,assigns_you_initial_owner:true,owner_authority_scope:kencleng_internal,organization_information_source:organization_provided}`; `consequence_version=c1-v1`. Uncommitted stale/unknown version needs fresh preparation + disclosure. Committed receipt replay requires valid same-person session but is historical, not new grant subject to hold/receipt expiry.

`OrganizationView={id,display_name,established_at,established_by_person_id,initial_owner_person_id,your_role:owner,owner_authority_scope:kencleng_internal,information_source:organization_provided,establishment_effects:{organization_review_performed:false,external_authority_verified:false,campaign_curation_performed:false}}`. Effects describe what establishment did, not all current/future review outcomes; no public person directory/provider subject/legal badge.

Safe localized error envelope `{error:{code,message,request_id}}`. Reads/mutations/errors private/no-store. Fixed local callback error destination, no arbitrary return URL/referrer/code leakage; no raw SQL/provider error/stack/path/secret.

| HTTP / code | Meaning / client behavior |
|---|---|
| 400 invalid_request | Invalid JSON/path/body/unknown authority-bearing field; no mutation. |
| 422 invalid_organization_input | Name correction; no success. |
| 401 authentication_required | Sign-in; no inferred success. |
| 403 request_not_allowed | Origin/CSRF denied before processing. |
| 404 not_found | Unknown/outside-person identical shape. |
| 409 consequence_changed | New preparation + redisclosure for uncommitted intent. |
| 410 preparation_expired | Uncommitted expired; new preparation allowed. Committed receipt retained. |
| 503 establishment_unavailable | Closed/missing/unreadable guard; temporary unavailability, not dispute verdict for every candidate. |
| 503 processing_unavailable | Confirmed pre-commit DB/lock/service failure; safe same-command retry. |
| 503 outcome_unknown / transport loss | Retain preparation ID, read/replay same command; no definite rollback or blind replacement. |
| 429 too_many_requests | Bounded login/preparation abuse controls; thresholds operational tuning. |

### Cross-layer/security boundary

Backend owns identity, grant, persistence/guard and inspection. Same-person scoped queries prevent another actor locking/learning preparation. Auth alone cannot take existing Organization or assign another Owner. Centralized typed client sends cookie credentials and `X-CSRF-Token` on JSON mutations; no bearer storage/refresh or automatic mutation retry. Callback uses state/nonce/browser binding rather than ordinary mutation Origin rule. Frontend toast/cache/MSW fixture never becomes authority.

## 9. Architecture / Plan

### Tasks and scoped writes

Future approved Build tasks only; this Run does not self-dispatch. Backend/frontend production edits stay separate; shared API and operational docs require the explicit concern coordination below.

| Task | Owned writes / dependencies | Completion evidence / gates |
|---|---|---|
| T1 Shared contract | API concern owner coordinates §8 with backend/frontend; source, validate/bundle, generated types; no consumer production writes in this batch. | Exact Techplan approved; contract captures auth/state/errors/scope; no G1–G3 implementation permission implied. |
| T2 Identity/session | Backend platform auth/config/person-session migrations/HTTP auth+CSRF/dependencies. Depends T1. | G1 before core writes; G2/G3 too if authority-bearing schema touched. Protocol negatives/person-upsert/session-lock checks authored. |
| T3 Aggregate/guard | Backend Organization service/repository/handlers, additive aggregate/preparation/guard/audit schema and restricted credentials/control. Depends T1/T2. | G2+G3 before protected writes; atomic confirm/scoped reads/privilege/DB fault-race tests authored. |
| T4 Guard operation docs | Scoped new docs/project/c1-establishment-guard-runbook.md + backend setup/config examples; coordinate T3. | Parameterized operating path/privileges/audit/all-holds reopen/rollback. G3 for executable control/provisioning; no legal resolution authority. |
| T5 Frontend C1 | Frontend only after T1; labeled contract-faithful MSW possible, real T2/T3 integration later. Form/summary/confirm/uncertainty/list/detail/nav. | Generated types, private server state, no local role authority, rendered states. Contract/backend defects route to owner, not cross-stack edits. |
| T6 Review/verification | Build authorship checks; separate Code Review reasons across contract/G1–G3; independent Testing runtime checks; Human semantic/gate acceptance. | Final evidence §12, real provider prerequisite, no delivery claim from compile. |

### Runtime order

1. Form display name without legal proof; sign-in if needed before server preparation. Optional tab-local name through redirect; no token/role authority there. Callback/login/page-load never automatically confirms.
2. Preparation freezes person/name/version/expiry. Render all four consequences + frozen name before final action. Name edit means new preparation + redisclosure. Disable action while preparing, summary not rendered, or confirmation unresolved.
3. Short READ COMMITTED transaction locks session → guard → preparation, with same-person lookup. Recheck session revocation/expiry/person under lock; logout locks same session, rotation uses deterministic order and never enters Organization path. No network calls inside transaction.
4. Committed receipt returns stored result without another grant or rejection solely due new hold/expiry. Uncommitted receipt enforces guard/expiry/version/name, inserts complete aggregate and links result, commits before success. Historical authoritative read resolves result even when new-grant permission is unavailable; mutating path must preserve selected lock order.
5. Known deadlock/serialization rollback may bounded-retry same preparation; possibly committed acknowledgment/encode/disconnect outcome remains unknown to caller until same-person read/replay. No blind retries/replacement. Cleanup keeps committed receipts/attribution.
6. My-list → detail authoritative after restart/re-login, available during hold. Query cache and CSRF context cleared on logout/re-login. Distinct loading/invalid/login-required/expired/hold/definite failure/unknown/empty/success presentations.

### Operator runbook / privileges (T3/T4)

- Owner/migration credential provisions constraints/audit and initializes guard closed. Runtime non-owner/non-superuser, no schema CREATE/role management. Baseline Compose bootstrap credential is infrastructure provisioning, never auth-enabled application runtime.
- Runtime guard SELECT + column UPDATE(id) solely for FOR UPDATE; CHECK id=1 prevents enabling by key edit. No state UPDATE/INSERT/DELETE/audit write/constraint alteration/ownership. Test actual DB credentials.
- Separate restricted operator performs controlled parameterized transaction: lock singleton, expected revision check, update permission/revision, append operator ref/time/reason/evidence ref, commit together. A narrowly executable DB routine with fixed search_path, migration owner and no PUBLIC execution can enforce this without operator arbitrary state/audit changes. Exact mechanics remain G3 Build freedom, preserving append-only evidence/no bypass.
- Before initial open: named operator/input channel for all accepted conflict knowledge, inspect outstanding holds, record explicit environment permission. Opening states available platform knowledge, not independent verification of real-world conflict absence.
- Received unresolved conflict closes guard before further grants. Record safe evidence refs without raw claimant/PII. Track all holds; clearing one is insufficient. Reopening requires external owning decision/evidence that all outstanding holds no longer apply; C1 does not adjudicate.
- Monitor permission/revision/audit and safe availability/error counts. Failed control operation retains prior committed state; if safe closure cannot be confirmed stop grants operationally. Missing/unreadable guard fails closed. Repair privileges/input path before opening, never override in application.
- Rollback/incident: audited close + disable new confirmation if needed; preserve Organizations/Owner/committed receipts for inspection/replay. No automatic reopen/destructive rollback/retroactive revoke. New asynchronous conflict source must atomically update guard or trigger solution reconciliation before compliance claims.

### Continuation recommendation

Independent Techplan review: **Recommend** — auth/session/PII, privileged Owner/object scope, API and guard/transaction consistency meet current Complex signals. Reviewer needs fresh independent context, exact plan hash + Solution Contract/Exploration/authorities. This Run self-checks only.

Decomposition: **Consider after approval** — contract/backend security-data/frontend rendered experience/operator runbook offer useful independent context/review boundaries. Length alone is not reason. Spine keeps all material obligations; no child task/manifest created now. Planner owning semantics generates report only after any invoked review/resolution converges and exact revision enters Human approval gate; no report now.

## 10. Implementation Details

| Anchor | Why relevant | Intended change / precedent |
|---|---|---|
| backend/cmd/server/main.go — run/loadEnvironment/healthz | Sole health route/config/composition | Compose auth + Organization routes, safe config validation, preserve health. |
| backend/internal/platform/db/db.go — Open | Connectivity-only pgxpool | Reuse pool with narrow repo/transaction dependencies and restricted runtime credential. |
| backend/internal/domain/, migrations/ — .gitkeep | No live C1 schema/domain | Add Organization domain/platform auth/additive constraints; physical names ordinary freedom. |
| backend/go.mod — go 1.24.8; pgx/v5 v5.8.0 | Current module/compiler substrate | Add exact auth pins below; derived go.sum in Build. |
| backend/Makefile — verify/security/migrate-up | Existing tooling | Use checks + proposed explicit isolated integration route; no destructive migrate-down default. |
| api/openapi/index.yaml — paths: {}; api/package.json — validate/bundle | Source contract + tooling | §8 source/bundle/types coordinated, no parallel handwritten model. |
| frontend/app/page.tsx/layout.tsx; components/README.md | Neutral home, fonts/layout; broad contracts empty | C1 routes/nav/route-local components; register shared primitives only if demonstrated. |
| frontend/package.json — verify/build/test:browser | Existing Vitest/RTL/MSW/Playwright/Next | Reuse locked tooling/typed client/local form+query; no additional UI dependency required. |
| Caddyfile — handle_path /api/*; docker-compose.yml — postgres/caddy; Makefile — up-podman | Same-origin prefix stripping/Postgres16/localhost8080 | Preserve topology; coordinate restricted provisioning T3/T4; MinIO unused. |
| .harscode-spaces/.local-config.yaml — runtime | Podman/podman-compose | Probe actual configured route; Docker CLI absence is not runtime unavailability. |

**Dependency selection — read-only primary evidence checked 2026-10-09:**

- Pin `github.com/coreos/go-oidc/v3 v3.16.0` + `golang.org/x/oauth2 v0.34.0`. [OIDC go.mod](https://raw.githubusercontent.com/coreos/go-oidc/v3.16.0/go.mod) and [OAuth2 go.mod](https://raw.githubusercontent.com/golang/oauth2/v0.34.0/go.mod) require Go1.24. OIDC requires OAuth2≥v0.28.0 and go-jose/v4 v4.1.3; preserve go-jose≥v4.1.3, whose [module](https://raw.githubusercontent.com/go-jose/go-jose/v4.1.3/go.mod) requires Go1.24. Metadata compatible with baseline Go1.24.8; not compile/security evidence or latest-version claim.
- [Tagged verifier source](https://raw.githubusercontent.com/coreos/go-oidc/v3.16.0/oidc/verify.go) leaves nonce/authorized-party checks to caller and allows a Google issuer alias. Configure exact client/RS256, all skip flags false; after verified token require exact issuer `https://accounts.google.com`, exact nonce and applicable azp. Use protocol-library S256 helpers; no handwritten JWT/signature implementation.
- Build inspects tagged library APIs/source and resolved graph; `go mod verify` + `govulncheck ./...` before dependency commit, actual toolchain compile/tests and resolved pins recorded. Vulnerable/incompatible selection requires bounded dependency reconciliation, never weaker trust/silent toolchain upgrade. Existing frontend/API lockfiles retained; no dependencies installed by Planner.

Applicable current Harscode: Go authorization/IDOR, privilege separation, token lifecycle (conditional D9), dependency/supply chain/integration setup; REST CSRF/cookie; PostgreSQL locking; React server/client boundary, centralized client (D9), visual verification. Current source files read, not prior paraphrases. Build reopens live anchors. Naming/narrow interfaces/SQL organization/preparation TTL/timeouts/list size/thresholds/copy/layout remain ordinary freedom; login≤10m/session8h and material trust/atomicity/guard/scope fixed.

## 11. Files Changed / Files NOT Changed

Expected future implementation, not this Run's writes:

| File / area | Change type | Description |
|---|---|---|
| api/openapi/ + derived api/openapi.yaml/types | Shared coordinated | T1 §8 contract. |
| backend/platform auth (proposed), domain Organization (proposed), HTTP composition | New backend | T2/T3; exact file names Build freedom. |
| backend/migrations, go.mod/go.sum, config examples/verification tests | Additive backend | Person/session/aggregate/preparation/guard/audit/restricted provisioning; G1–G3. |
| docs/project/c1-establishment-guard-runbook.md + scoped backend operating docs | Explicit coordinated docs | T4 execution/privileges/audit/holds/reopening/rollback. |
| frontend C1 routes/nav/client/generated types/route-local components/tests | Frontend-only batch | T5 material states; shared registry only when broad contract actually emerges. |

| File / area intentionally untouched | Why |
|---|---|
| All protected Product/Pilot pre-engineering authorities | Upstream read-only; route contradictions to owner. |
| Control Tower/Space/WU/events/invocation projections | Orchestrator reconciliation, outside Planner writes. |
| Exploration/Run001 evidence; Solution Contract | Immutable historical/pinned inputs. |
| Money movement/MinIO/general roles/review engine | Outside C1. |
| Harscode guidance | Read-only/proposal-gated. |

## 12. Testing Checklist

Future ownership, **not executed evidence**. Build runs authored focused checks enough to establish executability. Independent Testing owns final runtime correctness. Code Review reasons on diff/gates, targeted reproduction only when needed, no ceremonial broad final suite.

| Rule | Verification / evidence | Primary owner | Why this is worth running / risk if skipped |
|---|---|---|---|
| R1, R9 | Real Google login maps stable person across sessions; real DB concurrent first-login unique upsert. No token/credential payload evidence. | Testing | Fakes cannot prove configured provider integration/durable identity. |
| R9 | Labeled protocol negatives: forged issuer/audience/azp/signature/alg, expired token, nonce/state/browser mismatch, code replay, missing key/provider failure, body/email/domain spoof. | Testing | Library defaults alone do not prove selected trust; simulated provider explicitly labeled. |
| R2 | Callback/preparation creates no Organization/Owner; invalid version cannot commit; name edit rediscloses; no auto confirm after callback/load. | Testing | Consequence opportunity before effect could be silently bypassed. |
| R2, R4, R5 | Rendered Human walkthrough: all four B2 meanings before action; detail/list internal Owner and provenance/effects intelligible without legal/review inference. | Human | Semantic intelligibility requires conscious acceptance, not compilation. |
| R3 | Isolated real PostgreSQL constraint/migration tests; injected failure between aggregate insert/result linkage leaves neither; commit yields complete state; FK/non-null/equality/unique failures. | Testing | Unit-only mocks cannot prove DB atomicity/constraints. |
| R4, R8 | Fresh browser context/login/restart navigates list/detail without original response; second-person list/detail/preparation/result/confirm denied, unknown/outside same404, absent session401. | Testing | Durable discovery/IDOR cannot rely on UI gates. |
| R5 | Persisted organization_provided/effects through confirm/replay/read; no review/verified/curation taxonomy in contract/code. | Testing | Silent truth upgrade changes downstream meaning. |
| R6, R11 | Real DB concurrent confirms one result; replay committed after expiry/hold; changed input denied. Inject commit acknowledgment loss/post-commit encode/disconnect; resolve same result by read/replay. Definite rollback/unknown UI distinct, no blind replacement. | Testing | Transport loss otherwise causes duplicate grants or false success/failure. |
| R7, R11 | Guard-close/confirm races with controlled barriers; hold after preparation; session logout/revoke race and expired/revoked under-lock deny; inspect lock/commit ordering. | Testing | Default isolation/Go race cannot prove DB shared decisions. |
| R7, R12 | Real default closed/missing/unreadable guard; audited atomic operator change; runtime cannot update state/audit/constraints/roles/insert/delete but can lock singleton. User bypass denied; hold permits historical inspection/replay. | Testing | Mock clear predicate/superuser tests hide privilege bypass. |
| R12, R14 | Named operator/input channel, initial permission, outstanding holds/owning reopen evidence and runbook rehearsal with safe refs. | Human | Software cannot prove external knowledge receipt or adjudicate resolution. |
| R10 | Cookie flags/dev-only exception/production HTTP rejection, rotation/fixation/8h expiry/logout; Origin/CSRF denial/no credentialed CORS/fixed redirect; callback log/no-store secrecy. | Testing | Auth/browser bypass compromises every Owner surface. |
| R13 | Unicode trim/empty/200/201/controls, duplicate names/escaped render, strict bodies, safe envelopes/no-store, bounded keyset ties; two-person cache isolation after logout/login. | Testing | Validity/injection/enumeration/private-cache leaks. |
| R2, R4, R5, R6, R13 | Playwright narrow/wide rendered sample: long name, keyboard/focus, loading/invalid/login/expired/hold/definite failure/unknown/empty/success. Inspect hierarchy/overflow/reachable actions vs Sunlit Editorial/patterns. | Testing | RTL/compile cannot prove visible consequences or responsive usability. |
| R14 | Dependency source/graph review, go mod verify/govulncheck before dependency commit; focused go/httptest; API validate/bundle/types; frontend verify/build/authored tests executable. | Build | Proportional edit-loop/toolchain/supply-chain evidence. |
| R8, R14 | Exact approval + separate G1–G3 permissions before protected writes, limitations recorded. | Human | Product/solution/model approval is not implementation authorization. |

**Future execution route:** backend `make verify` (lint/unit/race/contract/security), `go mod verify`; proposed bounded isolated `go test -tags=integration ./...` tests for actual PostgreSQL behavior (integration artifacts not present today). Use configured Podman/podman-compose (`make up-podman`) with disposable state/distinct runtime-operator-migration roles. Never production data/bootstrap-superuser-only proof. Deterministic cleanup/bounded timeout; after abort inspect leaked processes/resources. Race detector for process-local pending-login/shared state; real DB races separately required. No broad load suite for bounded pilot.

API `npm run validate`/`npm run bundle`; frontend `npm run verify`, `npm run build`, targeted authored `npm run test:browser`. `--passWithNoTests` is not C1 evidence. Check actual tool/environment availability later. Real Google client/secret/callback/consent/test user/network/Human interaction prerequisites; production Secure-cookie evidence needs HTTPS besides dev test. Missing prerequisite = not run, never simulated login relabeled real attribution.

### Test Focus Pointer

Paths relative WU-C1-ENG-001; Solution §11 refines oracles, not earlier executed Exploration evidence.

| Area | Why sensitive | Evidence anchor from Exploration | Still relevant post-synthesis? |
|---|---|---|---|
| Initial Owner/person/auth | Privileged grant and scope | runs/RUN-C1-ENG-EXPLORATION-001/evidence/stage-3-solutioning.md#material-questions-and-constraints (1,3) | Yes — D3/R8–R10; Solution §§5–6. |
| Organization + Owner atomicity | Persistent shared success/rollback | runs/RUN-C1-ENG-EXPLORATION-001/evidence/stage-3-solutioning.md#organization--initial-owner-consistency | Yes — D2/R3/R6/R11; Solution §9 locks/unknown result. |
| Known conflict | Privileged guard/actual input | runs/RUN-C1-ENG-EXPLORATION-001/evidence/stage-3-solutioning.md#known-unresolved-representation-conflict | Yes — D4/R7/R12; Solution §8 real guard. Detection/matching N/A — D4-alt/outside C1. |

Refined sensitive concerns not explicit in earlier Exploration: nonce/azp/issuer library defaults, UPDATE(id) privilege, session-lock revocation race, indeterminate commit, private-cache isolation. Preserved in R6/R9–R13/checklist with Solution §§5,8–11/primary source; no claim Exploration already specified/tested them.

## 13. Open Items

### Active — needs external input or verification

5. **Provider/runtime/operator/dependency evidence.** Fixed Google configuration/registered callback/consent/test user/real login, HTTPS, configured Podman/PostgreSQL isolated roles, named guard input/operator/initial open evidence; actual module integrity/vulnerability/compile. Owners Build/Testing/operator/Human per §§9,12. Not tested here; no credentials persisted. No unresolved material planning blocker identified.

### Resolved — retained as decision history

1. ~~**Valid Organization context (prior Active 1).**~~ **RESOLVED — H3 + Solution §§2,4:** display name/UUID/attributed Owner/provenance technical validity within Stage6/7 deliberate field freedom; not missing Product/legal validity.
2. ~~**Trusted actor/auth boundary (prior Active 2).**~~ **RESOLVED — H1 + Solution §§5–6:** Google OIDC/local person/session direction; G1 separately pending permission.
3. ~~**Transaction boundary (prior Active 4).**~~ **RESOLVED — Solution §§4,9:** single PostgreSQL ownership selected, no conditional second service.
4. ~~**Known conflict input (prior Active 5).**~~ **RESOLVED — H2 + Solution §8:** durable restricted operator guard/default closed; G3/operational evidence separate; no matching/adjudication.
5. ~~**Recovery required?**~~ **RESOLVED — No general recovery product.** B6/ER6 optionality preserved; same-receipt read/replay specifically selected for transport uncertainty (D6).
6. ~~**Real-world representation proof/external prerequisite?**~~ **RESOLVED — No.** B1/B5/ER1.3/ER2.2; internal Owner only.
7. ~~**Conflict matching/legal adjudication?**~~ **RESOLVED — Outside C1.** B7/ER7.2; D4 avoids inventing policy.
8. ~~**Exact execution contract approval (prior Active 1).**~~ **RESOLVED — Anhar, 2026-10-09:** “ya aku setuju dengan techplan nya bro” approves the exact reviewed/reported source SHA-256 `577673c2b03bc93526362a848636699ebd888b8917519bae65b1adde15860954`. [Approval event](../events.md) records lifecycle reconciliation only; G1–G3 and provider/runtime/operator/dependency evidence remain separately pending. No material plan semantics changed.
9. ~~**G1 implementation permission (prior Active 2).**~~ **RESOLVED — Anhar, 2026-10-09:** “ya aku izinkan” in response to the separate exact G1–G3 implementation request authorizes Google OIDC code/PKCE/state/nonce/token verification, local person mapping, opaque issuance/digest/expiry/rotation/revocation, cookies including the explicit localhost exception, CSRF and secret handling, within the approved Techplan and Solution Contract §12. [Authorization event](../events.md) owns the decision; no expanded auth/key-handling scope or runtime verification is implied.
10. ~~**G2 initial Owner authorization (prior Active 3).**~~ **RESOLVED — same explicit Human decision:** authenticated initiator-only fresh Organization/Owner aggregate and authority-bearing constraints, same-person confirmation/replay, and Owner-scoped object inspection. No general roles, Staff, transfer, additional Owner, or legal-authority rule is authorized.
11. ~~**G3 guard control authorization (prior Active 4).**~~ **RESOLVED — same explicit Human decision:** trusted operator-only guard provisioning/control/privileges/audit/runbook and commit-time enforcement. No application user override, administrative API, legal/conflict-resolution judgment, or claim that an environment's guard has been safely opened. Provider/runtime/operator/dependency evidence remains Active 5.
