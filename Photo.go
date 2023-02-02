package facebook

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"net/url"
)

type CreatePhotoConfig struct {
	PageId          string
	Caption         string
	Url             string
	Published       *bool
	PageAccessToken string
}

func (service *Service) CreatePhoto(config *CreatePhotoConfig) (string, string, *errortools.Error) {
	if config == nil {
		return "", "", errortools.ErrorMessage("CreatePhotoConfig must not be a nil pointer")
	}

	var values = url.Values{}
	values.Set("access_token", config.PageAccessToken)
	values.Set("caption", config.Caption)
	values.Set("url", config.Url)
	if config.Published != nil {
		values.Set("published", fmt.Sprintf("%v", *config.Published))
	}

	var response struct {
		Id     string `json:"id"`
		PostId string `json:"post_id"`
	}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url(fmt.Sprintf("%s/photos?%s", config.PageId, values.Encode())),
		ResponseModel: &response,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return "", "", e
	}

	return response.Id, response.PostId, nil
}
