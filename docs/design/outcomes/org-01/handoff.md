# ORG-01 — Organization Onboarding / Entry: Design Candidate 01

> **Standing: DESIGN CANDIDATE / NOT HUMAN APPROVED.** Feature-specific, experimental hi-fidelity static design, not canonical Kencleng Design Authority, implementation authorization, engineering acceptance, or verification of runtime.
>
> **Owner / next judge:** Anhar (Human Product/Design acceptance). **Next consumer only after exact-revision acceptance:** NextJS Builder, using approved Design Authority and accepted ORG-01 Rancang handoff in addition to this bounded candidate.
>
> **Repository target:** `anhsbolic/kencleng`, `pilot/station5-c1-space-only`, `docs/design/outcomes/org-01/`. Initial branch HEAD inspected as `d5f5ce3ed8f005e75524f09796f4dec0a18c497f` on 2026-11-02. During this session the branch advanced independently by two commits (`docs/design/README.md` plus a new optional static-HTML kickoff template); updated guidance was read before persistence. The candidate is based on current branch state, not a reset to the old checkpoint. Candidate source files: `prototype.html`, `handoff.md`. **Exact initial prototype SHA-256: `6cb9406ecb3915edeb85b481f55eab9e9a2efe03ddac77d4168c3dcb0ef18a72`.** No Product, Rancang, project-wide design, frontend production, or `api/` mutation.

## 1. Purpose / visual thesis

**Direction: An Editorial Beginning.** Translate Sunlit Editorial into the first act of creating organizational context, without treating onboarding as donation, application for external verification, or campaign setup. The screen asks users to **understand → attribute → review → commit**. It gives organization and internal initial Owner equal conceptual weight, then makes the one primary starting action understandable.

Composition: warm paper field, asymmetric editorial typography, a bounded explanatory two-part organization/Owner diagram instead of fake organizational imagery, a warmer instructional rail, sparse Sun action treatment, limited Berry detail, border/spacing-first hierarchy. Plain `kencleng` text in header is *noncanonical*, not a proposed final logo/wordmark. No photography, fabricated verification cue, or real-world organization evidence.

**Challenge resolved for the candidate:** A generic SaaS onboarding hero could wrongly imply “organization approved” or “Google sign-in creates the organization.” The chosen hierarchy explains the internal effect and boundaries before the CTA. Trade-off: more information above the primary action; validate reading cost and visual rhythm in human review.

## 2. Authority and evidence read

### Binding Kencleng Product

- `docs/product/outcomes/org-01/product-contract.md` — approved Product Contract v0.3; `docs/product/outcomes/org-01/approval.md` — Product approval/standing. The only success is a truthful, attributable pair: new organization context **and** first **internal** Owner. Initial Owner is not verified legal representation, Reviewer approval, public listing, campaign eligibility, or donation readiness. Similar names do not automatically block or merge separate contexts.
- `docs/product/outcomes/org-01/privacy-decision.md` and `privacy-notice-draft.md` — privacy/product boundaries. Identity of controller `Yayasan Kencleng Indonesia` and contact `care@kencleng.dev` settled. Notice text still **DRAFT**, not approved/publishable. Operational readiness and providers/locations remain open/held.

### Binding Kencleng Design

- `docs/design/README.md` (including newly added optional `templates/hi-fi-static-html-kickoff.md`, non-authoritative), `product-design-principles.md`, `brand-product-ui-brief.md`, `design-guidelines.md`, `patterns.md`, `asset-governance.md`.
- `docs/design/visual-references/selected-direction/README.md`: both approved reference filenames registered as `sunlit-editorial-public.png` and `sunlit-editorial-evidence-journal.png`. **Limitation:** connector could retrieve directory metadata/README, but could not return PNG binary pixels. Their direct visual content was NOT inspected; do not claim pixel/reference fidelity. Concrete palette/typography/radius guidance was read from canonical `design-guidelines.md` instead.

### Accepted Engineering solution: constraint, not Design Authority

