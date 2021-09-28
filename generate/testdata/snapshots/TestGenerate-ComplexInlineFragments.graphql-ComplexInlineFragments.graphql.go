package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// ComplexInlineFragmentsConflictingStuffArticle includes the requested fields of the GraphQL type Article.
type ComplexInlineFragmentsConflictingStuffArticle struct {
	Typename  string                                                               `json:"__typename"`
	Thumbnail ComplexInlineFragmentsConflictingStuffArticleThumbnailStuffThumbnail `json:"thumbnail"`
}

// ComplexInlineFragmentsConflictingStuffArticleThumbnailStuffThumbnail includes the requested fields of the GraphQL type StuffThumbnail.
type ComplexInlineFragmentsConflictingStuffArticleThumbnailStuffThumbnail struct {
	Id           testutil.ID `json:"id"`
	ThumbnailUrl string      `json:"thumbnailUrl"`
}

// ComplexInlineFragmentsConflictingStuffContent includes the requested fields of the GraphQL interface Content.
//
// ComplexInlineFragmentsConflictingStuffContent is implemented by the following types:
// ComplexInlineFragmentsConflictingStuffArticle
// ComplexInlineFragmentsConflictingStuffVideo
// ComplexInlineFragmentsConflictingStuffTopic
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type ComplexInlineFragmentsConflictingStuffContent interface {
	implementsGraphQLInterfaceComplexInlineFragmentsConflictingStuffContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
}

func (v *ComplexInlineFragmentsConflictingStuffArticle) implementsGraphQLInterfaceComplexInlineFragmentsConflictingStuffContent() {
}

// GetTypename is a part of, and documented with, the interface ComplexInlineFragmentsConflictingStuffContent.
func (v *ComplexInlineFragmentsConflictingStuffArticle) GetTypename() string { return v.Typename }

func (v *ComplexInlineFragmentsConflictingStuffVideo) implementsGraphQLInterfaceComplexInlineFragmentsConflictingStuffContent() {
}

// GetTypename is a part of, and documented with, the interface ComplexInlineFragmentsConflictingStuffContent.
func (v *ComplexInlineFragmentsConflictingStuffVideo) GetTypename() string { return v.Typename }

func (v *ComplexInlineFragmentsConflictingStuffTopic) implementsGraphQLInterfaceComplexInlineFragmentsConflictingStuffContent() {
}

// GetTypename is a part of, and documented with, the interface ComplexInlineFragmentsConflictingStuffContent.
func (v *ComplexInlineFragmentsConflictingStuffTopic) GetTypename() string { return v.Typename }

func __unmarshalComplexInlineFragmentsConflictingStuffContent(b []byte, v *ComplexInlineFragmentsConflictingStuffContent) error {
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
		*v = new(ComplexInlineFragmentsConflictingStuffArticle)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(ComplexInlineFragmentsConflictingStuffVideo)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(ComplexInlineFragmentsConflictingStuffTopic)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"Response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for ComplexInlineFragmentsConflictingStuffContent: "%v"`, tn.TypeName)
	}
}

// ComplexInlineFragmentsConflictingStuffTopic includes the requested fields of the GraphQL type Topic.
type ComplexInlineFragmentsConflictingStuffTopic struct {
	Typename string `json:"__typename"`
}

// ComplexInlineFragmentsConflictingStuffVideo includes the requested fields of the GraphQL type Video.
type ComplexInlineFragmentsConflictingStuffVideo struct {
	Typename  string                                               `json:"__typename"`
	Thumbnail ComplexInlineFragmentsConflictingStuffVideoThumbnail `json:"thumbnail"`
}

// ComplexInlineFragmentsConflictingStuffVideoThumbnail includes the requested fields of the GraphQL type Thumbnail.
type ComplexInlineFragmentsConflictingStuffVideoThumbnail struct {
	Id           testutil.ID `json:"id"`
	TimestampSec int         `json:"timestampSec"`
}

// ComplexInlineFragmentsNestedStuffArticle includes the requested fields of the GraphQL type Article.
type ComplexInlineFragmentsNestedStuffArticle struct {
	Typename string `json:"__typename"`
}

