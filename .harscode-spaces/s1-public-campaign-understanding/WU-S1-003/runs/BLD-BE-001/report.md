# BLD-BE-001 — Build Report

> Phase: Build
> Work Unit: `WU-S1-003`
> Run: `BLD-BE-001`
> Author: Codex CLI agent
> Role: Implementer
> Specialization: Backend / Go / Campaign public delivery
> Session: Fresh Build session
> Created: 2026-09-23
> Model / Reasoning: `gpt-5.6-terra` / `high`
> Target revision: `a053b48ef3fef35a073c5fe57e5ee4581f848e07`
> Workflow revision: `d46358563942c7e015b97aa7c5767c880ef1bc63`

## What changed

- `backend/migrations/000011_create_public_campaigns.*.sql` → menambah state persisted minimum `organizations`, `campaigns`, dan `campaign_media`, termasuk FK, funding nullable-pair/non-negative constraints, MIME/order constraints, dan down migration child-before-parent.
- `backend/internal/domain/campaign/` → menambah repository `goqu`/`pgx` dengan predicate public `published`, closed public projection, decimal `shopspring/decimal` untuk funding/progress tanpa `float64`, recheck parent/member media, serta sentinel `404`/`503` boundary.
- `backend/internal/platform/storage/private_reader.go` → menambah MinIO reader/seed writer untuk satu private bucket; reader menilai object/type sebelum stream dan tidak menyediakan public/signed URL atau redirect.
- `backend/internal/transport/http/campaign_public.go` → menambah dua handler GET tanpa auth middleware, exact public wire model, no-store pada `200`/`404`/`503`, 404 body bersama untuk malformed/absent/non-public/non-member, dan stream JPEG/PNG yang menutup reader.
- `backend/cmd/server/main.go` → mempertahankan client MinIO setelah bucket checks, inject private reader ke Campaign service, dan mendaftarkan direct backend paths `/campaigns/...` tanpa mengubah Caddy/path `/api` ownership.
- `backend/cmd/campaign-seed/` → menambah operator-only manifest command dengan UUID eksplisit, default collision failure, opt-in `--replace`, validasi funding/media state dan JPEG/PNG lokal, opaque object key, serta cleanup best-effort. Tidak ada startup seed.
- `backend/go.mod`, `backend/go.sum` → menambah direct dependency `github.com/shopspring/decimal`.

## Tests run

- `GOCACHE=/tmp/kencleng-go-build go test ./internal/domain/campaign ./internal/platform/storage ./cmd/campaign-seed ./cmd/server` → Build unit/compile verification → PASS. Menjalankan unit projection/funding/media classification dan manifest validation; storage/server package compiled.
- `GOCACHE=/tmp/kencleng-go-build go test ./internal/transport/http` → Build HTTP handler/wire verification → PASS. Memeriksa exact top-level public projection, no-store, shared malformed 404 body, safe 503, JPEG stream, dan reader close.
- `go vet ./internal/domain/campaign ./internal/platform/storage ./internal/transport/http ./cmd/campaign-seed ./cmd/server` → static targeted verification → PASS.
- `git diff --check` → patch integrity → PASS.

## Verification scope confirmation

Tidak ada race/concurrency, performance/load, atau security-class test yang dijalankan pada iterasi Build ini. Tidak ada broad Testing-owned suite yang dijalankan; `go test ./...`, `go test -race ./...`, `make verify`, testcontainers Postgres/MinIO, dan topology verification sengaja ditinggalkan untuk independent Testing/Human sesuai Testing Checklist.

## Contract check

- [x] Current build target satisfied in full
- [x] Live-code re-grounding did not invalidate a material contract assumption

## Deferred / not tested here

- Isolated Postgres migration apply/down/re-up, constraint behavior, dan repository runtime scan behavior: independent Testing menggunakan disposable testcontainers, bukan shared/manual DB.
- Isolated MinIO adapter against private bucket, cancellation, dan missing-object classification: independent Testing.
- Broad regression, race, contract/security toolchain, dan `make verify`: independent Testing.
- Caddy `/api` stripping, persisted MinIO anonymous policy, proxy preservation of no-store, root retraction scenario, dan frontend/browser integration: `WU-S1-005`/downstream integration/Human. Tidak ada klaim topology atau `INTEGRATED_VERIFIED`.

## Flagged for Techplan / Testing

Tidak ada material assumption break. Testing perlu menjalankan testcontainers evidence yang sudah ditetapkan Techplan sebelum capability atau migration evidence dipromosikan.

## Phase handoff

- Completed: minimum persisted Campaign backend, closed public detail projection, controlled private media delivery seam, direct public transport/routes, dan operator seed command untuk `WU-S1-003`.
- Artifacts: `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/BLD-BE-001/report.md`
- Human decision: none
- Open / deferred: independent Testcontainers/topology evidence seperti di atas.
- Recommended next step: Code Review setelah initial Build, kemudian independent Testing sesuai Techplan.
- Session transition: start fresh Code Review/Testing untuk independence; gunakan parent Techplan, Build report ini, diff backend Campaign, dan test terarah sebagai context pointers.
