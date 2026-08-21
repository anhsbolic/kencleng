# Patch — Agentic Workflow: drop wireframe references, narrow Finalize
# state-handling scope (2026-08-20)

> Target file: `kencleng-agentic-workflow.md`
> How to apply: two small in-place edits, §14 and §15.

## 1. §14, step 1 — replace

Old:
```
1. Analyze (Explore + Plan) → same feature spec in
   `spec/<domain>/features/<fitur>.md` already written for the backend track
   covers this — frontend doesn't get a separate spec doc, it
   implements against the same acceptance criteria plus
   `kencleng-ux-page-map.md` / wireframes / design guidelines
```

New:
```
1. Analyze (Explore + Plan) → same feature spec in
   `spec/<domain>/features/<fitur>.md` already written for the backend track
   covers this — frontend doesn't get a separate spec doc, it
   implements against the same acceptance criteria plus
   `docs/ui-ux/page-map.md` (which page, which pattern),
   `docs/ui-ux/patterns.md` (what shape/states that pattern requires),
   and `docs/ui-ux/design-guidelines.md` (visual tokens). Wireframes
   are retired — `patterns.md` is now the structural reference.
```

## 2. §15, Finalize — replace the "Resolve empty/loading/error states"
and "Check conformance" bullets

Old:
```
- Resolve empty/loading/error states concretely per page (this was a
  deliberate, documented deferral to implementation time — see
  `kencleng-roadmap-next-steps.md` Open Item #13 — this is where that
  deferral gets closed out)
- Check conformance against wireframes and `kencleng-design-
  guidelines.md`
```

New:
```
- Apply `docs/ui-ux/patterns.md`'s state definitions (loading/empty/
  error/success) to this page's actual data shape — narrower than the
  original deferred task: the *convention* is now pre-defined per
  pattern, so this step is applying it correctly to the page's
  specific fields/copy, not inventing state handling from scratch.
  Still a real per-page check (e.g. does this Detail page's empty
  state need a role-specific CTA per `page-map.md`), just smaller
  scope than before.
- Check conformance against `docs/ui-ux/patterns.md` (structural/
  state) and `docs/ui-ux/design-guidelines.md` (visual)
```
