# Stage 2 — Area 6: Cross-Domain Dependencies

## Current State

**Tables referenced by the spec:**
- `organization_representatives` — needed for INV-account-10 (admin ⊥ representative check on assign)
- `organization_curation_assignments` — needed for INV-account-14 (pending curation check on revoke kurator)

**Migration status:** Neither table exists. The `backend/migrations/` directory contains only `account`-domain migrations (000001-000010). No `organization`-domain migrations exist.

**ERD definitions** (`kencleng-erd.md`):
- `organization_representatives` (line 766): `{id, user_id, organization_id, level, created_at}` with `UNIQUE (user_id, organization_id)` and indexes on `user_id`, `organization_id`
- `organization_curation_assignments` (line 803): `{id, organization_id, kurator_id, assigned_by, assigned_at, decision, decision_note, decided_at}` with indexes on `kurator_id`, `organization_id`, and a partial unique index `ux_org_curation_one_pending` on `(organization_id) WHERE decision = 'pending'`

**Cross-domain ownership** (from invariants.md):
- INV-account-10 (line 175-180): `account` enforces the assign direction; `organization` will enforce the invite direction later
- INV-account-14 (line 269-276): `account` reads `organization`'s table; `organization` owns the table and the `pending → approved/rejected` transition

## Requirement

From `08-role-assignment.md`:
- **INV-account-10** (line 35-36): assigning `admin` to a user who is an `organization_representatives` row of any organization → 409
- **INV-account-14** (line 67-73): revoking `kurator` from a user with a `pending` row in `organization_curation_assignments` → 409

From invariants.md:
- **INV-account-10** (line 159): `SELECT EXISTS(...)` against `organization_representatives WHERE user_id = ?`
- **INV-account-14** (line 249): `SELECT EXISTS(...)` against `organization_curation_assignments WHERE kurator_id = ? AND decision = 'pending'`

## Gap

**The tables don't exist yet.** The repository methods `ExistsOrganizationRepresentative` and `ExistsPendingCurationAssignment` (identified in Area 3) can't be implemented or tested without the target tables.

**Options (not solutions — flagging for human decision):**

1. **Create the tables as part of this feature's migration** — a new migration (e.g. `000011`) creates both tables. This is the simplest path: the tables are defined in the ERD, the indexes are specified, and the feature needs them. The `organization` domain will later add its own application logic on top of these tables, but the schema itself is stable (ERD is finalized).

2. **Defer INV-account-10/14 enforcement** — implement the feature without the cross-domain guards, add a TODO/assumption note, and ship the guards when `organization` domain is built. This violates the spec (both invariants are marked RESOLVED), but may be pragmatic if the tables shouldn't be created before `organization`'s own spec is written.

3. **Create the tables but skip the application-layer enforcement** — create the empty tables so the repository methods can run (returning `false` since no rows exist), but don't test the "user IS a representative" or "user HAS pending curation" paths until the tables have data. This is a half-measure — the code exists but isn't exercised.

## Sniffing

1. **Risk**: If the tables are created now but `organization` domain later changes the schema (e.g. adds columns, changes constraints), the `account` domain's queries could break. The ERD is finalized, but the ERD is a design document — migrations are the source of truth. The risk is low for `SELECT EXISTS` queries (they only check existence, not specific columns), but worth noting.

2. **Edge cases**: `ExistsOrganizationRepresentative` checks `user_id` — if the user is a representative of *any* organization (any level, any org), the check blocks admin assignment. The spec says "any `level`" (invariants.md line 161). The query is simple: `SELECT EXISTS(SELECT 1 FROM organization_representatives WHERE user_id = ?)`.

3. **Miscontext**: The spec says INV-account-14 is a "cross-domain read against a table `organization` owns" (spec line 72). This implies the table should exist before this feature ships. But the `organization` domain isn't built yet — creating its tables from `account`'s migration is a cross-domain schema ownership question.

4. **Misleading signals**: The ERD defines both tables in detail — looks like they "exist." But the ERD is a design document, not a migration. The actual DB doesn't have these tables.

5. **Inconsistency**: The invariants.md says INV-account-10's "Verification" section: "Today (account-only): test that assigning `admin` to a user with an existing `organization_representatives` row is rejected" (line 171-172). This implies the table should exist for testing. But the table doesn't exist yet. The "Deferred" note (line 173) is about the *reverse* direction (organization → account), not this direction.
