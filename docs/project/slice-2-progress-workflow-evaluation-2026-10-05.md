# Slice 2 — Progress dan Evaluasi Workflow Harscode

> Tanggal: 5 Oktober 2026, Asia/Jakarta. Disusun oleh Orchestration Operator dari durable workspace dan keterangan Human.
> Tujuan: evaluasi dan perbaikan Harscode berdasarkan hasil Kencleng Slice 2.
> Status: **development HOLD atas instruksi Human**; WU-S2-003 dan WU-S2-008 `ACTIVE / PARKED`.
> Dokumen ini merupakan assessment dan proposal perbaikan. Ia tidak mengubah Product/MVP, canonical Harscode, acceptance, approval, atau batas keamanan.

## 1. Penilaian utama

**Effort dan volume artefak belum sebanding dengan hasil produk yang sudah terverifikasi.** Dalam 11 tanggal kalender pengerjaan, dengan effort Human minimal 44 jam menurut keterangannya, Slice 2 menghasilkan fondasi kontrak yang diterima, frontend yang terverifikasi terhadap mock, dan bagian backend public Campaign cap projection. Perjalanan donasi dengan backend nyata belum memperoleh bukti completion.

Investment ini menghasilkan nilai engineering: authority lebih jelas, kontrak uang dan state lebih presisi, temuan kritis ditemukan sebelum implementasi, serta keputusan dan approval dapat direkonstruksi. Namun konversi hasil planning/review menjadi capability backend dan bukti integrasi masih rendah. Pada checkpoint ini, keuntungan assurance dan pembelajaran workflow lebih terlihat daripada hasil produk yang dapat digunakan end-to-end.

Penyebabnya merupakan gabungan gap produk/kontrak yang nyata, detail readiness yang ditunda sampai delivery, kesalahan fidelity Participant, dan overhead orchestration. Bukti tidak mendukung penjelasan tunggal seperti kekurangan jam kerja Human atau ketidakmampuan satu model. Report ini juga tidak membuktikan bahwa menghapus gate correctness/security akan menghasilkan delivery yang aman lebih cepat.

**Keputusan HOLD masuk akal untuk memisahkan evaluasi workflow dari kelanjutan development.** Tahap berikutnya yang disarankan adalah memperbaiki cara menentukan readiness, mengurangi pekerjaan metadata dan state berulang, serta mengukur kemajuan melalui capability/bukti. Proposal di report ini perlu dipertimbangkan oleh pemilik Harscode sebelum menjadi kebijakan.

## 2. Periode dan effort Human

| Ukuran | Hasil | Dasar dan batas |
|---|---|---|
| Mulai pengerjaan Slice 2 yang teramati | **25 September 2026** | Bootstrap dan actual launch EXP-S2-001-001; 14:13:34 UTC / 21:13:34 WIB. |
| Tanggal checkpoint/HOLD | **5 Oktober 2026** | Instruksi Human dan tanggal saat evaluasi. |
| Selisih tanggal | **10 hari** | 5 Oktober dikurangi 25 September. |
| Rentang kalender inklusif | **11 hari** | 25–30 September: 6 tanggal; 1–5 Oktober: 5 tanggal. |
| Effort harian menurut Human | **minimal 4 jam/hari** | Keterangan Human dalam sesi evaluasi, bukan timesheet. |
| Batas bawah effort Human | **minimal 44 jam** | 11 × 4 jam; termasuk tanggal awal dan tanggal hari ini sesuai pernyataan bekerja setiap hari. |

Sumber tanggal awal: [first launch record](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/launch-record.md) dan [Events](../../.harscode-spaces/s2-guest-donation-truthful-state/events.md). Keterangan effort Human diterima sebagai self-report; workspace tidak menyediakan pencatatan per jam untuk coding, membaca artefak, memutuskan authority, dispatch, atau pengembangan Harscode. Karena itu 44 jam adalah batas bawah effort yang dinyatakan, bukan pengukuran waktu coding aktif. Total biaya/token AI, durasi masing-masing Run, waktu tunggu, serta pembagian jam Kencleng versus eksperimen Harscode tidak dapat dihitung dari bukti yang tersedia.

