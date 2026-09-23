# Review findings — WU-S1-005 / TPR-TOP-001

> Phase             : Techplan independent review
> Author            : Codex CLI agent (Reviewer)
> Created           : 2026-09-23
> Model             : gpt-5.6-terra
> Reasoning         : medium
> Session           : Fresh independent session
> Target revision   : 9b5dd8bfbef3c0a8815624f0355347ca00ef9f03

**Gate:** Complex — melintasi kontrak proxy/API same-origin, policy object storage persisten, dan batas ownership topology/backend; juga menyentuh batas keamanan media privat dan retraction.

**Sections resolved:** Background §1; Scope §2; Requirements §3; Rules & Validation §4; Decision Log §5; Backward Compatibility §6; Edge Cases & Risks §7; Interface Contract §8; Architecture / Plan §9; Implementation Details §10; Files Changed / Files NOT Changed §11; Testing Checklist dan Test Focus Pointer §12; Open Items §13.

### Blocking

- **MATERIAL / BLOCKING — Verification / ownership:** §12, row R4 dan R5; Test Focus Pointer baris “Private bucket / direct object bypass / retraction”; serta §13 Active item 2 menyatakan retraction tetap relevan dan perlu evidence bersama, tetapi tidak ada checklist yang menguji fresh fetch melalui `localhost:8080` setelah parent eligibility atau membership ditarik. R4 hanya menguji 200/404/503 dan header; R5 hanya menguji anonymous direct-object denial serta ketiadaan URL/redirect. Ini meninggalkan bukti yang dibutuhkan untuk memastikan `private, no-store` dan proxy benar-benar mendukung withdrawal pada URL yang telah diketahui. `EXP-TOP-001/evidence/solutioning.md#runtime-verification-required-not-performed-here` item 6 mewajibkan skenario tersebut, dan `WU-S1-003/runs/TP-BE-001/techplan.md` §7 RISK-3 serta §13 Active item 1 secara eksplisit mendelegasikan bukti policy/proxy/**retraction** kepada WU-S1-005. Planning resolution harus menambahkan rule/checklist terkoordinasi dengan WU-S1-003 yang menarik eligibility atau membership lalu melakukan fresh request pada URL yang sama melalui Caddy, beserta owner dan batas klaim bila fixture belum tersedia.

### Non-blocking

- Tidak ada.

### Clean

- Keputusan `handle_path /api/*`, penegasan `mc anonymous set none` hanya pada `kencleng-private`, dan batas ownership WU-S1-003/WU-S1-004 setia pada EXP-TOP-001 dan otoritas Slice 1.
- Spot-check source hidup mengonfirmasi mismatch saat ini (`Caddyfile` meneruskan `/api/*`, sedangkan `backend/cmd/server/main.go` mendaftarkan `/healthz` tanpa prefix), dua bucket serta volume persisten di `docker-compose.yml`, dan kontrak `/api`/`content_url`/`Cache-Control: private, no-store` di sumber OpenAPI.
- Setiap R1–R6 memiliki setidaknya satu baris Testing Checklist; tidak ada diagram untuk divalidasi. Active Open Items dengan benar non-blocking untuk narrow source Build, tetapi memblokir klaim runtime/integrasi sampai evidence tersedia.
- Tidak ada perubahan pada source topology/contract yang direview sejak target revision Draft `a053b48ef3fef35a073c5fe57e5ee4581f848e07`; perubahan menuju revision review hanya artefak orkestrasi.

## Phase handoff

- Completed: independent review gate + review ketika Complex warranted.
- Artifacts: `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/TPR-TOP-001/review-findings.md`.
- Human decision: revise Draft Techplan untuk menutup Finding material, lalu jalankan human gate setelah resolution yang diperlukan.
- Open / deferred: Finding retraction-through-proxy di atas; runtime Compose/Caddy/MinIO dan fixture WU-S1-003 tetap deferred sesuai §13 setelah kewajiban test direkam.
- Recommended next step: satu resolution pass pada Techplan untuk menambahkan evidence retraction terkoordinasi, kemudian human gate; re-review diperlukan bila resolution mengubah scope, ownership, kontrak, atau strategi verifikasi secara material.
- Session transition: fresh resolution/human-gate action; Build tetap fresh-preferred setelah Approval.
- Context pointers: `TP-TOP-001/techplan.md` §12/§13; `EXP-TOP-001/evidence/solutioning.md#runtime-verification-required-not-performed-here`; `WU-S1-003/runs/TP-BE-001/techplan.md` §7/§13.
