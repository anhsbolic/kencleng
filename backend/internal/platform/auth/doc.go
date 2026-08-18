// Package auth owns the JWT (ES256) and session primitives: token issuance,
// verification, and refresh-token lifecycle.
//
// Tier 0 fencing (root AGENTS.md §3): this package may only be modified by
// an agent with an explicit per-session human ask. The backend scaffold was
// authorized to create the ES256 key-pair loader below and nothing else —
// JWT signing/verification and session logic belong to a human-paired
// session.
package auth
