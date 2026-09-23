# EXP-TOP-001 — Gap Analysis

> Phase/Stage        : Exploration / Stage 2 — Gap Analysis
> Work Unit          : `WU-S1-005`
> Run                : `EXP-TOP-001`
> Role               : Explorer
> Specialization     : root Caddy proxy, Docker Compose, MinIO policy, same-origin `/api` routing, controlled Campaign media topology
> Participant        : Codex CLI agent
> Session            : Fresh session
> Created            : 2026-09-23
> Model              : `gpt-5.6-terra`
> Reasoning          : medium
> Target revision    : `70ef5c6af1a9befce8492512ee11fec4b22f779f`
> Workflow revision  : not exposed by the invocation

## Scope and boundary observed

This Run owns only root proxy, Docker Compose / MinIO policy, and runtime
integration boundaries for Slice 1. Campaign eligibility, media membership,
HTTP handler semantics, persistence, and frontend product behavior remain
owned by `WU-S1-003` and `WU-S1-004` respectively. No code/configuration was
modified during this Stage.

The current-effective authority requires the public `content_url` to remain a
same-origin `/api/.../content` reference, with controlled origin delivery,
private Campaign media, parent/member checks on every request, identical
public `404`, eligible-dependency `503`, and `Cache-Control: private,
no-store` on `200`/`404`/`503`:

- `docs/spec/4-campaign/features/03-campaign-media.md` — Summary, Behavior,
  Concurrency & correctness notes, Test checklist.
- `docs/spec/4-campaign/invariants.md` — `INV-campaign-14`.
- `docs/spec/4-campaign/threat-model.md` — “Slice-1 Campaign media delivery”.
- `api/openapi/campaign.yaml` — `getPublicCampaignMediaContent` and
  `PublicCampaignMediaItem.content_url`.
- `docs/project/kencleng-integration-map.md` — “Public Campaign Detail” row.

## Area 1 — Root Caddy same-origin `/api` boundary

### Current state

- `Caddyfile` sends `handle /api/*` to
  `host.containers.internal:8090`, and all other requests to the native
  frontend at port `3000`.
- That handler forwards the path unchanged. The established topology authority
  records this explicitly as a known caveat:
  `docs/project/kencleng-repo-setup.md` §8 says `/api` is not stripped before
  the backend receives it.
- The currently wired backend mux is rooted at paths such as `/healthz`,
  `/auth/...`, `/account/...`, `/docs`, and `/openapi.yaml`; it does not mount
  an `/api` prefix (`backend/cmd/server/main.go`, server mux setup). The
  backend Swagger development helper bypasses Caddy for the same reason
  (`backend/internal/transport/http/swagger.go`, `devServerURL`).
- The proxy configuration contains no response-header manipulation. The
  actual Caddy runtime was not inspectable: this environment has no `docker`
  executable (`docker compose ps` returned `command not found`).

### Requirement

The OpenAPI root declares server URL `/api` (`api/openapi/index.yaml`,
`servers`), and the public media schema permits only the controlled
same-origin `/api/campaigns/{campaignId}/media/{mediaId}/content` URL. The
integration map assigns `/api` routing and preservation of `private, no-store`
to the backend/topology coordination boundary.

### Gap

The configured root proxy cannot currently deliver the contract path shape to
the backend route shape: a browser request to the contract URL is forwarded
with `/api` still present, while existing backend routes are unprefixed. The
current setup documentation and Swagger bypass both corroborate this as a
live known mismatch. No runtime evidence yet proves that a proxied success or
public error response preserves the contract cache header.

### Sniffing

