# Review findings — RUN-C1-ENG-CODEREVIEW-T2-001

- Phase / Role: Independent Code Review / Reviewer.
- Author: Codex; Created: 2026-10-11.
- Work Unit / Run: `WU-C1-ENG-001` / `RUN-C1-ENG-CODEREVIEW-T2-001`.
- Participant / Session (Invocation): `PARTICIPANT-C1-ENG-T2-CODEREVIEW-001` / `SESSION-C1-ENG-T2-CODEREVIEW-001`.
- Model / Reasoning: approval tercatat untuk `gpt-6-sol / medium`; identitas runtime tidak ditetapkan secara independen.
- Target revision / branch teramati: `059d145d5df64fcda427e5a625486a8e95ce1fd8` / `pilot/3-c1-engineering`.
- Workflow revision teramati: `91a97af8ba64f98b0a692a83fbff866a210c45c0`.
- Outcome: **STALLED pada preflight authority; empat pass belum dijalankan.**

## Preflight authority — R1 (blocking untuk Run ini)

Invocation §Review scope and baseline mewajibkan: “Stop and ask Orchestrator to re-ground the Run if a scoped file/hash or authority changes.” Enam file scope, target HEAD/branch, Approved Techplan, T2 task, manifest, laporan/gate/capture Build terikat, root/backend instructions, registry identity, dan tujuh sumber workflow/routing cocok. Namun tiga identity coordination authority pada §Bound authority and evidence tidak cocok:

| Authority | SHA-256 terikat Invocation | SHA-256 teramati |
|---|---|---|
| `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001.md` | `47b8db1cbbc3c09082f8e591f1553400447daf0c54c1d6203a26fd0e89b7ad22` | `8e704f034af45e0b3b7ad0e0fb5b8892c52e861eb1dfe366724962927ed1a42d` |
| `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/events.md` | `7c3a2ed42a7750067c333a9e79a557348c0ce292edb6d78bb6bf24ceb9433606` | `b60508f3fc3c0a3718b54d4afbf6fc51287c5919667591e0506f1adbd67739a1` |
| `docs/product/pilot-3-control-tower.md` | `f7abf20f887e30542bc8d6d3ae88d85ed25c2e975a828e97541602b83f336d4c` | `bd043b68a224b9e665790550e68f1219c190afc734c66fa1ce37c420552f4b63` |

Events terakhir mencatat model approval dan READY_FOR_HUMAN_DISPATCH; WU serta Control Tower juga memproyeksikan approval/launch pending. Ini kemungkinan pembaruan lifecycle setelah preparation, tetapi reviewer tidak memiliki pengecualian untuk mengganti binding sendiri. Perbedaan isi terhadap snapshot lama belum diverifikasi karena snapshot tidak disediakan. Tidak ada temuan cacat production atau perubahan Product yang disimpulkan dari drift ini.

Resolution / owner: Orchestrator memeriksa delta dan merekonsiliasi binding tiga authority tersebut dengan state current yang sah, lalu menentukan Invocation/Run/Session lanjutan sesuai overlay. Tidak memerlukan patch production; `patch-plan-1.md` tidak dibuat. Approval model sebelumnya tidak diminta ulang oleh reviewer.

## 1. Safety

Belum dijalankan: stop condition preflight R1 terjadi sebelum inspeksi semantik current diff. Tidak ada kesimpulan Safety.

## 2. Quality

Belum dijalankan karena R1. Tidak ada kesimpulan Quality.

## 3. Stack-Specific Best Practices

Belum dijalankan karena R1. Sumber routing dibaca/diperiksa identity; matching guidance belum diterapkan ke kode. Tidak ada kesimpulan best practices.

## 4. Consistency

Belum dijalankan karena R1. Root/backend AGENTS dan workflow Review dibaca sebagai authority preflight; tidak ada kesimpulan konsistensi implementasi.

## Verification executed during Review

Verified melalui read-only `git rev-parse HEAD`, `git branch --show-current`, `git status --short --untracked-files=all`, `sha256sum`, dan Python `hashlib.sha256`: baseline/branch serta 28 identity yang tercantum di tabel berikut cocok dengan Invocation. Pemeriksaan SHA-256 tiga authority di atas mengonfirmasi drift. Semua identity tersebut dan scope status diperiksa kembali segera sebelum penulisan evidence terminal. Workflow revision teramati dicatat di provenance.

