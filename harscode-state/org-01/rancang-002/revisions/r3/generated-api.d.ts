export interface paths {
    "/auth/google/attempts": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * createGoogleAttempt
         * @description Body <=1024 bytes. {} normal login uses exact Origin+JSON, no session CSRF. purpose=deleteAccount additionally requires AccessCookie + X-CSRF-Token; security alternatives here express conditional branches, not optional proof auth. At most one browser attempt: issuing another replaces binding and invalidates prior attempt. Attempt expires5m through final callback effect. No organization effect. Normative state transitions, evaluation order, cookie transport and conformance cases: interface-decisions.md. Collection behavior: not applicable unless this is listOrganizations; no hidden filters/sort or total count.
         */
        post: operations["createGoogleAttempt"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/auth/google/callback": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * completeGoogleAttempt
         * @description Protocol GET exception; requires exactly one state plus exactly one of code/error. Callback is the sole unknown-query exception: ignore bounded provider extension parameters (including scope/authuser/prompt/error_description/error_uri) as non-authoritative; never use as claims, reflect or log them. Bound total raw request-target8192 bytes, <=16 query pairs, decoded names<=64bytes, extension values<=2048bytes. Duplicate state/code/error, malformed encoding/oversize or any validation/provider/fence failure ->303 /auth/failure, no JSON and no provisioning effect. Consume attempt once before exchange; recheck five-minute final-effect expiry after network/lock waits. Success normal ->/auth/complete; special same-account proof ->/account-deletion/review. No arbitrary returnTo, no query tokens on redirects. Always clear attempt cookie; only success issues app cookies. On special failure invalidate prior relevant deletion previews; no account creation/linking. Normative state transitions, evaluation order, cookie transport and conformance cases: interface-decisions.md. Collection behavior: not applicable unless this is listOrganizations; no hidden filters/sort or total count.
         */
        get: operations["completeGoogleAttempt"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/auth/session": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * getSession
         * @description Access-first inspection. Absent/expired access alone permits unused RefreshCookie bootstrap without consuming/rotating; malformed/signature-invalid or revoked/fenced/account-mismatch access forbids fallback. Neither branch grants org authority via refresh. Both validate current account/fence/session and idle/absolute expiry. Read only: no idle extension, no cookies issued. Do not reveal invalidity subtype. Normative state transitions, evaluation order, cookie transport and conformance cases: interface-decisions.md. Collection behavior: not applicable unless this is listOrganizations; no hidden filters/sort or total count.
         */
        get: operations["getSession"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/auth/refresh": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * refreshSession
         * @description Body <=1024 bytes. Effective refresh verifier+session and CSRF; atomic single-use rotation +access issuance. Consumed-token replay after valid CSRF revokes entire family ->401. Concurrent loser may revoke winner; lost response503 result-unknown, no blind old-token retry; explicit Google flow may be necessary. Access cookie does not substitute for refresh. Normative state transitions, evaluation order, cookie transport and conformance cases: interface-decisions.md. Collection behavior: not applicable unless this is listOrganizations; no hidden filters/sort or total count.
         */
        post: operations["refreshSession"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/auth/logout": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * logoutSession
         * @description Body <=1024 bytes. Origin+JSON required for every branch. Effective session resolved from validated access or refresh verifier; when both identify different sessions ->401 no mutation. Live session requires matching CSRF; revoke durable then204 clear cookies. Already terminated/absent/invalid cookies have no effective authority:204 clear cookies after Origin+JSON, no CSRF required. Explicit fenced/deleted account likewise no new authority/effect and204 cookie-clear only. Unknown DB/journal state ->503 unavailable, not assumed absent; ambiguous revoke ->503 result-unknown. This no-authority branch is harmless repeat, not a CSRF bypass for live sessions. Normative state transitions, evaluation order, cookie transport and conformance cases: interface-decisions.md. Collection behavior: not applicable unless this is listOrganizations; no hidden filters/sort or total count.
         */
        post: operations["logoutSession"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/organizations": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * listOrganizations
         * @description Primary current authorized list only, filter ownership before keyset limit. Order created_at DESC,id DESC. Cursor opaque position only, never authority; changed data can disappear or new higher rows be absent from later pages, no snapshot promise. Unknown query parameters ->400. No name search/total count. Normative state transitions, evaluation order, cookie transport and conformance cases: interface-decisions.md. Collection behavior: not applicable unless this is listOrganizations; no hidden filters/sort or total count.
         */
        get: operations["listOrganizations"];
        put?: never;
        /**
         * createOrganization
         * @description Body <=4096 bytes. Authenticated current account is assigned Owner; never client role/account. UUIDv7 key per intended submission, timestamp <=server now+5m; absent key admission only now<key+24h rechecked after locks. Atomic org+Owner+receipt; verified commit201 and exact replay200. Receipt expiry=succeeded_at+48h immutable. Canonical same name replay only; same name/different keys or accounts valid distinct IDs. Unknown outcome preserves same key/account/name. Normative state transitions, evaluation order, cookie transport and conformance cases: interface-decisions.md. Collection behavior: not applicable unless this is listOrganizations; no hidden filters/sort or total count.
         */
        post: operations["createOrganization"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/organization-creation-requests/{requestId}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * getOrganizationCreationResult
         * @description Primary receipt join current organization+Owner. Readable only now<expires_at, never extends retention. Miss/expiry/unowned ->404 result-not-observed; original may still commit. Corrupt pair ->503 result-unavailable. No fake pending row/202. Normative state transitions, evaluation order, cookie transport and conformance cases: interface-decisions.md. Collection behavior: not applicable unless this is listOrganizations; no hidden filters/sort or total count.
         */
        get: operations["getOrganizationCreationResult"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/organizations/{organizationId}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * getOrganization
         * @description Primary current pair, authorized own relationship. Missing or unowned404 uniform resource-unavailable; invariant failure503 result-unavailable; never partial representation. Normative state transitions, evaluation order, cookie transport and conformance cases: interface-decisions.md. Collection behavior: not applicable unless this is listOrganizations; no hidden filters/sort or total count.
         */
        get: operations["getOrganization"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/account-deletion-previews": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * createAccountDeletionPreview
         * @description Body <=1024 bytes. Require effective access and fresh explicit deleteAccount flow started<=5m ago, same account. Prepared expiry10m; one per session replaces prior preview. Locked complete simple dependency set;64KiB bound; empty set valid. No journal intent, acceptance or closure. Proof/session switch invalidates prior preview. Normative state transitions, evaluation order, cookie transport and conformance cases: interface-decisions.md. Collection behavior: not applicable unless this is listOrganizations; no hidden filters/sort or total count.
         */
        post: operations["createAccountDeletionPreview"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/account-deletions": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /**
         * deleteAccount
         * @description Body <=1024 bytes. Require access, CSRF, same-session prepared preview unexpired and explicit true; purpose proof started<=10m. Compare current exact set and eligibility before intent. Missing/wrong/expired preview409; stale proof409; set change409; complex dependency without accepted intent422. Durable exact intent append precedes one active erasure transaction; no cancellation thereafter. Confirmed pending503 deletion-pending; append/commit uncertainty503 result-unknown. Concurrent confirms use the one preview-derived intent. After durable fence a user retry is401; only originally authorized in-flight confirm may report503 pending/unknown, and trusted reconciliation applies same intent, never a second intent. After erase old session401 is not failure evidence.200 only verified active erase, expire app+attempt cookies. No status GET/public UUID/email or retained session. Normative state transitions, evaluation order, cookie transport and conformance cases: interface-decisions.md. Collection behavior: not applicable unless this is listOrganizations; no hidden filters/sort or total count.
         */
        post: operations["deleteAccount"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
}
export type webhooks = Record<string, never>;
export interface components {
    schemas: {
        /** @example {} */
        Empty: Record<string, never>;
        /**
         * @description Exactly {} for normal login; exactly {purpose: deleteAccount} requires access cookie and X-CSRF-Token. No null/default purpose accepted.
         * @example {}
         */
        GoogleAttempt: {
            /** @enum {string} */
            purpose?: "deleteAccount";
        };
        /**
         * @example {
         *       "id": "b86bda16-6c68-4d13-85fb-3f7729d95445",
         *       "name": "Organisasi K",
         *       "createdAt": "2026-10-30T00:00:00Z"
         *     }
         */
        Organization: {
            /** Format: uuid */
            id: string;
            name: string;
            /**
             * Format: date-time
             * @description UTC RFC3339 instant; never null.
             */
            createdAt: string;
        };
        /**
         * @example {
         *       "accountId": "b86bda16-6c68-4d13-85fb-3f7729d95445",
         *       "role": "Owner",
         *       "scope": "internal"
         *     }
         */
        Owner: {
            /** Format: uuid */
            accountId: string;
            /** @enum {string} */
            role: "Owner";
            /** @enum {string} */
            scope: "internal";
        };
        /**
         * @example {
         *       "organization": {
         *         "id": "b86bda16-6c68-4d13-85fb-3f7729d95445",
         *         "name": "Organisasi K",
         *         "createdAt": "2026-10-30T00:00:00Z"
         *       },
         *       "owner": {
         *         "accountId": "b86bda16-6c68-4d13-85fb-3f7729d95445",
         *         "role": "Owner",
         *         "scope": "internal"
         *       }
         *     }
         */
        OrganizationPair: {
            organization: components["schemas"]["Organization"];
            owner: components["schemas"]["Owner"];
        };
        /**
         * @example {
         *       "data": {
         *         "organization": {
         *           "id": "b86bda16-6c68-4d13-85fb-3f7729d95445",
         *           "name": "Organisasi K",
         *           "createdAt": "2026-10-30T00:00:00Z"
         *         },
         *         "owner": {
         *           "accountId": "b86bda16-6c68-4d13-85fb-3f7729d95445",
         *           "role": "Owner",
         *           "scope": "internal"
         *         }
         *       }
         *     }
         */
        OrganizationResponse: {
            data: components["schemas"]["OrganizationPair"];
        };
        /**
         * @example {
         *       "data": {
         *         "requestId": "019a3438-9000-7000-8000-000000000001",
         *         "organization": {
         *           "id": "b86bda16-6c68-4d13-85fb-3f7729d95445",
         *           "name": "Organisasi K",
         *           "createdAt": "2026-10-30T00:00:00Z"
         *         },
         *         "owner": {
         *           "accountId": "b86bda16-6c68-4d13-85fb-3f7729d95445",
         *           "role": "Owner",
         *           "scope": "internal"
         *         }
         *       }
         *     }
         */
        CreationResponse: {
            data: {
                /** Format: uuid */
                requestId: string;
                organization: components["schemas"]["Organization"];
                owner: components["schemas"]["Owner"];
            };
        };
        /**
         * @description Omit meta entirely on final page; no meta:null or empty meta. Each data item is a pair, with no nested data envelope.
         * @example {
         *       "data": [
         *         {
         *           "organization": {
         *             "id": "b86bda16-6c68-4d13-85fb-3f7729d95445",
         *             "name": "Organisasi K",
         *             "createdAt": "2026-10-30T00:00:00Z"
         *           },
         *           "owner": {
         *             "accountId": "b86bda16-6c68-4d13-85fb-3f7729d95445",
         *             "role": "Owner",
         *             "scope": "internal"
         *           }
         *         }
         *       ]
         *     }
         */
        OrganizationList: {
            data: components["schemas"]["OrganizationPair"][];
            meta?: {
                nextCursor: string;
            };
        };
        /**
         * @example {
         *       "data": {
         *         "accountId": "b86bda16-6c68-4d13-85fb-3f7729d95445",
         *         "csrfToken": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
         *       }
         *     }
         */
        SessionResponse: {
            data: {
                /** Format: uuid */
                accountId: string;
                csrfToken: string;
            };
        };
        /**
         * @example {
         *       "data": {
         *         "authorizationUrl": "https://accounts.google.com/o/oauth2/v2/auth?scope=openid&state=example&code_challenge=example"
         *       }
         *     }
         */
        AttemptResponse: {
            data: {
                /**
                 * Format: uri
                 * @description HTTPS accounts.google.com/o/oauth2/v2/auth only, server constructed, openid scope only; no application tokens or returnTo.
                 */
                authorizationUrl: string;
            };
        };
        /**
         * @example {
         *       "name": "Organisasi K"
         *     }
         */
        CreateOrganization: {
            /** @description Trim Unicode White_Space at edges; canonical result1..200 Unicode code points, reject Unicode Cc controls; preserve case/internal whitespace. Semantic invalid ->422 validation; UTF-8 body>4096 bytes ->400 invalid-request. No uniqueness rule. */
            name: string;
        };
        /**
         * @description Complete sorted-by-ID dependency set, including empty list, total UTF-8 response <=65536 bytes. No pagination/truncation; otherwise422 manual-review-required and no actionable preview. prepared only, not accepted intent. Owner internal meaning and residual-copy explanation are normative UI obligations in interface-decisions.md.
         * @example {
         *       "data": {
         *         "previewId": "b86bda16-6c68-4d13-85fb-3f7729d95445",
         *         "expiresAt": "2026-10-30T00:00:00Z",
         *         "organizations": [],
         *         "effect": "closeContextsAndDeleteAccount",
         *         "backupCopies": "pendingExpiry"
         *       }
         *     }
         */
        DeletionPreviewResponse: {
            data: {
                /** Format: uuid */
                previewId: string;
                /**
                 * Format: date-time
                 * @description UTC RFC3339 instant; never null.
                 */
                expiresAt: string;
                organizations: {
                    /** Format: uuid */
                    id: string;
                    name: string;
                }[];
                /** @enum {string} */
                effect: "closeContextsAndDeleteAccount";
                /** @enum {string} */
                backupCopies: "pendingExpiry";
            };
        };
        /**
         * @example {
         *       "previewId": "b86bda16-6c68-4d13-85fb-3f7729d95445",
         *       "closeOrganizations": true
         *     }
         */
        DeleteAccount: {
            /** Format: uuid */
            previewId: string;
            /** @enum {boolean} */
            closeOrganizations: true;
        };
        /**
         * @description 200 only verified active DB commit. noneVerified only with actual disposal evidence; no allCopiesDeleted field. No request/status bearer capability remains.
         * @example {
         *       "data": {
         *         "accessEnded": true,
         *         "activeDataErased": true,
         *         "remainingCopies": "pendingVerification"
         *       }
         *     }
         */
        DeletionResponse: {
            data: {
                /** @enum {boolean} */
                accessEnded: true;
                /** @enum {boolean} */
                activeDataErased: true;
                /** @enum {string} */
                remainingCopies: "pendingVerification" | "noneVerified";
            };
        };
        /**
         * @description Malformed JSON, duplicate keys, unknown/missing fields, invalid type/UUID/header/query, unsupported content type, or input size bound exceeded. No detail, instance, code, correlation ID, payload echo or additional members on this bounded interface; match type, never parse title.
         * @example {
         *       "type": "urn:kencleng:problem:invalid-request",
         *       "title": "Invalid request",
         *       "status": 400
         *     }
         */
        ProblemInvalidRequest: {
            /** @enum {string} */
            type: "urn:kencleng:problem:invalid-request";
            /** @enum {string} */
            title: "Invalid request";
            /** @enum {integer} */
            status: 400;
        };
        /**
         * @description Required authority is absent, expired, invalid, revoked, mismatched or fenced. Does not determine prior mutation outcome. No detail, instance, code, correlation ID, payload echo or additional members on this bounded interface; match type, never parse title.
         * @example {
         *       "type": "urn:kencleng:problem:unauthenticated",
         *       "title": "Authentication required",
         *       "status": 401
         *     }
         */
        ProblemUnauthenticated: {
            /** @enum {string} */
            type: "urn:kencleng:problem:unauthenticated";
            /** @enum {string} */
            title: "Authentication required";
            /** @enum {integer} */
            status: 401;
        };
        /**
         * @description Missing/null/wrong exact Origin or required session CSRF token. No protected effect. No detail, instance, code, correlation ID, payload echo or additional members on this bounded interface; match type, never parse title.
         * @example {
         *       "type": "urn:kencleng:problem:forbidden",
         *       "title": "Request forbidden",
         *       "status": 403
         *     }
         */
        ProblemForbidden: {
            /** @enum {string} */
            type: "urn:kencleng:problem:forbidden";
            /** @enum {string} */
            title: "Request forbidden";
            /** @enum {integer} */
            status: 403;
        };
        /**
         * @description Organization absent or not currently owned; identical safe response for both. No detail, instance, code, correlation ID, payload echo or additional members on this bounded interface; match type, never parse title.
         * @example {
         *       "type": "urn:kencleng:problem:resource-unavailable",
         *       "title": "Resource unavailable",
         *       "status": 404
         *     }
         */
        ProblemResourceUnavailable: {
            /** @enum {string} */
            type: "urn:kencleng:problem:resource-unavailable";
            /** @enum {string} */
            title: "Resource unavailable";
            /** @enum {integer} */
            status: 404;
        };
        /**
         * @description No visible unexpired complete creation receipt. Does not prove original failure or cancellation. No detail, instance, code, correlation ID, payload echo or additional members on this bounded interface; match type, never parse title.
         * @example {
         *       "type": "urn:kencleng:problem:result-not-observed",
         *       "title": "Result not observed",
         *       "status": 404
         *     }
         */
        ProblemResultNotObserved: {
            /** @enum {string} */
            type: "urn:kencleng:problem:result-not-observed";
            /** @enum {string} */
            title: "Result not observed";
            /** @enum {integer} */
            status: 404;
        };
        /**
         * @description Same account/key has unexpired receipt for different canonical name; no new effect. No detail, instance, code, correlation ID, payload echo or additional members on this bounded interface; match type, never parse title.
         * @example {
         *       "type": "urn:kencleng:problem:idempotency-mismatch",
         *       "title": "Request identity conflict",
         *       "status": 409
         *     }
         */
        ProblemIdempotencyMismatch: {
            /** @enum {string} */
            type: "urn:kencleng:problem:idempotency-mismatch";
            /** @enum {string} */
            title: "Request identity conflict";
            /** @enum {integer} */
            status: 409;
        };
        /**
         * @description Live session lacks purpose-specific same-account proof or proof age exceeds bound. Start explicit special flow; no intent. No detail, instance, code, correlation ID, payload echo or additional members on this bounded interface; match type, never parse title.
         * @example {
         *       "type": "urn:kencleng:problem:deletion-proof-required",
         *       "title": "Fresh account proof required",
         *       "status": 409
         *     }
         */
        ProblemDeletionProofRequired: {
            /** @enum {string} */
            type: "urn:kencleng:problem:deletion-proof-required";
            /** @enum {string} */
            title: "Fresh account proof required";
            /** @enum {integer} */
            status: 409;
        };
        /**
         * @description Current complete dependency set or simple eligibility differs from preview; invalidate preview, no intent. No detail, instance, code, correlation ID, payload echo or additional members on this bounded interface; match type, never parse title.
         * @example {
         *       "type": "urn:kencleng:problem:deletion-scope-changed",
         *       "title": "Deletion scope changed",
         *       "status": 409
         *     }
         */
        ProblemDeletionScopeChanged: {
            /** @enum {string} */
            type: "urn:kencleng:problem:deletion-scope-changed";
            /** @enum {string} */
            title: "Deletion scope changed";
            /** @enum {integer} */
            status: 409;
        };
        /**
         * @description Preview missing, expired, replaced, wrong account/session or already invalidated without confirmed intent; safe uniform body; no intent. No detail, instance, code, correlation ID, payload echo or additional members on this bounded interface; match type, never parse title.
         * @example {
         *       "type": "urn:kencleng:problem:deletion-preview-invalid",
         *       "title": "Deletion preview invalid",
         *       "status": 409
         *     }
         */
        ProblemDeletionPreviewInvalid: {
            /** @enum {string} */
            type: "urn:kencleng:problem:deletion-preview-invalid";
            /** @enum {string} */
            title: "Deletion preview invalid";
            /** @enum {integer} */
            status: 409;
        };
        /**
         * @description No replayable receipt and key timestamp+24h <= current time; no new effect, original may still be unknown. No detail, instance, code, correlation ID, payload echo or additional members on this bounded interface; match type, never parse title.
         * @example {
         *       "type": "urn:kencleng:problem:request-window-expired",
         *       "title": "Request window expired",
         *       "status": 410
         *     }
         */
        ProblemRequestWindowExpired: {
            /** @enum {string} */
            type: "urn:kencleng:problem:request-window-expired";
            /** @enum {string} */
            title: "Request window expired";
            /** @enum {integer} */
            status: 410;
        };
        /**
         * @description Canonical organization name empty, over200 code points or contains Cc controls. No detail, instance, code, correlation ID, payload echo or additional members on this bounded interface; match type, never parse title.
         * @example {
         *       "type": "urn:kencleng:problem:validation",
         *       "title": "Validation failed",
         *       "status": 422,
         *       "errors": [
         *         {
         *           "field": "name",
         *           "detail": "A name is required."
         *         }
         *       ]
         *     }
         */
        ProblemValidation: {
            /** @enum {string} */
            type: "urn:kencleng:problem:validation";
            /** @enum {string} */
            title: "Validation failed";
            /** @enum {integer} */
            status: 422;
            errors: {
                /** @enum {string} */
                field: "name";
                /** @enum {string} */
                detail: "A name is required." | "Name exceeds 200 Unicode code points." | "Name contains a control character.";
            }[];
        };
        /**
         * @description Simple closure cannot safely include complete dependency set within64KiB or supported dependency registry. Before accepted intent only; no actionable preview/erase. No detail, instance, code, correlation ID, payload echo or additional members on this bounded interface; match type, never parse title.
         * @example {
         *       "type": "urn:kencleng:problem:manual-review-required",
         *       "title": "Manual review required",
         *       "status": 422
         *     }
         */
        ProblemManualReviewRequired: {
            /** @enum {string} */
            type: "urn:kencleng:problem:manual-review-required";
            /** @enum {string} */
            title: "Manual review required";
            /** @enum {integer} */
            status: 422;
        };
        /**
         * @description Admission refused before this request creates a new effect; Retry-After seconds. Original concurrent request may remain unknown. No detail, instance, code, correlation ID, payload echo or additional members on this bounded interface; match type, never parse title.
         * @example {
         *       "type": "urn:kencleng:problem:rate-limited",
         *       "title": "Too many requests",
         *       "status": 429
         *     }
         */
        ProblemRateLimited: {
            /** @enum {string} */
            type: "urn:kencleng:problem:rate-limited";
            /** @enum {string} */
            title: "Too many requests";
            /** @enum {integer} */
            status: 429;
        };
        /**
         * @description Required DB, journal, provider/configuration or integrity evidence unavailable; no success claim. After possible consequential commit use result-unknown instead. No detail, instance, code, correlation ID, payload echo or additional members on this bounded interface; match type, never parse title.
         * @example {
         *       "type": "urn:kencleng:problem:unavailable",
         *       "title": "Service unavailable",
         *       "status": 503
         *     }
         */
        ProblemUnavailable: {
            /** @enum {string} */
            type: "urn:kencleng:problem:unavailable";
            /** @enum {string} */
            title: "Service unavailable";
            /** @enum {integer} */
            status: 503;
        };
        /**
         * @description Receipt/pair integrity cannot be established; never partial success. No detail, instance, code, correlation ID, payload echo or additional members on this bounded interface; match type, never parse title.
         * @example {
         *       "type": "urn:kencleng:problem:result-unavailable",
         *       "title": "Result unavailable",
         *       "status": 503
         *     }
         */
        ProblemResultUnavailable: {
            /** @enum {string} */
            type: "urn:kencleng:problem:result-unavailable";
            /** @enum {string} */
            title: "Result unavailable";
            /** @enum {integer} */
            status: 503;
        };
        /**
         * @description Bounded creation lock wait exceeded; this request rolled back; another request may commit. Keep same key. No detail, instance, code, correlation ID, payload echo or additional members on this bounded interface; match type, never parse title.
         * @example {
         *       "type": "urn:kencleng:problem:request-busy",
         *       "title": "Request busy",
         *       "status": 503
         *     }
         */
        ProblemRequestBusy: {
            /** @enum {string} */
            type: "urn:kencleng:problem:request-busy";
            /** @enum {string} */
            title: "Request busy";
            /** @enum {integer} */
            status: 503;
        };
        /**
         * @description Refresh issuance, create commit, journal append acknowledgment or deletion commit outcome unknown. Never infer rollback or automatically change key/login. No detail, instance, code, correlation ID, payload echo or additional members on this bounded interface; match type, never parse title.
         * @example {
         *       "type": "urn:kencleng:problem:result-unknown",
         *       "title": "Result unknown",
         *       "status": 503
         *     }
         */
        ProblemResultUnknown: {
            /** @enum {string} */
            type: "urn:kencleng:problem:result-unknown";
            /** @enum {string} */
            title: "Result unknown";
            /** @enum {integer} */
            status: 503;
        };
        /**
         * @description Durable confirmed intent known accepted/fenced; active erase not yet verified. No automatic cancellation, repeat login or public status. No detail, instance, code, correlation ID, payload echo or additional members on this bounded interface; match type, never parse title.
         * @example {
         *       "type": "urn:kencleng:problem:deletion-pending",
         *       "title": "Deletion pending",
         *       "status": 503
         *     }
         */
        ProblemDeletionPending: {
            /** @enum {string} */
            type: "urn:kencleng:problem:deletion-pending";
            /** @enum {string} */
            title: "Deletion pending";
            /** @enum {integer} */
            status: 503;
        };
    };
    responses: never;
    parameters: never;
    requestBodies: never;
    headers: never;
    pathItems: never;
}
export type $defs = Record<string, never>;
export interface operations {
    createGoogleAttempt: {
        parameters: {
            query?: never;
            header: {
                /** @description Exactly configured browser-facing origin including scheme/host/port; absent, null, multiple or mismatch ->403. No CORS grant. Not derived from untrusted forwarded headers. */
                Origin: string;
                /** @description Required only when purpose=deleteAccount. Normal {} ignores it; not a credential. */
                "X-CSRF-Token"?: string;
            };
            path?: never;
            cookie?: never;
        };
        /** @description Required application/json, UTF-8, closed object, no duplicate JSON keys. Other media types ->400 invalid-request; no simple form fallback. */
        requestBody: {
            content: {
                "application/json": components["schemas"]["GoogleAttempt"];
            };
        };
        responses: {
            /** @description Success */
            201: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    /** @description Multiple separate Set-Cookie header lines; never comma fold. Exact cookie tuples and expiry obligations in interface-decisions.md. */
                    "Set-Cookie": string;
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AttemptResponse"];
                };
            };
            /** @description Invalid request */
            400: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemInvalidRequest"];
                };
            };
            /** @description Authentication required */
            401: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemUnauthenticated"];
                };
            };
            /** @description Request forbidden */
            403: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemForbidden"];
                };
            };
            /** @description Too many requests */
            429: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    /** @description Positive integer seconds1..60; fixed admission bound, no new key instruction. */
                    "Retry-After": number;
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemRateLimited"];
                };
            };
            /** @description Service unavailable */
            503: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemUnavailable"];
                };
            };
        };
    };
    completeGoogleAttempt: {
        parameters: {
            query?: {
                /** @description Required on success, must be absent when error present. */
                code?: string;
                /** @description Required for both success and provider denial; exact browser-bound one-use match. */
                state?: string;
                /** @description Provider denial identifier, treated as untrusted; must be absent with code. */
                error?: string;
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Fixed same-origin browser redirect only, no body. All protocol outcomes use303. App cookies only after validated success, otherwise existing app cookies preserved. */
            303: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    /** @description Multiple separate Set-Cookie header lines; never comma fold. Exact cookie tuples and expiry obligations in interface-decisions.md. */
                    "Set-Cookie": string;
                    Location: "/auth/complete" | "/account-deletion/review" | "/auth/failure";
                    "Referrer-Policy": "no-referrer";
                    [name: string]: unknown;
                };
                content?: never;
            };
        };
    };
    getSession: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Success */
            200: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["SessionResponse"];
                };
            };
            /** @description Invalid request */
            400: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemInvalidRequest"];
                };
            };
            /** @description Authentication required */
            401: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemUnauthenticated"];
                };
            };
            /** @description Service unavailable */
            503: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemUnavailable"];
                };
            };
        };
    };
    refreshSession: {
        parameters: {
            query?: never;
            header: {
                /** @description Exactly configured browser-facing origin including scheme/host/port; absent, null, multiple or mismatch ->403. No CORS grant. Not derived from untrusted forwarded headers. */
                Origin: string;
                /** @description Constant-time checked against effective session derivation; wrong/missing ->403 after effective authentication. Stable across refresh; reload via GET session. Never authorization alone. */
                "X-CSRF-Token": string;
            };
            path?: never;
            cookie?: never;
        };
        /** @description Required application/json, UTF-8, closed object, no duplicate JSON keys. Other media types ->400 invalid-request; no simple form fallback. */
        requestBody: {
            content: {
                "application/json": components["schemas"]["Empty"];
            };
        };
        responses: {
            /** @description No body; confirmed atomic rotation only. */
            204: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    /** @description Multiple separate Set-Cookie header lines; never comma fold. Exact cookie tuples and expiry obligations in interface-decisions.md. */
                    "Set-Cookie": string;
                    [name: string]: unknown;
                };
                content?: never;
            };
            /** @description Invalid request */
            400: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemInvalidRequest"];
                };
            };
            /** @description Authentication required */
            401: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemUnauthenticated"];
                };
            };
            /** @description Request forbidden */
            403: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemForbidden"];
                };
            };
            /** @description Result unknown | Service unavailable */
            503: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemResultUnknown"] | components["schemas"]["ProblemUnavailable"];
                };
            };
        };
    };
    logoutSession: {
        parameters: {
            query?: never;
            header: {
                /** @description Exactly configured browser-facing origin including scheme/host/port; absent, null, multiple or mismatch ->403. No CORS grant. Not derived from untrusted forwarded headers. */
                Origin: string;
                /** @description Required iff an effective session remains. Missing/wrong live-session token ->403 no revoke. Harmless already-ended repeat branch needs Origin+JSON only. */
                "X-CSRF-Token"?: string;
            };
            path?: never;
            cookie?: never;
        };
        /** @description Required application/json, UTF-8, closed object, no duplicate JSON keys. Other media types ->400 invalid-request; no simple form fallback. */
        requestBody: {
            content: {
                "application/json": components["schemas"]["Empty"];
            };
        };
        responses: {
            /** @description No body; current resolved session authority ended or no effective session exists. Does not establish when or how an earlier request ended it. */
            204: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    /** @description Multiple separate Set-Cookie header lines; never comma fold. Exact cookie tuples and expiry obligations in interface-decisions.md. */
                    "Set-Cookie": string;
                    [name: string]: unknown;
                };
                content?: never;
            };
            /** @description Invalid request */
            400: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemInvalidRequest"];
                };
            };
            /** @description Authentication required */
            401: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemUnauthenticated"];
                };
            };
            /** @description Request forbidden */
            403: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemForbidden"];
                };
            };
            /** @description Service unavailable | Result unknown */
            503: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemUnavailable"] | components["schemas"]["ProblemResultUnknown"];
                };
            };
        };
    };
    listOrganizations: {
        parameters: {
            query?: {
                limit?: number;
                /** @description Server-produced unpadded base64url UTF-8 JSON tuple [UTC createdAt,UUID id], position only. Decode/type validation required; no client authorization authority. Unknown members/invalid encoding ->400. */
                cursor?: string;
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Success */
            200: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["OrganizationList"];
                };
            };
            /** @description Invalid request */
            400: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemInvalidRequest"];
                };
            };
            /** @description Authentication required */
            401: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemUnauthenticated"];
                };
            };
            /** @description Service unavailable */
            503: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemUnavailable"];
                };
            };
        };
    };
    createOrganization: {
        parameters: {
            query?: never;
            header: {
                /** @description Exactly configured browser-facing origin including scheme/host/port; absent, null, multiple or mismatch ->403. No CORS grant. Not derived from untrusted forwarded headers. */
                Origin: string;
                /** @description Constant-time checked against effective session derivation; wrong/missing ->403 after effective authentication. Stable across refresh; reload via GET session. Never authorization alone. */
                "X-CSRF-Token": string;
                /** @description UUIDv7 required, canonical case-insensitive UUID identity. >5m future ->400; no duplicate header. */
                "Idempotency-Key": string;
            };
            path?: never;
            cookie?: never;
        };
        /** @description Required application/json, UTF-8, closed object, no duplicate JSON keys. Other media types ->400 invalid-request; no simple form fallback. */
        requestBody: {
            content: {
                "application/json": components["schemas"]["CreateOrganization"];
            };
        };
        responses: {
            /** @description Exact replay of unexpired receipt, current complete authorized pair. */
            200: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["CreationResponse"];
                };
            };
            /** @description Success */
            201: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    /** @description Backend resource reference /organizations/{id}, not frontend navigation and not a directly fetchable browser URL. Browser transport uses returned organization.id with explicit /api/organizations/{id}; never automatically follow arbitrary Location. No ingress Location rewrite required. */
                    Location: string;
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["CreationResponse"];
                };
            };
            /** @description Invalid request */
            400: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemInvalidRequest"];
                };
            };
            /** @description Authentication required */
            401: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemUnauthenticated"];
                };
            };
            /** @description Request forbidden */
            403: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemForbidden"];
                };
            };
            /** @description Request identity conflict */
            409: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemIdempotencyMismatch"];
                };
            };
            /** @description Request window expired */
            410: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemRequestWindowExpired"];
                };
            };
            /** @description Validation failed */
            422: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemValidation"];
                };
            };
            /** @description Too many requests */
            429: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    /** @description Positive integer seconds1..60; fixed admission bound, no new key instruction. */
                    "Retry-After": number;
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemRateLimited"];
                };
            };
            /** @description Service unavailable | Result unknown | Request busy */
            503: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemUnavailable"] | components["schemas"]["ProblemResultUnknown"] | components["schemas"]["ProblemRequestBusy"];
                };
            };
        };
    };
    getOrganizationCreationResult: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                requestId: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Success */
            200: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["CreationResponse"];
                };
            };
            /** @description Invalid request */
            400: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemInvalidRequest"];
                };
            };
            /** @description Authentication required */
            401: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemUnauthenticated"];
                };
            };
            /** @description Result not observed */
            404: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemResultNotObserved"];
                };
            };
            /** @description Result unavailable | Service unavailable */
            503: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemResultUnavailable"] | components["schemas"]["ProblemUnavailable"];
                };
            };
        };
    };
    getOrganization: {
        parameters: {
            query?: never;
            header?: never;
            path: {
                organizationId: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Success */
            200: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["OrganizationResponse"];
                };
            };
            /** @description Invalid request */
            400: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemInvalidRequest"];
                };
            };
            /** @description Authentication required */
            401: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemUnauthenticated"];
                };
            };
            /** @description Resource unavailable */
            404: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemResourceUnavailable"];
                };
            };
            /** @description Result unavailable | Service unavailable */
            503: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemResultUnavailable"] | components["schemas"]["ProblemUnavailable"];
                };
            };
        };
    };
    createAccountDeletionPreview: {
        parameters: {
            query?: never;
            header: {
                /** @description Exactly configured browser-facing origin including scheme/host/port; absent, null, multiple or mismatch ->403. No CORS grant. Not derived from untrusted forwarded headers. */
                Origin: string;
                /** @description Constant-time checked against effective session derivation; wrong/missing ->403 after effective authentication. Stable across refresh; reload via GET session. Never authorization alone. */
                "X-CSRF-Token": string;
            };
            path?: never;
            cookie?: never;
        };
        /** @description Required application/json, UTF-8, closed object, no duplicate JSON keys. Other media types ->400 invalid-request; no simple form fallback. */
        requestBody: {
            content: {
                "application/json": components["schemas"]["Empty"];
            };
        };
        responses: {
            /** @description Success */
            201: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["DeletionPreviewResponse"];
                };
            };
            /** @description Invalid request */
            400: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemInvalidRequest"];
                };
            };
            /** @description Authentication required */
            401: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemUnauthenticated"];
                };
            };
            /** @description Request forbidden */
            403: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemForbidden"];
                };
            };
            /** @description Fresh account proof required */
            409: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemDeletionProofRequired"];
                };
            };
            /** @description Manual review required */
            422: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemManualReviewRequired"];
                };
            };
            /** @description Too many requests */
            429: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    /** @description Positive integer seconds1..60; fixed admission bound, no new key instruction. */
                    "Retry-After": number;
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemRateLimited"];
                };
            };
            /** @description Service unavailable */
            503: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemUnavailable"];
                };
            };
        };
    };
    deleteAccount: {
        parameters: {
            query?: never;
            header: {
                /** @description Exactly configured browser-facing origin including scheme/host/port; absent, null, multiple or mismatch ->403. No CORS grant. Not derived from untrusted forwarded headers. */
                Origin: string;
                /** @description Constant-time checked against effective session derivation; wrong/missing ->403 after effective authentication. Stable across refresh; reload via GET session. Never authorization alone. */
                "X-CSRF-Token": string;
            };
            path?: never;
            cookie?: never;
        };
        /** @description Required application/json, UTF-8, closed object, no duplicate JSON keys. Other media types ->400 invalid-request; no simple form fallback. */
        requestBody: {
            content: {
                "application/json": components["schemas"]["DeleteAccount"];
            };
        };
        responses: {
            /** @description Success */
            200: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    /** @description Multiple separate Set-Cookie header lines; never comma fold. Exact cookie tuples and expiry obligations in interface-decisions.md. */
                    "Set-Cookie": string;
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["DeletionResponse"];
                };
            };
            /** @description Invalid request */
            400: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemInvalidRequest"];
                };
            };
            /** @description Authentication required */
            401: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemUnauthenticated"];
                };
            };
            /** @description Request forbidden */
            403: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemForbidden"];
                };
            };
            /** @description Fresh account proof required | Deletion scope changed | Deletion preview invalid */
            409: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemDeletionProofRequired"] | components["schemas"]["ProblemDeletionScopeChanged"] | components["schemas"]["ProblemDeletionPreviewInvalid"];
                };
            };
            /** @description Manual review required */
            422: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemManualReviewRequired"];
                };
            };
            /** @description Service unavailable | Result unknown | Deletion pending */
            503: {
                headers: {
                    /** @description Required on every response, including errors/redirects. */
                    "Cache-Control": "no-store";
                    [name: string]: unknown;
                };
                content: {
                    "application/problem+json": components["schemas"]["ProblemUnavailable"] | components["schemas"]["ProblemResultUnknown"] | components["schemas"]["ProblemDeletionPending"];
                };
            };
        };
    };
}
