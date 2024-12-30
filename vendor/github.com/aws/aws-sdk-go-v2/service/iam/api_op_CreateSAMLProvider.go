// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an IAM resource that describes an identity provider (IdP) that supports
// SAML 2.0.
//
// The SAML provider resource that you create with this operation can be used as a
// principal in an IAM role's trust policy. Such a policy can enable federated
// users who sign in using the SAML IdP to assume the role. You can create an IAM
// role that supports Web-based single sign-on (SSO) to the Amazon Web Services
// Management Console or one that supports API access to Amazon Web Services.
//
// When you create the SAML provider resource, you upload a SAML metadata document
// that you get from your IdP. That document includes the issuer's name, expiration
// information, and keys that can be used to validate the SAML authentication
// response (assertions) that the IdP sends. You must generate the metadata
// document using the identity management software that is used as your
// organization's IdP.
//
// This operation requires [Signature Version 4].
//
// For more information, see [Enabling SAML 2.0 federated users to access the Amazon Web Services Management Console] and [About SAML 2.0-based federation] in the IAM User Guide.
//
// [Signature Version 4]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [About SAML 2.0-based federation]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_saml.html
// [Enabling SAML 2.0 federated users to access the Amazon Web Services Management Console]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_enable-console-saml.html
func (c *Client) CreateSAMLProvider(ctx context.Context, params *CreateSAMLProviderInput, optFns ...func(*Options)) (*CreateSAMLProviderOutput, error) {
	if params == nil {
		params = &CreateSAMLProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSAMLProvider", params, optFns, c.addOperationCreateSAMLProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSAMLProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSAMLProviderInput struct {

	// The name of the provider to create.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	Name *string

	// An XML document generated by an identity provider (IdP) that supports SAML 2.0.
	// The document includes the issuer's name, expiration information, and keys that
	// can be used to validate the SAML authentication response (assertions) that are
	// received from the IdP. You must generate the metadata document using the
	// identity management software that is used as your organization's IdP.
	//
	// For more information, see [About SAML 2.0-based federation] in the IAM User Guide
	//
	// [About SAML 2.0-based federation]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_saml.html
	//
	// This member is required.
	SAMLMetadataDocument *string

	// A list of tags that you want to attach to the new IAM SAML provider. Each tag
	// consists of a key name and an associated value. For more information about
	// tagging, see [Tagging IAM resources]in the IAM User Guide.
	//
	// If any one of the tags is invalid or if you exceed the allowed maximum number
	// of tags, then the entire request fails and the resource is not created.
	//
	// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Contains the response to a successful CreateSAMLProvider request.
type CreateSAMLProviderOutput struct {

	// The Amazon Resource Name (ARN) of the new SAML provider resource in IAM.
	SAMLProviderArn *string

	// A list of tags that are attached to the new IAM SAML provider. The returned
	// list of tags is sorted by tag key. For more information about tagging, see [Tagging IAM resources]in
	// the IAM User Guide.
	//
	// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSAMLProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateSAMLProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateSAMLProvider{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSAMLProvider"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpCreateSAMLProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSAMLProvider(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateSAMLProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSAMLProvider",
	}
}