- `harscode-state/org-01/rancang-002/handoff.md` and `experience.md`. Chosen progression: public explanation and privacy summary/full notice access **before** explicit Google login, then confirmed current account, name entry, explicit pre-effect review, explicit submit. Google callback alone creates nothing. Failure/unknown/recovery semantics do not permit fake success or blind retries.
- Relevant O3: no name draft carried through full-page sign-in, user input preserved on ordinary submission failure, uncertain creation outcome requires authoritative inspection, and review identity/effect before creation.
- Existing `frontend/features/org-01/workspace.tsx`, `provider.tsx`, and `harscode-state/org-01/nextjs-builder-001/handoff.md` consulted as **existing implementation evidence**, not design authority. Current frontend reports MOCK_DEVELOPMENT-only and browser verification unverified.

### Supporting Harscode research, not Kencleng authority

`anhsbolic/harscode-workspace` branch `refactor/harscode-foundation`: `design/README.md`, `design/exploration.md`, `design/authority-and-handoff.md`. Used for bounded exploration/authority separation/handoff rigor, not for new ORG-01 semantics.

## 3. Bounded UI scope

`prototype.html` is one stand-alone file with embedded CSS and minimal JavaScript **only** for switching illustrative views and explaining the authentication boundary. Open locally in a modern browser. Use its bottom-right **Preview state** controls or `#entry`, `#disclosure`, `#unavailable` hashes.

| State | Actor sees and understands | Action / non-action |
| --- | --- | --- |
| `entry` — public onboarding | Editorial promise, organization + first internal Owner pairing, three consequences, controller/contact and pre-login data summary. | Primary **Lanjutkan dengan Google** is visible only after context and data summary in reading order. In prototype it opens an explicit *simulation-only* explanation; it does NOT navigate to Google. |
| `disclosure` — before Google | More focused explanation of account identifiers, actor-provided name, internal role, controller/contact, and explicit **DRAFT** / incomplete notice warning. | Explicit route to repo's complete **draft source** for review. It is not advertised as final/published notice. Google CTA still simulation-only. |
| `unavailable` — illustrative pre-sign-in entry failure | Calm state, specific stage/context, explanation that **no organization creation was submitted** in this example. | Back to entry or disclosure. No fake network retry or assertion about a *post-creation* uncertain outcome. |

**Not within this page:** live account state, provider callback, name form, creation review/submit, creation unknown result, recovery, organization workspace, account deletion, Reviewer workflow, production notice, or legal consent design. These need their own controlled future design/review, using accepted Rancang semantics. Do not extrapolate this example unavailable state to unknown mutation results.

## 4. Material design decisions to review

1. **Information first, action second.** On large screens the hero conveys meaning; the right instructional rail keeps privacy and CTA in a single controlled reading path. The CTA is not a floating conversion widget.
2. **One conceptual pair.** The two outlined editorial blocks communicate organization context plus first internal Owner as one success condition. Numbers are conceptual identifiers **not progress %, evidence IDs, verification scores, status labels, or real organization data**.
3. **No fabricated trust icons.** Typography, provenance wording and consequence information carry confidence. Sun means action/emphasis, not success or legitimacy. No verified badge, shield, photos or fake person.
4. **Open full notice before Google.** Separate disclosure view maintains pre-login explanation; no checkbox claiming all-purpose consent. The link is to a *draft source* for candidate review, not a release-ready legal notice.
5. **Truthful unavailable state.** The example is explicitly scoped before Google/sign-in and before organization submission. More sensitive result-unknown states belong to the later creation/recovery design.
6. **Static-only interactions.** Preview state switcher is persistently labeled candidate tooling, not intended as a production navigation component. 'Lanjutkan' never performs a side effect.

## 5. Responsive, type and accessibility intent

