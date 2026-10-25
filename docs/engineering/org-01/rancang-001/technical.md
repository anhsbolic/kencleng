# ORG-01 — Technical solution r2 (O1/O2/O4)

Assignment `RANCANG-ORG-01-001`; producer Solution Architect `/root/org01_architect`, 2026-10-25. Intended consumers: direct backend/frontend owners, Experience Specialist and independent assessors. **Delegated Engineering proposal**, pending independent assessment and exact-bundle Human acceptance. This is neither adopted OpenAPI/schema nor implementation permission. Only this file is authored by this actor.

Revision r1 → r2: address independent Security SEC-01 (LOW) with transient admission controls for login allocation and formation/replay, safe 429/backoff semantics and added T-08 coverage. No Product quota, eligibility rule or project-wide standard is introduced. Prior r1 SHA256: `a6818aaa7a31205fc1d9b9901ac8b999cfe30486fa6ab5af621829cf8b4f6ef6`; blob `9c873a536b2d7a4ce71185a006d28c7ffa26abfe`.

## Sources and inspected baseline

Product meaning comes from [approved Product Contract](../../../product/outcomes/org-01/product-contract.md), blob `da19d8e7cdb9baac8cf4af1c5e6b61763607013e`, and [approval](../../../product/outcomes/org-01/approval.md), blob `65211210c102afd02143ab1921b2486fb600b241`. Both originals were read; no historical C1 rules are imported. This consumes exact [X1](obligations.md) blob `e1529de3001ae7d261dff30f8b047f5fcd6bb11d`, SHA256 `93359cf29788af494ba6c98ca485640aaabd6e956a2837157ed2dfff3b5a04e1`, checked against the real file. [State](state.md) and [handoff](handoff.md) supply assignment authorization, documented R2 deviation and R5/R6 limitations, not final solution acceptance.

Kencleng inspected HEAD: `3faf3b8a959cdb617958568b158c0805e84107ed`; Harscode HEAD: `df21f47ecb85877a571b15f0793ced2f5c33e003`. Formed definition is `9663a3fa73ac0101e66764d63af8711f291a66ab`, approved by the 2026-10-25 formation record despite embedded candidate labels. Workspace/Harscode root and Engineering AGENTS, Rancang contracts/coordination and Architect agent/knowledge originals were read. Role blobs at that definition: agent `325672c9753f7c9cef6cc3edef4efcaac64c0d2d`, knowledge `e2e6cd726806c224e10b52779cbae63b1c39079b`.

Material project originals read (paths relative to Kencleng; Git blobs at baseline):

| Source | Blob |
| --- | --- |
| `AGENTS.md` | `14fd68fe8e25306270be6e3e259f181079ea9390` |
| `backend/AGENTS.md` / `frontend/AGENTS.md` | `c463cf204ae5dbf32067eaf20c16ceaa76dd6bb9` / `8d5836d6450b2556c91995e57a3c408a7da6b14b` |
| `docs/project/kencleng-backend-tech-stack.md` / frontend equivalent | `f451e63af760ad07991b5c072a3530ed9296799f` / `a1cdd73b816cbeab5509a8729c8f20a3bf012b34` |
| `docs/ui-ux/README.md` / `product-design-principles.md` / `patterns.md` | `426e0d499c73e28b8ffa8af55353d107369f5afa` / `9eb9e5893fc7c1090d28024c51ed507c18cea945` / `da439232e2d17547a986c8a17914b6858b012a79` |
| `backend/cmd/server/main.go` / `internal/platform/db/db.go` / `go.mod` | `745ff37cba7471ff4d15c6e624da0d64341301f8` / `1abb2743ee6b991d5d20fb7b070c04897fe9d0fe` / `eb3283649337f1c61c1a1720b539ee812b09ae33` |
| `frontend/next.config.ts` / `app/page.tsx` | `f6e153f76d600f522cf4565ae89974010ad2bf6a` / `95ea11781e110824f4413999e25f102d1939df18` |
| `api/README.md` / `api/openapi/index.yaml` | `3a7be794640f1b243c47e043b035065ba7651856` / `13fb4f31e04d64228eda04c7995548c93c2f649a` |