| Lens | Evidence / concern |
|---|---|
| Risk | Every same-origin API request, including the controlled-media URL, can fail at the proxy boundary before backend authorization/retraction checks run. If a later proxy-level error is substituted, it is not the contract’s `404`/`503` Problem Details response or cache guarantee. |
| Edge cases | Requests to `/api` without a trailing slash are not selected by the visible `handle /api/*` matcher. Backend unavailable, malformed paths, and non-media API paths need observation through the actual proxy, not only a successful happy path. |
| Miscontext | `url: /api` in OpenAPI and the correct `content_url` pattern can appear integration-ready while backend Swagger succeeds only because it rewrites the server URL to direct `:8090` access. |
| Misleading signals | Caddy’s presence and a reverse-proxy target do not establish path compatibility; neither the Caddyfile nor Compose starts the native backend/frontend processes. |
| Inconsistency | `api/openapi/index.yaml` promises `/api`, whereas `Caddyfile` preserves `/api` and current backend registration uses unprefixed routes. `docs/project/kencleng-repo-setup.md` records, rather than resolves, that contradiction. |

### Code / authority anchors

| Anchor | Why it matters later |
|---|---|
| `Caddyfile` — `handle /api/*` | Root-owned request-path behavior to re-open for a topology change and proxy verification. |
| `docker-compose.yml` — `caddy` service | Defines the container image, Caddyfile mount, and host entry port `8080`. |
| `docs/project/kencleng-repo-setup.md` §1 and §8 | Topology rationale and recorded prefix caveat; not a substitute for executable verification. |
| `api/openapi/index.yaml` — `servers` | Contract’s same-origin API base. |
| `api/openapi/campaign.yaml` — `PublicCampaignMediaItem.content_url`, `getPublicCampaignMediaContent` | Exact route and cache-bearing response expectation that the proxy must carry. |
| `backend/cmd/server/main.go` — root `ServeMux` registrations | Consumer path shape; backend ownership remains with `WU-S1-003`. |
| `backend/internal/transport/http/swagger.go` — `devServerURL` | Existing direct-backend workaround that must not be mistaken for real proxy integration. |

## Area 2 — Docker Compose and MinIO public/private policy

### Current state

- `docker-compose.yml` defines MinIO with persistent volume
  `kencleng_miniodata`, host API port `9087`, and console port `9088`.
- `minio-init` creates `kencleng-public` and `kencleng-private`, then executes
  `mc anonymous set download local/kencleng-public`. The Compose source has no
  matching anonymous-read grant for `kencleng-private`.
- The native backend configuration names both buckets through
  `MINIO_BUCKET_PUBLIC` and `MINIO_BUCKET_PRIVATE` (`.env.example` and
  `backend/cmd/server/main.go`).
- The actual persisted policy and bucket contents were not observable in this
  session: Docker/Compose is unavailable and the named volume may retain
  policy from earlier initialization.

### Requirement

Slice-1 Campaign media must stay private in object storage. Public access is
only through the controlled same-origin byte operation; a public bucket,
direct object URL, redirect, and long-lived signed URL are forbidden for this
media. Other storage uses remain deferred to their owning slice
(`docs/project/kencleng-backend-tech-stack.md`, File Storage;
`docs/spec/4-campaign/features/03-campaign-media.md`, Summary and Deferred
historical breadth).

### Gap

The topology provides a private bucket but does not encode a Campaign-media
bucket assignment or a policy assertion that Campaign objects cannot be put in
the anonymously readable bucket. Its persisted runtime policy is unverified.
Because the public bucket remains deliberately provisioned for unrelated
deferred storage, removing or globally changing it cannot be inferred from
the Slice-1 requirement.

### Sniffing

| Lens | Evidence / concern |
|---|---|
| Risk | A Campaign object written to `kencleng-public` becomes independently retrievable from MinIO and bypasses parent/member rechecks and retraction. Persistent volume state can retain a prior policy even after source changes. |
| Edge cases | First initialization, a reused `kencleng_miniodata` volume, failed or delayed `minio-init`, absent bucket, changed environment bucket names, and a backend starting before initializer completion produce materially different states. |
| Miscontext | Having bucket names labelled `public` and `private` does not bind a given data class to either one; current startup only verifies that both exist. |
| Misleading signals | `mc anonymous set download` appears limited to the public bucket in source, but it does not prove the policy currently stored in a pre-existing MinIO volume or how future Campaign code chooses its bucket. |
| Inconsistency | The active architecture says Slice-1 Campaign media is private, while Compose continues to make a public bucket available and no executable Campaign-storage binding exists yet. This is an unimplemented boundary, not authority to remove unrelated public storage. |

