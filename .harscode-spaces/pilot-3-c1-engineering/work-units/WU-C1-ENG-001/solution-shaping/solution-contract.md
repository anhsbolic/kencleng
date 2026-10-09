# C1 Solution Contract / Planning Readiness Evidence

> Work Unit: `WU-C1-ENG-001`  
> Commitment: C1 — Legitimate Organization representation  
> Date: 2026-10-09  
> Baseline inspected: `pilot/3-c1-engineering`, HEAD `54d73d1fbc63865a1b97373b0faf71e78e8507ed`  
> Activity: Human-requested experimental Solution Shaping; not a Techplan, Build, or new Product/pre-engineering authority  
> Verdict: **PLANNING READY — YES**  
> Build Ready: **NO**; an approved execution contract and protected implementation authorizations remain required

## 1. Authority, evidence, and limits

Stage 5 B1–B7 and Stage 6 XR/ER1–7 remain binding within C1, routed by Stage 7. Root/scoped `AGENTS.md`, current stack architecture, and reusable UX authority constrain the derived solution. This artifact owns bounded engineering solution evidence; it does not amend those inputs or create whole-product truth.

Inputs consumed:

- Root, backend, and frontend `AGENTS.md`.
- `docs/product/product-intent.md` and `pilot-3-stage-7-c1-engineering-handoff.md`.
- `docs/product/pilot-3-stage-5-c1-confirmed-behavior.md` and `pilot-3-stage-6-c1-requirements.md`.
- [Exploration Stage 2](../runs/RUN-C1-ENG-EXPLORATION-001/evidence/stage-2-gap-analysis.md) and [Stage 3](../runs/RUN-C1-ENG-EXPLORATION-001/evidence/stage-3-solutioning.md).
- [Draft Techplan](../techplan/techplan.md), especially D2–D5, RISK-1–6, §§8–9, and §13; [its phase handoff](../runs/RUN-C1-ENG-TECHPLAN-001/evidence/phase-handoff.md).
- Current `docs/project/` architecture, `docs/ui-ux/product-design-principles.md`, and `docs/ui-ux/patterns.md`.
- Harscode root/workflow/orchestration routing and targeted Go authorization/IDOR, privilege separation, token lifecycle, and PostgreSQL transaction guidance.
- `../harscode-workspace/proposals/0040-engineering-readiness-boundaries-and-solution-shaping.md`, **Draft experimental framing**, applied only under this session's explicit Human instruction. It is not adopted as mandatory Harscode policy.

Verified by read-only repository inspection: backend serves only health; the domain/migration directories have no C1 implementation; PostgreSQL connectivity uses `pgxpool`; frontend is the neutral Next.js scaffold; OpenAPI has `paths: {}`; Caddy strips `/api` and routes to the backend on one browser origin. No current principal, session, Organization, Owner, or conflict source exists. This corroborates Exploration F1. Exploration F2 establishes one earlier reader's handoff comprehension, not implementation correctness.

The inspected HEAD is later than the Exploration/Draft target `e9cb1f31f031ed3905181eb86ad3c3e3af4e248f`; live anchors were reopened rather than assuming revision equivalence. Initial `git status --short` was clean. Historical implementations were not reused as authority.

**Evidence class:** repository facts are inspected; the solution is selected engineering design, with the three Human decisions below; its runtime/security/behavior correctness is **not tested**. No source, executable API specification, migration, Techplan, Product authority, or Harscode guidance was changed. No tests, builds, or runtime checks were run. This artifact does not claim C1 delivery or verification.

## 2. Audit of material unresolved concerns

Classifications distinguish an unresolved design choice from a missing Product decision and distinguish design selection from permission to implement it.

