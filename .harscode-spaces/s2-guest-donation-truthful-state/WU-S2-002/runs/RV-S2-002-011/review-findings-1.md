# Code Review Findings — `RV-S2-002-011`

> Phase: Code Review  
> Work Unit / Run: `WU-S2-002` / `RV-S2-002-011`  
> Author: Codex Reviewer  
> Role / specialization: Reviewer / Independent four-pass review of Task 01 post-approval Donation spec reconciliation  
> Participant: `P-S2-002-RV-011-1`  
> Session: Fresh Reviewer Session; Session ID not exposed  
> Created: 2026-10-01  
> Model / reasoning: `gpt-6-luna` / `high` (Invocation configuration; runtime selection not independently exposed)  
> Target revision: `650e73c5d646c29c0ddf1931618f02685d15f7b7` plus current durable approval/task snapshots and working-tree artifacts  
> Workflow revision: `33b03a3f62cc3aacba6534b8a011465613c64b09`

## 1. Safety

### F-001 — Persyaratan credential status yang sulit ditebak hilang dari acceptance hasil rekonsiliasi

- **Location:** `docs/spec/5-donation/invariants.md#INV-donation-05`; `docs/spec/5-donation/features/02-donation-status-check.md` acceptance criteria and threat breakdown; `docs/spec/5-donation/tasks.md` Task 02 acceptance; `docs/spec/5-donation/threat-model.md` temporary guest status URL / Spoofing row.
- **Problem:** Acceptance hasil rekonsiliasi sudah mencatat fragment handoff, URL cleanup, one-way HMAC verifier, hard expiry 24 jam, dan response status-only, tetapi tidak lagi mensyaratkan bearer credential/token yang sulit ditebak. Baris threat model justru memasukkan “token strength” sebagai hal yang masih terbuka di O4 tanpa mempertahankan sifat sulit ditebak yang sudah diputuskan. Tabel klasifikasi tugas historis di `tasks.md` masih menyebut “difficult-to-guess”, tetapi active acceptance tidak.
- **Dampak:** Kepemilikan credential memberi akses ke status donasi privat. Hilangnya syarat anti-guessing melemahkan perlindungan yang disetujui terhadap credential guessing dan dapat membuat rekonsiliasi API/implementasi berikutnya memperlakukannya sebagai kontrol opsional yang belum diputuskan.
- **Authority:** TP-015 Q7 mensyaratkan “hard-to-guess token”; Product/MVP `docs/product/mvp-scope.md` §5 dan `docs/product/mvp-delivery-slices.md` §5 juga mensyaratkannya. `../harscode-workspace/best-practices/restapi/anti-enumeration.md` merutekan concern secret comparison dan enumeration; spec saat ini tepat membiarkan comparison implementation dan bukti empiris tetap terbuka, tetapi harus mempertahankan syarat tingkat Product bahwa credential sulit ditebak.
- **Resolusi yang disarankan:** Kembalikan “difficult-to-guess” sebagai properti credential yang sudah ditetapkan pada INV-donation-05, acceptance Task 02, dan acceptance Feature 02. Selaraskan baris Spoofing agar membedakan properti yang sudah ditetapkan itu dari bukti konkret token generation/strength dan kontrol implementasi O4 yang masih terbuka. Jangan menetapkan target entropy/panjang numerik atau mengklaim bukti implementasi.
- **Blocking:** Ya. Minta perubahan sebelum rekonsiliasi spec ini dapat diterima sebagai sesuai TP-015 dan Product/MVP.

Tidak ada temuan Safety lain. Spec yang berubah mempertahankan kewajiban lifecycle terminal notice O3 dan secara eksplisit membiarkan kontrol/bukti tetap terbuka; safeguards implementasi O4/O5, parity response/timing empiris, abuse controls, dan residual-risk acceptance tidak diklaim terbukti.

## 2. Quality