| Input cocok | SHA-256 |
|---|---|
| `backend/go.mod` | `90b7c1c6e870e32828fa5932a6c4e9a28a65d0671a74bdb807ec29dbc5d27cae` |
| `backend/go.sum` | `6923543aab63e09da2b894699afb70b2ff1072a2eb014e063c8542ffed7deaa0` |
| `backend/internal/platform/auth/store.go` | `eb70175ebe756698a3a83b627e0fd6606928d8ac4c355f9dd8a05dab705552b3` |
| `backend/internal/platform/auth/store_test.go` | `bac0d93f0aa6af1591b10dc9a1df81fba0bf77ea2bdb5a165c129e2611795d60` |
| `backend/internal/platform/auth/oidc_test.go` | `2ffd30f781b3aeaa09f6beb0fddd00bea90abcab99b221e126867acfc8331ab4` |
| `backend/internal/platform/auth/store_integration_test.go` | `1b38058a177e0df5f87cf6c732478369e8f798265f2ad91c0762a1ea1879b1e0` |
| `AGENTS.md` | `d6bb74b1ffa8cbd1a115006191a40574b5f16508d22477d38e09744681367462` |
| `backend/AGENTS.md` | `817d01ab80598637644c64be9bfd7ec18a9dc5e3f0d22485d505f0760f72769f` |
| `.harscode-spaces/.local-config.yaml` | `ffbd0521361eac90e82dbee98ba59f01f3e62641afd79bf8e71c4fc8a0e633ce` |
| `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.md` | `6d2e41e177e4adaf74f53a06d81dd568def75019c8dc4e4792d938a869896d79` |
| `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/tasks/T2-backend-identity-session.md` | `dfe17b7a3dc0dc89127b040a67f547506a14a4ab591ed0a6343f1a8dd8bcce37` |
| `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/tasks/manifest.md` | `2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d` |
| `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/build-report.md` | `4fc08a8cb5533ebf702d364c86be11ce06b0937837ed8941cac828d5e3436191` |
| `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/build-report.md` | `4a63608a42dd89c9978c62161b0f6924700115fd0ed06eeed9af42aa04d6b326` |
| `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/authorization.md` | `3c0b06c256b773124fcb9af01207eadc570d90084a1577c551d5a783b1aabb98` |
| `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated-gate.md` | `0bbe936da2225601f2e9253d6dc08d432f6044c92da3009109974e99de4da012` |
| `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual-gate.md` | `c3a7f83133ae9e016c1451a9408bab66d3c8fff9492cde608e8177de9e19c063` |
| `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/changed-file-hashes.json` | `5daa2a9b25d13b706bdc8aa5f78ab04894e7cb8523fd031c6661c910328a404d` |
| `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/capture-hashes.json` | `f5913ccb140311022a487704fb57f9e3025993bde086eced2e9af45658c4cad7` |
| `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/final-source-hashes.json` | `9db9b5239ef9be3361d705e2d518d05b895e270579928e2b2c0896ce6d5a1ecc` |
| `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/final-diff.patch` | `b26f0a28185d57acff16c668240fd7b20ea3aacaaf23b53d2d21a0be36e84ea4` |
| `../harscode-workspace/workflow/4-code-review-prompt.md` | `d152c1419ea8b0e67a08d1ec2cec7c0235cc15b0fccbd032722aea1ed2332f8a` |
| `../harscode-workspace/workflow/4-code-review/guidelines.md` | `513f145a5bc1980589a4967f6a0894a3f59f089ade4fa199ba252b2604a735c0` |
| `../harscode-workspace/workflow/4-code-review/checklist.md` | `5621854c75f8834ff9c32a8a67a326ab03388e11cceac95f0f2c6bdb0e03e079` |
| `../harscode-workspace/workflow/orchestrated-run-overlay.md` | `f806673aa5e7d1d82dae9f433ec54c3562e15ce778b709a141c57d3d22dfaf87` |
| `../harscode-workspace/best-practices/AGENTS.md` | `9dc2b38e5175f4da7e21cbd7c78e4e8aaea07d2325510f24cd9527625980fc6e` |
| `../harscode-workspace/best-practices/index.md` | `876a14da409b3a486a790daee542fbafd02ecae82283b892dd189023cfbaaa34` |
| `../harscode-workspace/best-practices/go/index.md` | `0b51468a6cf4bdc1037422376ef2aeb6884898e0d4a38c9a79b594f4f10bebf6` |

Full changed-file scope yang teramati sebelum evidence ini dibuat:

