// Package crypto owns the PII encryption-at-rest primitives: AES-GCM
// ciphertext plus HMAC-SHA256 lookaside hashes, per the pattern documented
// in docs/project/kencleng-backend-tech-stack.md.
//
// Tier 0 fencing (root AGENTS.md §3): this package may only be modified by
// an agent with an explicit per-session human ask. The key holder (keys.go)
// was created by the backend scaffold. The encrypt/decrypt/HMAC helpers
// (crypto.go) were agent-authored under an explicit per-session §3 fence
// lift, reviewed by the human in that session. Future modifications still
// require a new per-session human ask — the lift was one-time, not
// permanent.
package crypto