| Concern / original evidence | Classification | Resolution / readiness effect |
| --- | --- | --- |
| Minimum valid Organization information; Draft §13.1 / RISK-4 | **DESIGN DECISION**, plus **ENGINEERING DECISION** for technical validity | Stage 7 §7 and Stage 6 §4 expressly preserve field/form choices. Choose a display name and server identifier; valid persisted context includes attributable Owner and provenance. No legal-validity criterion is added. Human selected the bounded flow/field set (H3). Draft's Product-blocker classification is not retained. |
| Exact flow, sign-in timing, disclosure, checkbox, inspection destination; Exploration Area 2 / D5 | **DESIGN DECISION** | H3 selects Form → sign-in if needed → consequence summary → final confirmation → durable detail/list. Copy, layout, and visual composition remain safe design freedom. |
| Principal source / auth mechanism; Exploration Stage 3 question 3, D3 / §13.2 | **ENGINEERING DECISION**; implementation also **PROTECTED HUMAN AUTHORIZATION** | H1 selects Google OIDC and local person/session mapping. §§5–6 settle trust and security direction; gate G1 remains before implementation. No missing Product identity meaning is inferred. |
| Initial Owner grant; Exploration question 1, §13.3 / RISK-2 | **ENGINEERING DECISION**; implementation **PROTECTED HUMAN AUTHORIZATION** | Grant only to the authenticated initiating person, only for the fresh aggregate created by that confirmation. No assignment to another person or existing Organization. Exact rule and reads are settled in §6; G2 remains. |
| Same-database ownership assumption; D2 / §13.4 | **ENGINEERING DECISION** | Select Organization aggregate, attribution references, preparation/result, session checks, and guard in the same PostgreSQL database/transaction boundary. No other service owns either side of C1 success. The earlier conditional assumption is now a selected boundary. |
| API, IDs, data constraints, provenance representation; Draft §8 | **ENGINEERING DECISION** | §§4 and 7 define the material HTTP/data interface. No parallel frontend-authoritative model or review-status taxonomy. |
| Consequence before effect; ER2.1 / Draft R2 | **DESIGN DECISION** and **ENGINEERING DECISION** | Authenticated preparation returns the frozen name and four consequence meanings; separate confirmation makes establishment effective. UI must render the summary before enabling final confirmation. Receipt/version proves a protocol step, not that a human actually read it. |
| Durable Organization/Owner inspection; Exploration Stage 3 / Draft R4 | **DESIGN DECISION** and **ENGINEERING DECISION** | Owner-scoped list and detail read committed aggregate state, with stable attribution and internal-only meaning. Fresh navigation/login works without a transient success payload. |
| Truth-class preservation / review meaning; B4–B5 / RISK-6 | **ENGINEERING DECISION** for representation; established Product invariant | Persist Organization-provided provenance and the internal Owner relationship. Express what the C1 transition did, without inventing a general review state. |
| Known-conflict integration; Exploration question 2, D4 / §13.5 / RISK-5 | **ENGINEERING DECISION**; bypass/enable controls also **PROTECTED HUMAN AUTHORIZATION** | H2 selects a durable conservative establishment guard with an operator input path (§8). It is not an always-clear stub. No matching source is presumed. |
| Conflict detection, relevance matching, legal adjudication, resolution, duplicate guarantees | **DEFER / NOT NEEDED** | Remain outside C1. The conservative guard avoids requiring a matching policy for this commitment. Any later selective relevance source requires separate owning decisions/integration. |
| Partial persistence, concurrent confirmation, response loss, unknown commit outcome | **ENGINEERING DECISION** | §9 settles atomicity, lock ordering, confirmation idempotency, and indeterminate outcome. Build must not equate timeout with rollback. |
| Backend/frontend ownership and production topology | **ENGINEERING DECISION** | Backend owns identity, authorization, persistence, success, guard, and reads; frontend owns form and consequence presentation. Same-origin HTTP boundary; no browser provider-token authority. |
| Verification strategy still conditional on unselected solution; Draft §12 | **ENGINEERING DECISION** | §11 fixes the material oracles, negative security cases, real PostgreSQL concurrency/failure evidence, and real-provider evidence boundary. Specific test organization remains planning freedom. |
| Save-draft/resume, general retry/recovery product, multi-owner/Staff matrix, invitations/transfers, review taxonomy, campaigns, payments | **DEFER / NOT NEEDED** | Excluded. Technical confirmation replay is selected to resolve uncertain transport outcomes; it does not introduce a saved-draft product journey. |
| Independent planning review, execution approval, protected write authorization | **PROTECTED HUMAN AUTHORIZATION** / downstream planning gate | These gate Build readiness; their bounded surfaces are already specified. They are not unsolved architecture/security design. |
| A true missing C1 product/domain meaning | **PRODUCT GAP: NONE IDENTIFIED** | B1–B7 and XR/ER supply the needed meaning; deliberate downstream freedom explains the remaining solution choices. Return upstream only if subsequent evidence requires changing those meanings. |

The Draft and Exploration remain historical evidence. This contract settles their material solution questions for the next synthesis; it neither edits the Draft nor promotes its requests into Product authority.

## 3. Human decisions recorded in this session

These are bounded solution/design decisions, not implementation authorizations.

| ID | Problem and recommendation presented | Exact Human response | Consequence |
| --- | --- | --- | --- |
| H1 | No existing trusted principal; recommend Google OIDC with backend validation, local issuer/subject mapping, and server-side session rather than local email/password auth. | **“Pilih Google OIDC untuk C1”** | Real users of this C1 path need a Google account. OAuth configuration is a verification/runtime prerequisite. Account authentication does not establish legal identity or external Organization authority. |
| H2 | No existing conflict/relevance source; recommend durable operator-controlled guard that blocks all new establishments when unresolved conflict knowledge cannot safely be scoped. | **“Setujui guard konservatif C1”** | Prefer conservative availability over invented matching. No assertion that every blocked Organization is disputed. Existing durable inspection remains available. |
| H3 | Deliberately open Organization fields and material interaction structure; recommend display name only, sign-in before server preparation, four-consequence summary, final confirmation, then reopenable detail/list. | **“Setujui struktur C1 ini”** | No legal-evidence prerequisite, mandatory checkbox, or saved-draft/resume feature. Ordinary copy/layout/visual decisions remain downstream. |

No Human-owned solution decision remains waiting for an answer. Protected implementation gates are separately retained in §12 because this session authorizes shaping and expressly excludes Build.

## 4. Architecture, ownership, and persistence contract

