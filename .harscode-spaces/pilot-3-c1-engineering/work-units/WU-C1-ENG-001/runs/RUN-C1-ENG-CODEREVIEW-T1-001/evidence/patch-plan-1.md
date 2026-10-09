# Patch plan — T1 shared contract F1–F4

> Phase: Independent Code Review; Author / Participant: `PARTICIPANT-C1-ENG-CONTRACT-REVIEWER-001`; Role: Reviewer; Created/Updated: 2026-10-09. Work Unit / Run: `WU-C1-ENG-001` / `RUN-C1-ENG-CODEREVIEW-T1-001`; Session: `SESSION-C1-ENG-CONTRACT-REVIEWER-001` (invocation binding). Model / Reasoning: `gpt-6-sol` / `medium` menurut invocation, runtime tidak diverifikasi independen. Target baseline/observed HEAD: `112c35f8322b53bac39d511ed4f22982cea619af`; workflow: `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`.

Owning findings: [review-findings-1.md](review-findings-1.md), verdict **Request changes**. Exact reviewed target/input hashes berada di provenance findings. Plan ini rekomendasi untuk Orchestrator/Build; tidak mengeksekusi fix, mengubah authority, atau mendispatch Run.

## Authorized correction boundary

Patch owner memperbaiki editable `api/openapi/index.yaml`, lalu menghasilkan ulang `api/openapi.yaml` dan `api/openapi.d.ts` dengan existing locked tooling. README hanya jika dibutuhkan untuk menjelaskan environment/browser contract. Tidak ada kebutuhan mengubah dependency/lockfiles, backend/frontend production, Product, Techplan/task/manifest, Build evidence terdahulu, atau lint policy.

## Required corrections

1. **F1 — Cookie contract:** Cocokkan production security scheme dengan R10 `__Host-kencleng_session`, host-only/no Domain, HttpOnly/Secure/SameSite=Lax/Path=/. Jelaskan issuance, rotation dan logout expiry untuk cookie yang sama. Development cookie harus berbeda dan hanya explicit localhost:8080 insecure development; jangan menambahkan production auth alternative yang menerima development cookie. Environment representation adalah pekerjaan contract authoring dalam R10 yang sudah settled.
2. **F2 — Mutation browser contract:** Nyatakan shared exact configured Origin + JSON + session-bound X-CSRF-Token, no credentialed CORS, denial sebelum processing untuk logout/prepare/confirm. Deklarasikan permitted JSON logout request sehingga generated consumers dapat mengirimnya (closed empty object adalah salah satu representasi yang cukup); tetap tidak menerima authority-bearing fields. Tambahkan applicable private/no-store invalid-request response. Auth start/callback tetap state/nonce/browser-bound PKCE exception, bukan ordinary mutation Origin/CSRF enforcement. Jangan meminta client JavaScript menetapkan Origin manual.
3. **F3 — Callback failure:** Deklarasikan actual redirection response untuk safe fixed local failure destination; 302 yang existing dapat mencakup success/failure dengan issuance hanya pada successful verified login. Hapus deklarasi 400-as-redirect yang kontradiktif. Nyatakan legitimate provider-error query branch tanpa required code, mempertahankan required state/browser validation untuk trusted processing dan safe handling malformed/missing state. Tidak memantulkan provider error/code/state/return URL dan tidak memberi sesi/Owner pada failure. Tidak ada fabricated 2xx.
4. **F4 — Read invalid input:** Gunakan existing 400 InvalidRequest pada GET preparation, GET Organization, dan list untuk malformed UUID/cursor/limit; unknown/outside-person valid IDs tetap 404 dengan identical shape. Pertahankan private/no-store envelope. Jangan memperluas read menjadi public atau mengganti error dengan handwritten consumer types.

## Focused Build verification and return evidence

- Re-ground input hashes/target diff pada fresh patch Run. Jika parent materially berubah, route reconciliation sebelum patch.
- Jalankan existing API validate/bundle/types dalam Build, inspeksi resulting generated diff, dan catat exact commands/results/hashes. Provision dependencies hanya jika diperlukan melalui locked route; tidak mengganti versions untuk membetulkan semantik.
- Periksa contract secara statis: production versus development cookie, mutation JSON/token/Origin rules, callback success/error branches dengan redirect status, dan 400 read responses. Periksa types logout dan operation response keys setelah generation serta source/bundle equivalence/reproducibility.
- Pertahankan receipt union, four consequences, literal OrganizationView effects, same-person/object scope, guard hold/expiry historical replay, dan processing-unavailable versus outcome-unknown behavior. Tidak menambah authority/version/product choices.
- Tiga known lint warnings bukan acceptance target patch; tidak memilih license atau melonggarkan lint.
- Return bounded Build evidence dan exact patched hashes ke Orchestrator. Perubahan auth/browser/HTTP semantics memerlukan independent review terhadap patched diff. Runtime provider/session/DB/guard/browser correctness dan Human acceptance tetap milik execution/Testing/Human downstream; jangan menyatakan C1 complete dari contract checks.
