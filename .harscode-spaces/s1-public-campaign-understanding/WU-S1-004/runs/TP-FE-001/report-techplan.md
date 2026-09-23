# Laporan: Detail Campaign Publik Slice 1 Frontend Mock-Parallel

> Fase              : laporan review Human Techplan
> Ticket            : `WU-S1-004`
> Status            : Draft
> Penulis sumber    : Codex CLI agent
> Dibuat oleh       : Codex CLI agent
> Model             : `gpt-5.6-terra`
> Penalaran         : medium
> Session           : sesi fokus baru
> Dibuat            : 2026-09-23 13:51:43 WIB
> Revisi target     : `a053b48ef3fef35a073c5fe57e5ee4581f848e07`
> Revisi workflow   : tidak diekspos oleh Run ini
> Sumber             : `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/TP-FE-001/techplan.md` pada revisi target `a053b48ef3fef35a073c5fe57e5ee4581f848e07`

---

## Apa & mengapa

Work Unit ini merencanakan detail Campaign publik untuk Slice 1 agar pengunjung dapat memahami tujuan, pengelola, sumber cerita, tahapan lifecycle, pendanaan, media, dan keadaan aksi berikutnya secara jujur. Hasilnya adalah bukti frontend yang parallel terhadap contract dengan mock pada network boundary; ini memungkinkan klaim `FRONTEND_MOCK_VERIFIED` setelah seluruh bukti dan Human rendered acceptance terpenuhi, tanpa menyatakan integrasi backend nyata sudah terbukti.

## Lingkup

**Termasuk**

- Route publik App Router `frontend/app/campaigns/[campaignId]/` untuk detail satu Campaign.
- Satu request bertipe `GET /api/campaigns/{campaignId}` yang memakai generated OpenAPI types, TanStack Query untuk server state, dan MSW pada browser/Vitest network boundary yang sama.
- State loading, success, media `available`/`absent`/`unavailable`, pendanaan unavailable, `404` yang public-safe, `503` yang dapat dicoba lagi, dan kegagalan jaringan umum.
- Penyajian pendanaan, provenance, media, dan aksi berikutnya yang jujur; plain text organizer yang aman; aksesibilitas, bukti responsive/rendered, Human rendered acceptance, serta dokumentasi mock browser yang dapat direproduksi.

**Secara eksplisit tidak termasuk pada tahap ini**

- Perubahan backend, OpenAPI/generated type, API proxy/rewrite, Caddy/topology, storage, cache, atau enforcement controlled media nyata.
- Donation Flow maupun link/tombol Donate aktif; mutation, polling, optimistic update, perhitungan eligibility, dan detail Campaign Discovery/Organization.
- Fetch Organization tambahan, penggunaan field Campaign internal yang publik, atau state/lifecycle pasca fundraising yang belum direkonsiliasi.
- Shared design-system/component registry, Zustand store, reusable expressive placeholder asset, serta perubahan root layout/global visual system.
- Klaim `INTEGRATED_VERIFIED`; bukti integrasi nyata tetap menjadi rendezvous WU-S1-003, WU-S1-005, lalu WU-S1-006.

## Arsitektur / Rencana

Route server hanya menyediakan identitas Campaign dari URL. Leaf Client Component yang kecil mengelola lifecycle TanStack Query, lalu memanggil satu fungsi API bertipe. Pada browser mock runtime yang diaktifkan secara eksplisit, MSW mulai sebelum query berjalan dan mengintersepsi HTTP path yang sama; tanpa mock, request yang sama menuju jaringan same-origin nyata saat integrasi tersedia. Tampilan route-local memperoleh seluruh state dari hasil query dan discriminant contract, tanpa salinan data pada Zustand, context, atau state yang disinkronkan melalui effect.

**Komponen yang disentuh**

