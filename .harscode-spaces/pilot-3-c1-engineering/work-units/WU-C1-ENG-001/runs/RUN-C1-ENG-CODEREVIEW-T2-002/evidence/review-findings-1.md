# Review findings — RUN-C1-ENG-CODEREVIEW-T2-002

- Phase / Role: Independent Code Review / Reviewer.
- Author: Codex; Created: 2026-10-11 Asia/Jakarta.
- Work Unit / Run: `WU-C1-ENG-001` / `RUN-C1-ENG-CODEREVIEW-T2-002`.
- Participant / Session bindings: `PARTICIPANT-C1-ENG-T2-CODEREVIEW-002` / `SESSION-C1-ENG-T2-CODEREVIEW-002`; review baru dalam konteks ini, bukan kelanjutan Build atau Review T2-001.
- Model / Reasoning: Invocation mencatat approval `gpt-6-sol / medium`; identitas model/effort runtime tidak diverifikasi secara independen.
- Target HEAD / branch teramati: `059d145d5df64fcda427e5a625486a8e95ce1fd8` / `pilot/3-c1-engineering`.
- Workflow HEAD teramati: `91a97af8ba64f98b0a692a83fbff866a210c45c0`.
- Outcome: **COMPLETED — empat pass berurutan selesai; verdict Approve untuk exact scope enam file.**

## Preflight dan scope

Preflight lulus: seluruh **32 binding SHA-256 unik** pada `invocation.md` cocok, termasuk enam implementation files, tiga coordination authority pasca-approval, Techplan/task/manifest, laporan/gate/capture Build, prior stalled Review, root/backend AGENTS, local config, serta tujuh workflow/routing sources. HEAD, branch dan workflow revision cocok. Prior Review tetap terminal STALLED dan tidak menghasilkan code verdict.

Enam file berikut ditelaah sebagai current diff terhadap target HEAD; ketiga test untracked dibaca lengkap sebagai additions. File lain hanya menjadi authority/context, bukan implementation review scope.

| File | SHA-256 reviewed |
|---|---|
| `backend/go.mod` | `90b7c1c6e870e32828fa5932a6c4e9a28a65d0671a74bdb807ec29dbc5d27cae` |
| `backend/go.sum` | `6923543aab63e09da2b894699afb70b2ff1072a2eb014e063c8542ffed7deaa0` |
| `backend/internal/platform/auth/store.go` | `eb70175ebe756698a3a83b627e0fd6606928d8ac4c355f9dd8a05dab705552b3` |
| `backend/internal/platform/auth/store_test.go` | `bac0d93f0aa6af1591b10dc9a1df81fba0bf77ea2bdb5a165c129e2611795d60` |
| `backend/internal/platform/auth/oidc_test.go` | `2ffd30f781b3aeaa09f6beb0fddd00bea90abcab99b221e126867acfc8331ab4` |
| `backend/internal/platform/auth/store_integration_test.go` | `1b38058a177e0df5f87cf6c732478369e8f798265f2ad91c0762a1ea1879b1e0` |

Full changed-file scope sebelum report: 278 entries, terdiri dari 10 tracked modifications, empat untracked non-Run files, dan 264 untracked Run files. Tracked modifications ialah WU, Events, root/backend AGENTS, go.mod/go.sum/store.go, Control Tower, dependency-verification guide dan scanner runbook. Untracked non-Run files ialah tiga test scope dan network-recovery guide. Untracked Run entries: Build T2-006 (257), prior Review T2-001 (2), Run ini (dua Invocation), Loop Health 002 (3). Tidak ada production delta frontend/API atau backend lain. Instruction/runbook/projection/orchestration edits yang sudah ada dikecualikan sesuai Invocation.

## 1. Safety

**No findings.** Pass ini memeriksa kegagalan, cancellation, lifecycle/resource, shared-state dan auth/session authority pada diff secara tersendiri.

