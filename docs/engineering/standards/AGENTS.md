# AGENTS.md — docs/engineering/standards/

Scope: `docs/engineering/standards/` only. Follow root `AGENTS.md` and read [Engineering Standards index](README.md) before writing here.

- **Tulis hanya project-wide Engineering decisions yang memang sudah diotorisasi** atau koreksi non-semantik yang dapat dibuktikan. Discussion proposal, research, historical implementation, dan scoped Rancang output bukan authority otomatis.
- **Satu owner per concern.** Backend/frontend stack tetap stack-specific; identity, authorization, data/privacy, API, dan testing lintas-komponen mendapat owning document tersendiri bila telah diputuskan. Gunakan link untuk cross-reference, jangan salin ulang rule normatif.
- **Pertahankan batas authority.** Product meaning dan feature-level commitments tetap di Product/scoped artifacts; API operation detail dan runtime implementation tidak diciptakan dari Engineering Standards.
- **Bedakan status.** Tulis secara eksplisit settled/persisted, proposed, not implemented, atau not verified. Approved design policy bukan bukti bahwa mekanismenya sudah bekerja.
- **Change discipline.** Jangan merevisi settled policy tanpa Human-approved material decision atau material new evidence. Perubahan besar memerlukan review impact ke konsumennya; perbarui katalog, links, dan routing bila ownership berubah.
- **Documentation style.** Ikuti [T09 — Documentation & Language](documentation-conventions.md) untuk dokumen baru/yang disunting. Bahasa Indonesia adalah default human-facing; pertahankan technical terms dan identifiers ketika lebih presisi. Jangan melakukan mass translation atau mengubah approved semantics lewat editorial cleanup.

File ini hanya mengatur perilaku Agent **ketika bekerja dalam direktori ini**. Agent yang mengedit `backend/` atau `frontend/` tetap mengikuti scoped `AGENTS.md` miliknya dan rujukan standards pada root routing.
