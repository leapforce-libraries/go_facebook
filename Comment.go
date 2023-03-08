package facebook

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"net/url"
)

type CreateCommentConfig struct {
	ObjectId           string
	AttachmentId       *string
	AttachmentShareUrl *string
	AttachmentUrl      *string
	Source             *string
	Message            *string
	PageAccessToken    string
}

func (service *Service) CreateComment(config *CreateCommentConfig) (string, *errortools.Error) {
	if config == nil {
		return "", errortools.ErrorMessage("CreateCommentConfig must not be a nil pointer")
	}

	var values = url.Values{}
	values.Set("access_token", config.PageAccessToken)
	if config.AttachmentId != nil {
		values.Set("attachment_id", *config.AttachmentId)
	}
	if config.AttachmentShareUrl != nil {
		values.Set("attachment_share_url", *config.AttachmentShareUrl)
	}
	if config.AttachmentUrl != nil {
		values.Set("attachment_url", *config.AttachmentUrl)
	}
	if config.Message != nil {
		values.Set("message", *config.Message)
	}
	if config.Source != nil {
		values.Set("source", *config.Source)
	}

	var response struct {
		Id string `json:"id"`
	}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.urlV16(fmt.Sprintf("%s/comments?%s", config.ObjectId, values.Encode())),
		ResponseModel: &response,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return "", e
	}

	return response.Id, nil
}
