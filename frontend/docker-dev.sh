#!/bin/sh
set -eu

# Existing named volumes outlive image rebuilds. Refresh when lock/runtime changes.
fingerprint="$(node --version):$(sha256sum package.json package-lock.json | sha256sum | cut -d ' ' -f 1)"
stamp=node_modules/.kencleng-dev-deps
if [ ! -f "$stamp" ] || [ "$(cat "$stamp")" != "$fingerprint" ]; then
    npm ci
    printf '%s\n' "$fingerprint" > "$stamp"
fi

if [ "${1:-}" = '--install' ]; then
    exit 0
fi

# Webpack polling is deliberate here; the native host runner stays on Turbopack.
exec ./node_modules/.bin/next dev --hostname 0.0.0.0 --webpack
