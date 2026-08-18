# Feature Spec — Account Profile (read)

> File: `docs/spec/account/features/07-account-profile.md`
> Status: draft — all open items resolved, ready for human review
> Risk tier: 2
> Domain: account

## Endpoint

- `GET /account/me`

## Acceptance criteria

- Given a valid access token, When called, Then the response is the
  caller's own `User` resource — `id`, `name`, `email` (decrypted,
  since this is always the resource owner viewing their own data —
  the `MaskedField` masking concern only applies when *other* users'
  PII is displayed, e.g. Admin's user list, per
  `kencleng-actors-entities.md`), `email_verified`, `roles`,
  `auth_providers`, `mfa_enabled`, `created_at`. `200`.
- Given no access token, or an expired/invalid one, When called, Then
  `401` — standard auth middleware behavior, nothing endpoint-specific.
- There is no request body, no path/query parameters carrying any
  identifier — the resource is entirely determined by the
  authenticated session, so there's no ID to tamper with (no IDOR
  surface).

## Error cases

| Condition | Expected response |
|---|---|
| Missing/expired/invalid access token | `401` |

## Applicable invariants

None — pure read, no state mutation, no invariant from
`docs/spec/account/invariants.md` is exercised or at risk here.

## Threat breakdown

Derived from `docs/spec/account/threat-model.md` component 5:

| Threat | Mitigation at this endpoint's level | Test that proves it |
|---|---|---|
| IDOR — fetching another user's profile | No ID parameter exists; resource is keyed entirely by the authenticated session | `TestAccountMe_NoIDParameter_SessionScoped` |
| Unauthenticated access | Standard auth middleware `401` | `TestAccountMe_RequiresAuth` |

## Risk tier & rationale

**Tier 2** — standard CRUD read, DTO shape taken directly from
`openapi.yaml`'s `User` schema, no invariant or concurrency concern.
Automated gates (contract test against the schema, auth middleware
test) are sufficient; no property/invariant test or mandatory human
review beyond a spot-check, per the Tier 2 definition in
`kencleng-agentic-workflow.md` §4.

## Assumptions / open questions

None.

## Audit log entry?

No — a read of one's own profile isn't a sensitive action in Fitur 9's
scope (that scope is about *changes*: role assign/revoke, MFA
enable/disable, account linking, PII reveal **of another user's**
data by Admin/Kurator — not a user reading their own already-known
data).

## References

- `docs/project/kencleng-actors-entities.md` — PII masking scope note
- `docs/spec/account/threat-model.md` component 5
- `api/openapi.yaml` — `GET /account/me`, `User` schema