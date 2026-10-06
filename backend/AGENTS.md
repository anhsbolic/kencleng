# AGENTS.md — backend/

Backend-specific rules on top of root `AGENTS.md`.

Scope: `backend/`.

## Baseline posture

This backend is a neutral engineering substrate. No historical domain, route, migration, security mechanism, or implementation is current merely because it exists in Git history.

## Authority routing

- product/business meaning → `../docs/product/product-intent.md`;
- backend architecture → `../docs/project/kencleng-backend-tech-stack.md`;
- reusable product-design concerns when relevant → `../docs/ui-ux/README.md`;
- API contract → `../api/` only after an active contract exists;
- generic lifecycle / engineering practice → current Harscode authority.

Release scope and delivery order are intentionally open. Do not infer them from historical code/specs/contracts.

## Structure

```text
backend/
├── cmd/server/main.go
├── internal/
│   ├── domain/          # created only when real delivery needs a product domain
│   └── platform/db/     # neutral database connectivity
└── migrations/          # created by active delivery needs
```

Do not create empty architecture layers merely to make the repository look complete.

## Engineering rules

- Use `net/http` unless demonstrated need justifies otherwise.
- Preserve error chains with `%w`; do not swallow expected failures.
- Parameterize SQL; never interpolate user-controlled values into SQL.
- Do not use `float64` for money.
- Do not log secrets, raw tokens, or unnecessary PII.
- Backend authorization, when introduced, is explicit; frontend gates are not security authority.
- Choose security mechanisms from the active risk/data semantics, not historical implementation inertia.
- Simulation/fixtures must not be presented as real settlement, verification, provenance, or external evidence.

## Testing and evidence

Use unit, integration, race, contract, and security checks proportionally to the active work. Never claim a verification command ran unless it actually ran.

## Write boundaries

A backend-scoped implementation should not casually modify frontend production code. Shared docs/API changes require explicit coordination with the concern they own.
