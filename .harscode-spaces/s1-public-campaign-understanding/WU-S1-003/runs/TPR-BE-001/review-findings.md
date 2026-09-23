# Review findings — WU-S1-003

> Phase             : Techplan Independent Review
> Work Unit         : `WU-S1-003`
> Run               : `TPR-BE-001`
> Role              : Reviewer
> Specialization    : Backend / Campaign / public-data security
> Author            : Codex CLI agent
> Session           : Fresh independent session
> Created           : 2026-09-23
> Target revision   : `9b5dd8bfbef3c0a8815624f0355347ca00ef9f03`
> Workflow revision : `d46358563942c7e015b97aa7c5767c880ef1bc63`

## Review findings — WU-S1-003

**Gate:** Complex — rencana melintasi persistensi PostgreSQL, domain/service,
MinIO private-object, HTTP public contract, migration, dan batas
authorization/non-disclosure media; public data/security boundary ini juga
Tier 1.

**Sections resolved:** Background §1; Scope §2; Requirements §3; Rules &
Validation §4; Decision Log §5; Backward Compatibility §6; Edge Cases & Risks
§7; Interface Contract §8; Architecture / Plan §9; Implementation Details
§10; Files Changed / Files NOT Changed §11; Testing Checklist + Test Focus
Pointer §12; Open Items §13.

### Blocking

- None.

### Non-blocking

- None.

### Clean

- **Rule fidelity:** R1–R8 faithfully retain `INV-campaign-14`, Feature 02/03,
  and the closed OpenAPI public graph: `published` fundraising eligibility,
  exact projection, decimal funding, media membership/private delivery,
  uniform public error classes, and `Cache-Control: private, no-store`.
  Every rule has one or more explicit §12 evidence rows.
- **Decision fidelity:** D1–D7 retain the chosen minimal persisted state,
  explicit service/transport boundary, decimal dependency, private MinIO
  seam, isolated integration evidence, opt-in seed command, and the
  WU-S1-005 topology handoff from `EXP-BE-001/evidence/solutioning.md`.
- **Technical-fact spot checks:** current `backend/cmd/server/main.go`
  confirms `initMinIO` currently verifies both buckets and discards its
  client; `backend/go.mod` has neither a decimal nor `testcontainers-go`
  dependency; `backend/internal/platform/storage/doc.go` is an empty
  placeholder; and `api/openapi/campaign.yaml` confirms the unprefixed backend
  routes, `security: []`, exact closed projection, controlled `/api/...`
  `content_url`, JPEG/PNG-only success, and shared no-store 404/503 responses.
- **Ownership/open items:** §13 correctly keeps WU-S1-005 Caddy/private-policy/
  proxy proof and operator-supplied real seed facts Active but non-blocking for
  backend Build. §11 retains the required root, frontend, API/spec, and Tier-0
  fences.
- **Test Focus Pointer:** all surviving public-boundary, decimal, controlled
  media, and integration-sensitive Exploration risks have exact evidence
  anchors and Testing ownership; the no-new-hot-path concurrency item is
  explicitly N/A rather than dropped.

## Phase handoff

- Completed: independent review gate + review.
- Artifacts: `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/TPR-BE-001/review-findings.md`.
- Human decision: approve the Draft Techplan after the required planning
  resolution/approval control step; no review finding requires revision.
- Open / deferred: topology runtime proof remains with `WU-S1-005`; truthful
  operator seed facts remain external; neither blocks the planned backend Build.
- Recommended next step: one resolution/control pass to record this clean
  review, then Human approval gate.
- Session transition: Build remains fresh-preferred after Approval so the
  implementer reopens current source anchors and the approved Techplan.
- Context pointers: `TP-BE-001/techplan.md`; `EXP-BE-001/evidence/gap-analysis.md`;
  `EXP-BE-001/evidence/solutioning.md`; `docs/spec/4-campaign/invariants.md#inv-campaign-14`;
  `api/openapi/campaign.yaml` public operations and schemas.
