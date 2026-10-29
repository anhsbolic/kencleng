# ORG-01 / rancang-002 — Engineering Solution Handoff

**Produksi dokumentasi COMPLETE — bundle `ORG-01-rancang-002-r1`; Human acceptance PENDING.** Pemilik assembly dan standing: Koordinator Rancang. Exact final file revisions dipin pada [output-revisions.json](output-revisions.json); artifact tetap proposal untuk pertimbangan Human, dengan external held edges di bawah.

## Ringkasan keputusan

Outcome approved adalah pembentukan organisasi bersama hubungan `Owner` internal pertama yang teratribusi, dengan pemahaman sebelum efek, pemeriksaan kembali setelah berhasil, dan penyampaian hasil negatif/unknown secara jujur. Ini tidak memberi status legal verification, approval Reviewer, publikasi atau kelayakan donasi.

Pilihan utama berada pada [solution.md](solution.md), bagian **O1 — Keputusan bersama dan batas** dan seluruh bagian O2; kontrak pengalaman berada pada [experience.md](experience.md):

1. Google login dilakukan eksplisit sebelum field nama, dilanjutkan review konteks/akun dan submit eksplisit. Experience dan Architect memilih ini setelah membandingkan popup/draft preservation: biaya login lebih dini diterima untuk mengurangi state lintas-window dan penyimpanan nama sebelum login. Login sendiri tidak membentuk organisasi.
2. Satu transaksi PostgreSQL membentuk organisasi, relasi initial Owner melalui `first_owner_account_id`, dan account-bound creation receipt. Model ini cukup untuk initial Owner scoped; perpindahan/multi-Owner/full RBAC memerlukan desain baru saat diotorisasi. Nama tidak unique dan tidak menjadi bukti identitas/legal.
3. Same-key replay dan lookup menjaga hasil unknown agar tidak menjadi sukses palsu atau otomatis menciptakan konteks baru. UUIDv7, bounded new-effect window 24 jam dan penolakan expired absent key mencegah receipt cleanup membuka ulang efek lama. Ini pilihan teknis retry; bukan kebijakan retensi bisnis.
4. Bounded own-org discovery/detail memberi pemeriksaan keadaan terkini setelah toast atau browser marker hilang. Receipt membuktikan event pembentukan, sementara current reads memeriksa hubungan sekarang. Nama serupa tidak membuktikan bahwa suatu unknown attempt committed.
5. Identity/session prerequisite mengikuti T02, dengan trusted internal account, JWT HttpOnly, rotating refresh dan backend effective authorization. Data diperkecil menjadi kebutuhan atribusi/konteks; tidak menyimpan email/profil Google. Parallel refresh dapat berakhir pada pencabutan session dan login ulang sebagai trade-off fail-safe; seamless recovery tidak dijanjikan.

Source-level review dan planned verification belum membuktikan aplikasi bekerja. **PD-PRIV-01 masih keputusan Product/privacy owner** untuk retensi/penghapusan/basis pemrosesan data nyata dan backup. Rincian serta alternatif ada pada solution, bagian **PD-PRIV-01 — precise owner edge, not an invented lifecycle rule**. Dampaknya menahan real-person persistence release serta lifecycle/cleanup readiness; tidak menambah kebijakan privasi melalui output Engineering. Tidak ada perubahan standar yang diusulkan.

## Identitas, sumber dan standing

- Assignment/otorisasi/runtime: [state](state.md).
- Kencleng baseline `d0d78fab26a976f9f48c6fc97a45c0adc5cec282`, branch `pilot/station5-c1-space-only`.
- [Product Contract](../../../docs/product/outcomes/org-01/product-contract.md), blob `da19d8e7cdb9baac8cf4af1c5e6b61763607013e`; [approval](../../../docs/product/outcomes/org-01/approval.md).
- [Engineering Standards](../../../docs/engineering/standards/README.md), [Design authority](../../../docs/design/README.md), [exact source revisions](source-revisions.json).
- Rancang formed definition `47297c6f369dad4d4e4c4f90f0e81eb0310be7cf` pada repo Harscode; formation approval terpisah dari invocation, runtime verification dan acceptance bundle ini.

