package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// UserQueryInput is the argument to Query.users.
//
// Ideally this would support anything and everything!
// Or maybe ideally it wouldn't.
// Really I'm just talking to make this documentation longer.
type MyInput struct {
	Email *string `json:"email"`
	Name  *string `json:"name"`
	// id looks the user up by ID.  It's a great way to look up users.
	Id         *testutil.ID      `json:"id"`
	Role       *Role             `json:"role"`
	Names      []*string         `json:"names"`
	HasPokemon *testutil.Pokemon `json:"hasPokemon"`
	Birthdate  *time.Time        `json:"-"`
}

func (v *MyInput) MarshalJSON() ([]byte, error) {

	var fullObject struct {
		*MyInput
		Birthdate json.RawMessage `json:"birthdate"`
		graphql.NoMarshalJSON
	}
	fullObject.MyInput = v

	{

		dst := &fullObject.Birthdate
		src := v.Birthdate
		if src != nil {
			var err error
			*dst, err = testutil.MarshalDate(
				src)
			if err != nil {
				return nil, fmt.Errorf(
					"Unable to marshal MyInput.Birthdate: %w", err)
			}
		}
	}

	return json.Marshal(&fullObject)
}

// MyMultipleDirectivesResponse is returned by MultipleDirectives on success.
type MyMultipleDirectivesResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User  *MyMultipleDirectivesResponseUser        `json:"user"`
	Users []*MyMultipleDirectivesResponseUsersUser `json:"users"`
}

// MyMultipleDirectivesResponseUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type MyMultipleDirectivesResponseUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id *testutil.ID `json:"id"`
}

// MyMultipleDirectivesResponseUsersUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type MyMultipleDirectivesResponseUsersUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id *testutil.ID `json:"id"`
}

// Role is a type a user may have.
type Role string

const (
	// What is a student?
	//
	// A student is primarily a person enrolled in a school or other educational institution and who is under learning with goals of acquiring knowledge, developing professions and achieving employment at desired field. In the broader sense, a student is anyone who applies themselves to the intensive intellectual engagement with some matter necessary to master it as part of some practical affair in which such mastery is basic or decisive.
	//
	// (from [Wikipedia](https://en.wikipedia.org/wiki/Student))
	RoleStudent Role = "STUDENT"
	// Teacher is a teacher, who teaches the students.
	RoleTeacher Role = "TEACHER"
)

// UserQueryInput is the argument to Query.users.
//
// Ideally this would support anything and everything!
// Or maybe ideally it wouldn't.
// Really I'm just talking to make this documentation longer.
type UserQueryInput struct {
	Email *string `json:"email"`
	Name  *string `json:"name"`
	// id looks the user up by ID.  It's a great way to look up users.
	Id         *testutil.ID      `json:"id"`
	Role       *Role             `json:"role"`
	Names      []*string         `json:"names"`
	HasPokemon *testutil.Pokemon `json:"hasPokemon"`
	Birthdate  *time.Time        `json:"-"`
}

func (v *UserQueryInput) MarshalJSON() ([]byte, error) {

	var fullObject struct {
		*UserQueryInput
		Birthdate json.RawMessage `json:"birthdate"`
		graphql.NoMarshalJSON
	}
	fullObject.UserQueryInput = v

	{

		dst := &fullObject.Birthdate
		src := v.Birthdate
		if src != nil {
			var err error
			*dst, err = testutil.MarshalDate(
				src)
			if err != nil {
				return nil, fmt.Errorf(
					"Unable to marshal UserQueryInput.Birthdate: %w", err)
			}
		}
	}

	return json.Marshal(&fullObject)
}

// __MultipleDirectivesInput is used internally by genqlient
type __MultipleDirectivesInput struct {
	Query   MyInput           `json:"query,omitempty"`
	Queries []*UserQueryInput `json:"queries,omitempty"`
}

func MultipleDirectives(
	client graphql.Client,
	query MyInput,
	queries []*UserQueryInput,
) (*MyMultipleDirectivesResponse, error) {
	__input := __MultipleDirectivesInput{
		Query:   query,
		Queries: queries,
	}
	var err error

	var retval MyMultipleDirectivesResponse
	err = client.MakeRequest(
		nil,
		"MultipleDirectives",
		`
query MultipleDirectives ($query: UserQueryInput, $queries: [UserQueryInput]) {
	user(query: $query) {
		id
	}
	users(query: $queries) {
		id
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}