- `store.go:122–145`: public `LockSession(context.Context, pgx.Tx, string, time.Time)` mempertahankan tipe/signature pemanggil. Timestamp pemanggil sengaja diabaikan dan didokumentasikan; private helper menerima clock per-call, tanpa mutable clock global. Clock dipanggil setelah `QueryRow(... FOR UPDATE).Scan` berhasil, sehingga waktu sebelum lock wait tidak dapat mengesahkan session yang sudah expired ketika locked row kembali. Revocation non-nil dan expiry equality/past ditolak. Karena revocation sudah cukup untuk menolak, short-circuit tanpa sampling clock pada revoked row aman.
- Malformed token ditolak sebelum storage; missing row menjadi `ErrUnauthenticated`; query/Scan error dibungkus `%w`; setiap kegagalan helper mengembalikan `Session{}`. `ctx` diteruskan ke pgx tanpa membuat Background baru di jalur production. Helper tidak mengambil ownership transaksi atau menambah goroutine/resource; caller tetap wajib commit/rollback. Existing logout caller memakai transaksi yang sama dan memetakan unauthenticated/storage failure ke safe envelope; retained timestamp tidak lagi menentukan keputusan validitas under-lock.
- Unit barrier/clock bersifat lokal dan serial; tidak ada race global baru. Fake transaksi hanya mengimplementasikan QueryRow yang diperlukan. Nil transaction pada test malformed token membuktikan early rejection, bukan dukungan nil transaction untuk token valid.
- Protocol fixtures memakai signature/JWKS melalui production constructor/verifier, kontrol valid serta negatives, dan tidak menjadikan email/domain/body-person sebagai attribution. Seluruh material key/token/client values adalah fixture; tidak ada production secret atau runtime provider call dalam fixture transport.
- Integration helper memakai generated sanitized schema identifier, parameterized identity values, integration tag dan explicit disposable DSN. Timeout setup/query/cleanup dibatasi; cleanup schema terdaftar segera setelah creation dan pool ditutup sebelum schema/admin cleanup melalui urutan LIFO. Cleanup failure terlihat sebagai test error. Ini tidak membuktikan penghentian proses jika terjadi abnormal suite abort; pemeriksaan tersebut tetap milik execution owner.

Sensitive real DB lock/revocation/expiry interleavings sudah ada di Techplan §9, §12 dan Test Focus Pointer. Tidak ditemukan specialized concern baru yang memerlukan Finding Techplan drift. App wall-clock versus persisted expiry tetap membutuhkan lingkungan waktu yang benar; diff tidak menambah trust/config policy.

## 2. Quality

**No findings.** Pass ini menilai maintainability, kontrak/nama, duplikasi, kompleksitas dan kualitas oracle, tanpa mengandalkan hasil pass Safety sebagai pengganti.

Clock injection private cukup kecil untuk oracle expiry setelah Scan dan tidak memperluas exported API. Komentar public menjelaskan semantic correction walaupun parameter timestamp tetap tersedia bagi pemanggil. Sentinel/error handling tetap jelas; tidak ada dead branch atau logging baru yang membutuhkan observability tambahan.

`store_test.go` memisahkan deterministic simulated lock wait, revalidation matrix, public stale-time regression, dan malformed-input early rejection. Oracle memeriksa digest/FOR UPDATE, hasil attribution valid, zero authority pada error, expiry boundary dan error identity. Public regression memakai waktu nyata dengan margin menit/jam, sedangkan boundary oracle memakai clock tetap; tidak bergantung pada sleep atau scheduling lock nyata. Keduanya memiliki tujuan berbeda. Ini bukan PostgreSQL contention test.

`oidc_test.go` memiliki valid control, 14 negatives dan PKCE/input checks. Fresh verifier/provider per subtest mencegah key-cache state dari satu case mengesahkan case lain. Fixture transport tidak melakukan network; fixture claims menguji issuer/subject-only mapping. Pembangkitan RSA acak menghasilkan key yang valid untuk kedua sisi fixture, sehingga oracle tidak bergantung pada exact key bytes. Margin expiry satu jam menghindari boundary flakiness biasa. Callback/browser/token-provider replay dan cancellation behavior nyata tetap tidak dibuktikan oleh test ini.

`store_integration_test.go` menguji serial person upsert, case sensitivity, distinct issuer/subject dan SQL parameter safety; pending-login matching, one-time consumption dan expiry. Test skip saat DSN absent eksplisit; captured Build execution memperlihatkan dua tests PASS tanpa skip. Tests tidak dinamai atau diklaim sebagai concurrency/SCRAM/privilege acceptance. Tidak ada kebutuhan patch atau minor comment yang dipaksakan menjadi patch loop.

## 3. Stack-Specific Best Practices

**No findings.** Routing dilakukan melalui Go subindex dan targeted PostgreSQL locking row pada global clue map. Matching guidance dibaca dan diterapkan:

- `../harscode-workspace/best-practices/go/context-propagation.md`, `error-handling.md`, `error-wrapping.md`, `nil-and-zero-values.md`, `defer-pitfalls.md`, `struct-copying-and-time.md`: context/error/result semantics, optional revoked timestamp, cleanup order dan Before/Equal comparisons sesuai. Perbandingan `Session{}` pada error menguji seluruh zero result, bukan kesamaan dua populated time instants.
- `go/jwt-and-token-lifecycle.md`: provider verification/algorithm/issuer/audience/expiry dan negative fixtures sesuai concern yang berlaku. Key separation/refresh-family rules tidak diterapkan ke local JWT karena surface ini memakai opaque local session dan tidak memperkenalkan refresh/local signing system.
- `go/testing-concurrency.md`, `go/integration-testing-setup.md`, `postgresql/transactions-and-locking.md`: fake Scan barrier cukup untuk time-order regression; real lock serialization tidak dianggap terbukti. Lock tetap berada dalam caller transaction dan session-first order terdokumentasi. Tidak ada external provider call baru di dalam transaksi. Integration isolation/cleanup sesuai; specialized contention/race checks tetap Testing-owned sesuai Techplan.
- `go/dependency-and-supply-chain.md`: exact pinning, bounded relevant source/API/config review, retained integrity dan full source scan evidence diperiksa; tidak ada floating/latest/replacement/exclusion atau waiver baru.

Dependency review terhadap module files dan retained captures:

1. Go directive `1.26.9`, OIDC `v3.16.0`, OAuth2 `v0.34.0`, jose `v4.1.4` tetap; pgx `v5.9.2` dan text `v0.41.0` sesuai approved D8. Jose menjadi direct karena test addition mengimport package tersebut; versi/source tidak berubah. Sync terpilih `v0.22.0`; graph juga memilih mod `v0.38.0` dan tools `v0.48.0` sebagaimana anticipated transitive reconciliation.
2. Parsed isolated dan final actual graphs memiliki 20 external modules dengan path/version identik dan edges identik. Raw graph JSON berbeda pada Main Dir/GoMod karena probe versus backend serta jose Indirect annotation; bukan graph/version drift. Captured Go minima maksimum 1.25.0 berada di bawah recorded compiler 1.26.9. Tidak ada selected x/net/module addition. go.sum delta konsisten dengan pgx/sync/text versions.
3. Manifest retry memuat 190 capture hashes; semuanya cocok. Origin/upstream validation memuat 12 tag matches serta dua pseudo-version timestamp matches; legacy-source report memuat tujuh modules, 133 file comparisons, zero mismatch. Ini retained Build provenance yang ditelaah, bukan live upstream fetch baru oleh Reviewer.
4. Full isolated, actual pre-G1 dan actual post-G1 JSON streams diparse ulang: masing-masing config 1, SBOM 1, progress 2, OSV 200, finding **0**; source/symbol mode, scanner v1.7.0, Go1.26.9, DB modification `2026-10-08T22:31:09Z`. Command records menunjukkan exit 0; stderr captures kosong. Kesimpulan tidak diambil hanya dari exit code. Retained disposition berisi 55 unique prior IDs. Lima pgx/text IDs memilih versi di luar affected ranges: GO-2026-4771/4772 fixed 5.9.0, GO-2026-5004 fixed 5.9.2, GO-2026-5970 fixed 0.39.0, GO-2026-6629 fixed 0.41.0; named pgx major paths lain absent.
5. Retained source excerpts/release notes menunjukkan pool.New/BeginTx/QueryRow compatibility, CacheStatement default dan URL/per-query override, pgx SCRAM OpaqueString normalization/fallback versus Nickname profile, protocol defaults serta channel binding prefer. Current session query tetap positional tanpa dollar-quoted literal. Tidak ada new application auth/config toggle pada diff. Captured focused compile/integrity mendukung bounded source compatibility; configured SCRAM/DB acceptance masih terpisah.

Tidak ada alasan concrete untuk menjalankan ulang scanner, dependency acquisition, race suite atau DB tests dalam Review. Freshness/konfigurasi pada dispatch Testing selanjutnya tetap tanggung jawab owning Run; review atas retained scan tidak menjamin security universal atau masa depan.

## 4. Consistency

**No findings.** Pass ini kembali memeriksa diff terhadap target-repo authority: root `AGENTS.md`, `backend/AGENTS.md`, `docs/project/kencleng-backend-tech-stack.md`, Approved Techplan §§9–10/12, current T2 task dan accepted manifest.