- **Desktop:** two-column editorial main with a split between contextual headline and action rail. Entry hierarchy: headline → pair illustration → responsibility summary, with instructional rail containing consequence steps → pre-login summary → action.
- **Narrow mobile (~390px):** editorial intro first, semantic pair compressed without disappearing, then full instructional rail. The privacy summary and consequence appear before the CTA in document order. No sticky CTA covering the necessary context. Supplemental bottom statements follow after task region.
- CSS adaptation at 1100/790/520px is **prototype geometry, not mandated engineering breakpoints**. Do not force NextJS Builder to copy DOM or CSS.
- Canonical colors from `design-guidelines.md`; requested Newsreader/Instrument Sans font *roles* declared, but local-file fallback used (Georgia/Arial) because fonts are not bundled. Font-render equivalence not established.
- Semantic headings, semantic ordered steps, keyboard-operable state controls, focus-visible styles, focus moved to changed view heading, state preview not signaled only by color, reduced-motion support. Contrast, focus behavior and screen-reader announcements **not tested in an actual browser**.

## 6. Evidence, limitations and acceptance checklist

Evidence gathered:

- GitHub verified initial branch HEAD equal to `d5f5ce3...`, no existing candidate files/directory found prior to writing.
- Text authority and relevant implementation evidence read; Harscode supporting research read.
- Static HTML parsed successfully; duplicate IDs and unresolved `aria-labelledby`/`aria-describedby` checks returned zero; three headings and state controls present.
- `prototype.html` is self-contained and has **no external requests** in its CSS/JS. The linked draft resource is optional external navigation for review.

**Unverified:** Actual rendered desktop/mobile pixels, typography equivalence, keyboard journey, screen-reader behavior, contrast measured at rendered text sizes, potential clipped/overflow regions. Chromium headless screenshot attempts timed out in this environment; no screenshots or visual acceptance proof exist. Visual ref PNG pixels also unavailable. These are real acceptance gaps, not reasons to claim the prototype is final.

Human review should decide whether:

- [ ] Desktop editorial balance/hierarchy feels distinctively Sunlit Editorial without distracting from the primary task.
- [ ] Narrow mobile maintains comprehension before CTA; text density and rail rhythm are comfortable.
- [ ] Organization+Owner concept illustration feels clear, rather than mistakenly suggesting two steps/records/verified states.
- [ ] Pre-login transparency/copy is enough but not excessive, with strong DRAFT separation.
- [ ] The explicitly pre-login unavailable/error illustration is appropriately calm and recoverable.
- [ ] The exact HTML revision is accepted or receives bounded visual changes. **Only explicit Human acceptance may change standing.**

## 7. Open questions / boundary

| Classification | Question / follow-up owner |
| --- | --- |
| OPEN EXPERIENCE / DESIGN DECISION | Whether the pair graphic is the best expression of organization + initial Owner, and whether its dominance stays appropriate at narrow widths. Anhar after rendered visual inspection. |
| OPEN EXPERIENCE / DESIGN DECISION | Whether the amount of public pre-login explanation should be adjusted, without losing consequential and privacy information. Anhar after review. |
| OPEN PRODUCT / PRIVACY READINESS (already held, **not reopened here**) | Final notice text, provider roles/locations, legal/operational readiness and verified contact handling before production. Owning Product/privacy authority; this candidate cannot settle it. |
| OBSERVATION / EVIDENCE GAP | Selected-direction PNG actual pixels and actual rendered viewport verification not obtained. Browser review required. |
| OPEN ENGINEERING CONSTRAINT (later) | Actual font bundling, browser behavior and NextJS integration; Builder owns implementation mechanics after design acceptance. |

## 8. Next consumer and promotion rule

**Now:** Anhar visual review of exact `prototype.html` and this handoff. After an explicitly accepted file revision, record acceptance separately with revision IDs and scoped decisions. **Then:** NextJS Builder may interpret approved design with Product/Rancang/Design Authority; no direct copy-paste of mock behavior as API/authentication truth. Future pages (creation/review, unknown/recovery, workspace, auth failure, privacy/account deletion) are separate iterations and are **not** accepted from this entry candidate.