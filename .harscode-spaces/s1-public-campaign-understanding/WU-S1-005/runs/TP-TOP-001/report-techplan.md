# Laporan: Slice 1 Topology & Controlled Media Enablement

> Fase              : Techplan human review report
> Tiket             : WU-S1-005 / TP-TOP-001
> Status            : Draft
> Penulis sumber    : Planner
> Dihasilkan oleh   : Codex CLI agent
> Model             : gpt-5.6-terra
> Reasoning         : medium
> Session           : Fresh focused session
> Dihasilkan        : 2026-09-23
> Revisi target     : a053b48ef3fef35a073c5fe57e5ee4581f848e07
> Revisi workflow   : tidak diekspos oleh invocation
> Sumber            : `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/TP-TOP-001/techplan.md` pada revisi target `a053b48ef3fef35a073c5fe57e5ee4581f848e07`

---

## Apa & mengapa

Work Unit ini menyiapkan boundary topology Slice 1 agar URL API same-origin tetap memakai `/api`, tetapi backend menerima path native tanpa prefix. Ia juga menegaskan bahwa bucket media Campaign bersifat private pada state MinIO yang baru maupun persisten. Hasil yang dituju adalah jalur yang siap dipakai delivery Campaign terkontrol tanpa mengambil alih ownership Campaign tersebut.

## Scope

**Termasuk**

- Mengubah handler API Caddy menjadi `handle_path /api/*`, sehingga `/api` dihapus hanya sebelum request diteruskan ke backend dan URL browser tidak berubah.
- Menegaskan `mc anonymous set none local/kencleng-private` pada setiap `minio-init`, sambil mempertahankan policy download `kencleng-public`.
- Memperbarui dokumentasi setup dari caveat prefix menjadi invariant executable, termasuk batas bukti runtime yang masih diperlukan.
- Menyerahkan `MINIO_BUCKET_PRIVATE` sebagai handoff konfigurasi kepada `WU-S1-003`.

**Secara eksplisit tidak termasuk pada putaran ini**

- Persistence, handler/rute, authorization, streaming, MIME, `Problem Details`, cache-header production, dan retraction implementation Campaign milik `WU-S1-003`.
- UI/consumer Campaign, mock behavior, atau Next.js rewrite milik `WU-S1-004`.
- Perubahan OpenAPI, base URL, `content_url`, CORS, route-prefix backend, redirect/signed URL, response synthesis, atau cache override.
- Penggantian bucket/volume/port, perubahan policy `kencleng-public`, serta hardening security-header/TLS yang lebih luas.

## Arsitektur / Rencana

Request browser `/api/<path>` masuk melalui root Caddy. `handle_path` mengirim `<path>` ke backend native; request non-API tetap menuju fallback frontend. Untuk media Campaign, Caddy hanya meneruskan respons dari endpoint terkontrol backend—ia bukan titik authorization dan tidak membuat redirect, object URL, signed URL, cache override, atau respons pengganti.

`minio-init` tetap membuat kedua bucket, mempertahankan anonymous download pada `kencleng-public`, dan secara konvergen menetapkan anonymous policy `none` pada `kencleng-private`, termasuk ketika `kencleng_miniodata` dipakai ulang.

**Komponen yang disentuh**

| Komponen | Tujuan / perubahan |
|---|---|
| `Caddyfile` | Translasi `/api` ke path backend native sambil mempertahankan fallback frontend. |
| `docker-compose.yml` | Assertion policy anonymous private yang konvergen pada `minio-init`. |
| `docs/project/kencleng-repo-setup.md` | Dokumentasi invariant path dan batas bukti runtime. |

**Batas blast radius:** Tidak ada perubahan pada `backend/**`, `frontend/**`, `api/openapi/**`, generated types, `.env.example`, nama bucket, volume, port, atau policy `kencleng-public`.

## Kontrak antarmuka

| Perhatian | Kontrak |
|---|---|
| URL API publik | Browser tetap menggunakan base `/api`; Caddy meneruskan remainder tanpa `/api` ke backend. |
| Otoritas | Caddy hanya meneruskan; authorization Campaign dan recheck eligibility/membership tetap milik `WU-S1-003`. |
| Media Campaign | `content_url` tetap opaque same-origin reference menuju endpoint terkontrol, bukan URL object storage. |
| Respons media | `Problem Details` dan `Cache-Control` diproduksi backend; Caddy tidak mengubahnya. |

## Keputusan utama

| Keputusan | Alasan / konsekuensi |
|---|---|
| Gunakan `handle_path /api/*`. | Memenuhi kontrak `/api` tanpa mengubah URL browser atau konvensi rute backend. |
| Set anonymous `none` hanya pada `kencleng-private`. | Menutup akses anonymous pada bucket Campaign secara konvergen tanpa mengganggu `kencleng-public`. |
| Gunakan `MINIO_BUCKET_PRIVATE` yang sudah ada. | Menjadi handoff konfigurasi; ownership storage/delivery Campaign tetap di `WU-S1-003`. |
| Backend memproduksi `Problem Details` dan `Cache-Control`; Caddy meneruskan. | Menjaga controlled origin dan menghindari perubahan proxy yang dapat merusak semantik retraction. |
| Security headers/TLS tidak ditambahkan. | Hardening tersebut di luar scope dan tidak diklaim sebagai remediation oleh Work Unit ini. |