Observed code supplies `net/http` health routing, pgx pool/configuration and a Next page/dev rewrite; no actor, authentication, organization domain or product API exists. API source has `paths: {}`. These absences are implementation work, not unresolved Product decisions. Applicable security originals are the root/backend golden rules, protected authentication/privileged authorization boundary and stack rules; no separate mandatory identity implementation was found. Frontend component governance was also read: route-local composition is sufficient; no broad shared component contract is needed here.

## O1 — Selected design and invariants

Use one Go domain-oriented monolith and PostgreSQL transaction. Add only actor/session plumbing needed for attribution, organization formation and private inspection. No queue, object storage, distributed saga, Reviewer workflow or conflict subsystem is justified.

1. An authenticated session binds one stable actor. The server assigns that actor as the first internal Owner of the newly generated organization ID; client-supplied actor/Owner/role/status fields cannot grant authority.
2. Organization, attributable first Owner relationship, creation provenance and successful request identity become durable in one commit. No normal reader can see a committed organization without its Owner. Success is emitted only after commit is positively confirmed.
3. Names are claims, not identifiers. Equal names across distinct requests/actors produce distinct organization IDs and effective Owner relationships. No name uniqueness index, conflict-based denial, merge or legal-verification prerequisite.
4. Inspection reads current committed server state under current authorization. It is separate from a toast, cached creation response or browser memory.
5. Stored name is explicitly organization-provided. Formation proves the internal relationship only; it creates no Reviewer approval, external representation verification, Campaign, publication or donation status/effect.

These choices implement X1-01..09 / Product REQ-01..08, AC-01..09. A single organization row with a non-null Owner foreign key is smaller and gives a stronger bounded invariant than separately committing an organization and membership. A generic multi-role membership/permission system is unnecessary for this outcome; later transfer/Staff work must design its own authorized evolution. PostgreSQL transactions plus request deduplication solve partial writes/lost responses without a recovery service.

## O2-A — Identity, session and authorization trust boundary

Choose **OIDC authorization-code flow with PKCE, completed by the Go backend, then an opaque server-side session**. Use a maintained OIDC/OAuth library in future implementation; do not hand-roll signature validation or use historical auth code as authority. This is a feature-level proposal, not a project-wide identity standard.

Browser may draft before login. Before effective submission, it obtains a backend session through these proposed operations (browser paths add `/api`):

| Backend operation | Contract |
| --- | --- |
| `GET /auth/login` | Begin login with one configured trusted issuer/client and fixed callback URI. Generate random state, nonce and PKCE verifier/challenge; store a short-lived (10 minute), single-use login transaction bound to the initiating browser. Return provider redirect. Allow only fixed local return destinations, never arbitrary redirect URLs. |
| `GET /auth/callback` | Validate browser binding, state, one-use transaction, code exchange/PKCE and ID token signature, allowed algorithm, exact issuer, audience/authorized party, expiry and nonce. Reject missing/invalid subject. Upsert actor by exact `(issuer, subject)` uniqueness, then create fresh session. Consume transaction; clear login cookie. Fail closed; no actor from unvalidated claims. |
| `GET /session` | Authenticated, non-cacheable `{actor:{id, identity_provider}, csrf_token}`; missing/expired session returns 401. `identity_provider` is a server-configured display label, not a browser claim. Before effect, display this provider and the stable Kencleng account ID as the acting account; no email/name is required or treated as identity authority. Actor ID is internal, not a legal identity assertion. |
| `POST /auth/logout` | CSRF-protected invalidation of this server session and cookie; does not promise provider-wide logout. |

