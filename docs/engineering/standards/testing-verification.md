# Kencleng — Testing & Verification Standards (T07)

> **Status:** Project-wide Engineering Standard T07 — Human-settled dan persisted di research branch `pilot/station5-c1-space-only`. **Belum berarti seluruh test tooling, coverage enforcement, browser suite, atau runtime verification sudah tersedia/terbukti.** Dokumen ini tidak mempromosikan apa pun ke `main`.
> **Scope:** Pemilihan test yang proporsional, backend/frontend automation, optional browser/integration verification, quality/coverage, evidence, dan resource discipline.
> **Authority:** [Engineering Standards](README.md). Approved Product/scoped requirements menetapkan behavior/invariant; T07 menetapkan cara memilih dan membuktikan verification. T01–T06 tetap owning standards untuk architecture, auth, authorization, data protection, DB, dan API. Source code + executed checks adalah evidence, bukan pengganti requirement atau accepted contract.

## Prinsip dasar: sufficient evidence, not ceremonial testing

**Tidak semua test suite wajib pada setiap perubahan.** Agent wajib memahami *changed behavior*, consequence, reversibility, blast radius, dan failure modes yang relevan, lalu memilih **verification minimum yang cukup dan dapat dipercaya**. Tidak menjalankan test irrelevant adalah keputusan yang baik; mengabaikan material risk tanpa bukti bukan demikian.

- **Wajib:** verification yang proporsional untuk perubahan consequential; klaim evidence yang akurat; applicable unit test quality/coverage pada T07-I; security/financial safeguard pada T07-F.
- **Conditional:** race, DB integration, contract, Playwright/browser, full build, security scan, extended load/performance hanya saat perubahan atau acceptance requirement membenarkannya.
- **Tidak diwajibkan:** exhaustive full-repository/full-stack test suite untuk setiap code generation, setiap refactor, atau setiap perubahan styling/docs.
- Suatu test yang tersedia bukan berarti harus dijalankan. Sebaliknya, shortcut untuk menghindari pengujian invariant kritis tidak dibenarkan.
- Ketersediaan `make verify`, `npm run verify`, atau test scripts tidak menetapkan kebijakan mandatory. Pilih perintah secara independen sesuai scope; jangan menganggap green command tanpa cases sebagai bukti behavior.

## T07-A — Risk-based verification selection

- Mulai dari *apa yang berubah* dan *apa yang bisa gagal*, bukan urutan tetap semua tool. Tentukan business impact, security boundary, data durability, concurrency, external dependencies, dan UI behavior; pilih focused checks yang dapat menangkap failure modes tersebut.
- **Documentation-only:** source/link/meaning review yang relevan, bukan menjalankan Go race/Playwright. **Pure backend logic:** targeted unit tests. **Backend DB transaction:** DB integration/concurrency bila invariant menuntut. **Frontend form behavior:** focused Vitest/RTL; visual/navigation/browser-only risk dapat memerlukan rendered/browser verification.
- Pengecualian dari test suite yang tidak relevan tidak memerlukan prosedur panjang; berikan alasan singkat jika hasil verifikasi/risiko bisa disalahpahami. Test penting yang tidak bisa dilakukan menjadi **Not verified/Blocked**, bukan `N/A`.
- Selected test tidak selalu harus dibuat baru. Existing valid checks dapat dipakai bila benar-benar meng-cover changed behavior.
- Test depth harus meningkat sesuai high-impact/hard-to-reverse operations, bukan berdasarkan jumlah file atau kesan kompleks.

## T07-B — Backend unit, integration, race and concurrency

- Untuk **new/changed testable Go behavior**, tulis/jalankan targeted unit tests yang memvalidasi contract/business outcomes. Gunakan `go test` dengan package scope yang relevan, bukan mewajibkan full `./...` setiap iterasi kecil.
- `go test -race` dipilih ketika ada **shared mutable state, goroutine, concurrent access, locks, cache, shared process resources**, atau evidence race yang masuk akal. Race detector bukan kewajiban semua API dan bukan alat untuk membuktikan database transaction isolation. Volume request rendah **tidak otomatis meniadakan race risk**.
- Untuk **DB transactional correctness**, gunakan targeted integration/concurrency/rollback/retry tests terhadap database yang sesuai bila correctness mencakup cross-row/cross-request interleavings. Go race detector tidak menggantikan itu; unit mocks tidak membuktikan transaction semantics PostgreSQL.
- Boundary external services, real DB integration, external callback dan worker recovery diuji ketika perubahan sungguh bergantung pada semantik sistem tersebut. Mocks/stubs sesuai untuk unit isolation tetapi **tidak boleh diklaim sebagai real integration evidence**.
- Jika concurrent execution atau idempotency effect critical, uji duplikasi, partial commit, ambiguous commit, dan negative cases sesuai requirement; jangan membuat generalized stress test infra hanya untuk formalitas.

