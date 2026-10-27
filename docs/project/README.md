# Kencleng — Project & Engineering Documentation

`docs/project/` berisi dokumentasi repository/setup dan **current Engineering Standards**. Ini tidak menentukan Product scope, urutan delivery, atau hasil implementasi.

| Kebutuhan | Dokumen yang memiliki authority |
| --- | --- |
| Project-wide dan stack-specific Engineering Standards | [Engineering Standards index](standards/README.md) |
| Repository topology, toolchain, dan local setup | [Repository setup](kencleng-repo-setup.md) |
| Mode host/container dan hot-reload verification | [Local development modes](local-development-modes.md) |
| Optional S3-compatible local object storage | [Local object storage](local-object-storage-options.md) |

**Boundary:** Product berada di [Product index](../product/README.md), reusable UI/UX di [Product Design index](../ui-ux/README.md), dan feature-level Rancang/evidence di `docs/engineering/`. Historical Station 5/pilot entries dalam direktori ini merupakan catatan episode, bukan active engineering architecture.

Dokumen backend/frontend yang dahulu berada langsung di `docs/project/` sudah dipindahkan ke `standards/`. Referensi dengan old path dan pinned blob dalam artefak Rancang lama merupakan **historical provenance** dan tidak ditulis ulang untuk mengklaim Agent pernah membaca standard baru. Konsumen aktif harus mengikuti [current standards index](standards/README.md).
