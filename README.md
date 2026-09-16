# Kencleng

Kencleng is a sandbox donation/crowdfunding platform for Indonesian non-profit organizations — campaign curation, guest donation, progress tracking, disbursement workflows, and fund-usage reporting.

**This is a learning project, not a production system.** It exists to practice:

- Go and Next.js development through a realistic, non-trivial domain;
- concurrency-safe, correctness-critical backend patterns;
- secure-by-design coding across a full stack, using a financial/donation domain as a natural showcase for common pitfalls;
- product/design-to-engineering translation;
- an agentic development workflow where implementation is guided by durable specs, architecture, design authorities, review, testing, and explicit human decisions;
- preserving representative workflow evidence so other people can inspect the learning process.

If you're here to fork, study, or reuse parts of this for your own purposes — go ahead. That's the intent.

## Stack

- **Backend**: Go (`net/http` stdlib, Go 1.22+ pattern routing), PostgreSQL, `pgx`, `goqu`, `golang-migrate`, MinIO (S3-compatible)
- **Frontend**: Next.js App Router, React, TypeScript, Tailwind CSS v4, TanStack Query, React Hook Form/Zod, optional Zustand for justified shared client state, Vitest/RTL/MSW, Playwright capability
- **API contract**: split OpenAPI source under `api/openapi/` with `api/openapi.yaml` as the bundled aggregate used where a full generated view is needed
- **Local orchestration**: Docker Compose + Caddy (reverse proxy)

Full rationale and active decisions live under `docs/`.

## Current frontend generation

The previous frontend product/UI implementation is intentionally being retired so the next generation can start from the current canonical design authority:

```text
Sunlit Editorial
→ Evidence-Led Optimism
```

The reboot retains the engineering scaffold/tooling but removes old route UI, components, implementation-specific state/API plumbing, mocks/tests, and old active agent-work history that could accidentally become precedent.

See:

```text
docs/project/frontend-reboot-plan.md
docs/project/kencleng-frontend-tech-stack.md
docs/ui-ux/README.md
```

The next real frontend feature-development run begins only after the reboot baseline is verified and frozen.

## Repository structure

```text
kencleng/
├── docs/          # specs, architecture, UI/UX authority, project/workflow docs
├── api/           # OpenAPI source + bundled aggregate
├── backend/       # Go module
├── frontend/      # Next.js frontend scaffold / implementation
├── AGENTS.md      # root rules/router for AI coding agents
└── docker-compose.yml
```

See `docs/kencleng-repo-setup.md` for repository setup/rationale and `docs/kencleng-agentic-workflow.md` for Kencleng-specific project orchestration. Generic per-feature lifecycle is owned by Harscode.

## Running locally

### Prerequisites

- Podman + Podman Compose
- Go 1.22+
- Node.js 20+
- `openssl` (for generating keys)

> **Note**: infrastructure such as Postgres, MinIO, and Caddy runs in containers. Backend/frontend can run natively on the host for development; follow current repository setup docs when exercising same-origin authentication/integration behavior.

### Setup

1. Clone the repo:

   ```bash
   git clone https://github.com/anhsbolic/kencleng.git
   cd kencleng
   ```

2. Generate an ES256 keypair (used to sign JWT access tokens):

   ```bash
   mkdir -p keys
   openssl ecparam -genkey -name prime256v1 -noout -out keys/es256-private.pem
   openssl ec -in keys/es256-private.pem -pubout -out keys/es256-public.pem
   ```

3. Copy the env template and fill in the generated secrets:

   ```bash
   cp .env.example .env
   # ENCRYPTION_KEY and HMAC_KEY: generate each with `openssl rand -base64 32`
   ```

4. Start the infrastructure:

   ```bash
   make up-podman
   ```

5. Verify the infrastructure came up correctly before continuing.

6. Run database migrations when the active backend workflow requires them:

   ```bash
   cd backend && make migrate-up
   ```

7. Run backend/frontend as needed:

   ```bash
   cd backend && go run ./cmd/server
   cd frontend && npm install && npm run dev
   ```

For integrated authentication/session behavior, use the project proxy/origin setup described in the current repository setup/architecture docs rather than assuming a direct dev-server URL is equivalent.

### Verifying infrastructure

After `make up-podman`:

```bash
podman ps -a
```

Expected infrastructure includes Postgres, MinIO, and Caddy plus the one-shot MinIO initialization container where configured.

The MinIO Console is normally available at `http://localhost:9001` using local development credentials from the project's environment configuration.

### Verifying changes

Use the verification appropriate to the active work and current Harscode/Kencleng guidance. Repository-level commands such as:

```bash
make verify
```

may provide a broad sweep, but phase-specific ownership and risk-driven verification remain governed by the active workflow and scoped stack guidance.

Never claim checks that were not run.

## Process evidence

Kencleng intentionally commits representative agent/Harscode process artifacts under locations such as `frontend/.local-agents/`.

These artifacts exist for learning and historical inspection. They are **task evidence**, not automatic project-wide authority. Reusable truth must be promoted into the spec/design/architecture document that owns the concern.

## Contributing

This is a personal learning sandbox rather than a project actively seeking contributions, but issues, forks, and questions are welcome.

If you're exploring the agentic-workflow side of the repository, start with `AGENTS.md`, `docs/kencleng-agentic-workflow.md`, and the applicable Harscode workflow entrypoint.

## License

TBD.
