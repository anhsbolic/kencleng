# Template — Kickoff Hi-Fi UI/UX Design (Multi-Page Static HTML)

> **Standing:** Optional reusable **session prompt** for bounded Kencleng Product Design work. Not a new Design Standard, Product/Design approval, Space formation, or repository-write authorization.
> **Method:** Human–AI IA discussion → hi-fi **separate HTML page files** → rendered feedback → optional project-owned Design Candidate persistence after explicit Human authorization.
> **Authority:** Current Kencleng Product and Design sources, not this prompt. Harscode `design/` is supporting research guidance only.

## How to use

1. Open a **new ChatGPT session in Project Kencleng**. Copy the prompt between the markers below and fill `{{OUTCOME_ID}}`, `{{OUTCOME_ID_LOWER}}`, `{{TARGET_BRANCH}}`, `{{INITIAL_SURFACE}}`, `{{SCOPE_NOTES}}`.
2. **Attach the actual selected-direction PNG reference images with the kickoff message.** GitHub textual connector access alone may not expose image pixels. Add prior UI screenshots only if useful as *existing-implementation evidence*.
3. The **first reply** must propose bounded information architecture (IA) and one recommended visual approach, then **stop for Human feedback**. It must not create/persist a full prototype or commit/push.
4. After Human accepts the IA direction, make one or a few **separate static HTML pages** for the agreed scope, present rendered views/preview files, critique and iterate with Human. **Stop again before any repository write.**
5. A later **explicit Human authorization** controls writing candidate files, and separate explicit authorization controls committing/pushing. Accepted design standing also needs a distinct exact-revision Human decision. **Neither images nor this template grant write/commit/push permission.**

The two review stops apply to this **material feature-specific design experiment**, not every routine micro-decision or technical UI fix.

---

## BEGIN COPY-PASTE KICKOFF PROMPT

Bro, kita mulai **Product UI/UX Design — `{{OUTCOME_ID}}`** dalam Project Kencleng. Kita akan berdiskusi dan menguji desain secara visual, bukan menyerahkan satu hasil jadi tanpa Human review.

### Goal and scope

Hasil akhir yang ingin kita capai: **feature-specific hi-fidelity UX/UI design** yang enak dipakai, jelas secara Information Architecture (IA), konsisten dengan approved Kencleng visual direction, dan cukup matang untuk dijadikan input NextJS Builder.

- Repository: `anhsbolic/kencleng`.
- Target branch: `{{TARGET_BRANCH}}`; discover actual current HEAD and original authority, jangan bergantung pada stale commit.
- Outcome: `{{OUTCOME_ID}}`.
- Initial surface/page: `{{INITIAL_SURFACE}}`.
- Boundaries and goals: `{{SCOPE_NOTES}}`.

**Ini sesi Product Design, bukan Next.js implementation, Rancang, backend, auth integration, ataupun produksi.**

### Authoritative inputs

Temukan dan baca sources yang efektif serta relevan, secara proporsional:

1. **Product:** current approved scoped Product Contract, approval, actor goals, consequences, negative paths. Do not invent Product/domain/privacy decisions or make a held production release gate into a blanket blocker for synthetic design.
2. **Kencleng Design:** `docs/design/README.md`, canonical `product-design-principles.md`, `patterns.md`, `brand-product-ui-brief.md`, `design-guidelines.md`, `asset-governance.md`, and `visual-references/selected-direction/`. Also use the **working supplemental guidance** `information-architecture.md` and `copywriting-id.md` for disclosure placement and natural Indonesian UX writing; they do **not** override canonical project authority.
3. **Actual attached reference images:** inspect and use the PNG files **sent by Human with this kickoff** as approved-direction evidence, together with their repository README and canonical guidelines. If no image is attached or its actual pixels cannot be inspected, **tell Human and request the images before making the visual-direction decision**. Do not pretend to have viewed images merely from filenames/README. Attached images are **not** pixel-perfect screen specs, canonical logos, or permissions to invent factual imagery.
4. **Accepted Engineering solution, if available:** current Rancang handoff and relevant experience/interaction/interface/state obligations. Those constraints ensure design is truthful and implementation-consumable, but do not make Engineering the owner of visual design.
5. **Existing frontend and screenshots (if any):** observation/critique evidence, **not Design Authority**. Harscode research `design/README.md`, `exploration.md`, and `authority-and-handoff.md` may provide supporting guidance without overriding Kencleng.

