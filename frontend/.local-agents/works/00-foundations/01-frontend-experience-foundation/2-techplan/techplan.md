# Tech Plan: Frontend Experience Foundation

> Phase             : Techplan  
> Ticket            : none  
> Author            : Codex  
> Model             : GPT-5  
> Reasoning         : not exposed  
> Session           : not exposed  
> Created           : 2026-09-16  
> Target revision   : `a30ee75307ff14a7053dbdb749dae030fbcbb727`  
> Workflow revision : `4199c6db1b26ef1920ba670f222aff0c6d0f9e59`  
> Status            : Approved by @anhsbolic 
> Approach          : Establish a deliberately bounded, static public calibration surface at `/`; no Campaign data or speculative design system.  
> Refs              : `docs/spec/0-foundations/features/01-frontend-experience-foundation.md`; `1-exploration/logs/01-scaffold-and-root-route.md` through `06-phase-handoff.md`

---

## 1. Background

The clean reboot left `/` as a technical placeholder. This task creates the first rendered public precedent for Sunlit Editorial / Evidence-Led Optimism so later surfaces inherit coherent visual and interaction foundations without treating a full landing page or Campaign semantics as in scope.

## 2. Scope

**In scope:**

- Replace the root placeholder with a static, public, text-led calibration surface: semantic header, editorial hero, same-page transparency/clarity explanation, and a compact closing route-local action treatment.
- Establish only the global font roles and visual values consumed by this root surface, plus the matching title/description metadata.
- Deliver semantic landmarks, visible-label in-page navigation, responsive desktop/mobile composition, and focused behavior coverage for the meaningful anchor/landmark contract.
- Render and inspect the new root surface before human desktop/mobile acceptance.

**Out of scope (explicit):**

- Campaign discovery/listing/detail, organizer data, donations/payments, authentication, API calls, mocks, generated API types, client/server data state, forms, or global stores.
- Invented campaign, progress, amount, outcome, ranking, curation, verification, popularity, urgency, or beneficiary claims.
- A final wordmark/logo, production illustration family, campaign placeholder, campaign photography, motion system, or any Level 3/4 expressive asset.
- A `components/ui/` or `components/shared/` contract, registry entry, generic starter component library, or Playwright scenario absent a newly discovered repeatable browser regression.

## 3. Requirements

| ID | Requirement | Source / evidence |
|---|---|---|
| Q1 | `/` must be a coherent representative public slice, sufficient for desktop/mobile human brand and experience judgement, rather than a complete landing page. | Feature spec — Feature surface, Goal, Acceptance criteria; `1-exploration/logs/01-scaffold-and-root-route.md#Area-clean-scaffold-and-the--calibration-surface` |
| Q2 | The rendered precedent must express the approved visual system: Sunlit Editorial / Evidence-Led Optimism, Newsreader display plus Instrument Sans UI/body roles, warm-neutral foundation, restrained Sun/Berry, and border/spacing-first grouping. | `docs/ui-ux/design-guidelines.md` §§2–12, §20, §27; `1-exploration/logs/02-public-design-and-truth-authorities.md#Requirement` |
| Q3 | Public hierarchy must earn action through early transparency context, a calm clear CTA, and responsive preservation of task, trust information, and reachable action. | `product-design-principles.md` §§1, 8, 13; `brand-product-ui-brief.md` §§5, 8; `design-guidelines.md` §§11, 21 |
| Q4 | Content and visual treatment must not imply unsupported Campaign, trust, evidence, outcome, or ranking facts; unknown/open asset and product decisions remain explicit/deferred. | Feature spec — Requirements and Explicit non-goals; `product-design-principles.md` §§2, 5, 7; `asset-governance.md` §§1–5 |
| Q5 | Component/state ownership must remain route-local and Server Component-first; only actually needed global styling may be added. No broad primitive, data/client-state, or API contract is introduced. | `kencleng-frontend-tech-stack.md` §§4, 6–10; `components/README.md` §§1–13; `1-exploration/logs/03-architecture-and-component-governance.md#Requirement` |
| Q6 | Verification must include proportional static/build and focused behavior evidence, desktop/mobile real-browser inspection, and human rendered acceptance; Playwright is conditional, not routine. | Feature spec — Verification expectations; `kencleng-frontend-tech-stack.md` §§15–20; `1-exploration/logs/04-verification-assets-and-baseline-health.md#Requirement` |

