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
		Url:           service.urlV20(relativeUrl),
		ResponseModel: &response,
	}

	return &requestConfig, &response, nil
}

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

type CreatePagePostConfig struct {
	PageId          string
	Message         string
	Link            *string
	MediaIds        *[]string
	PageAccessToken string
}

func (service *Service) CreatePagePost(config *CreatePagePostConfig) (string, *errortools.Error) {
	if config == nil {
		return "", errortools.ErrorMessage("CreatePagePostConfig must not be a nil pointer")
	}

	var values = url.Values{}
	values.Set("access_token", config.PageAccessToken)
	values.Set("message", config.Message)

	if config.Link != nil {
		values.Set("link", *config.Link)
	}

	if config.MediaIds != nil {
		for i, mediaId := range *config.MediaIds {
			values.Set(fmt.Sprintf("attached_media[%v]", i), fmt.Sprintf("{\"media_fbid\":\"%s\"}", mediaId))
		}
	}

	var response struct {
		Id string `json:"id"`
	}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url(fmt.Sprintf("%s/feed?%s", config.PageId, values.Encode())),
		ResponseModel: &response,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return "", e
	}

	return response.Id, nil
}