```text
 M .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001.md
 M .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/events.md
 M AGENTS.md
 M backend/AGENTS.md
 M backend/go.mod
 M backend/go.sum
 M backend/internal/platform/auth/store.go
 M docs/product/pilot-3-control-tower.md
 M docs/project/go-dependency-verification.md
 M docs/project/go-scanner-runbook.md
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/baseline.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/baseline.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/build-report.md
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/capture-hashes.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/commands.jsonl
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/edges.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/edges.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/final-boundary.txt
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/focused-compile.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/focused-compile.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/go-env.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/go-env.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/go-version.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/go-version.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/graph.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/graph.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/integrity.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/integrity.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/inventory.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/isolated.go.mod
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/isolated.go.sum
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/legacy-provenance.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/legacy-source-match.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/legacy-source.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/origin-github.com_davecgh_go-spew.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/origin-github.com_davecgh_go-spew.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/origin-github.com_jackc_pgpassfile.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/origin-github.com_jackc_pgpassfile.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/origin-github.com_kr_pretty.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/origin-github.com_kr_pretty.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/origin-github.com_pmezard_go-difflib.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/origin-github.com_pmezard_go-difflib.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/origin-github.com_stretchr_objx.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/origin-github.com_stretchr_objx.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/origin-gopkg.in_check.v1.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/origin-gopkg.in_check.v1.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/origin-gopkg.in_yaml.v3.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/origin-gopkg.in_yaml.v3.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/pins.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/pins.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/preflight.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/provenance.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual-gate.md
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/actual-identity-check.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/advisory-dispositions.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/baseline.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/commands-context.jsonl
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/commands.jsonl
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/compatibility-review.md
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/edges.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/edges.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/final-edges.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/final-edges.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/final-go-env.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/final-go-env.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/final-graph.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/final-graph.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/final-integrity.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/final-integrity.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/final.go.mod
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/final.go.sum
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/fixture-cleanup.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/fixture-cleanup.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/fixture-context.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/fixture-image.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/fixture-image.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/fixture-port.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/fixture-port.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/fixture-stop.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/fixture-stop.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/fixture-version.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/fixture-version.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/focused-compile.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/focused-compile.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/focused-postgres.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/focused-postgres.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/focused-unit-final.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/focused-unit-final.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/focused-unit.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/focused-unit.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/graph.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/graph.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/integrity.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/integrity.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/inventory.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/legacy-source-match.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/missing-sums.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/pins.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/pins.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/post-g1-advisory-dispositions.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/post-g1-graph.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/post-g1-graph.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/post-g1-integrity.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/post-g1-integrity.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/post-g1-scan-summary.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/post-g1-source-scan.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/post-g1-source-scan.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/post-tests-tidy.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/post-tests-tidy.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/regression-control.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/regression-original.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/regression-original.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/regression-patched.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/regression-patched.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/scan-summary.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/source-identities.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/source-scan.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/source-scan.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/tidy.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/tidy.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/actual/upstream-validation.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/authorization.md
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/build-report.md
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/capture-hashes.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/changed-file-hashes.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/final-diff.patch
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/final-report.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/final-source-hashes.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/final-status.txt
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/fixture-check.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/inventory.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated-gate.md
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/advisory-dispositions.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/baseline.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/baseline.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/capture-hashes.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/commands-context.jsonl
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/commands.jsonl
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/commit-github.com_jackc_pgservicefile.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/commit-github.com_jackc_pgservicefile.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/commit-gopkg.in_check.v1.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/commit-gopkg.in_check.v1.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/compatibility-review.md
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/delta.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/edges.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/edges.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/focused-compile.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/focused-compile.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/go-env.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/go-env.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/go-version.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/go-version.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/graph.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/graph.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/integrity.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/integrity.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/inventory.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/inventory.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/isolated.go.mod
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/isolated.go.sum
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/legacy-provenance.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/legacy-source-match.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/legacy-source.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/missing-sums.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/missing-sums.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/origin-github.com_davecgh_go-spew.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/origin-github.com_davecgh_go-spew.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/origin-github.com_jackc_pgpassfile.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/origin-github.com_jackc_pgpassfile.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/origin-github.com_kr_pretty.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/origin-github.com_kr_pretty.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/origin-github.com_pmezard_go-difflib.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/origin-github.com_pmezard_go-difflib.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/origin-github.com_stretchr_objx.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/origin-github.com_stretchr_objx.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/origin-gopkg.in_check.v1.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/origin-gopkg.in_check.v1.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/origin-gopkg.in_yaml.v3.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/origin-gopkg.in_yaml.v3.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/pgx-source.diff
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/pins.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/pins.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/preflight.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/provenance.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-cloud.google.com_go_compute_metadata.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-cloud.google.com_go_compute_metadata.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-github.com_coreos_go-oidc_v3.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-github.com_coreos_go-oidc_v3.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-github.com_go-jose_go-jose_v4.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-github.com_go-jose_go-jose_v4.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-github.com_jackc_pgx_v5.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-github.com_jackc_pgx_v5.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-github.com_jackc_puddle_v2.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-github.com_jackc_puddle_v2.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-github.com_joho_godotenv.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-github.com_joho_godotenv.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-github.com_stretchr_testify.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-github.com_stretchr_testify.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-golang.org_x_mod.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-golang.org_x_mod.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-golang.org_x_oauth2.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-golang.org_x_oauth2.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-golang.org_x_sync.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-golang.org_x_sync.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-golang.org_x_text.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-golang.org_x_text.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-golang.org_x_tools.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/ref-golang.org_x_tools.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/run.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/scan-review.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/scan-summary.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/scanner-version.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/scanner-version.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/source-excerpts.txt
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/source-identities.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/source-review.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/source-scan.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/source-scan.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/text-source.diff
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/tidy.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/tidy.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/upstream-github.com_davecgh_go-spew.tar.gz
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/upstream-github.com_davecgh_go-spew.tar.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/upstream-github.com_jackc_pgpassfile.tar.gz
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/upstream-github.com_jackc_pgpassfile.tar.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/upstream-github.com_kr_pretty.tar.gz
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/upstream-github.com_kr_pretty.tar.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/upstream-github.com_pmezard_go-difflib.tar.gz
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/upstream-github.com_pmezard_go-difflib.tar.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/upstream-github.com_stretchr_objx.tar.gz
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/upstream-github.com_stretchr_objx.tar.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/upstream-gopkg.in_check.v1.tar.gz
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/upstream-gopkg.in_check.v1.tar.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/upstream-gopkg.in_yaml.v3.tar.gz
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/upstream-gopkg.in_yaml.v3.tar.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/upstream-validation.json
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/isolated/upstream.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/legacy-source.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/provenance.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/regression-control.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/run.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/scan-review.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/source-review.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1/upstream.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/run.py
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/scanner-version.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/scanner-version.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/tidy-escalated.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/tidy-escalated.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/tidy.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/tidy.stdout
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/upstream-github.com_davecgh_go-spew.tar.gz
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/upstream-github.com_davecgh_go-spew.tar.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/upstream-github.com_jackc_pgpassfile.tar.gz
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/upstream-github.com_jackc_pgpassfile.tar.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/upstream-github.com_kr_pretty.tar.gz
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/upstream-github.com_kr_pretty.tar.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/upstream-github.com_pmezard_go-difflib.tar.gz
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/upstream-github.com_pmezard_go-difflib.tar.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/upstream-github.com_stretchr_objx.tar.gz
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/upstream-github.com_stretchr_objx.tar.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/upstream-gopkg.in_check.v1.tar.gz
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/upstream-gopkg.in_check.v1.tar.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/upstream-gopkg.in_yaml.v3.tar.gz
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/upstream-gopkg.in_yaml.v3.tar.stderr
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/invocation.md
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-CODEREVIEW-T2-001/invocation.md
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-LOOPHEALTH-002/evidence/loop-health-audit.md
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-LOOPHEALTH-002/evidence/orchestrator-reconciliation.md
?? .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-LOOPHEALTH-002/invocation.md
?? backend/internal/platform/auth/oidc_test.go
?? backend/internal/platform/auth/store_integration_test.go
?? backend/internal/platform/auth/store_test.go
?? docs/project/go-network-recovery.md
```

