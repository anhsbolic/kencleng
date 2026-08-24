/**
 * Global, account-wide role. The API's `Role` schema
 * (`lib/api/schema.d.ts`) only enumerates `'admin' | 'kurator'` —
 * elevated roles a user is explicitly granted. `'donatur'` isn't a
 * granted role at all; it's the implicit status every authenticated
 * user has, so it's added on the frontend side when deriving a
 * user's effective roles (see `lib/hooks/use-has-role.ts`) rather
 * than expected from the API response.
 */
export type GlobalRole = "donatur" | "kurator" | "admin";

/**
 * Organization-scoped role level (`Organization`/representative
 * context, e.g. the `my_level` field returned by organization
 * endpoints) — distinct from `GlobalRole`, which is account-wide.
 */
export type OrgRoleLevel = "owner" | "staff";
