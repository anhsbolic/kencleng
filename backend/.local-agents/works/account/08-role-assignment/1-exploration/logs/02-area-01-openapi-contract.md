# Stage 2 — Area 1: OpenAPI Contract

## Current State

Three endpoints are fully defined in `api/openapi.yaml`:

**`GET /admin/users`** (lines 555-573):
- Parameters: `CursorParam` (uuid, optional), `AdminUsersLimitParam` (int, 1-20, default 20)
- Response 200: `UserListResponse` (array of `UserListItem` + `Pagination`)
- Response 403: `Forbidden`
- No 401 response defined (relies on `RequireSession` middleware which returns 401)

**`POST /admin/users/{userId}/roles`** (lines 574-632):
- Path param: `userId` (uuid)
- Body: `AssignRoleRequest` → `{ role: Role }` where Role is enum `[admin, kurator]`
- Response 201: `User` schema
- Response 403: `Forbidden`
- Response 404: `NotFound`
- Response 409: three Problem examples with distinct `type` URIs:
  - `https://kencleng.dev/errors/role-conflict` (admin/kurator conflict, representative conflict)
  - `https://kencleng.dev/errors/duplicate-role` (already has this role)

**`DELETE /admin/users/{userId}/roles`** (lines 633-678):
- Path param: `userId` (uuid)
- Query param: `role` (Role enum)
- Response 204: no body
- Response 403: `Forbidden`
- Response 404: `NotFound`
- Response 409: two Problem examples:
  - `https://kencleng.dev/errors/last-admin` (INV-account-13)
  - `https://kencleng.dev/errors/pending-curation-assignment` (INV-account-14)

**Schemas:**
- `UserListItem` (line 2627): `allOf` ref to `User` — same shape as `User`, just a named subtype for the list context
- `UserListResponse` (line 2644): `{ data: UserListItem[], pagination: Pagination }`
- `Pagination` (line 2633): `{ next_cursor: uuid|null, has_more: boolean }`
- `AssignRoleRequest` (line 2656): `{ role: Role }`
- `Role` (line 2437): enum `[admin, kurator]`
- `Problem` (common.yaml line 121): RFC 9457 — `{ type: uri, title: string, status: int, detail?: string }`

## Requirement

The spec (`08-role-assignment.md`) references these endpoints and schemas. The contract is the source of truth per AGENTS.md §1 item 3.

## Gap

**No gap at the contract level.** All three endpoints, all request/response schemas, all error types, and the hard-capped `AdminUsersLimitParam` are fully defined. The contract is complete and ready to implement against.

One minor observation: `GET /admin/users` does not define a 401 response explicitly, but `RequireSession` middleware handles that uniformly for all authenticated endpoints — this is consistent with other authenticated endpoints in the spec (e.g. `POST /account/security/set-password` also omits 401).

## Sniffing

1. **Risk**: The `User` schema returned by `POST /admin/users/{userId}/roles` (201) includes `email` (decrypted). The spec says the response is the updated `User` — this means the assign endpoint also decrypts PII on the response path. Same pattern as `GET /account/me`, but worth noting for the "every row requires a decrypt" concern in the threat model.
2. **Edge cases**: `DELETE` uses query param `role` (not body) — the handler must parse and validate the enum value from query string. Invalid enum values need a 422 or 400 response (not explicitly defined in the contract — the spec only lists 403/404/409).
3. **Miscontext**: None found — contract matches spec.
4. **Misleading signals**: None — the `UserListItem` allOf inheritance looks like it adds fields, but it doesn't — it's just a type alias. The actual list item shape is identical to `User`.
5. **Inconsistency**: The `POST` 409 uses two different `type` URIs (`role-conflict` and `duplicate-role`) for three different scenarios. The spec (line 37-39) says the two exclusivity cases get "distinct messages" — the contract achieves this via different `detail` strings under the same `type` URI, which is fine. But `duplicate_assignment` uses a different `type` URI entirely. This is consistent with the spec's intent (Assumption A says duplicate is a separate case).