## T07-C — Frontend lightweight automated baseline

- Default frontend adalah **lightweight automated-first**, meminimalkan ketergantungan pada Human manual testing. Gunakan ESLint dan **focused Vitest/React Testing Library** untuk applicable changed behavior (validation, interaction, conditional state, submit/error/loading, accessibility semantics yang dapat diuji di komponen).
- `npm run lint` tetap useful static check; `npm run test` atau focused test selection harus benar-benar mengeksekusi relevant cases ketika behavior diklaim tested. **Vitest `--passWithNoTests` dapat berstatus sukses tanpa satu test; itu hanya tooling invocation, bukan behavior verification**.
- Styling/copy/docs-only change tidak otomatis perlu unit test baru. Material visual/UI work memerlukan proportional rendered verification; lint atau jsdom component tests sendiri tidak membuktikan screenshot layout, browser interaction, responsive state, atau rendering real.
- Build check (`npm run build`) dipilih bila deploy/build-specific risk relevan; build yang sukses bukan bukti business/UX correctness. MSW dapat digunakan untuk contract-faithful FE isolation; tidak boleh membuat mocked success seolah real backend sudah berfungsi.
- Tooling tersedia (Vitest, RTL, MSW, Playwright) **bukan kewajiban mengaktifkan semuanya** pada setiap generated component.

## T07-D — Selective real-browser, Playwright and Human opt-in

- Playwright adalah **optional targeted browser/integration capability**, **bukan mandatory gate setiap frontend code generation**. Human dapat meminta automated integration test untuk selected flow/milestone; Agent juga dapat merekomendasikan/menjalankan scoped browser checks bila real-browser-only risk cukup material dan environment/authority mendukung.
- Sebelum menjalankan browser integration, verifikasi readiness: local FE, BE, PostgreSQL, browser executable, trusted test data/auth/session, available ports, dan side effects. Gunakan [local development modes](../development/local-development-modes.md) serta current Playwright config, bukan asumsikan FE/BE/DB semua sudah berjalan.
- Jangan mengubah DB volume, melakukan destructive fixture setup, reset secrets, menyalakan layanan eksternal, atau mengganti active local environment Human tanpa scope/approval yang sesuai.
- `npm run test:browser`, headed/UI modes, dan existing opt-in `make check-hot-reload` hanya dijalankan bila relevan; hot-reload probe **mengubah source sementara** dan tidak boleh otomatis dijalankan di tengah edits.
- Jika browser/runtime/dependency tidak tersedia, laporkan **Blocked/Not verified** beserta kondisi konkret. Mock-based component checks **tidak boleh** diberi label real FE↔BE integration.
- Tidak mewajibkan screenshot suite, broad E2E coverage, atau penulisan Playwright spec untuk setiap component. Selective tests harus memeriksa real accepted outcomes, bukan sekadar halaman berhasil dibuka.

## T07-E — API contract, persistence and environmental evidence

- Ketika active OpenAPI berubah, lakukan applicable validation/bundling pada `api/`, serta conformance tests untuk consumer/server ketika diimplementasikan. Lint/bundle hanya membuktikan syntax/schema validity, **bukan** backend/FE behavior.
- Untuk migration/schema change, gunakan relevant fresh-DB apply dan existing-data upgrade/constraint checks; evaluate backward compatibility, transactional data, rollback/recovery sesuai [T05](database-data-model.md). Jangan pernah menjalankan destructive `migrate-down`/`down -all` sebagai generic verification.
- Gunakan real integration checks bila correctness membutuhkan actual PostgreSQL constraints, SQL transactions, cookie/CSRF in browser, backend service identity, atau actual file/object-store semantics. Unit mocks dan generated types tidak setara dengan live evidence.
- Kebutuhan environment runtime yang lebih mahal ditentukan oleh consequence dan requirement, bukan dilakukan untuk semua changes.

