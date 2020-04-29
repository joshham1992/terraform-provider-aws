// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeBadRequestException for service response error code
	// "WAFBadRequestException".
	ErrCodeBadRequestException = "WAFBadRequestException"

	// ErrCodeDisallowedNameException for service response error code
	// "WAFDisallowedNameException".
	//
	// The name specified is invalid.
	ErrCodeDisallowedNameException = "WAFDisallowedNameException"

	// ErrCodeInternalErrorException for service response error code
	// "WAFInternalErrorException".
	//
	// The operation failed because of a system problem, even though the request
	// was valid. Retry your request.
	ErrCodeInternalErrorException = "WAFInternalErrorException"

	// ErrCodeInvalidAccountException for service response error code
	// "WAFInvalidAccountException".
	//
	// The operation failed because you tried to create, update, or delete an object
	// by using an invalid account identifier.
	ErrCodeInvalidAccountException = "WAFInvalidAccountException"

	// ErrCodeInvalidOperationException for service response error code
	// "WAFInvalidOperationException".
	//
	// The operation failed because there was nothing to do. For example:
	//
	//    * You tried to remove a Rule from a WebACL, but the Rule isn't in the
	//    specified WebACL.
	//
	//    * You tried to remove an IP address from an IPSet, but the IP address
	//    isn't in the specified IPSet.
	//
	//    * You tried to remove a ByteMatchTuple from a ByteMatchSet, but the ByteMatchTuple
	//    isn't in the specified WebACL.
	//
	//    * You tried to add a Rule to a WebACL, but the Rule already exists in
	//    the specified WebACL.
	//
	//    * You tried to add a ByteMatchTuple to a ByteMatchSet, but the ByteMatchTuple
	//    already exists in the specified WebACL.
	ErrCodeInvalidOperationException = "WAFInvalidOperationException"

	// ErrCodeInvalidParameterException for service response error code
	// "WAFInvalidParameterException".
	//
	// The operation failed because AWS WAF didn't recognize a parameter in the
	// request. For example:
	//
	//    * You specified an invalid parameter name.
	//
	//    * You specified an invalid value.
	//
	//    * You tried to update an object (ByteMatchSet, IPSet, Rule, or WebACL)
	//    using an action other than INSERT or DELETE.
	//
	//    * You tried to create a WebACL with a DefaultAction Type other than ALLOW,
	//    BLOCK, or COUNT.
	//
	//    * You tried to create a RateBasedRule with a RateKey value other than
	//    IP.
	//
	//    * You tried to update a WebACL with a WafAction Type other than ALLOW,
	//    BLOCK, or COUNT.
	//
	//    * You tried to update a ByteMatchSet with a FieldToMatch Type other than
	//    HEADER, METHOD, QUERY_STRING, URI, or BODY.
	//
	//    * You tried to update a ByteMatchSet with a Field of HEADER but no value
	//    for Data.
	//
	//    * Your request references an ARN that is malformed, or corresponds to
	//    a resource with which a web ACL cannot be associated.
	ErrCodeInvalidParameterException = "WAFInvalidParameterException"

	// ErrCodeInvalidPermissionPolicyException for service response error code
	// "WAFInvalidPermissionPolicyException".
	//
	// The operation failed because the specified policy is not in the proper format.
	//
	// The policy is subject to the following restrictions:
	//
	//    * You can attach only one policy with each PutPermissionPolicy request.
	//
	//    * The policy must include an Effect, Action and Principal.
	//
	//    * Effect must specify Allow.
	//
	//    * The Action in the policy must be waf:UpdateWebACL, waf-regional:UpdateWebACL,
	//    waf:GetRuleGroup and waf-regional:GetRuleGroup . Any extra or wildcard
	//    actions in the policy will be rejected.
	//
	//    * The policy cannot include a Resource parameter.
	//
	//    * The ARN in the request must be a valid WAF RuleGroup ARN and the RuleGroup
	//    must exist in the same region.
	//
	//    * The user making the request must be the owner of the RuleGroup.
	//
	//    * Your policy must be composed using IAM Policy version 2012-10-17.
	ErrCodeInvalidPermissionPolicyException = "WAFInvalidPermissionPolicyException"

	// ErrCodeInvalidRegexPatternException for service response error code
	// "WAFInvalidRegexPatternException".
	//
	// The regular expression (regex) you specified in RegexPatternString is invalid.
	ErrCodeInvalidRegexPatternException = "WAFInvalidRegexPatternException"

	// ErrCodeLimitsExceededException for service response error code
	// "WAFLimitsExceededException".
	//
	// The operation exceeds a resource limit, for example, the maximum number of
	// WebACL objects that you can create for an AWS account. For more information,
	// see Limits (https://docs.aws.amazon.com/waf/latest/developerguide/limits.html)
	// in the AWS WAF Developer Guide.
	ErrCodeLimitsExceededException = "WAFLimitsExceededException"

	// ErrCodeNonEmptyEntityException for service response error code
	// "WAFNonEmptyEntityException".
	//
	// The operation failed because you tried to delete an object that isn't empty.
	// For example:
	//
	//    * You tried to delete a WebACL that still contains one or more Rule objects.
	//
	//    * You tried to delete a Rule that still contains one or more ByteMatchSet
	//    objects or other predicates.
	//
	//    * You tried to delete a ByteMatchSet that contains one or more ByteMatchTuple
	//    objects.
	//
	//    * You tried to delete an IPSet that references one or more IP addresses.
	ErrCodeNonEmptyEntityException = "WAFNonEmptyEntityException"

	// ErrCodeNonexistentContainerException for service response error code
	// "WAFNonexistentContainerException".
	//
	// The operation failed because you tried to add an object to or delete an object
	// from another object that doesn't exist. For example:
	//
	//    * You tried to add a Rule to or delete a Rule from a WebACL that doesn't
	//    exist.
	//
	//    * You tried to add a ByteMatchSet to or delete a ByteMatchSet from a Rule
	//    that doesn't exist.
	//
	//    * You tried to add an IP address to or delete an IP address from an IPSet
	//    that doesn't exist.
	//
	//    * You tried to add a ByteMatchTuple to or delete a ByteMatchTuple from
	//    a ByteMatchSet that doesn't exist.
	ErrCodeNonexistentContainerException = "WAFNonexistentContainerException"

	// ErrCodeNonexistentItemException for service response error code
	// "WAFNonexistentItemException".
	//
	// The operation failed because the referenced object doesn't exist.
	ErrCodeNonexistentItemException = "WAFNonexistentItemException"

	// ErrCodeReferencedItemException for service response error code
	// "WAFReferencedItemException".
	//
	// The operation failed because you tried to delete an object that is still
	// in use. For example:
	//
	//    * You tried to delete a ByteMatchSet that is still referenced by a Rule.
	//
	//    * You tried to delete a Rule that is still referenced by a WebACL.
	ErrCodeReferencedItemException = "WAFReferencedItemException"

	// ErrCodeServiceLinkedRoleErrorException for service response error code
	// "WAFServiceLinkedRoleErrorException".
	//
	// AWS WAF is not able to access the service linked role. This can be caused
	// by a previous PutLoggingConfiguration request, which can lock the service
	// linked role for about 20 seconds. Please try your request again. The service
	// linked role can also be locked by a previous DeleteServiceLinkedRole request,
	// which can lock the role for 15 minutes or more. If you recently made a DeleteServiceLinkedRole,
	// wait at least 15 minutes and try the request again. If you receive this same
	// exception again, you will have to wait additional time until the role is
	// unlocked.
	ErrCodeServiceLinkedRoleErrorException = "WAFServiceLinkedRoleErrorException"

	// ErrCodeStaleDataException for service response error code
	// "WAFStaleDataException".
	//
	// The operation failed because you tried to create, update, or delete an object
	// by using a change token that has already been used.
	ErrCodeStaleDataException = "WAFStaleDataException"

	// ErrCodeSubscriptionNotFoundException for service response error code
	// "WAFSubscriptionNotFoundException".
	//
	// The specified subscription does not exist.
	ErrCodeSubscriptionNotFoundException = "WAFSubscriptionNotFoundException"

	// ErrCodeTagOperationException for service response error code
	// "WAFTagOperationException".
	ErrCodeTagOperationException = "WAFTagOperationException"

	// ErrCodeTagOperationInternalErrorException for service response error code
	// "WAFTagOperationInternalErrorException".
	ErrCodeTagOperationInternalErrorException = "WAFTagOperationInternalErrorException"

	// ErrCodeWAFEntityMigrationException for service response error code
	// "WAFEntityMigrationException".
	//
	// The operation failed due to a problem with the migration. The failure cause
	// is provided in the exception, in the MigrationErrorType:
	//
	//    * ENTITY_NOT_SUPPORTED - The web ACL has an unsupported entity but the
	//    IgnoreUnsupportedType is not set to true.
	//
	//    * ENTITY_NOT_FOUND - The web ACL doesn't exist.
	//
	//    * S3_BUCKET_NO_PERMISSION - You don't have permission to perform the PutObject
	//    action to the specified Amazon S3 bucket.
	//
	//    * S3_BUCKET_NOT_ACCESSIBLE - The bucket policy doesn't allow AWS WAF to
	//    perform the PutObject action in the bucket.
	//
	//    * S3_BUCKET_NOT_FOUND - The S3 bucket doesn't exist.
	//
	//    * S3_BUCKET_INVALID_REGION - The S3 bucket is not in the same Region as
	//    the web ACL.
	//
	//    * S3_INTERNAL_ERROR - AWS WAF failed to create the template in the S3
	//    bucket for another reason.
	ErrCodeWAFEntityMigrationException = "WAFEntityMigrationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"WAFBadRequestException":                newErrorBadRequestException,
	"WAFDisallowedNameException":            newErrorDisallowedNameException,
	"WAFInternalErrorException":             newErrorInternalErrorException,
	"WAFInvalidAccountException":            newErrorInvalidAccountException,
	"WAFInvalidOperationException":          newErrorInvalidOperationException,
	"WAFInvalidParameterException":          newErrorInvalidParameterException,
	"WAFInvalidPermissionPolicyException":   newErrorInvalidPermissionPolicyException,
	"WAFInvalidRegexPatternException":       newErrorInvalidRegexPatternException,
	"WAFLimitsExceededException":            newErrorLimitsExceededException,
	"WAFNonEmptyEntityException":            newErrorNonEmptyEntityException,
	"WAFNonexistentContainerException":      newErrorNonexistentContainerException,
	"WAFNonexistentItemException":           newErrorNonexistentItemException,
	"WAFReferencedItemException":            newErrorReferencedItemException,
	"WAFServiceLinkedRoleErrorException":    newErrorServiceLinkedRoleErrorException,
	"WAFStaleDataException":                 newErrorStaleDataException,
	"WAFSubscriptionNotFoundException":      newErrorSubscriptionNotFoundException,
	"WAFTagOperationException":              newErrorTagOperationException,
	"WAFTagOperationInternalErrorException": newErrorTagOperationInternalErrorException,
	"WAFEntityMigrationException":           newErrorWAFEntityMigrationException,
}
