package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

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

// StructOptionResponse is returned by StructOption on success.
type StructOptionResponse struct {
	Root StructOptionRootTopic `json:"root"`
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User StructOptionUser `json:"user"`
}

// StructOptionRootTopic includes the requested fields of the GraphQL type Topic.
type StructOptionRootTopic struct {
	// ID is documented in the Content interface.
	Id       testutil.ID                            `json:"id"`
	Children []StructOptionRootTopicChildrenContent `json:"children"`
}

// StructOptionRootTopicChildrenContent includes the requested fields of the GraphQL type Content.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type StructOptionRootTopicChildrenContent struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id     testutil.ID                                     `json:"id"`
	Parent StructOptionRootTopicChildrenContentParentTopic `json:"parent"`
}

// StructOptionRootTopicChildrenContentParentTopic includes the requested fields of the GraphQL type Topic.
type StructOptionRootTopicChildrenContentParentTopic struct {
	// ID is documented in the Content interface.
	Id                testutil.ID                                                               `json:"id"`
	Children          []StructOptionRootTopicChildrenContentParentTopicChildrenContent          `json:"children"`
	InterfaceChildren []StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent `json:"-"`
}

func (v *StructOptionRootTopicChildrenContentParentTopic) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*StructOptionRootTopicChildrenContentParentTopic
		InterfaceChildren []json.RawMessage `json:"interfaceChildren"`
		graphql.NoUnmarshalJSON
	}
	firstPass.StructOptionRootTopicChildrenContentParentTopic = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.InterfaceChildren
		src := firstPass.InterfaceChildren
		*dst = make(
			[]StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			if len(src) != 0 && string(src) != "null" {
				err = __unmarshalStructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent(
					src, dst)
				if err != nil {
					return fmt.Errorf(
						"Unable to unmarshal StructOptionRootTopicChildrenContentParentTopic.InterfaceChildren: %w", err)
				}
			}
		}
	}
	return nil
}

// StructOptionRootTopicChildrenContentParentTopicChildrenContent includes the requested fields of the GraphQL type Content.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type StructOptionRootTopicChildrenContentParentTopicChildrenContent struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

// StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenArticle includes the requested fields of the GraphQL type Article.
type StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

// StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent includes the requested fields of the GraphQL interface Content.
//
// StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent is implemented by the following types:
// StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenArticle
// StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo
// StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenTopic
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent interface {
	implementsGraphQLInterfaceStructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
}

func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenArticle) implementsGraphQLInterfaceStructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent() {
}

// GetTypename is a part of, and documented with, the interface StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent.
func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenArticle) GetTypename() string {
	return v.Typename
}

// GetId is a part of, and documented with, the interface StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent.
func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenArticle) GetId() testutil.ID {
	return v.Id
}

func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo) implementsGraphQLInterfaceStructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent() {
}

// GetTypename is a part of, and documented with, the interface StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent.
func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo) GetTypename() string {
	return v.Typename
}

// GetId is a part of, and documented with, the interface StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent.
func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo) GetId() testutil.ID {
	return v.Id
}

func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenTopic) implementsGraphQLInterfaceStructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent() {
}

// GetTypename is a part of, and documented with, the interface StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent.
func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenTopic) GetTypename() string {
	return v.Typename
}

// GetId is a part of, and documented with, the interface StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent.
func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenTopic) GetId() testutil.ID {
	return v.Id
}

func __unmarshalStructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent(b []byte, v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenArticle)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenTopic)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"Response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenContent: "%v"`, tn.TypeName)
	}
}

// StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenTopic includes the requested fields of the GraphQL type Topic.
type StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

// StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo includes the requested fields of the GraphQL type Video.
type StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id          testutil.ID `json:"id"`
	VideoFields `json:"-"`
}

func (v *StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo
		graphql.NoUnmarshalJSON
	}
	firstPass.StructOptionRootTopicChildrenContentParentTopicInterfaceChildrenVideo = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.VideoFields)
	if err != nil {
		return err
	}
	return nil
}

// StructOptionUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type StructOptionUser struct {
	Roles []Role `json:"roles"`
}

// VideoFields includes the GraphQL fields of Video requested by the fragment VideoFields.
type VideoFields struct {
	Duration int `json:"duration"`
}

func StructOption(
	client graphql.Client,
) (*StructOptionResponse, error) {
	var err error

	var retval StructOptionResponse
	err = client.MakeRequest(
		nil,
		"StructOption",
		`
query StructOption {
	root {
		id
		children {
			__typename
			id
			parent {
				id
				children {
					__typename
					id
				}
				interfaceChildren: children {
					__typename
					id
					... VideoFields
				}
			}
		}
	}
	user {
		roles
	}
}
fragment VideoFields on Video {
	duration
}
`,
		&retval,
		nil,
	)
	return &retval, err
}

