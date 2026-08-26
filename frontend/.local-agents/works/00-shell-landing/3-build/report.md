# Implementation Report — 00-shell-landing

> Ticket    : 00-shell-landing
> Feature   : Landing page (`/`) — Public Shell nav + highlighted campaigns mock
> Date      : 2026-08-26
> Spec ref  : `docs/ui-ux/page-map.md` (Public Shell + highlighted-campaigns decisions, 2026-08-24), `docs/ui-ux/prototype-reference.md` (`/` Tier 1 entry)
> Techplan  : `.local-agents/works/00-shell-landing/2-plan/techplan.md`
> Explore   : `.local-agents/works/00-shell-landing/1-explore/logs/` (4 areas + solutioning)
> Tasks     : none — techplan decomposition was evaluated and explicitly declined (see `2-plan/` history); executed directly from the techplan's Architecture/Plan (§9)

---

## 1. Summary

`/` moves from the untouched `create-next-app` scaffold (behind a nav-less
`(public)/layout.tsx` stub left deliberately unbuilt at Phase 0) to a real
Guest landing page: a Public Shell nav (desktop top nav, mobile hamburger +
drawer reusing the existing `useFocusTrap` hook) and a highlighted-campaigns
section backed by a real (mocked) `GET /campaigns` data layer. Includes two
new `components/ui/` primitives (`Badge`, `ProgressBar`), the codebase's
first typography-scale token addition to `globals.css`, a new one-time Shell
playbook (`scaffold-public-shell.md`), and the Hero/HighlightedCampaigns/
HowItWorks composition translated from the Tier 1 Claude Design prototype
per this repo's AI-prototype-to-production conventions.

The defining correctness/trust properties, all implemented and tested:

- **No fabricated data**: the highlighted-campaign card shows no
  organization name or verified badge — `GET /campaigns`'s real response
  shape (`CampaignListItem`) structurally can't supply either (only
  `organization_id`, a UUID). Applied the same reasoning to drop the
  prototype's Hero badge ("120 organisasi terverifikasi") and the whole
  `TrustStrip` section — none of these numbers have a real backing source.
- **No client-side recomputation of money/progress figures**: `CampaignCard`
  never receives raw `collected_amount`/`target_amount` at all, only the
  already-computed `progress.percentage` — enforced structurally by the
  prop interface, not just by convention.
- **No color-alone signaling**: the ProgressBar's 100%-fill color switch
  (`primary-600` → `success-500`, both green) is paired with a "Target
  tercapai" text label, not relied on alone.
- **Two known issues from the Tier 1 prototype did not carry over**: the
  campaign card's photo area is a plain non-interactive placeholder (not
  the prototype's upload-dropzone `image-slot` element), and all typography
  uses `design-guidelines.md`'s real type scale (30px/36px), not the
  prototype's drifted hardcoded values (44px/40px/48px).

Executed by one agent session end to end (explore → techplan → decomposition
check → implementation), with a human decision point on the one genuinely
product-facing open item (the "Jelajahi Kampanye" nav-link destination).
No commits made; everything is working-tree changes.

---

## 2. Files changed

### New files (16)

