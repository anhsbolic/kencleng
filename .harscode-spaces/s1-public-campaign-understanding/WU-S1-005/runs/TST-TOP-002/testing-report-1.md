# Testing Report 1 — Joint Runtime Verification Re-entry

> Phase: Testing  
> WORK_UNIT_ID / RUN_ID: `WU-S1-005` / `TST-TOP-002`  
> Role / Specialization: Verifier / Joint Testing re-entry — controlled media proxy/retraction closure  
> Participant: Codex CLI agent  
> Created: 2026-09-23  
> Target revision: `855a9f7d6fa57b35421ed893da223420653ad4c5`

## 0. Verdict

**Pass — runtime evidence for R4, Campaign-specific R5, and R7 is closed.**

The browser-facing root route `localhost:8080` delivered eligible private PNG
bytes through Caddy, preserved the Campaign cache contract for `200`, all
tested public `404` cases, and an induced eligible dependency `503`. A known
opaque `content_url` stopped delivering bytes on fresh requests after each of
two separate test-only persisted withdrawals: parent public eligibility and
media membership. No production code, API, proxy configuration, or product
operator capability was added or changed; this run used a disposable runtime
fixture only.

## 1. Isolated fixture and runtime method

- The existing root Caddy (`localhost:8080`) and configured private MinIO
  bucket were used. The MinIO initializer/persisted policy had already been
  independently verified in `TST-TOP-001`.
- A new Podman PostgreSQL 16 container, `tst-top-002-postgres`, was started
  on host port `15435`; it was not the Compose database or a manual/shared DB
  migration target. Migrations were applied only to this disposable container.
- Backend was run locally on `:8090` with `DATABASE_URL` pointing at that
  disposable database and harmless process-only OAuth dummy values needed by
  the existing server startup validation. Caddy therefore reached the native
  backend through its normal `host.containers.internal:8090` upstream.
- `campaign-seed` persisted only fixture IDs ending in `...000005`, with a
  clearly test-only title/story and a valid local PNG. It stored one generated
  `campaign/...` object in `kencleng-private`; no shared/manual Campaign state
  was read or changed.
- Parent status and media-row withdrawal were direct, narrowly scoped SQL
  mutations of this disposable fixture database. They are Testing seams, not
  product/operator features. The seed command was used to restore the fixture
  between independent scenarios.
- Every temporary database/container/process, local fixture/output file, and
  final private fixture object was removed after the run. Post-cleanup no
  listener remained on `:8090` or `:15435`.

## 2. Runtime coverage

| Requirement / scenario | Actual observation through root Caddy | Result |
|---|---|---|
| Eligible controlled media (`R4`) | `GET` of the `content_url` obtained from the public detail response returned `200`, `Content-Type: image/png`, `Cache-Control: private, no-store`, and `Via: 1.1 Caddy`. Body was `image/png`, 18,945 bytes, SHA-256 `38037bcf18c7f9b950da2e1a2d8800d38fb27637ccda293d38368a3704b67f18`. | Pass |
| Opaque public response (`R5`) | Detail returned the exact same-origin `/api/campaigns/{id}/media/{id}/content` URL. Its media item contained only `id`, `content_url`, `content_type`, `alt_text`, `caption`, and `source`; serialized detail contained neither MinIO endpoint/name nor an object/signed URL. | Pass |
| Anonymous direct private object (`R5`) | Anonymous request to the fixture's known private object URL on `:9087` returned `403 Forbidden`, had no `Location`, and did not return image bytes. | Pass |
| Public absent / non-member / non-public (`R4`) | All three fresh root requests returned `404`, `application/problem+json`, `Cache-Control: private, no-store`, and `Via: 1.1 Caddy`. Their complete bodies had the same SHA-256: `127665197f7a1f57238268064e188c8f025262a03fbe020e9797b6e5e08a463e`. | Pass |
| Eligible dependency failure (`R4`) | After a successful eligible `200` precondition, the fixture's private object alone was removed with authenticated MinIO test tooling. Fresh root request returned `503`, `application/problem+json`, `Cache-Control: private, no-store`, `Via: 1.1 Caddy`, and no `Location`; body was 182 bytes. | Pass |
| Retraction: parent eligibility (`R7`) | After a prior `200` for the exact `content_url`, only fixture Campaign status was changed to `unpublished`. A fresh request to that unchanged URL returned `404`, `application/problem+json`, `Cache-Control: private, no-store`, `Via: 1.1 Caddy`, and a 165-byte Problem Details body—not PNG bytes. | Pass |
| Retraction: media membership (`R7`) | Fixture was restored and delivered a new `200` precondition. After only its `campaign_media` row was deleted, a fresh request to the same known URL returned the same `404` contract/body hash as above, no redirect, and no image bytes. | Pass |

## 3. Commands and evidence boundaries

Commands actually executed included:

```text
podman run ... postgres:16                       # disposable database only
migrate -path backend/migrations -database <disposable-db> up
go run ./cmd/server                              # DATABASE_URL overridden to disposable DB
go run ./cmd/campaign-seed --manifest <test-only-json> --media-file <PNG>
curl ... http://localhost:8080/api/campaigns/.../media/.../content
podman run --rm ... minio/mc ... mc rm <fixture-private-object>
```

The `503` trigger was an eligible object removal, not a Caddy upstream outage:
the preceding request was a valid `200`, and the resulting `503` retained the
backend Problem Details/cache response through Caddy. Existing R1–R3/R6
topology and persisted-policy evidence remains in `TST-TOP-001`; this run did
not repeat unrelated frontend or broad repository suites.

## 4. Error handling and residual scope

- No redirect, direct object URL, or signed URL was observed in any public
  success/error/retraction response.
- The direct MinIO `403` is bucket-policy evidence. The public response
  inspection and parent/member fresh-fetch tests supply the distinct
  Campaign-origin/retraction evidence.
- The guarantee tested is a new request after withdrawal. Bytes obtained by a
  client before withdrawal are intentionally outside R7's stated guarantee.
- The repository-wide pre-existing `gosec` findings recorded by
  `TST-TOP-001` and `TST-BE-002` remain outside this no-production-change
  runtime run; no aggregate `make verify` is claimed here.

## 5. Handoff

- No production defect, missing test seam, or Testing patch plan was found.
- WU-S1-005 may use this report as its proxy/policy/retraction runtime
  closure evidence, subject to the existing Tier-1 human review/milestone
  gates.