## 4. Rules & Validation

- **R1** — Given a visitor opens `/`, when the page renders, then it exposes a semantic public-page structure with a textual product identifier, one editorial primary heading, a visible-label same-page action, and an actual target section explaining the platform’s clarity/transparency posture; every in-page link targets an element that exists in the document.
- **R2** — Given the root surface is rendered, when display/UI text, surfaces, grouping, and actions are inspected, then Newsreader and Instrument Sans have their approved role separation; the warm-neutral, border/spacing-first system is apparent; Sun/Berry are restrained accents; and status/trust is not conveyed by brand color, badges, or decorative symbols.
- **R3** — Given the static root content is reviewed, when copy, controls, and visuals are considered, then they remain platform-level and non-transactional: no fictional campaign/organizer/amount/progress/outcome/ranking/verification/urgency claim, donation action, or evidence-like imagery appears.
- **R4** — Given the implementation is organized, when source ownership and runtime boundaries are inspected, then root composition remains route-local and Server Component-compatible; no `components/ui`/`components/shared` contract, API/data layer, client store, form, or expressive production asset is introduced. Global CSS contains only values/font hooks consumed by this surface.
- **R5** — Given representative desktop and mobile widths with realistic text wrapping, when the page is inspected and its same-page action is used, then the primary heading, transparency explanation, and action remain visible/reachable; no horizontal overflow or clipped critical content occurs; and landmarks/link focus remain understandable without color alone.
- **R6** — Given the change is prepared for handoff, when verification is completed, then focused page behavior checks and static production checks pass, rendered desktop/mobile evidence is recorded, and a human has accepted the rendered result before the task is treated as accepted.

## 5. Decision Log

| ID | Decision / option | Status | Rationale / consequence |
|---|---|---|---|
| D1 | Static, text-led platform/trust calibration slice at `/` | Chosen | Establishes visual hierarchy, CTA restraint, truth-first language, and responsive precedent without requiring Campaign data. Functional public journeys remain deferred. |
| D1-alt | Complete landing page with discovery and donation CTA | Rejected | Would require Campaign semantics/data or fabricated campaign, progress, steward, and trust context; both violate explicit scope/truth boundaries. |
| D2 | Route-local Server Component composition plus only surface-used global styling | Chosen | One route does not demonstrate a broad semantic component contract or a client-state/data need. `components` registry remains unchanged. |
| D2-alt | Build generic shell/Button/Card/Badge/Progress primitives first | Rejected | Recreates retired abstraction behavior and establishes unsupported broad contracts before real consumers exist. |
| D3 | No campaign/expressive production asset; textual product identifier only | Chosen | Avoids fake documentary evidence and avoids silently canonicalizing OPEN logo/illustration decisions. Plain text identifies the product but is not a wordmark decision. |
| D3-alt | Use selected-direction reference, stock, or synthetic documentary hero imagery | Rejected | References are not production assets; such media can imply campaign evidence and bypass asset governance/human approval. |
| D4 | No Playwright scenario unless Build finds a stable, repeatable browser regression | Chosen | The selected route has one same-page navigation contract; focused browser inspection plus human acceptance are proportionate. Any later automation proposal must state value, omission risk, and owner. |
| D5 | Proceed without final logo/illustration/motion tokens | Chosen | These remain documented OPEN details but do not block a text-led, no-asset foundation. |

## 6. Backward Compatibility

- Existing data: none; the root route currently contains only a reboot placeholder.
- API/contracts/clients: no API, event, or generated-type contract changes; no existing product client behavior is preserved or migrated.
- Migration/deprecation compatibility: none. Retired frontend UI is explicitly not compatibility authority.

## 7. Edge Cases & Risks