// ComplexInlineFragmentsNestedStuffContent includes the requested fields of the GraphQL interface Content.
//
// ComplexInlineFragmentsNestedStuffContent is implemented by the following types:
// ComplexInlineFragmentsNestedStuffArticle
// ComplexInlineFragmentsNestedStuffVideo
// ComplexInlineFragmentsNestedStuffTopic
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type ComplexInlineFragmentsNestedStuffContent interface {
	implementsGraphQLInterfaceComplexInlineFragmentsNestedStuffContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
}

func (v *ComplexInlineFragmentsNestedStuffArticle) implementsGraphQLInterfaceComplexInlineFragmentsNestedStuffContent() {
}

// GetTypename is a part of, and documented with, the interface ComplexInlineFragmentsNestedStuffContent.
func (v *ComplexInlineFragmentsNestedStuffArticle) GetTypename() string { return v.Typename }

func (v *ComplexInlineFragmentsNestedStuffVideo) implementsGraphQLInterfaceComplexInlineFragmentsNestedStuffContent() {
}

// GetTypename is a part of, and documented with, the interface ComplexInlineFragmentsNestedStuffContent.
func (v *ComplexInlineFragmentsNestedStuffVideo) GetTypename() string { return v.Typename }

func (v *ComplexInlineFragmentsNestedStuffTopic) implementsGraphQLInterfaceComplexInlineFragmentsNestedStuffContent() {
}

// GetTypename is a part of, and documented with, the interface ComplexInlineFragmentsNestedStuffContent.
func (v *ComplexInlineFragmentsNestedStuffTopic) GetTypename() string { return v.Typename }

func __unmarshalComplexInlineFragmentsNestedStuffContent(b []byte, v *ComplexInlineFragmentsNestedStuffContent) error {
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
		*v = new(ComplexInlineFragmentsNestedStuffArticle)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(ComplexInlineFragmentsNestedStuffVideo)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(ComplexInlineFragmentsNestedStuffTopic)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"Response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for ComplexInlineFragmentsNestedStuffContent: "%v"`, tn.TypeName)
	}
}

// ComplexInlineFragmentsNestedStuffTopic includes the requested fields of the GraphQL type Topic.
type ComplexInlineFragmentsNestedStuffTopic struct {
	Typename string                                                  `json:"__typename"`
	Children []ComplexInlineFragmentsNestedStuffTopicChildrenContent `json:"-"`
}

func (v *ComplexInlineFragmentsNestedStuffTopic) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*ComplexInlineFragmentsNestedStuffTopic
		Children []json.RawMessage `json:"children"`
		graphql.NoUnmarshalJSON
	}
	firstPass.ComplexInlineFragmentsNestedStuffTopic = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.Children
		src := firstPass.Children
		*dst = make(
			[]ComplexInlineFragmentsNestedStuffTopicChildrenContent,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			if len(src) != 0 && string(src) != "null" {
				err = __unmarshalComplexInlineFragmentsNestedStuffTopicChildrenContent(
					src, dst)
				if err != nil {
					return fmt.Errorf(
						"Unable to unmarshal ComplexInlineFragmentsNestedStuffTopic.Children: %w", err)
				}
			}
		}
	}
	return nil
}

// ComplexInlineFragmentsNestedStuffTopicChildrenArticle includes the requested fields of the GraphQL type Article.
type ComplexInlineFragmentsNestedStuffTopicChildrenArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id     testutil.ID                                                      `json:"id"`
	Text   string                                                           `json:"text"`
	Parent ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentTopic `json:"parent"`
}

// ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopic includes the requested fields of the GraphQL type Topic.
type ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopic struct {
	Children []ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent `json:"-"`
}

func (v *ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopic) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopic
		Children []json.RawMessage `json:"children"`
		graphql.NoUnmarshalJSON
	}
	firstPass.ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopic = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.Children
		src := firstPass.Children
		*dst = make(
			[]ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			if len(src) != 0 && string(src) != "null" {
				err = __unmarshalComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent(
					src, dst)
				if err != nil {
					return fmt.Errorf(
						"Unable to unmarshal ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopic.Children: %w", err)
				}
			}
		}
	}
	return nil
}

// ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenArticle includes the requested fields of the GraphQL type Article.
type ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent includes the requested fields of the GraphQL interface Content.
//
// ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent is implemented by the following types:
// ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenArticle
// ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenVideo
// ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenTopic
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent interface {
	implementsGraphQLInterfaceComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
	// GetName returns the interface-field "name" from its implementation.
	GetName() string
}

func (v *ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenArticle) implementsGraphQLInterfaceComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent() {
}

// GetTypename is a part of, and documented with, the interface ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent.
func (v *ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenArticle) GetTypename() string {
	return v.Typename
}

// GetId is a part of, and documented with, the interface ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent.
func (v *ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenArticle) GetId() testutil.ID {
	return v.Id
}

// GetName is a part of, and documented with, the interface ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent.
func (v *ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenArticle) GetName() string {
	return v.Name
}

func (v *ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenVideo) implementsGraphQLInterfaceComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent() {
}

// GetTypename is a part of, and documented with, the interface ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent.
func (v *ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenVideo) GetTypename() string {
	return v.Typename
}

// GetId is a part of, and documented with, the interface ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent.
func (v *ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenVideo) GetId() testutil.ID {
	return v.Id
}

// GetName is a part of, and documented with, the interface ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent.
func (v *ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenVideo) GetName() string {
	return v.Name
}

func (v *ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenTopic) implementsGraphQLInterfaceComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent() {
}

// GetTypename is a part of, and documented with, the interface ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent.
func (v *ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenTopic) GetTypename() string {
	return v.Typename
}

// GetId is a part of, and documented with, the interface ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent.
func (v *ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenTopic) GetId() testutil.ID {
	return v.Id
}

// GetName is a part of, and documented with, the interface ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent.
func (v *ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenTopic) GetName() string {
	return v.Name
}

func __unmarshalComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent(b []byte, v *ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent) error {
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
		*v = new(ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenArticle)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenVideo)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenTopic)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"Response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenContent: "%v"`, tn.TypeName)
	}
}

// ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenTopic includes the requested fields of the GraphQL type Topic.
type ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenVideo includes the requested fields of the GraphQL type Video.
type ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopicChildrenVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentTopic includes the requested fields of the GraphQL type Topic.
type ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentTopic struct {
	Name   string                                                                        `json:"name"`
	Parent ComplexInlineFragmentsNestedStuffTopicChildrenArticleParentContentParentTopic `json:"parent"`
}

// ComplexInlineFragmentsNestedStuffTopicChildrenContent includes the requested fields of the GraphQL interface Content.
//
// ComplexInlineFragmentsNestedStuffTopicChildrenContent is implemented by the following types:
// ComplexInlineFragmentsNestedStuffTopicChildrenArticle
// ComplexInlineFragmentsNestedStuffTopicChildrenVideo
// ComplexInlineFragmentsNestedStuffTopicChildrenTopic
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type ComplexInlineFragmentsNestedStuffTopicChildrenContent interface {
	implementsGraphQLInterfaceComplexInlineFragmentsNestedStuffTopicChildrenContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
}

func (v *ComplexInlineFragmentsNestedStuffTopicChildrenArticle) implementsGraphQLInterfaceComplexInlineFragmentsNestedStuffTopicChildrenContent() {
}

// GetTypename is a part of, and documented with, the interface ComplexInlineFragmentsNestedStuffTopicChildrenContent.
func (v *ComplexInlineFragmentsNestedStuffTopicChildrenArticle) GetTypename() string {
	return v.Typename
}

// GetId is a part of, and documented with, the interface ComplexInlineFragmentsNestedStuffTopicChildrenContent.
func (v *ComplexInlineFragmentsNestedStuffTopicChildrenArticle) GetId() testutil.ID { return v.Id }

func (v *ComplexInlineFragmentsNestedStuffTopicChildrenVideo) implementsGraphQLInterfaceComplexInlineFragmentsNestedStuffTopicChildrenContent() {
}

