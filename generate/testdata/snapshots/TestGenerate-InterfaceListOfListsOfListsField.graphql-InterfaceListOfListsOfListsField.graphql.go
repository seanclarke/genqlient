package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// InterfaceListOfListOfListsFieldListOfListsOfListsOfContent includes the requested fields of the GraphQL type Content.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceListOfListOfListsFieldListOfListsOfListsOfContent interface {
	implementsGraphQLInterfaceInterfaceListOfListOfListsFieldListOfListsOfListsOfContent()
}

func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle) implementsGraphQLInterfaceInterfaceListOfListOfListsFieldListOfListsOfListsOfContent() {
}
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo) implementsGraphQLInterfaceInterfaceListOfListOfListsFieldListOfListsOfListsOfContent() {
}
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic) implementsGraphQLInterfaceInterfaceListOfListOfListsFieldListOfListsOfListsOfContent() {
}

func __unmarshalInterfaceListOfListOfListsFieldListOfListsOfListsOfContent(v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContent, m json.RawMessage) error {
	if string(m) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(m, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle)
		return json.Unmarshal(m, *v)
	case "Video":
		*v = new(InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo)
		return json.Unmarshal(m, *v)
	case "Topic":
		*v = new(InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic)
		return json.Unmarshal(m, *v)
	default:
		return fmt.Errorf(
			`Unexpected concrete type for InterfaceListOfListOfListsFieldListOfListsOfListsOfContent: "%v"`, tn.TypeName)
	}
}

// InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle includes the requested fields of the GraphQL type Article.
type InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic includes the requested fields of the GraphQL type Topic.
type InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo includes the requested fields of the GraphQL type Video.
type InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// InterfaceListOfListOfListsFieldResponse is returned by InterfaceListOfListOfListsField on success.
type InterfaceListOfListOfListsFieldResponse struct {
	ListOfListsOfListsOfContent [][][]InterfaceListOfListOfListsFieldListOfListsOfListsOfContent `json:"-"`
	WithPointer                 [][][]*InterfaceListOfListOfListsFieldWithPointerContent         `json:"-"`
}

func (v *InterfaceListOfListOfListsFieldResponse) UnmarshalJSON(b []byte) error {

	type InterfaceListOfListOfListsFieldResponseWrapper InterfaceListOfListOfListsFieldResponse

	var firstPass struct {
		*InterfaceListOfListOfListsFieldResponseWrapper
		ListOfListsOfListsOfContent [][][]json.RawMessage `json:"listOfListsOfListsOfContent"`
		WithPointer                 [][][]json.RawMessage `json:"withPointer"`
	}
	firstPass.InterfaceListOfListOfListsFieldResponseWrapper = (*InterfaceListOfListOfListsFieldResponseWrapper)(v)

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		target := &v.ListOfListsOfListsOfContent
		raw := firstPass.ListOfListsOfListsOfContent
		*target = make(
			[][][]InterfaceListOfListOfListsFieldListOfListsOfListsOfContent,
			len(raw))
		for i, raw := range raw {
			target := &(*target)[i]
			*target = make(
				[][]InterfaceListOfListOfListsFieldListOfListsOfListsOfContent,
				len(raw))
			for i, raw := range raw {
				target := &(*target)[i]
				*target = make(
					[]InterfaceListOfListOfListsFieldListOfListsOfListsOfContent,
					len(raw))
				for i, raw := range raw {
					target := &(*target)[i]
					err = __unmarshalInterfaceListOfListOfListsFieldListOfListsOfListsOfContent(
						target, raw)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	{
		target := &v.WithPointer
		raw := firstPass.WithPointer
		*target = make(
			[][][]*InterfaceListOfListOfListsFieldWithPointerContent,
			len(raw))
		for i, raw := range raw {
			target := &(*target)[i]
			*target = make(
				[][]*InterfaceListOfListOfListsFieldWithPointerContent,
				len(raw))
			for i, raw := range raw {
				target := &(*target)[i]
				*target = make(
					[]*InterfaceListOfListOfListsFieldWithPointerContent,
					len(raw))
				for i, raw := range raw {
					target := &(*target)[i]
					*target = new(InterfaceListOfListOfListsFieldWithPointerContent)
					err = __unmarshalInterfaceListOfListOfListsFieldWithPointerContent(
						*target, raw)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

// InterfaceListOfListOfListsFieldWithPointerArticle includes the requested fields of the GraphQL type Article.
type InterfaceListOfListOfListsFieldWithPointerArticle struct {
	Typename *string `json:"__typename"`
	// ID is the identifier of the content.
	Id   *testutil.ID `json:"id"`
	Name *string      `json:"name"`
}

// InterfaceListOfListOfListsFieldWithPointerContent includes the requested fields of the GraphQL type Content.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceListOfListOfListsFieldWithPointerContent interface {
	implementsGraphQLInterfaceInterfaceListOfListOfListsFieldWithPointerContent()
}

func (v *InterfaceListOfListOfListsFieldWithPointerArticle) implementsGraphQLInterfaceInterfaceListOfListOfListsFieldWithPointerContent() {
}
func (v *InterfaceListOfListOfListsFieldWithPointerVideo) implementsGraphQLInterfaceInterfaceListOfListOfListsFieldWithPointerContent() {
}
func (v *InterfaceListOfListOfListsFieldWithPointerTopic) implementsGraphQLInterfaceInterfaceListOfListOfListsFieldWithPointerContent() {
}

func __unmarshalInterfaceListOfListOfListsFieldWithPointerContent(v *InterfaceListOfListOfListsFieldWithPointerContent, m json.RawMessage) error {
	if string(m) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(m, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(InterfaceListOfListOfListsFieldWithPointerArticle)
		return json.Unmarshal(m, *v)
	case "Video":
		*v = new(InterfaceListOfListOfListsFieldWithPointerVideo)
		return json.Unmarshal(m, *v)
	case "Topic":
		*v = new(InterfaceListOfListOfListsFieldWithPointerTopic)
		return json.Unmarshal(m, *v)
	default:
		return fmt.Errorf(
			`Unexpected concrete type for InterfaceListOfListOfListsFieldWithPointerContent: "%v"`, tn.TypeName)
	}
}

// InterfaceListOfListOfListsFieldWithPointerTopic includes the requested fields of the GraphQL type Topic.
type InterfaceListOfListOfListsFieldWithPointerTopic struct {
	Typename *string `json:"__typename"`
	// ID is the identifier of the content.
	Id   *testutil.ID `json:"id"`
	Name *string      `json:"name"`
}

// InterfaceListOfListOfListsFieldWithPointerVideo includes the requested fields of the GraphQL type Video.
type InterfaceListOfListOfListsFieldWithPointerVideo struct {
	Typename *string `json:"__typename"`
	// ID is the identifier of the content.
	Id   *testutil.ID `json:"id"`
	Name *string      `json:"name"`
}

func InterfaceListOfListOfListsField(
	client graphql.Client,
) (*InterfaceListOfListOfListsFieldResponse, error) {
	var retval InterfaceListOfListOfListsFieldResponse
	err := client.MakeRequest(
		nil,
		"InterfaceListOfListOfListsField",
		`
query InterfaceListOfListOfListsField {
	listOfListsOfListsOfContent {
		__typename
		id
		name
	}
	withPointer: listOfListsOfListsOfContent {
		__typename
		id
		name
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}

