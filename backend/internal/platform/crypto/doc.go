// Package crypto owns the PII encryption-at-rest primitives: AES-GCM
// ciphertext plus HMAC-SHA256 lookaside hashes, per the pattern documented
// in docs/project/kencleng-backend-tech-stack.md.
//
// Tier 0 fencing (root AGENTS.md §3): this package may only be modified by
// an agent with an explicit per-session human ask. The backend scaffold was
// authorized to create the key holder below and nothing else — the actual
// encrypt/decrypt/HMAC helpers belong to a human-paired session.
package crypto