## Risiko / trade-off material

| Risiko / trade-off | Paparan / mitigasi | Status |
|---|---|---|
| Prefix tidak diterjemahkan atau API tertangkap fallback frontend. | Validasi source dan runtime `GET localhost:8080/api/healthz`; request non-API juga diuji. | Runtime/testing deferred |
| Volume persisten menyimpan policy private yang salah. | Inspeksi runtime policy kedua bucket pada first boot dan reuse `kencleng_miniodata`. | Runtime/testing deferred |
| Object Campaign bypass origin atau URL yang telah diketahui masih mengirim byte setelah withdrawal. | Private bucket, controlled endpoint, direct-object denial, dan R7 fresh fetch melalui root Caddy setelah withdrawal. | Runtime/testing deferred |
| Proxy menyamarkan `404`/`503` atau membuang `private, no-store`. | Uji end-to-end setelah endpoint/fixture `WU-S1-003` tersedia. | Runtime/testing deferred |
| Source change dianggap bukti runtime/integrasi. | Dokumentasi dan status evidence wajib membedakan keduanya. | Dimitigasi pada source; runtime tetap deferred |

## Riwayat review & resolution

Independent review `TPR-TOP-001` menemukan satu Finding `MATERIAL / BLOCKING`: Techplan belum memiliki bukti executable bahwa `content_url` yang sudah diketahui tidak lagi mengirim byte baru setelah public eligibility parent Campaign atau membership media ditarik melalui proxy root.

Resolution `TPR-RES-TOP-001` menutup Finding secara sempit dengan menambahkan Q8, R7, mitigasi `RISK-3`, langkah rencana, dan checklist R7. R7 menetapkan urutan: fetch eligible sebagai precondition, mutasi persisted oleh `WU-S1-003`, lalu fresh request tanpa cache client ke URL identik melalui `localhost:8080`; hasilnya harus `404` public non-disclosure dengan `Cache-Control: private, no-store`, tanpa byte baru, redirect, object URL, signed URL, atau respons pengganti proxy. Byte yang telah diunduh/dipegang client berada di luar jaminan ini.

Tidak ada perubahan scope, arsitektur, ownership, atau semantik API/security. Konfirmasi terarah `TPR-CONF-TOP-001` memutuskan `CONFIRMED_CLOSED`; tidak ada Finding blocking yang tersisa.

## Keputusan yang diperlukan sebelum approval

Tidak ada.

## Tindak lanjut deferred / milik Human

- Testing/environment operator perlu menjalankan bukti runtime Compose/Caddy/MinIO: validasi konfigurasi Caddy yang dirender dan inspeksi policy persisted pada `kencleng_miniodata`.
- `WU-S1-003` perlu menyediakan endpoint dan fixture Campaign/media eligible/non-eligible serta mutasi persisted eligibility/membership; Testing bersama `WU-S1-005` kemudian menjalankan R4, R5, dan R7.
- R7 retraction-through-proxy tetap **runtime evidence**, bukan pembuktian tingkat source. Sampai capability dan environment tersedia, statusnya `deferred/not tested`; ini tidak menghapus kewajiban R7 dan tidak mengizinkan klaim retraction atau integrasi selesai.

## Batas approval

Approval Human mengotorisasi Build narrow source-boundary sesuai Techplan: perubahan `Caddyfile`, `docker-compose.yml`, dan `docs/project/kencleng-repo-setup.md`, dengan ownership Campaign tetap pada `WU-S1-003`.

Approval tidak mengotorisasi perubahan Campaign/backend/frontend/OpenAPI di luar scope, tidak membuktikan runtime Caddy/Compose/MinIO atau integrasi media, dan tidak mengizinkan klaim bahwa R7 retraction-through-proxy telah verified maupun bahwa retraction/integrasi telah selesai.

## Persetujuan

- [ ] Scope dikonfirmasi
- [ ] Perubahan material pada antarmuka/otoritas dipahami, bila berlaku
- [ ] Risiko/trade-off material dipahami dan diterima/dimitigasi sebagaimana tercatat, bila ada
- [ ] Item keputusan blocking telah diselesaikan, bila ada
- [ ] Tindak lanjut Human/eksternal yang deferred telah disadari ownership-nya, bila ada
- [ ] Batas approval dipahami

---
*Detail eksekusi penuh, rule ID, implementation anchor, ownership/rationale verifikasi, dan riwayat risiko/keputusan lengkap tersedia pada source Techplan. Jika report ini dan Techplan berbeda, Techplan yang berlaku.*