| File | Lines | Description |
|---|---|---|
| `app/(public)/_components/public-shell-client.tsx` | 97 | Mobile hamburger + drawer, mirrors `DashboardShellClient`'s shape, reuses `useFocusTrap` |
| `app/(public)/_components/public-shell-client.test.tsx` | 42 | R2–R5: `aria-expanded` toggle, focus-trap in/out, nav-item click closes drawer |
| `app/(public)/_components/nav-items.ts` | 16 | Shared nav item list (desktop + mobile read the same definition) — not in the original techplan file list, split out during implementation |
| `components/ui/badge.tsx` | 48 | New primitive — 6 semantic tones, `radius-full`, `caption`/`text-[11px]` sizing per `design-guidelines.md` |
| `components/ui/badge.test.tsx` | 17 | Renders label; tone maps to the right bg/text classes |
| `components/ui/progress-bar.tsx` | 41 | New primitive — `primary-600`→`success-500` fill switch at 100%, `role="progressbar"` with `aria-valuenow/min/max` |
| `components/ui/progress-bar.test.tsx` | 23 | R10/R11: fill color below/at 100%, value clamping |
| `lib/api/campaign.ts` | 31 | `getCampaigns()` — thin `apiFetch` wrapper typed off `schema.d.ts`, generic error on non-OK |
| `lib/hooks/use-campaigns.ts` | 26 | `useCampaigns()` + `campaignKeys` query-key factory, `staleTime: 60_000` |
| `components/features/campaign/campaign-card.tsx` | 79 | `CampaignCard` + `CampaignCardSkeleton` — no `organization` prop by design |
| `components/features/campaign/campaign-card.test.tsx` | 57 | R7 (no org/badge), R11 (goal-reached label), R12 (null days omitted), R14 (no upload affordance) |
| `components/features/landing/hero.tsx` | 46 | Static hero — real `text-h1`/`text-display` tokens, no fabricated org count |
| `components/features/landing/hero.test.tsx` | 22 | R13 (type-scale tokens present), no fabricated count in the badge |
| `components/features/landing/how-it-works.tsx` | 63 | Static 3-step "Cara kerja" section |
| `components/features/landing/how-it-works.test.tsx` | 13 | All three steps render |
| `components/features/landing/highlighted-campaigns.tsx` | 92 | The one `'use client'` leaf on `/` — loading/empty/error/success + stale indicator |
| `components/features/landing/highlighted-campaigns.test.tsx` | 78 | R6 (skeleton), R7 (success, no org/badge), R8 (empty, no CTA), R9 (error + retry, no raw error text), retry behavior |
| `frontend/.agents/docs/scaffold-public-shell.md` | 106 | One-time Shell playbook (techplan Decision 5), mirrors `phase0-shared-infra.md` Step 5's shape |

### Modified files (5)

| File | Diff | Description |
|---|---|---|
| `app/globals.css` | +45 | Full typography-scale tokens (`display`/`h1`–`h4`/`body-lg`/`body`/`body-sm`/`caption`) as CSS vars + `@theme inline` `--text-*` entries, wired identically to the existing color/radius/shadow pattern |
| `app/(public)/layout.tsx` | +72/-4 | Real Public Shell — desktop nav server-rendered, delegates mobile to `PublicShellClient` |
| `app/(public)/page.tsx` | +20/-65 | Replaced the `create-next-app` scaffold with `<Hero /><HighlightedCampaigns /><HowItWorks /></>` |
| `mocks/handlers.ts` | +87 | New `GET /campaigns` handler, 3-item fixture typed against `CampaignListResponse` — no `organization` field |
| `frontend/.agents/docs/README.md` | +1 | Indexed the new `scaffold-public-shell.md` playbook |

### Pre-existing changes (NOT this feature — flagged)

| File | Note |
|---|---|
| `api/openapi.yaml`, `api/openapi/account.yaml` | Modified outside this session — not touched here |
| `docs/ui-ux/page-map.md` | Modified outside this session — not touched here |
| `backend/.local-agents/works/account/03-login-session-management/{1-explore/logs/,2-plan/prompt.md,2-plan/techplan.md}` | A separate, unrelated backend task with its own live work in progress in the same working tree — confirmed via `git status` before this session's decomposition step (see that step's write-up), not touched here |
| `api/node_modules/`, `api/package-lock.json` | Untracked, unrelated to `frontend/` |

`git status` was checked before finishing so the diff for this feature can be isolated cleanly from the above at commit time.

---

## 3. Routes & components delivered

