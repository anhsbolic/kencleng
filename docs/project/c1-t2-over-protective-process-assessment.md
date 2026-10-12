# C1/T2 — Evaluasi proses over-protective

- **Tanggal:** 2026-10-12
- **Scope:** proses engineering/verifikasi C1/T2, bukan perubahan Product/auth contract.
- **Status:** evaluasi dan rekomendasi proses yang diminta Human; bukan approval scanner disposition, policy exemption atau acceptance.
- **Owner tindak lanjut:** Orchestrator untuk routing; security/Human/Testing sesuai concern.

## Penilaian utama

Proteksi menjadi berlebihan ketika tambahan audit, Run, refactor atau rerun tidak menjawab pertanyaan material baru. Untuk C1/T2, perbaikan yang disarankan terutama mengurangi pengulangan proses. Bukti security backend yang belum established tetap diperlukan.

Bedakan tindakan yang sudah terlihat berlebihan, risiko proses jika diulang, dan proteksi yang masih justified. Tidak semua Review atau scanner failure adalah overhead yang boleh dihapus.

## Evaluasi area konkret

| Area | Fakta dan penilaian | Risiko proteksi berlebihan | Rekomendasi dan batas |
|---|---|---|---|
| **Empat kandidat Gitleaks history** | Independent Review memetakan tiga fixture inputs dan satu komentar JWT; satu diagnostic v8.30.1 mereproduksi empat fingerprint. Attribution gap sudah terjawab. Mengulang audit yang sama tanpa changed premise akan berlebihan. | Run berulang untuk kesimpulan sama, keputusan owner tertunda, worktree gate tetap belum selesai. | Minta keputusan atas empat exact fingerprints, implementasikan treatment yang disahkan secara terbatas, lalu verifikasi actual gate. Jangan blanket rule/file/test-path exclusion. Proposal false-positive bukan acceptance yang sudah enacted. |
| **Tiga G124** | Source/unit menunjukkan flags sesuai valid production/dev constructor/composition; arbitrary mutable config dan actual logout/browser belum terbukti. Refactor auth besar belum dilakukan dan tidak wajib hanya karena scanner flags. | Perubahan security core yang tidak diperlukan menambah regression risk dan scope; mengejar clean scanner menggantikan tujuan perilaku aman. | Prioritaskan disposition sempit bila owner menerima rationale. Jika batas tidak dapat dijaga atau defect terbukti, lakukan scoped remediation. Runtime evidence tetap wajib; disposition tidak membuktikan semua auth entry paths aman. |
| **Review berulang setelah narrow patch** | Review terakhir justified: menyelesaikan candidate attribution yang konkret dan menilai batas cookie oracle. Bukan contoh Review tanpa hasil. Mengulang four-pass Review otomatis pada tiap metadata/literal/policy iteration bisa berlebihan. | Phase bounce, context/approval overhead dan reviewer mengulang fakta tanpa material delta. | Tentukan applicability dari actual delta, risk dan unresolved question; kembali ke requesting phase setelah patch bila authority mengizinkan. Independent Review tetap dipakai untuk keputusan/risk nyata; scope/perubahan policy yang material jangan disamakan dengan cosmetic change. |
| **Dependency/provenance gate diulang dari nol** | Build retry-1 punya bounded dependency evidence; perubahan dua error literals tidak mengubah module pins. Tidak ada alasan mengulang seluruh provenance acquisition hanya karena Run ID berganti. | Network/cache recovery berulang, waktu habis pada tooling, bukti substantive tidak bertambah. | Periksa relevansi source/toolchain/module/config evidence, referensikan yang tetap valid, jalankan affected checks dan final/fresh gates yang memang diwajibkan owning authority. Source dependency, runtime atau vulnerability database perubahan bisa membutuhkan pemeriksaan baru. Cached source saja bukan integrity/provenance clearance. |
| **Broad suites pada setiap phase** | PATCH-002 memakai focused Staticcheck/cookie units; Review memakai satu disputed-history diagnostic. Keduanya proporsional. Mengulang make verify/race/browser/security penuh di Build dan Review demi confidence tambahan bisa berlebihan. | Final suite menjadi ritual yang menduplikasi ownership Testing; biaya naik tanpa oracle baru. | Affected verification dahulu, broaden dengan pertanyaan/risk yang spesifik. Final backend gate tetap dijalankan oleh owning Testing pada target yang relevan; constituent PASS tidak menggantikannya. |
| **Production HTTPS dipaksakan sebelum environment tersedia** | HTTPS tidak bergantung FE, tetapi real production-cookie evidence membutuhkan environment HTTPS. Belum ada keputusan hold khusus HTTPS dari Human. | Setup yang di luar scope, dispatch outcome yang sudah diketahui tidak reachable, atau salah menyebut environment gap sebagai defect. | Resolve prerequisite atau narrow scope dan catat production HTTPS belum verified sesuai active authority. Bukan automatic waiver/hold tambahan, bukan localhost PASS yang dilabel production. |