Use a cryptographically random session secret with at least 256 bits entropy; persist only its SHA-256 digest, actor FK and absolute expiry (proposed 8 hours), with revocation represented by row deletion. Cookie is HttpOnly, Secure in HTTPS, SameSite=Lax, Path=/, host-only (no Domain); rotate on login, never accept session IDs from URL/body. Tokens/code/verifier remain backend-only and are not logged; discard provider tokens after login unless a separately justified capability needs them. No browser token/local-storage session authority. Local HTTP development needs an explicitly loopback-only non-Secure cookie configuration, never enabled for production. Provider refresh-token storage, local passwords, MFA, account recovery and broad account management are outside this proposal.

An actor is the authenticated account controlling a validated issuer/subject, **not externally proven real-world/legal identity**. Store no email/name/avatar unless independently needed. Do not link accounts by email, display name or claim similarity; subject changes yield a different actor. Re-login to the same issuer/subject recovers the same actor and ownership. Provider/client migration or account linking is a separate security-sensitive decision, not an automatic ORG-01 operation. This mechanism supplies accountable account attribution without pretending to prove legal representation.

Middleware interface: `Authenticate(request) -> SessionPrincipal{actor_id, session_digest}` or unauthenticated/unavailable. The formation transaction must re-read and lock the matching session row (`FOR SHARE`) and check its expiry against database wall clock immediately before inserting effects; logout deletion locks the same row, giving a defined order. Never rely solely on an earlier middleware check. Do not hold the DB transaction across provider calls. Authorization effective time is this locked database decision; expiry/revocation that linearizes after it does not retroactively invalidate a committed formation. Upstream provider revocation is not instant local-session revocation; the bounded session lifetime is explicit. Missing session denies; DB/auth dependency failure does not become anonymous success.

For state-changing cookie-authenticated requests require exact configured application Origin and a session-bound unpredictable CSRF token in `X-CSRF-Token`, compared securely; reject missing/wrong/cross-origin values. `/session` must not expose tokens cross-origin. OIDC callback uses state/browser binding instead of the mutation CSRF header. Keep CORS closed, JSON content type and body bounds, no state-changing GET, no trusting identity headers from the browser/proxy. Use transport TLS outside loopback. Same-origin production routing must forward cookies and trusted origin consistently; the current development-only Next rewrite is not production gateway readiness.

**Bounded authorization:** authenticated actor may form a new context and becomes its initial Owner under approved ORG-01. No Reviewer approval or name exclusivity check. Backend inspection/list authorizes only when `organization.owner_actor_id == principal.actor_id`. Other actors get indistinguishable 404 for missing/inaccessible IDs; unauthenticated gets 401. No public lookup, owner PII exposure, Staff/admin privilege or full permission matrix is introduced. The scope grants no additional Owner actions. Do not add unsupported eligibility restrictions during implementation.

Apply transient admission controls before allocating login transactions (bounded per-source request rates, bounded outstanding records, automatic expiry) and before formation/replay transactions (bounded per-actor request rates and concurrent work). Bound transaction lifetime and total in-flight backend work; failed/replayed requests still consume admission capacity. Thresholds belong to downstream capacity configuration, not organization eligibility or lifetime creation quotas; do not key controls by organization name or Reviewer/conflict state. Reject before mutation/allocation with 429 and a bounded `Retry-After`; do not evict durable idempotency records as overload cleanup. Log only safe aggregate abuse/availability signals. A throttled retry cannot settle any earlier unknown formation.

Provider issuer/client registration, redirect URI, credentials, TLS/public origin/gateway configuration, admission-control thresholds, operational availability and deployed library validation are **downstream configuration/delivery prerequisites**. They are not undecided solution mechanisms or permission to configure them now. A provider unable to supply a validated stable subject is incompatible; hold that deployment edge rather than accepting user-asserted identity.

## O2-B — Persistent state and atomic formation

Proposed relational fields, deliberately kept here rather than in canonical migrations:

