# Build report — RUN-C1-ENG-BUILD-T2-004

> Phase: Build; Role: Implementer T2/G1; Participant / Session bindings: `PARTICIPANT-C1-ENG-T2-IMPLEMENTER-004` / `SESSION-C1-ENG-T2-IMPLEMENTER-004`. WU: `WU-C1-ENG-001`. Created/Updated: 2026-10-11. Target HEAD: `70f0ca000054001cbc7b1d1fd23e1e8e984e8395`; observed workflow HEAD: `91a97af8ba64f98b0a692a83fbff866a210c45c0`, clean. Invocation selects `gpt-6-sol` / `medium`; runtime model/effort tidak terverifikasi secara independen. Human memberikan kickoff spesifik Run; Invocation masih mencatat model approval PENDING. Tidak ada klaim approval model terpisah atau perubahan registry.

## What changed

**BLOCKED — berhenti pada isolated graph gate D8 karena fresh full source scan gagal.** Backend aktual tidak diubah. Pin D8 hanya diterapkan dalam probe `/tmp/c1-t2-004-8n2xu6i1`. Koreksi expiry `LockSession` dan focused protocol-negative/person-upsert/session-lock checks belum dilakukan. Tidak ada commit, perubahan upstream, penerimaan dependency/risk, atau dispatch task lain.

Probe dibuat dari live `backend/go.mod`, `go.sum`, `cmd/`, dan `internal/` melalui `tempfile.mkdtemp`, `shutil.copy2`, dan `shutil.copytree`. Tidak menyalin `.env` atau konfigurasi rahasia. Hanya dua pin diubah secara eksplisit: pgx/v5 v5.8.0 → v5.9.2 dan x/text v0.29.0 → v0.41.0. `go mod tidy` memilih x/sync v0.22.0 menggantikan v0.17.0 sebagai delta transitif; pin Go/OIDC/OAuth2/jose dipertahankan. Tidak memilih versi alternatif.

Satu-satunya artifact repository yang ditulis Run ini adalah `evidence/build-report.md` ini. Identitas changed backend files: **none**, dikonfirmasi `git diff --exit-code -- backend` exit 0. Existing Orchestrator edits pada Space README, WU, events, T2 snapshot, Control Tower, serta direktori T2-003/T2-004 dipertahankan. Report tidak memuat checksum dirinya sendiri.

## Input identities

Semua bound hashes cocok dengan Invocation. Path berikut relatif WU kecuali `backend/` dan `api/`:

```text
6d2e41e177e4adaf74f53a06d81dd568def75019c8dc4e4792d938a869896d79  techplan/techplan.md
dfe17b7a3dc0dc89127b040a67f547506a14a4ab591ed0a6343f1a8dd8bcce37  techplan/tasks/T2-backend-identity-session.md
2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d  techplan/tasks/manifest.md
18562231cc39a6c1f85ce7c7c20cc8e161d00a203a2b612f2f60117e02b2e611  techplan/report-techplan.md
f87ea3ae133cb8279008b31dadd3cb501f9e50c31a56817313e3d5f42dd93853  runs/RUN-C1-ENG-BUILD-T2-003/evidence/build-report.md
474e39ddd7358c4bf0f1c94a1cb217406cd81285046b06207a8cb4332e91bf93  runs/RUN-C1-ENG-BUILD-T2-002/evidence/build-report.md
6095e0b6d7d80d83c10411602246c2cf0c02b1d78888f5579f0350ad57823b9f  runs/RUN-C1-ENG-BUILD-T2-002/evidence/dependency-graph.txt
15ea30ec5eb124d75299be0109f632cd4884648464525b06d9d5bfa02d4bf4fe  runs/RUN-C1-ENG-BUILD-T2-002/evidence/dependency-vuln.json
b4537ed75f533f993f371954de47e42a793b8e5b0587577de7e27fb3e50696bd  runs/RUN-C1-ENG-BUILD-T2-002/evidence/dependency-integrity.txt
2c089366f9b54e86da2993dadbeffe3cd7d3786a7b5f0535dd9163521e4db5d3  backend/go.mod
979a350418d3d0e4906ea8af162cb69f8cfcfd38507eaa7d8657f6c7c5d7960f  backend/go.sum
dfeed8714bc1881b9c94e8b9ac8f29ed2226f7396b1f2c789b0bc7dcfd0032b9  backend/internal/platform/auth/store.go
ddefffbe588459cb0cba9f68dfedaccba3da8a1313f1956511cf628df7199965  api/openapi/index.yaml
38bb6c531804b54e9d5ce56d623c397a590722c56dcb7a3c12e646bb8280fa09  api/openapi.yaml
d21e8aa1eec92a95cc1fed87d976f726f59b48bd334aa8e463844d48e16cb80c  api/openapi.d.ts
```

