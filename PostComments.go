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
	ID          string `json:"id"`
	CreatedTime string `json:"created_time"`
	Message     string `json:"message"`
}

type GetPostCommentsConfig struct {
	PostID      string
	After       *string
	Limit       *int64
	Summary     bool
	AccessToken *string
	Filter      *PostCommentsFilter
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
	if config.AccessToken != nil {
		values.Set("access_token", *config.AccessToken)
	}
	if config.Filter != nil {
		values.Set("filter", string(*config.Filter))
	}

	response := PostCommentsResponse{}
	relativeURL := fmt.Sprintf("%s/comments?%s", config.PostID, values.Encode())
	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		RelativeURL:   relativeURL,
		URL:           service.url(relativeURL),
		ResponseModel: &response,
	}

	return &requestConfig, &response, nil
}

// GetPostComments returns Facebook post comments for a post
//
func (service *Service) GetPostComments(config *GetPostCommentsConfig) (*PostCommentsResponse, *errortools.Error) {
	requestConfig, response, e := service.GetPostCommentsRequest(config)
	fmt.Println(requestConfig.URL)
	if e != nil {
		return nil, e
	}
	_, _, e = service.httpRequest(requestConfig)
	if e != nil {
		return nil, e
	}

	return response, nil
}

/*
// PostCommentsCount returns Facebook post comments count for a post
//
func (service *Service) PostCommentsCount(postID string, accessToken string, filter *PostCommentsFilter) (*int64, *errortools.Error) {
	path := fmt.Sprintf("/%s/comments", postID)

	params := fb2.Params{
		"limit":        0,
		"access_token": accessToken,
		"summary":      true,
	}

	if filter != nil {
		params["filter"] = *filter
	}

	result, e := api.GetWithRetry(service.session, path, params)
	if e != nil {
		return nil, e
	}

	response := PostCommentsResponse{}
	err := result.DecodeField("", &response)
	//err = mapstructure.Decode(result, &response)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	return &response.Summary.TotalCount, nil
}
*/
