# Stage 2 — Area 3: Dashboard Shell + `/dashboard/security` page +
page-consolidation check

## Current state

- **Dashboard Shell is fully built already** (`app/(dashboard)/
  layout.tsx` — thin Server Component wrapper, `_components/
  dashboard-shell-client.tsx` — 154 lines): top-nav desktop / top-bar +
  hamburger mobile per `patterns.md` Pattern 4, `FilteredNavLinks`
  renders `nav-items.ts`'s three items through `useHasRole`, mobile
  drawer has a working focus trap (`useFocusTrap`, `onEscape` closes),
  `LogoutButton` gated on `useAccountMe()`'s `data` (not role-based),
  `NotificationBadge` present. Nothing here needs to change for this
  task — the "Keamanan" nav item already exists and already points at
  `/dashboard/security`.
- **`app/(dashboard)/dashboard/security/page.tsx` is confirmed a
  12-line placeholder** (re-confirmed, full content): renders a static
  `h1` "Keamanan" + "Placeholder — Account Task #5." paragraph, no
  imports beyond React/JSX. Genuinely empty — no prior real work exists
  to preserve or consolidate with.
- **Sibling placeholder**: `/dashboard/profile` (Task #7, not yet
  built) is the same 12-line shape — confirms this is Phase 0's uniform
  scaffolding pattern across all three Dashboard Shell pages, not
  something specific to security.
- **`docs/spec/1-account/tasks.md`**: Task #5 (this task, "Account
  Linking") is in **Serial group S1** (tasks #1-#5, same core tables,
  strict dependency order — Linking explicitly "depends on Google OAuth
  and set-password both already existing," both of which are
  implemented per the git log). Task #6 ("MFA TOTP") is a **separate,
  independent Group B** ("independent tables `mfa_totp_secrets`,
  `mfa_backup_codes`, can run in parallel with S1 or Group C") — no
  dependency either direction between #5 and #6.
- **No MFA frontend code exists anywhere** — confirmed by directory
  listing across `components/features/account/`, `lib/hooks/`,
  `lib/api/`: no `mfa`-named file of any kind. Task #6 has not started.
- **`page-map.md`'s `/dashboard/security` row is one Form
  (multi-section) page covering both tasks**: "Enable/disable MFA (QR
  scan + confirm code), view/regenerate backup codes, link/unlink
  Google identity. Google-only users also see 'Atur Password' here."
  This is the **only** page-map row for either task's frontend surface
  — Task #5 has no row of its own, Task #6 has no row of its own; they
  share this single entry.
- **`mocks/handlers.ts`'s `mockUser` fixture defaults to
  `auth_providers: ["email_password"]`** (line 108) — an
  email_password-only user, i.e. already in the Branch-2/no-Google
  state by default. No Google-only fixture variant exists yet anywhere
  in the mock layer, and no mock handlers exist yet for either of this
  task's two endpoints (expected — "one handler per endpoint, added as
  the page/component that needs it gets built," per the file's own
  header comment).

## Requirement

- Per `page-map.md`, the real `/dashboard/security` page must
  eventually contain both the MFA section and the linking/set-password
  section as one multi-section Form page.
- Per `phase0-shared-infra.md`'s revision note, Task #5 was
  specifically the task Phase 0 anticipated would first turn this
  placeholder into a real route ("Task #5, Account Linking, maps to
  `/dashboard/security`, already in this phase's Dashboard Shell nav
  scope") — Task #6 (MFA) isn't mentioned in that note at all, i.e.
  Phase 0's own author was aware of Task #5's ownership of this route
  but the note is silent on how Task #6 is meant to slot in later.
- This task's own feature spec (`05-account-linking.md`) only specifies
  the two linking-related endpoints/behaviors — it does not mention the
  MFA section, does not claim ownership of the whole page, and has no
  "how this composes with Task #6" note of its own.

## Gap

- The page needs to go from a 12-line placeholder to a real
  multi-section form, but **only half the sections this page is
  ultimately supposed to have are in this task's scope** (linking +
  conditional "Atur Password" for Google-only users). The MFA
  enable/disable + backup-codes section is Task #6's scope, not built,
  not started, no dependency link forcing an order between the two.
  Neither task's own spec document says explicitly whether:
  (a) Task #5 builds the whole page shell (with an empty/deferred slot
  where Task #6's section will later go), or
  (b) Task #5 builds a page scoped to only its own two sections, and
  Task #6 later edits this same file to add its section, or
  (c) some other split (e.g. a shared page-level component composing
  independently-owned section components).
  This is exactly the "page-map action with no single owning task"
  shape the page-consolidation check is meant to surface — not a
  missing-endpoint gap, but an ownership-boundary gap between two
  already-scoped tasks sharing one page-map row.
- No mock fixture exists yet for a Google-only user (`auth_providers:
  ["google"]`) — needed to exercise Branch 1 (add password) and the
  unlink-blocked (`409`, only-identity) case in tests; the default
  fixture only covers the Branch-2 (already has email_password) shape
  out of the box.

## Page-consolidation check (workflow §14 step 1)

- **Does this task's page already exist from an earlier frontend task
  in this domain?** Yes, as an empty Phase-0 placeholder only — not
  prior real work from another *feature* task. No consolidation
  conflict with already-shipped functionality (unlike, hypothetically,
  editing a page Task #3 had already filled in).
- **Any page-map action with no backing endpoint in this domain's
  `tasks.md`, or vice versa?** No orphan at the endpoint level (checked
  in Area 1: both of Task #5's endpoints are in both `tasks.md` and
  `page-map.md`'s implied action list). The real finding is one level
  up: **one page-map row is backed by two separate, independently-
  scheduled tasks (#5 and #6) with no stated composition plan between
  them** — not an orphan, but a shared-ownership case the standard
  "one row ↔ one task" check doesn't cleanly cover. Flagging explicitly
  per the instruction to note this rather than silently assume either
  task's spec has already resolved it (neither has).

## Sniffing

- **Risk**: if Task #5 builds the full page shell including a
  hardcoded/placeholder MFA section "for completeness," that section
  becomes throwaway work Task #6 must delete or heavily rework — wasted
  effort and a needless diff for whichever task runs second. Converse
  risk: if Task #5 renders *nothing* where MFA will go, a user with MFA
  already enabled at that point in time (impossible today, since #6
  hasn't shipped, but relevant once it has and this page is revisited)
  would see an incomplete-looking page until #6 lands — acceptable for
  now since #6 genuinely hasn't shipped, but worth being deliberate
  about in Stage 3 rather than accidental.
- **Miscontext**: `page-map.md` describes the finished-state page as if
  it's one atomic deliverable ("Form (multi-section)... Enable/disable
  MFA... link/unlink Google identity...") without acknowledging it's
  actually assembled from two independently-tiered, independently-
  scheduled backend tasks — a reader of `page-map.md` alone would not
  discover this split without cross-referencing `tasks.md`'s Parallel/
  serial grouping section, which is a different document entirely.
- **Edge case**: since Task #6 (MFA) can run *before*, *after*, or *in
  parallel with* Task #5 per the Group B independence, a future session
  building Task #6 could just as easily hit this same question from the
  other direction (finding Task #5's real linking section already
  built, and needing to add an MFA section alongside it without
  disturbing it) — whatever composition approach Stage 3 picks here
  should be legible enough for that future, unrelated session to extend
  correctly without re-deriving this same analysis.
- **Inconsistency**: none found between `tasks.md` and `page-map.md`'s
  facts (both correctly list the same two endpoints for Task #5) — the
  gap is an omission (no stated composition plan), not a contradiction.

Proceeding to Area 4 (feature components precedent).