Dibaca: root/backend AGENTS, Invocation, current T2 snapshot, approved Techplan §§9–10/12, prior T2-003 report, backend architecture, live auth store, server composition dan db.Open; current Harscode Build prompt/guidelines/checklist/overlay. HEAD dan semua hash di atas diperiksa dengan `git rev-parse HEAD`, `sha256sum`; worktree dengan `git status --short` dan backend diff. Tidak ditemukan bound-input mismatch. F-T2-003 sudah menjadi koreksi yang diotorisasi, bukan new finding Run ini.

## Tests run

Semua perintah Go berikut memakai `GOTOOLCHAIN=go1.26.9`; kecuali `go version`, cwd adalah probe. Graph/integrity/compile dan scan dimulai independen; setelah scan gagal tidak ada kelanjutan gate/patch.

| Command | Category | Result |
|---|---|---|
| `GOTOOLCHAIN=go1.26.9 go version` | Toolchain | exit 0: `go version go1.26.9 linux/amd64` |
| `GOTOOLCHAIN=go1.26.9 go mod tidy` | Isolated resolution | exit 0, approved pins retained, x/sync transitif v0.22.0 |
| `GOTOOLCHAIN=go1.26.9 go list -m -json all > /tmp/c1-t2-004-8n2xu6i1/graph.json` | Isolated graph | Success; complete selected module list below |
| `GOTOOLCHAIN=go1.26.9 go mod verify` | Integrity | exit 0: `all modules verified` |
| `GOTOOLCHAIN=go1.26.9 go test -run '^$' ./internal/platform/auth ./internal/platform/db ./cmd/server` | Focused compile | exit 0; ketiga package `[no test files]`, tidak membuktikan functional auth checks |
| `GOTOOLCHAIN=go1.26.9 govulncheck -json ./... > /tmp/c1-t2-004-8n2xu6i1/vuln.json` | Required fresh full source scan | **exit 1**, package loading/build-cache errors; tidak menghasilkan findings/traces lengkap |
| `GOTOOLCHAIN=go1.26.9 go mod download -json github.com/jackc/pgx/v5 golang.org/x/text github.com/coreos/go-oidc/v3 golang.org/x/oauth2 github.com/go-jose/go-jose/v4` | Module metadata/checksums | exit 0; lima pin exact tersedia di module cache; output tidak memuat Origin metadata |
| `git diff --exit-code -- backend` | Write boundary | exit 0, backend tetap baseline |

## Finding F-T2-004 — full source scan tidak lengkap

`govulncheck` melaporkan `loading packages` lalu delapan error `loading compiled Go files from cache: reading srcfiles list: cache entry not found: open ...: no such file or directory`. File yang disebut berada di `/home/anhar-solehudin/.cache/go-build/`, dengan hash berikut (semua suffix `-a`):

```text
7f/7f9e103c6047b92fca9782782f424ea428cae58d0984b65cc087b297657ba810-a
9c/9c466a7fe7fade589f460190f1bd53124220a097f8965f113fd0736e7bba8bd4-a
be/be916f508b52a0f3e6559fbb24c810eb6cb3335cfae32712f9fabfc5ce16f5d8-a
fd/fddb205bb960531635917f1d3fbac6446410f0712e6d4aa4137090fdd94f638d-a
d5/d52c5a9c89551bc0c86d81e07469d274e623bef51658f703999e009195a7c938-a
86/86f9a4afea11f5495cd1998386b2454550b70c96a41a134d6ce9c0e32a1e4751-a
73/73acd6bf6fe0164863bc0fb518b347c113e90cbd8a064b11a996a5a97c3edaf1-a
6c/6cde42cf426e79399cc3552b1780fb53e7801184637f9d5bcd97100291281df5-a
```

Output JSON hanya konfigurasi berikut; tidak ada bukti scan selesai:

```json
{
  "config": {
    "protocol_version": "v1.0.0",
    "scanner_name": "govulncheck",
    "scanner_version": "v1.7.0",
    "db": "https://vuln.go.dev",
    "go_version": "go1.26.9",
    "scan_level": "symbol",
    "scan_mode": "source"
  }
}
```

