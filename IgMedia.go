package facebook

import (
	f_types "github.com/leapforce-libraries/go_facebook/types"
)

type IgMedia struct {
	Caption          *string `json:"caption"`
	CommentsCount    *int64  `json:"comments_count"`
	Id               string  `json:"id"`
	IgId             *string `json:"ig_id"`
	IsCommentEnabled *bool   `json:"is_comment_enabled"`
	LikeCount        *int64  `json:"like_count"`
	MediaProductType *string `json:"media_product_type"`
	MediaType        *string `json:"media_type"`
	MediaUrl         *string `json:"media_url"`
	Owner            *struct {
		Id *string `json:"id"`
	} `json:"owner"`
	Permalink    *string                 `json:"permalink"`
	ShortCode    *string                 `json:"shortcode"`
	ThumbnailUrl *string                 `json:"thumbnail_url"`
	Timestamp    *f_types.DateTimeString `json:"timestamp"`
	Username     *string                 `json:"username"`
}

type IgMediaField string

const (
	IgMediaFieldCaption          IgMediaField = "caption"
	IgMediaFieldCommentsCount    IgMediaField = "comments_count"
	IgMediaFieldId               IgMediaField = "id"
	IgMediaFieldIgId             IgMediaField = "ig_id"
	IgMediaFieldIsCommentEnabled IgMediaField = "is_comment_enabled"
	IgMediaFieldLikeCount        IgMediaField = "like_count"
	IgMediaFieldMediaProductType IgMediaField = "media_product_type"
	IgMediaFieldMediaType        IgMediaField = "media_type"
	IgMediaFieldMediaUrl         IgMediaField = "media_url"
	IgMediaFieldOwner            IgMediaField = "owner"
	IgMediaFieldPermalink        IgMediaField = "permalink"
	IgMediaFieldShortcode        IgMediaField = "shortcode"
	IgMediaFieldThumbnailUrl     IgMediaField = "thumbnail_url"
	IgMediaFieldTimestamp        IgMediaField = "timestamp"
	IgMediaFieldUsername         IgMediaField = "username"
)