// GetTypename is a part of, and documented with, the interface ComplexInlineFragmentsNestedStuffTopicChildrenContent.
func (v *ComplexInlineFragmentsNestedStuffTopicChildrenVideo) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface ComplexInlineFragmentsNestedStuffTopicChildrenContent.
func (v *ComplexInlineFragmentsNestedStuffTopicChildrenVideo) GetId() testutil.ID { return v.Id }

func (v *ComplexInlineFragmentsNestedStuffTopicChildrenTopic) implementsGraphQLInterfaceComplexInlineFragmentsNestedStuffTopicChildrenContent() {
}

// GetTypename is a part of, and documented with, the interface ComplexInlineFragmentsNestedStuffTopicChildrenContent.
func (v *ComplexInlineFragmentsNestedStuffTopicChildrenTopic) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface ComplexInlineFragmentsNestedStuffTopicChildrenContent.
func (v *ComplexInlineFragmentsNestedStuffTopicChildrenTopic) GetId() testutil.ID { return v.Id }

func __unmarshalComplexInlineFragmentsNestedStuffTopicChildrenContent(b []byte, v *ComplexInlineFragmentsNestedStuffTopicChildrenContent) error {
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
		*v = new(ComplexInlineFragmentsNestedStuffTopicChildrenArticle)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(ComplexInlineFragmentsNestedStuffTopicChildrenVideo)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(ComplexInlineFragmentsNestedStuffTopicChildrenTopic)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"Response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for ComplexInlineFragmentsNestedStuffTopicChildrenContent: "%v"`, tn.TypeName)
	}
}

// ComplexInlineFragmentsNestedStuffTopicChildrenTopic includes the requested fields of the GraphQL type Topic.
type ComplexInlineFragmentsNestedStuffTopicChildrenTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

// ComplexInlineFragmentsNestedStuffTopicChildrenVideo includes the requested fields of the GraphQL type Video.
type ComplexInlineFragmentsNestedStuffTopicChildrenVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

// ComplexInlineFragmentsNestedStuffVideo includes the requested fields of the GraphQL type Video.
type ComplexInlineFragmentsNestedStuffVideo struct {
	Typename string `json:"__typename"`
}

// ComplexInlineFragmentsRandomItemArticle includes the requested fields of the GraphQL type Article.
type ComplexInlineFragmentsRandomItemArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Text string      `json:"text"`
	Name string      `json:"name"`
}

// ComplexInlineFragmentsRandomItemContent includes the requested fields of the GraphQL interface Content.
//
// ComplexInlineFragmentsRandomItemContent is implemented by the following types:
// ComplexInlineFragmentsRandomItemArticle
// ComplexInlineFragmentsRandomItemVideo
// ComplexInlineFragmentsRandomItemTopic
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type ComplexInlineFragmentsRandomItemContent interface {
	implementsGraphQLInterfaceComplexInlineFragmentsRandomItemContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
	// GetName returns the interface-field "name" from its implementation.
	GetName() string
}

func (v *ComplexInlineFragmentsRandomItemArticle) implementsGraphQLInterfaceComplexInlineFragmentsRandomItemContent() {
}

// GetTypename is a part of, and documented with, the interface ComplexInlineFragmentsRandomItemContent.
func (v *ComplexInlineFragmentsRandomItemArticle) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface ComplexInlineFragmentsRandomItemContent.
func (v *ComplexInlineFragmentsRandomItemArticle) GetId() testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface ComplexInlineFragmentsRandomItemContent.
func (v *ComplexInlineFragmentsRandomItemArticle) GetName() string { return v.Name }

func (v *ComplexInlineFragmentsRandomItemVideo) implementsGraphQLInterfaceComplexInlineFragmentsRandomItemContent() {
}

// GetTypename is a part of, and documented with, the interface ComplexInlineFragmentsRandomItemContent.
func (v *ComplexInlineFragmentsRandomItemVideo) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface ComplexInlineFragmentsRandomItemContent.
func (v *ComplexInlineFragmentsRandomItemVideo) GetId() testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface ComplexInlineFragmentsRandomItemContent.
func (v *ComplexInlineFragmentsRandomItemVideo) GetName() string { return v.Name }

func (v *ComplexInlineFragmentsRandomItemTopic) implementsGraphQLInterfaceComplexInlineFragmentsRandomItemContent() {
}

