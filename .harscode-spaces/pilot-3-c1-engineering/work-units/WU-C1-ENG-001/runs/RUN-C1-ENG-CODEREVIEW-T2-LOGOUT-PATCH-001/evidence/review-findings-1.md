# Review findings — logout patch

Phase: Code Review. Author: Codex / independent Reviewer, backend auth/session.
Created/Updated: 2026-10-12.
WORK_UNIT_ID: WU-C1-ENG-001. RUN_ID: RUN-C1-ENG-CODEREVIEW-T2-LOGOUT-PATCH-001.
Participant: PARTICIPANT-C1-ENG-T2-LOGOUT-PATCH-REVIEWER-001.
Session: SESSION-C1-ENG-T2-LOGOUT-PATCH-REVIEWER-001 / FRESH — binding invocation; bukan ID thread runtime yang diamati independen.
Selected model/reasoning: gpt-6-sol / medium menurut invocation dan approval; identitas model/effort runtime tidak diattestasi oleh Review ini.
Target revision: `dae55ac44beb0e967b94c2f791d8a0b8d8e72372`, branch `pilot/3-c1-engineering` (diamati).
Workflow revision: `651ce57f044bff23fba99082322cee6c066d25a6` (diamati).
Write scope: hanya evidence Review ini; source dan report Build dipertahankan.

## Binding dan scope

Manifest SHA-256 `71074a0657f102456ad598487dbc253595c76097657e5685bbc2346b92d73860` cocok. Seluruh **12/12** assignment-defining bindings cocok sebelum Review dan diperiksa kembali sebelum evidence ditulis. Tidak ada material drift.

| Input | SHA-256 terverifikasi |
| --- | --- |
| `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-CODEREVIEW-T2-LOGOUT-PATCH-001/scope.md` | `2d2e7c8153dcd3f7561e2a4856213f4aba2fdc893774d98130d1657a222228ad` |
| `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-LOGOUT-PATCH-001/evidence/patch-report-1.md` | `68db135f66a8755c900734bad86fd6dc4e94db9d3bcf0b82cc50a40d40594bc2` |
| `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-AUTH-HARNESS-001/evidence/build-report.md` | `3bbb10516378156635a1705fd3e360789ced6c59e91455b3499a22931971a8e4` |
| `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-LOGOUT-PATCH-001/scope.md` | `99551dca7d83cfdff3e103cbfec7d7322bccc86e4597b92d553d0423ccd9dad3` |
| `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.md` | `6d2e41e177e4adaf74f53a06d81dd568def75019c8dc4e4792d938a869896d79` |
| `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/tasks/T2-backend-identity-session.md` | `dfe17b7a3dc0dc89127b040a67f547506a14a4ab591ed0a6343f1a8dd8bcce37` |
| `api/openapi/index.yaml` | `ddefffbe588459cb0cba9f68dfedaccba3da8a1313f1956511cf628df7199965` |
| `backend/internal/platform/auth/http.go` | `83b575b0e43d17d85303ae7dfeb0273fa2516bb18952a9660da7b80a2d039d63` |
| `backend/internal/platform/auth/harness_test.go` | `c9997611413446641da20b73da2e251459a9ccbb33fbb7c88429f18206d0f45a` |
| `backend/scripts/auth_harness.py` | `04e8173381bb978e9ab52898969418b893672729f82b741645267aaef1e41e53` |
| `backend/scripts/scanner_gate.py` | `ea1aa24035c3532e93b0faa3419883e08306e4bb624d6debae845121d4746cd7` |
| `backend/README.md` | `38d6dbb59f19f7f936ac43932c8d3a2028e9b07585e34e5baf9844b94b41574e` |

Empat pass berikut memakai scope lima file yang sama. Live diff terhadap HEAD dibaca; karena harness/scanner belum tracked dan README memuat pekerjaan pendahulu, diff HEAD tidak dianggap seluruhnya delta logout. Identitas sebelum dari bound patch report diperiksa dengan inverse delta dalam memori: `http.go` kembali tepat ke `bd6a7706…b5009` setelah hanya menghapus ` || payload == nil`; `harness_test.go` kembali tepat ke `bdbfe0b5…28552` setelah hanya mengembalikan marker Run; `auth_harness.py` kembali tepat ke `01a2e728…6d7a6f6` setelah mengembalikan Run dan selector `^TestHarness`; `scanner_gate.py` kembali tepat ke `6c0f06e3…c615dc` setelah hanya mengembalikan hash `http.go`. SHA-256 penuh sebelum dimiliki patch report terikat. Tidak ada source temporer ditulis atau dieksekusi.

README sekarang dibaca penuh dan identitas sesudah cocok. Preimage README `38b5623e84b9b2fef704978b8f3100fd68e51d2b35f14fe939e02893ba85d99f` tersedia sebagai identitas Builder-reported; byte preimage tidak direkonstruksi secara independen. Penilaian dokumentasi mengacu teks live, batas scope, dan kedua report terikat, tanpa mengklaim pembuktian byte-diff README sebelum/sesudah.

