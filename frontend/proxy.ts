import { NextRequest, NextResponse } from "next/server";

// File named `proxy.ts`, not `middleware.ts` — phase0-shared-infra.md
// (Step 4) predates Next.js 16 renaming this file convention
// (`middleware.ts` → `proxy.ts`, same `export const config.matcher`
// shape, exported function renamed `proxy`). Same v-next drift
// pattern as scaffold-frontend.md's Tailwind v3→v4 config note:
// followed the installed framework's actual convention over the
// doc's now-stale name, not a new decision.
//
// Coarse session check only (resolved decision,
// kencleng-frontend-tech-stack.md): this is not the source of truth
// for role-gated access — that needs the actual role data fetched via
// `GET /account/me` (`lib/hooks/use-has-role.ts`), which this can't
// do without an extra network round trip on every navigation. This
// just keeps a request with no session at all from ever reaching a
// `/dashboard/*` page.
//
// "Session indicator" = presence of the refresh-token cookie the
// backend sets on login (`kencleng_refresh`, HttpOnly — readable
// here because this runs server-side, not by client JS). Its
// presence is checked, not its validity — an expired/revoked cookie
// still passes this coarse check and gets caught by the first
// `apiFetch` 401 instead; that's `lib/api/client.ts`'s job, not
// this one's.
const SESSION_COOKIE_NAME = "kencleng_refresh";

export function proxy(request: NextRequest) {
  const hasSession = request.cookies.has(SESSION_COOKIE_NAME);

  if (!hasSession) {
    const loginUrl = new URL("/login", request.url);
    return NextResponse.redirect(loginUrl);
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/dashboard/:path*"],
};
