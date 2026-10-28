# Kencleng — Privacy, PII & Cryptography Standards (T04)

> **Standing:** Settled/persisted project-wide Engineering policy on research branch `pilot/station5-c1-space-only`; **no runtime implementation or legal compliance claim**.
> **Scope:** Data classification, PII minimization, field encryption, decryption, blind indexing, key management, lifecycle, UI/API masking, and feature-level inventory.
> **Authority:** [Engineering Standards](README.md). Product Authority determines data purpose, required collection, public visibility, and lawful processing decisions. [T03 Authorization](authorization.md) owns access decisions; [T02 Identity](identity-authentication.md) owns authentication/token handling.

## T04-A — Classification

Engineering classification applies to all personal, inferred, and derived data. It is not a replacement for legal categories.

| Class | Examples (illustrative, not required fields) | Default protection |
| --- | --- | --- |
| **Public** | Campaign details approved for publication | Only approved publication, integrity maintained |
| **Internal** | Nonsecret operational metadata | Restricted internal access |
| **Confidential / PII** | Personal email, phone, address, identifiable donation history | Minimize, scoped authorized access, suitable masking, risk-assessed encryption |
| **Restricted** | Personal financial details, identity documents, credentials and cryptographic keys | Stronger isolation, no implicit public/API exposure |

A pseudonymous account identifier or metadata may still identify a person when joined with other information. Public status must come from approved Product semantics, not the existence of a field.

## T04-B — Data minimization and inventory

- Collect, persist, process and share only fields necessary for an approved purpose. A login provider or external service providing extra fields is **not** permission to store them.
- For each personal/sensitive field, document owner/purpose, classification, source, storage/recipients, scope, plaintext requirements, display/exposure, search, retention/deletion, and verification using T04-H.
- Unknown legal basis, privacy expectation, public exposure, retention or erasure exception is a **Product/privacy decision gap**, not an Engineering assumption. Do not add speculative fields for future features.
- Classification also applies to combined data, traceability relationships, data-derived identifiers, and logs.

## T04-C — Risk-based field encryption

- Avoid sensitive storage where possible. Choose protections proportional to threat model and data class. **Not every field needs application-level encryption**; encryption at rest alone may be insufficient where raw database reading is in the threat model.
- For sensitive fields requiring **reversible application-level encryption**, default to **AES-256-GCM (AEAD)** via maintained crypto libraries, with unique nonce per encryption under a key, authenticated metadata/AAD where context binding matters, and versioned stored envelope (key ID, nonce, ciphertext/tag). No homegrown primitives or deterministic encryption for query convenience.
- Plaintext password storage is prohibited. Passwords need appropriate one-way password hashing; authentication/refresh tokens follow T02 verifier/rotation policy, not a generic PII encrypt function.
- Encryption, hashing, blind indexes, and interface masking solve different problems; do not substitute one for another.

## T04-D — Authorized decryption

- Only authorized backend/domain operations can decrypt, according to T03 principal/action/resource/purpose and effective state. Key possession by a process does **not** authorize all its callers.
- Limit plaintext lifetime, propagation and privileged data reads. Avoid sending unnecessary raw values in API DTOs, browser state, logs, traces, metrics, job payloads, errors, exports, or backups.
- Authentication-tag/key failures are fail closed and do not reveal raw data or cryptographic details. Audit sensitive reveal operations proportionally without writing plaintext to audit records.
- Do not decrypt restricted data solely for a UI field that only needs a masked representation.

## T04-E — Blind indexing

- Use **keyed HMAC-SHA-256** blind indexing with a **separate key** and field/purpose domain separation **only when an approved feature needs exact-match lookup** of encrypted data. No mandatory blind indexes for every field.
- Specify safe, domain-appropriate normalization; do not silently merge account identities by email/provider-specific transformations or broad phone rewriting.
- Blind indexes can still expose equality/frequency patterns and enable enumeration through an abused lookup endpoint. Protect query access and consider low-entropy value spaces; plain unsalted SHA-256 is not an acceptable substitute.
- Prefix/fuzzy/substring search over encrypted PII is not a default commitment.

## T04-F — Key custody, rotation and recovery

