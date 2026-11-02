// Generated from accepted r4. Run npm run generate:org; do not edit.
import { z } from "zod";
export const Empty = z.strictObject({});
export const GoogleAttempt = z.strictObject({"purpose":z.literal("deleteAccount").optional()});
export const Organization = z.strictObject({"id":z.string().uuid(),"name":z.string().refine(value => [...value].length >= 1).refine(value => [...value].length <= 200),"createdAt":z.string().datetime({offset:false}).regex(new RegExp("Z$"))});
export const Owner = z.strictObject({"accountId":z.string().uuid(),"role":z.literal("Owner"),"scope":z.literal("internal")});
export const OrganizationPair = z.strictObject({"organization":Organization,"owner":Owner});
export const OrganizationResponse = z.strictObject({"data":OrganizationPair});
export const CreationResponse = z.strictObject({"data":z.strictObject({"requestId":z.string().uuid().regex(new RegExp("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-7[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")),"organization":Organization,"owner":Owner})});
export const OrganizationList = z.strictObject({"data":z.array(OrganizationPair).max(50),"meta":z.strictObject({"nextCursor":z.string().refine(value => [...value].length >= 1).refine(value => [...value].length <= 512)}).optional()});
export const SessionResponse = z.strictObject({"data":z.strictObject({"accountId":z.string().uuid(),"csrfToken":z.string().refine(value => [...value].length >= 43).refine(value => [...value].length <= 43).regex(new RegExp("^[A-Za-z0-9_-]{43}$"))})});
export const AttemptResponse = z.strictObject({"data":z.strictObject({"authorizationUrl":z.string().refine(value => [...value].length <= 8192)})});
export const CreateOrganization = z.strictObject({"name":z.string()});
export const DeletionPreviewResponse = z.strictObject({"data":z.strictObject({"previewId":z.string().uuid(),"expiresAt":z.string().datetime({offset:false}).regex(new RegExp("Z$")),"organizations":z.array(z.strictObject({"id":z.string().uuid(),"name":z.string().refine(value => [...value].length >= 1).refine(value => [...value].length <= 200)})),"effect":z.literal("closeContextsAndDeleteAccount"),"backupCopies":z.literal("pendingExpiry")})});
export const DeleteAccount = z.strictObject({"previewId":z.string().uuid(),"closeOrganizations":z.literal(true)});
export const DeletionResponse = z.strictObject({"data":z.strictObject({"accessEnded":z.literal(true),"activeDataErased":z.literal(true),"remainingCopies":z.enum(["pendingVerification","noneVerified"])})});
export const ProblemInvalidRequest = z.strictObject({"type":z.literal("urn:uuid:9ed32605-5d10-4f59-b811-042f45fdffed"),"title":z.literal("Invalid request"),"status":z.literal(400)});
export const ProblemUnauthenticated = z.strictObject({"type":z.literal("urn:uuid:0c56be0a-40a6-45c9-bb49-c20b075ba28b"),"title":z.literal("Authentication required"),"status":z.literal(401)});
export const ProblemForbidden = z.strictObject({"type":z.literal("urn:uuid:142cb767-7ce1-403f-88b4-b1fe7c6c2f82"),"title":z.literal("Request forbidden"),"status":z.literal(403)});
export const ProblemResourceUnavailable = z.strictObject({"type":z.literal("urn:uuid:5ed0b852-a063-4c56-85a7-a0d7c1775000"),"title":z.literal("Resource unavailable"),"status":z.literal(404)});
export const ProblemResultNotObserved = z.strictObject({"type":z.literal("urn:uuid:cbf15221-1fdb-47ca-8b8a-8285b9e031b1"),"title":z.literal("Result not observed"),"status":z.literal(404)});
export const ProblemIdempotencyMismatch = z.strictObject({"type":z.literal("urn:uuid:144ea3b4-83cc-4b0f-961a-892f3fc09b86"),"title":z.literal("Request identity conflict"),"status":z.literal(409)});
export const ProblemDeletionProofRequired = z.strictObject({"type":z.literal("urn:uuid:0ddfd153-87b2-4619-a11e-f48418401c21"),"title":z.literal("Fresh account proof required"),"status":z.literal(409)});
export const ProblemDeletionScopeChanged = z.strictObject({"type":z.literal("urn:uuid:bd8b4ade-aaef-4ea9-9d73-2585c7b31ea6"),"title":z.literal("Deletion scope changed"),"status":z.literal(409)});
export const ProblemDeletionPreviewInvalid = z.strictObject({"type":z.literal("urn:uuid:25c9699d-d474-496b-8214-65ceb830f197"),"title":z.literal("Deletion preview invalid"),"status":z.literal(409)});
export const ProblemRequestWindowExpired = z.strictObject({"type":z.literal("urn:uuid:22166aa9-86c0-4033-a5c6-49f7cb83f9ff"),"title":z.literal("Request window expired"),"status":z.literal(410)});
export const ProblemValidation = z.strictObject({"type":z.literal("urn:uuid:1b1d2eb4-2c68-4809-8910-f40bd94e1e42"),"title":z.literal("Validation failed"),"status":z.literal(422),"errors":z.array(z.strictObject({"field":z.literal("name"),"detail":z.enum(["A name is required.","Name exceeds 200 Unicode code points.","Name contains a control character."])})).min(1).max(1)});
export const ProblemManualReviewRequired = z.strictObject({"type":z.literal("urn:uuid:5a799183-67f6-49a9-90cd-ae9e06af1cf2"),"title":z.literal("Manual review required"),"status":z.literal(422)});
export const ProblemRateLimited = z.strictObject({"type":z.literal("urn:uuid:364a031e-17d4-42d0-af70-06c79716e98e"),"title":z.literal("Too many requests"),"status":z.literal(429)});
export const ProblemUnavailable = z.strictObject({"type":z.literal("urn:uuid:bc32fcd9-a6ce-4c42-bfa2-d09431017922"),"title":z.literal("Service unavailable"),"status":z.literal(503)});
export const ProblemResultUnavailable = z.strictObject({"type":z.literal("urn:uuid:59625f7b-1a32-46af-aae7-c42ddf7be1fd"),"title":z.literal("Result unavailable"),"status":z.literal(503)});
export const ProblemRequestBusy = z.strictObject({"type":z.literal("urn:uuid:cc31e0af-5ded-4ccd-afe6-cb0c592323cc"),"title":z.literal("Request busy"),"status":z.literal(503)});
export const ProblemResultUnknown = z.strictObject({"type":z.literal("urn:uuid:f67a74ed-f3a2-45bb-a2ee-8b14e9c4c96c"),"title":z.literal("Result unknown"),"status":z.literal(503)});
export const ProblemDeletionPending = z.strictObject({"type":z.literal("urn:uuid:6d143e2a-e3ba-4e71-8c67-1c8f7d4bae85"),"title":z.literal("Deletion pending"),"status":z.literal(503)});
export const operations = {
createGoogleAttempt:{path:"/auth/google/attempts",method:"POST",responses:{201:AttemptResponse,400:ProblemInvalidRequest,401:ProblemUnauthenticated,403:ProblemForbidden,429:ProblemRateLimited,503:ProblemUnavailable,}},
completeGoogleAttempt:{path:"/auth/google/callback",method:"GET",responses:{303:z.undefined(),}},
getSession:{path:"/auth/session",method:"GET",responses:{200:SessionResponse,400:ProblemInvalidRequest,401:ProblemUnauthenticated,503:ProblemUnavailable,}},
refreshSession:{path:"/auth/refresh",method:"POST",responses:{204:z.undefined(),400:ProblemInvalidRequest,401:ProblemUnauthenticated,403:ProblemForbidden,503:z.union([ProblemResultUnknown,ProblemUnavailable]),}},
logoutSession:{path:"/auth/logout",method:"POST",responses:{204:z.undefined(),400:ProblemInvalidRequest,401:ProblemUnauthenticated,403:ProblemForbidden,503:z.union([ProblemUnavailable,ProblemResultUnknown]),}},
createOrganization:{path:"/organizations",method:"POST",responses:{200:CreationResponse,201:CreationResponse,400:ProblemInvalidRequest,401:ProblemUnauthenticated,403:ProblemForbidden,409:ProblemIdempotencyMismatch,410:ProblemRequestWindowExpired,422:ProblemValidation,429:ProblemRateLimited,503:z.union([ProblemUnavailable,ProblemResultUnknown,ProblemRequestBusy]),}},
listOrganizations:{path:"/organizations",method:"GET",responses:{200:OrganizationList,400:ProblemInvalidRequest,401:ProblemUnauthenticated,503:ProblemUnavailable,}},
getOrganizationCreationResult:{path:"/organization-creation-requests/{requestId}",method:"GET",responses:{200:CreationResponse,400:ProblemInvalidRequest,401:ProblemUnauthenticated,404:ProblemResultNotObserved,503:z.union([ProblemResultUnavailable,ProblemUnavailable]),}},
getOrganization:{path:"/organizations/{organizationId}",method:"GET",responses:{200:OrganizationResponse,400:ProblemInvalidRequest,401:ProblemUnauthenticated,404:ProblemResourceUnavailable,503:z.union([ProblemResultUnavailable,ProblemUnavailable]),}},
createAccountDeletionPreview:{path:"/account-deletion-previews",method:"POST",responses:{201:DeletionPreviewResponse,400:ProblemInvalidRequest,401:ProblemUnauthenticated,403:ProblemForbidden,409:ProblemDeletionProofRequired,422:ProblemManualReviewRequired,429:ProblemRateLimited,503:ProblemUnavailable,}},
deleteAccount:{path:"/account-deletions",method:"POST",responses:{200:DeletionResponse,400:ProblemInvalidRequest,401:ProblemUnauthenticated,403:ProblemForbidden,409:z.union([ProblemDeletionProofRequired,ProblemDeletionScopeChanged,ProblemDeletionPreviewInvalid]),422:ProblemManualReviewRequired,503:z.union([ProblemUnavailable,ProblemResultUnknown,ProblemDeletionPending]),}},
} as const;
