# Harscode Handoff Staging

This directory is a **temporary handoff surface** for changes that belong in the separate `anhsbolic/harscode-workspace` repository.

It is intentionally non-authoritative for Kencleng.

## Rules

- Kencleng production code, project policy, specs, and runtime behavior must not depend on anything under `harscode-handoff/`.
- `full-files/` mirrors the final target paths in `harscode-workspace` so files can be copied without reconstructing snippets from chat.
- `PROPOSAL.md` explains the reasoning and review posture for the handoff.
- `MANIFEST.md` is the copy checklist and target-path map.
- Harscode remains authoritative for whether a proposal number is still available and whether a proposal is accepted.
- Before applying a handoff, compare it against the current Harscode `main`; if Harscode moved materially, reconcile rather than blindly copy.
- A proposal-only handoff may intentionally stage only the Harscode proposal/changelog entry. Policy-file edits should be produced in Harscode only after that proposal is reviewed/accepted.
- After the Harscode change is applied and verified, remove the corresponding handoff directory from Kencleng. Git history is sufficient if the staging package ever needs to be inspected later.

## Current Harscode baseline used for staged proposals

Unless a proposal says otherwise, the handoff was prepared against:

```text
anhsbolic/harscode-workspace
main @ 5f3c9bcf0d9a14d3ae72314864eab982ffa4947e
```

That baseline was read-only during preparation. No Harscode files were modified from Kencleng.
