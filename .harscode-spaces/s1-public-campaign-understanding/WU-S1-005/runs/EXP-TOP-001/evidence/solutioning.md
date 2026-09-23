# EXP-TOP-001 — Solutioning

> Phase/Stage        : Exploration / Stage 3 — Solutioning
> Work Unit          : `WU-S1-005`
> Run                : `EXP-TOP-001`
> Role               : Explorer
> Specialization     : root Caddy proxy, Docker Compose, MinIO policy, same-origin `/api` routing, controlled Campaign media topology
> Participant        : Codex CLI agent
> Session            : Fresh session
> Created            : 2026-09-23
> Model              : `gpt-5.6-terra`
> Reasoning          : medium
> Target revision    : `16234e0dd6b1a3751e3f7c4b1ea1c91a968320c9`
> Workflow revision  : not exposed by the invocation

## Decision

Choose a narrow root-topology enablement:

1. Make Caddy’s `/api/*` handler strip only the `/api` prefix before proxying
   to the native backend, while retaining the existing same-origin browser URL
   and frontend catch-all handler.
2. Keep both existing MinIO buckets and the current anonymous-download policy
   of `kencleng-public` unchanged. Have Compose initialization explicitly
   assert no anonymous policy for `kencleng-private` on every initialization,
   including a reused persistent volume.
3. Treat `MINIO_BUCKET_PRIVATE` as the existing configuration handoff for
   Slice-1 Campaign media. `WU-S1-003` must select that private bucket when it
   introduces Campaign storage; this Work Unit does not add Campaign domain
   wiring, authorization, object reads, MIME classification, or media
   response handling.

This direction directly satisfies the settled same-origin/private-media
boundary without redesigning unrelated storage or taking backend ownership.

## Rationale

`handle_path /api/*` is the Caddy directive matching the current handler’s
scope while stripping the matched prefix before forwarding. Caddy documents it
as the `handle` equivalent with `uri strip_prefix`; its documented API example
uses the same `/api/*` shape. This maps contract-visible
`/api/campaigns/...` to backend-visible `/campaigns/...` without changing the
OpenAPI server base or `content_url` pattern.

The active Slice-1 authority already decides that Campaign media is private
and controlled by the parent/member-authorizing origin. The existing two-bucket
topology is sufficient for that decision if the private bucket’s anonymous
policy is asserted and the backend explicitly consumes it. Deleting the public
bucket, hiding the local MinIO port, or changing policies for Organization and
other deferred media would broaden this Work Unit beyond its authority.

`mc anonymous set none` is the documented operation for clearing a bucket’s
anonymous policy. Applying it to only `kencleng-private` is idempotent
topology hardening for persistent local state; it does not alter
`kencleng-public`.

## Source-level implementation boundary

The future root-scoped Build should be limited to these source-owned changes:

| Source area | Intended narrow change | Explicit non-change |
|---|---|---|
| `Caddyfile` | Replace the `/api/*` `handle` block with its prefix-stripping `handle_path` counterpart, preserving the same backend upstream and frontend fallback. | No OpenAPI base-path change, backend route-prefix change, Next.js rewrite, CORS policy, response synthesis, or cache override. |
| `docker-compose.yml` — `minio-init` | Preserve bucket creation and the public bucket’s existing anonymous download policy; explicitly set the private bucket anonymous policy to `none` after creation. | No deletion/renaming of buckets or volumes; no policy change for `kencleng-public`; no new external storage service. |
| `docs/project/kencleng-repo-setup.md` §8 | Replace the now-resolved caveat with the executable path invariant and point to the required runtime verification. | Do not turn setup documentation into Campaign-handler or frontend behavior documentation. |

No new bucket environment variable is needed. The existing
`MINIO_BUCKET_PRIVATE` name is the cross-Work-Unit handoff. `WU-S1-003` owns
the later proof that Campaign object writes/reads use it and never expose
direct/object-storage URLs.

## Alternatives considered and rejected

| Alternative | Decision | Rationale / consequence |
|---|---|---|
| Keep `handle /api/*` and add `/api` routes in the backend | Rejected | It shifts a root proxy incompatibility into backend transport ownership, conflicts with the existing unprefixed routing convention, and changes the path boundary for all API routes. |
| Change OpenAPI server URL and `content_url` to remove `/api` | Rejected | It breaks the settled same-origin contract and avoids, rather than fixes, the documented topology mismatch. |
| Add a Next.js rewrite/proxy | Rejected | Root Caddy is the declared topology boundary; a frontend-local proxy would create a competing route path and does not address direct root integration. |
| Make `kencleng-public` private or remove it | Rejected | Slice 1 does not own unrelated storage policy. The active authority only requires Campaign media to use private storage. |
| Create a third Campaign-specific bucket | Rejected | Existing private bucket fulfils the settled requirement. A new bucket increases configuration and migration surface without a new authority requirement. |
| Use public object URLs, redirects, or long-lived signed URLs for media | Rejected | Explicitly forbidden by the reconciled feature, threat model, and OpenAPI contract because they bypass origin recheck/retraction. |
| Add proxy cache rules/header rewriting | Rejected | The contract requires backend-produced `Cache-Control: private, no-store` to pass through. Source config presently has no cache layer; extra cache behavior is unnecessary scope and could undermine retraction guarantees. |

