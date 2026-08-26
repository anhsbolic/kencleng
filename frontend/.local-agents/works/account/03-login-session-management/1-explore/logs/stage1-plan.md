# Stage 1 — Plan Announcement
## Feature: account/03-login-session-management (frontend surface)

## Docs read so far

- `docs/spec/README.md` — doc-type conventions (spec vs project docs);
  `docs/spec/*` is the executable-spec, wins over `docs/project/*` on
  conflict.
- `docs/spec/1-account/features/03-login-session-management.md` — four
  endpoints (`POST /auth/login`, `POST /auth/login/mfa`,
  `POST /auth/refresh`, `POST /auth/logout`), Tier 1 with a Tier 0
  fenced sub-area (JWT issuance/verification, refresh rotate-on-use +
  reuse detection — backend-side only, no frontend Tier 0 concern).
  Key items for the frontend track:
  - `mfa_pending_token`: stateless HS256 JWT, 5-minute TTL, carried by
    the client between `POST /auth/login` (password step) and
    `POST /auth/login/mfa` (code step) — no DB persistence, so the
    frontend is its only "storage."
  - `401`/`429` share **identical generic detail text** for wrong-
    credentials vs. lockout — only the status code differs
    (anti-enumeration).
  - **Assumption D**, verbatim: multi-tab refresh-token races are
    explicitly deferred to "when the `account` domain's frontend track
    starts," to be solved with `BroadcastChannel` cross-tab
    coordination — noted in the backend spec specifically so this
    frontend task wouldn't have to rediscover it from scratch.
- `docs/ui-ux/page-map.md` — Guest table: `/login` = "Login form +
  'Masuk dengan Google' button." No dedicated row anywhere for an MFA
  challenge step, and no Cross-Cutting UI Elements row for logout.
- `docs/ui-ux/patterns.md` — `/login` is Pattern 3 (Form Page), Auth
  sub-variant; standard Idle → Validating → Submitting → Submit error
  (banner vs. field-level, never conflated) → Success state table
  applies. No MFA-step sub-pattern named anywhere in this doc.
- `docs/ui-ux/prototype-reference.md` — `/login` is **Tier 1**
  (`design-reference/login-register.html`), desktop modal / mobile
  full page. Known Issue #1: the prototype's login failure renders as
  a field-level error on the Email input instead of a banner —
  "status: not confirmed fixed in the final export, verify before
  implementing."
- `docs/design-reference/` listing — confirmed `login-register.html`
  exists (same single file `/login`+`/register` share, already
  established by task 02's exploration). No MFA-step file or state
  anywhere in the export (confirmed later in Stage 2/5 by extracting
  and reading the actual JSX, not just the filename list).
- `docs/ui-ux/design-guidelines.md` — color/type/shape/component
  tokens; nothing session/auth-specific beyond what task 02 already
  applied (Auth Shell modal styling, button/input tokens).

## Shell & current repo state

Per `frontend/AGENTS.md` and `page-map.md`, `/login` sits under the
**Auth Shell** — one of the three Phase-0 shells (Public / Auth /
Dashboard). The Auth Shell exists and is already partially populated
by tasks 01/02:

```
app/(auth)/layout.tsx
app/(auth)/_components/auth-shell-client.tsx (+ .test.tsx)
app/(auth)/login/page.tsx        ← real content (task 02), explicitly
                                    scoped: Google button + callback-
                                    error handling only; own comment
                                    names THIS task as the owner of
                                    "the form above/alongside the
                                    Google button"
app/(auth)/register/page.tsx     ← real (task 01)
app/(auth)/forgot-password/page.tsx   ← Phase 0 placeholder, task 04's scope
app/(auth)/reset-password/page.tsx    ← Phase 0 placeholder, task 04's scope
```

Also already scaffolded ahead of this task, by task 02:

- `lib/stores/auth-store.ts` — in-memory `accessToken` Zustand store,
  own comment: "shape only, no login logic — that's Account Task #3's
  job."
- `lib/api/client.ts` — `apiFetch`'s 401→`tryRefreshOnce()`→retry-once
  path is already fully wired to `POST /auth/refresh`, exported so
  `AuthBootstrapProvider` can reuse it on app mount.

`docs/spec/1-account/tasks.md`'s status tracker says task #3 is "in
progress... build not started," but the actual `backend/.local-agents/
works/account/03-login-session-management/` working tree already has
populated `3-build/`, `4-code-review/`, `4-patch/`, and `5-testing/`
directories — the tracker text looks stale against real working-tree
state (confirmed more concretely in Stage 2).

This means Stage 2's page-consolidation check is central: `/login`
already exists as a route with real (if partial) content — this task
is an *extension*, not new-page creation — but I haven't yet opened
`register-form.tsx` (the composition precedent `/login`'s own comment
points to), `auth-store.ts`, or `client.ts` in full to confirm exactly
where the boundary sits.

## Areas I intend to explore in Stage 2, in order

1. **Auth Shell + `/login` page current state** (`app/(auth)/`,
   `components/features/account/`) — first, since it's the primary
   page-map.md surface and directly establishes what task 01/02 left
   for this task, plus `register-form.tsx`'s composition pattern
   (named by `/login`'s own comment as the precedent to match).
2. **Session/token infrastructure** (`lib/stores/auth-store.ts`,
   `lib/api/client.ts`, `components/providers/auth-bootstrap-
   provider.tsx`) — login/MFA/refresh/logout all write into or read
   from this scaffolding; also where Assumption D's cross-tab
   coordination lands. Second because area 1's form can't be designed
   without knowing what plumbing already exists.
3. **API layer & generated types** (`lib/api/account.ts`,
   `lib/api/schema.d.ts`, `lib/hooks/`) — which of the four endpoints
   already have typed fetch functions/hooks vs. need adding.
4. **Dashboard Shell logout entry + session-expiry handling**
   (`app/(dashboard)/_components/dashboard-shell-client.tsx`,
   `nav-items.ts`) — downstream consumer of area 2's session state.
5. **Visual precedent + MFA boundary check**
   (`design-reference/login-register.html`, `/dashboard/security`) —
   confirm the login-time MFA challenge is cleanly separated from MFA
   *enrollment* (a different feature/page). Done last since it's a
   validation pass, best done once the functional shape from 1–4 is
   clear.

Order rationale: page/shell first (establishes the extension baseline
and the composition precedent to match), then session/token infra
(area 1's form logic depends on knowing what already exists there),
then the API/type layer (concrete wiring these depend on), then the
Dashboard Shell (a downstream consumer of area 2's state), then the
visual/boundary check last as a validation pass over the functional
shape established by 1–4.

Confirmed by Anhar — proceeded to Stage 2 as planned, no changes to
area list or order.
