package facebook

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"net/url"
)

type CreateIgCommentConfig struct {
	MediaId string
	Message string
}

func (service *Service) CreateIgComment(config *CreateIgCommentConfig) (string, *errortools.Error) {
	if config == nil {
		return "", errortools.ErrorMessage("CreateIgCommentConfig must not be a nil pointer")
	}

	var values = url.Values{}
	values.Set("message", config.Message)

	var response struct {
		Id string `json:"id"`
	}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.urlV18(fmt.Sprintf("%s/comments?%s", config.MediaId, values.Encode())),
		ResponseModel: &response,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return "", e
	}

	return response.Id, nil
}