## 1. Safety

Tidak ada blocking finding.

- `backend/internal/platform/auth/http.go:250–280`: `noStore` dan `CheckMutation` tetap mendahului decoding. Exact Origin, media type JSON, sesi terautentikasi, serta perbandingan CSRF tetap diperlukan. Perubahan tunggal membedakan nil map hasil JSON `null` dari map kosong hasil `{}`. Decode error, nil, atau field tambahan menghasilkan safe 400 `invalid_request` dan immediate return, sebelum `Store.RevokeSession` serta sebelum penerbitan cookie. Defect awal adalah pelanggaran kontrak input pada request authenticated/CSRF-valid; tidak ada bukti bypass CSRF/privilege dari defect itu.
- Batas 16 byte, penutupan body dan second decode yang harus `io.EOF` tetap. Non-object, JSON rusak, field tambahan, trailing value, dan body terlalu besar tidak memperoleh jalur revokasi. Empty object (termasuk whitespace JSON yang masih dalam batas) tetap lolos; bukan syarat literal byte persis `{}`.
- Revocation masih memakai request context dan mekanisme transaksi/lock yang sama. Error 401/503 tetap aman; cookie penghapusan hanya dibuat sesudah revokasi berhasil. Constructor/caller cookie di tiga lokasi G124 tidak berubah, termasuk production Secure/HttpOnly/Lax/host-only dan development exception.
- Ownership marker test dan label runner sama-sama menunjuk Build logout yang baru. Nama resource acak, loopback port, environment credential privat, setup/grants, dan cleanup ownership checks tidak berubah. Selector tidak mengeksekusi provider/concurrency suite atau membuang assertion dalam dua test terpilih.
- Tidak ada concern concurrency/security baru yang absen dari Test Focus Pointer: parent R10/R11/R13 dan §12 sudah mencakup strict bodies, CSRF, cookie, serta session revocation/lock. Tidak ada Techplan drift yang ditemukan pada delta ini.

## 2. Quality

**R1 — nonblocking: oracle null belum memeriksa header cookie secara eksplisit.**

Lokasi: `backend/internal/platform/auth/harness_test.go:597`, `TestHarnessLogoutRejectsNull`, dan helper `safeResponse` sekitar baris 186.

Problem: test null memeriksa 400, safe envelope/no-store/CORS dan sesi tetap valid, tetapi helper tidak memeriksa `Set-Cookie`; karena itu hasil test tersebut sendiri tidak membuktikan assertion negatif “tidak ada deletion cookie”. Dampak: perubahan masa depan yang menambahkan cookie ke jalur error dapat lolos oracle ini meskipun status dan DB tetap benar. Saran: pada Build yang berwenang berikutnya, tambahkan assertion negatif header `Set-Cookie` untuk null. Ini enhancement nonblocking: source live dan delta eksak sudah menunjukkan immediate return sebelum satu-satunya jalur penerbitan cookie logout. Tidak ada source change yang diwajibkan untuk menerima patch ini, sehingga tidak dibuat patch plan atau patch loop.

Oracle utama tetap kuat dan tidak dilemahkan: inverse hash membuktikan test hanya berganti marker. Null menggunakan sesi nyata dari restricted Store dengan Origin/CSRF valid dan `ReadSession` yang menyaring `revoked_at IS NULL`; salah-revokasi membuat assertion gagal. Test positif `{}` tetap memeriksa 204, tepat satu cookie penghapusan bernama environment development, value kosong, MaxAge=-1, Path=/, sesi tidak lagi authenticated, serta repeated logout 401. Invalid-body variants tetap memeriksa 400 dan sesi valid. Tidak ada klaim bahwa assertion ini exhaust seluruh T2 atau production browser cookie behavior.

Runner tetap memakai regex berjangkar untuk tepat `TestHarnessHTTPMeMutationLogout` dan `TestHarnessLogoutRejectsNull`, `-count=1`, `-failfast`, authharness tag, timeout Go 90s serta subprocess 120s, dan mengembalikan exit test. Seluruh subtest dalam test terpilih tetap aktif. Tidak ada kompleksitas, duplikasi atau logging baru pada delta.

## 3. Stack-Specific Best Practices

Tidak ada finding tambahan. Routing melalui Go subindex dan Security Concern Map; guidance yang dibaca dan diterapkan:

- `../harscode-workspace/best-practices/go/nil-and-zero-values.md`: nil/empty merupakan keadaan berbeda pada boundary; explicit nil check menyelesaikan semantik object-only tanpa dereference hazard.
- `../harscode-workspace/best-practices/go/input-validation-and-injection.md`: validasi dilakukan di HTTP boundary sebelum mutation; runner menggunakan argv subprocess, bukan shell interpolation. Credential/SQL setup tidak berubah oleh patch.
- `../harscode-workspace/best-practices/restapi/csrf-and-cookie-security.md`: session-bound header plus exact Origin tetap mendampingi SameSite dan production HttpOnly/Secure; development exception tetap dibatasi konfigurasi repository.
- `../harscode-workspace/best-practices/go/jwt-and-token-lifecycle.md`: konteks auth ditinjau; aturan JWT key separation/refresh rotation N/A pada opaque local session tanpa refresh dalam scope ini. Tidak diperkenalkan JWT/key handling.
- `../harscode-workspace/best-practices/go/integration-testing-setup.md`: authharness tag, fixture terisolasi, cleanup ownership, bounded execution, serta batas bukti runtime tetap eksplisit. Forced interruption/timeout cleanup tidak diuji dalam Review; tidak disimpulkan dari pembacaan runner.