| Route/Component | Type | Behavior |
|---|---|---|
| `/` | Page (Server Component) | Real landing page: Hero → highlighted campaigns → "Cara kerja". Statically prerendered (confirmed in `next build` output). |
| `(public)/layout.tsx` | Layout (Server Component) | Public Shell: desktop top nav; delegates mobile hamburger/drawer to a Client Component leaf |
| `HighlightedCampaigns` | Client Component | The only `'use client'` boundary on `/` — calls `useCampaigns()`, handles all four List/Browse states |
| `GET /campaigns` | Consumed (unchanged, mocked) | `security: []`, no request params sent — `/`'s section is a fixed fixture, not a filterable browse UI |

No backend endpoints were created or changed — this is a pure frontend consumption task.

---

## 4. Rule coverage (R1–R17)

| Rule | Test(s) | Status |
|---|---|---|
| R1 (desktop nav, no hamburger at ≥`md`) | — | ⚠️ Not automated: jsdom doesn't evaluate CSS breakpoints (same pre-existing limitation as `dashboard-shell-client.test.tsx`) |
| R2 (hamburger + `aria-expanded=false` at <`md`) | — | ⚠️ Same jsdom limitation as R1; the hamburger's *behavior* once rendered is fully covered by R3–R5 |
| R3 (drawer open → focus to first item, `aria-expanded=true`) | `public-shell-client.test.tsx` | ✅ |
| R4 (Escape/nav-click → close, focus returns to hamburger) | `public-shell-client.test.tsx` | ✅ |
| R5 (Tab/Shift+Tab cycles within drawer only) | `public-shell-client.test.tsx` | ✅ |
| R6 (loading → skeleton, not spinner) | `highlighted-campaigns.test.tsx` | ✅ |
| R7 (success → N cards, no org name/badge) | `highlighted-campaigns.test.tsx`, `campaign-card.test.tsx` | ✅ |
| R8 (empty → icon+message, no CTA) | `highlighted-campaigns.test.tsx` | ✅ |
| R9 (error → generic banner+retry, no raw text) | `highlighted-campaigns.test.tsx` | ✅ |
| R10 (`percentage < 100` → `primary-600`) | `progress-bar.test.tsx` | ✅ |
| R11 (`percentage >= 100` → `success-500` + text label) | `progress-bar.test.tsx`, `campaign-card.test.tsx` | ✅ |
| R12 (`days_remaining: null` → segment omitted) | `campaign-card.test.tsx` | ✅ |
| R13 (headings use real type-scale tokens) | `hero.test.tsx` | ✅ |
| R14 (photo placeholder, no upload affordance) | `campaign-card.test.tsx` | ✅ |
| R15 (`percentage` displayed as-is, no recomputation) | Structural, not a runtime test — see §1 | ✅ (by construction) |
| R16 (stale/revalidating indicator) | — | ⚠️ Implemented (`staleTime`, `isFetching && !isLoading`) but not exercised by a test — reliably triggering a background revalidation needs fake-timer setup not written here |
| R17 ("Jelajahi Kampanye" → `#kampanye`, not `/campaign`) | `public-shell-client.test.tsx`, `hero.test.tsx` | ✅ |

14 of 17 rules have automated coverage; R1/R2 are a known, pre-existing tooling limitation (documented, not silently skipped) and R16 is implemented but untested. R15 is enforced by the type system rather than a runtime assertion, which is a stronger guarantee than a test would give.

---

## 5. Verification results

| Gate | Command | Result |
|---|---|---|
| Lint | `npm run lint` | ✅ clean |
| Tests | `npm run test` | ✅ 36/36 passed (14 test files) |
| Type-check + build | `npm run build` | ✅ compiles, type-checks, `/` prerenders as static content |
| `npm run verify` (lint+test) | re-run after every batch of changes | ✅ clean at each checkpoint |
| Static output sanity check | grepped `.next/server/app/index.html` | ✅ expected copy present (nav labels, hero heading, section titles); no "browse files"/hardcoded "verified" strings baked into the prerendered shell |

