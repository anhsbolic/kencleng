# Kencleng Frontend

Next.js (PWA) frontend for Kencleng. See the [root README](../README.md)
for what Kencleng is and how to run the full stack via Docker Compose.

This README is for working **inside this folder specifically** — e.g.
running the dev server, running tests, regenerating API types —
without necessarily spinning up the whole stack.

## Stack

Next.js (App Router), TypeScript, Tailwind CSS, Zustand, TanStack
Query, `react-hook-form` + `zod`, manual PWA setup. Full rationale in
`../docs/project/kencleng-frontend-tech-stack.md`.

## Layout

```
frontend/
├── app/            # Next.js App Router — routes, layouts, pages
├── components/     # shared UI components
├── lib/api/        # types generated from ../api/openapi.yaml — do not hand-edit
└── public/         # static assets + PWA manifest
```

See `AGENTS.md` in this folder for Next.js/TypeScript-specific coding
conventions — including the "no business logic in the frontend"
boundary, which is a hard rule, not a style preference.

## Prerequisites

- Node.js 20+
- The backend reachable — either via the root `docker-compose up -d`
  (recommended, and required for the same-origin/`SameSite=Strict`
  auth flow to behave correctly), or your own local backend instance

## Running standalone

```bash
npm install
npm run dev
```

Opening the app at `localhost:3000` directly (bypassing the Caddy
proxy) will not correctly exercise the auth flow, since it depends on
the frontend and backend sharing an origin. For anything auth-related,
go through `http://localhost` via the root `docker-compose` setup.

## Regenerating API types

Whenever `../api/openapi.yaml` changes:

```bash
npx openapi-typescript ../api/openapi.yaml -o lib/api/schema.d.ts
```

Don't hand-edit the generated file — regenerate it instead.

## Testing

```bash
npm run test    # component/unit tests via vitest
npm run lint
```

## Verification

```bash
npm run verify
```

Runs lint + tests. See `../docs/kencleng-agentic-workflow.md` §7-8 for
the full gate sequence across the whole stack.