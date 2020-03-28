package example

// Code generated by github.com/Khan/genql, DO NOT EDIT.

import (
	"context"

	"github.com/Khan/genql/graphql"
)

type getViewerResponse = struct {
	Viewer struct {
		MyName *string
	} `json:"viewer"`
}

func getViewer(ctx context.Context, client *graphql.Client) (*getViewerResponse, error) {
	var retval getViewerResponse
	err := client.MakeRequest(ctx, `
query getViewer {
	viewer {
		MyName: name
	}
}
`, &retval, nil)
	return &retval, err
}

type getUserResponse = struct {
	User *struct {
		TheirName *string `json:"theirName"`
	} `json:"user"`
}

// getUser gets the given user's name from their username.
func getUser(ctx context.Context, client *graphql.Client, login string) (*getUserResponse, error) {
	variables := map[string]interface{}{
		"Login": login,
	}

	var retval getUserResponse
	err := client.MakeRequest(ctx, `
query getUser ($Login: String!) {
	user(login: $Login) {
		theirName: name
	}
}
`, &retval, variables)
	return &retval, err
}