| Komponen | Tujuan / perubahan |
|---|---|
| `frontend/app/campaigns/[campaignId]/` | Shell route, bootstrap mock route-local, lifecycle query, dan presentasi detail. |
| `frontend/lib/api/client.ts` dan `frontend/lib/api/public-campaign.ts` | Satu request boundary tingkat rendah dan consumer generated `PublicCampaignDetail`. |
| `frontend/lib/hooks/use-public-campaign-detail.ts` | Query key stabil dan hook detail per `campaignId`. |
| `frontend/mocks/` | Fixture yang setia pada contract serta handler MSW browser/node untuk detail dan controlled media. |
| `frontend/public/mockServiceWorker.js` | Worker MSW yang digenerasi untuk browser mock runtime. |
| Tests, `vitest.setup.ts`, dan `frontend/README.md` | Bukti yang dapat diamati melalui node MSW, setup lifecycle test, dan instruksi mock browser. |

**Tidak disentuh / batas jangkauan dampak:** Tidak ada perubahan pada `backend/**`, topology/proxy, `api/openapi/**`, `api/openapi.yaml`, atau `frontend/lib/api/generated/openapi.ts`. `frontend/app/layout.tsx`, `frontend/app/globals.css`, shared component areas, Zustand stores, dan sumber authority product/design/spec juga tetap tidak diubah.

## Kontrak antarmuka

| Aspek | Kontrak |
|---|---|
| Endpoint | `GET /api/campaigns/{campaignId}` |
| Autentikasi / otoritas | Operasi baca publik; backend secara eksklusif menentukan public eligibility dan makna `404` tetap non-disclosing. |
| Pemicu / pemanggil | Route publik `frontend/app/campaigns/[campaignId]/` melalui `getPublicCampaignDetail`. |

**Request / masukan**

| Field / masukan | Arti / kebutuhan |
|---|---|
| `campaignId` | Satu route segment yang di-path-encode oleh request boundary; frontend tidak memvalidasi atau mengklasifikasikan eligibility publik. |

**Response / keluaran**

| Field / keluaran | Arti |
|---|---|
| `200` | Hanya generated `components["schemas"]["PublicCampaignDetail"]` yang digunakan untuk detail. |
| `404` | Outcome safe not-found/not-public tanpa alasan atau detail diagnostik. |
| `503` | Outcome temporary-unavailable yang dapat dicoba lagi oleh pengguna. |
| Kegagalan lain/jaringan | Outcome kegagalan umum yang aman dan dapat dicoba lagi oleh pengguna. |
| `PublicCampaignMediaItem.content_url` | Referensi opaque yang dimuat sebagai native image; MSW hanya mensimulasikan JPEG/PNG pada path controlled yang sama. |

**Perilaku error:** Problem Details tidak diparse menjadi copy pengguna; tidak ada `error.message`, stack, header, atau alasan spesifik untuk `404` yang ditampilkan. `404` tidak menyediakan retry, sedangkan `503` dan kegagalan umum menyediakan retry eksplisit.

## Keputusan utama

| Keputusan | Alasan / konsekuensi |
|---|---|
| Dynamic route dengan Client query boundary route-local | Identitas tetap dapat di-bookmark, sementara browser server-state terbatas pada boundary terkecil; shell/layout tetap server-capable. |
| Focused typed API function dan helper tingkat rendah terpusat | Menjaga ownership request/response generated-contract dan satu titik klasifikasi transport tanpa menambahkan semantik auth atau eligibility. |
| MSW hanya pada browser/node network boundary | Browser mock dan test menggunakan HTTP path produksi yang sama; tidak ada cabang mock service pada `lib/api/`. |
| Native `<img src={content_url}>` untuk media tersedia | URL controlled yang disuplai contract tetap opaque dan tidak menciptakan image optimizer/proxy baru. |
| Query key per `campaignId`, tanpa retry/polling otomatis | Tidak ada cache key ad-hoc atau re-request otomatis untuk funding/status; retry hanya melalui tindakan pengguna pada state retryable. |
| Decimal/progress dari API sebagai input tampilan | Tidak ada float conversion, perhitungan ulang, capping, atau eligibility inference di frontend. |
| Treatment media yang tidak tersedia bersifat route-local dan struktural | Tidak menciptakan asset ekspresif/reusable sebelum Human Design decision tersedia. |
| Tidak ada regresi Playwright yang dikomit pada tahap awal | RTL + node MSW melindungi perilaku contract; rendered inspection dan Human acceptance tetap wajib untuk penilaian spasial/produk. |

