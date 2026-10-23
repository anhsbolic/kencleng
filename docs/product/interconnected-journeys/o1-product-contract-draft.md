# Kencleng — Product Contract O1: Pembentukan Organisasi dan Kewenangan `Owner` Internal Pertama (Draf v0.2)

> **Status:** DRAF PRODUCT CONTRACT / PROPOSED — **BELUM DISETUJUI SEBAGAI INPUT `I1` RANCANG**. Penulisan ulang Bahasa Indonesia menggunakan [template I1 Harscode — Bahasa Indonesia](https://github.com/anhsbolic/harscode-workspace/blob/00ed2881a608c47d3c0afde423252798f6ffe8f7/product/product-contract-template-id-draft.md) (template itu juga masih proposal riset).
> **Lingkup:** Hanya Product O1 pada `anhsbolic/kencleng`, branch `pilot/station5-c1-space-only`. Bukan perubahan Kencleng `main`, bukan prioritas Engineering, dan bukan perintah menjalankan Rancang.
> **Riwayat yang diadaptasi:** Draf Product O1 v0.1 Inggris, blob `00497ab3ab0a6183ff54e90089cf54131e560b08`. Semua `PD-O1-01..05`, `O1-R01..08`, dan `O1-AC01..09` tetap proposal menurut makna aslinya; Success Criteria dan trace antarbagian ditambahkan sebagai **struktur usulan untuk review**, bukan persetujuan aturan baru.
> **Batas keputusan:** Persetujuan Anhar untuk **menulis draf mengikuti template** tidak berarti approval atas kebijakan pemberian `Owner` pertama, keberhasilan gabungan, atau perlakuan Potential Conflicts; keputusan itu tetap milik Anhar.

## Identitas artefak dan authority

| Informasi | Isi dan status |
| --- | --- |
| Proyek dan outcome | Kencleng — `O1`: pembentukan konteks organisasi dan kewenangan `Owner` internal pertama |
| Product owner | Anhar |
| Status Product Contract | **DRAFT / PROPOSED**; belum berlaku sebagai `I1` yang binding |
| Artefak saat ini | `docs/product/interconnected-journeys/o1-product-contract-draft.md` pada `pilot/station5-c1-space-only`; rujuk blob/commit yang nyata saat review |
| Approval | **BELUM ADA** untuk `PD-O1-01..05`, Success Criteria, requirements, dan acceptance dalam cakupan O1 |
| Sumber utama | [Product Intent yang disetujui](../product-intent.md), blob `18fafc31121b49277a6c1b45a7ff6ef57edf0825`; [Peta Journey Level 2 v0.3](level-2-network.md), blob `5783fdb303e34ecef7988d6ec3c7fb8b15326a95` (node O1 dan hubungan terkait) |
| Pendukung keputusan riset | [Selection and Product Output research](selection-and-handoff-research.md), blob `6323010a854aa4eb585500e462fdf5d42960add8`, **working research**, bukan aturan O1 yang disetujui |
| Yang tidak menjadi sumber ketentuan O1 | Paket perilaku/requirements pilot historis dan draft PRD lama yang pernah dirakit dari komitmen tersebut |
| Consumer yang dituju | Calon `I1` Rancang, **setelah** persetujuan Product exact revision; `I2/I3/I4` dan invocation tetap terpisah |

**Urutan authority:** Product Intent dan Level 2 yang disetujui menentukan batas riset. Semua penjabaran O1 yang dinyatakan `PROPOSED` di bawah **memerlukan keputusan baru** dari Product owner. Riwayat bahasa Inggris bisa dilacak melalui Git dan bukan sumber Product kedua.

## 1. Outcome Product dan Success Criteria

**Pernyataan outcome O1 (arah riset yang dipilih, bukan acceptance yang sudah disetujui):** Seseorang dapat membentuk konteks organisasi baru yang dapat diatribusikan di Kencleng, sehingga organisasi tersebut mempunyai hubungan kewenangan **internal** dengan `Owner` pertama yang sah menurut aturan Product O1 yang nantinya disetujui; keberhasilannya jujur, bertahan sebagai keadaan yang nyata, dan dapat diperiksa kembali.

Outcome ini membantu menyiapkan konteks kewenangan organisasi yang bermakna bagi tindakan berikutnya, **tanpa** mewajibkan peninjauan organisasi, persiapan Campaign, donasi, atau hasil Journey lain. Organisasi yang sudah memiliki konteks sah tidak perlu otomatis mengulangi O1.

### Success Criteria yang diusulkan

Keempat SC berikut adalah **kandidat kondisi keberhasilan semantik**, bukan bukti runtime maupun syarat yang sudah disetujui. Untuk menyebut O1 berhasil sebagaimana diusulkan, kondisi materialnya harus berlaku **bersama**.

| ID | Kapan outcome dapat dinyatakan berhasil menurut proposal | Batas makna yang tidak ikut dibuktikan | Dasar/status |
| --- | --- | --- | --- |
| **SC-O1-01** | Konteks organisasi **dan** hubungan `Owner` internal pertama **sama-sama sah dan benar-benar berlaku** sebagai satu hasil; tindakan dan hubungan `Owner` dapat diatribusikan kepada orang yang bertindak pada saat akibatnya efektif. | Tidak membuktikan kewenangan legal eksternal, tidak mengharuskan identitas sudah lengkap saat awal interaksi. | `PD-O1-01/02` **PROPOSED**, selaras dengan batas O1 Level 2 |
| **SC-O1-02** | Sebelum akibat berlaku, aktor bisa memahami organisasi yang dibentuk, arti dan batas `Owner` internal; setelahnya hubungan aktual tetap dapat diperiksa, bukan hanya melalui pesan sukses sesaat. | Tidak menetapkan layar, waktu sign-in, bentuk persetujuan UI, atau teknologi penyimpanan. | `PD-O1-03` **PROPOSED**; prinsip keterbukaan Product Intent |
| **SC-O1-03** | Klaim keberhasilan sesuai fakta: keadaan tidak lengkap, gagal, tidak dapat diatribusikan, atau **Potential Conflicts yang material, sudah benar-benar diketahui, dan belum selesai terkait klaim kewenangan awal** tidak diperlakukan sebagai keberhasilan tanpa masalah. | Tidak mewajibkan deteksi seluruh kemungkinan konflik, maupun satu mekanisme penyelesaian konflik. | `PD-O1-04` **PROPOSED** |
| **SC-O1-04** | Keberhasilan O1 **tidak menaikkan klaim organisasi** menjadi hasil peninjauan/verifikasi eksternal, kurasi Campaign, penayangan publik, atau kelayakan donasi. | Tidak menjamin kualitas organisasi atau kesiapan Campaign. | Batas umum Product Intent diterapkan pada O1 lewat `PD-O1-05` (konsekuensi O1 tetap **PROPOSED**) |

**Definisi keberhasilan gabungan yang diusulkan:** `O1 berhasil = organisasi sah + hubungan Owner internal pertama yang sah dan teratribusi + makna/hasil yang jujur serta dapat dipahami dan diperiksa`. Perlindungan dari klaim sukses palsu tetap mengikat definisi keberhasilan proposal ini. Definisi tersebut **belum diputuskan Anhar**.

## 2. Pengalaman Aktor dan Makna Konsekuensinya

Ini gambaran pengalaman bermakna, **bukan urutan layar** atau desain autentikasi. **Potential Conflicts** adalah label umum untuk kondisi konflik yang berpotensi mengubah kebenaran klaim Product; pada O1 yang dibahas hanya konflik yang **material terhadap klaim kewenangan awal, benar-benar diketahui Kencleng, dan belum terselesaikan**. Label ini tidak berarti setiap dugaan konflik otomatis menghalangi O1 atau Kencleng wajib mendeteksi semua konflik.

| Momen | Apa yang dialami atau dipahami aktor | Konsekuensi Product menurut usulan |
| --- | --- | --- |
| **Kondisi awal** | Seseorang ingin membentuk kehadiran organisasi di Kencleng dan memahami bahwa konteks internal berbeda dari pembuktian hak mewakili organisasi nyata. | Orang boleh memulai tanpa produk mengklaim kewenangan legal eksternalnya telah diverifikasi (`PD-O1-01` dan `O1-R01`, **PROPOSED**). |
| **Sebelum hasil efektif** | Aktor memahami organisasi yang hendak dibentuk, bahwa dirinya akan menjadi `Owner` internal pertama, serta batas makna peran tersebut. | Atribusi orang yang cukup harus tersedia **sebelum akibat berhasil menjadi efektif**, bukan harus sejak awal mengisi informasi (`PD-O1-02/03`). |
| **Tindakan** | Aktor melakukan pembentukan organisasi menurut kewenangan dan syarat Product yang berlaku. | Tindakan **tidak** otomatis menjadi bukti mewakili organisasi secara hukum atau hasil peninjauan. |
| **Berhasil** | Aktor mendapat informasi keberhasilan yang tidak berlebihan; organisasi dan hubungan `Owner` internal benar-benar ada, kemudian tetap dapat diperiksa. | Satu keberhasilan gabungan yang bermakna; bukan pesan sukses untuk salah satu keadaan saja. |
| **Gagal / tidak lengkap / belum pasti** | Aktor tidak menerima kesan keliru bahwa pembentukan sukses. | Atribusi tidak cukup, sebagian proses gagal/terputus, atau hasil tidak diketahui **tidak boleh** diperlakukan sebagai keberhasilan biasa. |
| **Potential Conflicts yang diketahui** | Bila ada konflik material terkait klaim kewenangan awal yang benar-benar diketahui dan belum selesai, aktor tidak mendapat kesan keliru bahwa kewenangan tersebut tanpa konflik. | Klaim yang terdampak harus tetap belum terselesaikan atau ditolak secara jujur; **bentuk pembatasan hubungan `Owner` yang efektif belum diputuskan lebih jauh**. |

## 3. Perilaku Product, Kewenangan, dan Invariants

### 3.1 Satu paket keputusan Product yang masih diusulkan

Lima butir ini merupakan **satu kebijakan pembentukan O1 yang saling terkait**, bukan lima fakta yang sudah disetujui terpisah.

| ID | Usulan keputusan O1 | Standing |
| --- | --- | --- |
| **PD-O1-01** | Seseorang dapat memulai pembentukan konteks organisasi baru. **Jika pembentukan berhasil secara efektif**, orang yang bertindak dan dapat diatribusikan secara sah memperoleh hubungan `Owner` **internal Kencleng** yang pertama. Ini bukan hasil verifikasi perwakilan eksternal. Verifikasi hukum eksternal yang telah selesai **tidak diusulkan sebagai syarat universal** untuk kewenangan internal terbatas ini. | **PROPOSED** — Product Intent tidak menetapkan model pemberian `Owner` pertama. |
| **PD-O1-02** | Hasil berhasil merupakan **satu akibat gabungan**: konteks organisasi sah **dan** hubungan `Owner` awal sah/teratribusi. Atribusi aktor wajib tersedia saat akibat menjadi efektif, tidak harus saat awal interaksi. Satu komponen saja tidak boleh dinyatakan berhasil penuh. | **PROPOSED** — aturan keberhasilan gabungan dan waktu atribusi masih keputusan Product. |
| **PD-O1-03** | Aktor dapat memahami akibat dan batas `Owner` sebelum efektif. Setelah berhasil, organisasi serta hubungan `Owner` aktual dapat diperiksa kembali secara durable, bukan hanya dari notifikasi sementara. | **PROPOSED** — penjabaran scoped dari prinsip kejelasan/kebenaran Product Intent. |
| **PD-O1-04** | Pembentukan yang tidak lengkap, gagal, atau tak dapat diatribusikan **tidak boleh dinyatakan berhasil**. Jika **Potential Conflicts yang material bagi klaim kewenangan awal** benar-benar diketahui dan belum terselesaikan, Kencleng tidak boleh menyatakan keberhasilan kewenangan awal seolah tanpa konflik; klaim terkait tetap unresolved atau ditolak dengan makna yang jujur. | **PROPOSED** — tidak mewajibkan deteksi konflik universal, penyelesaian konflik eksternal, UX khusus konflik, atau layanan teknis baru. |
| **PD-O1-05** | Informasi yang diberikan organisasi tetap **informasi dari organisasi**. Pembentukan organisasi saja tidak melakukan atau menyiratkan peninjauan organisasi, verifikasi legal eksternal, kurasi Campaign, penayangan publik, atau izin menerima donasi. | Pembedaan jenis informasi **didukung Product Intent**; aplikasi konsekuensinya pada O1 tetap bagian paket **PROPOSED**. |

**Keputusan yang masih OPEN:** Usulan untuk **melarang seluruh hubungan `Owner` menjadi efektif selama Potential Conflicts yang relevan dan benar-benar diketahui belum selesai** pernah muncul sebagai rekomendasi diskusi, tetapi **belum disetujui Anhar dan tidak dimasukkan sebagai aturan baru** dalam draf ini. `PD-O1-04` mempertahankan safeguard kejujuran pada proposal asal; konsekuensi tepatnya membutuhkan keputusan Product sebelum menjadi I1 binding.

### 3.2 Requirements dan invariants yang diusulkan

| ID | Perilaku Product yang diusulkan | Syarat/kewenangan | Invariant / yang tidak otomatis dijamin | Dasar |
| --- | --- | --- | --- | --- |
| **O1-R01** | Seseorang dapat memulai pembentukan tanpa produk menyatakan representasi legal eksternal telah dibuktikan. | Kelayakan aktor spesifik di luar batas ini memerlukan keputusan Product sendiri. | Memulai tidak otomatis berarti berwenang secara hukum atau berhasil menjadi `Owner`. | `PD-O1-01`, **PROPOSED** |
| **O1-R02** | Saat berhasil, orang yang bertindak benar-benar teratribusi ke organisasi dan hubungan `Owner` internal pertama. | Atribusi yang cukup tersedia **pada saat akibat efektif**. | Hubungan `Owner` tidak boleh tak bertuan/tidak teratribusi. | `PD-O1-01/02`, **PROPOSED** |
| **O1-R03** | Akibat pembentukan dan batas kewenangan `Owner` internal dapat dipahami sebelum berlaku. | Sebelum tindakan menghasilkan perubahan material. | Tidak berarti sudah diverifikasi sebagai perwakilan legal. | `PD-O1-03`, **PROPOSED** |
| **O1-R04** | Keberhasilan hanya diklaim jika organisasi **dan** `Owner` internal pertama sama-sama sah/efektif. | Kedua fakta Product benar-benar terjadi. | Satu keadaan yang baru sebagian tidak boleh diklaim sebagai sukses penuh. | `PD-O1-02`, **PROPOSED** |
| **O1-R05** | Organisasi dan hubungan `Owner` saat ini dapat diperiksa kembali beserta makna internalnya. | Setelah keberhasilan nyata. | Pesan sukses sementara tidak menggantikan keadaan durable. | `PD-O1-03`, **PROPOSED** |
| **O1-R06** | Klaim dari organisasi tidak otomatis meningkat menjadi hasil peninjauan, bukti eksternal, kurasi, atau kategori fakta independen. | Setiap kali klaim organisasi disajikan/ditafsirkan. | Pembentukan O1 bukan peninjauan organisasi atau legitimasi legal eksternal. | `PD-O1-05`, prinsip Product Intent + aplikasi O1 **PROPOSED** |
| **O1-R07** | Kegagalan, keadaan parsial, hasil tidak teratribusi, atau interupsi dibedakan dari keberhasilan asli. | Ketika keadaan negatif atau belum pasti terjadi. | Retry/pemulihan bukan kewajiban sistem Product baru kecuali diputuskan kemudian; tidak boleh membuat sukses palsu. | `PD-O1-04`, **PROPOSED** |
| **O1-R08** | **Potential Conflicts** yang material bagi klaim kewenangan awal, benar-benar diketahui dan belum selesai tidak menghasilkan klaim kewenangan **seolah tanpa konflik**. | Hanya jika ada **sinyal material yang benar-benar diketahui**; asal/tersedianya sinyal belum diverifikasi. | Tidak mengharuskan deteksi atau penyelesaian konflik universal. | `PD-O1-04`, **PROPOSED** |

**Invariants yang diusulkan:** Tidak ada `Owner` internal awal yang tidak dapat diatribusikan; tidak ada keberhasilan penuh dari pembentukan yang parsial; tidak ada klaim kewenangan awal seolah tanpa konflik ketika Potential Conflicts material yang belum selesai benar-benar diketahui; tidak ada label telah ditinjau/terverifikasi hanya karena O1.

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
| **O1-AC08** | `R08 / SC-O1-03` | Saat **Potential Conflicts** material terhadap klaim kewenangan awal **benar-benar diketahui** dan belum selesai, sistem tidak menyatakan kewenangan awal seolah tanpa konflik atau berpura-pura konflik telah tuntas. | Klaim `Owner` tanpa konflik yang tidak didukung keadaan nyata. |
| **O1-AC09** | `R06 / SC-O1-04` | O1 selesai tidak otomatis menghasilkan persiapan Campaign, penayangan publik, izin menerima donasi, atau peninjauan organisasi. | Klaim side effect lintas-Journey yang tidak termasuk O1. |

Acceptance ini adalah konsekuensi Product yang bisa diperiksa bila nanti diimplementasikan, **bukan** resep database transaction, daftar tes API, atau bukti pengujian yang sudah dijalankan.

## 5. Dependencies, Cakupan, dan Kebebasan Solusi

| Hal | Keadaan untuk O1 |
| --- | --- |
| **Prasyarat Product** | Identitas/atribusi orang yang sah saat hasil efektif dan kemampuan menyatakan organisasi serta hubungan `Owner` internal yang nyata. **Ketersediaannya dalam implementasi sekarang belum diverifikasi.** |
| **Sinyal penting lintas concern** | Akses pada sinyal Potential Conflicts yang **benar-benar diketahui dan material bagi klaim kewenangan awal**, bila ada; sumber dan keterjangkauannya belum diverifikasi. Juga butuh aturan identity/security proyek yang berlaku; bukan diciptakan oleh Product Contract. |
| **Dalam cakupan** | Pembentukan konteks organisasi teratribusi, hubungan `Owner` internal pertama, pemahaman sebelum akibat, keberhasilan gabungan, pemeriksaan kembali, dan hasil negatif/unknown yang material. |
| **Di luar cakupan** | Matriks permission lengkap `Owner`/`Staff`; undangan/perpindahan/delegasi staf; pembuktian representasi legal eksternal; peninjauan organisasi menyeluruh; Campaign; publikasi; kelayakan donasi; dana dan pertanggungjawaban; deteksi konflik/duplikasi universal; penyelesaian konflik formal; subsystem pemulihan wajib. |
| **Kebebasan Design** | Formulir dan isinya sesuai pengenalan organisasi yang bermakna, penyampaian konsekuensi, alur identitas/sign-in, interaksi, copy, tampilan gagal/belum pasti serta penyampaian Potential Conflicts yang benar-benar diketahui. |
| **Kebebasan Engineering** | Mekanisme autentikasi, penyimpanan/data model, konsistensi/transaksi, API/komponen, pemulihan bila dibutuhkan, teknik pengujian; tanpa melemahkan makna dan safeguard yang nanti disetujui. |
| **I2/I3 yang belum diuji** | Sumber Design, arsitektur, keamanan, protected paths, repository/ref/kode runtime dan consumer expectations harus disuplai dan diverifikasi terpisah saat assignment. Tidak ada klaim kesiapan Engineering dari draf ini. |

**Pembatas pemaknaan:** Keberadaan node O1 di peta yang disetujui tidak membuktikan kemampuan runtime atau menjadi persyaratan bahwa seluruh J-ORG harus selesai sebelum Campaign bisa disiapkan.

## 6. Open Decisions & Blockers

Status `OPEN` di sini tidak otomatis berarti mendesak secara keseluruhan Product. Namun keputusan yang material terhadap **kontrak O1 ini** harus diselesaikan sebelum kontrak dapat menjadi `I1` yang mengikat.

| Pertanyaan / konsekuensi yang belum disetujui | Pemilik | Status | Bagian terdampak / syarat pengesahan |
| --- | --- | --- | --- |
| **Paket kebijakan pembentukan O1**: siapakah aktor yang dapat memulai dan memperoleh `Owner` internal pertama, serta arti kewenangannya? | Anhar | **PD-O1-01 PROPOSED** | `SC-O1-01, R01/R02, AC01/AC02`; keputusan O1 mandiri diperlukan. |
| **Satu hasil gabungan dan waktu atribusi**: organisasi dan `Owner` wajib terbentuk bersama pada saat efek berhasil? | Anhar | **PD-O1-02 PROPOSED** | `SC-O1-01, R02/R04, AC02/AC04`; belum ada kebijakan final. |
| **Kejelasan akibat dan pemeriksaan durable**: apa klaim yang wajib dapat dipahami/diperiksa? | Anhar | **PD-O1-03 PROPOSED** | `SC-O1-02, R03/R05, AC03/AC05`; bentuk UI/penyimpanan didelegasikan. |
| **Kegagalan dan Potential Conflicts yang benar-benar diketahui**: apakah klaim terdampak ditunda atau ditolak, dan bagaimana hubungan `Owner` efektif diperlakukan? | Anhar | **PD-O1-04 PROPOSED; akibat konflik terhadap hubungan efektif perlu dipertegas** | `SC-O1-03, R07/R08, AC07/AC08`; jangan mengubah rekomendasi ketat sebelumnya menjadi aturan tanpa keputusan. |
| **Pemisahan informasi organisasi dan hasil verifikasi**: apakah aplikasi ketentuan umum ini pada O1 diterima? | Anhar | **PD-O1-05 dalam paket PROPOSED**; pembedaan umum sudah didukung Product Intent | `SC-O1-04, R06, AC06/AC09`. |
| **Sinyal Potential Conflicts yang tersedia, identity authority, persistence, dan constraint keamanan nyata** | Pemilik teknis/Design terkait pada assignment yang sah | **UNVERIFIED — bukan keputusan Product baru secara otomatis** | `I2/I3` harus dibuktikan sesuai scope; bila tidak tersedia, hentikan hanya klaim/pekerjaan terdampak. |

**Satu keputusan Product terpadu yang direkomendasikan untuk review:** bahas `PD-O1-01..05` sebagai paket yang koheren, termasuk konsekuensi material Potential Conflicts; jangan melakukan approval parsial diam-diam hanya karena tabel terlihat lengkap.

## 7. Kesiapan Konsumsi dan Batas Handoff

| Pertanyaan kesiapan | Keadaan yang benar-benar diketahui |
| --- | --- |
| **Exact Product authority dan approval** | Product Intent dan Level 2 v0.3 disetujui sebagai sumber; **PD-O1-01..05, SC, requirements dan AC masih PROPOSED**, tidak ada approval exact revision kontrak O1. |
| **Konsistensi outcome → SC → rules → AC** | Jejak antar-ID dirumuskan dalam draf ini dari proposal v0.1 untuk review. Ini **pemeriksaan penulisan berbasis dokumen**, belum independent fresh-reader validation. |
| **Blocker Product tersembunyi** | Model pemberian `Owner`, keberhasilan gabungan, batas sebelum akibat, dan penanganan Potential Conflicts masih memerlukan keputusan Anhar; tidak boleh dianggap telah diselesaikan oleh template. |
| **Syarat Product dan ketersediaan teknis** | Atribusi dan persistensi dibutuhkan pada akibat yang diusulkan; status kode/lingkungan, ketersediaan sinyal Potential Conflicts dan mekanisme autentikasi **UNVERIFIED**. |
| **Rancang `I1` generic fit** | Struktur memuat outcome, aturan, acceptance, source revisions, exclusions dan delegated freedom. **Candidate only; belum diuji oleh consumer Rancang independen.** |
| **Rancang project-bound fit** | **NOT CHECKED / BLOCKED for binding use** sampai Product approval exact revision dan pemeriksaan I2/I3/I4 secara terpisah, plus readiness Rancang. |
| **Hasil handoff** | **DRAFT — NOT I1 READY.** Tidak ada izin menjalankan Rancang, tidak ada perubahan operational `main`, dan tidak ada bukti perilaku sudah delivered. |

**Tindakan Product berikutnya:** Rekonsiliasi lima usulan `PD-O1-01..05` terhadap Product Intent dan Journey, putuskan hanya pertanyaan O1 yang material, lalu ajukan **revisi artefak tepat** untuk approval Anhar. Sesudahnya baru dapat dilakukan pemeriksaan calon `I1` consumer fit, terpisah dari I2/I3/I4.

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

**Status saat ini tetap DRAFT / PROPOSED.** Catatan ini tidak mempercepat approval `PD-O1-01..05` atau menjadikan O1 input `I1` yang mengikat.