## Consumer entrypoints

| Consumer diketahui | Owning sections yang harus dibaca saat tersedia | Project-bound dependency dan next permitted action |
| --- | --- | --- |
| Pemilik implementasi backend | [X1](requirements.md), [solution](solution.md) bagian O1, O2 identity/permission/data/transaction/API dan O4 TECH; [experience](experience.md) tabel intent dan state; [Security](security-findings.md) dan [Solution Review](solution-review.md) | Root/backend AGENTS, T01–T08 dan baseline source telah diperiksa sebagai design constraints. Fit **CONDITIONAL**: canonical OpenAPI/SQL belum diadopsi, protected auth/session/authz/key implementation belum diotorisasi, PD-PRIV-01 dan deployment prerequisites tetap terbuka. Next permitted action: inspeksi/assessment proposal ini; tidak ada izin menulis implementasi dari handoff. |
| Pemilik implementasi frontend | [X1](requirements.md), [experience](experience.md) bagian kontrak intent/state, auth/reload/accessibility dan EXP01–08; [solution](solution.md) API/result certainty, current reads dan TECH02/04/05/06/07; Design originals dan independent findings | Frontend AGENTS, state ownership, Design dan T02/T04/T06/T07 diperiksa sebagai design constraints. Fit **CONDITIONAL**: typed active OpenAPI belum ada, cookie integration dan rendered/accessibility/FE↔BE evidence NOT RUN. Next permitted action: inspeksi kebutuhan komponen/alur terhadap proposal; canonical contract/implementation memerlukan acceptance dan otorisasi terpisah. |
| Pemilik verifikasi | [X1 coverage](requirements.md), [TECH01–09](solution.md), [EXP01–08](experience.md), seluruh independent findings | T07 dan sumber consumer diperiksa; 17 O4 punya producer tunggal dan intended downstream verifier. Fit **CONDITIONAL** terhadap actual implemented contract/environment yang belum tersedia. Next permitted action: menilai kecukupan kewajiban yang sudah ditulis; aplikasi/DB/browser/services tidak dijalankan atau diotorisasi oleh assignment ini. |

**Generic compatibility:** NOT APPLICABLE untuk tiga pemilik langsung di atas karena tidak ada declared reusable Space Input Contract mereka yang diidentifikasi. **Project-bound compatibility:** selalu diperiksa, berdasar artifact/revisi asli di atas; design constraints checked, consumer adoption/implementation fit CONDITIONAL sebagaimana tabel. Pemeriksaan ini tidak merupakan executed receiver acceptance. Future/unknown consumer fit NOT CHECKED. Reviewer navigation discovery adalah observer evidence; real downstream F-01/cold-resume probe tetap NOT RUN.

## Index producer dan review

| Payload / owner | Revisi dan standing |
| --- | --- |
| X1 / Requirements Analyst `/root/requirements` | [requirements.md](requirements.md); SHA-256 `055eb247a004a724b130c79e6a07ce662e20f520bb9f44df11802037af8ff951`; 4 SC / 8 REQ / 9 AC terpetakan |
| O1/O2 dan O4 TECH / Architect `/root/architect` | [solution.md](solution.md); reviewed SHA-256 `058c4f00912ea41f882982109aa20b153a6ee983237655ccb3fc9bc04ee7945c` |
| O3 dan O4 EXP / Experience `/root/experience` | [experience.md](experience.md); SHA-256 `475683ebe5e94b83e0930ccdbc5c4f92a4c1368da17229c802fe2b6e2fad1604`; current-state/retry/auth sudah direkonsiliasi |
| Security Findings / `/root/security` | [security-findings.md](security-findings.md), SHA-256 `0b3ec3ecf0be352c54a54d58c1a1254307e7be3f4a0d51360646652b7b10fbce`; independent bounded design assessment selesai |
| Independent Solution Review / `/root/reviewer` | [solution-review.md](solution-review.md), verdict **ELIGIBLE** untuk bounded design completion/Human consideration; exact report revision mengikuti output manifest |

Koordinator mengonfirmasi Architect sebagai producer seluruh TECH01–09 dan Experience sebagai producer EXP01–08. Intended verifier pada O4 bukan dispatch/otorisasi tes. Temuan material harus dikoreksi producer dan dinilai ulang oleh pemilik review; Koordinator tidak menulis ulang findings.