Verified: scan exit 1, error cache di atas, JSON hanya config; integrity/compile yang berjalan independen sukses. **Root cause belum ditentukan**; tidak menyimpulkan sandbox, compiler incompatibility, concurrent-cache race, vulnerability baru, atau graph clean. Tidak membersihkan cache, mengganti scanner/toolchain, mererun dengan workaround, suppress/waive finding, ataupun menerima risiko. Sesuai stop condition Invocation §3 dan instruksi Human, kegagalan scan menutup kelanjutan Run ini sebelum tracked writes. Owner: Orchestrator/Anhar untuk bounded environment/scanner reconciliation dan fresh execution scope; Techplan owner bila bukti selanjutnya memerlukan pilihan material.

## Isolated graph and partial dependency evidence

| Module | Selected version | Go minimum from `go list -m -json all` |
|---|---|---|
| github.com/anhsbolic/kencleng/backend | main | 1.26.9 |
| cloud.google.com/go/compute/metadata | v0.3.0 | tidak tersedia pada output |
| github.com/coreos/go-oidc/v3 | v3.16.0 | 1.24.0 |
| github.com/davecgh/go-spew | v1.1.1 | tidak tersedia pada output |
| github.com/go-jose/go-jose/v4 | v4.1.4 | 1.24.0 |
| github.com/jackc/pgpassfile | v1.0.0 | 1.12 |
| github.com/jackc/pgservicefile | v0.0.0-20240606120523-5a60cdf6a761 | 1.14 |
| github.com/jackc/pgx/v5 | v5.9.2 | 1.25.0 |
| github.com/jackc/puddle/v2 | v2.2.2 | 1.19 |
| github.com/joho/godotenv | v1.5.1 | 1.12 |
| github.com/kr/pretty | v0.3.0 | tidak tersedia pada output |
| github.com/pmezard/go-difflib | v1.0.0 | tidak tersedia pada output |
| github.com/stretchr/objx | v0.1.0 | tidak tersedia pada output |
| github.com/stretchr/testify | v1.11.1 | 1.17 |
| golang.org/x/mod | v0.38.0 | tidak tersedia pada output |
| golang.org/x/oauth2 | v0.34.0 | 1.24.0 |
| golang.org/x/sync | v0.22.0 | 1.25.0 |
| golang.org/x/text | v0.41.0 | 1.25.0 |
| golang.org/x/tools | v0.48.0 | tidak tersedia pada output |
| gopkg.in/check.v1 | v1.0.0-20201130134442-10cb98267c6c | tidak tersedia pada output |
| gopkg.in/yaml.v3 | v3.0.1 | tidak tersedia pada output |

Graph tidak memuat x/net. Ini tidak menyatakan semua package/runtime exposure telah diaudit. x/mod dan x/tools terpilih dalam graph; metadata Go minima keduanya tidak tersedia pada output ini, dan review source/compatibility lengkap belum tercapai. Maximum minimum yang tersedia adalah main Go1.26.9; pgx/text/sync menyatakan Go1.25.0. Compile terfokus sukses, bukan validasi perubahan protokol PostgreSQL.

Checksum metadata exact pins (`Sum`; `GoModSum`) dari download:

```text
pgx/v5 v5.9.2: h1:3ZhOzMWnR4yJ+RW1XImIPsD1aNSz4T4fyP7zlQb56hw= ; h1:mal1tBGAFfLHvZzaYh77YS/eC6IX9OWbRV1QIIM0Jn4=
x/text v0.41.0: h1:vz/seA0lnX87Othu2f/0L24RcgrXD9/YFTSuGjj3rH8= ; h1:jvf1O8ajNzZqhSrQBPbutR/EB83Cc0CFrezNQIwbb5M=
OIDC v3.16.0: h1:qRQUCFstKpXwmEjDQTIbyY/5jF00+asXzSkmkoa/mow= ; h1:wqPbKFrVnE90vty060SB40FCJ8fTHTxSwyXJqZH+sI8=
OAuth2 v0.34.0: h1:hqK/t4AKgbqWkdkcAeI8XLmbK+4m4G5YeQRrmiotGlw= ; h1:lzm5WQJQwKZ3nwavOZ3IS5Aulzxi68dUSgRHujetwEA=
jose/v4 v4.1.4: h1:moDMcTHmvE6Groj34emNPLs/qtYXRVcd6S7NHbHz3kA= ; h1:x4oUasVrzR7071A4TnHLGSPpNOm2a21K9Kf04k1rs08=
```

Temporary capture identities (bukan durable artifact terpisah; temporary directory tidak dijamin tersedia untuk Run berikutnya):

```text
06b8dabb343d09d03f1cfa708c9870e1f064151281c7db549cbd4f08bfadfa84  probe/go.mod
6923543aab63e09da2b894699afb70b2ff1072a2eb014e063c8542ffed7deaa0  probe/go.sum
e9e654184307ede9c43a45f10cc3a06a75d23d72d9a8f92a796cbde9a24beb01  probe/graph.json
c04210b43013a8833759174431f31fdc1ef799c6d3de2ce152e6f2c7df74f6a6  probe/vuln.json
```

