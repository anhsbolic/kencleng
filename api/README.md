# Kencleng API Contract Tooling

> Status: clean delivery baseline

`api/` provides OpenAPI authoring, validation, and bundling capability. The baseline intentionally defines **no product operations**. Add contract surface only when an authorized delivery requirement is sufficiently clear.

Reusable API design, minimal success response envelope, RFC 9457 errors, and verification conventions are owned by [T06 — API & Contract Standards](../docs/project/standards/api-contract.md). This tooling document does not establish any new API operation.

## Structure

```text
api/
├── openapi/
│   └── index.yaml      # editable source root
├── openapi.yaml        # generated bundled view
├── package.json
└── package-lock.json
```

## Setup and commands

Gunakan Node.js **24.22.1** melalui nvm, sesuai `.nvmrc` root dan `engines`
di `package.json`. Dari direktori `api/`, dengan nvm sudah dimuat di shell:

```bash
nvm use
npm ci
npm run validate
npm run bundle
```

`npm ci` memasang dependency sesuai lockfile. Tooling ini tidak menjalankan
service/API server dan tidak memerlukan Caddy, PostgreSQL, atau object storage.
Untuk menjalankan aplikasi, ikuti [setup repo](../docs/project/kencleng-repo-setup.md),
[backend](../backend/README.md), dan [frontend](../frontend/README.md).

Do not reconstruct endpoints from historical contracts by inertia. Historical OpenAPI remains available in Git history as evidence only.