Distinguish binding project decisions, approved visual direction, working Design Candidate, historical examples, and unverified implementation behavior.

### First decision: Information Architecture before visuals

**FIRST RESPONSE ONLY:** after source discovery, perform a focused IA audit for `{{INITIAL_SURFACE}}` and stop for my feedback. No final multi-page HTML, no repository writes, no commit/push in this first reply.

For the specific user's **current decision** and journey stage, explain concisely:
- What the person wants to accomplish **now**, what they must know **before this specific action**, and which implications belong at a **later consequential step**.
- Which information should be **always visible** (brief, comprehensible, action-relevant), **available on demand**, **shown at a later journey step**, or **omitted** from this screen. Preserve easy access to consequential information rather than dumping everything above the CTA.
- Recommended placement/treatment per material information class: inline summary, link to a dedicated detail page, expandable section, contextual modal/dialog, or tooltip/help text. **Choose by content length, importance, timing, revisitability, mobile/keyboard accessibility and interruption cost**; don't prescribe one control universally.
- Keep trust-critical information, critical uncertainty, required action, and necessary pre-effect consequences visible/accessible at the correct time. **Do not hide mandatory disclosure or essential decision information inside hover-only tooltips**.
- Propose **natural Indonesian candidate copy** for the main heading, primary CTA, and one disclosure/error message. Avoid literal contract/English translations, verbose caveats and legalistic UI language; preserve scoped truth.
- If lengthy explanation competes with the current action, recommend whether a **meaningful icon, labeled diagram, illustration, or sourced photo** would communicate its relationship/context better. Explain its role and whether an approved asset exists; do not generate filler or fabricate evidence.
- For sign-in, distinguish **login itself** from later organization creation/commit. E.g., a short pre-login processing summary plus easy access to the complete notice **before Google sign-in** can satisfy accepted ORG-01 requirements without displaying the full policy or creation consequences on the login screen. Do not assume this example is global policy for other features.

Give **one recommended IA/interaction direction** with the key trade-off and a compact page/surface map. Ask me to assess it. Then **stop**. Don't run ahead to prototype or persist.

### After I approve the IA direction: create the hi-fi design

Translate **Sunlit Editorial / Evidence-Led Optimism** into actual page composition: purposeful visual hierarchy, editorial typography, restrained colors, brand character, spacing, navigation, interaction clarity, responsive priorities, and appropriate accessibility. **Good IA and idiomatic Indonesian UX copy matter as much as good-looking CSS.** Before filling unused space with explanations, challenge information priority and consider a justified meaningful visual. Recommend/brief missing material illustration rather than silently inventing brand or fake evidence.

Produce **one HTML file for each actual page/surface in the agreed scope**, **not one HTML document toggling entire pages through JavaScript**. Keep ordinary state variations in the relevant page where appropriate; additional page files only if a distinct page is genuinely useful.

Possible *candidate* structure if later persisted (adapt to actual scope, do not create every example):

```text
docs/design/outcomes/{{OUTCOME_ID_LOWER}}/
  pages/
    onboarding.html
    privacy-information.html    # only if designed as a distinct page
    auth-failure.html           # only if in scope
  prototype.css                 # optional shared styles
  handoff.md                     # only at authorized candidate persistence
  assets/                        # only if necessary, with provenance
```

The prototype must work as local static pages with relative links/navigation; **lightweight JS only for disclosure or illustrative non-business UI interactions**. No Next.js implementation, real API, OAuth, cookies, backend, MSW, personal data, invented product claims or fake live operations. No mandatory design tooling such as Figma.

Use correct page-level content priorities on desktop and narrow mobile. Show meaningful states with truthful synthetic text. Prefer deliberate, minimal design complexity; avoid new visual assets when existing expression suffices.

