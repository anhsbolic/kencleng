# C1 — Hold verifikasi terintegrasi sampai frontend/backend siap

- **Decision ID:** D-C1-INTEGRATED-VERIFICATION-HOLD-001
- **Status:** Hold aktif untuk scope terintegrasi di bawah; evidence final belum established.
- **Tanggal:** 2026-10-12
- **Work Unit:** WU-C1-ENG-001 / C1
- **Human owner:** Anhar
- **Coordination owner:** Orchestrator
- **Verification owners:** Testing dan Human, sesuai Approved Techplan §12.

## Tujuan dan sumber keputusan

Anhar meminta dokumen untuk hal-hal lain yang dapat di-hold sampai integrasi selesai, setelah sebelumnya menetapkan [hold real Google login](c1-google-login-verification-hold.md). Dokumen ini mencatat penundaan bukti **final terintegrasi** untuk empat area yang telah dibahas: full journey C1, rendered UI/accessibility, cache isolation, dan manual walkthrough makna Owner/konsekuensi.

Scope ini mengikuti owner dan dependency yang sudah ada pada Approved Techplan serta accepted task manifest. Hold mengatur kapan evidence integrasi dikumpulkan; requirement, aturan backend dan kondisi acceptance tidak dihapus atau dianggap lulus. Bounded component/fixture checks yang sudah executable dapat tetap berjalan, dengan label evidence yang tepat.

## Daftar verifikasi yang di-hold

| Area / rule terkait | Bukti final yang ditunda | Owner / execution mode setelah integrasi | Prasyarat khusus |
|---|---|---|---|
| **Full journey C1** — R1/R2/R4/R5/R6/R8/R11/R13 pada surface terintegrasi | Form → sign-in bila diperlukan → preparation → frozen name/consequence summary → explicit confirm → hasil → my-list/detail; durable navigation/reload/relogin dan outcome gagal/unknown yang tersedia. Tidak ada auto-confirm setelah callback/load. | Testing; real Google step manual Human sesuai dokumen terpisah. Bagian lain dapat diuji agent/automation bila aman dan faithful. | FE T5 tersambung backend real T2/T3 serta endpoint/data/guard prerequisites untuk scenario yang dipilih. |
| **Rendered UI, responsive dan accessibility** — R2/R4/R5/R6/R13 | Final rendered narrow/wide, long name, keyboard/focus, hierarchy/overflow, reachable actions; loading/invalid/login-required/expired/hold/definite-failure/unknown/empty/success pada aplikasi terintegrasi. | Testing; browser automation dan inspection terarah dapat digunakan. Manual jika scenario faithful memang memerlukannya. | UI terintegrasi dan state fixture/runtime yang sah, tanpa relabel simulasi menjadi real outcome. |
| **Private frontend cache / pergantian pengguna** — R8/R10/R13 | Dua-person browser journey, logout/relogin, clearing cache/CSRF context, tidak ada data/nama/authority pengguna sebelumnya, scoped list/detail dalam real integration. | Testing; agent/automation yang aman dapat menjalankan bagian yang tidak memerlukan private Google login. Google interaction tetap manual Human. | Session/person setup sah, FE cache/client nyata dan backend scoped reads tersedia. |
| **Human walkthrough makna dan konsekuensi** — R2/R4/R5 | Human memahami empat consequence meanings sebelum tindakan; detail/list menjelaskan initial Owner, internal authority, organization-provided provenance dan effects tanpa kesan legal verification/review/curation. | **Human-owned, manual**; agent menyiapkan skenario/oracle dan mencatat explicit Human decision, tidak menilai lulus atas nama Human. | Pengalaman rendered yang cukup lengkap untuk memahami sebelum/sesudah efek pada integrasi real. |

Tidak menunda seluruh rule ID di tabel. Rule dapat memiliki bukti backend, unit dan Human/UI yang berbeda; hold berlaku hanya pada bagian final terintegrasi yang disebutkan. Contoh: UI presentation unknown outcome ditunda, tetapi DB atomicity/commit-ack loss/replay checks tidak otomatis ikut ditunda.

## Pekerjaan yang tidak ikut di-hold

- Scanner G124/Gitleaks disposition, policy/remediation yang disahkan owner, integrity/dependency gates dan final backend make verify.
- Backend HTTP/contract assertions, authorization/scoped reads/IDOR, cookie flags, Origin/CSRF/logout/revocation; checks yang tidak memerlukan real Google login dapat berjalan dalam authorized scope.
- Restricted-runtime PostgreSQL/SCRAM, concurrent first-login upsert, session-row locking/expiry/revocation/cancellation; real DB constraint/atomicity/guard/privilege checks setelah owning surface tersedia.
- T3/T4 operator/provisioning/runbook obligations; kesiapan mengikuti owning implementation dan explicit permission, bukan menunggu frontend secara otomatis.
- Build-authored focused checks, isolated frontend/component/browser development checks dengan fixtures/MSW dan Review. Lulus mock/component bukan final real integration PASS.
- Production HTTPS/cookie prerequisite tetap terpisah. Dokumen ini tidak memberi blanket hold atau mengganti HTTPS evidence dengan localhost.

