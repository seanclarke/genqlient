package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"github.com/Khan/genqlient/graphql"
)

// QueryWithSlicesResponse is returned by QueryWithSlices on success.
type QueryWithSlicesResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User QueryWithSlicesUser `json:"user"`
}

// A User is a user!
type QueryWithSlicesUser struct {
	Emails                []string `json:"emails"`
	EmailsOrNull          []string `json:"emailsOrNull"`
	EmailsWithNulls       []string `json:"emailsWithNulls"`
	EmailsWithNullsOrNull []string `json:"emailsWithNullsOrNull"`
}

func QueryWithSlices(
	client graphql.Client,
) (*QueryWithSlicesResponse, error) {
	var retval QueryWithSlicesResponse
	err := client.MakeRequest(
		nil,
		"QueryWithSlices",
		`
query QueryWithSlices {
	user {
		emails
		emailsOrNull
		emailsWithNulls
		emailsWithNullsOrNull
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}