**Render and inspect** the actual HTML if possible. Provide preview/screenshots or accessible temporary artifact files to Human for review. If rendering fails, clearly label `NOT VISUALLY VERIFIED`, supply locally openable files where possible, and invite Human screenshots. Do not substitute source inspection for visual approval.

**SECOND REVIEW STOP:** present the designed pages and specific critique/limitations to me, then **wait for my visual/IA feedback and revision requests**. Do not silently treat a generated design as approved, and do not write to GitHub merely to make the prototype inspectable. Iterate until Human is satisfied; no need to request approval for minor CSS tweaks.

### Persistence and exact-revision acceptance — only after separate Human instruction

The kickoff **does not authorize repository writes, creating commits, or pushing**. You may inspect sources and create temporary chat/downloadable prototype artifacts for review. A user saying "start", "lanjut" or "buat prototipe" is permission for the design step being discussed, **not automatic persistence permission**.

After visual review, **ask for the specific repository action** only if it is needed. Human may authorize:
- Writing Design Candidate files within a verified project-owned destination (proposed `docs/design/outcomes/{{OUTCOME_ID_LOWER}}/`; confirm existing convention and current HEAD).
- Committing, and separately pushing, that exact candidate scope. Never infer this from the template, GitHub access, or earlier examples.

Maintain `DESIGN CANDIDATE / NOT HUMAN APPROVED` until Human explicitly accepts the **exact design revision**. A commit/push only persists a candidate and is **not** design approval. No write to `frontend/`, `api/`, Product Contract, Rancang sources, or canonical Design Standards under this assignment. If approved later, record discoverable exact-revision Design handoff and Engineering freedom, without duplicating Product/Engineering source of truth.

### Collaboration

Use **Challenge first, then commit**. Challenge information overload, poor action hierarchy, premature disclosure, hidden consequences, misleading language and gratuitous visual decoration. Do not be contrarian without concrete UX consequence. Be a collaborative Designer who brings one strong recommendation and waits at meaningful Human review boundaries.

**Start now with the FIRST RESPONSE ONLY: current-source verification + recommended IA and page-level disclosure strategy for `{{INITIAL_SURFACE}}`. Do not generate/push the full HTML in the first answer.**

## END COPY-PASTE KICKOFF PROMPT

---

## Example filling — ORG-01 (illustrative, not implicit authorization)

| Placeholder | Example |
| --- | --- |
| `{{OUTCOME_ID}}` | `ORG-01` |
| `{{OUTCOME_ID_LOWER}}` | `org-01` |
| `{{TARGET_BRANCH}}` | `pilot/station5-c1-space-only` |
| `{{INITIAL_SURFACE}}` | `Organization Onboarding / Entry and Google sign-in` |
| `{{SCOPE_NOTES}}` | `Prioritize a fast, clear login decision: concise upfront data-processing summary and link to full DRAFT notice before Google login; defer organization/Owner creation consequences until their appropriate review step. Discuss IA first. No direct GitHub writes at kickoff.` |

Current ORG-01 discovery candidates (verify current standing rather than hardcode):
- `docs/product/outcomes/org-01/product-contract.md`, `approval.md`, `privacy-decision.md` and `privacy-notice-draft.md` (notice still draft until approved).
- `harscode-state/org-01/rancang-002/handoff.md`, `experience.md`, relevant `interface-decisions.md`.
- Two approved `docs/design/visual-references/selected-direction/*.png` images **attached directly to the chat by Human**.
- `docs/design/outcomes/org-01/` may already contain an **unapproved candidate**; treat it as prior exploration/evidence, not authoritative design, and do not overwrite without explicit permission.
- `docs/design/information-architecture.md` and `docs/design/copywriting-id.md` are working project-level aids for review; apply them without displacing canonical Product/Design authority.

**Design Guidelines already include "Ordered Transparency, Not Information Dumping" and the UX Pattern System covers Progressive Disclosure.** Evaluate application of these established principles before proposing a new canonical IA standard.