**Selected architecture:** current domain-oriented Go monolith with `net/http`, `pgx`, PostgreSQL, and Next.js App Router. C1 is a backend Organization establishment capability, not a new service. Identity/session is a narrow backend platform concern providing a trusted local `person_id`; Organization domain logic consumes that principal without knowing provider tokens. Composition stays in backend bootstrap/transport. Frontend accesses same-origin `/api/*`; no new Next.js authentication authority or BFF token store is introduced.

All rows that decide C1 success are in the existing PostgreSQL database. No network call, object upload, message publish, or external identity exchange participates in the establishment transaction. No outbox/event bus/object storage is needed for C1.

Material logical data model (physical SQL names may be adjusted without changing this contract):

| Record | Material content / constraints |
| --- | --- |
| `persons` | Server-issued UUID; immutable authenticated identity mapping `(issuer, subject)` unique and case-sensitive. Google subject is opaque. Email/name/domain never identify or authorize the person. No automatic account linking. Attribution rows must not be deleted while referenced. |
| `sessions` | Digest of a high-entropy opaque session token, `person_id` FK, expiry, revocation, and a separate CSRF token. A provider token is never the Kencleng session. |
| `organization_establishment_preparations` | Server UUID, authenticated initiating `person_id`, frozen validated `display_name`, consequence version, creation/expiry, and nullable committed result reference. Preparation is an incomplete command receipt; it is **not** an Organization or Owner. It grants no rights. |
| `organizations` | Server UUID, `display_name`, non-null `initial_owner_person_id` FK, `established_by_person_id` FK, establishment timestamp, preparation UUID unique, consequence version, and `information_source = organization_provided`. C1 requires establishing person = initial Owner. These attribution/establishment facts remain durable. |
| `c1_establishment_guard` | One durable guard row controlling whether new confirmations may take effect; operator change evidence and timestamps (§8). Missing/unreadable row fails closed. |

**Owner relationship representation:** for C1, the non-null `initial_owner_person_id` on the Organization aggregate is the initial Owner relationship. Do not build a generic role/membership/permission engine just to represent this one relationship. Its FK and equality constraint with the establishing person make a committed row incapable of representing an unattributable or separately incomplete initial Owner. No current C1 API may mutate that pointer, transfer ownership, or assign additional roles. Future role lifecycle work must separate immutable initial-establishment provenance from current role changes under its own authority; it must not reinterpret this C1 constraint silently.

**Technical Organization validity:** name is a user-provided display label, not legal registration or a uniqueness key. Trim leading/trailing whitespace; accept 1–200 Unicode scalar values, reject control characters, render as escaped text. Server validates. Use UUIDs for references; allow repeated names and do not deduplicate. The name, durable identity reference, Owner relationship, and provenance constitute the bounded C1 record shape. A valid C1 context is one satisfying these technical constraints and the confirmed success invariant, not a newly invented claim about legal validity.

No attachment, legal identifier, address, review-status enum, `verified` badge field, real-world Organization matching key, or Staff table is required. Scope is additive migrations from a baseline with no C1 schema/data; no legacy conversion or destructive data migration is selected.

## 5. Trusted person attribution / authentication integration

**Selected direction:** Google OIDC Authorization Code flow in the Go backend, with PKCE `S256`, browser-bound one-time `state` and `nonce`, fixed callback, and server-side ID-token verification. Use maintained protocol libraries (`golang.org/x/oauth2` and `github.com/coreos/go-oidc/v3/oidc` are the selected integration direction), not handwritten JWT/signature logic. Pin compatible dependency versions during Techplan; an incompatible dependency is a reconciliation trigger, not permission to weaken verification.

Trust only the configured Google issuer `https://accounts.google.com` and this application's client ID. Validate signature against trusted issuer keys, allowed algorithm `RS256`, issuer, audience/authorized party as applicable, expiry, and the exact nonce. Reject caller-selected issuers, unsigned tokens, verification bypasses, and claims from unverified tokens. Key retrieval must fail safely when no valid verification key is available. State/nonce/PKCE material is tied to the initiating browser, expires within ten minutes, and is consumed once. A process-local bounded pending-login store is sufficient for this single-process C1 direction; restart loses pending login and requires a new login, never grants authority. Multi-instance callback coordination would require solution reconciliation.

Request only `openid`; no email/profile, Google API access, offline access, or refresh token is needed to create the durable `(issuer, subject)` mapping. On verified login, upsert the local person by that unique pair; rotate any prior browser session and issue a new opaque local session. Identity exchange completes before any Organization transaction. Do not retain Google access/ID tokens after verification. Do not log callback queries, tokens, secrets, or subject payloads.

**Session direction:** 32 random bytes of bearer entropy; store only its SHA-256 digest for lookup. Local session has an absolute eight-hour lifetime, with server revocation on logout; re-login may create a fresh session. No local signed JWT, password credentials, MFA implementation, recovery, provider-token refresh, or account linking is selected. Loss of a Google account is not solved by silently reassigning Owner to another identity.

