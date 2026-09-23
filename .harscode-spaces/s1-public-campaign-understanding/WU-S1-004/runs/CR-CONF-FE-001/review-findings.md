# Targeted Code Review Confirmation — CR-CONF-FE-001

- Phase: Code Review (targeted confirmation)
- Work Unit / Run: `WU-S1-004` / `CR-CONF-FE-001`
- Role / specialization: Reviewer / confirmation F1 dan F2 dari `CR-FE-001`
- Session: Fresh independent focused review session
- Reviewed revision: `f07e24295590f49ddf63cf68542288a4f3a82c22`
- Created: 2026-09-23

## Scope reviewed

Review ini terbatas pada patch untuk F1/F2: lifecycle route-local MSW, pengumuman dan pemulihan focus sesudah retry, serta test yang menyertainya. Tidak dilakukan four-pass re-review penuh.

## Confirmed improvements

- `mock-service-worker.tsx` sekarang menahan `start()` bila effect sudah dibersihkan, menyimpan worker yang selesai dimulai, dan menghentikannya pada unmount. Ini menutup jalur unmount biasa dan jalur import yang selesai setelah unmount.
- Success sesudah `temporarily-unavailable` kini memfokuskan landmark utama; state loading dan kedua state retryable mempunyai `role="status"`/`aria-live="polite"`, sementara 404 tetap non-live dan non-disclosing.
- Patch hanya menyentuh lima file route-local frontend (tiga implementasi dan dua test). Tidak ada perubahan pada `frontend/lib/api/`, backend, topology, OpenAPI, atau scope donation pada patch ini.

## Remaining blocking findings

### F1-R — Startup stale dapat menghentikan worker milik effect aktif pada React Strict Mode

- **Location:** `frontend/app/campaigns/[campaignId]/mock-service-worker.tsx:16-32`
- **Problem:** `worker` adalah singleton yang dibagi semua effect. Dalam lifecycle React Strict Mode/dev, effect pertama dapat memanggil `worker.start()`, lalu dibersihkan; effect kedua kemudian memulai worker yang sama. Jika startup effect kedua selesai lebih dahulu dan startup stale effect pertama selesai sesudahnya, cabang `!active` pada baris 27-29 memanggil `worker.stop()`. Panggilan itu menghentikan singleton yang saat itu dipakai effect kedua yang masih mounted, sehingga provider aktif dapat mencapai `ready` tetapi intersepsi telah dimatikan.
- **Why blocking:** Patch tidak aman/appropriate untuk lifecycle dev yang menjadi eksplisit dalam objective. Ia mencegah orphan interception, tetapi stale completion masih dapat mengganggu owner aktif.
- **Evidence gap:** Test lifecycle hanya menguji satu effect/unmount (`mock-service-worker.test.tsx:20-69`); tidak ada test Strict Mode atau dua startup tertunda dengan urutan penyelesaian silang.
- **Required resolution:** Serialisasi atau koordinasikan ownership startup/stop untuk singleton worker sehingga completion stale hanya menghentikan worker bila tidak ada owner aktif yang lebih baru; tambahkan test yang membuktikan urutan Strict Mode/overlap tersebut tidak meninggalkan provider mounted dalam keadaan worker berhenti.

### F2-R — Jalur `request-failure → retry → success` belum dilindungi sampai completion/focus

- **Location:** `frontend/app/campaigns/[campaignId]/campaign-detail-client.test.tsx:37-74`
- **Problem:** Implementasi memang memperlakukan `request-failure` dan `temporarily-unavailable` sebagai predecessor focus recovery. Namun test baru hanya membuktikan `temporarily-unavailable (503) → success` dan focus landmark. Test generic network failure hanya menekan retry lalu memeriksa focus masih berada pada tombol sebelum respons recovery selesai (baris 45-54).
- **Why blocking:** Objective meminta test fokus benar-benar mencakup intended transitions untuk kedua predecessor retryable. Tanpa respons handler yang berubah ke success dan assertion heading/focus setelah completion, cabang `request-failure` belum memiliki bukti regresi observable.
- **Required resolution:** Tambahkan test node-MSW yang membuat generic request failure, mengganti handler ke fixture success, menekan retry, lalu menunggu success dan memastikan landmark/heading success menerima focus. Pertahankan assertion bahwa initial successful load tidak memindahkan focus.

## Verification executed

- `cd frontend && npm run test -- 'app/campaigns/[campaignId]/mock-service-worker.test.tsx' 'app/campaigns/[campaignId]/campaign-detail-client.test.tsx'` — passed: 2 files, 8 tests.
- `git diff --check 1b767c94732f95649299e33e8f4478c2c781f690..HEAD -- frontend/` — passed.
- Patch-path inspection against predecessor revision — hanya lima file route-local yang disebut di atas; tidak ada branch mock production/API, perubahan backend/topology/OpenAPI, atau perluasan donation.

Vitest mengeluarkan warning konfigurasi Vite yang sudah ada mengenai native config loader; tidak ada kegagalan test.

## Verdict

**STILL_BLOCKING**

F1 dan F2 belum dapat dikonfirmasi tertutup sepenuhnya karena F1-R dan F2-R di atas. Jangan lanjut ke independent Testing sebelum patch sempit berikutnya dan targeted confirmation baru selesai.

## Handoff

- Required next step: Build/Patch terbatas untuk F1-R/F2-R, tanpa perubahan production scope di luar route-local MSW dan test terkait.
- Setelah itu: fresh targeted Code Review confirmation; baru rekomendasikan independent Testing apabila semua temuan ini tertutup.
