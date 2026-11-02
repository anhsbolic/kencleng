# Template — Kickoff Hi-Fi UI/UX Design (Static HTML)

> **Standing:** Reusable **optional session prompt**, not an approved feature design, new project Design Standard, reusable Space formation, or execution authorization.
> **Owner:** Kencleng Product Design guidance (`docs/design/`).
> **Use for:** Bounded feature-specific visual/interaction design when current approved design is not sufficient for implementation.
> **Method:** ChatGPT design discussion + hi-fi HTML/CSS prototype + rendered review + explicit Human design acceptance.

## How to use

1. Copy **only** the prompt between the delimiters below into a new ChatGPT session in Project Kencleng.
2. Fill `{{OUTCOME_ID}}`, `{{TARGET_BRANCH}}`, `{{INITIAL_SURFACE}}`, `{{SCOPE_NOTES}}` and `{{PERSISTENCE_AUTHORIZATION}}`. Do not assume example commit SHAs remain current.
3. Let the new session discover effective Product, Design and accepted Engineering sources from the current repository. Give the session repository access explicitly if needed.
4. Start with a single representative screen and material states. Extend scope after inspecting actual rendered evidence.
5. Decide whether to **persist a Design Candidate** and, separately, whether to **commit/push**. The template alone grants neither. Human approval of the exact design revision is a later decision.

---

## BEGIN COPY-PASTE KICKOFF PROMPT

Bro, kita mulai sesi **Product UI/UX Design — `{{OUTCOME_ID}}`** dalam Project Kencleng.

### Objective

Hasilkan **feature-specific high-fidelity UI/UX design** yang kuat secara visual, implementable, accessible dan benar secara Product meaning; **static HTML/CSS** menjadi primary prototype medium. Fokus pada *design quality and experience*, bukan engineering production.

Kita sedang menyiapkan **Design Candidate** untuk direview Human. Ini bukan Next.js implementation, bukan Rancang, dan bukan authorization untuk mengubah canonical Product atau Engineering semantics.

### Current context dan authority discovery

- Project repository: `anhsbolic/kencleng`.
- Target branch: `{{TARGET_BRANCH}}`. Verifikasi actual HEAD dan current authority sebelum mengubah file; jangan menebak checkpoint.
- Outcome/work identifier: `{{OUTCOME_ID}}`.
- Initial design surface: `{{INITIAL_SURFACE}}`.
- Scope/exclusions/constraints tambahan: `{{SCOPE_NOTES}}`.

Baca hanya authoritative originals yang relevan:

1. **Product:** accepted scoped Product Contract, approval/status, actor journey dan critical acceptance for `{{OUTCOME_ID}}`; temukan melalui current `docs/product/` routing. Product owns *what and why*, including legal/security/privacy claims and consequences.
2. **Kencleng Design:** `docs/design/README.md`, `product-design-principles.md`, `brand-product-ui-brief.md`, `design-guidelines.md`, `patterns.md`, `asset-governance.md` dan relevant approved `visual-references/`. **Lihat actual visual reference images**, bukan sekadar membaca deskripsi; bila tool tidak bisa membukanya, laporkan limitation dan minta gambar dari Human bila diperlukan.
3. **Accepted Engineering/Rancang:** jika sudah tersedia dan relevan, temukan exact accepted handoff dan baca experience obligations, material interface/state/error/unknown/auth constraints. Ini membatasi desain agar truthful dan implementable; bukan undangan mendesain API baru.
4. **Existing frontend:** boleh dilihat sebagai *implementation/evidence of current state*, bukan visual Design Authority atau source of new Product semantics.
5. **Harscode Design research guidance** (jika tersedia): `anhsbolic/harscode-workspace`, current research `design/README.md`, `design/exploration.md`, `design/authority-and-handoff.md`. Ini supporting collaboration guidance, bukan authority di atas current Kencleng Design.

Distinguish **approved/current**, **working/proposed**, **visual reference/evidence**, dan **unverified**. Jika satu material Product decision tidak tersedia, jangan mengarangnya. Jangan memperluas scope ke fitur lain secara otomatis.

### Quality expectation

Terjemahkan approved **Sunlit Editorial / Evidence-Led Optimism** sebagaimana berlaku untuk surface ini. Jangan hanya menyalin warna kuning, font, dan rounded cards. Perhatikan:

- Clear hierarchy, purposeful editorial composition, visual rhythm, typography roles, warm restrained palette, action emphasis, truthful copy, accessible controls, and realistic content.
- Perbedaan ekspresi public vs authenticated/operational surfaces. Visual references memberi *character and relationships*, bukan pixel-perfect geometry atau permission memakai data/figur ilustratif sebagai product truth.
- Relevant loading, empty, error, denied, pending, uncertain, recovery dan consequential confirmation states **hanya bila relevan** terhadap current surface.
- Responsive intent: desktop dan narrow mobile; jangan menghilangkan material information/action di mobile.
- Assets/iconography sesuai current approved project authority; jangan membuat arbitrary logo, trust badge, documentary image atau verification claim. Synthetic demo content harus terlihat sebagai demo, bukan bukti operasional.