| ID | Risk / edge case | Likelihood | Severity | Mitigation / accepted exposure |
|---|---|---:|---:|---|
| RISK-1 | A decorative treatment, badge, color, copy, or image implies verification, outcome, popularity, or real Campaign evidence. | Medium | High | Use platform-level copy only; add no campaign/expressive asset; inspect content against R3 and require human truthfulness review. |
| RISK-2 | Narrow layouts hide the transparency context, make the in-page action unreachable, or overflow with realistic text. | Medium | Medium | Inspect desktop and mobile at representative widths with text wrapping; preserve information priority before decorative composition. |
| RISK-3 | The first global visual rules become an unhealthy precedent or recreate retired tokens/components. | Medium | Medium | Limit globals to values/font hooks demonstrably consumed at `/`; keep composition route-local and assess it with the design guideline final test. |
| RISK-4 | The approved font roles degrade to arbitrary system fallbacks or cause an environment-specific production-build failure. | Low | Medium | Use a standard Next-compatible loading path for the approved fonts and confirm the selected mechanism in `npm run build`; stop if available project/build constraints prevent it rather than silently substituting unapproved roles. |
| RISK-5 | A green verification result gives false confidence because the baseline previously had zero tests and no visual evidence. | Medium | Medium | Add focused observable page coverage; run static checks; use real-browser inspection and human acceptance rather than treating lint/unit output as visual proof. |

## 8. Interface Contract

**Persistence/data shape:** None. The route consumes no remote, persisted, or mocked data.

**API/event/external interface:** No API, event, URL-query, form, or external-service contract is introduced. The only intentional navigation interface is an in-document fragment link to the transparency/clarity section.

**Cross-layer/business boundary:** The frontend renders platform-level explanation only. It must not calculate, source, label, or imply Campaign/business truth. If implementation requires a campaign-specific claim/value, stop and route the gap to its domain spec/API owner.

## 9. Architecture / Plan

1. Replace the bootstrap root composition with one static public document flow: header → editorial hero and same-page action → early transparency/clarity section → compact closing treatment.
2. Configure the approved font roles and the small global visual foundation consumed by that flow; update baseline metadata so it no longer describes a reboot placeholder.
3. Keep page-specific markup and styling ownership in `app/`. No runtime data, interaction state, or reusable component layer is needed for fragment navigation.
4. Add focused observable test coverage for document landmarks and the valid same-page action; then gather rendered desktop/mobile and human acceptance evidence.

## 10. Implementation Details

| Anchor | Why relevant | Intended change / precedent |
|---|---|---|
| `app/layout.tsx` — `metadata`, `RootLayout` | Current document boundary, `lang="id"`, and baseline metadata. | Keep document language; load/expose the approved Newsreader/Instrument Sans roles through a standard Next-compatible mechanism; update title/description to describe the public experience rather than the reboot. Do not introduce a global interactive shell. |
| `app/globals.css` — global document rules | Existing Tailwind v4/global-style boundary. | Define only root-surface-consumed visual values and font hooks from `design-guidelines.md`: warm canvas/surface/ink/border roles, restrained Sun/Berry, approved public spacing/radius/grouping posture, focus-visible treatment, and responsive foundation. Avoid a speculative legacy-style token/variant taxonomy. |
| `app/page.tsx` — `Home` | Current `/` placeholder and route-local composition owner. | Replace placeholder with the static flow in §9. Use semantic landmarks/headings and descriptive anchor text; target the early transparency section with a real fragment ID. Keep copy factual, calm, and platform-level. No client directive, fetch, campaign card, money/progress/status presentation, image, icon package, or non-existent route link. |
| `app/page.test.tsx` — new focused root-page behavior test | No current component behavior is protected; Vitest/RTL supports this scope. | Assert observable landmark/heading presence and that the visible same-page link references the existing transparency target. Keep it contract-focused rather than asserting layout classes or implementation structure. |
| `components/README.md` — registry | Registry is intentionally empty. | Do not change it: no broad `ui`/`shared` contract is established by this one route. |
| `package.json` — scripts | Defines repository verification commands. | Do not change. Use existing `verify` and `build`; do not add browser tooling or dependencies for this static scope. |

## 11. Files Changed / Files NOT Changed

| File / area | Change type | Description |
|---|---|---|
| `app/layout.tsx` | Modify | Approved font-role setup and production-appropriate root metadata. |
| `app/globals.css` | Modify | Minimal consumed global visual foundation and responsive/focus rules. |
| `app/page.tsx` | Modify | Route-local static public calibration composition. |
| `app/page.test.tsx` | Add | Focused observable semantic/fragment-navigation behavior coverage. |

| File / area intentionally untouched | Why |
|---|---|
| `components/ui/`, `components/shared/`, `components/README.md` | No demonstrated broad component contract; preserve the intentionally empty registry. |
| `lib/`, `mocks/`, API/OpenAPI sources, backend | No data/API/business behavior belongs to the static slice. |
| `public/`, asset directories, selected-direction reference files | No campaign or expressive production asset is authorized/needed. |
| `tests/browser/`, `playwright.config.ts`, `package.json` | No repeatable browser regression justifies automation or dependency/config changes. |

