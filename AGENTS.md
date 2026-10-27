# AGENTS.md — Kencleng

Root instruction map and cross-repository hard rules for AI coding agents.

Read this file before writing code. More-specific `AGENTS.md` files apply within their directory scope.

## Authority routing

```text
whole-product purpose / actors / durable product truth
→ docs/product/product-intent.md

reusable product-design / UX / visual guidance
→ docs/ui-ux/README.md

repository topology/setup
→ docs/project/README.md

project-wide and stack-specific Engineering Standards
→ docs/project/standards/README.md

backend/frontend architecture
→ docs/project/standards/backend.md
→ docs/project/standards/frontend.md

shared API contract
→ api/ (only after an active contract exists)

stack-specific execution
→ backend/AGENTS.md
→ frontend/AGENTS.md

generic lifecycle / engineering practice
→ current Harscode authority
```

Release scope, delivery order, detailed delivery requirements, API operations, and implementation state are intentionally open at this baseline **unless a current Human-approved scoped artifact explicitly establishes them for a bounded active commitment**. Do not infer them from historical specs, contracts, code, migrations, handoffs, or Git history.

Existing historical material may be useful evidence when an active decision/work item calls for it. It is not automatic authority.

## Local development

- Setup/topology authority: `docs/project/kencleng-repo-setup.md`; quick starts: root, backend and frontend `README.md`.
- Default development runs FE and BE on the host, with PostgreSQL in Compose. Optional app-container configuration is in `docker-compose.apps.yml`; runtime compatibility/hot reload still require actual evidence. Follow `docs/project/local-development-modes.md`.
- Activate Go 1.27.3 with gvm using `.go-version` and Node.js 24.22.1 with `nvm use` using `.nvmrc` before setup or verification.
- App setup and run belong to each app directory: `make setup` and `make run` from `backend/` or `frontend/`. From repo root, start infrastructure with `make COMPOSE=podman-compose up` (Podman) or `make up` (Docker Engine/Desktop).
- Run `make gvm-run` from `backend/` and `make nvm-run` from `frontend/` in separate terminals to activate each pinned runtime and start its runner. Plain `make run` uses the already-active runtime. Browser entry is `http://localhost:3000`; development API requests use `/api/*` through Next.js, without Caddy.
- From either app directory, `make docker-run COMPOSE=podman-compose` starts that app in a dev container; omit the provider override for Docker Engine/Desktop. Container apps use project `kencleng-apps`, sharing the existing database network. Stop the host apps before switching modes; never reset data just to switch modes. `docker-down` from root removes app containers while preserving database/cache volumes.
- `frontend/: make check-hot-reload` is an opt-in verification that temporarily edits/restores the health handler and FE page. It requires both apps/browser tooling, refuses non-loopback targets, and must not run concurrently with edits to those files. Mock/static checks do not establish actual reload readiness.
- Compose provider selection is explicit through `COMPOSE`; do not silently select or switch engines. Each engine owns its own containers and data volumes.
- Preserve the PostgreSQL image OS family when reusing data: a prior volume reported collation 2.41, and the pinned image uses Trixie. `make up` checks recorded versus actual database collation; inspect current values rather than assuming a reset/new volume has historical metadata. Do not silence a mismatch by refreshing metadata without assessing affected objects.
- Optional S3 storage is SeaweedFS in `docker-compose.storage.yml`: root `make storage-up`, `storage-check`, and `storage-persistence-check` with explicit `COMPOSE`. The last check restarts only SeaweedFS; checks create/delete their own UUID-scoped probe objects. See `docs/project/local-object-storage-options.md` for endpoints, image evidence, and deployment requirements. Human reported these Podman commands completed without error on 24 October 2026; logs were not attached. Do not present this as agent-executed evidence or Docker/production verification.
- Legacy MinIO remains in `docker-compose.minio.yml`; it is not selected by `storage-*`. Never mount its volume into SeaweedFS or infer format compatibility from S3 compatibility.
- Optional storage uses its own Compose project (`kencleng-storage`) to avoid coupling PostgreSQL recreation to storage configuration changes. Never silently attach/migrate a legacy storage volume.
- SeaweedFS is selected for optional local preparation. Storage is not required for the scaffold; production deployment/readiness is a separate decision with evidence.
- Setup preserves existing local env files. Do not reset databases, remove volumes, overwrite credentials, or stop unrelated containers as setup cleanup.
- For a separately requested Podman reset, `scripts/reset-kencleng-podman.py` previews targets by default. `--apply` requires typed confirmation; `--delete-data` additionally destroys selected volumes. Creating/providing the script is not permission to execute it. Never substitute global prune/reset commands.

## Golden rules

- Handle expected failures intentionally; do not silently swallow errors.
- Never introduce binary floating-point arithmetic for money.
- Parameterize SQL; never interpolate user-controlled values into SQL.
- Never log secrets, raw tokens, or unnecessary PII payloads.
- Client-facing failures must not leak raw SQL, stack traces, filesystem paths, secrets, or other implementation internals.
- Backend authorization, when introduced, is explicit; frontend role gates are UX, not security authority.
- Simulation, fixtures, or operator assistance must not be presented as real settlement, independent verification, provenance, authorization, or real-world evidence.
- If product meaning is materially missing, surface the decision gap instead of inventing it downstream.

## Authority separation

Product intent owns whole-product meaning. Product-design authority owns reusable experience/visual guidance. Engineering architecture owns technical constraints. Human-approved commitment-specific product behavior and requirements may be **binding within that bounded commitment** without becoming whole-product canonical Product Authority. Detailed contracts and implementation are derived only when active delivery work needs them.

An implementation agent must follow the current scoped owning artifacts identified by the active handoff and must not change higher-level requirements merely to make code pass. If authority appears wrong or incomplete, report the contradiction/gap and route it to the owner of that concern.

## Scope and write boundaries

Work on one coherent unit at a time. Backend and frontend production writes remain separate by default. Shared `docs/` or `api/` changes require explicit coordination with the concern they own.

A small task unexpectedly touching many unrelated files is a risk signal; investigate before expanding scope.

## High-risk write boundary

Explicit Human authorization is required before an agent introduces or materially changes implementation that owns any of these concerns, unless the current authorized work already records that approval for the same surface:

- balance, ledger, settlement, disbursement, or other money-movement transaction/locking logic;
- encryption, HMAC, signing-key, or other cryptographic key-handling core logic;
- authentication token/session/MFA core security logic;
- privileged authorization rules whose failure could grant or exercise materially elevated access.

This is a semantic safety boundary, not a commitment to historical file paths or architecture. Agents may inspect, critique, test, or propose changes without treating old implementations as current precedent.

## Security scrutiny

Security gets proportional scrutiny. Do not copy historical authentication, encryption, authorization, or data-protection mechanisms into a new active path merely because they existed before. Reuse only when current data semantics and risk justify it.

## Evidence

Distinguish verified, assumed, deferred, and not tested. Never claim a check ran when it did not.

Git history and prior Harscode artifacts are evidence/history. Promote reusable truth into the current source that owns the concern rather than treating process artifacts as project-wide precedent.