Error wrapping/sentinels, parameterized SQL, Go/net-http/pgx platform boundary, no-secret logging, serta pemisahan fixture dari real-world evidence sesuai instruksi repository. Exported API tetap compatible pada source/type level; perubahan makna timestamp sengaja memperbaiki aturan under-lock expiry pada Techplan, bukan mengubah requirement. Integration build tag dan scoped disposable state sesuai guidance yang diroute. Tidak ada production/frontend/API/Product/Techplan/task/manifest/blocker write, authorization expansion atau general auth feature baru.

Manifest parent hash dan task baseline-anchor prose adalah historical snapshot yang sudah dijelaskan oleh current scoped Invocation/task; keduanya tidak menggantikan current approved Techplan binding. Tidak ada drift material baru yang perlu diselesaikan melalui perubahan authority dalam Review.

Initial Build terminal report beserta deviation/recovery history tetap dipertahankan; retry `authorization.md` dan separate retry report menjadi evidence occurrence yang Human-authorized. Prior Review STALLED tidak diperlakukan sebagai approval. Report ini tidak merekonsiliasi lifecycle atas nama Orchestrator.

## Verification executed during Review

Executed hanya read-only repository/evidence checks:

- `git rev-parse HEAD`, `git branch --show-current`, workflow HEAD, `git status --short --untracked-files=all` dan scoped `git diff`: memastikan baseline, full changed-file boundary dan actual six-file scope. Ketiga additions dibaca lengkap.
- Python `hashlib.sha256` terhadap 32 bound identities dan 190 retry capture identities: semua cocok. Binding diperiksa kembali pada final boundary; production dan coordination authority tetap sama.
- Python JSON stream parsing/counts, retained command-record inspection, graph path/version/edge comparison, advisory disposition/source/provenance inspection: menguji klaim Build terhadap capture, tanpa mengeksekusi Build commands.
- Captured final unit/protocol stdout menunjukkan PASS; captured two serial PostgreSQL tests PASS tanpa skip. Original-source public regression capture FAIL exit 1 dan patched capture PASS exit 0; ini kontrol Build yang dibaca, bukan reproduction Reviewer. Captured final integrity menyatakan all modules verified dan command exit 0.

**Runtime verification oleh Reviewer: none.** Reasoning dan capture inspection cukup untuk scope ini; tidak ada unresolved concrete uncertainty yang membutuhkan targeted execution. Tidak ada go test/build/race/scanner/fetch/browser/provider/DB execution atau matriks Build/Testing yang diulang. Tidak ada claim bahwa Reviewer independently reran hasil runtime tersebut.

Unverified/deferred: configured real Google login dan real authorization-code replay; browser/HTTPS-cookie behavior; deployed PostgreSQL query/auth/SCRAM, restricted credentials/privileges; concurrent first-login upsert dan session logout/revoke/confirmation contention; real post-lock expiry/cancellation/resource behavior under DB contention; independent final Testing/Human acceptance. Serial fixture menggunakan trust/provisioning identity dan bukan privilege evidence. B-T2-002/B-T2-003 tetap tidak ditutup; T2 completion/T3 readiness tidak ditetapkan.

## Verdict

**Approve — 0 blocking findings, 0 non-blocking code findings.** Berlaku hanya untuk enam file dengan hashes di atas pada target HEAD dan authority Invocation ini. Tidak memerlukan code changes; `patch-plan-1.md` tidak dibuat. Approval Review ini tidak merupakan dependency commit, independent Testing acceptance, blocker closure, atau downstream dispatch.

## Phase handoff

- **Outcome:** COMPLETED — empat independent review passes selesai; Approve untuk exact six-file scope.
- **Result refs:** `evidence/review-findings-1.md` (Run-local terminal carrier); six-file identities pada §Preflight dan scope.
- **Findings:** none; no code patch atau Techplan drift baru.
- **Decision requests:** tidak ada semantic/security decision baru; Orchestrator merekonsiliasi verdict dengan bounded Build retry dan owning acceptance.
- **Blockers:** tidak ada blocker baru dari Review; B-T2-002/B-T2-003 tetap berada pada owning reconciliation, tanpa closure oleh Reviewer.
- **Open / unverified:** configured provider/browser/DB/SCRAM/restricted-role, real contention dan independent Testing/Human evidence sebagaimana §Verification.
- **Recommended continuation:** Orchestrator reconcile lalu route fresh independent Testing sesuai Techplan untuk scope/risk yang applicable; jangan simpulkan T2/T3 readiness atau dispatch otomatis dari Approve ini.
- **Context refs:** `invocation.md`; Approved `techplan/techplan.md` §§9–10/12 dan T2 task; `RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/build-report.md`, gate/capture manifest; reviewed six files.
