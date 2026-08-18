### Structural principles
 
- **Feature components co-located in `components/features/`** first (not
  nested inside `app/`), so they're reusable across pages without
  duplication. Move to route-nested components only if something
  genuinely isn't reusable.
- **`lib/api/` is separate from `lib/hooks/`** — `api/` holds pure fetch
  functions (usable from both Server Components and inside hooks),
  `hooks/` wraps them with TanStack Query for use in Client Components.
  **[UPDATED — Step 2]** Request/response types in `lib/api/` are now
  generated from `api/openapi.yaml` via `openapi-typescript` (types
  only, not a full client SDK) — the fetch functions themselves stay
  hand-written, just typed against the generated interfaces instead of
  hand-written ones. See `kencleng-backend-tech-stack.md` "API Contract
  & Codegen" for the full decision.
- **Zustand stores are split per domain**, not one giant store — easier
  to read and maintain.
- **`components/shared/` is for cross-domain non-primitive components
  [NEW]** — different from `components/ui/` (which is pure design-system
  primitives with no business awareness) and different from
  `components/features/` (which is domain-specific). `shared/` sits in
  between: components like `MaskedField` know about the *concept* of
  PII across multiple domains (donation, organization, user) but aren't
  domain-specific themselves.
## Shared Component Notes **[NEW]**
 
### `MaskedField`
Central component for all PII masking, per `kencleng-actors-entities.md`
PII Handling Note. Used wherever `guest_email`, `User.primary_email`,
`NPWP`, or future banking details are displayed — **regardless of
viewer role, including Admin**.
 
- Default: masked display (e.g. `j***@***.com`)
- Explicit reveal toggle/button per field instance (not a global
  "show all PII" switch — each field revealed individually,
  intentional friction)
- On reveal by Admin or Kurator viewing another party's data: fires a
  call to log the reveal action to Audit Log (`kencleng-phase0-detail.md`
  Fitur 9) — this means `MaskedField` needs to know the *context*
  (whose data, what field, who's viewing) to pass along, not just be a
  dumb visual toggle
- Open question: does reveal auto-re-mask after a timeout / on
  navigation, or stay revealed until manual re-toggle or page refresh?
  Not yet decided.
### `SecureUploadNote`
Small reassurance note/popup used on every non-public file upload form
(organization legal docs, fund-usage-report attachments) — communicates
that the file is stored securely and privately. Purely informational,
no logic — just consistent copy/placement across the 2-3 forms that
need it.
 
### `CurationDecisionPanel`
Reused across the 3 curation contexts (organization curation, campaign
curation, fund-usage-report verification) — approve/reject buttons +
mandatory `decision_note` textarea on reject. Same interaction pattern,
different underlying entity — good candidate for one component
parameterized by curation type.
 
## Layout Patterns [NEW — RESOLVED 2026-07-20]
 
Decided during Step 1.5 wireframing (see `kencleng-ux-page-map.md`
"Dashboard Shell" / "Benchmark Design Reference" for the reasoning and
benchmark sources).
 
### Dashboard shell
Horizontal top-nav on desktop (no sidebar), top-bar + hamburger on
mobile. Applies to every authenticated route. Worth revisiting only if
the nav gets genuinely crowded once Organization Owner/Staff, Kurator,
and Admin personas are implemented — a concrete, demonstrated need,
not a reason to add a sidebar preemptively.
 
### Auth pages (desktop)
`/login`, `/register`, `/forgot-password` render as a **modal overlay**
on top of whatever page the user was on — not a dedicated route change.
`/reset-password` is the deliberate exception: it's a full page on
desktop too, since the user always arrives via an emailed link (no
"current page" exists to overlay a modal on top of).
 
### Auth pages (mobile)
Full-page routes for all four (`/login`, `/register`,
`/forgot-password`, `/reset-password`) — no modal pattern on mobile,
consistent with limited screen real estate.
 
### Google OAuth flow
**Full-page redirect** on both desktop and mobile, for every trigger
point (login, register, and account-linking from `/dashboard/security`)
— **not** a popup window. This was evaluated against a popup +
`postMessage` approach (which would have preserved the desktop auth
modal's context) and rejected for added complexity (popup-blocker edge
cases, extra handshake code) relative to the benefit, consistent with
the project's "lowest complexity" principle. Consequence: clicking
"Masuk/Daftar dengan Google" from the desktop modal navigates the
browser away from that modal entirely; the user lands on
`/auth/google/callback` and is then redirected to their final
destination.
 
### Public campaign pages
`/campaign/[id]` and `/campaign/[id]/donate` follow the GoFundMe/
Kitabisa benchmark structure: hero image → progress bar prominent →
sticky Donate CTA → donor list separate from narrative body (mobile);
two-column layout with sticky donate sidebar (desktop). Donation form
field order (nominal → payment method → optional donor info) already
matched the existing field order in `kencleng-phase2-detail.md` Fitur
1, so no rework was needed there.
 
## Testing [RESOLVED — Step 6]
 
- **Test runner: Vitest** — native ESM, faster than Jest, minimal
  config for a TypeScript + Next.js App Router project. Consistent
  with the rest of the modern tooling already in use (Vite-based
  ecosystem, Tailwind, TanStack Query). Jest was considered but
  rejected: more established, but ESM + App Router config friction
  outweighs any familiarity benefit for a fresh project.
- **Component testing library: React Testing Library** — de-facto
  standard for React, philosophy of testing behavior (query by
  role/text) rather than implementation detail. Effectively the only
  reasonable choice, not treated as a real trade-off decision.
- **Scope for v1: unit + component tests only, no E2E yet.** Covers
  hook/util logic and individual component behavior (form validation,
  conditional rendering, etc.). Playwright (or similar) for E2E is
  explicitly deferred — consistent with "lowest complexity, add only
  when there's a demonstrated need": E2E setup (browser automation,
  test data seeding) isn't justified before there's an actual page/flow
  to test end-to-end. Revisit once the first vertical slice
  (Registrasi & Login) exists.
- **API mocking: MSW (Mock Service Worker)** — intercepts requests at
  the network layer, so components under test exercise the same code
  path as runtime (real `fetch` calls, real TanStack Query hooks)
  rather than mocking `fetch`/the query client directly. Chosen over
  manual mocking (`vi.mock` on fetch functions) because the API
  contract is already type-safe via the OpenAPI-generated types (Step
  2) — MSW keeps tests aligned with that same contract shape instead
  of hand-rolled mock responses that can silently drift.
## Open Items — Needs Further Discussion
 
- Exact CORS / cookie configuration for the refresh-token flow (depends
  on final deployment topology — same-origin vs separate origins)
- ~~API contract format (REST plain JSON vs OpenAPI spec-first)~~ →
  **resolved: OpenAPI 3.x spec-first**, types generated via
  `openapi-typescript` (types only, fetch functions stay hand-written)
  **[RESOLVED — Step 2]**
- Whether `middleware.ts`-based route protection is sufficient, or if
  dashboard guarding needs to happen at the layout/component level too
- `MaskedField` reveal persistence behavior (auto-re-mask timing) **[NEW]**
- ~~Guest donation status page: server-side state (token-based lookup) vs
  client-side state~~ → **resolved: token-in-URL, server-side lookup by
  token** — see `kencleng-phase2-detail.md` Fitur 1 and
  `kencleng-ux-page-map.md` **[RESOLVED — NEW]**
## Not Yet Discussed
 
- Deployment/hosting target for the frontend (same host as backend via
  Docker Compose vs separate hosting like Vercel)
- ~~Testing approach for the frontend~~ → see **Testing** section
  above **[RESOLVED — Step 6]**