// GetTypename is a part of, and documented with, the interface ComplexInlineFragmentsRandomItemContent.
func (v *ComplexInlineFragmentsRandomItemTopic) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface ComplexInlineFragmentsRandomItemContent.
func (v *ComplexInlineFragmentsRandomItemTopic) GetId() testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface ComplexInlineFragmentsRandomItemContent.
func (v *ComplexInlineFragmentsRandomItemTopic) GetName() string { return v.Name }

func __unmarshalComplexInlineFragmentsRandomItemContent(b []byte, v *ComplexInlineFragmentsRandomItemContent) error {
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
		*v = new(ComplexInlineFragmentsRandomItemArticle)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(ComplexInlineFragmentsRandomItemVideo)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(ComplexInlineFragmentsRandomItemTopic)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"Response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for ComplexInlineFragmentsRandomItemContent: "%v"`, tn.TypeName)
	}
}

// ComplexInlineFragmentsRandomItemTopic includes the requested fields of the GraphQL type Topic.
type ComplexInlineFragmentsRandomItemTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// ComplexInlineFragmentsRandomItemVideo includes the requested fields of the GraphQL type Video.
type ComplexInlineFragmentsRandomItemVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id       testutil.ID `json:"id"`
	Name     string      `json:"name"`
	Duration int         `json:"duration"`
}

// ComplexInlineFragmentsRepeatedStuffArticle includes the requested fields of the GraphQL type Article.
type ComplexInlineFragmentsRepeatedStuffArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id  testutil.ID `json:"id"`
	Url string      `json:"url"`
	// ID is the identifier of the content.
	OtherId   testutil.ID `json:"otherId"`
	Name      string      `json:"name"`
	Text      string      `json:"text"`
	OtherName string      `json:"otherName"`
}

// ComplexInlineFragmentsRepeatedStuffContent includes the requested fields of the GraphQL interface Content.
//
// ComplexInlineFragmentsRepeatedStuffContent is implemented by the following types:
// ComplexInlineFragmentsRepeatedStuffArticle
// ComplexInlineFragmentsRepeatedStuffVideo
// ComplexInlineFragmentsRepeatedStuffTopic
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type ComplexInlineFragmentsRepeatedStuffContent interface {
	implementsGraphQLInterfaceComplexInlineFragmentsRepeatedStuffContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
	// GetUrl returns the interface-field "url" from its implementation.
	GetUrl() string
	// GetOtherId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetOtherId() testutil.ID
	// GetName returns the interface-field "name" from its implementation.
	GetName() string
	// GetOtherName returns the interface-field "name" from its implementation.
	GetOtherName() string
}

func (v *ComplexInlineFragmentsRepeatedStuffArticle) implementsGraphQLInterfaceComplexInlineFragmentsRepeatedStuffContent() {
}

// GetTypename is a part of, and documented with, the interface ComplexInlineFragmentsRepeatedStuffContent.
func (v *ComplexInlineFragmentsRepeatedStuffArticle) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface ComplexInlineFragmentsRepeatedStuffContent.
func (v *ComplexInlineFragmentsRepeatedStuffArticle) GetId() testutil.ID { return v.Id }

// GetUrl is a part of, and documented with, the interface ComplexInlineFragmentsRepeatedStuffContent.
func (v *ComplexInlineFragmentsRepeatedStuffArticle) GetUrl() string { return v.Url }

// GetOtherId is a part of, and documented with, the interface ComplexInlineFragmentsRepeatedStuffContent.
func (v *ComplexInlineFragmentsRepeatedStuffArticle) GetOtherId() testutil.ID { return v.OtherId }

// GetName is a part of, and documented with, the interface ComplexInlineFragmentsRepeatedStuffContent.
func (v *ComplexInlineFragmentsRepeatedStuffArticle) GetName() string { return v.Name }

// GetOtherName is a part of, and documented with, the interface ComplexInlineFragmentsRepeatedStuffContent.
func (v *ComplexInlineFragmentsRepeatedStuffArticle) GetOtherName() string { return v.OtherName }

func (v *ComplexInlineFragmentsRepeatedStuffVideo) implementsGraphQLInterfaceComplexInlineFragmentsRepeatedStuffContent() {
}

