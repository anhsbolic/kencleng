Read .local-agents/works/account/02-google-oauth-login-register/2-plan/techplan.md

### 0. Answer first: is it worth decomposing this techplan?

> **Is it worth it to decompose this techplan?**

Before doing anything else, read the techplan and answer this question
explicitly — state the answer and the reasoning behind it. Do not proceed
to Step 1 on autopilot just because this prompt was invoked.

Answer **no** (stop here, do not decompose) if:

- The techplan is small/linear enough that an execution agent can run
  through it start to finish without losing focus.
- Splitting it would mostly produce single-task manifests or trivial
  boundaries — decomposition for its own sake.

Answer **yes** (proceed to Step 1) only if at least one of the trigger
signals from "When To Use This" genuinely applies to this specific
techplan.

If the answer is no, stop and report that decomposition isn't warranted,
with a one-line reason. Don't generate task files or a manifest just
because the prompt was run.

### 1. Read the techplan in full

Don't start splitting before understanding the overall scope and the
dependencies between its parts.

### 2. Choose a splitting axis

Pick one of the following options (or a combination, if the techplan is
complex enough), based on the characteristics of the techplan you read:

- **Dependency/sequence chain** — tasks are split based on hard
  dependencies; task B cannot start until task A is done. Choose this when
  the techplan has stages that genuinely wait on each other.
- **Component/module boundary** — tasks are split along independent code
  boundaries (different package/service/domain). Choose this when there
  are multiple modules that don't technically depend on each other.
- **Layer (vertical slice)** — data layer → business logic → interface/API
  layer as separate tasks. Choose this when a single change cuts across
  multiple layers and each layer has a different review concern.
- **Review-able/PR-sized chunk** — tasks are split so each task equals one
  PR that can be merged independently. Choose this when the main goal is
  reviewability rather than execution efficiency.
- **Risk/blast-radius** — tasks are split so sensitive changes
  (security-critical, data-destructive, breaking-change) are separated
  from low-risk changes. Choose this when the techplan mixes high-risk and
  low-risk changes within a single scope.

**When it's ambiguous, default to dependency/sequence chain** — this axis
carries the least assumption compared to the others, especially
risk/blast-radius, which requires additional judgment.

State the chosen axis and a short rationale before moving to the next
step — this will go into the manifest.

### 3. Redistribute detail into task files

For each resulting task:

- Copy/redistribute the derived-section detail relevant to that task's
  scope — do not summarize it.
- Include a back-reference to the originating contract techplan
  (path/title), so the execution agent can cross-check high-level
  decisions whenever needed.
- Each task should be executable without needing to read other tasks
  first, except where dependency is genuinely required by the chosen
  splitting axis.

### 4. Generate the manifest

The manifest is a snapshot-at-generation-time, not a living document — it
does not track progress status (done/in-progress/blocked). Status tracking
belongs to the PR/ticket domain, not this workspace.

Minimum manifest contents:

- List of task files + a short title for each.
- The splitting axis used (+ brief rationale from step 2).
- Dependency graph between tasks (if the axis is sequence/dependency) — or
  an explicit "no hard dependency" marker if the axis is parallel
  (component/module).
- Back-reference to the originating contract techplan.
- llm models to execute each tasks, you can read /home/anhar-solehudin/kencleng-workspace/harscode-workspace/best-practices/model-routing.md

## Output

- N task files (count depends on the resulting split), each self-contained
  for its scope.
- 1 manifest file indexing all the tasks above.
- write files in .local-agents/works/account/02-google-oauth-login-register/2-plan/tasks

## Cross-reference

- Source techplan: see `../../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md` for the
  contract/derived authoring process that precedes this decomposition
  step.
- Task files produced here are still subject to the review checklist in
  `../../harscode-workspace/workflow/4-code-review/checklist.md`, same as any other derived-section content.