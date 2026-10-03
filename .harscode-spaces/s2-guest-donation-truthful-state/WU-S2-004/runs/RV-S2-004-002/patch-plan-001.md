Phase: Code Review Patch Plan
Author: P-S2-004-RV-002-1 (Reviewer / KC-REVIEWER)
Created: 2026-10-03
Model / Reasoning: Invocation-configured gpt-6-luna / high; active runtime values not independently exposed
Session: Not independently exposed
Target revision: bb69cd002b3f1a1056837affcd77bb2b001007b0 plus current working tree
Workflow revision: Harscode pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4 (ordinary applicable guidance current-effective)
Work Unit / Run: WU-S2-004 / RV-S2-004-002
Source review: `review-findings-001.md`; exact reviewed source set is recorded there.

# Patch Plan

Build/Patch owns implementation. Do not change Product, API, spec, or test acceptance criteria to accommodate these findings.

## Required fixes

1. **F01 — Associate the email error with its input.** In `frontend/app/campaigns/[campaignId]/donate/donation-client.tsx`, add an ID to the rendered `guest_email` validation message and conditionally reference it from the email input with `aria-describedby`. Keep `aria-invalid` tied to the same error state.
2. **F02 — Announce asynchronous status results.** In `frontend/app/donations/[donationId]/status/status-client.tsx`, provide a stable polite live region (or an equivalent accessible announcement pattern) for initial and manual status-check outcomes, including pending, terminal, and unavailable states. Keep the recheck control usable and do not move focus unexpectedly.

## Acceptance for the patch

- The email field’s accessible description exposes its validation message when invalid and no stale error description when valid.
- A completed manual status check announces the resulting pending, success, failed, or unavailable state to assistive technology.
- Existing truthful wording, generic unavailable behavior, fragment cleanup, request header, and status-only presentation remain intact.
- Build reports only its scoped implementation feedback. Independent Testing remains responsible for the approved R8 browser check and relevant accessible behavior coverage; R10 Human rendered acceptance remains separate.

## Handoff

Create a new Build/Patch Run/Participant/Session under Orchestration, re-ground it on this plan and the approved Techplan, then return the scoped Build evidence for orchestration routing. No production change was made during Review.