## 4. Consistency

Tidak ada finding tambahan terhadap root/backend `AGENTS.md`, Approved Techplan R10/R13 dan §8/§12, T2 snapshot, bound patch scope, atau OpenAPI `api/openapi/index.yaml:74`.

Object-only schema dengan `additionalProperties: false` dan 400 InvalidRequest sesuai dengan nil rejection; safe response memakai helper lokal yang sama tanpa SQL/stack/path/secret leak. Tidak ada API, Product, dependencies, migration, authorization, lock, atau session persistence change.

`scanner_gate.py` hanya mengganti satu `SOURCE` hash `http.go`. Inverse hash membuktikan seluruh sisa file identik dengan before identity: SOURCE set, HISTORY, tool/version guards, coverage/error handling, `nosec` handling dan accepted candidate set tetap. Tiga candidate tetap G124 `(49,21)`, `(158,23)`, `(277,23)` untuk file yang sama; disappearing/extra/drifted candidates masih fail closed. Ini source identity refresh setelah Builder-reported re-triage, bukan allowance expansion atau scanner clearance independen.

README menjelaskan binding Run baru, hanya dua regression terpilih, simulated Protocol/httptest, dan larangan inferensi real Google/browser/HTTPS/deployed credentials/independent T2 acceptance. Cleanup description adalah mekanisme runner, bukan attestation Review. Patch report melaporkan ownership-checked container/anonymous-volume cleanup, dan menyatakan independent post-run Podman listing tidak tersedia. Review mempertahankan limitation itu. Tidak ada klaim cleanup unrelated resources atau fault-injection cleanup. AUTH-HARNESS-001 tetap terminal FAILED; passing groups sebelumnya dan hasil patch tetap bukti milik Builder.

## Verification executed during Review

**none** — tidak ada runtime verification, test, scanner, fixture, DB, Podman, provider, atau browser execution.

Pemeriksaan read-only yang dilakukan: membaca authority/contract/report/live source; `git status`, branch/HEAD/workflow revision dan scoped `git diff`; SHA-256 manifest serta 12 input; inverse-delta SHA-256 empat file di memori; penalaran empat pass. Tidak ada source diedit. Output Review ini merupakan independent reasoning, bukan independent runtime/T2 acceptance. PASS harness dan effective Gosec PASS dengan tiga raw findings hanya Builder-reported dari bound patch report; tidak dijalankan ulang atau diubah menjadi bukti Review.

## Verdict

**Approve with minor comments** — 0 blocking findings, 1 komentar nonblocking (R1). Penerimaan terbatas pada patch lima file: null ditolak dengan safe 400 sebelum revokasi, `{}` dan boundary/error/cookie behavior tetap, ownership fixture diperbarui, serta satu sourceguard G124 identity diperbarui tanpa perluasan kebijakan.

Tidak menetapkan whole-T2/T3 readiness, release acceptance, dependency commit, real Google login, configured deployment runtime, atau closure `B-T2-002`/`B-T2-003`/historical `B-T2-001`. Status koordinasi Finding/Blocker tetap milik Orchestrator.

## Phase handoff

- Outcome: COMPLETED — empat canonical Review pass selesai; Approve with minor comments untuk exact patch claim.
- Result refs: `RUN-C1-ENG-CODEREVIEW-T2-LOGOUT-PATCH-001/evidence/review-findings-1.md`; tanpa patch plan.
- Findings: R1 nonblocking, oracle null belum mengassert ketiadaan `Set-Cookie`; 0 blocking findings.
- Decision requests: none.
- Blockers: none untuk claim Review ini; tidak menutup blocker T2 yang ada.
- Open / unverified: runtime/test/scanner/fixture cleanup tidak diverifikasi independen; preimage byte README belum direkonstruksi; B-T2-002/B-T2-003/historical B-T2-001 dan provider/browser/HTTPS/integrated/T2 acceptance tetap di luar Run.
- Recommended continuation: Orchestrator rekonsiliasi verdict dan evidence identity terhadap claim ini, lalu tentukan applicability langkah berikutnya. R1 tidak memerlukan patch loop. Additional Testing untuk closure mismatch ini telah dinyatakan N/A pada coordination; Review ini tidak mendispatch Testing/T3. Jika ada phase/re-entry berikutnya, gunakan fresh Run/Participant/Session.
- Context refs: invocation dan input-bindings Run ini; bound patch report; Approved techplan/T2; API logout; `coordination/2026-10-12-auth-logout-patch-reconciliation.md`.
