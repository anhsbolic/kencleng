## Step 0 — Gate question (mandatory, answer explicitly before proceeding)

**Is this techplan Complex-tier?** 

- If **no** — stop here. State why review isn't warranted at this tier and exit. Don't review "on autopilot" just because this prompt was invoked.
- If **yes** — proceed. State which Complex-tier criteria applied.

## Inputs required

1. The synthesized `./techplan.md` (the artifact under review)
2. The raw exploration docs it was synthesized from (`../1-explore`)
3. `rules.md`, `guardrails.md`, `diagram-guidelines.md`, `template.md` — inside /home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/2-techplan — the review is against these, not against the reviewer's own judgment of what a techplan should look like

## Review checklist

Run each check against the actual techplan and actual source docs — not from memory of what a typical techplan contains.

### 1. Rule fidelity (rules.md §4)
- Count rule IDs in §4. Confirm every rule ID has a corresponding line in the testing checklist (§12). Flag any missing.
- For each rule, re-derive it from the raw exploration doc it claims to come from. Flag any rule that doesn't trace back, or that paraphrases the source in a way that changes meaning.

### 2. Diagram validation (diagram-guidelines.md)
- **Syntax**: every edge uses `-->`, no single-dash edges.
- **Semantic**: for every branch condition, re-check against the source table it claims to represent. Confirm the inequality direction is correct and the range is actually satisfiable. Confirm no gaps or overlaps between branches.

### 3. Summary accuracy (template.md Summary section)
- Confirm Summary's Top Risks lists **High-severity only** — flag any Medium or lower.
- Confirm every item in Summary's Open Items list matches the current state of §14 exactly — no item resolved in §14 but still shown Active in Summary, or vice versa.
- Confirm Summary contains no decision not present in §1-13 (guardrails §8: Summary must not introduce new decisions).

### 4. Open Items lifecycle (rules.md §8)
- Every item in §14 is in exactly one state (Active or Resolved) — flag ambiguous or dual-state items.
- Every Resolved item has its resolution recorded, not deleted.

### 5. Guardrail compliance spot-check
- No invented facts — spot-check 2-3 non-obvious technical claims against source docs.
- No silent overwrite of a locked (Approved/Implemented) contract, or of Draft-status content mid-edit without Summary resync (guardrails §11).

## Output format

Do **not** output a rewritten techplan. Output a findings report:

```
## Review findings — <story-code>

**Gate check:** [Complex-tier criteria that applied]

### Blocking (must fix before Approved)
- [finding] — [location in techplan] — [what's wrong] — [what source says instead]

### Non-blocking (worth a look, doesn't block approval)
- [finding]

### Clean
- [checklist items that passed, briefly — confirms the check was actually run, not skipped]
```

## What happens with findings

Findings go back to the primary model (or the human lead) for resolution — this prompt does not auto-fix the techplan. If the same category of finding recurs across 2+ stories, that's the proposal threshold (`guidelines.md`) — write a proposal to fold it into `rules.md`/`guardrails.md`/`diagram-guidelines.md` as a permanent self-check item, the same way proposal 0003 converted recurring prose instructions into checklist items.

## Explicitly out of scope

- Re-litigating architectural decisions already made in §5 (Decision Log) — this is a compliance/consistency review, not a second design review.
- Style or wording nitpicks — not a proposal-worthy friction on their own (see `guidelines.md` Proposal Threshold).