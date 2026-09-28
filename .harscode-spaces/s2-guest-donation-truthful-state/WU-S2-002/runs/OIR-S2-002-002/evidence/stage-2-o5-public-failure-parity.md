# Stage 2 Gap Evidence — O5 Public Failure Parity

> Phase/Stage: Exploration / Stage 2  
> Author: `P-S2-002-OIR-002-1` (`KC-EXPLORER`)  
> Created: 2026-09-28  
> Work Unit / Run: `WU-S2-002` / `OIR-S2-002-002`  
> Model / reasoning: invocation dispatch metadata says `gpt-6-luna` / `high`; active runtime model is not independently exposed  
> Session: not exposed  
> Target revision: `02e6bc77aea8630afb3d4ffdff25a20d5af3a535` (clean working tree at inspection)  
> Workflow revision: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`

## Current state

- Product sets one generic public behavior/copy for invalid, missing, and expired guest status links, e.g. “Link status tidak tersedia atau mungkin kedaluwarsa.” Source: `docs/product/mvp-scope.md` line 116 and `docs/product/mvp-delivery-slices.md` §5 line 174. Techplan §13 O5 says transport-level parity and anti-enumeration remain open; shared frontend copy is not proof of transport parity.
- Historical Donation API operation lists `401` for missing/mismatched token and also `404` via common NotFound for a missing donation (`api/openapi/donation.yaml` `/donations/{donationId}/status`). Historical status feature spec instead says absent donation, wrong token, and missing token all return identical `401`. Thus code behavior is contradictory before considering headers/timing.
- The shared `Unauthorized` Problem example says “Access token tidak valid atau sudah kedaluwarsa.” This is a JWT-oriented message for an unauthenticated guest status credential. Shared `NotFound` returns the Problem schema but has no example. The authored API does not define a complete failure-response matrix or cache headers for the guest status operation.
- Historical threat model calls for identical 401 on wrong token versus absent donation and flags explicit runtime testing; it does not settle current Product's 24-hour expiry or all transport surfaces. It is a draft and cannot override the current Product direction.
- The current backend has no Donation route/runtime to inspect. `backend/cmd/server/main.go` wires a per-IP `RateLimit` around `/auth/` only; the status endpoint has no configured rate/abuse behavior. `backend/internal/transport/http/middleware.go` uses `r.RemoteAddr` and notes that behind a proxy it identifies the proxy unless trusted forwarding support is later added. Its limiter sweeps idle entries, but no evidence establishes use for Donation.
- Existing Campaign detail/threat-model material provides a precedent for identical 404 body/cache headers and later timing-parity checks for absent/malformed/non-public resources. It is evidence of a project pattern, not an approved Donation response choice.
- Harscode `restapi/anti-enumeration.md` calls for generic identical responses, constant-time secret comparison, review of response timing, and avoiding error details that leak existence. These are relevant analysis clues, not settled owner choices for this endpoint.

## Requirement and gap

The generic user-facing behavior is fixed, but the current source set does not specify or implement a single response matrix across missing token, invalid token, expired token, and nonexistent donation. Status code, Problem type/title/detail, headers, cache policy, material timing behavior, and abuse/rate handling remain unsettled. Any future secret comparison and rate policy need explicit evidence; a shared message alone cannot satisfy the requirement.

## Sniffing lenses

- **Risk:** Different status codes, Problem bodies, headers, cache behavior, or large timing differences can reveal whether a donation or credential exists. Weak or absent abuse protection can permit repeated online guesses. Proxy address handling can make a per-IP limit ineffective or overly broad.
- **Edge cases:** query parameter absent; empty or malformed token; valid but expired token; valid token with absent donation ID; malformed/non-resolvable donation ID; 24-hour boundary; high-volume requests through one proxy; limiter restart/idle eviction; intermediary cache serving a private response or caching distinguishable errors.
- **Miscontext:** A single frontend error string does not imply API parity. The old “all failures are 401” feature spec conflicts with the API's additional 404 and does not prove body/header/cache/timing equality. Product's generic behavior includes expiry, which the old spec lacks.
- **Misleading signals:** Both 401 and 404 reference a shared Problem schema, but the response codes and examples differ. Existing Campaign anti-enumeration and general auth rate limiting look like controls but do not establish Donation runtime behavior.
- **Inconsistency:** Product generic failure behavior is current; Donation feature spec and OpenAPI disagree on missing-resource status, and the shared 401 example describes an Account access token. This must be reconciled through the API owner with Security review, without importing either historical choice silently.

## Code and authority anchors

- `docs/product/mvp-scope.md` line 116 — generic guest status-link behavior/copy and 24-hour expiry.
- `docs/product/mvp-delivery-slices.md` §5 lines 174–177 — same failure requirement and Security ownership of credential exposure controls.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-007/techplan.md` §13 O5 and §12 R6 — response matrix, parity, and future verification obligations.
- `api/openapi/donation.yaml` `/donations/{donationId}/status`, lines 80–106 — historical `401`/`404` operation responses.
- `api/openapi/common.yaml` `Unauthorized` and `NotFound`, lines 67–96 — shared Problem response definitions; the 401 example is Account-token-specific.
- `docs/spec/5-donation/features/02-donation-status-check.md` §Validation and test checklist — old uniform-401 behavior, including no-expiry assumption.
- `backend/cmd/server/main.go` lines 154–188 — current route wiring and auth-only rate-limit mount.
- `backend/internal/transport/http/middleware.go` `RateLimit` — existing limiter mechanics and proxy-address limitation.
- `docs/spec/4-campaign/features/02-campaign-detail-listing.md` lines 23–25, 39–40, 75–76 and `docs/spec/4-campaign/threat-model.md` lines 26–28 — adjacent anti-enumeration precedent.
- `../harscode-workspace/best-practices/restapi/anti-enumeration.md` — applicable generic response, constant-time comparison, timing, and error-detail guidance.

Stage 2 makes no response-code or limiter choice and accepts no residual risk. The named API/contract owner can decide contract detail after the Stage 3 gate with Security review; residual-risk acceptance remains a Human gate. Runtime parity and abuse evidence belong to later Testing.
