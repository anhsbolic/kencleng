# Patch — Housekeeping: stray filename references (2026-08-20)

> Cosmetic only — content of each reference is still accurate, just
> the filename changed. Apply via find/replace where noted.

## `kencleng-phase1-detail.md`

Line ~245: `(see kencleng-ux-page-map.md ...)` → `(see docs/ui-ux/page-map.md ...)`
Line ~269: `kencleng-roadmap-next-steps.md and kencleng-ux-page-map.md` →
`kencleng-roadmap-next-steps.md and docs/ui-ux/page-map.md`

## `kencleng-erd.md`

Line ~782: `kencleng-roadmap-next-steps.md and kencleng-ux-page-map.md` →
`kencleng-roadmap-next-steps.md and docs/ui-ux/page-map.md`

## `api/openapi.yaml` (or `api/openapi/campaign.yaml` post-split)

Line ~855: `(/campaign/[id] per kencleng-ux-page-map.md)` →
`(/campaign/[id] per docs/ui-ux/page-map.md)`

No content changes anywhere in this patch — every reference still
points at the same page-map content, just at its new path.
