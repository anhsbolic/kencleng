# Code Review Findings — CR-TOP-001

> Phase: Code Review  
> WORK_UNIT_ID: WU-S1-005  
> RUN_ID: CR-TOP-001  
> Author: Codex CLI agent (Reviewer; Independent Code Review — Root topology Caddy/MinIO policy)  
> Created: 2026-09-23  
> Model / Reasoning / Session: gpt-5.6-terra / medium / Fresh independent Code Review session  
> Target revision: 1b767c94732f95649299e33e8f4478c2c781f690  
> Workflow revision: not exposed by the invocation

Review dilakukan terhadap perubahan source Build pada commit `c25a6959aee97c0d32994a01924da6fb5c04be40` (`Caddyfile`, `docker-compose.yml`, dan `docs/project/kencleng-repo-setup.md`). Ketiga file tersebut tidak berubah lagi sampai target revision di atas.

## 1. Safety

### S1 — Kegagalan policy public dapat tertutup oleh policy private yang berhasil

- **Lokasi:** `docker-compose.yml:30-36`, entrypoint `minio-init`.
- **Masalah:** Shell menjalankan perintah yang dipisahkan `;` tanpa `set -e`/`-e`. Setelah perubahan ini, `mc anonymous set download local/kencleng-public` dapat gagal tetapi `mc anonymous set none local/kencleng-private` berikutnya berhasil; exit status container menjadi sukses. Ini menghilangkan sinyal gagal yang seharusnya memicu retry/diagnosis dan dapat meninggalkan `kencleng-public` tanpa policy `download` yang diwajibkan R3.
- **Mengapa penting:** R3 menuntut konvergensi dua policy bucket pada bootstrap maupun volume reuse. Initializer yang sukses palsu berarti keadaan policy aktual tidak dapat dipercaya dan perubahan private baru secara khusus menutupi kegagalan policy public sebelumnya.
- **Resolusi yang disarankan:** Jalankan skrip dengan fail-fast, misalnya ganti `/bin/sh -c` menjadi `/bin/sh -ec` (atau tambahkan `set -e;` di awal skrip), sehingga setiap kegagalan `mc mb` atau `mc anonymous set` membuat `minio-init` gagal. Perilaku retry pada `until mc alias set ...` tetap valid karena perintah kondisi `until` dikecualikan dari exit langsung `-e`.
- **Status:** **Blocking.**

Tidak ada hazard concurrency, nullable state, atau resource-lifecycle lain pada diff konfigurasi ini. Safety concern akses anonymous/private sudah ada di Test Focus Pointer; tidak ada Techplan drift baru.

## 2. Quality

Tidak ada Finding tambahan. Perubahan Caddy satu-baris menggunakan primitive lokal yang tepat (`handle_path`) dan dokumentasi membedakan invariant source dari bukti runtime/deferred. S1 juga merupakan masalah reliability/observability initializer dan ditangani lewat patch plan, bukan diduplikasi sebagai Finding Quality.

## 3. Stack-Specific Best Practices

Sumber yang dirutekan: `../harscode-workspace/best-practices/infra/security-headers-and-tls.md` (trigger: Caddy/reverse proxy; Security Concern Map: network boundary).

Tidak ada Finding baru yang dapat diatribusikan ke diff. `Caddyfile` memang belum memasang CSP, `X-Content-Type-Options`, `Referrer-Policy`, maupun HSTS sebagaimana checklist sumber tersebut, tetapi kondisi ini sudah ada sebelum perubahan dan `TP-TOP-001` §2/D5/RISK-6 secara eksplisit menempatkan hardening header/TLS di luar Work Unit ini. Tidak tepat memperluas patch plan topology sempit untuk remediasi tersebut; residual risk tetap perlu ditangani oleh Work Unit terpisah bila diprioritaskan manusia.

`handle_path /api/*` sendiri mempertahankan boundary same-origin yang diperlukan kontrak `/api` dan tidak menambah redirect/object URL atau cache override.

## 4. Consistency

Tidak ada Finding tambahan.

- `Caddyfile:6-11` memenuhi `TP-TOP-001` R1/R2 dan `api/openapi/index.yaml:39-41`: browser memakai base `/api`, sedangkan backend native menerima remainder path dan fallback tetap frontend.
- `docker-compose.yml:32-35` mempertahankan bucket, volume, port, dan public-policy scope, lalu menegaskan private anonymous `none`, sesuai `TP-TOP-001` R3/R5 dan `docs/spec/4-campaign/features/03-campaign-media.md:11-14`.
- `docs/project/kencleng-repo-setup.md:217-247` mengikuti authority root `AGENTS.md` §1/§4/§7: dokumentasi topology tidak mengklaim endpoint Campaign atau runtime/integrated verification yang belum terjadi.
- S1 justru perlu diperbaiki agar implementasi memenuhi R3 secara andal.

## Verification executed during Review

- `git status --short`, source history, dan `git diff c25a695^ c25a695 -- Caddyfile docker-compose.yml docs/project/kencleng-repo-setup.md` — memastikan actual Build scope tiga file dan tidak ada perubahan source scope setelah Build.
- `git diff --check` pada current worktree — lulus. Pemeriksaan diff historis seluruh commit Build juga menemukan trailing whitespace hanya pada Build report, di luar current diff scope yang ditugaskan; tidak menjadi Finding production pada Run ini.
- `sh -c 'false; true'` menghasilkan exit status `0`; `sh -ec 'false; true'` menghasilkan `1` — reproduksi terarah untuk S1.
- `podman-compose -f docker-compose.yml config` dicoba untuk rendered-Compose check sesuai runtime lokal yang ditetapkan, tetapi tidak dapat berjalan di sandbox ini: Podman gagal mengatur `/run/user/1000/libpod` karena filesystem read-only, lalu `podman-compose` melaporkan Podman tidak tersedia. Ini bukan bukti runtime unavailable pada environment Human dan tidak menggantikan Testing-owned runtime evidence.

## Verdict

**Request changes.**

Blocking Finding: **S1**. Terapkan patch plan yang menyebabkan `minio-init` gagal ketika salah satu operasi policy/bucket gagal, lalu kembali ke Review confirmation terarah atau lakukan validasi sebagai bagian dari fase patch yang berwenang.

## Phase handoff

- **Completed:** empat pass independent Code Review terhadap current implementation/diff scope.
- **Artifacts:** `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/CR-TOP-001/review-findings-1.md`; `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/CR-TOP-001/patch-plan-1.md`.
- **Human decision:** none.
- **Open / deferred:** evidence runtime Caddy/MinIO persisted serta R4-R7 tetap Testing-owned dan memerlukan Compose-capable environment + capability/fixture WU-S1-003. Hardening security headers/TLS tetap outside scope dengan residual risk yang sudah dicatat TP-TOP-001.
- **Recommended next step:** Build/Patch menggunakan patch plan spesifik S1, kemudian targeted confirmation; jangan mulai Testing dalam Run ini.
- **Session transition:** kembali ke healthy Build/Patch session yang mengerjakan BLD-TOP-001, karena patch satu-baris, tidak memperluas arsitektur/kontrak, dan patch plan sudah mengunci perubahan serta verification question.
- **Context pointers:** `TP-TOP-001` R3/R5 dan §12; S1 di dokumen ini; `docker-compose.yml:30-36`.
