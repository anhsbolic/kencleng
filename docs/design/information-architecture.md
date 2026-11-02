# Kencleng — Information Architecture & Disclosure Guidance

> **Standing:** Working, project-level design guidance for experimentation and review. Supplements, but does not replace or amend, the canonical [Product Design Principles](product-design-principles.md) §4 **Ordered Transparency, Not Information Dumping** and [UX Patterns](patterns.md) §18 **Progressive Disclosure**.
> **Concern:** What a person needs to see, understand, find, or revisit at each decision point; not Product/domain meaning, legal-policy approval, route/API design, or mandatory UI widgets.
> **Review:** Validate with representative rendered pages and Human judgment before promoting this working guidance to new canonical design precedent.

## Problem this guidance addresses

A truthful interface can still be hard to use when it asks people to read every relevant fact before doing a small, reversible action. **Transparency is access to the right information at the right decision—not an information dump.** Excessive pre-action prose competes with the primary task, hides the useful disclosure inside other text, and makes a design look complete without improving comprehension.

The opposite mistake is equally harmful: hiding trust-critical facts, relevant uncertainty, money, authorization, or irreversible consequences behind decoration or hard-to-discover disclosure.

## 1. Begin with the decision, not with a content inventory

For each page/state, write one sentence answering **“What is this person deciding or doing now?”** Identify what success, failure and recovery mean *for that step*, using approved Product and accepted Engineering Experience sources. Then distinguish:

- **Needed now:** the primary task/action, indispensable context, relevant trust/identity fact, and material consequences of *this action*. Make these directly perceivable in reading and visual hierarchy.
- **Needed before a later consequential action:** preserve it at that later review/confirmation step; do not front-load the full consequence onto an earlier navigation or sign-in step.
- **Useful on demand:** supporting explanation, evidence details, policy text, provenance/history, and answers to likely questions. Keep a visible, descriptively labeled way to reach it at the time it matters.
- **Not needed on this surface:** redundant assurances, repeated policy paragraphs, internal engineering terminology, or details unrelated to the current decision. Remove or relocate; do not bury everything in a drawer by default.

**Accessibility of information is not the same as presenting all of it simultaneously.** Legal or Product requirements may demand specific upfront disclosure or a full notice path: preserve those requirements exactly; this guidance does not reinterpret them.

## 2. Design a decision-stage content map

Use a small working map *only when a page is overloaded or the placement choice is consequential*. Do not maintain a giant feature-wide IA spreadsheet by default.

| Question | Practical judgment |
| --- | --- |
| Current person/task | First-time visitor, returning owner, contributor, reviewer? What is the immediate action? |
| Decision-critical fact | What would cause a material misunderstanding if not noticed before this action? |
| Required short explanation | Can one clear sentence state it without changing meaning? |
| More detail | Where can people find the complete explanation and come back without losing context? |
| Later commitment | What belongs at the next meaningful review/confirm step, not here? |
| Negative/uncertain state | How does the page distinguish “not started”, “not available”, “unknown”, “pending”, and “completed” where the approved contract requires it? |
| Mobile / assistive use | Is task context preserved in logical order, with labeled controls and keyboard access? |

Do not rely on a fixed word limit. Use the **first-read burden** test: can a new reader identify the page purpose, main action, and its immediate implication without parsing several equally prominent blocks?

## 3. Choose disclosure treatment by the information's role

| Treatment | Suitable use | Avoid when |
| --- | --- | --- |
| **Inline sentence or small callout** | Needed-now consequence, important uncertainty, mandatory warning or concise pre-action summary | It becomes an entire policy document or repeats every detail |
| **Descriptively labeled link → dedicated page** | Full notice, substantial reference content, material that must be revisitable/bookmarkable | It is the only place a critical one-sentence warning exists |
| **Accordion / expandable detail** | Optional explanation or bounded evidence that helps a reader stay in context | Essential content is hidden by default or opening it is hard on keyboard/mobile |
| **Dialog / sheet** | Short focused explanation or confirmation that merits deliberate interruption and a clear return path | Long legal reading, multiple unrelated topics, normal navigation, or a surprise commitment |
| **Tooltip / help text** | Very brief clarification of a local term or control, supplementary to a visible label | Critical disclosure, error message, consent, nontrivial instructions, or anything hover-only |
| **Dedicated review/confirmation step** | Exact target, identity/authority, irreversible consequences and explicit commitment just before consequential mutation | It adds ceremonial friction to ordinary reversible navigation |

All interactive disclosures require discoverable labels and meaningful keyboard, focus, dismissal and return behavior. **A small link that says “Baca informasi privasi” is more useful than a vague “Selengkapnya.”** Do not make dark-pattern disclosure: the route to important details must be genuine and available before the relevant action.

## 4. Example: ORG-01 is two different decisions

The example below interprets the current approved ORG-01 scope; confirm exact effective sources for any later implementation.

| Journey moment | Make immediately clear | Defer / provide on demand |
| --- | --- | --- |
| **Before Google login** | Why sign-in is needed; that sign-in alone does **not** form an organization; concise data-use/controller context as approved; **accessible full DRAFT notice before sign-in** | Long account/Owner creation explanation, complete legal-policy text, organization creation recovery mechanics; link to the draft for detail without claiming production readiness |
| **Before organization creation commit** | Proposed organization name, the currently acting account, first **internal** Owner role, what this does **not** verify, and an explicit “buat” action | Background explanation of earlier login and long policy text already available elsewhere |
| **After uncertain creation result** | Effect is **not confirmed failed or successful**; original request reference and allowed inspection/recovery; do not create fresh unintended effects | Unrelated onboarding storytelling; never replace uncertainty with a generic success or retry prompt |
| **Before account deletion commit** | Actual complete scope, account/context consequences, confirmation and any relevant copy/retention uncertainty specified by accepted contract | General introductory explanations that distract from the destructive decision |

**Illustrative visual design, not new ORG-01 semantics.** The UX task is to express accepted meaning with lower cognitive load, not to weaken disclosure requirements.

## 5. Balance text with evidence and visual support

When the page becomes text-heavy, first audit repetition and timing. Then ask whether a labeled process illustration, meaningful icon, status graphic or properly sourced image would communicate structure faster than an extra paragraph. Follow [Asset Governance](asset-governance.md) for truth, provenance and asset standing. Never use imagery in place of required precise wording or as fake evidence.

A visual can replace **redundant explanation**; it cannot replace **consequential facts**.

## 6. Review against actual screens

Check the rendered desktop and narrow-mobile experience where possible, including initial, focused, expanded, error, pending, and recovery states as applicable. Inspect:

- Can the user explain the immediate action and its consequence without long reading?
- Are Product-critical claims and uncertainty still discoverable **at the right time**?
- Is the first-screen hierarchy calm and recognizable, rather than several competing headings and equal-weight paragraphs?
- Do links, disclosures and navigation let people return without losing input/context?
- Does the mobile layout preserve reading order and primary-action accessibility?
- Has any visual shorthand implied unapproved legitimacy, authorization, verification, financial success or deletion completion?

When actual rendering is unavailable, label visual validation **NOT VERIFIED**. A source-only claim is not rendered UX evidence.

## Authority and handoff

Current Kencleng Product authority and canonical Design principles/patterns prevail over this working guidance. Record only material page-level IA decisions in a **feature-specific Design Candidate/Handoff**, with Human acceptance and exact revision when required. Do not silently change Product semantics, canonical guidance, or production code by treating an exploratory HTML prototype as binding.
