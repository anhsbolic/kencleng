# Validation #1 Retrospective — Frontend Experience Foundation

> Phase: CRTV retrospective / operator evidence
> Author: ChatGPT
> Model: GPT-5.6 Sol
> Created: 2026-09-15
> Target revision reviewed: `341767481d3f49b4822ca9be7c5d5c176f7d1d58`
> Kencleng run baseline: `4ef5a50caa89e58e2cc2a5feea3e693a17bf8c3f`
> Harscode workflow used by the run: `workflow-v2@fce5721d996b013c7c9357745978cd9511592a78`
> Status: retrospective evidence; not a product/spec authority and not a required input to future workflow phases

## Outcome

The first Continuous Real-Task Validation run completed Exploration → Techplan → Build → Code Review → Build/Patch → Testing → Build/Patch → Testing with durable artifacts and a satisfactory correctness outcome.

The run showed that workflow-v2's progressive context/handoff model works, but also exposed operator and efficiency friction around provenance, human-facing handoff wording, optional Techplan routing, repeated verification, and benchmark capture.

The detailed cross-workspace learning and accepted Harscode changes are preserved in:

- `harscode-workspace/audits/workflow-v2-validation-01-retrospective.md`
- `harscode-workspace/proposals/0032-workflow-v2-provenance-handoff-and-harness-observability.md`
- `harscode-workspace/proposals/0033-workflow-v2-proportional-techplan-and-verification-ownership.md`

The post-run candidate refinement checkpoint is frozen separately in Harscode after those changes are verified.

## Project-side lessons

### Engineering workflow

- Independent Code Review and Testing both found real signal, so they remain valuable.
- The main efficiency problem was repeated broad verification across Build/Review/Patch/Testing, not Exploration/Techplan correctness.
- Techplan independent review and decomposition should remain optional and should be recommended early rather than invoked as ritual.
- Human decisions should be surfaced explicitly at phase boundaries rather than mixed with ordinary deferred/open items.

### Product / UI design

The implementation followed the approved conservative contract correctly. Exploration/Techplan explicitly avoided inventing a new brand-defining visual direction or asset system, so Build retaining much of the existing visual character was expected behavior rather than a Build failure.

The human owner wants a substantially stronger creative/brand direction before the next frontend engineering iteration. That creative work should happen upstream in a dedicated product/brand/design exploration session, then become a durable design brief/spec before Harscode engineering phases begin.

Harscode should not be expanded into a separate creative-design lifecycle for this purpose.

## Next frontend iteration intent

The next frontend validation should **not** treat the current frontend component/layout implementation as the foundation to incrementally preserve.

After the new brand/design direction is prepared and consciously accepted, the human owner intends to restart frontend implementation from a deliberately clean foundation — effectively from zero at the component/layout/application-shell level where appropriate — because the new design may differ substantially from the current implementation.

This is a future baseline/setup decision, not authorization to delete/reset frontend code now.

Before that validation begins:

1. finish the upstream brand/product UI exploration;
2. convert the accepted direction into durable Kencleng design/spec authority;
3. decide the exact clean-start boundary (what is deleted/recreated versus retained as runtime/tooling infrastructure);
4. prepare and freeze a new Kencleng validation baseline;
5. use the new frozen Harscode workflow-v2 candidate for the next real task/run.

Do not reuse Validation #1 chat conclusions or implementation details as hidden input to the next workflow run unless the new canonical project authorities independently route to them.

## Benchmark note

`benchmarks.md` remains the raw session-usage evidence for this run. Its limitations are known: telemetry began after the ideal pre-Exploration point, and the first session combined Exploration + Techplan + decomposition-gate activity.

Future validation should capture benchmark/status state before the first workflow session and after every measured agent session, while keeping telemetry outside correctness authority.