## Overhead preparation yang sudah dikoreksi

PATCH-001 mengikat 62 inputs, termasuk lifecycle projections dan banyak guidance. Perubahan approval/scheduling bisa membuat assignment tampak stale tanpa semantic scope drift. PATCH-001 kemudian dipertahankan sebagai prepared-but-never-dispatched/superseded; PATCH-002 menggunakan 18 assignment bindings. Ini contoh koreksi terhadap terlalu luasnya exact-binding, bukan izin mengabaikan perubahan material plan/source.

Preservation diagnostic di PATCH-001 memuat 38 identities. Bukti historis tetap dipertahankan; current evidence-sufficiency guidance tidak mewajibkan setiap output/cache disalin ulang pada Run berikutnya. Dua safe captures pada Review terakhir justru justified karena attribution diperselisihkan. Jumlah file bukan satu-satunya ukuran overhead.

## Proteksi yang tetap sebanding dengan risiko

| Pemeriksaan | Mengapa masih diperlukan |
|---|---|
| HTTP contract/Origin/CSRF/logout assertions | Memberi bukti boundary auth, bukan sekadar command tagged contract yang tidak memiliki assertions. |
| Restricted-runtime DB/SCRAM/privileges | Serial fixture/provisioning-owner connection tidak membuktikan identity aplikasi memiliki batas akses yang benar. |
| Concurrent first login/session locking/expiry/revocation | Unit fake-row dan Go race detector tidak membuktikan invariant lock/transaction PostgreSQL. |
| Final backend gate setelah relevant changes | Focused lint/unit success tidak menetapkan security/final verification keseluruhan. |
| Exact protected implementation approval dan secret hygiene | Mengurangi risiko perubahan trust/elevated access dan credential disclosure; biaya proses bukan alasan mengabaikan boundary ini. |
| Independent/real/Human-attested evidence separation | Menghindari acceptance palsu dari fixture, laporan Human atau simulated provider. |

## Rekomendasi urutan tindak lanjut

1. Selesaikan decision-ready scanner disposition dengan owner; jangan audit ulang attribution tanpa uncertainty baru.
2. Jika disahkan, fresh bounded Build membuat treatment minimal yang dapat direview, atau remediation source bila itu opsi owner. Tidak ada authorization policy implementation dari dokumen ini.
3. Rekonsiliasi actual delta dan tentukan Review applicability dengan alasan spesifik; jangan automatic full-loop replay.
4. Siapkan independent backend Testing untuk executable scope, termasuk substantive HTTP/DB/final gates yang masih perlu bukti.
5. Bawa [hold Google login](c1-google-login-verification-hold.md) dan [hold integrated verification](c1-integrated-verification-hold.md) sebagai deferred, lalu lakukan final integrated evidence saat prerequisites tersedia.

B-T2-002/B-T2-003 tetap open; D-PATCH002-01 belum accepted. Tidak ada T2/T3 readiness/acceptance, dependency commit atau Run baru dari dokumen ini. Approved Techplan/task/manifest tetap; bila perampingan membutuhkan perubahan material authority, route ke owner.

## Dasar evidence dan guidance

- [Independent patch/security Review](../../.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-CODEREVIEW-T2-PATCH-001/evidence/review-findings-1.md), SHA-256 `7ea1a91074bdb5e064c167ca26003f3e99b5d79ce201c40bda2aa32bb1e1fe74`.
- [PATCH-002 report](../../.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-PATCH-002/evidence/patch-report-1.md) dan [checkpoint reconciliation](../../.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/coordination/2026-10-12-patch-checkpoint-reconciliation.md).
- [Candidate-specific Human proposal](../../.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/coordination/2026-10-12-scanner-candidate-decision.md).
- Harscode current workflow/3-build/guidelines.md, workflow/4-code-review/guidelines.md, orchestration/run-contract.md (binding/evidence sufficiency) dan orchestration/AGENTS.md (readiness/routing). Observed workflow revision: `f6856df25a13b004e5fb242eb71f3ecd80728f70`. Referensi ini guidance sumber; evaluasi manfaat tambahan pada tabel adalah judgment Orchestrator, bukan rule Harscode baru.
- [Panduan umum proportionality](engineering-protection-proportionality.md).

Tidak ada benchmark waktu/biaya yang diukur untuk evaluasi ini. Penilaian berfokus pada material information delta dan outcome, bukan angka efisiensi yang belum diverifikasi.
