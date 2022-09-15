// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tokens

import (
	"context"
)

// The Token API allows you to create, list, and revoke tokens that can be used
// to authenticate and access Databricks REST APIs.
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type TokensService interface {

	// Creates and returns a token for a user. If this call is made through
	// token authentication, it will create the token that has the same client
	// id with the authenticated token. This call returns an error
	// ``QUOTA_EXCEEDED`` if over the token quota for the user.
	Create(ctx context.Context, request CreateTokenRequest) (*CreateTokenResponse, error)

	// Revokes an access token. This call returns an error
	// ``RESOURCE_DOES_NOT_EXIST`` if a token with the given ID is not valid.
	Delete(ctx context.Context, request RevokeTokenRequest) error

	// DeleteByTokenId calls Delete, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteByTokenId(ctx context.Context, tokenId string) error

	// Lists all the valid tokens for a user-workspace pair.
	List(ctx context.Context) (*ListTokensResponse, error)
}