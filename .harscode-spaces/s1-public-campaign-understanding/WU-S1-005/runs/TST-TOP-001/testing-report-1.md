# Testing Report 1 — Root Topology & Controlled Media Enablement

> Phase: Testing  
> WORK_UNIT_ID: WU-S1-005  
> RUN_ID: TST-TOP-001  
> Role / Specialization: Verifier / Independent Testing — root topology and controlled media enablement  
> Participant: Codex CLI agent  
> Created / Updated: 2026-09-23  
> Model / Reasoning / Session: gpt-5.6-terra / high / Fresh independent Testing session  
> Target revision: 6b7a103  
> Workflow revision: not exposed by the invocation

## 0. Sweep Summary

- Confirmed: R1 source claim was independently exercised after a Caddy reload: the real native backend returned `200 {"status":"ok"}` for both direct `/healthz` and root `localhost:8080/api/healthz`; the latter contained `Via: 1.1 Caddy`.
- Confirmed: R2 was exercised with the real Next.js development server at `:3000`; root `localhost:8080/` returned its frontend document with `Via: 1.1 Caddy`, and an unknown non-API path returned Next.js `404`, not a backend response.
- Closed from the Build deferred list: rendered `podman-compose config`, runtime Caddy validation, persisted MinIO policy inspection, and policy convergence on the existing `kencleng_kencleng_miniodata` volume.
- Confirmed: `minio-init` was force-recreated against the existing volume and exited `0`; authenticated `mc anonymous get` then reported `download` for `kencleng-public` and `private` for `kencleng-private`.
- Confirmed: a temporary, authenticated private-bucket object was denied to an anonymous direct MinIO request (`403`, no redirect); the exact probe object was removed and absence checked. This establishes bucket-policy enforcement, not the Campaign-object/public-response half of R5.
- Still requires fresh Testing: R4 success and public-404 media scenarios, the Campaign-object/public-response portion of R5, and R7 require a seeded eligible Campaign-media fixture plus persisted eligibility/membership mutation capability from WU-S1-003. No integrated or retraction claim is made.

An initially running Caddy instance was serving its previously loaded configuration despite the bind-mounted Caddyfile already containing `handle_path`; the first probe therefore delivered `/api/healthz` unchanged. `caddy reload --config /etc/caddy/Caddyfile --adapter caddyfile` applied the validated current configuration, and the same probe then reached the actual backend successfully. This is deployment/runtime staleness of an instance started before the source change, not a source defect; an already-running instance must be reloaded or recreated to apply it.

## 0a. Test Focus Pointer Execution

| Area | Evidence anchor opened | Specialized verification | Result |
|---|---|---|---|
| Private bucket / direct bypass / retraction | `EXP-TOP-001/evidence/solutioning.md#runtime-verification-required-not-performed-here`, items 2, 5–6 | Inspected running policy, reran initializer on reused volume, then requested a temporary private object anonymously; removed probe afterward. | Policy/reused-volume/direct-bucket denial passed. Campaign fixture and retraction remain deferred. |
| Proxy cache and public error preservation | `EXP-TOP-001/evidence/gap-analysis.md#stage-2-carry-forward-evidence` | Sent a well-formed nonexistent media URL through root Caddy while actual backend was running. | `503 application/problem+json` and `Cache-Control: private, no-store` reached client unchanged. This is a dependency-unavailable path, not proof of required eligible 200/404 matrix. |
| Same-origin API path boundary | `EXP-TOP-001/evidence/gap-analysis.md#area-1--root-caddy-same-origin-api-boundary` | Validated loaded Caddyfile; used an upstream path probe, then actual backend `/healthz`; exercised Next frontend fallback. | Passed after Caddy reload. |

No sensitive area was missing from the Techplan pointer; no Techplan drift found.

## 1. Test Coverage

| Rule / scenario | Category | Observable verification | Result |
|---|---|---|---|
| R1 — `/api` path translation | Happy / integration | `podman-compose config`; `caddy validate`; actual backend direct `/healthz` and proxied `/api/healthz`. Before reload an upstream probe exposed stale loaded configuration; after reload proxy response was backend `200` with `Via: 1.1 Caddy`. | Pass |
| R2 — frontend preservation | Happy / negative | Native `npm run dev -- --hostname 0.0.0.0`; root and unknown non-API requests through `:8080` returned Next.js (`200`/`404`) with `Via: 1.1 Caddy`. | Pass |
| R3 — private-policy convergence | Happy / reused-volume edge | Existing named volume was retained; `podman-compose up --force-recreate minio-init` exited `0`; authenticated policy inspection reported public `download`, private `private`. | Pass |
| R4 — controlled boundary preservation | Error propagation | Actual backend response through root for a well-formed unavailable media URL was `503 application/problem+json`, `Cache-Control: private, no-store`, no redirect. | Partial — 503 preservation passed; eligible 200 and public 404 matrix need fixture. |
| R5 — no topology ownership leak / direct bypass | Scope / security | `git diff --check` passed. Build/review scope remains Caddyfile, Compose, and setup docs. Anonymous request to a temporary known private object returned `403`, no redirect; object then deleted. | Partial — direct policy denial passed. Seeded Campaign object and public response inspection unavailable. |
| R6 — documentation truth | Documentation | Re-read setup §8 against rendered Caddy/Compose. It names path translation, private/public policy, deferred runtime/integration/retraction proof, and makes no false verification claim. | Pass |
| R7 — retraction through proxy | Retraction / security | No seeded eligible content URL or persisted parent/membership withdrawal fixture was available in this run. | Deferred / not tested — mandatory follow-up, not removed. |