## 12. Testing Checklist

| Rule | Verification / evidence | Primary owner | Why this is worth running / risk if skipped |
|---|---|---|---|
| R1 | Run the new Vitest/RTL page test asserting semantic landmark/heading presence and visible fragment link → existing target relationship. | Build | Protects the only interactive contract from regressions to a non-functional navigation cue; class/style-only checks would not establish it. |
| R1 | Independently run `npm run verify` after the Build change. | Testing | Confirms lint plus the authored observable test in the repository’s normal baseline; skipping leaves a green-no-tests baseline risk. |
| R2 | Inspect the rendered route at representative desktop and mobile widths against the approved type, color, grouping, and action hierarchy rules. | Build | Static checks cannot prove font role use, hierarchy, or excess card/color treatment; omission risks a poor first visual precedent. |
| R2 | Perform human rendered acceptance at representative desktop and mobile scope. | Human | Brand coherence and whether optimism is evidence-led rather than decorative need product/design judgement that automation cannot decide. |
| R3 | Review implemented copy, controls, and asset imports against the prohibited-claim boundary before handoff. | Build | This is a truthfulness boundary, not a generic visual style preference; omission risks fabricated Campaign/trust semantics. |
| R3 | Include truthfulness/no-implied-evidence in human rendered acceptance. | Human | Human review validates that actual wording and visual emphasis do not imply unsupported facts. |
| R4 | Inspect changed-file scope and source boundaries; run `npm run build` to confirm the selected font path and Server Component-compatible route compile. | Build | Prevents accidental client/data/shared-component expansion and catches font/build integration failure; omission leaves an unverified global precedent. |
| R5 | Exercise the fragment link and keyboard focus in a real browser at desktop and mobile widths; inspect for wrapping, clipping, and horizontal overflow. | Build | jsdom cannot establish actual scrolling, focus visibility, or responsive layout; omission risks inaccessible or broken narrow presentation. |
| R5 | Manually exercise the same responsive/action path during human acceptance. | Human | Confirms interaction comprehension and priority preservation in the actual product context. |
| R6 | Record Build commands/results: `npm run verify`, `npm run build`, and focused desktop/mobile browser inspection. | Build | Supplies implementation confidence and a reproducible evidence trail; static/unit results alone are insufficient for material UI. |
| R6 | Independently repeat final `npm run verify` and `npm run build`; review Build’s rendered evidence. | Testing | Provides independent baseline/build confidence before delivery. The build is especially valuable because font integration crosses CSS/framework build processing. |
| R6 | Complete human desktop/mobile rendered acceptance before marking the task accepted. | Human | Required by the feature spec; without it, the first production brand/experience calibration lacks the mandated acceptance evidence. |

### Test Focus Pointer

| Area | Why sensitive | Evidence anchor from Exploration | Still relevant post-synthesis? |
|---|---|---|---|
| Campaign/auth/payment/PII/concurrency boundaries | No such runtime boundary survives the selected static slice: it consumes no campaign data, action, credential, or shared state. | `1-exploration/logs/02-public-design-and-truth-authorities.md#Sniffing`; `1-exploration/logs/03-architecture-and-component-governance.md#Sniffing` | N/A — D1 and D2 intentionally exclude these boundaries. Truthfulness remains covered by R3, not a specialized security test class. |

## 13. Open Items

### Active — needs external input or verification

None. Human approval of this Draft Techplan is the required phase gate, not an unresolved implementation-direction item.

### Resolved — retained as decision history

1. ~~**Minimum representative `/` slice was unspecified.**~~ **RESOLVED — static text-led platform/trust calibration surface chosen in D1.** Campaign journeys and data remain deferred.
2. ~~**Whether a broad component foundation should precede the route was unspecified.**~~ **RESOLVED — route-local Server Component composition only in D2.** No registry update is warranted.
3. ~~**Whether a public visual needs a campaign/expressive asset before implementation.**~~ **RESOLVED — no production asset in D3.** OPEN logo/illustration/motion decisions stay deferred rather than silently filled.
4. ~~**Whether browser automation is required.**~~ **RESOLVED — no default Playwright coverage in D4.** Real-browser inspection and human acceptance remain required; reconsider automation only for a concrete repeatable regression.