Tagged-source/origin validation, relevant source diff, full transitive delta versus prior graph, missing Go minima, pgx query-mode/config/SCRAM/OpaqueString source audit, dan per-finding traces **belum lengkap**. Live `db.Open` memakai `pgxpool.New(ctx, databaseURL)` tanpa override query mode dalam source; konfigurasi runtime/password tidak dibuka. Tidak menyimpulkan configured simple-protocol atau SCRAM exposure. Bukti probe parsial tidak dapat menutup B-T2-002.

## Five-advisory disposition

Ranges berikut berasal dari approved Techplan, bukan advisory verification baru. Tidak ada finding/traces baru yang lengkap dari scanner Run ini.

| Advisory | Approved affected range / fixed version | Run disposition |
|---|---|---|
| GO-2026-4771 | pgx <v5.9.0 / v5.9.0 | Probe v5.9.2 di luar recorded range; scan/source gate gagal, belum closed |
| GO-2026-4772 | pgx <v5.9.0 / v5.9.0 | Sama; belum closed |
| GO-2026-5004 | pgx <v5.9.2 / v5.9.2 | Probe v5.9.2 di luar recorded range; query-mode/path verification belum lengkap |
| GO-2026-5970 | text <v0.39.0 / v0.39.0 | Probe v0.41.0 di luar recorded range; SCRAM/profile source verification belum lengkap |
| GO-2026-6629 | text <v0.41.0 / v0.41.0 | Probe v0.41.0 di luar recorded range; source verification belum lengkap |

Backend aktual masih pgx v5.8.0/text v0.29.0, sehingga actual graph clearance tidak didapat. Tidak ada residual-risk acceptance atau klaim absence of surviving/new advisories.

## Verification scope confirmation

Tidak ada race/concurrency, performance/load, final security-class suite, atau broad Testing-owned suite dijalankan. Required full `govulncheck` adalah explicit Build dependency gate di §10/§12; scan dicoba namun gagal, bukan final security verification. Tidak ada real PostgreSQL contention, real Google, configured auth/SCRAM, browser/HTTPS, migration/privilege test, atau exploit reproduction. Fake-clock session-lock check dan protocol/person-upsert checks belum ditulis/dijalankan karena graph gate gagal.

## Contract check

- [x] Bound hashes dan live anchors diperiksa; unrelated edits dipertahankan.
- [x] Stop-on-failed-gate dipatuhi sebelum tracked dependency/backend changes.
- [ ] Isolated D8 gate passed (scan dan source/compatibility evidence belum lengkap).
- [ ] Actual backend graph verified.
- [ ] Bounded G1 correction dan focused T2 checks completed.

T2/G1 tetap incomplete. Existing B-T2-001/B-T2-002/B-T2-003 tidak ditutup; tidak ada T3 readiness atau Work Unit completion claim.

## Phase handoff

- **Outcome:** BLOCKED — berhenti pada isolated D8 graph gate; G1 tidak dilanjutkan.
- **Result refs:** `RUN-C1-ENG-BUILD-T2-004/evidence/build-report.md` ini, khususnya F-T2-004 dan partial graph/check identities.
- **Findings:** F-T2-004 — required full source scan exit 1 saat package loading karena build-cache entries tidak ditemukan; root cause belum ditentukan, scan clearance tidak tersedia.
- **Decision requests:** Orchestrator/Anhar menentukan bounded scanner/environment reconciliation dan fresh Run scope; tidak ada permintaan waiver atau versi alternatif.
- **Blockers:** Failed/incomplete isolated scan gate menghalangi tracked D8 writes, expiry correction, commit, dan T2 completion. B-T2-002/B-T2-003 tetap open; task lain tidak dievaluasi/di-dispatch.
- **Open / unverified:** Complete isolated source/origin/compatibility/scan/traces dan Go minima; actual backend gate; post-lock expiry correction/focused checks; real provider/PostgreSQL/browser/independent Testing.
- **Recommended continuation:** Rekonsiliasi kegagalan scanner/cache, lalu fresh authorized Build mulai dari isolated D8 gate lengkap; actual backend verification dan G1 hanya setelah gate lolos. Jangan konsumsi interface T2 sebagai accepted dari report ini.
- **Context refs:** Invocation Run ini; Approved Techplan §§9–10/12 (hash di atas); current T2 snapshot; live `backend/internal/platform/auth/store.go:118–138`; prior T2-003 finding dan T2-002 dependency evidence.