Disposition Security: **SEC01 ADDRESSED AT DESIGN LEVEL** setelah Architect memperjelas CSRF nonce/HMAC derivation dan Security menilai ulang actual bytes; runtime closure tetap NOT VERIFIED. **SEC02 OPEN / AFFECTED EDGE HELD**: PD-PRIV-01 tetap keputusan owner untuk real-person persistence/lifecycle release, bukan waiver. Security menyatakan tidak ada material safeguard defect yang menghalangi proposal dokumentasi scoped dengan dependency tersebut tetap eksplisit. Original findings tetap dimiliki Security dan tidak digantikan ringkasan ini.

Independent Solution Reviewer memeriksa source fidelity dua arah, trade-off Experience/Engineering, 17 O4, actual Security findings/disposition dan consumer constraints, lalu menilai **tidak ada material producer blocker** untuk scoped documentation completion. PD-PRIV-01 dinilai external release/lifecycle gate yang telah diidentifikasi cukup konkret, sehingga tetap OPEN dan HELD tanpa menambah authority. Final administrative status/reference/manifest changes diizinkan report; perubahan mekanisme, source authority, gate atau consumer claim membuka ulang affected review.

## Evidence dan batas runtime

Executed: source inspection oleh producer; X1 coverage check; role package diff terhadap approved definition; source SHA checks (Analyst 21/21, Architect 20/20, Experience 14/14; Koordinator 31 baseline sources, hanya authorized README navigation berubah). Security/Reviewer mempunyai fresh contexts tanpa inherited producer conversation dan harus membuka original serta final producer bytes sendiri. Exact independent findings/checks dilaporkan pemiliknya, bukan diasumsikan dari label.

R2 scoped deviation diotorisasi: tool write surface lebih luas daripada assignment, sehingga narrow writes bergantung pada instruksi/supervisi. R5/R6 cold-resume/process-loss/backup dan comprehensive native gate enforcement belum dibuktikan. Specialist model `gpt-6.1-sol / medium`; Security pernah menerima capacity error, kemudian dilanjutkan pada konteks/model sama tanpa fallback. Overall Harness standing tetap PARTIAL.

NOT RUN: application/unit/coverage/DB/session/security/browser/accessibility/integration tests, actual downstream acceptance/discovery probe, production recovery. Source/link/hash checks tidak dianggap runtime conformance. Tidak ada canonical API/code/schema/standards/env/service/deploy/commit/push perubahan dalam assignment.

Hasil pemeriksaan dokumentasi final dicatat pada [documentation-checks.json](documentation-checks.json). Output manifest mengikat final bytes seluruh payload, state, handoff, findings, checks dan navigation; manifest tidak meng-hash dirinya sendiri. Intermediate checkpoint yang disebut producer/review adalah riwayat pembacaan yang terpisah dari final artifact hashes, bukan circular final-hash dependency. Tidak ada klaim semua intermediate bytes telah diarsipkan Git; cold-resume/process-loss/backup tetap NOT RUN.

## Gate dan keputusan Human

1. Exact-bundle acceptance PENDING. Final revision manifest dan review disposition harus dibaca sebelum menerima bundle.
2. PD-PRIV-01 PENDING pada owning Product/privacy source; tidak ada periode retensi atau erasure exception yang diputuskan Engineering. Resolusi memerlukan scope owning-source persistence terpisah dan revalidation affected output/review.
3. Canonical API adoption, migration dan protected implementation authorization PENDING; penerimaan bundle tidak membuka izin tersebut.
4. Real-user/public deployment memerlukan privacy, security, ingress/key custody dan actual verification/recovery evidence tersendiri.

## Batas penggunaan

Rancang COMPLETE (bila tercapai), Human acceptance pada exact bundle, canonical contract adoption, implementation authorization dan actual runtime evidence adalah status terpisah. Pilot runtime PARTIAL dengan scoped R2 deviation; application/security/browser/integration verification NOT RUN. Review independence akan dilaporkan berdasarkan konteks dan source inspection aktual, bukan role label.
