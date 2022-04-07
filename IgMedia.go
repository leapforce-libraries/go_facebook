package facebook

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	f_types "github.com/leapforce-libraries/go_facebook/types"
	go_http "github.com/leapforce-libraries/go_http"
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

type IgMediaResponse struct {
	Data    []IgMedia           `json:"data"`
	Paging  Paging              `json:"paging"`
	Summary PostCommentsSummary `json:"summary"`
}

type GetIgMediasConfig struct {
	IgUserId string
	Fields   *[]IgMediaField
	Since    *time.Time
	Until    *time.Time
}

func (service *Service) GetIgMedias(config *GetIgMediasConfig) (*[]IgMedia, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetIgMediasConfig must not be a nil pointer")
	}

	values := url.Values{}

	fields := []string{string(IgMediaFieldId)}
	if config.Fields != nil {
		for _, field := range *config.Fields {
			if field == IgMediaFieldId {
				continue
			}
			fields = append(fields, string(field))
		}
	}
	values.Set("fields", strings.Join(fields, ","))

	if config.Since != nil {
		values.Set("since", fmt.Sprintf("%v", config.Since.Unix()))
	}

	if config.Until != nil {
		values.Set("until", fmt.Sprintf("%v", config.Until.Unix()))
	}

	url := service.url(fmt.Sprintf("%s/media?%s", config.IgUserId, values.Encode()))

	igMedias := []IgMedia{}

	for {
		response := IgMediaResponse{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           url,
			ResponseModel: &response,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		igMedias = append(igMedias, response.Data...)

		if response.Paging.Next == "" {
			break
		}

		url = response.Paging.Next
	}

	return &igMedias, nil
}

type GetIgMediaConfig struct {
	IgMediaId string
	Fields    *[]IgMediaField
}

func (service *Service) GetIgMedia(config *GetIgMediaConfig) (*IgMedia, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetIgMediasConfig must not be a nil pointer")
	}

	values := url.Values{}

	fields := []string{string(IgMediaFieldId)}
	if config.Fields != nil {
		for _, field := range *config.Fields {
			if field == IgMediaFieldId {
				continue
			}
			fields = append(fields, string(field))
		}
	}
	values.Set("fields", strings.Join(fields, ","))

	url := service.url(fmt.Sprintf("%s?%s", config.IgMediaId, values.Encode()))

	igMedia := IgMedia{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           url,
		ResponseModel: &igMedia,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &igMedia, nil
}