### Code / authority anchors

| Anchor | Why it matters later |
|---|---|
| `docker-compose.yml` — `minio`, `minio-init`, `kencleng_miniodata` | Authoritative local service/policy initialization and persistence topology. |
| `.env.example` — MinIO variables | Native backend endpoint and named-bucket expectations; do not copy local credentials into artifacts. |
| `backend/cmd/server/main.go` — `requireEnv`, `initMinIO` | Shows startup requires both bucket names and only checks existence. |
| `docs/project/kencleng-backend-tech-stack.md` — File Storage row | Reconciled private-Campaign-media architectural constraint. |
| `docs/spec/4-campaign/features/03-campaign-media.md` — Summary | Active storage/public-delivery rule. |

## Area 3 — Backend storage and controlled-media runtime boundary

### Current state

- The server creates a MinIO client in `initMinIO`, calls `BucketExists` for
  both configured buckets, and then discards the client. It neither supplies a
  storage dependency to a Campaign domain nor performs object reads/writes.
- `backend/internal/platform/storage/doc.go` declares the storage package an
  empty skeleton whose implementation belongs to a later task.
- No Campaign domain package, Campaign HTTP handler, Campaign route
  registration, object read, redirect, signed-URL generation, or media
  persistence implementation exists under `backend/` at the inspected target
  revision.
- The current public contract has already defined the public controlled-media
  operation, including binary-only JPEG/PNG success and cache-bearing
  `404`/`503`, but this is contract evidence only.

### Requirement

Every content request must authorize public eligibility and Campaign/media
membership at its origin. The origin must read only eligible JPEG/PNG bytes,
not reveal an object-storage URL or redirect, distinguish unavailable eligible
storage/object state as `503`, and put `private, no-store` on every specified
public boundary response. These delivery semantics are backend-owned by
`WU-S1-003`; this Run only identifies the topology/storage dependency.

### Gap

There is no runtime controlled-media delivery path to integrate with Caddy or
MinIO yet. The existing MinIO connectivity check cannot prove object privacy,
content type enforcement, retraction, parent/member recheck, response
classification, or cache-header behavior. The storage contract between
`WU-S1-003` and this topology Work Unit is therefore not yet executable.

### Sniffing

| Lens | Evidence / concern |
|---|---|
| Risk | If object access is later connected without an explicit private-bucket boundary, a direct/public object path can silently bypass the anti-enumeration and revocation properties. If dependency errors are flattened, consumers may receive false absence rather than the required `503`. |
| Edge cases | Object missing after eligible metadata lookup, storage outage/time-out, changed Campaign visibility between requests, media linked to another Campaign, unsupported object MIME type, and client disconnect during byte streaming each need backend-owned runtime evidence. |
| Miscontext | Successful application startup proves only MinIO endpoint credentials and bucket existence, not controlled Campaign-media availability. |
| Misleading signals | The MinIO SDK dependency, two bucket variables, and a storage package directory can look like completed storage integration, but the client has no retained consumer and the storage package has no implementation. |
| Inconsistency | The contract and active threat model require controlled origin delivery now, while current backend contains only storage scaffolding. This is expected downstream work under `WU-S1-003`, not a contradiction that topology work can resolve alone. |

### Code / authority anchors

| Anchor | Why it matters later |
|---|---|
| `backend/cmd/server/main.go` — `initMinIO` | Existing initialization and failure boundary; verify its interaction with real storage work rather than treating it as delivery. |
| `backend/internal/platform/storage/doc.go` | Confirms no reusable storage abstraction currently exists. |
| `backend/cmd/server/main.go` — server route registrations | Shows Campaign routes must be introduced by the backend-owned Work Unit before proxy integration can be tested. |
| `docs/spec/4-campaign/features/03-campaign-media.md` — Behavior and Test checklist | Backend/public-boundary assertions to carry into integration verification. |
| `docs/spec/4-campaign/threat-model.md` — “Slice-1 Campaign media delivery” | Security rationale for private origin recheck and cache behavior. |
| `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/manifest.md` — Scope | Ownership rendezvous: it explicitly owns controlled delivery and private-storage integration boundary; it excludes root policy mutation. |