| Relation | Minimum fields / constraints |
| --- | --- |
| `actors` | `id UUID PK`, `issuer text NOT NULL`, `subject text NOT NULL`, unique `(issuer,subject)`; immutable mapping, FK references restrict deletion while attributed records exist. |
| `sessions` | `secret_digest PK`, `actor_id FK NOT NULL`, `expires_at timestamptz NOT NULL`, `csrf_secret NOT NULL`; login transactions are short-lived one-use backend records with browser binding/state/nonce/verifier/expiry. Protect secrets at rest/access; no raw session secret stored. |
| `organizations` | `id UUID PK` server-generated; `name text NOT NULL`; `owner_actor_id UUID NOT NULL FK actors`; `created_by_actor_id UUID NOT NULL FK actors`; `created_at timestamptz NOT NULL`; `name_source` constrained to `organization_provided`; `meaning_version` constrained to `org01-internal-owner-v1`. At formation owner equals creator. These fields encode the initial current Owner relationship; no separate uncommitted membership exists. |
| `organization_formations` | `(actor_id FK, request_key UUID) PK`; `normalized_name text NOT NULL`, `meaning_version NOT NULL`; `organization_id UUID NOT NULL UNIQUE FK organizations`. Records committed only with the organization. This is private request deduplication metadata, not an externally accessible operation queue. |

Technical input bounds: body at most 4 KiB, name trim outer whitespace then require 1–200 Unicode scalar values, reject control characters, strict JSON object/known fields. These are storage/transport hygiene, not organization eligibility/legal-name rules. No global case folding, slug uniqueness or name-based identity. Preserve claimed display spelling after outer trim; render as text, never HTML. The exact normalized name and meaning version define equality for idempotency; comparing these small values avoids a custom cryptographic payload scheme.

Transaction at ordinary Read Committed is sufficient: lock/revalidate session; serialize this actor/key using PostgreSQL transaction-scoped advisory lock derived deterministically from actor/key (hash collision merely serializes unrelated requests); read existing formation. Same key/equal normalized payload returns its existing committed organization; same key/different payload returns conflict. If absent, generate UUID, insert organization with Owner=creator=principal, insert formation, commit. Composite unique constraint remains the final safeguard. All insert failures roll back both rows; session/actor FK and non-null ownership prevent orphan success. Concurrent different keys do not serialize by organization name. No pre-create lookup or separate remote call decides ownership. Keep transaction short with bounded database timeouts; do not retry a possibly committed transaction under a new key.

No edit/delete/transfer endpoints are supplied here. Current Owner remains the first Owner in this bounded slice; later authorized lifecycle work must preserve creation attribution while updating current-state inspection. Do not expose stale initial membership as current if later work changes it. Detection of an impossible/inconsistent row is an internal integrity failure, not ordinary success or a client-repairable partial organization.

## O2-C — Feature operations, results and retry

Backend paths below are unprefixed; browser uses relative `/api/...` through the existing development gateway. UUIDs are opaque; timestamps are RFC3339 UTC. Protected responses use `Cache-Control: no-store`; caches must be scoped to actor and cleared on session change.

`POST /organization-formations` requires session, Origin/CSRF and `Idempotency-Key: <random UUID>`. Body:

```json
{"name":"Organisasi K","meaning_acknowledgement":"org01-internal-owner-v1"}
```

Acknowledgement binds the submission to the consequence explanation shown before effect; it is not evidence that a person read or legally consented. Frontend must show the organization name, acting account context and internal-Owner limits before this action, including after sign-in returns. Server rejects missing/unsupported acknowledgement; it cannot prove comprehension. No actor, role, Reviewer status or external authority field is accepted.

Confirmed initial commit returns **201**, same-key/equal-payload replay **200**, both with `Location: /organizations/{id}` and the same logical result:

```json
{
  "outcome":"formed",
  "organization":{
    "id":"<uuid>","name":"Organisasi K",
    "name_source":"organization_provided","created_at":"<timestamp>",
    "current_relationship":{"actor_id":"<self uuid>","role":"Owner","scope":"internal"},
    "meaning_version":"org01-internal-owner-v1"
  }
}
```

