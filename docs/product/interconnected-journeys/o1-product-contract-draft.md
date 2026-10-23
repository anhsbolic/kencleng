# Kencleng — Product Contract O1: Pembentukan Organisasi dan Kewenangan `Owner` Internal Pertama (Draf v0.3)

> **Status:** DRAF PRODUCT CONTRACT / PROPOSED — **BELUM DISETUJUI SEBAGAI INPUT `I1` RANCANG**. Penulisan ulang Bahasa Indonesia menggunakan [template I1 Harscode — Bahasa Indonesia](https://github.com/anhsbolic/harscode-workspace/blob/00ed2881a608c47d3c0afde423252798f6ffe8f7/product/product-contract-template-id-draft.md) (template itu juga masih proposal riset).
> **Lingkup:** Hanya Product O1 pada `anhsbolic/kencleng`, branch `pilot/station5-c1-space-only`. Bukan perubahan Kencleng `main`, bukan prioritas Engineering, dan bukan perintah menjalankan Rancang.
> **Riwayat dan standing:** Berasal dari draf O1 v0.1 Inggris (blob `00497ab3ab0a6183ff54e90089cf54131e560b08`). Anhar menyetujui keputusan Product scoped `PD-O1-01..05` pada 23 Oktober 2026, termasuk keputusan final bahwa `Owner` internal tidak menunggu approval Reviewer. SC/REQ/AC adalah rumusan kontrak yang masih memerlukan review dan approval **exact revision**.
> **Keputusan Human (23 Oktober 2026):** Seluruh kebijakan Product `PD-O1-01..05` dalam scope O1 **DECIDED — scoped**. Hubungan `Owner` internal pertama efektif sejak keberhasilan pembentukan O1, sementara approval Reviewer menilai klaim/status organisasi secara terpisah. **Human policy decision gate CLOSED**; artefak tetap DRAFT, **belum final-approved sebagai `I1` Rancang**. Lihat Section 6.

## Identitas artefak dan authority

| Informasi | Isi dan status |
| --- | --- |
| Proyek dan outcome | Kencleng — `O1`: pembentukan konteks organisasi dan kewenangan `Owner` internal pertama |
| Product owner | Anhar |
| Status Product Contract | **DRAFT / PROPOSED**; belum berlaku sebagai `I1` yang binding |
| Artefak saat ini | `docs/product/interconnected-journeys/o1-product-contract-draft.md` pada `pilot/station5-c1-space-only`; rujuk blob/commit yang nyata saat review |
| Approval | **Keputusan Product O1 `PD-O1-01..05` settled oleh Anhar, 23 Oktober 2026.** Approval Reviewer atas klaim organisasi bukan prasyarat `Owner` internal. **Approval final exact-revision Product Contract O1 belum diberikan**, sehingga SC/REQ/AC belum menjadi `I1` binding. |
| Sumber utama | [Product Intent yang disetujui](../product-intent.md), blob `18fafc31121b49277a6c1b45a7ff6ef57edf0825`; [Peta Journey Level 2 v0.3](level-2-network.md), blob `5783fdb303e34ecef7988d6ec3c7fb8b15326a95` (node O1 dan hubungan terkait) |
| Pendukung keputusan riset | [Selection and Product Output research](selection-and-handoff-research.md), blob `6323010a854aa4eb585500e462fdf5d42960add8`, **working research**, bukan aturan O1 yang disetujui |
| Yang tidak menjadi sumber ketentuan O1 | Paket perilaku/requirements pilot historis dan draft PRD lama yang pernah dirakit dari komitmen tersebut |
| Consumer yang dituju | Calon `I1` Rancang, **setelah** persetujuan Product exact revision; `I2/I3/I4` dan invocation tetap terpisah |

**Urutan authority:** Product Intent dan Level 2 yang disetujui menentukan batas riset. **Standing `PROPOSED` pada Success Criteria, requirements, dan acceptance di bawah berarti turunan dokumen belum mendapat approval final sebagai kontrak**, bukan keputusan Product yang sudah disetujui di Section 6 harus diputuskan ulang. Riwayat bahasa Inggris tetap terlacak di Git dan bukan sumber Product kedua.

## 1. Outcome Product dan Success Criteria

**Pernyataan outcome O1 (arah riset yang dipilih, bukan acceptance yang sudah disetujui):** Seseorang dapat membentuk konteks organisasi baru yang dapat diatribusikan di Kencleng, sehingga organisasi tersebut mempunyai hubungan kewenangan **internal** dengan `Owner` pertama yang sah menurut aturan Product O1 yang nantinya disetujui; keberhasilannya jujur, bertahan sebagai keadaan yang nyata, dan dapat diperiksa kembali.

Outcome ini membantu menyiapkan konteks kewenangan organisasi yang bermakna bagi tindakan berikutnya, **tanpa** mewajibkan peninjauan organisasi, persiapan Campaign, donasi, atau hasil Journey lain. Organisasi yang sudah memiliki konteks sah tidak perlu otomatis mengulangi O1.

### Success Criteria yang diusulkan

Keempat SC berikut adalah **kandidat kondisi keberhasilan semantik**, bukan bukti runtime maupun syarat yang sudah disetujui. Untuk menyebut O1 berhasil sebagaimana diusulkan, kondisi materialnya harus berlaku **bersama**.

| ID | Kapan outcome dapat dinyatakan berhasil menurut proposal | Batas makna yang tidak ikut dibuktikan | Dasar/status |
| --- | --- | --- | --- |
| **SC-O1-01** | Konteks organisasi **dan** hubungan `Owner` internal pertama **sama-sama sah dan benar-benar berlaku** sebagai satu hasil; tindakan dan hubungan `Owner` dapat diatribusikan kepada orang yang bertindak pada saat akibatnya efektif. | Tidak membuktikan kewenangan legal eksternal, tidak mengharuskan identitas sudah lengkap saat awal interaksi. | `PD-O1-01/02` **DECIDED scoped**; `SC` masih draf untuk approval final |
| **SC-O1-02** | Sebelum akibat berlaku, aktor bisa memahami organisasi yang dibentuk, arti dan batas `Owner` internal; setelahnya hubungan aktual tetap dapat diperiksa, bukan hanya melalui pesan sukses sesaat. | Tidak menetapkan layar, waktu sign-in, bentuk persetujuan UI, atau teknologi penyimpanan. | `PD-O1-03` **DECIDED scoped**; `SC` masih draf untuk approval final |
| **SC-O1-03** | Kegagalan, hasil parsial, atau tidak teratribusi tidak diklaim berhasil. Kesamaan nama/klaim atau **Potential Conflicts** antar-konteks dengan ID berbeda **tidak otomatis** menunda, menolak, menggabungkan, atau menutup pembentukan O1 maupun `Owner` internalnya. | Keberhasilan internal bukan bukti approval Reviewer atau validitas klaim organisasi di dunia nyata; tidak menuntut deteksi konflik universal. | `PD-O1-04` **DECIDED scoped**; `SC` masih draf |
| **SC-O1-04** | Keberhasilan O1 **tidak menaikkan klaim organisasi** menjadi hasil peninjauan/verifikasi eksternal, kurasi Campaign, penayangan publik, atau kelayakan donasi. | Tidak menjamin kualitas organisasi atau kesiapan Campaign. | Batas umum Product Intent dan `PD-O1-05` **DECIDED scoped**; `SC` masih draf |

**Definisi keberhasilan gabungan yang diusulkan:** `O1 berhasil = organisasi sah + hubungan Owner internal pertama yang sah dan teratribusi + makna/hasil yang jujur serta dapat dipahami dan diperiksa`. Perlindungan dari klaim sukses palsu tetap mengikat definisi keberhasilan proposal ini. Definisi tersebut **belum diputuskan Anhar**.

## 2. Pengalaman Aktor dan Makna Konsekuensinya

Ini gambaran pengalaman bermakna, **bukan urutan layar** atau desain autentikasi. **Potential Conflicts** adalah istilah umum untuk kemungkinan pertentangan informasi atau klaim. Pada O1, kemiripan nama/klaim tidak otomatis menghalangi pembentukan konteks organisasi terpisah atau pemberian `Owner` internal. Penilaian klaim dan approval organisasi merupakan concern Reviewer yang berbeda dari O1.

| Momen | Apa yang dialami atau dipahami aktor | Konsekuensi Product menurut usulan |
| --- | --- | --- |
| **Kondisi awal** | Seseorang ingin membentuk kehadiran organisasi di Kencleng dan memahami bahwa konteks internal berbeda dari pembuktian hak mewakili organisasi nyata. | Orang boleh memulai tanpa produk mengklaim kewenangan legal eksternalnya telah diverifikasi (`PD-O1-01` dan `O1-R01`, **PROPOSED**). |
| **Sebelum hasil efektif** | Aktor memahami organisasi yang hendak dibentuk, bahwa dirinya akan menjadi `Owner` internal pertama, serta batas makna peran tersebut. | Atribusi orang yang cukup harus tersedia **sebelum akibat berhasil menjadi efektif**, bukan harus sejak awal mengisi informasi (`PD-O1-02/03`). |
| **Tindakan** | Aktor melakukan pembentukan organisasi menurut kewenangan dan syarat Product yang berlaku. | Tindakan **tidak** otomatis menjadi bukti mewakili organisasi secara hukum atau hasil peninjauan. |
| **Berhasil** | Aktor mendapat informasi keberhasilan yang tidak berlebihan; organisasi dan hubungan `Owner` internal benar-benar ada, kemudian tetap dapat diperiksa. | Satu keberhasilan gabungan yang bermakna; bukan pesan sukses untuk salah satu keadaan saja. |
| **Gagal / tidak lengkap / belum pasti** | Aktor tidak menerima kesan keliru bahwa pembentukan sukses. | Atribusi tidak cukup, sebagian proses gagal/terputus, atau hasil tidak diketahui **tidak boleh** diperlakukan sebagai keberhasilan biasa. |
| **Potential Conflicts / nama organisasi serupa** | User A dan User B dapat masing-masing membentuk konteks bernama organisasi K, sepanjang syarat O1 masing-masing terpenuhi. | Setiap konteks mempunyai ID sendiri, dengan `Owner` internal masing-masing yang efektif saat O1 berhasil. Tidak ada penggabungan, penolakan, penundaan, atau penutupan otomatis hanya karena kemiripan; tidak tersirat approval Reviewer. |

## 3. Perilaku Product, Kewenangan, dan Invariants

### 3.1 Keputusan Product O1 yang telah diselesaikan

Anhar menyelesaikan seluruh kebijakan Product O1 `PD-O1-01..05` pada 23 Oktober 2026. Semua keputusan di bawah adalah **DECIDED — scoped**; rumusan SC, requirements, dan AC tetap membutuhkan review dan approval final *exact revision* sebelum dapat menjadi `I1` binding.

| ID | Kebijakan Product O1 | Standing keputusan Product |
| --- | --- | --- |
| **PD-O1-01** | Seseorang dapat memulai pembentukan konteks organisasi baru. **Jika pembentukan berhasil secara efektif**, orang yang bertindak dan dapat diatribusikan secara sah memperoleh hubungan `Owner` **internal Kencleng** yang pertama. Ini bukan hasil verifikasi perwakilan eksternal. Verifikasi hukum eksternal yang telah selesai **tidak menjadi syarat universal** untuk kewenangan internal terbatas ini. | **DECIDED — scoped**, Anhar 23 Oktober 2026; revisi kontrak final belum approved. |
| **PD-O1-02** | Hasil berhasil merupakan **satu akibat gabungan**: konteks organisasi sah **dan** hubungan `Owner` awal sah/teratribusi. Atribusi aktor wajib tersedia saat akibat menjadi efektif, tidak harus saat awal interaksi. Satu komponen saja tidak boleh dinyatakan berhasil penuh. | **DECIDED — scoped**, Anhar 23 Oktober 2026; revisi kontrak final belum approved. |
| **PD-O1-03** | Aktor dapat memahami akibat dan batas `Owner` sebelum efektif. Setelah berhasil, organisasi serta hubungan `Owner` aktual dapat diperiksa kembali secara durable, bukan hanya dari notifikasi sementara. | **DECIDED — scoped**, Anhar 23 Oktober 2026; revisi kontrak final belum approved. |
| **PD-O1-04** | Pembentukan yang gagal, parsial, atau tidak dapat diatribusikan **tidak boleh dinyatakan berhasil**. Kesamaan nama/klaim atau Potential Conflicts **tidak otomatis** menghalangi dua konteks organisasi dengan ID internal berbeda atau menunda/menolak/menggabungkan/menutup hubungan `Owner` internal masing-masing. Penilaian serta approval atas klaim/status organisasi dilakukan oleh Reviewer dalam proses terpisah dan **tidak menentukan efektivitas `Owner` internal pertama**. | **DECIDED — scoped**, Anhar 23 Oktober 2026. Tidak mewajibkan conflict detection, deduplication, adjudication, atau otomatisasi penutupan. |
| **PD-O1-05** | Informasi yang diberikan organisasi tetap **informasi dari organisasi**. Pembentukan organisasi saja tidak melakukan atau menyiratkan peninjauan organisasi, verifikasi legal eksternal, kurasi Campaign, penayangan publik, atau izin menerima donasi. | **DECIDED — scoped**, Anhar 23 Oktober 2026; selaras prinsip Product Intent. Revisi kontrak final belum approved. |

**Batas yang disetujui:** Dua konteks organisasi dengan ID berbeda tidak otomatis menjadi salah hanya karena nama/klaim yang sama. Jika kelak Reviewer menyetujui dua klaim organisasi yang ternyata merujuk pihak nyata yang sama, itu adalah masalah kualitas keputusan peninjauan/pengelolaan organisasi, **bukan kegagalan otomatis O1**. Mekanisme koreksi Reviewer atau conflict resolution tidak dibahas dalam O1.

### 3.2 Requirements dan invariants yang diusulkan

| ID | Perilaku Product yang diusulkan | Syarat/kewenangan | Invariant / yang tidak otomatis dijamin | Dasar |
| --- | --- | --- | --- | --- |
| **O1-R01** | Seseorang dapat memulai pembentukan tanpa produk menyatakan representasi legal eksternal telah dibuktikan. | Kelayakan aktor spesifik di luar batas ini memerlukan keputusan Product sendiri. | Memulai tidak otomatis berarti berwenang secara hukum atau berhasil menjadi `Owner`. | `PD-O1-01` **DECIDED scoped**; requirement **DRAFT** |
| **O1-R02** | Saat berhasil, orang yang bertindak benar-benar teratribusi ke organisasi dan hubungan `Owner` internal pertama. | Atribusi yang cukup tersedia **pada saat akibat efektif**. | Hubungan `Owner` tidak boleh tak bertuan/tidak teratribusi. | `PD-O1-01/02` **DECIDED scoped**; requirement **DRAFT** |
| **O1-R03** | Akibat pembentukan dan batas kewenangan `Owner` internal dapat dipahami sebelum berlaku. | Sebelum tindakan menghasilkan perubahan material. | Tidak berarti sudah diverifikasi sebagai perwakilan legal. | `PD-O1-03` **DECIDED scoped**; requirement **DRAFT** |
| **O1-R04** | Keberhasilan hanya diklaim jika organisasi **dan** `Owner` internal pertama sama-sama sah/efektif. | Kedua fakta Product benar-benar terjadi. | Satu keadaan yang baru sebagian tidak boleh diklaim sebagai sukses penuh. | `PD-O1-02` **DECIDED scoped**; requirement **DRAFT** |
| **O1-R05** | Organisasi dan hubungan `Owner` saat ini dapat diperiksa kembali beserta makna internalnya. | Setelah keberhasilan nyata. | Pesan sukses sementara tidak menggantikan keadaan durable. | `PD-O1-03` **DECIDED scoped**; requirement **DRAFT** |
| **O1-R06** | Klaim dari organisasi tidak otomatis meningkat menjadi hasil peninjauan, bukti eksternal, kurasi, atau kategori fakta independen. | Setiap kali klaim organisasi disajikan/ditafsirkan. | Pembentukan O1 bukan peninjauan organisasi atau legitimasi legal eksternal. | `PD-O1-05` **DECIDED scoped** + prinsip Product Intent; requirement **DRAFT** |
| **O1-R07** | Kegagalan, keadaan parsial, hasil tidak teratribusi, atau interupsi dibedakan dari keberhasilan asli. | Ketika keadaan negatif atau belum pasti terjadi. | Retry/pemulihan bukan kewajiban sistem Product baru kecuali diputuskan kemudian; tidak boleh membuat sukses palsu. | `PD-O1-04` **DECIDED scoped**; requirement **DRAFT** |
| **O1-R08** | Kesamaan nama/klaim atau Potential Conflicts antar-konteks tidak otomatis menghalangi keberhasilan pembentukan O1 masing-masing; tiap konteks memiliki ID sendiri dan `Owner` internal yang efektif bila syarat internalnya terpenuhi. | Approval Reviewer atas klaim/status organisasi merupakan proses berbeda, bukan syarat `Owner` internal O1. | Tidak memaksa deteksi konflik, penggabungan, penundaan, penolakan, penutupan otomatis, atau klaim bahwa organisasi sudah disetujui Reviewer. | `PD-O1-04` **DECIDED scoped**; requirement **DRAFT** |

**Invariants Product O1 (rumusan untuk approval final):** Tidak ada `Owner` internal awal yang tidak dapat diatribusikan; tidak ada keberhasilan penuh dari hasil parsial; nama/klaim serupa tidak memicu penutupan atau pembatalan otomatis konteks/`Owner` internal yang berbeda ID; dan keberhasilan O1 tidak menjadi label telah ditinjau/diapprove Reviewer.

## 4. Acceptance Criteria dan Pemeriksaan Batas Makna

Semua AC berikut **PROPOSED / BELUM DIUJI DI RUNTIME**. ID tetap mengikuti draf Inggris agar riwayat bisa ditelusuri. Tautan `SC` menunjukkan kondisi keberhasilan yang diperiksa, **bukan** bahwa SC tersebut sudah di-approve.

| ID | Requirement / SC terkait | Skenario dan hasil Product yang dapat diperiksa | Klaim salah yang harus dicegah |
| --- | --- | --- | --- |
| **O1-AC01** | `R01 / SC-O1-04` | Seseorang dapat memulai tanpa dianggap sudah memiliki kewenangan mewakili organisasi secara eksternal yang terbukti. | Kesan legal authority sudah diperiksa. |
| **O1-AC02** | `R02 / SC-O1-01` | Keberhasilan tidak mungkin menjadi efektif jika tindakan atau hubungan `Owner` pertama tidak dapat diatribusikan secara nyata kepada aktor. | Pemberian `Owner` tanpa atribusi. |
| **O1-AC03** | `R03 / SC-O1-02` | Sebelum efek berlaku, aktor bisa memahami organisasi yang dibentuk, `Owner` internal dan hal yang tidak dibuktikan peran itu. | `Owner` internal dianggap legal representation eksternal. |
| **O1-AC04** | `R04 / SC-O1-01` | Organisasi sah bersama hubungan `Owner` awal teratribusi dapat dialami/diperiksa sebagai **satu keberhasilan**; pasangan parsial tidak boleh menjadi sukses biasa. | Keberhasilan hanya karena salah satu keadaan ada. |
| **O1-AC05** | `R05 / SC-O1-02` | Setelah berhasil, pemeriksaan terpisah dari pesan sukses menampilkan hubungan organisasi/`Owner` terkini dan batas maknanya. | Menganggap notifikasi sementara bukti keadaan durable. |
| **O1-AC06** | `R06 / SC-O1-04` | Klaim informasi organisasi tetap dapat diketahui asalnya, tidak otomatis menjadi hasil peninjauan, verifikasi eksternal, atau kurasi Campaign. | Klaim organisasi diangkat menjadi fakta independen. |
| **O1-AC07** | `R07 / SC-O1-03` | Hasil tidak teratribusi, gagal, sebagian, atau belum pasti tidak tampak sebagai sukses biasa; opsi pemulihan tidak boleh mengubah keadaan parsial menjadi sukses palsu. | Menyamakan kegagalan/uncertain dengan berhasil. |
| **O1-AC08** | `R08 / SC-O1-03` | User A dan User B membentuk konteks bernama organisasi K dengan dua ID berbeda. Jika masing-masing memenuhi syarat O1, setiap konteks dan `Owner` internal efektif tanpa auto-merge, auto-close, atau conflict-based blocking; pembentukan tidak menciptakan status approval Reviewer. | Menyimpulkan nama serupa berarti satu ID, menentukan pemenang otomatis, atau menganggap organisasi telah disetujui Reviewer. |
| **O1-AC09** | `R06 / SC-O1-04` | O1 selesai tidak otomatis menghasilkan persiapan Campaign, penayangan publik, izin menerima donasi, atau peninjauan organisasi. | Klaim side effect lintas-Journey yang tidak termasuk O1. |

Acceptance ini adalah konsekuensi Product yang bisa diperiksa bila nanti diimplementasikan, **bukan** resep database transaction, daftar tes API, atau bukti pengujian yang sudah dijalankan.

## 5. Dependencies, Cakupan, dan Kebebasan Solusi

| Hal | Keadaan untuk O1 |
| --- | --- |
| **Prasyarat Product** | Identitas/atribusi orang yang sah saat hasil efektif dan kemampuan menyatakan organisasi serta hubungan `Owner` internal yang nyata. **Ketersediaannya dalam implementasi sekarang belum diverifikasi.** |
| **Sinyal penting lintas concern** | Tiap konteks organisasi harus berbeda secara identitas internal dan memiliki atribusi `Owner` yang sah. **Sinyal Potential Conflicts dan approval Reviewer tidak menjadi prasyarat universal pembentukan O1.** Aturan identity/security proyek diperiksa terpisah saat Engineering. |
| **Dalam cakupan** | Pembentukan konteks organisasi teratribusi, hubungan `Owner` internal pertama, pemahaman sebelum akibat, keberhasilan gabungan, pemeriksaan kembali, dan hasil negatif/unknown yang material. |
| **Di luar cakupan** | Matriks permission lengkap `Owner`/`Staff`; perpindahan/delegasi staf; verifikasi legal eksternal; tata cara, standar serta koreksi keputusan Reviewer; Campaign; publikasi/donasi; dana/pertanggungjawaban; universal duplicate/conflict detection, penggabungan/penutupan otomatis, adjudikasi konflik, atau subsystem recovery wajib. |
| **Kebebasan Design** | Formulir dan isinya, penyampaian konsekuensi, alur identitas/sign-in, interaksi, copy, serta keadaan gagal/belum pasti sesuai makna Product. Status peninjauan organisasi, bila ditampilkan, harus dibedakan dari keberhasilan O1 tanpa membuat proses Reviewer sebagai syarat pembentukan. |
| **Kebebasan Engineering** | Mekanisme autentikasi, penyimpanan/data model, konsistensi/transaksi, API/komponen, pemulihan bila dibutuhkan, teknik pengujian; tanpa melemahkan makna dan safeguard yang nanti disetujui. |
| **I2/I3 yang belum diuji** | Sumber Design, arsitektur, keamanan, protected paths, repository/ref/kode runtime dan consumer expectations harus disuplai dan diverifikasi terpisah saat assignment. Tidak ada klaim kesiapan Engineering dari draf ini. |

**Pembatas pemaknaan:** Keberadaan node O1 di peta yang disetujui tidak membuktikan kemampuan runtime atau menjadi persyaratan bahwa seluruh J-ORG harus selesai sebelum Campaign bisa disiapkan.

## 6. Open Decisions & Blockers — Human Decision Gate

> **PRODUCT POLICY DECISIONS SETTLED — FINAL CONTRACT APPROVAL STILL REQUIRED**
>
> Anhar telah menyelesaikan keputusan Product `PD-O1-01..05` pada **23 Oktober 2026**. `Owner` internal pertama efektif sejak pembentukan O1 berhasil; **Reviewer menilai klaim/status organisasi, bukan memberikan `Owner` internal pertama**. Tidak ada Open Product decision material tersisa dalam scope O1 berdasarkan keputusan saat ini.
>
> **WARNING — NOT READY FOR ENGINEERING HANDOFF:** Draf ini perlu konsistensi SC/REQ/AC dan **approval final exact revision** Product Contract sebelum diserahkan sebagai input `I1` binding. Keputusan Product tidak otomatis memenuhi `I2/I3/I4` atau mengizinkan invocation Rancang.

### 6.1 Keputusan Human yang sudah dituntaskan

| ID | Keputusan Product | Status |
| --- | --- | --- |
| **PD-O1-01** | Pembentukan konteks organisasi menghasilkan `Owner` internal pertama yang teratribusi; bukan representasi legal eksternal. | **DECIDED — scoped / 23 Oktober 2026** |
| **PD-O1-02** | O1 berhasil ketika konteks organisasi dan hubungan `Owner` internal sah dan efektif bersama. | **DECIDED — scoped / 23 Oktober 2026** |
| **PD-O1-03** | Makna `Owner` dijelaskan sebelum efektif; hubungan aktual dapat diperiksa kembali setelah berhasil. | **DECIDED — scoped / 23 Oktober 2026** |
| **PD-O1-04** | Dua konteks organisasi dengan nama/klaim sama dan ID internal berbeda boleh mempunyai `Owner` internal efektif masing-masing jika memenuhi syarat O1. Potential Conflicts tidak otomatis menunda/menolak/menggabungkan/menutup salah satu. Reviewer berwenang menilai klaim/status organisasi secara terpisah. | **DECIDED — scoped / 23 Oktober 2026** |
| **PD-O1-05** | O1 bukan approval Reviewer, verifikasi legal eksternal, kurasi Campaign, publikasi, atau kelayakan donasi. | **DECIDED — scoped / 23 Oktober 2026** |

**Tidak ditambahkan ke O1:** conflict resolution engine, auto-deduplication, aturan koreksi approval Reviewer, atau perancangan seluruh pipeline Review. Hal tersebut tidak diperlukan untuk keputusan Product O1 sekarang.

### 6.2 Status gate dan persyaratan handoff

- **Human Product policy decisions:** **CLOSED** — `PD-O1-01..05` sudah diputuskan dalam scope O1.
- **Product Contract final consistency review:** **PENDING** — cek outcome, `SC`, pengalaman aktor, `REQ`, invariant, dan `AC` pada revisi terbaru sebagai satu kontrak.
- **Exact-revision Human Product Contract approval:** **PENDING** — Anhar perlu menyetujui artefak revisi tepat setelah final review.
- **Engineering handoff:** **BLOCKED** sampai approval Product Contract final; `I2/I3/I4` serta kesiapan Rancang ditangani terpisah.

## 7. Kesiapan Konsumsi dan Batas Handoff

| Pertanyaan kesiapan | Keadaan yang benar-benar diketahui |
| --- | --- |
| **Exact Product authority dan approval** | Product Intent dan Level 2 v0.3 approved; `PD-O1-01..05` **DECIDED — scoped, Anhar 23 Oktober 2026**. Seluruh `SC`, requirements, dan `AC` telah direvisi agar selaras tetapi belum disetujui sebagai Product Contract pada exact revision. |
| **Konsistensi outcome → SC → rules → AC** | SC, requirements dan AC diselaraskan dengan keputusan scoped O1 terbaru, termasuk pembatasan Potential Conflicts pada O1. **Pemeriksaan silang dokumen internal telah dilakukan**; ini belum menjadi persetujuan Product Contract ataupun bukti independent fresh-reader validation. |
| **Human Decision Gate** | **POLICY DECISIONS CLOSED — lihat Section 6.** Final consistency review dan approval exact revision Product Contract **PENDING**. **NOT READY FOR ENGINEERING HANDOFF.** |
| **Syarat Product dan ketersediaan teknis** | Atribusi dan persistensi diperlukan untuk keberhasilan internal O1. Keadaan kode/lingkungan dan autentikasi **UNVERIFIED**. Sinyal konflik dan proses Reviewer bukan dependency universal O1. |
| **Rancang `I1` generic fit** | Struktur memuat outcome, aturan, acceptance, source revisions, exclusions dan delegated freedom. **Candidate only; belum diuji oleh consumer Rancang independen.** |
| **Rancang project-bound fit** | **NOT CHECKED / BLOCKED for binding use** hingga Product Contract exact revision approved, input `I2/I3/I4` diperiksa terpisah, dan readiness Rancang terbukti. |
| **Hasil handoff** | **DRAFT — NOT I1 READY.** Tidak ada izin menjalankan Rancang, tidak ada perubahan operational `main`, dan tidak ada bukti perilaku sudah delivered. |

**Tindakan Product berikutnya:** Lakukan final Product consistency review terhadap seluruh SC/REQ/AC draf v0.3, lalu minta approval Anhar untuk **exact revision Product Contract final**. Tidak perlu kembali membahas kebijakan `PD-O1-01..05` tanpa evidence baru. Engineering dimulai terpisah setelah Product selesai.

## 8. Informasi Material Tambahan (Kondisional)

**Konteks penting yang tidak tertampung sebagai perilaku O1:** Dalam eksperimen sebelum jalur Journey saat ini, pernah ada PRD specimen yang dirakit dari paket requirement komitmen historis (label lama `C1`). Hasil penghitungan coverage spesimen tersebut **hanya membuktikan kesetiaan penyalinan paket historis**, bukan kesesuaian Product Contract O1 baru.

- **Mengapa dicatat di bagian tambahan:** Ini **risiko pembacaan provenance** yang khas untuk consumer O1, bukan requirement atau Success Criteria O1 baru.
- **Sumber dan standing:** [Selection and Product Output research](selection-and-handoff-research.md), working research yang menerima keputusan pemisahan historical requirements dari input Product baru.
- **Konsekuensi bila diabaikan:** Consumer mungkin menyamakan contoh atau persyaratan `C1` dengan keputusan O1 yang baru, kemudian menganggap Product authority/acceptance telah terpenuhi padahal belum.
- **Pemilik tindakan:** Product owner Kencleng menjaga keputusan O1 hanya dari authority yang berlaku; Rancang, jika kelak diinvokasi, harus membaca `I1` approved asli dan tidak mengisi celah semantik dari arsip.

**Bagian tambahan ini tidak menjadi sumber aturan baru.** Keputusan Product yang masih blocking tetap dicatat di bagian 6; semua aturan yang diusulkan tetap di bagian 3.

---

**Status keseluruhan:** Draf O1 ini mematuhi **struktur usulan** Harscode v0.1 Bahasa Indonesia, tetapi **bukan kontrak Product yang approved** dan tidak boleh digunakan sebagai `I1` binding sampai persetujuan Product exact revision serta pemeriksaan consumer yang sesuai.


## Catatan untuk Rancang dan consumer downstream — referensi upstream

**Kode upstream** seperti `PD-O1-01..05`, kode node Journey `O1`, serta tautan ke Product Intent dan riset Product merupakan **penanda provenance dan keputusan yang sudah direkonsiliasi oleh Product**. Setelah kontrak O1 ini *benar-benar disetujui pada revisi yang tepat*, Rancang/consumer downstream **tidak perlu membuka, memvalidasi ulang, atau menafsirkan kode dan referensi upstream tersebut untuk menemukan aturan Product tambahan**. Isi kontrak yang disetujui harus cukup untuk memahami seluruh perilaku yang mengikat. `SC-O1-*`, `O1-R*`, dan `O1-AC*` adalah ID **di dalam kontrak ini** untuk membantu keterlacakan hasil, requirement, dan acceptance—bukan kewajiban pencarian sumber upstream.

**Pemisahan pemeriksaan:** Larangan *upstream Product source re-check* **tidak** berarti melewati verifikasi persetujuan dan revisi tepat `I1`, pemeriksaan konsistensi kontrak yang diserahkan, `I2/I3`, keamanan, atau kewenangan `I4`. Ketidaksesuaian Product yang material harus dikembalikan kepada Anhar, bukan diperbaiki dengan melakukan Product discovery ulang. **Definisi Rancang v3 yang terbentuk saat ini masih mewajibkan source inspection**; pembatasan baru ini adalah arah usulan dan baru berlaku pada Rancang setelah revisi definisi dan approval yang sah.

**Status kontrak saat ini tetap DRAFT / NOT I1 READY.** Keputusan scoped `PD-O1-01..05` sudah settled. Konsistensi akhir dan approval final *exact revision* Product Contract masih pending. Catatan downstream tidak menjadikan draf otomatis `I1` binding.
