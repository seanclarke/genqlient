package test

// Code generated by github.com/Khan/genql, DO NOT EDIT.

import (
	"github.com/Khan/genql/graphql"
)

type QueryWithStructsResponse struct {
	User QueryWithStructsUser `json:"user"`
}

type QueryWithStructsUser struct {
	AuthMethods []QueryWithStructsUserAuthMethodsAuthMethod `json:"authMethods"`
}

type QueryWithStructsUserAuthMethodsAuthMethod struct {
	Provider string `json:"provider"`
	Email    string `json:"email"`
}

func QueryWithStructs(client *graphql.Client) (*QueryWithStructsResponse, error) {
	var retval QueryWithStructsResponse
	err := client.MakeRequest(nil, `
query QueryWithStructs {
	user {
		authMethods {
			provider
			email
		}
	}
}
`, &retval, nil)
	return &retval, err
}