## T07-F — Security-critical and consequential change verification

- Untuk authentication/session rotation, authorization/grants, PII encryption/reveal/masking, money movement, ledger/settlement, or other protected operations: pilih targeted negative/abuse, replay/idempotency, privilege boundary, transactional/concurrency, dan recovery tests yang benar-benar membuktikan critical invariants.
- Risiko tinggi **bukan alasan mewajibkan semua jenis suite sekaligus**: tuntut evidence tepat sesuai failure mode, dengan test depth yang cukup. Misal CSRF perlu negative Origin/token checks; DB settlement concurrency perlu transaction/interleaving checks; Go race detector hanya bila relevant shared-memory risk.
- Root [repository AGENTS](../../../AGENTS.md) dan corresponding security standards tetap mengikat **protected implementation writes** dan Human authorization; T07 approval bukan izin otomatis mengubah protected code.
- Jika verification yang material tidak dapat dilakukan, nyatakan implemented but not fully verified/blocked dengan risk yang tersisa. Jangan menyebut perubahan payment/security `validated` hanya karena unit tests green.

## T07-G — Evidence, claims and handoff

- Catat secara ringkas: **executed + result** (command/scope dan relevant assertions), **not applicable** (kenapa tidak relevan), **not verified / blocked** (relevant tapi belum dieksekusi, alasan dan impact), dan failures yang masih ada.
- `Passed` hanya untuk checks yang betul-betul dijalankan; no test cases, mocked external integration, `build` pass, lint, coverage percentages, dan screenshots masing-masing membuktikan hal berbeda.
- Bedakan **implemented**, **contract-conformant**, **integration-verified**, **render-verified**, **security-verified**, dan **not verified**; tidak perlu mandatory report template panjang pada low-risk changes. Test selection rationale boleh hanya 1–2 kalimat.
- Failed test tidak diabaikan dengan menonaktifkan test atau menurunkan threshold hanya demi green status. Document known gaps and required next evidence.

## T07-H — Independent execution and resource discipline

- **Tidak ada single mandatory mega-script yang menjalankan optional suites** hanya karena script tersedia. Mandatory lightweight baseline boleh digabung seperlunya; race, integration, contract, browser/E2E, full security audit, dan performance suites tetap **individually invokable** dan dipilih berdasarkan relevansi.
- **Current tooling gap (not yet corrected):** `backend/Makefile` mempunyai `.DEFAULT_GOAL := verify` dan `verify: lint test-unit test-race test-contract security`. Ini mencampur optional full checks. **Jangan gunakan `make` tanpa target atau `make verify` sebagai mandatory default T07** sampai Makefile diselaraskan lewat scoped tooling change. Perintah individual tetap boleh dipakai sesuai relevance.
- **Existing frontend `npm run verify`** hanya menggabungkan lint + Vitest, tidak otomatis menjalankan browser. Boleh dipakai bila kedua checks relevan, tetapi `--passWithNoTests` tidak boleh dimaknai sebagai coverage/behavior tests.
- Do not automatically provision browsers, containers, DB, run migrations, mutate shared fixtures, or launch high-resource checks tanpa risk justification/permissions. Tooling simplification should solve observed friction; tidak perlu membangun testing orchestrator atau state machine.

## T07-I — Unit coverage ≥80% AND behavior quality

