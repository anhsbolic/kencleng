# Kencleng API Contract Tooling

> Status: clean delivery baseline

`api/` owns the coordinated C1 OpenAPI contract. `openapi/index.yaml` is the only editable source; `openapi.yaml` and `openapi.d.ts` are generated outputs. The contract describes implementation obligations and is not runtime verification or authorization evidence.

## Structure

```text
api/
├── openapi/
│   └── index.yaml      # editable source root
├── openapi.yaml        # generated bundled view
├── openapi.d.ts        # generated shared TypeScript types
├── package.json
└── package-lock.json
```

## Commands

```bash
npm install
npm run validate
npm run bundle
npm run types
```

`npm run types` uses the locked `openapi-typescript` installation declared by `frontend/package.json` and emits types here so consumers use the shared generated contract. Run `npm ci` in `frontend/` before it when dependencies are not installed. Do not edit either generated output by hand.