One real type error surfaced and was fixed during `build` (not `lint`/`test`,
since those don't run the TS compiler over `mocks/`): the mock fixture set
`unpublish_reason: null` / `closed_reason: null`, but `schema.d.ts` generates
these fields as optional-only (`Type | undefined`), not nullable, even
though `campaign.yaml` marks them `nullable: true` — a real `openapi-
typescript` codegen drift, not a fixture bug. Fixed by omitting the two
optional fields from the fixture rather than hand-editing the generated
file (`AGENTS.md`: "don't hand-edit the generated file").

---

## 6. Process deviations (flagged for audit trail)

All are implementation-level refinements within the techplan's stated
intent; none touch scope or the resolved decisions. Listed for the techplan
owner to ratify or reject.

### 6.1 Hero's fabricated-stat badge dropped, not just `TrustStrip`

The techplan's Decision 3 named `TrustStrip` specifically, but the Tier 1
prototype's Hero *also* carries a hardcoded stat ("120 organisasi
terverifikasi") with the identical no-backing-source problem. Applied
Decision 3's reasoning consistently: kept the badge's visual trust-signal
role, dropped the number ("Organisasi terverifikasi"). Not a new decision —
an application of an existing one to a spot the techplan's file-level
detail hadn't explicitly enumerated.

### 6.2 `Button`-in-`Link` nesting avoided (accessibility fix not in the techplan)

The techplan's Interface Contract didn't specify how Masuk/Daftar CTAs
should compose navigation with the `Button` component. `Button` always
renders a literal `<button>`; wrapping it in `next/link`'s `<Link>` would
nest interactive content inside interactive content (invalid, and
confusing keyboard/screen-reader semantics per
`accessibility-fundamentals.md`). Implemented as styled `<Link>` elements
matching `Button`'s outline/primary visual variants directly, in both
`layout.tsx` (desktop) and `public-shell-client.tsx` (mobile drawer) —
`components/ui/button.tsx` itself stays untouched, per the techplan's
Files-NOT-Changed list.

### 6.3 `nav-items.ts` split out as its own file

Not in the techplan's file list. `layout.tsx` (desktop) and
`public-shell-client.tsx` (mobile) both need the identical nav item array;
splitting it into one shared file avoids the two copies silently drifting,
the same reasoning `data-fetching-conventions.md` gives for a single query-
key factory.

### 6.4 `useCampaigns()` given an explicit `staleTime`

The techplan's Interface Contract specified `useQuery({ queryKey, queryFn })`
with no `staleTime`. Left at the TanStack Query default (`0`), R16's
"revalidating" indicator would fire on effectively every mount (data is
"stale" immediately), which isn't what `patterns.md`'s stale-data
convention intends. Set `staleTime: 60_000` and used
`isFetching && !isLoading` as the actual revalidating signal so the
indicator only appears during a genuine background refetch.

### 6.5 Mock fixture omits two `null`-valued optional fields

See §5 — `unpublish_reason`/`closed_reason` are typed as optional-only by
the current codegen output despite the OpenAPI spec marking them nullable.
Fixture omits them rather than passing `null`, to avoid a false type error
that has nothing to do with this feature's own logic.

---

## 7. Risk note

- **Assumptions made:**
  - The three fixture campaigns' `collected_amount`/`target_amount` pairs
    were chosen to be arithmetically consistent with their
    `progress.percentage` values (e.g. 10,200,000/15,000,000 = 68%) purely
    for fixture plausibility — this is mock data only; the real backend
    computes and returns `percentage` independently, and the frontend never
    derives it from the amounts (R15).
  - "Mulai berdonasi" (Hero primary CTA) and "Jelajahi Kampanye" (nav) both
    point at the in-page `#kampanye` anchor, matching the user's explicit
    decision on the nav link; not separately re-confirmed for the Hero CTA
    since it's the same underlying "no `/campaign` route yet" situation.
    "Galang dana untuk organisasi Anda" (Hero secondary CTA) links to
    `/register` — a judgment call (starting a fundraiser requires becoming
    a registered org representative), not explicitly specified anywhere.
- **Edge cases intentionally NOT handled (and why):**
  - Pagination controls for the highlighted section — page-map.md's mock
    scope is a fixed fixture set, not a filterable/paginated browse UI
    (that's `/campaign`'s job, a separate task).
  - A real campaign photo — no media field exists on any `Campaign`/
    `CampaignListItem`/`CampaignDetail` schema yet; the placeholder is
    permanent until that field exists, not a temporary stand-in with logic
    already half-wired.
  - Footer — explicitly deferred by `page-map.md`, not built.
- **What is not tested, and why:**
  - R1/R2 (desktop-vs-mobile breakpoint switching) — jsdom limitation, see
    §4. Needs a manual check: `npm run dev` (or `build && start`) and resize
    the viewport, or a Playwright-based check if/when E2E tooling is added
    (`kencleng-frontend-tech-stack.md` defers E2E until a demonstrated
    need).
  - R16 (stale indicator) — implemented, not exercised by a test; would
    need fake timers to reliably force a background revalidation window.
  - Visual/pixel-level parity with the Tier 1 prototype — no screenshot
    diffing set up for this task; verified structurally (tokens, component
    composition, states) and via a static-HTML grep sanity check (§5), not
    a rendered-pixel comparison.

---

## 8. Open items status (techplan §14)

| # | Item | Status | Note |
|---|---|---|---|
| 1 | Typography-scale token naming/wiring convention | ✅ Resolved | Implemented following the existing color/radius/shadow token pattern exactly; open to renaming in code review if a reviewer disagrees |
| 2 | Organization name/verified badge unavailable on `GET /campaigns` | ✅ Resolved (for this task) | Shipped without it, escalated to `campaign` domain's own techplan — unchanged by implementation |
| 3 | `TrustStrip` omitted for v1 | ✅ Resolved | Not built; Hero's equivalent fabricated-stat badge also dropped (§6.1) |
| 4 | "Jelajahi Kampanye" nav link destination | ✅ Resolved | User decision: anchor to `#kampanye` — implemented as R17, tested |
| 5 | Scope of "Organization trust signal" pattern requirement | ✅ Resolved | Moot for v1, per #2 — unchanged by implementation |

All 5 items were already moved to Resolved in the techplan before implementation began (per the Open Items Lifecycle rules); none reopened during the build.

---

## 9. How to run

```bash
cd frontend
npm install                 # if not already done
npm run dev                 # / renders the real Public Shell + highlighted campaigns (MSW-mocked)

# Verification
npm run lint
npm run test
npm run build && npm run start   # production build + serve

# Manual checks still needed (§4, §7):
#   - resize the browser window across the `md` breakpoint and confirm
#     the desktop nav / mobile hamburger switch correctly (R1/R2)
#   - open the mobile drawer with keyboard only and confirm focus
#     behavior visually matches the automated test's assertions
```

---

## 10. Next steps (workflow hand-off)

1. **Human review** of §6 deviations (five small items, none touching
   resolved scope) and §7 assumptions (Hero CTA link targets).
2. Manual browser check of R1/R2 (breakpoint switching) — no automated
   coverage exists for this, see §4/§7.
3. Optional: add a fake-timer test for R16's stale indicator if that
   behavior needs stronger regression protection before `/campaign` reuses
   `useCampaigns()`.
4. When `campaign` domain's own techplan starts: revisit Open Item #2
   (org name/badge on list cards) and swap the `#kampanye` anchor links
   (nav + Hero CTA) to a real `/campaign` href — both are one-line changes
   per `scaffold-public-shell.md`'s own notes.
5. Code review against `../../harscode-workspace/workflow/4-code-review/checklist.md`.
