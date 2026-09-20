# Kencleng — Project Communication Profile

Status:
Pilot-active on `validation-03-orchestrator-slice-1`

Default human-facing language:
Bahasa Indonesia

## Communication Rules

Gunakan Bahasa Indonesia untuk:

- penjelasan human-facing;
- rationale;
- handoff prose;
- report narrative;
- Finding / Decision descriptions;
- Control Surface prose;
- orchestration artifact prose.

Pertahankan dalam English ketika memiliki makna canonical atau presisi teknis:

- canonical Harscode terms, termasuk `Work Unit`, `Run`, `Finding`, `Blocker`, `Decision`, `Authority Decision`, `Delivery Decision`, `Participant`, `Session`, `Control Surface`, dan `Harscode Space`;
- protocol enums/status/type values seperti `ACTIVE`, `BLOCKED`, `WAITING_HUMAN`, `ENABLER`, `DELIVERY`, `RECONCILIATION`, `VERIFICATION`, `QUEUED`, `DISPATCHED`, dan `RUNNING`;
- technical terms ketika terjemahan Indonesia mengurangi presisi atau menimbulkan terminology drift.

Jangan menerjemahkan:

- code symbols;
- API fields;
- schema names;
- CLI commands;
- file paths;
- branch names;
- commit SHAs;
- identifiers;
- exact externally-defined contract terms.

Jika exact wording dari authority source materially important, pertahankan wording sumber apa adanya.

## Compact Directive for Runs

```text
Human-facing prose: Bahasa Indonesia.
Preserve canonical Harscode terms/enums and code/API/schema identifiers in English.
```

Normal Run cukup memakai compact directive di atas. Dokumen ini dibaca penuh hanya ketika language handling ambigu.
