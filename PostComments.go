package facebook

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type PostCommentsFilter string

const (
	PostCommentsFilterTopLevel PostCommentsFilter = "toplevel"
	PostCommentsFilterStream   PostCommentsFilter = "stream"
)

const postCommentsLimit int64 = 50 //limit 100 icm comments does not work...

type PostCommentsResponse struct {
	//Data    []PagePost          `json:"data"`
	Paging  Paging              `json:"paging"`
	Summary PostCommentsSummary `json:"summary"`
}

type PostCommentsSummary struct {
	Order      string `json:"order"`
	TotalCount int64  `json:"total_count"`
	CanComment bool   `json:"can_comment"`
}

type PostComment struct {
	Id          string `json:"id"`
	CreatedTime string `json:"created_time"`
	Message     string `json:"message"`
}

type GetPostCommentsConfig struct {
	PostId  string
	After   *string
	Limit   *int64
	Summary bool
	Filter  *PostCommentsFilter
}

// GetPostComments returns Facebook post comments for a post
//
func (service *Service) GetPostCommentsRequest(config *GetPostCommentsConfig) (*go_http.RequestConfig, *PostCommentsResponse, *errortools.Error) {
	if config == nil {
		return nil, nil, errortools.ErrorMessage("GetAccountsConfig must not be a nil pointer")
	}

	values := url.Values{}
	limit := postCommentsLimit
	if config.Limit != nil {
		limit = *config.Limit
	}
	values.Set("limit", fmt.Sprintf("%v", limit))
	if config.After != nil {
		values.Set("after", *config.After)
	}
	values.Set("summary", fmt.Sprintf("%v", config.Summary))
	if config.Filter != nil {
		values.Set("filter", string(*config.Filter))
	}

	response := PostCommentsResponse{}
	relativeUrl := fmt.Sprintf("%s/comments?%s", config.PostId, values.Encode())
	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		RelativeUrl:   relativeUrl,
		Url:           service.urlV16(relativeUrl),
		ResponseModel: &response,
	}

	return &requestConfig, &response, nil
}

// GetPostComments returns Facebook post comments for a post
//
func (service *Service) GetPostComments(config *GetPostCommentsConfig) (*PostCommentsResponse, *errortools.Error) {
	requestConfig, response, e := service.GetPostCommentsRequest(config)
	if e != nil {
		return nil, e
	}
	_, _, e = service.httpRequest(requestConfig)
	if e != nil {
		return nil, e
	}

	return response, nil
}
