# Patch — Repo Setup: docs/ui-ux/ directory added (2026-08-20)

> Target file: `kencleng-repo-setup.md`
> How to apply: replace the `docs/` block inside the "## 2. Directory
> structure" tree with the version below.

## Replace the `docs/` block

Old:
```
├── docs/
│   ├── project/                  # existing docs (narrative, for humans to read)
│   │   ├── kencleng-erd.md
│   │   ├── kencleng-backend-tech-stack.md
│   │   ├── kencleng-frontend-tech-stack.md
│   │   ├── kencleng-actors-entities.md
│   │   ├── kencleng-business-process-overview.md
│   │   ├── kencleng-ux-page-map.md
│   │   ├── kencleng-phase0-detail.md ... phase3-detail.md
│   │   ├── kencleng-design-guidelines.md
│   │   └── kencleng-roadmap-next-steps.md
│   ├── spec/                     # executable spec, for the agent — see kencleng-agentic-workflow.md
│   │   ├── domains/               # invariants per domain (account, organisasi, campaign, donation, disbursement, notification)
│   │   ├── features/              # acceptance criteria + threat breakdown per endpoint/vertical-slice
│   │   ├── threat-model/          # STRIDE per domain
│   │   └── README.md              # structure & blank templates for each doc type above
│   ├── wireframes/                # gray-box HTML/SVG, mobile + desktop
│   └── kencleng-agentic-workflow.md  # process reference doc (lives at docs/ root, not project/ or spec/ — this is process, not product spec)
```

New:
```
├── docs/
│   ├── project/                  # existing docs (narrative, for humans to read)
│   │   ├── kencleng-erd.md
│   │   ├── kencleng-backend-tech-stack.md
│   │   ├── kencleng-frontend-tech-stack.md
│   │   ├── kencleng-actors-entities.md
│   │   ├── kencleng-business-process-overview.md
│   │   ├── kencleng-phase0-detail.md ... phase3-detail.md
│   │   └── kencleng-roadmap-next-steps.md
│   ├── ui-ux/                     # frontend UX doc set — NEW 2026-08-20, replaces docs/wireframes/
│   │   ├── design-guidelines.md   # visual tokens (moved from docs/project/)
│   │   ├── page-map.md            # per-persona page inventory (moved from docs/project/, evolved)
│   │   └── patterns.md            # reusable page-shape + state-handling + shared component behavior — NEW
│   ├── spec/                     # executable spec, for the agent — see kencleng-agentic-workflow.md
│   │   └── <domain>/              # domain-first: invariants.md, threat-model.md, tasks.md, features/ — one tree per domain (account, notification, organization, campaign, donation, disbursement)
│   └── kencleng-agentic-workflow.md  # process reference doc (lives at docs/ root, not project/ or spec/ — this is process, not product spec)
```

**Note (pre-existing, unrelated to this patch):** the old `docs/spec/`
tree above (`domains/`, `features/`, `threat-model/` as separate
type-first folders) was already superseded by the domain-first layout
(`docs/spec/<domain>/`) per the 2026-08-05 resolution documented in
this same file's §2.1 — this patch also corrects that stale line
since it was touched anyway, but that decision itself is not new here.

## Update §3 "Placement notes" table — add a row

```markdown
| `docs/ui-ux/` | `docs/` (sibling to `project/` and `spec/`) | Frontend UX-specific docs (visual tokens, page inventory, page patterns) — grouped separately from `docs/project/`'s general narrative docs since they're referenced together as a set during frontend work, and separately from `docs/spec/` since they're not per-domain |
```