Pilih satu desain yang kuat berdasarkan authority. Explore alternatif hanya jika ada trade-off yang betul-betul material.

### Prototype and feedback

- Deliver **single bounded hi-fi static HTML/CSS prototype** untuk `{{INITIAL_SURFACE}}` terlebih dahulu, dengan ringan JavaScript hanya untuk demo state/navigation bila perlu.
- Tidak memakai Next.js production structure, API calls, backend, database, OAuth, MSW, cookies atau real user data.
- Buat prototype **bisa dibuka secara lokal** dan tidak menggantung pada layanan remote tanpa kebutuhan; jika fonts/assets external dipakai, beri fallback dan jelaskan dependencies/provenance.
- **Render dan lihat** hasil aktual jika tool mendukung; inspeksi composition, typography, overflow, realistic content, responsive and meaningful UX states. Jangan claim visual verification jika belum melihat rendering.
- Tunjukkan hasil dan kritik visual yang spesifik terhadap Kencleng design authority, lalu iterasi bersama Human sebelum memperluas ke surface lain.
- Hi-fi pada **material representative states**, cukup fidelity pada sisanya. Tidak perlu ratusan screenshot atau semua layar sekaligus.

### Candidate persistence and handoff

**Authorization supplied for this particular session:** `{{PERSISTENCE_AUTHORIZATION}}`

Jika Human memberi izin repository write/commit, proposed project-owned location:

`docs/design/outcomes/{{OUTCOME_ID_LOWER}}/`

Suggested minimal deliverables (use only if useful):
- `prototype.html`: runnable static design prototype.
- `handoff.md`: scope, original authority/revisions, exact prototype identity, material design/interaction decisions, representative states/responsive intent, evidence/limitations, open Human decisions and freedom for NextJS Builder.
- `assets/` and/or `screenshots/` only when they add real evidence/value.

Ini **proposed path, bukan mandatory layout**. Periksa existing project location/convention before creating it. Jangan menyalin Product Contract atau Rancang API sebagai duplicate design authority. Jangan menulis ke `frontend/`, `api/`, shared Product/Rancang sources, atau project-wide Design Standards dalam assignment ini.

Semua output bermula sebagai **DESIGN CANDIDATE / NOT HUMAN APPROVED**, sekalipun prototipe sudah dirender atau committed. Human akan menentukan acceptance terhadap **exact design revision** secara terpisah. Tanpa explicit Human write permission, hasilkan prototype di chat/attachment dan usulkan persist action; jangan mengklaim telah menulis repo.

### Working style and first action

Gunakan **Challenge first, then commit**. Tetap aktif sebagai UI/UX Designer; jangan sekadar menerjemahkan source ke daftar requirement. Pertanyakan masalah visual/interaction material, tetapi jangan reopen accepted Product/Rancang decisions tanpa new evidence.

Mulai dengan current authoritative source discovery yang proporsional, lalu **langsung buat dan render hi-fi initial prototype untuk `{{INITIAL_SURFACE}}`**. Jangan berhenti pada read-only assessment atau planning ceremony jika input cukup. Tutup dengan apa yang benar-benar terlihat/teruji, keputusan yang masih diminta dari Human, dan tindakan berikutnya.

## END COPY-PASTE KICKOFF PROMPT

---

## Example filling — ORG-01 (illustrative; verify live state)

| Placeholder | Example |
| --- | --- |
| `{{OUTCOME_ID}}` | `ORG-01` |
| `{{OUTCOME_ID_LOWER}}` | `org-01` |
| `{{TARGET_BRANCH}}` | `pilot/station5-c1-space-only` |
| `{{INITIAL_SURFACE}}` | `Organization Onboarding / Entry Experience` |
| `{{SCOPE_NOTES}}` | `Mulai dari public entry, pre-Google disclosure dan significant error/unavailable states; jangan sekaligus redesign account deletion/privacy; preserve accepted ORG-01 and Rancang semantics.` |
| `{{PERSISTENCE_AUTHORIZATION}}` | `Contoh: hanya candidate files di docs/design/outcomes/org-01/, termasuk commit/push ke target research branch; tanpa writes ke frontend/api/Product/Rancang.` **Only when Human explicitly grants this scope.** |

For ORG-01 the session can discover:
- `docs/product/outcomes/org-01/product-contract.md` and `approval.md`.
- `harscode-state/org-01/rancang-002/handoff.md`, `experience.md`, and material `interface-decisions.md`.
- `docs/design/visual-references/selected-direction/` (public and evidence-journal PNG references, **not production assets**).
- Existing implementation under `frontend/features/org-01/` as observation only.

Do **not** turn the example branch/paths/authority standing into universal conventions or silently grant write permission through this file.