Real Google login memiliki keputusan hold dan **manual-only execution** tersendiri. Batas manual-only itu tidak otomatis berlaku untuk seluruh E2E/browser check dalam dokumen ini. Testing-owned automated/agent verification yang aman dan representative tetap dapat digunakan setelah integrasi; Human-owned semantic walkthrough tetap milik Human.

## Risiko penundaan dan pengendalian

| Risiko | Cara menjaga scope/evidence |
|---|---|
| Cross-layer mismatch baru terlihat setelah integrasi | Pertahankan coordinated API contract, typed consumers, focused backend/component tests dan patch/report flags selama implementasi. |
| UI terlihat sukses walau authority/persistence salah | Backend tetap security authority; fixture/toast/frontend role tidak menjadi bukti authorization atau durable success. |
| Private data/cache bocor saat pengguna berganti | Cache-isolation final tetap wajib saat integrasi; checks backend scoping/no-store dan FE cache design tidak dihapus. |
| Hold terlupakan atau berubah menjadi acceptance implisit | Daftar scenario/owner/status deferred dibawa ke Invocation/report berikutnya; release dan hasil direkonsiliasi eksplisit. |
| FE integrasi terlambat | Bukti backend yang tidak membutuhkan FE dapat diteruskan dalam scope sah; acceptance yang memerlukan integration tetap belum established. |

## Release readiness dan status hold

Orchestrator mengevaluasi pelepasan per area ketika owning FE/backend implementation dan integration evidence menetapkan real surface yang diperlukan sudah tersedia. Tidak mensyaratkan seluruh environment sempurna untuk semua bounded Testing; cukup prerequisites material bagi scenario yang akan diklaim.

1. Identifikasi build/revision, scenario dan backend endpoint yang benar-benar terhubung; kompilasi/mock/MSW saja bukan evidence integrasi selesai.
2. Rekonsiliasi prerequisites sesuai scenario: T2/T3/session/guard/scoped reads, authorized disposable data/identities, browser route serta environment error/unknown paths yang diperlukan. Jangan provisioning atau role escalation tanpa authority.
3. Testing menyiapkan observable oracle dan kategori evidence. Human semantic walkthrough memiliki pertanyaan/criteria sebelum Human execution; real Google mengikuti [dokumen hold Google](c1-google-login-verification-hold.md) dan [environment guide](c1-auth-verification-environment.md).
4. Scope yang belum executable tetap deferred dengan owner/gap; scope executable dapat diroute secara bounded. Fresh phase execution memakai Run/Participant/Session baru; terminal Runs tidak dibuka kembali.
5. Catat release/readiness dan hasil tiap scenario pada owning evidence. Integrasi selesai bukan automatic PASS atau whole-C1 acceptance.

## Evidence dan acceptance boundaries

- Current scope berstatus **DEFERRED / belum terverifikasi** untuk final real integration. Dokumen ini bukan test report dan tidak menjalankan verifikasi.
- Pisahkan real runtime, simulated/fixture, command-only/no assertions, agent-executed dan Human-attested. Tidak infer observasi dari broad “E2E PASS” tanpa scenario/result yang jelas.
- Jangan merekam credentials, codes/tokens/cookies, atau account/person identifiers yang tidak diperlukan; Human menyampaikan hanya relevant observable outcomes.
- Hold tidak menutup B-T2-002/B-T2-003, menerima D-PATCH002-01, mengesahkan dependency commit atau menyatakan T2/T3/C1 accepted.
- Accepted manifest tetap: T1→T2→T3→T4, T1→T3, T1→T5; T5 bisa mengembangkan contract-faithful fixtures, dengan real integration evidence downstream. Edge satisfaction/readiness tetap harus direkonsiliasi terhadap owning conditions; hold sendiri bukan pemenuh edge.
- Tidak ada rewrite Approved Techplan/task snapshots atau upstream Product authority. Jika routing/acceptance membutuhkan perubahan material kontrak, route ke owner; jangan melemahkan aturan melalui dokumen hold.

## Referensi

- [Approved Techplan §12](../../.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.md)
- [Accepted execution manifest](../../.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/tasks/manifest.md)
- [Hold real Google login](c1-google-login-verification-hold.md)
- [C1 auth verification environment](c1-auth-verification-environment.md)
- [Work Unit](../../.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001.md) / [Events](../../.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/events.md)