Production session/pre-login cookies: host-only, `HttpOnly`, `Secure`, `SameSite=Lax`, `Path=/`; production session uses `__Host-kencleng_session`. Mutation requires an exact configured Origin, JSON/custom CSRF header, and a session-bound CSRF token. No credentialed cross-origin access. OIDC callback uses its browser-bound state/nonce protections and immediately redirects to a fixed local route; it is not covered by the ordinary mutation Origin rule. Callback/error pages and responses use no-store and avoid referrer/code leakage. No arbitrary return URL.

**Current local topology:** explicit development-only `http://localhost:8080` may use a distinct host-only development cookie without `Secure`/`__Host-`; keep `HttpOnly`, `SameSite`, Origin, and CSRF protection. The callback is `/api/auth/google/callback` on that origin. Non-development configuration requires HTTPS and must reject this downgrade. Direct UI origin `:3000` is not the authenticated C1 browser entrypoint. This exception is part of G1's protected scope, not a production default.

Operational configuration: fixed browser origin, Google Web OAuth client ID/secret, exact registered callback, and consent/test-user setup. Secret values stay outside committed artifacts. Their values need not be discovered by a Planner to decide the architecture; they must exist for genuine provider verification. Missing configuration yields unavailable auth, never a fake principal.

Provider protocol evidence: Google documents server-side OIDC and stable subject rather than email as the identity key ([Google OIDC](https://developers.google.com/identity/openid-connect/openid-connect)); its current [discovery metadata](https://accounts.google.com/.well-known/openid-configuration) advertises this issuer, RS256, and S256. The [web-server guide](https://developers.google.com/identity/protocols/oauth2/web-server) supports localhost test redirect registration. [RFC 9700](https://www.rfc-editor.org/rfc/rfc9700.html) recommends PKCE for confidential clients. These confirm the integration direction, not successful Kencleng configuration or login. Library surfaces were inspected in the [OIDC package](https://pkg.go.dev/github.com/coreos/go-oidc/v3/oidc) and [OAuth2 package](https://pkg.go.dev/golang.org/x/oauth2) documentation.

## 6. Privileged Owner authorization boundary

The backend rule for the fresh C1 grant is exact:

1. Resolve a currently valid, non-revoked local session to a durable person; no person/Owner/role field in a request can replace it.
2. Preparation belongs to that same person, is valid/unexpired/uncommitted, and carries the accepted consequence version.
3. Establishment guard permits this new commit; validation succeeds.
4. Create a **new** Organization with its initial Owner equal to that person, atomically with the preparation's committed result.
5. Expose success only after a confirmed database commit. Confirmation replay may return the same previously committed result; it never grants again.

Authentication alone grants no authority over an existing Organization. There is no endpoint to assign oneself Owner of an existing ID, choose a different initial Owner, or bypass the guard. Backend rejects unexpected owner/person/role/conflict fields rather than silently treating them as authority.

For list/detail/result access, scope queries to `initial_owner_person_id = authenticated person_id`. No public Organization/person directory is introduced. Other-person and nonexistent resource IDs receive the same `404` shape; absent/expired session receives `401`. Browser visibility does not authorize reads or writes. Owner in C1 conveys only this internal relationship and access to its C1 inspection surface; broader Owner/Staff permissions remain out of scope, not inferred from a role label.

The implementation of these rules, their Owner-bearing persistence surface, and guard bypass controls is protected by G2/G3. Product approval of B1–B7 and H1–H3 do not authorize those writes.

## 7. Material HTTP / interface contract

Paths below are backend/OpenAPI paths; the browser uses the `/api` prefix stripped by current Caddy. JSON keys below are the material contract; generate frontend types from the active OpenAPI source during later authorized work.

| Operation | Request / trusted context | Result / effect |
| --- | --- | --- |
| `GET /auth/google/start` | Fixed local return destination; create browser-bound login attempt | Redirect to Google; no Organization/Owner effect. |
| `GET /auth/google/callback` | Google code/state plus initiating-browser binding | Verify and map identity; issue local session, then redirect to C1 form. Failure grants no session/Owner. |
| `GET /me` | Local session | `{person_id, csrf_token}`; no provider token or public identity claims. |
| `POST /auth/logout` | Session, Origin, CSRF | Revoke current session and expire cookie; `204`. |
| `POST /organization-establishments` | Authenticated session, Origin, CSRF; `{display_name}` | `201` preparation receipt: `{id, display_name, expires_at, consequence_version, consequences}`. **Only preparation**; no established Organization or Owner. |
| `GET /organization-establishments/{id}` | Same-person authenticated access | `200` tagged union: `{state: prepared, ...receipt}`, `{state: expired, id}`, or `{state: established, organization}`. Supports resolving a lost confirmation response. Other-person/unknown `404`. |
| `POST /organization-establishments/{id}/confirm` | Same-person session, Origin, CSRF; `{consequence_version}` | `200 {state: established, organization}` on confirmed commit or replay of that exact committed result. No new name, person, role, or target Organization input. UUID preparation is the idempotency identity. |
| `GET /me/organizations` | Session | `{items: OrganizationView[], next_cursor: string|null}` scoped to this person. Stable timestamp/UUID keyset ordering; bounded page size. No public browsing or search/matching. |
| `GET /organizations/{id}` | Session and object-level Owner check | `200 OrganizationView`; otherwise `404`. Authoritative durable inspection. |

`consequences` is a versioned semantic object conveying all four B2 meanings: `creates_organization_context: true`, `assigns_you_initial_owner: true`, `owner_authority_scope: kencleng_internal`, and `organization_information_source: organization_provided`. Frontend expresses these understandably in Bahasa Indonesia. `consequence_version = c1-v1` binds the displayed meaning to confirmation; an unknown/stale version requires fresh preparation/display. This is not a forced-checkbox or proof-of-reading claim.

`OrganizationView` contains `id`, `display_name`, `established_at`, `established_by_person_id`, `initial_owner_person_id`, `your_role: owner`, `owner_authority_scope: kencleng_internal`, `information_source: organization_provided`, and `establishment_effects: {organization_review_performed: false, external_authority_verified: false, campaign_curation_performed: false}`. The last object describes **what this establishment transition did**, not a fabricated current Organization review-status taxonomy or a claim about all future review outcomes. A future review capability must introduce its own owning contract. No provider subject/email or legal-authority badge is returned.

Common failures use `{error: {code, message, request_id}}`, with safe localized text and stable codes. Exact messages can evolve without changing semantics. Material mappings:

| HTTP / code | Meaning / required UI behavior |
| --- | --- |
| `400 invalid_request` | Malformed JSON, unknown authority-bearing fields, invalid path/body shape; no mutation. |
| `422 invalid_organization_input` | Display-name validation; field correction, no success. |
| `401 authentication_required` | No valid session; sign-in, never success. |
| `403 request_not_allowed` | Origin/CSRF failure; no processing/grant. |
| `404 not_found` | Unknown or outside-person resource; no existence disclosure. |
| `409 consequence_changed` | Version mismatch; fetch fresh preparation and show consequences again. |
| `410 preparation_expired` | Uncommitted preparation expired; start new preparation. A committed receipt stays retrievable/replayable. |
| `503 establishment_unavailable` | Guard is closed or cannot be safely read; no new grant; do not label every candidate as conflicted. |
| `503 processing_unavailable` | Database/lock/service failure confirmed before commit; safe retry of same command. |
| `503 outcome_unknown` or transport loss | Commit/result not confirmed to caller; show uncertainty, inspect/replay **same preparation**, never claim definite rollback or start a replacement blindly. |
| `429 too_many_requests` | Bounded auth/preparation abuse control; safe explanation, no hidden success. Exact thresholds are operational tuning. |

Read/mutation responses are private/no-store, including errors. Auth callback failures use a fixed safe local error destination; no raw provider response, SQL, stack, or filesystem internals reach the user. Preparation IDs are opaque references with person-scoped access, not credentials.

## 8. Known-conflict integration contract

There is no baseline conflict source. C1 therefore selects a **real durable input boundary**, not `hasConflict() = false`, a client boolean, or a fixture standing in for production knowledge.

One PostgreSQL guard row controls new establishment. A trusted platform operator records a received known unresolved representation-conflict hold via a controlled, parameterized DB operation/runbook. Since C1 has no approved relevance/matching mechanism, the hold conservatively closes all new establishment. Guard opening/closing records operator reference, timestamp, reason/evidence reference, and guard revision in an append-only change record committed with the guard update. Avoid raw claimant/PII payloads. Separate restricted operator credentials perform the controlled operation; browser/API roles cannot change guard permission or its change evidence. Do not introduce an administrative endpoint or new operator-role matrix.

Database privilege detail is material: PostgreSQL locking reads require `UPDATE` privilege on at least one column ([SELECT documentation](https://www.postgresql.org/docs/16/sql-select.html)). Use a fixed singleton primary key constrained to `id = 1`; the backend runtime credential receives `SELECT` and column-level `UPDATE(id)` solely to permit `SELECT ... FOR UPDATE`, with no guard-state-column update, insert, delete, table ownership, schema-creation, or role-management privilege. Updating that constrained key cannot enable establishment. Runtime credentials are separate from migration/owner/operator credentials and cannot change guard constraints or audit evidence. Verification must exercise these actual DB privileges, not just API denial.

The guard initializes **closed** until an authorized operator establishes the input boundary and explicitly permits establishment for this environment. Opening is an operational statement about available platform knowledge; it is not independent verification that no real-world conflict exists. On receipt of unresolved conflict knowledge, the operator must close the guard before permitting additional grants. Reopening requires an explicitly owned external decision/evidence that all outstanding holds no longer apply; clearing one of several holds is insufficient. C1 neither determines nor performs resolution. The later Techplan must include this runbook and permissions boundary rather than treating operator input as automatic external evidence.

All accepted conflict input in this C1 environment must enter this guard. There is no second asynchronous conflict store to ignore. If a later integration adds such a source, it must atomically update the guard or reconcile the solution before claims of compliance; incomplete propagation must not default to allowed. Human awareness outside the recorded input is not magically detected by this software. Operational receipt/recording is therefore a material verification and operating prerequisite, openly stated rather than presented as conflict discovery.

Confirmation locks the guard and observes its current value within the commit transaction. Guard writes use the same row lock. If a close commits first, no new grant follows it; if establishment holds the lock and commits first, that is an earlier grant and the later hold controls subsequent grants. No claim of retroactive revoke or conflict-free real-world representation is made. Previously committed confirmation replay and Owner-scoped inspection remain available while the guard is closed; they are reads of historical successful establishment, not a fresh uncontested grant.

**Availability consequence accepted by H2:** unrelated Organizations may also be unable to establish during a hold. UI says establishment is temporarily unavailable; no selective relevance claim, dispute outcome, or Organization review state is inferred. Future selective conflict handling is separate work.

## 9. Transaction, concurrency, and material failure semantics

Use one short PostgreSQL transaction for confirmation, with explicit row locking under `READ COMMITTED`. Consistent acquisition order: authenticated session row → guard row → preparation row. Session validation is rechecked under lock, including revocation/expiry and person mapping; logout/revocation uses that session row so it cannot race an unchecked grant. Same-person lookup prevents another actor locking or learning the preparation. Guard writers lock only the guard; auth session rotation locks affected sessions in deterministic order and never enters the Organization path.

Within that transaction:

1. Validate session and same-person preparation binding.
2. If already committed, return its stored result after validating person scope; do not create another Organization, reapply the grant, or reject a historical result solely because a new guard hold exists.
3. For an uncommitted preparation, enforce guard permission, expiry, consequence version, and frozen input validity.
4. Insert the complete Organization aggregate (including valid initial Owner/provenance) and link its result to the preparation; the unique preparation reference is a second defense against duplicates.
5. Commit, then return success. Any definite pre-commit failure rolls back both aggregate creation and result linkage.

Logical validation/replay may inspect the preparation while holding the established lock order; an implementation must not reverse the order for convenience. No Owner-success read path bypasses the committed aggregate. Database FKs/non-null/check/unique constraints enforce attribution and one result per preparation in addition to service checks.

Concurrent confirmation of one preparation returns one Organization/result; a replay cannot be converted into a new establishment by sending new fields. Different preparations may establish different Organization contexts even with the same name; this does not promise real-world uniqueness. The guard lock intentionally serializes this bounded pilot's commits against guard updates. This is an availability/throughput tradeoff, not money-movement locking logic.

Do not use blind retries after a possibly successful commit. On a known deadlock/serialization rollback, a bounded retry of the same preparation is safe; otherwise return the safe failure. On commit acknowledgment loss, response encode failure, or disconnect after commit, the result is **indeterminate to the caller** until authoritative read/replay resolves it. Keep the preparation identifier in the browser until resolved; do not show success from local state or definite failure from timeout. PostgreSQL result linkage survives restart and subsequent login by the same person. Retain committed receipts with their attribution; cleanup may remove expired uncommitted preparation rows, not evidence needed to resolve a committed result.

This bounded technical replay/read path is required by the selected transport/atomicity design. It does not add automatic retries, persisted user drafts, multi-step domain completion, or a general recovery journey. No half-formed Organization/Owner becomes a visible successful context.

The locking/isolation direction follows PostgreSQL's documented [row locks](https://www.postgresql.org/docs/16/explicit-locking.html) and [Read Committed behavior](https://www.postgresql.org/docs/16/transaction-iso.html); correctness for this implementation still requires the concurrency evidence in §11.

## 10. Frontend responsibility and durable inspection

Human-selected structural flow:

```text
form: Organization display name; no legal-proof requirement
→ sign-in if needed, before authenticated server preparation
→ preparation receipt + visible summary of all four B2 consequences
→ explicit final confirmation
→ committed result / durable Organization detail
↳ later navigation: my Organizations list → same durable detail
```

Frontend keeps form lifecycle local, using current form tooling where useful. It may preserve the name through an auth redirect using tab-local storage; no bearer token/provider identity/role authority lives there. Backend preparation freezes the submitted name. Editing that name requires a new preparation and redisplayed consequence; final confirmation does not accept changed fields.

Final action is unavailable while preparing, while disclosure has not been rendered, or while confirmation is unresolved. No automatic confirmation on login/callback/page load. Loading, invalid input, sign-in required, expired preparation, availability hold, definite failure, indeterminate result, empty list, and durable success have distinguishable presentations. No mandatory checkbox or additional full-screen step is prescribed. Form/confirmation can be composed within one route; route grouping and visual details remain safe freedom.

Durable detail comes from backend reads, including after refresh/restart/re-login, and contains Organization identity, this person's Owner relationship, internal authority meaning, and the establishment's non-review/non-verification effect. List/detail provides a navigable inspection path without needing a retained success response or remembered UUID. Existing aggregate inspection remains readable during a new-establishment hold. No success toast, optimistic cache insert, or browser-stored role may become evidence of authority.

API/server state remains authoritative; generated types come from the later active OpenAPI contract. Server rendering or TanStack Query may be chosen for reads provided authentication/privacy/no-store behavior is preserved. Frontend MSW fixtures are contract-faithful development evidence only; no production mock fallback or independent verification claim.

## 11. Verification implications fixed before Techplan

These are requirements for later verification evidence, **not checks executed in this session**. Techplan chooses test organization/commands and the implementation task order; it does not discover the following correctness oracles.

| Meaning / risk | Required observable evidence |
| --- | --- |
| B1 / ER1: trusted attribution | Real provider login can map to a stable local person across sessions; local `(issuer, subject)` uniqueness survives concurrent first login. Test forged issuer/audience/signature/algorithm, expired token, wrong nonce/state/browser, code replay, and missing session. No email/domain/body-person spoof grants Owner. Protocol-library fakes are labeled simulations; they do not establish real Google integration. |
| Auth/session boundary | Cookie flags and explicit localhost-only exception, session fixation/rotation, absolute expiry, revocation/logout, Origin/CSRF denial, fixed redirects, callback secrecy/no-store, no credentialed CORS, and provider failure without fake principal. Verify expired/revoked session under the transaction lock cannot produce a new grant. |
| B2 / XR2, ER2 | Rendered walkthrough shows each of the four consequences before final action. Preparation/callback never creates Organization/Owner; direct invalid-version confirmation cannot commit. Version receipt is not treated as proof of reading. Check keyboard, responsive visibility, and error states in the material UI. |
| B3 / ER3 | Real PostgreSQL integration verifies one valid aggregate/Owner/result after commit, FK/check failures reject invalid attribution, and injected failure between aggregate insertion and result linkage leaves neither successful aggregate nor committed result. Unit-only mocks do not prove transaction behavior. |
| B4 / ER4 | After a fresh browser context/login and authoritative read, my-list discovery and detail show the same context and Owner/internal meaning without the original response. Other-person list/detail/preparation/result access is denied with the selected non-disclosing semantics. |
| B5 / ER5 | Persisted information stays Organization-provided; no establishment code/API/UI field implies review, external/legal verification, or curation. Rendered review checks the difference between establishment effect facts and invented current review status. |
| B6 / ER6 | Definite rollback never appears successful; prepared/expired state grants no rights. Commit acknowledgment loss and post-commit response loss resolve to the same durable result through read/replay. Concurrent confirmation yields one result; changed-input replay fails. No automatic new preparation after uncertain commit. |
| B7 / ER7 | Real guard state/runbook, not only a mocked predicate: close guard and confirm no new Organization/Owner; test absent/unreadable guard and actor bypass attempts. Race guard close against confirmation and establish the lock/commit ordering; test a hold arriving after preparation but before confirmation. Inspection and historical replay still work. No matching/adjudication suite is claimed. |
| Protected boundaries | Review exact auth/session, Owner-grant/read-scoping, and guard control surfaces against G1–G3. Direct backend requests must obey them even with bypassed UI. No claimed security approval from frontend gates or unit compilation. |

Use the configured Podman/PostgreSQL route for local integration; Docker CLI absence is not evidence that database verification is unavailable. Dependency/library compatibility, actual callback registration, provider credentials, and a real Google login remain downstream evidence prerequisites. If not available, report that verification as not run; do not label simulated login or manual fixtures as real attribution verification.

## 12. Protected Human gates and safe remaining freedom

### Protected implementation gates still applicable

Explicit authorization to **implement** was not requested or granted in this shaping-only session. The exact surfaces are decision-ready for the later protected gate; they are not undefined rules for a Planner to discover:

- **G1 — authentication/token/session core:** Google code/PKCE/state/nonce verification, local person mapping, opaque token issuance/digest storage, session expiry/rotation/revocation, cookies including the explicit development exception, CSRF, and secret configuration handling in §5. Existing architecture approval/H1 is not this authorization.
- **G2 — privileged initial Owner rule and scope:** server grant to the authenticated initiating person only, fresh aggregate formation and authority-bearing constraints, same-person confirmation/replay, and object-level inspection authorization in §§4, 6–7, 9. No general Owner permissions, transfer, Staff, or legal-authority rules are authorized.
- **G3 — privileged guard enable/bypass control:** trusted operator-only guard writes/runbook and commit-time enforcement in §§8–9. No application user can set the guard, supply an override, or clear unresolved knowledge. No conflict-resolution judgment is introduced.

Before any Build writes these protected surfaces, obtain and durably record explicit Human implementation authorization for the same bounded surface, together with the applicable approved Techplan/review gates. If authorization narrows the selected solution materially, reconcile it before dependent Build. No money movement, custom signing/encryption/HMAC core, or local signing-key system is selected; adding one is new protected scope.

These implementation permissions can remain pending while **Planning Ready** is true: the exact proposed security/authority design is settled, and a Planner can express the gates without inventing the rules. Build readiness is false. Session scope expressly stops before Techplan synthesis and implementation; a subsequent gate must facilitate the Human decision directly rather than send a generic blocker list around an Orchestrator loop.

### Remaining safe Techplan/Build freedom

- Exact Go package/file/function naming, narrow service/repository interfaces, SQL statement organization, migration names, and generated client organization, preserving the chosen monolith/transaction boundaries.
- Compatible pinned library versions, bounded timeout/retry budgets, preparation TTL within the selected short-lived intent model, list page sizes/cursor encoding, cleanup scheduling, rate-limit thresholds, and performance tuning that cannot weaken trust, atomicity, or guard enforcement. Pending login maximum ten minutes and session absolute eight-hour lifetime are settled security limits.
- UI route naming/grouping, inline versus a separate summary view, copy, spacing, typography, accessible focus/error presentation, and responsive composition under current design authority and H3. All four consequences and durable navigation remain mandatory.
- Server Components versus query-based reads, form composition, and contract-faithful MSW fixtures, with the selected API/private-state responsibility unchanged.
- Test file placement/tools and execution sequencing, decomposition of backend/frontend work, review/approval mechanics, and actual environment-specific OAuth secret/origin values. Verification oracles and explicit provider/guard operating prerequisites are already fixed.

Safe freedom does not include switching identity providers, weakening session checks, adding body-supplied Owner identity, making names unique/matching keys, introducing legal prerequisites/review states, changing conflict behavior to an always-clear default, moving Owner to another service/database, exposing partial success, or expanding role lifecycle. A material change requires bounded solution reconciliation; a changed Product meaning requires its upstream owner.

## 13. Material alternatives rejected

| Alternative | Why not selected for C1 |
| --- | --- |
| Send minimum Organization fields back as a missing Product validity decision | Deliberate field/form freedom is already granted by Stage 6/7. A display label and technical record constraints resolve the implementation question without inventing legal validity. |
| Local email/password auth | H1 selects delegated Google identity. Local credentials add password/recovery/security ownership not needed for this bounded C1 path. |
| Caller-supplied person ID, frontend-only principal, development header, or prepared fixture as production attribution | Cannot establish the trusted authenticated person at a privileged grant; violates attribution/security meaning. |
| Generic role/membership permission engine or separate Owner service | No current C1 need for Staff/multi-owner/transfer; expands authority and cross-service consistency unnecessarily. Inline attributable initial Owner relation satisfies the bounded invariant. |
| Staged Organization creation followed by later Owner grant/saga | Optional recovery becomes required complexity and risks exposed half-success. Both are owned by the same PostgreSQL aggregate/transaction here. |
| Single create request without a preparation/confirmation boundary or replay identity | Makes consequence timing and unknown commit response harder to enforce/resolve across this new capability. A frozen, same-person receipt provides both boundaries without a saved-draft product. |
| Ignore conflict until some future source exists; test only an always-clear adapter | Leaves the material guard direction and ER7 verification unresolved. H2 chooses an actual durable input and enforcement path. |
| Detect duplicates by Organization name, infer Google email-domain authority, or adjudicate legal representation | Would invent matching/authority policy outside C1. H2 accepts coarse availability rather than such policy. |
| Public details or success-state authority stored only in the browser | Durable Owner inspection and authorization must come from backend state; private scoped reads are sufficient for C1. |
| `verified`, `pending_review`, or equivalent review taxonomy on C1 success | Confuses Organization-provided provenance and internal role with review/verification meaning; explicitly excluded by Stage 5/6. |

## 14. Exact Planning Ready verdict

| Material readiness concern | State at this boundary |
| --- | --- |
| Product/domain meaning and upstream exclusions | Settled by existing Stage 5/6; **no true Product gap identified**. |
| Material experience structure and field choice | Settled by H3; ordinary visual freedom remains. |
| Architecture/ownership | Selected monolith, backend authority, single PostgreSQL ownership. |
| Interface/data | Material operations, shapes, constraints, provenance, attribution, and durable read path specified. |
| Identity/security direction | H1 + exact OIDC/local-session trust contract; no caller assertion or provider role inference. |
| Privileged Owner/guard rules | Exact proposed grant/scope/hold rules settled; explicit implementation permission gates retained. |
| Consistency/concurrency/failure | Aggregate transaction, lock order, repeat-confirmation identity, guard race, and indeterminate-result semantics settled. |
| Known-conflict integration | H2 + real operator-controlled durable input; no matching or fabricated no-conflict verification. |
| Verification implications | Material positive/negative/concurrent/real-provider/rendered oracles fixed; execution evidence not yet run. |
| True unresolved Planning Ready blocker | **NONE**. |

**PLANNING READY — YES.** The next Techplan can synthesize implementation and verification tasks from this contract without becoming the primary discovery mechanism for material product/domain, authority/security, architecture/ownership, interface/data, consistency, or verification decisions.

**This does not mean Build Ready, Techplan approved, implemented, independently verified, or C1 complete.** Google configuration and operator guard initialization/operation are explicit future evidence/runtime prerequisites; G1–G3 and execution approval remain Build gates. None is concealed as an unselected material solution.

The historical WU/Space dispatch projection predates the stored Draft handoff; it should not be read as proof that no planning activity occurred. This shaping evidence does not reuse or dispatch that Run, manufacture Run/Participant identity, or revise its Techplan. A later synthesis must consume this contract and use the current applicable fresh-run/approval routing. Work in this session stops at this completed Solution Contract and verdict.
