# CR-001 — Patch Plan 1

> Phase             : Code Review handoff to Build/Patch
> Work Unit         : WU-S1-002
> Run               : CR-001
> Author            : Reviewer
> Created/Updated   : 2026-09-22T10:36:27+07:00
> Model             : `gpt-5.6-terra`
> Reasoning         : `high`
> Session           : Fresh independent session
> Source finding    : `CR-001-F01` in `review-findings-1.md`

## Objective

Membuat same-origin controlled media path pada `PublicCampaignMediaItem.content_url` menjadi constraint executable, tanpa menambah route, runtime behavior, field publik, atau topology/storage implementation.

## Authorized patch

1. Di `api/openapi/campaign.yaml`, pada `components.schemas.PublicCampaignMediaItem.properties.content_url`, pertahankan `type`, `format`, deskripsi, dan contoh yang ada; tambahkan `pattern` yang hanya menerima:

   ```text
   /api/campaigns/<UUID>/media/<UUID>/content
   ```

   Pattern harus di-anchor dengan `^` dan `$`, menolak scheme/host, query, fragment, trailing path, serta non-UUID segment. Ia tidak perlu mencoba mengikat kedua UUID ke `id`/parent sibling karena constraint lintas-field itu tetap runtime-owned.

2. Regenerasi `api/openapi.yaml` dari split source dengan command resmi.
3. Regenerasi `frontend/lib/api/generated/openapi.ts` dari aggregate dengan `npm run generate:api-types`.

## Required focused verification

- `cd api && npm run validate` — zero errors dan tidak ada warning coordinate baru yang disentuh.
- `cd api && npm run bundle` — aggregate regenerated.
- `cd frontend && npm run generate:api-types` — generated type artifact regenerated.
- Jalankan focused schema assertion terhadap generated bundle: contoh controlled path valid; `https://bucket.example/object.png`, path tanpa `/api`, query/fragment, dan malformed UUID ditolak oleh `content_url.pattern`.
- Independent regenerate bundle/types ke temporary path dan `cmp` dengan artifacts committed.
- `git diff --check` dan scoped `git status --short`.

`tsc --noEmit` dan `npm run lint` tidak perlu diulang untuk perubahan schema yang tetap menghasilkan `content_url: string`, kecuali generated output unexpectedly changes its TypeScript syntax. Bila itu terjadi, jalankan keduanya sebagai diagnosis dan catat hasilnya.

## Out of scope

- Tidak mengubah backend/frontend runtime, object-storage policy, proxy, service worker, test semantics, Product/MVP/Design authority, atau contract field/operation lain.
- Tidak mengubah `frontend/tsconfig.json`; patch TP-002 sudah diverifikasi dan tidak terkait Finding ini.
- Tidak mempromosikan tracker/milestone atau mengubah Control Surface.
- CR-001-C01 tidak menjadi bagian patch loop ini karena non-blocking.

## Re-review posture

Patch ini material untuk boundary keamanan tetapi tetap lokal pada schema constraint dan generated artifacts. Setelah Build/Patch menunjukkan constraint, regeneration, dan focused negative evidence, lakukan targeted confirmation terhadap CR-001-F01 sebelum merutekan Work Unit ke Testing. Full four-pass Code Review hanya diperlukan jika patch memperluas route/field/behavior atau menambah scope lain.