## Area 4 — Frontend/runtime integration boundary

### Current state

- The generated frontend contract type exposes `content_url` as a string and
  preserves its same-origin route example
  (`frontend/lib/api/generated/openapi.ts`, `PublicCampaignMediaItem`).
- The current frontend source is a foundation landing page. It has no Campaign
  route, real API request helper, fetch call, image consumer, Next.js rewrite,
  or production proxy configuration (`frontend/app/page.tsx`,
  `frontend/next.config.ts`).
- The local ignored `frontend/.env.local` sets API mocking for standalone
  frontend development and says it should be disabled/unset when using the
  root same-origin topology. It is not production integration evidence.
- The integration map correctly identifies the opaque `content_url` and names
  private storage, `/api` routing, and no-store as backend/topology work to be
  enforced later.

### Requirement

Frontend consumers must use the opaque same-origin contract URL rather than a
direct object-storage URL. The root topology must make that URL reach the
controlled backend operation. No frontend product behavior is in scope for
this Run.

### Gap

No current frontend runtime consumer exists to exercise the contract through
the root proxy. The generated type is necessary contract correspondence, but
it cannot prove browser-origin routing, image-byte rendering, or no-store
response preservation. This leaves an integration verification dependency,
not frontend implementation ownership for this Work Unit.

### Sniffing

| Lens | Evidence / concern |
|---|---|
| Risk | A future consumer that substitutes an object URL, absolute alternate origin, or a mock-only code branch defeats the controlled same-origin boundary even if the generated type remains present. |
| Edge cases | A media `absent`/`temporarily_unavailable` union state must not initiate a byte request; a failed image request needs the public error behavior to remain non-enumerating. Exact presentation is frontend-owned. |
| Miscontext | Generated OpenAPI declarations may be mistaken for a live API integration although current frontend has no request/data layer for this operation. |
| Misleading signals | `NEXT_PUBLIC_API_MOCKING=true` facilitates independent local development but does not test Caddy, backend, MinIO, header preservation, or controlled-byte delivery. |
| Inconsistency | The contract requires a real same-origin URL, whereas current frontend foundation intentionally has no Campaign consumer. This is sequenced delivery, not a defect in the foundation. |

### Code / authority anchors

| Anchor | Why it matters later |
|---|---|
| `frontend/lib/api/generated/openapi.ts` — `PublicCampaignMediaItem.content_url` | Generated contract correspondence to retain when a frontend owner creates a data boundary. |
| `frontend/next.config.ts` | Current lack of Next-local rewrite/proxy; root Caddy is the designated integration boundary. |
| `frontend/.env.local` | Local-only mock posture; must not be treated as real topology verification. |
| `docs/project/kencleng-integration-map.md` — Public Campaign Detail row | Cross-stack ownership and required runtime coordination. |
| `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/manifest.md` | Frontend behavior ownership (not modified or explored beyond integration need). |

## Stage-2 carry-forward evidence

- Runtime verification cannot yet be performed end-to-end because this
  environment has no Docker CLI and the backend controlled-media operation is
  not implemented. These are evidence limitations, not a claim that the
  topology currently passes or fails a live request.
- A later topology verification needs to observe, through the root host entry
  point, the backend path received for a contract-valid `/api/...` request;
  prove the public/private MinIO policy state with the persisted volume in
  use; and preserve `Cache-Control: private, no-store` for backend-produced
  `200`, identical public `404`, and eligible-dependency `503` responses.
- Backend integration evidence must cover retraction as a fresh origin fetch:
  after public eligibility or membership is withdrawn, a previously known
  same-origin content URL cannot produce new bytes. Already client-held bytes
  are explicitly outside this guarantee.
- No material new Product, contract, or storage-exposure decision was found.
  The active authority already settles private Campaign media and controlled
  same-origin delivery. Stage 3 must decide the narrow enablement direction
  without broadening policy for unrelated storage uses.

## Stage-2 status

Gap analysis is complete for the four planned areas. Await Human confirmation
before Stage 3 solutioning.
