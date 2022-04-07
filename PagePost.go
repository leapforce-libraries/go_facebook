package facebook

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type GetPagePostConfig struct {
	PostId string
	Fields *[]PagePublishedPostField
}

// GetPagePosts returns Facebook post comments for a post
//
func (service *Service) GetPagePostRequest(config *GetPagePostConfig) (*go_http.RequestConfig, *PagePublishedPost, *errortools.Error) {
	if config == nil {
		return nil, nil, errortools.ErrorMessage("GetAccountsConfig must not be a nil pointer")
	}

	values := url.Values{}
	fields := []string{string(PagePublishedPostFieldId)}
	if config.Fields != nil {
		for _, field := range *config.Fields {
			if field == PagePublishedPostFieldId {
				continue
			}
			fields = append(fields, string(field))
		}
	}
	values.Set("fields", strings.Join(fields, ","))

	response := PagePublishedPost{}
	relativeUrl := fmt.Sprintf("%s?%s", config.PostId, values.Encode())
	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		RelativeUrl:   relativeUrl,
		Url:           service.url(relativeUrl),
		ResponseModel: &response,
	}

	return &requestConfig, &response, nil
}

// GetPagePosts returns Facebook post comments for a post
//
func (service *Service) GetPagePost(config *GetPagePostConfig) (*PagePublishedPost, *errortools.Error) {
	requestConfig, response, e := service.GetPagePostRequest(config)
	if e != nil {
		return nil, e
	}
	_, _, e = service.httpRequest(requestConfig)
	if e != nil {
		return nil, e
	}

	return response, nil
}