Tidak ada temuan. Wording lintas dokumen untuk O1, O3, O4/O5, dan O11/D19 cukup jelas untuk direview dan mempertahankan pembedaan antara arah yang sudah ditetapkan dan kewajiban implementasi/bukti yang masih terbuka. Kelima spec tetap berstatus `draft`.

## 3. Stack-Specific Best Practices

**Tidak ada temuan tambahan.** `../harscode-workspace/best-practices/restapi/anti-enumeration.md` sesuai dengan concern status credential dan enumeration. Panduan generic response dan constant-time comparison berlaku untuk implementasi berikutnya; dokumen saat ini membiarkan comparison controls dan parity empiris tetap terbuka di O4/O5 dan tidak mengklaim bukti yang belum ada. `../harscode-workspace/best-practices/go/decimal-and-money.md` sesuai dengan persyaratan moneter; diff mempertahankan exact-decimal/no-float sambil membiarkan precision dan rounding yang belum diputuskan tetap terbuka. Checklist comparison/rounding yang spesifik implementasi tidak menjadi alasan untuk menambah persyaratan domain-spec di sini.

## 4. Consistency

F-001 adalah satu-satunya ketidakselarasan material terhadap authority: active acceptance belum membawa syarat credential sulit ditebak dari Product/MVP dan TP-015. Tidak ada temuan consistency tambahan. Arah lain yang diperiksa selaras dengan root `AGENTS.md`, `docs/spec/README.md`, Product/MVP, `docs/ui-ux/README.md` beserta interaction guidance yang dirujuk, monetary standard, Task 01, dan TP-015. D1, wording O7, kepemilikan simulator, pengecualian Slice 2, kepemilikan exact API expression oleh Task 02, status `draft`, dan batas `CONTRACT_READY` tetap terjaga.

## Verification executed during Review

- Pemeriksaan read-only `git status --short` dan `git diff` yang dibatasi ke lima file Donation spec; memastikan scope source yang berubah dan melihat adanya working-tree artifacts lain tanpa mereviewnya.
- Membaca TP-015 current-effective, Task 01, report/launch record BLD-004, root `AGENTS.md` yang diberikan untuk sesi ini, workflow Review yang dirutekan, `docs/spec/README.md`, Product/MVP authorities, Design routing/pattern source, monetary standard, authority map, dan best-practice files yang cocok.
- Pemeriksaan `rg` terarah memastikan syarat “hard-to-guess” ada di Product/MVP dan TP-015 tetapi hilang dari acceptance hasil rekonsiliasi (meski disebut di tabel klasifikasi tugas historis).
- Tidak menjalankan tests, runtime checks, atau reproduction implementasi; hal tersebut tidak diperlukan untuk review dokumentasi ini.

## Verdict

**Request changes** — F-001 bersifat blocking. Lihat `patch-plan-1.md` untuk koreksi dengan scope terbatas. Tidak ada source yang diedit dalam Review Run ini.

## Phase handoff

- Selesai: independent four-pass review atas diff Donation spec yang terdiri dari lima file.
- Artifacts: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-011/review-findings-1.md`; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-011/patch-plan-1.md`.
- Keputusan Human: Review ini tidak meminta keputusan baru. Acceptance Human/domain-owner untuk draft Donation saat ini tetap menjadi gate terpisah.
- Terbuka / ditunda: F-001 yang blocking; seluruh keputusan dan bukti owner O1/O2/O3/O4/O5 tetap mengikuti TP-015.
- Langkah berikut yang disarankan: jalankan Build/Patch Run baru memakai `patch-plan-1.md`; setelah koreksi terbatas, teruskan ke review/acceptance yang berlaku. Tidak ada Testing dispatch atau milestone otomatis.
- Transisi Session: Build/Patch Run baru dengan Participant dan Session fresh.
- Context pointers: TP-S2-002-015 §§3–13; TPD-S2-002-002 Task 01; F-001; diff lima file.