## Consequences and ownership handoff

### Root topology owner

- Owns the narrow Caddy request-path correction and private-bucket anonymous
  policy assertion.
- Must not claim `BACKEND_VERIFIED`, `FRONTEND_MOCK_VERIFIED`, or
  `INTEGRATED_VERIFIED` merely from source configuration.

### `WU-S1-003` backend owner

- Re-open `MINIO_BUCKET_PRIVATE` and the post-change Compose policy before
  implementation; bind Campaign media to that bucket.
- Owns eligibility/membership recheck, byte streaming, allowed MIME types,
  identical `404`, eligible-dependency `503`, no-store headers, retraction,
  and the fact that no direct storage URL/redirect appears in a response.

### `WU-S1-004` frontend owner

- Owns the eventual Campaign consumer. It must consume the opaque same-origin
  `content_url`, not derive a MinIO URL or add an alternate production API
  path.

## Runtime verification required (not performed here)

This session has no Docker CLI, and no controlled-media backend operation yet
exists. The following is mandatory execution evidence on an environment that
can run the Compose topology; it is not satisfied by source inspection:

1. Validate the rendered Caddy configuration and start the local topology.
   Observe that a browser-facing request under `/api/...` reaches the backend
   with the expected unprefixed path; also check an ordinary frontend request
   still reaches the frontend fallback.
2. Inspect the *running persisted* MinIO anonymous policies. Confirm
   `kencleng-private` has no anonymous read policy after initialization and
   `kencleng-public` retains its existing behavior. This check must account for
   a reused `kencleng_miniodata` volume, not only an empty first boot.
3. With a seeded eligible Campaign and implemented backend operation, retrieve
   the contract-valid media URL through `localhost:8080`. Confirm byte status,
   JPEG/PNG content type, and exact `Cache-Control: private, no-store` reach
   the client unchanged.
4. Through the same root endpoint, test absent/non-public parent and
   absent/non-member media for identical public `404` body/header behavior;
   induce an eligible storage/object failure to confirm backend `503` and its
   no-store header are preserved rather than replaced by a proxy error.
5. Verify that a direct anonymous request to the private-bucket Campaign
   object is denied and that no public-detail/media response contains a direct
   object URL, redirect, or signed URL.
6. Perform the retraction scenario as a new fetch via a previously known
   same-origin content URL after parent eligibility or membership is withdrawn.
   It must not deliver new bytes. Previously downloaded client-held bytes are
   outside the contract guarantee.

## Evidence limitations

- The decision is source-level and authority-backed only. Docker/Compose and
  MinIO policy state were not inspected at runtime because `docker` is absent
  in this session.
- The Caddy source was not config-validated by its runtime binary here.
- Backend and frontend runtime implementation is intentionally absent from
  this Run’s scope, so end-to-end HTTP behavior remains deferred.

## External implementation references

- Caddy, [`handle_path` directive](https://caddyserver.com/docs/caddyfile/directives/handle_path): documents prefix stripping and the `/api/*` proxy
  pattern used by this direction.
- MinIO, [`mc anonymous set`](https://docs.min.io/aistor/reference/cli/mc-anonymous/mc-anonymous-set/): documents anonymous bucket policy values, including `none`.

## Phase handoff

- Completed: Stage 2 gap analysis and Stage 3 source-level topology decision for same-origin `/api` compatibility and private Campaign-media boundary.
- Artifacts: `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/EXP-TOP-001/evidence/gap-analysis.md`; `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/EXP-TOP-001/evidence/solutioning.md`.
- Human decision: none.
- Open / deferred: Runtime Caddy validation; running/persisted MinIO policy inspection; controlled-media endpoint, response/cache/retraction evidence; all backend Campaign and frontend product work.
- Recommended next step: Techplan synthesis for the root-scoped topology Build, coordinated with `WU-S1-003` at the `MINIO_BUCKET_PRIVATE` handoff.
- Session transition: Start a fresh Techplan session re-grounded on these artifacts and current root configuration; the concurrent backend Work Unit can change the target revision, and a fresh read will keep the Build contract anchored to live source.
- Context pointers: `Caddyfile`; `docker-compose.yml` (`minio-init` and `caddy`); `docs/project/kencleng-repo-setup.md` §8; `docs/spec/4-campaign/features/03-campaign-media.md`; `api/openapi/campaign.yaml` (`getPublicCampaignMediaContent`, `PublicCampaignMediaItem.content_url`); `backend/cmd/server/main.go` (`initMinIO`); `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/manifest.md`.
