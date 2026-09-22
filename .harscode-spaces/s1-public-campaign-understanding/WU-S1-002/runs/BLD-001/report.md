# BLD-001 — Build Report

> Phase             : Build
> Work Unit         : WU-S1-002
> Run               : BLD-001
> Author            : Codex CLI agent
> Created           : 2026-09-22T09:10:40+07:00
> Model             : `gpt-5.6-terra`
> Reasoning         : `high`
> Session           : Fresh session
> Target revision   : `b8d5a757cdf9943d087cc5e420f7e385c504ecb2`
> Workflow revision : `d46358563942c7e015b97aa7c5767c880ef1bc63`
> Status            : Finding — Build target is not fully satisfied

## What changed

- `docs/spec/4-campaign/{tasks.md,features/02-campaign-detail-listing.md,features/03-campaign-media.md,invariants.md,threat-model.md}` → merekonsiliasi Slice 1 menjadi public-only `GET /campaigns/{campaignId}` dengan exact public projection, anti-enumerating `404`, plain-text organizer provenance, decimal funding truth, unavailable donation action, serta controlled media delivery. Listing/upload/privileged historical breadth ditandai `DEFER`.
- `api/openapi/{index.yaml,campaign.yaml}` → memperbaiki root `bearerAuth` ref; mendaftarkan `GET /campaigns/{campaignId}/media/{mediaId}/content`; mengganti public detail dengan `getPublicCampaignDetail`/closed `PublicCampaign*` schemas; menambah `getPublicCampaignMediaContent`, local no-store Problem responses, dan menghapus anonymous attachment-list `GET`. Attachment upload historis tetap ada tetapi eksplisit di luar Slice 1.
- `api/openapi.yaml` → regenerated dari split authored sources.
- `frontend/package.json` dan `frontend/lib/api/generated/openapi.ts` → menambah `generate:api-types` dan artifact TypeScript hasil `openapi-typescript`; kedua operationId dan public schemas tersedia tanpa model response tulisan tangan.
- `api/README.md`, `docs/project/kencleng-backend-tech-stack.md`, dan `docs/project/kencleng-integration-map.md` → menyelaraskan authored-vs-generated authority, private controlled Campaign media, dan mapping Public Campaign Detail pertama.
- `docs/project/kencleng-development-tracker.md` → sengaja **tidak diubah**. Techplan mensyaratkan tracker diubah terakhir dan hanya setelah semua Build evidence lulus.

## Tests run

- `cd api && npm run validate` → Build contract validation → **PASS**: zero errors, 126 warnings. Baseline pada Run ini adalah 1 error/130 warnings; error `bearerAuth` terselesaikan dan tidak ada warning baru di touched coordinates.
- `cd api && npm run bundle` → Build generated-artifact verification → **PASS**: `api/openapi.yaml` regenerated.
- `cd frontend && npm run generate:api-types` → Build generated consumer artifact → **PASS**: menghasilkan `lib/api/generated/openapi.ts`.
- `cd frontend && ./node_modules/.bin/tsc --noEmit` → required Build type check → **FAIL**: `app/page.test.tsx` tidak mengenali `describe`, `it`, dan `expect`. `vitest.config.ts` sudah memakai `globals: true`, tetapi `tsconfig.json` tidak memasukkan deklarasi globals Vitest. Diff dari target revision menunjukkan konfigurasi/test tersebut tidak berubah pada Run ini; kegagalan adalah pre-existing verification/configuration gap, bukan error dari generated OpenAPI artifact.
- `cd frontend && npm run lint` → focused frontend static verification → **PASS**.
- `git diff --check` → patch hygiene → **PASS**.

## Verification scope confirmation

Tidak ada test race/concurrency, performance/load, atau security-class yang dijalankan pada Build ini. Tidak ada broad Testing-owned suite yang dijalankan. Pemeriksaan parity runtime optional-auth, retraction melalui private MinIO/proxy, timing anti-enumeration, dan semantic cross-file sweep tetap milik Testing/downstream implementation sesuai Techplan.

## Contract check

- [ ] Current build target satisfied in full — tidak, karena required `tsc --noEmit` gagal pada konfigurasi test globals yang sudah ada dan tracker tidak boleh dipromosikan sebelum evidence lengkap.
- [x] Live-code re-grounding did not invalidate a material API/product/security assumption — source anchors tidak berubah sejak target Techplan; perubahan HEAD hanya artifact orchestration. Finding terbatas pada executable frontend verification assumption.

## Deferred / not tested here

- Independent Testing: semantic allowlist/closed-object sweep, generated artifact correspondence via independent `/tmp` regeneration, optional-auth/error parity, money/media edge inspection, and path-by-path authority consistency.
- Downstream backend/topology/integration: runtime public eligibility, forbidden-field mapping, timing parity, controlled private-storage retrieval/retraction, cache-header preservation, storage `503`, decimal calculation, and safe text rendering.
- Human: final Indonesian provenance/action wording and milestone acceptance.

## Flagged for Techplan / Testing

**Finding — required frontend TypeScript verification is not currently executable.** `frontend/tsconfig.json` includes `app/page.test.tsx` but does not supply Vitest global declarations, while the test runner itself is configured with `globals: true`. This is outside the Techplan's authorized file list (other than the types-generation script) and prevents claiming `CONTRACT_READY`. A scoped corrective decision/patch plan must either make the designated `tsc --noEmit` command executable or revise the approved verification contract with equivalent evidence; Build must not silently bypass it.

## Phase handoff

- Completed: all authorized reconciliation edits, split-contract validation, bundle regeneration, generated TypeScript artifact, lint, and patch hygiene.
- Artifacts: `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/BLD-001/report.md`; `api/openapi.yaml`; `frontend/lib/api/generated/openapi.ts`.
- Human decision: required only if the owner chooses a verification-contract change rather than authorizing a narrowly scoped frontend TypeScript test-globals correction.
- Open / deferred: Build target remains incomplete; `CONTRACT_READY` was not claimed and tracker remains unchanged pending the `tsc` Finding.
- Recommended next step: route the Finding to Techplan/authorized corrective scope, then start a new Build/Patch Run for the narrow resolution before independent Code Review and Testing.
- Session transition: stop this Build session; a fresh planning/corrective session is needed because the remaining choice changes authorized scope/verification evidence.
- Context pointers: `TP-001/techplan.md`, this report, `frontend/tsconfig.json`, `frontend/vitest.config.ts`, `frontend/app/page.test.tsx`, and the changed contract files only.
