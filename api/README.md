# Kencleng API Spec

> File: `api/README.md`

## Structure

```
api/
  openapi.yaml          # GENERATED — bundled single-file spec, used by
                         # openapi-typescript (frontend) and as the
                         # canonical reference for backend handlers.
                         # Do not edit directly.
  openapi/
    index.yaml           # Root: openapi/info/servers/tags/security +
                          # one $ref per path into the owning domain file.
    common.yaml           # Shared across all domains: securitySchemes,
                          # parameters (CursorParam, LimitParam,
                          # AdminUsersLimitParam, IdempotencyKeyHeader),
                          # responses (ValidationError, Unauthorized,
                          # Forbidden, NotFound, Conflict,
                          # TooManyRequests), and shared envelope schemas
                          # (Problem, ValidationProblem, Pagination).
    account.yaml
    organization.yaml
    campaign.yaml
    donation.yaml
    disbursement.yaml
    notification.yaml
  package.json            # devDependency on @redocly/cli, "bundle" script
```

## Why split

`api/openapi.yaml` as a single hand-authored file grew past ~4,000
lines once all 6 domains were specced out — expensive for an agent (or
a human) to read in full on every task touching just one domain. Each
domain's paths + schemas now live in their own file
(`api/openapi/{domain}.yaml`), so a task scoped to `organization` only
needs to read `api/openapi/organization.yaml` (~740 lines) plus
`common.yaml` (~175 lines), not the whole spec.

## Editing workflow

1. Edit the relevant `api/openapi/{domain}.yaml` (or `common.yaml` for
   shared parameters/responses/envelope schemas) directly. Same-domain
   `$ref`s stay as `"#/components/schemas/X"`. Cross-file references
   (to `common.yaml`, or to another domain's schema in the rare case
   that's needed) use a relative path:
   `"./common.yaml#/components/schemas/Problem"`.
2. If you add or remove a path, also update `index.yaml`'s `paths:`
   list — one entry per path, `$ref`ing the owning domain file (see
   existing entries for the pattern). This is the one place that
   still needs a manual, mechanical update when the path list changes.
3. Regenerate the bundled single-file spec:
   ```
   cd api && npm install && npm run bundle
   ```
   This overwrites `api/openapi.yaml`. Commit both the source
   (`api/openapi/*.yaml`) and the regenerated bundle in the same
   change — same discipline as the existing "keep `openapi.yaml` in
   sync with the handler in the same commit" rule from
   `kencleng-backend-tech-stack.md`, just with one more generated
   artifact in the loop now.
4. `npm run validate` runs `redocly lint` against the split source
   directly (catches unresolved `$ref`s and basic spec issues without
   needing to bundle first).

## Known pre-existing issue (not introduced by this split)

32 places in the original single-file spec use a `description:` key as
a sibling of `$ref:` inside the same response object (e.g. a `"403"`
response with both a locally-written description *and* `$ref:
"#/components/responses/Forbidden"`). This isn't valid per the
JSON Reference semantics OpenAPI relies on — `$ref` is meant to be the
only key in its containing object — and tools resolve the ambiguity
differently; Redocly's bundler keeps the local `description` and drops
the `$ref` (losing the shared response's `content`/`example`).
Confirmed present in the source file before this split (not something
splitting introduced). Left as-is pending a decision on the intended
pattern — see chat history 2026-08-20.