`GET /organizations/{id}` returns **200** `{organization:<same current-state shape>}` from an authorized committed read, or 401/404/503. This is authoritative as of that read, not a guarantee against later changes. `GET /me/organizations?cursor=<opaque>&limit=20` returns **200** `{items:[<same shape>],next_cursor:null|<opaque>}`, owner-scoped, ID-ordered keyset pagination; clamp limit to 1–100. A cursor selects a position only, never an actor/permission. This private collection makes separate-session revisiting possible without retaining a response/key forever in the browser. No search/public listing or other owners' details. Read failure never substitutes a fixture or stale success response for current state.

Errors use `{error:{code, effect:"none"|"unknown", fields?:{name?:<safe code>}},request_id:<opaque>}`; safe localized presentation belongs to O3. No SQL, stack/path/credential detail.

| HTTP / code | Mutation meaning and safe next action |
| --- | --- |
| 400 `invalid_request`; 422 `invalid_name` / `meaning_required` | Validation rejected before effects: `none`. Correct fields/explanation and submit; preserve input. |
| 401 `authentication_required`; 403 `request_not_authorized` | This request did not mutate: `none`. Reauthenticate same account or correct request security context. This says nothing about an earlier interrupted request. |
| 409 `request_key_reused` | Key already has a committed different payload: this request `none`. Never overwrite/rebind it; inspect existing contexts and reconcile original attempt. New intent can receive a new key only after unresolved original intent is handled explicitly. |
| 429 `temporarily_throttled` | Admission rejected this request before effects: `none`, with `Retry-After`. Wait/back off and retain original key/payload; any earlier unknown remains unknown. No permanent eligibility/name denial is implied. |
| 503 `temporarily_unavailable` | If definitely rejected/rolled back: `none`. If commit acknowledgement is uncertain: `unknown`. Preserve original key/payload; bounded retry with backoff or inspect. |
| 500 `integrity_failure` | No normal success. If rollback proved use `none`, otherwise `unknown`. Stop automatic retries and surface safe unavailable/uncertain state for operational investigation. |

Transport timeout, aborted fetch, malformed/unexpected result or absent response means **client unknown**, regardless of an attempted cancellation. A committed transaction cannot be undone by a disconnected browser. Do not return ordinary success on SQL execution alone or infer failure from a missing toast. Read failures carry no mutation outcome and cannot settle a prior unknown attempt.

Generate one key per explicit formation intent before first transmission, retain it and exact normalized payload through pending/unknown retries; disable accidental repeated submits. Persist only the minimal pending key/payload locally if offering reload recovery; treat it as untrusted input, never session authority, and do not submit it under a different account. On account change suspend recovery and require the original account; do not disclose/replay original data in another account's UI. Same-key replay under the original authenticated actor resolves a lost response by returning the existing commit or forming once if no committed attempt exists. A replay may wait for an in-flight transaction then return the same result; timeout remains unknown. Private list/detail inspection can also establish actual success. An empty list/404 while another transaction is in flight does **not** prove failure.

Retain the actor/key→organization record as long as the organization exists; no timed key eviction is proposed, avoiding a late retry silently forming twice. Store no separate raw response/token. Key is not a bearer credential: all lookup/replay requires actor authentication and authorization; different actor/same key has a separate namespace and cannot read the first result. If data deletion/retention policy later requires removal, its tombstone/expiry semantics require explicit design before advertising replay safety beyond retention. This is bounded deduplication, not an exactly-once network guarantee. No automatic background retry is required; any retry must retain intent and obey backoff/attempt bounds.

## O4 — Assigned technical verification obligations

Coordinator confirmed this Architect as accountable **producer for every T-* below**. Intended downstream executors are separate from producer and must receive implementation authorization. All implementation tests below are **NOT RUN**. Present evidence is source/document inspection only, not independent assurance or runtime delivery.