// GetTypename is a part of, and documented with, the interface ComplexInlineFragmentsRepeatedStuffContent.
func (v *ComplexInlineFragmentsRepeatedStuffVideo) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface ComplexInlineFragmentsRepeatedStuffContent.
func (v *ComplexInlineFragmentsRepeatedStuffVideo) GetId() testutil.ID { return v.Id }

// GetUrl is a part of, and documented with, the interface ComplexInlineFragmentsRepeatedStuffContent.
func (v *ComplexInlineFragmentsRepeatedStuffVideo) GetUrl() string { return v.Url }

// GetOtherId is a part of, and documented with, the interface ComplexInlineFragmentsRepeatedStuffContent.
func (v *ComplexInlineFragmentsRepeatedStuffVideo) GetOtherId() testutil.ID { return v.OtherId }

// GetName is a part of, and documented with, the interface ComplexInlineFragmentsRepeatedStuffContent.
func (v *ComplexInlineFragmentsRepeatedStuffVideo) GetName() string { return v.Name }

// GetOtherName is a part of, and documented with, the interface ComplexInlineFragmentsRepeatedStuffContent.
func (v *ComplexInlineFragmentsRepeatedStuffVideo) GetOtherName() string { return v.OtherName }

func (v *ComplexInlineFragmentsRepeatedStuffTopic) implementsGraphQLInterfaceComplexInlineFragmentsRepeatedStuffContent() {
}

// GetTypename is a part of, and documented with, the interface ComplexInlineFragmentsRepeatedStuffContent.
func (v *ComplexInlineFragmentsRepeatedStuffTopic) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface ComplexInlineFragmentsRepeatedStuffContent.
func (v *ComplexInlineFragmentsRepeatedStuffTopic) GetId() testutil.ID { return v.Id }

// GetUrl is a part of, and documented with, the interface ComplexInlineFragmentsRepeatedStuffContent.
func (v *ComplexInlineFragmentsRepeatedStuffTopic) GetUrl() string { return v.Url }

// GetOtherId is a part of, and documented with, the interface ComplexInlineFragmentsRepeatedStuffContent.
func (v *ComplexInlineFragmentsRepeatedStuffTopic) GetOtherId() testutil.ID { return v.OtherId }

// GetName is a part of, and documented with, the interface ComplexInlineFragmentsRepeatedStuffContent.
func (v *ComplexInlineFragmentsRepeatedStuffTopic) GetName() string { return v.Name }

// GetOtherName is a part of, and documented with, the interface ComplexInlineFragmentsRepeatedStuffContent.
func (v *ComplexInlineFragmentsRepeatedStuffTopic) GetOtherName() string { return v.OtherName }

func __unmarshalComplexInlineFragmentsRepeatedStuffContent(b []byte, v *ComplexInlineFragmentsRepeatedStuffContent) error {
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
		*v = new(ComplexInlineFragmentsRepeatedStuffArticle)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(ComplexInlineFragmentsRepeatedStuffVideo)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(ComplexInlineFragmentsRepeatedStuffTopic)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"Response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for ComplexInlineFragmentsRepeatedStuffContent: "%v"`, tn.TypeName)
	}
}

// ComplexInlineFragmentsRepeatedStuffTopic includes the requested fields of the GraphQL type Topic.
type ComplexInlineFragmentsRepeatedStuffTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id  testutil.ID `json:"id"`
	Url string      `json:"url"`
	// ID is the identifier of the content.
	OtherId   testutil.ID `json:"otherId"`
	Name      string      `json:"name"`
	OtherName string      `json:"otherName"`
}

// ComplexInlineFragmentsRepeatedStuffVideo includes the requested fields of the GraphQL type Video.
type ComplexInlineFragmentsRepeatedStuffVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id  testutil.ID `json:"id"`
	Url string      `json:"url"`
	// ID is the identifier of the content.
	OtherId   testutil.ID `json:"otherId"`
	Name      string      `json:"name"`
	OtherName string      `json:"otherName"`
	Duration  int         `json:"duration"`
}

