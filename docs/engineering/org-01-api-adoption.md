# ORG-01 — Keputusan adopsi API kanonis

**Status: HUMAN APPROVED FOR CANONICAL ADOPTION; FILE ADOPTION DEFERRED.** Keputusan Anhar, 2 November 2026. Persetujuan berlaku untuk penggunaan kontrak API ORG-01 r4 sebagai dasar adopsi kanonis proyek. Atas instruksi Human, kontrak belum boleh disalin atau ditulis ke `api/`.

## Basis persetujuan

- Accepted Rancang handoff: [`rancang-002`](../../harscode-state/org-01/rancang-002/handoff.md).
- Sumber OpenAPI r4 yang diterima: [`openapi-handoff.yaml`](../../harscode-state/org-01/rancang-002/openapi-handoff.yaml), SHA-256 `a2d21ee3ca4ea93c7e691a53255ab825b784da00f45baeeb88115e76564df9aa`.
- Acceptance exact r4 mengikat manifest SHA-256 `7e0d408f2e438c79149fd3c00e65e25fb4c62a5539c11e379c979ee12e6cd0ba` pada immutable Git checkpoint `6ba8ebcf1a41373ba9ddea441a3659a591c07cb5`; packaging saat ini mencatat identitasnya terpisah.
- Human menyatakan kontrak siap diimplementasikan pada fase Engineering, menyetujui masuk ke dokumentasi API proyek, dan secara eksplisit meminta agar kontrak belum disimpan ke `api/`.

## Efek dan batas keputusan

1. Adopsi kanonis kontrak ORG-01 r4 telah disetujui, tetapi **belum diterapkan**. `api/openapi/index.yaml` dan generated `api/openapi.yaml` tetap pada clean baseline tanpa operasi ORG-01.
2. Engineering dapat memakai exact accepted handoff r4 pada lokasi assignment di atas sebagai input kontrak untuk pekerjaan ORG-01 yang memiliki scope dan otorisasi tersendiri. Selama belum diadopsi, jangan menyebutnya sebagai API kanonis proyek atau mengklaim validasi melalui tooling `api/`.
3. Persetujuan ini tidak mengubah semantik API, menambah operasi, mengotorisasi perubahan implementasi tertentu, migration, test/runtime action, atau deployment. Instruksi menahan penulisan ke `api/` tetap berlaku sampai Human meminta penerapan itu.
4. Sebelum penulisan kanonis kelak, rekonsiliasi metadata OpenAPI: deskripsi di header sumber r4 masih menyebut Human acceptance pending, sementara handoff/state mencatat exact r4 telah diterima. Ketidaksesuaian status ini tidak mengubah wire semantics, tetapi perlu dibereskan atau dijelaskan tanpa mengklaim metadata tersebut sudah diperbarui.

## Status saat ini

Kontrak tersedia sebagai accepted Engineering Solution Handoff, bukan sebagai project-canonical API. Implementasi dan runtime conformance belum dilakukan. Keputusan ini merekam persetujuan adopsi beserta penundaan penerapannya; keputusan ini tidak mengubah Product approval, privacy release gate, atau batas implementasi yang berlaku.
