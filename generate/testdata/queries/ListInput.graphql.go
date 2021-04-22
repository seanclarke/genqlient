package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"github.com/Khan/genqlient/graphql"
	"github.com/me/mypkg"
)

// ListInputQueryResponse is returned by ListInputQuery on success.
type ListInputQueryResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User ListInputQueryUser `json:"user"`
}

// A User is a user!
type ListInputQueryUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id mypkg.ID `json:"id"`
}

func ListInputQuery(
	client graphql.Client,
	names []string,
) (*ListInputQueryResponse, error) {
	variables := map[string]interface{}{
		"names": names,
	}

	var retval ListInputQueryResponse
	err := client.MakeRequest(
		nil,
		"ListInputQuery",
		`
query ListInputQuery ($names: [String]) {
	user(query: {names:$names}) {
		id
	}
}
`,
		&retval,
		variables,
	)
	return &retval, err
}
