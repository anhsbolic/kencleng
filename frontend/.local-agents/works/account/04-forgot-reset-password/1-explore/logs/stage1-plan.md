# Stage 1 — Plan Announcement

Feature: `docs/spec/1-account/features/04-forgot-reset-password.md`
(frontend surface: `POST /auth/forgot-password`, `POST /auth/reset-password`)

## Docs read (per task instructions)

- `docs/spec/README.md` — doc-type conventions, not feature-specific.
- `docs/spec/1-account/features/04-forgot-reset-password.md` — backend
  acceptance criteria for both endpoints (Tier 1, INV-account-05/08).
- `docs/ui-ux/page-map.md` — Guest table, rows: `/forgot-password` (Form
  — "Submit email to request password reset") and
  `/reset-password?token=...` (Form — "Submit new password"). Both
  Guest-only rows; no other row in any persona table references either
  page or a non-page surface (no email-link-landing entry distinct from
  `/reset-password` itself — unlike email verification, which page-map
  lists separately as "Email verification link (from email, not a full
  page)"). So the only frontend surfaces for this feature are these two
  full pages; the "non-page surface" is the reset email itself
  (backend/notification concern, not a frontend route).
- `docs/ui-ux/patterns.md` — Pattern 3 (Form Page), including the
  Cross-Pattern State Conventions (§B) for submit/error/success
  handling. `/forgot-password` and `/reset-password` are listed under
  Pattern 3's "Used by" only via `page-map.md`, not named explicitly in
  `patterns.md`'s own list — they're plain Form pages, not the
  "Revisable Submission" sub-pattern (no draft/curation cycle).
- `docs/ui-ux/prototype-reference.md` — both routes are **Tier 2** (no
  dedicated prototype). Closest Tier 1 precedent named explicitly:
  `/dashboard/campaign/new` (dashboard forms) **or** `/login` (auth
  modal/mobile split) — `/login` is the more relevant one since these
  are Guest/unauthenticated auth-adjacent forms, same Auth Shell.
  Checked `docs/design-reference/` directory listing directly (not just
  the table): 10 files, none named `forgot-password*` or
  `reset-password*`. `login-register.html` is the only plausible
  candidate for visual precedent — it covers `/login` and `/register`,
  not this feature, so it's precedent-only, not a literal target.
  Noted known issues (login field-level-vs-banner error conflation,
  campaign card upload-dropzone placeholder, typography drift) — the
  first is directly relevant since this feature also needs
  request-level vs field-level error separation.
- `docs/ui-ux/design-guidelines.md` — full token read: colors,
  typography, shape/elevation, button/input/badge component tokens,
  accessibility rules. No forgot/reset-specific tokens; standard Form
  Page + Input + Button tokens apply.

## Shell identification

Both pages sit under the **Auth Shell**
(`app/(auth)/layout.tsx` + `app/(auth)/_components/auth-shell-client.tsx`),
not Public Shell or Dashboard Shell — confirmed by directory location:
`app/(auth)/forgot-password/page.tsx` and
`app/(auth)/reset-password/page.tsx` already exist as routes today.

Current Auth Shell contents (from directory listing + line counts, not
yet read in depth):
- `app/(auth)/layout.tsx` (15 lines), `app/(auth)/_components/
  auth-shell-client.tsx` (66 lines) — the shell chrome shared by
  `/login`, `/register`, `/forgot-password`, `/reset-password`.
- `app/(auth)/login/page.tsx` (32 lines), `app/(auth)/register/
  page.tsx` (13 lines) — sibling, already-implemented Form pages in
  the same shell; the natural structural precedent.
- `app/(auth)/forgot-password/page.tsx` and `app/(auth)/reset-password/
  page.tsx` — **both already exist, 12 lines each, and are explicit
  placeholder stubs**: each file's own header comment reads "Placeholder
  — real form is Account Task #4's scope (docs/spec/1-account/
  tasks.md), not this playbook (phase0-shared-infra.md). Exists so the
  Auth Shell has a route to render against during this phase's
  verification." This confirms the routes were scaffolded during Phase
  0 shell setup specifically to be filled in by this task — not a
  page-consolidation conflict, not prior real work to preserve.

## Areas to explore, and order

1. **App routes + Auth Shell**
   (`app/(auth)/layout.tsx`, `app/(auth)/_components/auth-shell-client.tsx`,
   `app/(auth)/forgot-password/page.tsx`, `app/(auth)/reset-password/
   page.tsx`, plus sibling `app/(auth)/login/page.tsx` and
   `app/(auth)/register/page.tsx` as precedent). First, because
   everything else (what data the form needs, what hook to call) is
   downstream of confirming exactly what the shell provides today and
   how the two stub pages are meant to slot in.

2. **API client layer** (`lib/api/account.ts`, `lib/api/schema.d.ts`,
   `../api/openapi/account.yaml`, `../api/openapi.yaml`). Second,
   because the form components/hooks in area 4 can't be scoped until
   it's confirmed whether forgot/reset-password request/response types
   and fetch functions already exist in the generated schema and
   `account.ts`, or need to be added — a grep already shows zero
   `forgot`/`reset` hits in `lib/api/account.ts`, so this needs a full
   look at what pattern `login`/`register` follow there to mirror.

3. **Hooks layer** (`lib/hooks/use-login.ts`, `use-register.ts`, and
   confirming `use-forgot-password.ts`/`use-reset-password.ts` don't
   exist yet — directory listing shows they don't). Third, depends on
   area 2's API client shape.

4. **Form components** (`components/features/account/login-form.tsx`
   + `login-schema.ts`, `register-form.tsx` + `register-schema.ts` as
   precedent for the `react-hook-form` + `zod` + Form-pattern-state
   convention this task must follow). Fourth, depends on areas 2-3
   since the form's submit handler wires into the hook/API layer.

5. **Test/mock layer** (`mocks/handlers.ts` MSW conventions for
   `/auth/login`, `/auth/register` — grep already shows zero
   `forgot`/`reset` handlers today — plus existing `*.test.tsx`/
   `*.test.ts` conventions for login/register as precedent). Fifth,
   since mock/test shape mirrors whatever areas 2-4 establish.

6. **Visual/prototype precedent** (`login-register.html` in
   `design-reference/`, per `design-reference-usage.md`'s extraction
   method, plus re-confirming the `/login` error-conflation known issue
   doesn't get carried over). Last — informs Stage 3 visual decisions
   once the structural/data gaps from areas 1-5 are already mapped;
   doesn't change what's built, only how it looks.

Order rationale: 1 → 5 is a dependency chain (shell → data contract →
hooks → components → mocks/tests), each step needs the previous step's
findings to know what to look for. Area 6 is independent of that chain
and pushed last since it's precedent for visual polish, not for
behavior/data shape.

---

Awaiting go-ahead before Stage 2.