- Separate keys by purpose: data encryption, HMAC/blind index, JWT signing, and unrelated secrets. Do not hardcode, commit, log, or carelessly expose keys in environment/config.
- Production requires protected custody appropriate to deployment risk (KMS/vault/equivalent if justified); local development must use distinct non-production key material. No vendor choice is mandated.
- Persist identifiable key versions; plan rotation, re-encryption/reindexing, migration, backups, and recovery **before encrypted production data depends on them**. Retired keys cannot be destroyed if retained backups still require them for authorized recovery.
- Suspected compromise requires containment and rekey planning; key replacement does not undo historical plaintext disclosure.
- Root `AGENTS.md` explicit Human approval still applies to cryptographic and key-handling core implementation.

## T04-G — Privacy lifecycle and output surfaces

- Document retention/deletion and lawful/financial exceptions for each real feature; no universal retention period or automatic consent rule is assumed.
- Protection covers database, S3-compatible object storage, jobs, notifications, exports, backups, analytics, application logs, caches, and diagnostics. Avoid raw PII, auth secrets and keys in logs or errors.
- Identity/financial documents in object storage are private by default. Public and presigned links need scoped authorization and bounded lifetime. Verify actual denial for unauthorized recipients.
- Account for backup replicas and caches when implementing deletion; never claim complete erasure without evidence.

## T04-H — Feature-level PII Protection Inventory (agent scaffold)

Agent/Solution Architect must fill this **only for personal/sensitive data actually used** by an approved feature, not as a speculative product-wide data model.

| Object/field | Purpose & owning authority | Class | Storage & recipients | Read/decrypt permission | UI/API full vs masked + reveal | Search/index | Retention/deletion | Test/evidence |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| *(fill per approved feature)* | | | | | | | | |

Document storage treatment (omit/protected plaintext/encrypted/hash), necessary plaintext recipients, key/backup implications, and positive/negative security tests. Missing Product authority is escalated, not invented. A plan is not implementation evidence.

## T04-I — Interface masking, redaction, controlled reveal

- Apply **minimum necessary disclosure** across web pages, shared screens, API responses, exports, notifications and integration outputs. Mask/omit/full representation depends on data sensitivity, authorized actor, purpose and Product disclosure rules. An authenticated user editing their own email can be a legitimate full-view case; public donor pages are not automatically entitled to donor identity.
- **Minimize before sending to browser:** if UI requires only a masked value, the server/API serializer should return only masked data or omit the field. **Do not return full PII in a JSON response and mask only with React/CSS/hidden DOM**. Visual masking is not an authorization boundary. Avoid serializing full DB models to clients.
- Example formats, **not fixed mandates**: `a•••@example.com`, `08••••1234`, `•••• 1234`. Local part, domain or last digits can still leak information. In more sensitive contexts, completely redact/omit instead.
- **Controlled reveal** requires fresh effective backend authorization for actor, resource, purpose and context. Use risk-appropriate reauthentication/step-up and audit when justified. A UI Show button cannot merely unhide data already sent to the browser; bound cache and reveal duration where needed.
- A masked string is **not a canonical editable field**: saving an unchanged form must not overwrite the original with mask characters. Editing full values needs authorized read/write semantics.
- Apply safeguards to aria/screen-reader labels, tooltips, HTML attributes, clipboard/copy, URLs, browser caches, CSV/PDF, email/SMS, errors and telemetry. Keep UI usable and accessible without leaking hidden values.
- Verify raw **network/API payload and rendered UI**, denial/cross-scope accesses, reveal behavior, copy/edit/save, export/notification channels, and accessibility text. A screenshot of masking alone is insufficient.
- Reusable mask components are optional and should be created only after real needs arise. [Product Design standards](../../ui-ux/README.md) own presentation/UX; T04 owns minimum data exposure/security boundaries.

## References / boundary

Supporting guidance: [OWASP Cryptographic Storage](https://cheatsheetseries.owasp.org/cheatsheets/Cryptographic_Storage_Cheat_Sheet.html), [OWASP Key Management](https://cheatsheetseries.owasp.org/cheatsheets/Key_Management_Cheat_Sheet.html), [OWASP Logging](https://cheatsheetseries.owasp.org/cheatsheets/Logging_Cheat_Sheet.html), [OWASP Excessive Data Exposure testing](https://owasp.github.io/www-project-web-security-testing-guide/latest/4-Web_Application_Security_Testing/12-API_Testing/03-Excessive_Data_Exposure).

This standard does **not** require an encryption framework, full schema, new Product fields, KMS vendor, privacy admin UI, or production implementation now. Each protected implementation still needs scoped approval and verification.