// ComplexInlineFragmentsResponse is returned by ComplexInlineFragments on success.
type ComplexInlineFragmentsResponse struct {
	Root             ComplexInlineFragmentsRootTopic               `json:"root"`
	RandomItem       ComplexInlineFragmentsRandomItemContent       `json:"-"`
	RepeatedStuff    ComplexInlineFragmentsRepeatedStuffContent    `json:"-"`
	ConflictingStuff ComplexInlineFragmentsConflictingStuffContent `json:"-"`
	NestedStuff      ComplexInlineFragmentsNestedStuffContent      `json:"-"`
}

func (v *ComplexInlineFragmentsResponse) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*ComplexInlineFragmentsResponse
		RandomItem       json.RawMessage `json:"randomItem"`
		RepeatedStuff    json.RawMessage `json:"repeatedStuff"`
		ConflictingStuff json.RawMessage `json:"conflictingStuff"`
		NestedStuff      json.RawMessage `json:"nestedStuff"`
		graphql.NoUnmarshalJSON
	}
	firstPass.ComplexInlineFragmentsResponse = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.RandomItem
		src := firstPass.RandomItem
		if len(src) != 0 && string(src) != "null" {
			err = __unmarshalComplexInlineFragmentsRandomItemContent(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal ComplexInlineFragmentsResponse.RandomItem: %w", err)
			}
		}
	}

	{
		dst := &v.RepeatedStuff
		src := firstPass.RepeatedStuff
		if len(src) != 0 && string(src) != "null" {
			err = __unmarshalComplexInlineFragmentsRepeatedStuffContent(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal ComplexInlineFragmentsResponse.RepeatedStuff: %w", err)
			}
		}
	}

	{
		dst := &v.ConflictingStuff
		src := firstPass.ConflictingStuff
		if len(src) != 0 && string(src) != "null" {
			err = __unmarshalComplexInlineFragmentsConflictingStuffContent(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal ComplexInlineFragmentsResponse.ConflictingStuff: %w", err)
			}
		}
	}

	{
		dst := &v.NestedStuff
		src := firstPass.NestedStuff
		if len(src) != 0 && string(src) != "null" {
			err = __unmarshalComplexInlineFragmentsNestedStuffContent(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal ComplexInlineFragmentsResponse.NestedStuff: %w", err)
			}
		}
	}
	return nil
}

// ComplexInlineFragmentsRootTopic includes the requested fields of the GraphQL type Topic.
type ComplexInlineFragmentsRootTopic struct {
	// ID is documented in the Content interface.
	Id          testutil.ID `json:"id"`
	SchoolGrade string      `json:"schoolGrade"`
	Name        string      `json:"name"`
}

// We test all the spread cases from docs/DESIGN.md, see there for more context
// on each, as well as various other nonsense.  But for abstract-in-abstract
// spreads, we can't test cases (4b) and (4c), where I implements J or vice
// versa, because gqlparser doesn't support interfaces that implement other
// interfaces yet.
func ComplexInlineFragments(
	client graphql.Client,
) (*ComplexInlineFragmentsResponse, error) {
	var err error

	var retval ComplexInlineFragmentsResponse
	err = client.MakeRequest(
		nil,
		"ComplexInlineFragments",
		`
query ComplexInlineFragments {
	root {
		id
		... on Topic {
			schoolGrade
		}
		... on Content {
			name
		}
	}
	randomItem {
		__typename
		id
		... on Article {
			text
		}
		... on Content {
			name
		}
		... on HasDuration {
			duration
		}
	}
	repeatedStuff: randomItem {
		__typename
		id
		id
		url
		otherId: id
		... on Article {
			name
			text
			otherName: name
		}
		... on Content {
			id
			name
			otherName: name
		}
		... on HasDuration {
			duration
		}
	}
	conflictingStuff: randomItem {
		__typename
		... on Article {
			thumbnail {
				id
				thumbnailUrl
			}
		}
		... on Video {
			thumbnail {
				id
				timestampSec
			}
		}
	}
	nestedStuff: randomItem {
		__typename
		... on Topic {
			children {
				__typename
				id
				... on Article {
					text
					parent {
						... on Content {
							name
							parent {
								... on Topic {
									children {
										__typename
										id
										name
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}

