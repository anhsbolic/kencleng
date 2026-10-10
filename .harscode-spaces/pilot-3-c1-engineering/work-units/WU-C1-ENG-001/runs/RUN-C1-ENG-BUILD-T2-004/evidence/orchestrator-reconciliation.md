# Orchestrator reconciliation — RUN-C1-ENG-BUILD-T2-004

Date: 2026-10-11. Work Unit: `WU-C1-ENG-001`.

## Terminal Run

The terminal Build report is preserved unchanged at `evidence/build-report.md`, SHA-256 `eff0a3b9c2f0392bbbcaf0f2aa33cf0b93d872afe3eb8b8573f6cef2a81ed051`. It contains exactly one `## Phase handoff`. All bound input identities reported by the Run match its Invocation. Outcome remains **BLOCKED** at the isolated D8 gate: `govulncheck -json ./...` could not load packages from missing Go build-cache entries. The report is not edited to incorporate later diagnosis.

Anhar confirmed in this conversation that `gpt-6-sol` / `medium` was approved for this exact Run and that it was launched. Invocation lifecycle metadata now records this approval and dispatch. The report says runtime model/effort was not independently exposed, so runtime identity remains unverified.

## Separate scanner diagnosis evidence

After T2-004 terminated, the scanner/tooling diagnosis was completed outside the Participant Run. The temporary source `/tmp/c1-scanner-diagnostic-x9_cedbf/diagnosis.md` was copied byte-for-byte to `evidence/scanner-diagnosis/diagnosis.md`; both SHA-256 values are `07c0330674fb1b43f1d15596e77c2af4a60e3a3260158221ce4c5b590175b78b`. The actual-backend scan capture was likewise preserved byte-for-byte; both source and copy hash to `99219c4bb8e965eb8601d86e3ecb13c437019b2497bd8fb733c8c05f9fd471e9`. Scanner stdout/stderr and supporting isolated-probe module inputs are in the same directory. The terminal Build report remains the only Participant-authored Run artifact; these are explicitly post-Run Orchestrator-preserved diagnostic captures.

The diagnosis attributes the original failure to environment constraints: the sandbox could not write the default Go build cache, and sandbox network access to `vuln.go.dev` was denied. With writable cache and authorized network access, scanner v1.7.0 completed successfully; the approved D8 probe returned zero finding objects, while the actual backend graph returned 15 finding objects across GO-2026-4771, GO-2026-4772, GO-2026-5004, GO-2026-5970, and GO-2026-6629. This resolves the scanner execution-method uncertainty. It does **not** complete the isolated D8 gate: tagged-source/origin validation, full transitive delta and Go-minimum review, pgx query-mode/config and SCRAM/OpaqueString source-path checks, and complete per-finding traces/dispositions remain incomplete. It also does not clear actual-backend advisories or B-T2-002.

No backend/dependency files, scanner inputs in the repository, or T2 implementation were changed by this reconciliation. No scanner, build, or test command was rerun. T2-003's source-derived F-T2-003 and B-T2-003 remain open; no PostgreSQL reproduction or post-lock expiry correction is established.

The diagnosis also reports a local `docs/project/go-scanner-runbook.md` and additions to root/backend `AGENTS.md`. Those are post-Run working-tree documentation/guidance changes, outside T2-004's `ARTIFACT_TARGET`; they are not represented as Participant output or as Build-report edits. Their content is bounded to scanner operation and does not change approved dependency pins, Product/Techplan authority, or gate disposition.

## Routing

The meaningful delta since T2-004 is the completed scanner diagnosis and preserved captures. A fresh Build re-entry is prepared as `RUN-C1-ENG-BUILD-T2-005`. It must complete every isolated D8 evidence gate using the current target instructions and scanner runbook before any tracked dependency write or G1 continuation. Only if that gate passes may the Run verify the actual backend graph, apply the already-approved bounded expiry correction, and run focused T2 checks. B-T2-002 and B-T2-003 remain open until their separate evidence conditions pass; no commit, T3 readiness, or T2 completion is established.
