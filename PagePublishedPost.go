package facebook

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	f_types "github.com/leapforce-libraries/go_facebook/types"
	go_http "github.com/leapforce-libraries/go_http"
)

const pagePublishedPostsLimit int64 = 50 //limit 100 icm comments does not work...

type PagePublishedPostsResponse struct {
	Data    []PagePublishedPost `json:"data"`
	Paging  Paging              `json:"paging"`
	Summary PostCommentsSummary `json:"summary"`
}

type PagePublishedPost struct {
	ID           string                  `json:"id"`
	Attachments  *Attachments            `json:"attachments"`
	CreatedTime  *f_types.DateTimeString `json:"created_time"`
	From         PagePostFrom            `json:"from"`
	FullPicture  *string                 `json:"full_picture"`
	Message      *string                 `json:"message"`
	PermalinkURL *string                 `json:"permalink_url"`
	Shares       PagePostShares          `json:"shares"`
	StatusType   *string                 `json:"status_type"`
	UpdatedTime  *f_types.DateTimeString `json:"updated_time"`
}

type PagePostFrom struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PagePostShares struct {
	Count int64 `json:"count"`
}

type PagePublishedPostField string

const (
	PagePublishedPostFieldID           PagePublishedPostField = "id"
	PagePublishedPostFieldAttachments  PagePublishedPostField = "attachments"
	PagePublishedPostFieldCreatedTime  PagePublishedPostField = "created_time"
	PagePublishedPostFieldFrom         PagePublishedPostField = "from"
	PagePublishedPostFieldFullPicture  PagePublishedPostField = "full_picture"
	PagePublishedPostFieldMessage      PagePublishedPostField = "message"
	PagePublishedPostFieldPermalinkURL PagePublishedPostField = "permalink_url"
	PagePublishedPostFieldShares       PagePublishedPostField = "shares"
	PagePublishedPostFieldStatusType   PagePublishedPostField = "status_type"
	PagePublishedPostFieldUpdatedTime  PagePublishedPostField = "updated_time"
)

type GetPagePublishedPostsConfig struct {
	PageID string
	After  *string
	Limit  *int64
	Fields *[]PagePublishedPostField
}

// GetPagePublishedPosts returns Facebook post comments for a post
//
func (service *Service) GetPagePublishedPostsRequest(config *GetPagePublishedPostsConfig) (*go_http.RequestConfig, *PagePublishedPostsResponse, *errortools.Error) {
	if config == nil {
		return nil, nil, errortools.ErrorMessage("GetAccountsConfig must not be a nil pointer")
	}

	values := url.Values{}
	limit := pagePublishedPostsLimit
	if config.Limit != nil {
		limit = *config.Limit
	}
	values.Set("limit", fmt.Sprintf("%v", limit))
	if config.After != nil {
		values.Set("after", *config.After)
	}
	fields := []string{string(PagePublishedPostFieldID)}
	if config.Fields != nil {
		for _, field := range *config.Fields {
			fields = append(fields, string(field))
		}
	}
	values.Set("fields", strings.Join(fields, ","))

	response := PagePublishedPostsResponse{}
	relativeURL := fmt.Sprintf("%s/published_posts?%s", config.PageID, values.Encode())
	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		RelativeURL:   relativeURL,
		URL:           service.url(relativeURL),
		ResponseModel: &response,
	}

	return &requestConfig, &response, nil
}

// GetPagePublishedPosts returns Facebook post comments for a post
//
func (service *Service) GetPagePublishedPosts(config *GetPagePublishedPostsConfig) (*PagePublishedPostsResponse, *errortools.Error) {
	requestConfig, response, e := service.GetPagePublishedPostsRequest(config)
	if e != nil {
		return nil, e
	}
	_, _, e = service.httpRequest(requestConfig)
	if e != nil {
		return nil, e
	}

	return response, nil
}