Tracked production/dependency scope tetap `backend/go.mod`, `backend/go.sum`, `backend/internal/platform/auth/store.go`; tiga test baru tetap untracked. Perubahan instructions/runbook/scanner/Control Tower/orchestration dikecualikan dari implementation review sesuai Invocation. Tidak ada full diff/code review, runtime check, test, build, scanner, dependency fetch, browser/provider atau PostgreSQL execution. Tidak ada matriks Build/Testing yang dijalankan ulang. Capture Build hanya diverifikasi identity; kebenaran claim runtime belum ditelaah.

## Verdict

**Tidak diterbitkan — STALLED pada authority preflight R1.** Empat pass belum selesai; hasil ini bukan Approve, Request changes pada kode, atau acceptance T2. B-T2-002/B-T2-003 tetap tidak ditutup; readiness T2/T3 tidak ditetapkan.

## Phase handoff

- Outcome: STALLED — identity authority tidak cocok sebelum empat-pass Review.
- Result refs: `evidence/review-findings-1.md` (Run-local; carrier terminal ini).
- Findings: R1, drift SHA-256 Work Unit / Events / bounded Control Tower; tidak ada code finding yang ditetapkan.
- Decision requests: Orchestrator re-ground binding authority dan menentukan route Review berikutnya; Human semantic/security decision baru belum teridentifikasi.
- Blockers: Review enam file pada Run ini terhenti pada R1; tidak menyimpulkan blocker baru untuk seluruh WU.
- Open / unverified: seluruh empat pass; seluruh runtime/independent Testing dan acceptance yang sudah pending; identitas runtime model belum terverifikasi.
- Recommended continuation: Orchestrator reconcile R1 terlebih dahulu, kemudian dispatch Review independen dengan authority current yang konsisten; jangan dispatch Testing dari hasil ini.
- Context refs: `invocation.md` §Review scope and baseline / §Bound authority and evidence; R1 dan verification table di atas; `workflow/orchestrated-run-overlay.md` §Chronology and re-entry.
