# Kencleng — Engineering Standards

> **Authority:** Project-wide Engineering constraints yang sudah disetujui Human untuk branch `pilot/station5-c1-space-only`. Ini tidak mempromosikan standar ke `main`.
> **Scope:** Reusable lintas-feature, termasuk standar yang secara khusus berlaku pada backend atau frontend. Jangan menjadikan feature-level Rancang proposal sebagai standard secara otomatis.

## Current standards

| Concern | Owning document | Standing |
| --- | --- | --- |
| T01 — Backend architecture, Modular Monolith, REST API dan Worker boundaries, Engineering decision authority | [Backend stack & architecture](backend.md) | **Settled / persisted** |
| Frontend technology dan state architecture | [Frontend stack & architecture](frontend.md) | **Existing active baseline** |
| T02 — Login methods, Kencleng-owned account identity, multi-method readiness | [Identity & Authentication](identity-authentication.md) | **Settled / persisted** |

T03–T09 masih dalam pembahasan atau belum diputuskan; tidak ada standar baru yang berlaku hanya karena suatu topik tercantum pada agenda. Tambahkan owning document jika keputusan nyata membutuhkan ownership tersendiri, bukan untuk melengkapi template folder.

## Batas authority dan konsumsi

- **Product meaning dan kapan login diperlukan:** [Product Authority](../../product/README.md). Engineering Standards tidak menambahkan syarat Product yang belum disetujui.
- **Product Design/UX reusable:** [Product Design Standards](../../ui-ux/README.md).
- **Kontrak API spesifik:** [API tooling dan source contract](../../../api/README.md); baseline OpenAPI masih kosong sampai ada authorized delivery requirement.
- **Feature-level Engineering output:** `docs/engineering/` (misalnya Rancang ORG-01) adalah scoped proposal/evidence. Keputusan Agent di sana tidak otomatis berlaku lintas-feature.
- **Executable code, configuration, dan current database state:** sumber evidence implementasi, bukan pengganti kebijakan atau Product authority. Hasil verifikasi harus dibedakan dari proposal.
- **Agent instructions:** root serta stack-scoped `AGENTS.md` mengarahkan implementer ke dokumen ini; [scoped AGENTS](AGENTS.md) berlaku untuk penulis yang mengubah direktori standards ini.

Gunakan owning document per concern. Jika ada gap/konflik material, eskalasikan ke owner yang tepat; jangan membuat aturan baru dengan menduplikasi kebijakan dari domain lain.