## Risiko / trade-off material

| Risiko / trade-off | Paparan / mitigasi | Status |
|---|---|---|
| Data publik atau alasan unavailable bocor | Konsumsi dibatasi pada `PublicCampaignDetail`; satu `404` aman dan copy failure tetap; raw error tidak dirender. | mitigated |
| Funding/media dipalsukan atau maknanya runtuh | Decimal hanya ditampilkan, media memakai discriminant dan `content_url` tersuplai; fixture/test mencakup zero, above-target, unavailable, dan absent. | mitigated |
| Teks organizer menjadi markup/XSS atau provenance tampak diverifikasi platform | Render sebagai React plain text dan tampilkan attribution organizer; hostile-looking text diuji. | mitigated |
| Mock dapat disalahartikan sebagai integrasi nyata | Bukti menyatakan batas mock; backend projection, controlled delivery, proxy, storage/retraction, cache, dan timing parity tetap deferred. | open |
| Copy provenance/action serta placeholder dianggap final | Keduanya tetap provisional dan route-local sampai Human Design review menyetujuinya. | open |

## Riwayat review & penyelesaian

Independent review `TPR-FE-001` telah selesai pada 2026-09-23 dan tidak menemukan Finding `MATERIAL / BLOCKING` maupun `MECHANICAL / NON-BLOCKING`. Review mengonfirmasi fidelity R1–R9, D1–D8, spot-check contract/API, batas ownership terhadap WU-S1-003 dan WU-S1-005, lifecycle Open Items, serta Test Focus Pointer. Tidak ada perubahan material pada Techplan dan tidak diperlukan re-review.

## Keputusan yang diperlukan sebelum persetujuan

Tidak ada.

## Tindak lanjut yang ditunda / dimiliki Human

- Human Design harus menyetujui wording Indonesia final untuk provenance dan unavailable action sebelum milestone promotion/delivery. Ini non-blocking untuk mock Build.
- Human Design harus memutuskan reusable campaign placeholder asset/system apabila kebutuhan tersebut muncul; Build hanya boleh memakai treatment route-local yang struktural.
- Human rendered acceptance terhadap route mock desktop dan mobile untuk success/media, not-found, serta retryable unavailable/failure wajib selesai setelah independent Testing sebelum `FRONTEND_MOCK_VERIFIED` dapat diklaim.
- WU-S1-003/WU-S1-005/WU-S1-006 tetap harus membuktikan backend projection, controlled byte delivery, same-origin routing, `private, no-store`, storage/retraction, cache, serta response/timing parity sebelum `INTEGRATED_VERIFIED` atau slice delivery.

## Batas persetujuan

Persetujuan mengotorisasi sesi Build frontend baru untuk melaksanakan Techplan Draft ini sebagai route detail publik mock-parallel: request generated-type-backed yang sama, MSW network boundary, presentasi state yang ditentukan, tests/evidence yang dialokasikan, dan batas scope yang tercatat. Persetujuan tidak mengotorisasi perubahan backend, OpenAPI/generated types, proxy/topology/storage/cache, Donation Flow atau Donate action, perubahan authority product/design, promosi `FRONTEND_MOCK_VERIFIED` tanpa independent Testing dan Human rendered acceptance, maupun klaim `INTEGRATED_VERIFIED` tanpa bukti integration Work Units yang telah ditetapkan.

## Persetujuan

- [ ] Lingkup dikonfirmasi
- [ ] Perubahan antarmuka/authority material dipahami, jika berlaku
- [ ] Risiko/trade-off material dipahami dan diterima/dimitigasi sebagaimana tercatat, jika ada
- [ ] Item keputusan blocking telah diselesaikan, jika ada
- [ ] Tindak lanjut Human/eksternal yang deferred telah disadari kepemilikannya, jika ada
- [ ] Batas persetujuan dipahami

---
*Untuk detail eksekusi lengkap, rule ID, implementation anchor, ownership/rationale verification, serta riwayat risiko/keputusan lengkap, lihat Techplan sumber. Jika laporan ini dan Techplan berbeda, Techplan yang berlaku.*