| ID / X1 trace | Checkable obligation | Intended downstream verification executor |
| --- | --- | --- |
| T-01 / 01,02 | OIDC fixtures reject bad issuer/audience/algorithm/signature/nonce/state/PKCE, reused callback and browser-binding loss. Same issuer/subject reconnects same actor; same email/different subject does not link. No externally verified/legal identity inference. | Backend owner with security verifier |
| T-02 / 02,04,07 | Direct unauthenticated requests, spoofed actor/Owner fields, revoked/expired session and auth/DB outage cannot form; test logout/session expiry ordering at the transaction decision. Session rotation, cookies and token/log exclusion verified. | Backend owner with security verifier |
| T-03 / 02,05 | CSRF missing/wrong token/Origin and cross-origin reads rejected; another actor cannot inspect/list/replay first actor's data; 404 anti-enumeration and caches/session changes do not leak. | Backend + frontend integration/security verifier |
| T-04 / 04,07 | Real PostgreSQL integration fault injection before/after each insert and around commit: rollback has neither row; confirmed commit has Owner/provenance/request record; lost acknowledgement is unknown then same-key replay resolves once. Never ordinary partial success. | Backend owner, independent downstream test executor |
| T-05 / 07,08 | Concurrent identical actor/key/payload yields one organization; changed payload yields 409; equal names under distinct keys/actors yield separate IDs and Owners; no Reviewer/conflict gate. Test delayed retry with retained key and simultaneous request timeouts. | Backend integration/race verifier |
| T-06 / 03,05,06,09 | Form and read contract tests preserve bounded internal meaning/provenance and reject unsupported fields; post-restart/new-session list→detail reads actual current Owner independently of toast/cache. No review/Campaign/publication/donation writes/events/status claims. | Backend + frontend contract/integration verifier |
| T-07 / 01,03,07 | FE pre-effect explanation survives sign-in, actor switch suspends pending intent, input survives error, double-click is controlled; lost response/cancel/401 after prior unknown never becomes false failure/success or new-key retry. Coordinate with E-* visible/accessibility obligations. | Frontend owner / browser verifier |
| T-08 / standards | Parameterized SQL, bounded strict parsing, escaped claim display, safe errors and sanitized telemetry; evaluate maintained auth dependency/configuration, production cookie/TLS/origin/gateway behavior. Exercise login allocation expiry/caps, per-source login and per-actor formation/replay throttling, bounded concurrent work/transaction lifetime, 429/Retry-After and preservation of prior unknown/key records under overload. Run proportional Go unit/httptest, PG integration/race, API validation/types and frontend checks after authorized implementation. | Respective implementation owners, downstream security verifier |

## Consumer fit, remaining gates and limits

**Backend project-bound design fit: CONDITIONAL, inspected.** O1/O2 map to current Go/net-http/pgx/PostgreSQL monolith and explicit backend authorization requirements. No standards change is necessary: feature-level identity/transaction decisions suffice. Actual code currently lacks them; canonical contract adoption, migrations and protected authentication/Owner implementation require separate authorized downstream actions, then T-* evidence. This document does not authorize those actions.

**Frontend project-bound design fit: CONDITIONAL, inspected.** Proposed paths fit the actual dev rewrite and frontend instructions; route-local form/inspection composition, React Hook Form/local pending intent and server-state queries fit current architecture. Once adopted, generate types from `api/openapi/index.yaml` and bundled output rather than maintaining a parallel model. Current empty API cannot yet generate these shapes. O3 must reconcile exact status/effect/unknown/meaning semantics before bundle completion; MSW can represent these contracts but is not runtime success evidence. No production gateway or rendered UI verification is established.

**Generic reusable consumer contract fit: NOT APPLICABLE** for both direct implementation owners; neither is a downstream Space with a reusable Input Contract. This does not waive the project-bound conditions above.

No unresolved Product decision, material standards change or scope expansion is required by this proposal. Provider/deployment configuration, downstream canonical API adoption, implementation authorization, actual verification, O3 reconciliation, independent security/solution review and exact Human acceptance remain distinct gates. No installation, network call, service action, test, migration, canonical-contract write or effective standards change was performed. Reviewer findings may require producer correction; this actor does not self-approve or rewrite independent findings.