## 3. Hasil dan posisi delivery pada HOLD

Kriteria produk Slice 2 adalah guest Donation dari Campaign eligible, QRIS simulation yang truthful, persisted pending/success/failed, revisit status yang aman selama 24 jam, perilaku optional name/email yang sesuai, exact-once successful Funding, serta admission/closure dan accepted-pending settlement yang benar. [Completion evidence Product](../product/mvp-delivery-slices.md#slice-2-completion-evidence) dan [finalisasi Kencleng](../kencleng-agentic-workflow.md#12-slice-finalization-and-delivery) mensyaratkan hasil terintegrasi beserta bukti yang berlaku.

| Bagian | Hasil yang tersedia | Batas completion |
|---|---|---|
| Authority dan current-state exploration | WU-S2-001 DONE | Hasilnya evidence/context, tanpa implementasi. |
| Donation baseline spec/API | WU-S2-002 DONE / `CONTRACT_READY` | Baseline diterima pada 1 Oktober; detail delivery/security yang deferred tetap memiliki gate sendiri. |
| Campaign donation-entry contract | WU-S2-005 DONE | Contract action diterima; backend producer dan same-predicate GET/POST evidence tetap delivery-owned. |
| Monetary/capacity sources | WU-S2-006 DONE | Tujuh exact source revisions diterima; generated/internal counterpart reconciliation dan handoff tersedia. Acceptance tidak mencakup retry credential OI9 yang muncul kemudian. |
| Funding-unavailable API | WU-S2-007 DONE | Lima exact source/counterpart revisions diterima untuk generic 503; tidak menerima aturan baru retry credential. |
| Guest Donation frontend | WU-S2-004 DONE / `FRONTEND_MOCK_VERIFIED` | Implementasi UI, scoped Review/Testing, dan Human rendered acceptance tercatat. Real backend integration belum terbukti; R3/R5/R8 assertion follow-ups tetap terlihat. |
| Backend public cap projection | Bounded Build/Review/Testing tersedia | Eight-file cap-projection slice, exact wire/amount mapping dan unapplied migration 000012. Ini belum memenuhi whole backend target. |
| Backend Donation delivery | WU-S2-003 ACTIVE / PARKED | Submission/status, simulator, D1 settlement/reservation/funding, email lifecycle dan producer/integration masih belum memperoleh whole-spine implementation/verification. |
| Retry credential source reconciliation | WU-S2-008 ACTIVE / PARKED | Invocation Exploration disiapkan tetapi belum di-dispatch. Exact source/counterpart acceptance belum tersedia. |

Milestone yang earned adalah **`CONTRACT_READY`** untuk baseline dan **`FRONTEND_MOCK_VERIFIED`** untuk scope WU004. **`BACKEND_VERIFIED`, `INTEGRATED_VERIFIED`, `SLICE_FINALIZED`, dan `DELIVERED` belum earned.** Enam WU DONE dari delapan tidak menyatakan Slice 2 selesai 75%; ukuran dan kontribusi tiap WU sangat berbeda. [Work Graph](../../.harscode-spaces/s2-guest-donation-truthful-state/work-graph.md), [WU003](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/manifest.md), dan [WU004](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/manifest.md) menjadi anchor state.

Estimasi percakapan sebelumnya, 45% selesai / 55% tersisa, adalah penilaian kasar tanpa bobot resmi. Report ini tidak mempromosikannya menjadi metrik. Data yang lebih dapat dipertanggungjawabkan adalah milestone dan capability yang telah memperoleh bukti, seperti tabel di atas; belum ada dasar untuk menghitung persentase sisa effort atau ETA.

### Timeline ringkas

| Periode | Perubahan material | Implikasi terhadap hasil |
|---|---|---|
| 25–26 September | Bootstrap, Exploration, initial Techplan, early Review terhadap atomic funding dan request idempotency; masalah visibility/report ownership juga dikoreksi. | Task berubah menjadi pekerjaan rekonsiliasi sebelum implementasi. Sebagian effort menguji orchestration pilot. |
| 27–30 September | Product amendment, atribusi owner, email/status/privacy direction, monetary standard dan D1 ordering; material plan revisions dan source authoring awal. | Kebutuhan yang semula belum cukup lengkap menjadi lebih jelas; beberapa keputusan datang setelah approval/review terdahulu. |
| 1 Oktober | Donation `CONTRACT_READY`, FE/BE topology, Campaign action reconciliation; backend planning mengekspos monetary/capacity source gap dan WU006 diturunkan. | Baseline memungkinkan delivery planning, tetapi dependencies tambahan langsung membutuhkan source reconciliation dan plan refresh. |
| 2 Oktober | Monetary/capacity Product/spec/API acceptance dan internal counterpart work; checkpoint report sebelumnya dibuat. | Source readiness maju; backend/frontend runtime belum tersedia. |
| 3 Oktober | WU006 handoff selesai; frontend mock outcome selesai; backend cap slice lalu dua Build berakhir STALLED; minimum schema Review menemukan enam blocking findings. | Frontend menunjukkan konversi ke delivery. Backend critical path kembali ke design/authority reconciliation. |
| 4 Oktober | Funding-unavailable dan Donation amount decisions, WU007 503 reconciliation/acceptance, candidate refresh; Review menemukan credential-authority mismatch lalu retry-token design gap. | Kontrol dan response recovery belum konvergen saat schema-design approval dicoba. |
| 5 Oktober | TP12/TP13, RV9/RV10; OI9 tetap unaccepted, WU008 prepared; Human memutuskan HOLD. | Kandidat Review konvergen secara material, tetapi source acceptance, exact candidate approval dan schema-design gate tetap terpisah. |

Timeline berasal dari heading dan material coordination facts di [Events](../../.harscode-spaces/s2-guest-donation-truthful-state/events.md), dibaca bersama Run reports/manifest yang relevan. Timeline tidak mengukur durasi aktif masing-masing aktivitas.

## 4. Volume effort yang terlihat dari artefak

Snapshot inventaris mencakup **126 direktori Run, 401 file total, dan 396 file Markdown** di Slice-2 Space. Isi Markdown sekitar **507.312 kata berdasarkan pemisahan whitespace**, dengan total semua file **4.403.866 byte**. Angka mencakup Invocation, riwayat plan, source snapshots, task files dan bukti; ia tidak menyatakan unique knowledge atau seluruh isi pernah dibaca Human. Metode, batas, dan daftar lengkap Run tersedia di [lampiran evidence](slice-2-harscode-evaluation-evidence-2026-10-05.md).

| Prefix aktivitas | Direktori Run | Persentase dari direktori Run |
|---|---:|---:|
| Techplan (TP) | 49 | 38,9% |
| Review (RV) | 37 | 29,4% |
| Build/Patch (BLD) | 20 | 15,9% |
| Exploration (EXP) | 7 | 5,6% |
| Open-item facilitation (OIR) | 6 | 4,8% |
| Decomposition (TPD) | 3 | 2,4% |
| Testing (TST) | 4 | 3,2% |
| **Total** | **126** | **100%** |

Persentase dibulatkan dan mengukur komposisi direktori, bukan pembagian waktu. RV mencampur Techplan, code/source dan migration-design Review; TST count tidak mengukur luas coverage. Tiga direktori hanya memiliki Invocation: RV-S2-003-002 superseded sebelum dispatch, TP-S2-006-009 tidak di-dispatch karena status reconciliation kemudian dilakukan langsung, serta EXP-S2-008-001 yang baru prepared dan kini parked. Direktori lain memiliki output selain Invocation; keberadaan output sendiri tidak membuktikan semua Run completed atau dispatched.

**TP + RV mencapai 86 dari 126 direktori, sekitar 68,3%.** Dua delivery WU memiliki empat direktori Build: dua frontend, dua backend. Dua backend Build berakhir STALLED; Build pertama menghasilkan safe cap slice, Build kedua berhenti sebelum production/schema write. Enam belas direktori Build lain berada di WU rekonsiliasi, termasuk source/counterpart authoring, patch atau perubahan metadata. Jadi jumlah Build keseluruhan juga tidak dapat disamakan dengan jumlah capability runtime yang berhasil dikirim.

Konsentrasi effort terlihat pada WU002: **46 Run / 153 Markdown**, dan WU003: **28 Run / 78 Markdown**. WU003 sendiri memiliki 13 TP dan 10 RV directories. Sebagai perbandingan, WU004 membutuhkan 11 Run directories dan mencapai frontend mock verification. Ini memperlihatkan bentuk hambatan yang berbeda; tidak membuktikan frontend secara intrinsik lebih mudah atau bahwa 11 Run menjadi standar ideal.

Invocation berjumlah 126 file dengan sekitar 101.409 kata. Techplan/candidate berjumlah 28 file dengan sekitar 160.374 kata. Dua kategori itu mencakup sekitar **51,6%** kata Markdown snapshot. Hanya tiga salinan ekstra ditemukan identik secara byte; ini tidak menilai duplikasi makna karena perubahan provenance kecil sudah mengubah hash. Banyak revisi sah untuk traceability, tetapi jumlah teks yang harus dire-ground bertambah dan dapat menutupi current frontier jika snapshot/history tidak dipisahkan dengan baik.

## 5. Mengapa hasil delivery masih terbatas

### A. Banyak keputusan fundamental diselesaikan setelah siklus planning berjalan

**Observed:** WU002 membutuhkan beberapa material revisions seiring Product/MVP, email/terminal-notification, monetary standard dan D1 decisions. Setelah `CONTRACT_READY`, delivery planning masih mengekspos Campaign action, cap/capacity, precision/storage, unavailable Funding, dan retry credential semantics. WU005/006/007/008 muncul untuk menutup gap yang berbeda.

**Interpretasi:** baseline cukup untuk contract planning tertentu, tetapi belum cukup untuk mengeksekusi seluruh donation spine. Penanda readiness membawa banyak deferred obligations. Saat obligations itu bertemu persistence/API behavior konkret, konsekuensinya memerlukan amendment dan approval baru. Ini adalah sumber rework terbesar yang terlihat, walaupun workspace tidak menyediakan jam per penyebab.

Sebagian kebutuhan memang baru atau memerlukan owner judgement. Namun beberapa interaksi dapat diuji lebih dini pada level scenario: request retry setelah response hilang, Campaign closed, expired credential, Funding absent, dan raw bearer yang tidak boleh disimpan. Scope/authority penting tetap harus dirutekan; perbaikannya adalah memperjelas konsekuensi sebelum full plan dipakai sebagai input Build. Evidence: [Events](../../.harscode-spaces/s2-guest-donation-truthful-state/events.md), [RV003005](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-005/review-findings-1.md), dan [RV003008](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-008/review-findings-1.md).

### B. Dispatch readiness tidak selalu mengikuti dependency yang sudah tertulis

**Observed:** BLD-S2-003-002 di-dispatch setelah exact Tier-0 authorization/model approval dan kesiapan Human pairing. Participant kemudian menemukan bahwa Techplan §10/§11 mensyaratkan migration-design Review yang belum tersedia dan berhenti sebelum writes. Orchestrator mencatat bahwa prerequisite tersebut terlewat sebelum dispatch. [Build report](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/BLD-S2-003-002/report.md) menetapkan hasil STALLED.

**Interpretasi:** ada gap pelaksanaan orchestration, bukan sekadar kekurangan izin atau model. Hasil aman karena Participant fail closed, tetapi dispatch, re-grounding dan kesiapan pairing Human sudah memakai effort tanpa menghasilkan implementasi. Gate yang telah tertulis seharusnya menjadi prasyarat operasional sebelum mobilisasi Build.

### C. Kontrak dan desain belum cukup diuji sebagai satu perjalanan recovery

**Observed:** RV003005 menemukan enam blocking design gaps: unavailable Funding, idempotency equivalence, rollback/operational impact, pending query/integrity, monetary storage, deletion/close-reason compatibility. Setelah revisi dan whole candidate approval, RV003008 masih menemukan response-token gap: accepted POST retry perlu `status_token`, tetapi verifier satu arah tidak dapat mereproduksi bearer yang hilang atau expired.

**Interpretasi:** Review efektif sebagai pengaman, tetapi interaction completeness datang terlambat. Dokumen terpisah bisa masing-masing terlihat masuk akal sementara combined scenario belum dapat dieksekusi. Mekanisme yang berulang ialah "review → temukan dependency baru → authority/plan/source reconciliation → review lagi". Kesimpulan ini tidak menghapus kebutuhan independent Review; ia menuntut scenario synthesis yang lebih baik sebelum Review/approval.

### D. Fidelity dan lifecycle artefak menyebabkan siklus tambahan

Contoh observed yang berbeda dari perubahan Product:

- RV-S2-002-005 menemukan tiga planning-reconciliation blockers; RV-S2-002-009 kemudian menemukan dua fidelity findings terhadap O4/O5. Keduanya memerlukan Planner repair.
- RV003006 menemukan concrete credential recipe dipresentasikan settled meski newer accepted authority masih membiarkannya open.
- TP006004 sudah memperbaiki citation, tetapi report-nya kehilangan Interface Contract yang applicable. TP006005 menjadi report-only repair sebelum Human gate.
- RV003009 menemukan approval predecessor diatribusikan ambigu pada current candidate; TP13 memperbaikinya, lalu RV10 mencatat stale pointer yang tersisa. Review RV10 atas revision itu secara khusus diminta Human; biaya tambahan tersebut tidak boleh diatribusikan semata-mata sebagai requirement universal Harscode.

Evidence: named Run findings dan Events headings 29 September, 1/2 Oktober, serta 4/5 Oktober. **Interpretasi:** masalah ini dapat dikurangi melalui first-pass fidelity dan pemeriksaan completeness pada output phase. Mengganti model belum merupakan solusi yang terbukti; tidak ada eksperimen pembanding model atau data waktu/cost yang cukup.

### E. Current state dan history berulang pada banyak surface

**Observed:** state/frontier diulang pada manifest, Work Graph, Control Surface, Parent Outcome, tracker dan Events. [Stage A baseline 29 September](../../.harscode-spaces/s2-guest-donation-truthful-state/experiments/current-state-simplification/stage-a-baseline.md) sudah mencatat pengulangan ini. Dokumen sekarang masih membawa historical checkpoints yang berjudul current, approval preimage/post-marker identities, dan nomor OI9 yang pernah merujuk unavailable Funding lalu merujuk retry credential.

Saat evaluasi ini, prepared WU008 metadata juga ditemukan membawa HEAD `6e78c49…` dari RV10, sementara git HEAD aktual `7e731f9…`, serta menyebut Feature 02 dengan nama file yang tidak ada. Pointer HEAD dan nama file dikoreksi sebelum dispatch reliance; path sebenarnya `02-donation-status-check.md`. Tidak ada source spec yang diubah.

**Interpretasi:** banyak metadata/provenance belum menjamin freshness atau kejelasan. State yang ditulis berulang menciptakan maintenance cost, peluang contradiction, dan pekerjaan rekonstruksi. Ini merupakan tanggung jawab kualitas Orchestrator juga, bukan hanya Participant.

### F. Risk, breadth dan pengalaman pilot menambah beban yang nyata

WU003 mencakup transaksi Campaign/Donation, money, reservation/closure, schema, status credential, email lifecycle, Campaign draft writes dan minimum Organization/representative source. Techplan Review menilai 20 rules; [TPD003001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TPD-S2-003-001/launch-record.md) memilih Step 0 = NO untuk decomposition. Dua Build berikutnya berhenti pada gate berbeda. Ini memberi alasan untuk menilai ulang execution readiness/batch scope sebelum resume, bukan bukti bahwa decomposition formal selalu diperlukan atau scope boleh dipangkas sepihak.

Pilot juga sempat mengurus visible dispatch, report ownership, authority routing dan current-state simplification. Kencleng delivery dan validasi Harscode berbagi effort Human. Besarnya bagian masing-masing tidak tercatat. Istilah "workflow berhasil" perlu mengukur safety, dispatch reliability dan outcome delivery secara bersama; artefak pilot merupakan hasil pembelajaran yang berguna, tetapi belum menjadi bukti produk selesai.

## 6. Perbandingan effort/artefak dengan outcome

| Investment | Hasil yang diperoleh | Penilaian pada checkpoint |
|---|---|---|
| Discovery/owner/contract reconciliation | Canonical decisions, exact accepted sources, shared money direction dan scoped dependencies | Bernilai dan reusable; jumlah re-entry memperlihatkan kesiapan awal belum cukup untuk delivery spine. |
| Independent Review dan patch loops | Atomic funding/request idempotency gaps, omitted currency code, authority mismatch, migration-design gaps, replay-token semantics dan frontend accessibility ditemukan | Nilai assurance konkret. Biaya dapat ditekan melalui better first-pass synthesis, tanpa meniadakan independent scrutiny. |
| Frontend Build/Review/Testing | Guest flow production UI terhadap contract mocks, rendered acceptance | Konversi ke capability yang jelas dalam scope mock. Integration masih outstanding. |
| Backend Build/re-grounding | Cap projection/wire behavior dan unapplied 000012; fail-closed stop ketika prerequisite belum terpenuhi | Bounded result berguna, tetapi core Donation/runtime belum ter-deliver. Dispatch kedua yang premature adalah avoidable effort. |
| Metadata, Invocation dan repeated projections | Traceability/reconstructability serta exact approval attribution | Maintenance burden tinggi; manfaat marginal lebih rendah ketika perubahan hanya marker/pointer atau state ditulis ulang. |
| Workflow pilot | Bukti kesalahan routing, report completeness dan state simplification | Bernilai bagi Harscode jika dipromosikan menjadi perubahan terukur. Satu report evaluasi tambahan belum menjadi perbaikan workflow. |

Lima contoh Run yang hasil utamanya status-only adalah TP-S2-002-005/008/013/017 dan TP-S2-005-004. Mereka mengikuti posture/guidance historis; report ini tidak menyatakan semuanya melanggar kebijakan saat dijalankan. Current [run contract](../../../harscode-workspace/orchestration/run-contract.md#deterministic-reconciliation-outside-the-run-path) sudah mengizinkan bounded deterministic reconciliation ketika pre/postconditions dan authority memadai. Approval TP006008 dan WU007 kemudian memang direkonsiliasi langsung tanpa dispatch status-only. Ini membuktikan simplifikasi sebagian sudah dapat diterapkan tanpa mengurangi authority gate.

Penilaian akhir hubungan effort/outcome adalah **belum proporsional untuk target delivery produk**, dengan **hasil assurance dan pembelajaran yang substansial**. Tidak semua 507 ribu kata harus dipangkas: evidence history, independent verdict dan exact acceptance memang diperlukan. Yang perlu dikurangi ialah penulisan ulang current truth, re-entry untuk fully determined metadata, repair output yang seharusnya lengkap, serta dispatch yang belum ready. Tidak tersedia denominator untuk ROI finansial, biaya per feature, atau klaim bahwa seluruh effort tertentu sia-sia.

## 7. Proposal perbaikan Harscode

Proposal berikut bertujuan meningkatkan conversion ke evidence/capability dalam lifecycle yang berlaku. Tidak ada canonical Harscode file atau phase policy diubah oleh report ini.

| Prioritas | Perbaikan yang diusulkan | Pemilik dan ukuran keberhasilan |
|---|---|---|
| P0 | Sebelum Build, Orchestrator membuktikan readiness untuk **batch yang akan dijalankan**: exact authority/source, approved plan, prerequisite design Review, protected pairing, dan dependency terkait. Gunakan existing state/receipt, bukan membuat dokumen readiness baru setiap dispatch. | Orchestration; tidak ada Build yang STALLED karena prerequisite yang sudah tertulis tetapi terlewat. |
| P0 | Planner menyintesis recovery scenarios sebelum approval: lost POST response → same-key retry; changed payload; closed Campaign; missing Funding; expired credential; accepted-pending settlement. Kaitkan scenario dengan source/API/persistence dan open owner decision. | Planner + independent Reviewer; contradiction lintas scenario ditemukan sebelum mobilisasi schema Build. Ini targeted application of authority/invariants, bukan template wajib untuk semua task. |
| P0 | Terapkan run qualification yang current: metadata yang fully determined memakai bounded mechanical path; semantic/source changes tetap melalui owning phase dan exact acceptance. Jangan menyamakan perubahan hash dengan perubahan executable/verification meaning. | Orchestration; tidak ada status-only Participant jika direct path memenuhi seluruh preconditions; material change tetap mendapat Review/Human gate. |
| P1 | Rapikan ownership current truth: Work Graph memiliki dependency; manifest memiliki active WU state; Events hanya history; Control Surface projection ringkas; tracker milestone/progress. Setiap surface mengacu pointer tepat, tanpa menyalin narasi panjang seluruh lineage. | Orchestration/Harscode owner; fresh reconstruction menemukan satu frontier konsisten, tanpa competing current states. Perubahan struktur dilakukan dengan preservasi evidence, tidak menghapus history sepihak. |
| P1 | Periksa output completeness sebelum Participant phase dinyatakan selesai: requirement/decision fidelity, applicable report sections, real file anchors, approval target. Periksa bagian relevan dari source yang current, bukan hanya hash metadata. | Role yang menghasilkan output + Orchestrator; tidak ada Run tambahan hanya untuk mengisi bagian report yang memang required atau membetulkan pointer yang sudah discoverable. |
| P1 | Nilai ulang execution batching WU003: pisahkan kapan gated scope membutuhkan owner/pairing dan kapan safe coherent capability dapat dikerjakan. Gunakan canonical decomposition applicability; jika NO, tetap buat urutan batch dan gate jelas dalam owning approved plan. | Planner/Orchestration + Human gates yang berlaku; tiap Build berikutnya punya achievable scope/evidence. Tidak ada pergeseran Product scope atau Tier-0 bypass. |
| P1 | Fasilitasi known-owner decision secara langsung saat evidence sudah decision-ready; Participant discovery baru dipakai hanya ketika ada evidence gap yang perlu diselesaikan. Saat resume, audit apakah WU008 membutuhkan Exploration baru atau cukup bounded planning dari existing evidence. | Orchestration/Human owner; setiap Run baru menyebut material knowledge/delta yang dihasilkan. Route prepared tidak otomatis menjadi mandatory route. |
| P2 | Ukur perubahan secara ringan dari existing reports: Run tujuan, observable output, alasan re-entry, prerequisite miss, first-pass completeness, dan milestone/capability yang bergerak. Catat actual Human/AI time bila Human memilih, tanpa mengarang telemetry yang tidak tersedia. | Harscode evaluator; evaluasi berikutnya dapat membandingkan safety dan conversion, bukan jumlah dokumen saja. |

Pengurangan report-only phases atau perubahan lifecycle/session policy yang lebih jauh harus diperlakukan sebagai proposal Harscode tersendiri. Current canonical fresh Participant/Session untuk new Run, independent Review yang applicable, serta protected human authority tetap berlaku. Safety threshold tidak diturunkan untuk mencapai target jumlah Run.

### Cara menilai apakah perbaikan benar-benar membantu

Untuk eksperimen resume yang disetujui Human, tentukan scope dan observasi sebelum menjalankan, lalu nilai: apakah prerequisite miss hilang; apakah first-pass output lengkap; apakah state dapat direkonstruksi tanpa koreksi; apakah Build menghasilkan scope yang dijanjikan; dan apakah evidence keamanan/financial/integration yang diwajibkan tetap lengkap. Catat jumlah re-entry beserta penyebabnya. Hindari target seperti "harus 50% lebih cepat" tanpa baseline jam, comparable scope, dan data quality.

Ukuran hasil berikutnya yang lebih bermakna adalah **capability backend terimplementasi dengan independent evidence**, kemudian **alur frontend/backend nyata yang verified**. Banyak plan yang direview bersih tidak sendiri menggerakkan milestone itu. Current outstanding source/security/Human gates tetap harus selesai sebelum scope terkait dapat menghasilkan bukti tersebut.

## 8. HOLD, preserved state dan batas evaluasi

WU003 dan WU008 tetap ACTIVE secara execution, dengan scheduling PARKED karena Human HOLD. WU001/002/004/005/006/007 tetap DONE dalam scope masing-masing. EXP-S2-008-001 tetap undispatched; tidak ada Participant development di-dispatch dalam evaluasi ini. HOLD tidak menghasilkan approval, source acceptance, schema verdict, atau risk acceptance.

Candidate exact tetap `e895a1da8b90e7f88c449651a9a46add59e1a1d610cce3cc7b739d0c12c30315`, Draft / In Review. RV10 findings hash `c70564f6219554f210f45f9906213c2050369faafe6d7a2cf35bf670c4a07171` dan launch hash `68d89f04775365c4dc17138e1c676b1be8e5f4a2bccc9f7fdf3de69f22960141` cocok pada pembacaan evaluasi. RV10 tidak merupakan approval kandidat atau migration-design verdict.

Jika Human kelak memberi explicit resume, re-ground git/durable state/current Harscode. Urutan gate yang dipertahankan: **OI9 source/counterpart reconciliation dan exact owner acceptance → koreksi mekanis §13 item 5 bila masih relevan → converged report dan Human approval kandidat exact → fresh positive migration-design Review sebelum schema Build**. Anhar, pada Donation/Security-PII dan API roles, menerima exact changed source/counterpart bytes. WU006/007 receipts tidak menggantikan acceptance baru itu.

Tetap terbuka secara terpisah: **Open Item 7, D1 Human-paired/Tier-0 work, O3/O4/O5, migration 000012 yang belum diaplikasikan, PostgreSQL/runtime, whole-spine Testing, real integration dan slice finalization**. Report ini tidak menetapkan policy baru atau menerima residual risk untuk area-area tersebut.

Metode evaluasi menggunakan inventories, source/manifest/Events/Run evidence, git read-only, dan hash; sampling substantive findings berfokus pada penyebab/critical path yang tercantum. Ini bukan full technical/security re-review semua artefak. Hasil tests/runtime sebelumnya dikutip sebatas carrier dan scope yang mencatatnya; tidak dijalankan ulang. Tidak ada tests, validators, code generators, SQL/migration, database, runtime, browser atau security actions dalam pekerjaan report ini. Production/spec/API/counterpart tetap tidak diedit; writes adalah report/evidence dan reconciliation state HOLD.

Klaim yang belum dapat dibuat: total AI spend/tokens, actual jam per phase, lamanya idle versus aktif, causal savings dari tiap Review, model superiority, persentase complete berbobot, atau ETA selesai. Report dan [lampiran evidence](slice-2-harscode-evaluation-evidence-2026-10-05.md) menyediakan basis untuk evaluasi, bukan pengganti pengukuran tersebut.
