package test

// Code generated by github.com/Khan/genql, DO NOT EDIT.

import (
	"github.com/Khan/genql/graphql"
)

type UsesEnumTwiceQueryMeUser struct {
	Roles []UsesEnumTwiceQueryMeUserRolesRole `json:"roles"`
}

type UsesEnumTwiceQueryMeUserRolesRole string

const (
	UsesEnumTwiceQueryMeUserRolesRoleStudent UsesEnumTwiceQueryMeUserRolesRole = "STUDENT"
	UsesEnumTwiceQueryMeUserRolesRoleTeacher UsesEnumTwiceQueryMeUserRolesRole = "TEACHER"
)

type UsesEnumTwiceQueryOtherUser struct {
	Roles []UsesEnumTwiceQueryOtherUserRolesRole `json:"roles"`
}

type UsesEnumTwiceQueryOtherUserRolesRole string

const (
	UsesEnumTwiceQueryOtherUserRolesRoleStudent UsesEnumTwiceQueryOtherUserRolesRole = "STUDENT"
	UsesEnumTwiceQueryOtherUserRolesRoleTeacher UsesEnumTwiceQueryOtherUserRolesRole = "TEACHER"
)

type UsesEnumTwiceQueryResponse struct {
	Me        UsesEnumTwiceQueryMeUser
	OtherUser UsesEnumTwiceQueryOtherUser
}

func UsesEnumTwiceQuery(client *graphql.Client) (*UsesEnumTwiceQueryResponse, error) {
	var retval UsesEnumTwiceQueryResponse
	err := client.MakeRequest(nil, `
query UsesEnumTwiceQuery {
	Me: user {
		roles
	}
	OtherUser: user {
		roles
	}
}
`, &retval, nil)
	return &retval, err
}
