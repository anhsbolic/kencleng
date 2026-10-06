# Kencleng API Contract Tooling

> Status: clean delivery baseline

`api/` provides OpenAPI authoring, validation, and bundling capability. The baseline intentionally defines **no product operations**. Add contract surface only when an authorized delivery requirement is sufficiently clear.

## Structure

```text
api/
├── openapi/
│   └── index.yaml      # editable source root
├── openapi.yaml        # generated bundled view
├── package.json
└── package-lock.json
```

## Commands

```bash
npm install
npm run validate
npm run bundle
```

Do not reconstruct endpoints from historical contracts by inertia. Historical OpenAPI remains available in Git history as evidence only.
