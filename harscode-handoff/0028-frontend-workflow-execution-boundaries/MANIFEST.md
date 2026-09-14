# Manifest — Harscode 0028 Frontend Workflow Execution Boundaries

Prepared against Harscode main `5f3c9bcf0d9a14d3ae72314864eab982ffa4947e`.

This is intentionally a **proposal-only** handoff. No existing Harscode policy file is staged for replacement yet.

## Target operation

```text
CREATE  proposals/0028-frontend-workflow-execution-boundaries.md
```

The staged proposal identifies the candidate Harscode files to edit if the proposal is accepted:

```text
workflow/README.md
workflow/2-2-techplan-review-prompt.md
harness-optimization/codex/README.md
best-practices/react/testing-automation-boundary.md
```

Those files are **not** included under `full-files/` in this handoff. Their exact edits should be produced in Harscode only after proposal review, because:

- the independent Techplan review prompt is still Draft;
- its own evidence threshold requires more real dogfood before formalization;
- generic model routing is being revised separately;
- this dogfood did not complete the full Build → Review → Testing → PR lifecycle.

## Copy source

```text
full-files/proposals/0028-frontend-workflow-execution-boundaries.md
→ proposals/0028-frontend-workflow-execution-boundaries.md
```

## Apply order

1. Re-check Harscode `main` and confirm proposal number `0028` is still available.
2. Compare the staged proposal against any Harscode changes made after baseline `5f3c9bcf0d9a14d3ae72314864eab982ffa4947e`.
3. Copy the staged proposal into `proposals/0028-frontend-workflow-execution-boundaries.md`.
4. Review/accept/reject it using Harscode's normal proposal process.
5. Only if accepted, make the smallest policy-file edits described by the proposal inside the Harscode repository.
6. Verify that no target-project model names, commands, paths, or lifecycle fork leaked into portable Harscode policy.
7. Once the Harscode follow-up is settled, remove this staging directory from Kencleng.

## Expected review checks

Before accepting the Harscode proposal, confirm:

- canonical phase prompts remain the single lifecycle invocation authority;
- no parallel frontend/backend lifecycle is introduced;
- execution profiles remain separate from lifecycle policy;
- generic Harscode does not inherit a project-specific Codex model table;
- the draft Techplan review mechanism gains a stopping/convergence rule without being prematurely promoted to a hard gate;
- browser automation remains optional tooling rather than a phase ritual;
- objective verification requirements are not weakened;
- the proposal still reflects limited evidence from one planning/review dogfood rather than claiming a fully proven lifecycle.