## 2. Error Verification

| Error case | Expected behavior/category | Actual | Actionable/propagated correctly? |
|---|---|---|---|
| Caddy before reload | Existing instance may retain old loaded configuration; source validation alone must not be treated as runtime proof. | Probe received `unexpected=/api/healthz`; this exposed the stale process configuration. | Yes — diagnosed, Caddy reloaded, then R1 passed. |
| Unavailable media dependency via root | Public backend `503` must retain Problem Details and `Cache-Control: private, no-store`; Caddy must not synthesize it. | `503`, `application/problem+json`, contractual no-store header, `Via: 1.1 Caddy`. | Yes for this 503 path. |
| Direct anonymous private object request | Access must be denied without redirect/object URL. | `403`, empty redirect value, XML MinIO denial body; no media bytes. | Yes. |
| Public absent/non-public/non-member media | Identical public `404` plus no-store propagated through root. | N/A — no eligible/non-eligible persisted Campaign fixture supplied. | Deferred to WU-S1-003 + Testing joint run. |

## 3. Final Verification

- Target repo required final build/lint/test commands: `make verify` was run with writable temporary Go/staticcheck caches. It failed in backend `gosec` on 12 already-recorded Account/auth/OAuth/breachcheck/server findings, so later targets did not run. This root work unit did not change those areas; the current WU-S1-003 patch report independently records the same repository-wide blocker.
- Broad checks intentionally not rerun: no additional race/load suite; this root configuration change has no concurrency behavior and the Techplan does not assign one to this run. Campaign integration-tag tests are owned/re-entered by WU-S1-003.
- Migration/schema collision: N/A — WU-S1-005 has no schema or migration change.
- Backward compatibility: passed at topology level: public `/api` remains browser-facing, native backend still receives unprefixed `/healthz`, public bucket policy remains download, bucket names/ports/volume unchanged, and frontend fallback remains reachable.
- Broader-suite requirement for cross-cutting change: `make verify` was attempted as above; its documented unrelated gosec blocker prevents a green aggregate claim.
- Fresh Techplan consistency read: completed end-to-end. No contradiction found. R4/R5/R7 remain expressly deferred until the Campaign runtime fixture/mutation capability exists.

## 4. New Recurring Bug Patterns

None. The stale Caddy process is deployment-state evidence rather than a reusable code defect; it confirms why rendered/current-runtime validation is required after a configuration change.

## Verdict

**Pass with flagged follow-ups.** R1–R3 and R6 have independent runtime evidence; R4 has a real 503 propagation check and R5 has real private-bucket anonymous-denial evidence. No production defect was found in this work unit, so no Testing patch plan is required.

Blocking for an integrated/retraction completion claim: WU-S1-003 must provide a seeded eligible Campaign media URL and persisted parent-eligibility/media-membership withdrawal mutation. Then rerun R4’s 200/404/503 matrix, Campaign-specific R5 checks, and R7 fresh-fetch proof through `localhost:8080`. The 12 aggregate `gosec` findings are pre-existing, out of WU-S1-005 scope, and remain a separate repository-wide final-verification blocker.

## Phase handoff

- Completed: independent root Caddy path/fallback and persisted MinIO policy testing, including reused-volume convergence and direct anonymous private-object denial.
- Artifacts: `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/TST-TOP-001/testing-report-1.md`.
- Human decision: none for the root source change. Tier-1 review/gates remain as required before milestone promotion.
- Open / deferred: WU-S1-003 fixture/mutation-dependent R4/R5/R7 integration evidence; repository-wide existing `gosec` findings.
- Recommended next step: coordinate a focused fresh Testing re-entry after WU-S1-003 exposes the fixture/mutations; do not claim `INTEGRATED_VERIFIED` or retraction verification meanwhile.
- Session transition: use a fresh joint Testing session for that integration re-entry, because it requires live seeded state and observing cross-work-unit behavior; no Build/Patch session is needed from this run.
- Context pointers: approved `TP-TOP-001`, this report, `BLD-TOP-PATCH-001/patch-report-1.md`, `CR-CONF-TOP-001/review-findings.md`, and WU-S1-003’s current Testing evidence.