- **Minimum 80% unit-test statement coverage pada affected testable feature/package/module scope**, bukan angka seluruh repo yang dapat menutupi logic baru yang belum diuji, atau memaksa empty baseline/irrelevant packages mencapai angka. Scope dan denominator harus eksplisit pada evidence. Target adalah *floor*, bukan reason untuk test yang hanya meningkatkan angka.
- **Applicable unit-test gate:** untuk new/modified production behavior yang layak diuji, ukur coverage setelah meaningful tests dijalankan. Branch/domain invariants yang consequential tetap harus diuji meskipun coverage ≥80%. Code yang genuinely generated/unreachable/boilerplate dapat dikecualikan hanya dengan transparent, stable rationale, bukan mengecilkan denominator ad hoc.
- **Go backend: table-driven tests + GIVEN/WHEN/THEN** untuk setiap new Go unit test function. Gunakan `tests := []struct{...}` dengan `name` pertama dan expected `want`/`wantErr` di akhir sesuai kebutuhan; iterate dengan `t.Run(tt.name, ...)`. Setiap case name memakai `GIVEN <context> WHEN <action> THEN <observable result>`; di body tiap subtest pakai **tiga baris komentar terpisah**: `// GIVEN`, `// WHEN`, `// THEN`. Satu-satunya case tetap boleh/harus ditulis table-driven untuk test unit baru.
- Unit tests harus membuktikan **observable business results**, bukan hanya lines executed/mock invoked. Saat relevant assert output **dan** error, meaningful dependency arguments, state transition, side effects, bookkeeping, dan forbidden effects. Kasus negative harus membuktikan tidak ada prohibited mutation; pilih relevant success, validation, boundary, empty/nil, dependency failure, alternate branch/error propagation **berdasarkan requirement**, bukan mencentang seluruh scenario families secara ritual.
- Isolasi mock/stub/dependencies per subtest, deterministik, tidak mengandalkan shared mutable globals, real network, DB, external clocks, atau order of execution untuk **unit tests**. Fixtures, fixed clock, dan fakes yang sesuai boleh dipakai. No hardcoded dependency pada `mockery`, `sqlmock`, BigQuery, atau harness project lain.
- Reference guidance berasal dari Anhar's supplied **"How To Write Unit Tests (product service)"**: reusable principle *coverage is the floor, behavior is the goal*, table-driven scenario design, assertion quality, isolation, dan determinism diadopsi; migration-specific mock APIs, mockery constructors, fixtures dan repo-target commands **tidak diadopsi sebagai Kencleng authority**.
- **Coverage implementation gap:** `go test -cover` / suitable profile dapat mengukur scoped Go coverage; existing frontend Vitest scripts belum otomatis menghasilkan/menegakkan threshold 80% dan coverage provider may need setup. **Jangan klaim gate enforced untuk FE/BE hanya karena policy disetujui.** Jika belum dapat mengukur affected scope secara kredibel, laporkan not verified, lalu tambahkan tooling sekecil yang diperlukan saat actual feature implementation.
- Unit coverage bukan pengganti database/HTTP/browser/security integration evidence. Quality ditentukan oleh relevant domain outcome assertions dan remaining risk, bukan hanya percentage.

## Agent verification selection scaffold (per active work, concise)

| Pertanyaan | Evidence yang ditentukan dari kebutuhan nyata |
| --- | --- |
| Apa behavior/invariant yang berubah? | Requirement/contract dan risk level |
| Apa failure modes yang masuk akal? | Validation, security, transaction, concurrency, UI/visual, compatibility |
| Apa smallest sufficient automated checks? | Target package/component, commands/cases/coverage scope |
| Perlu external integration/browser/race/security suite? | Hanya jika relevant; nyatakan trigger/justification dan environment readiness |
| Apa yang benar-benar dieksekusi? | Command/test scope, outcome, measured coverage bila applicable |
| Apa yang belum dibuktikan? | Not applicable vs relevant-but-blocked; residual risk/next action |

**T07 tidak mewajibkan pembuatan test framework atau suite untuk empty project baseline.** Saat feature nyata masuk, gunakan checklist ini sebagai reasoning aid yang proporsional, bukan bukti otomatis semua tests harus dijalankan.

## Reference & execution boundaries

- [Frontend architecture](frontend.md), [Backend architecture](backend.md), [API contract standard](api-contract.md).
- [Local development modes](../development/local-development-modes.md), [Frontend setup](../../../frontend/README.md), [Backend setup](../../../backend/README.md).
- [Go testing](https://pkg.go.dev/testing), [Go race detector](https://go.dev/doc/articles/race_detector), [Vitest](https://vitest.dev/), [Playwright](https://playwright.dev/).

**Status implementation:** T07 policy approved/persisted only. Backend `Makefile` optional-suite aggregation remains an observed follow-up; frontend coverage instrumentation/gate and first useful test cases require separate scoped implementation and evidence. Ini bukan perubahan `main` atau izin menjalankan destructive local integration.
