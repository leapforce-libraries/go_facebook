package facebook

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"net/url"
)

type ManagePage struct {
	AccessToken  string `json:"access_token"`
	Category     string `json:"category"`
	CategoryList []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"category_list"`
	Name  string    `json:"name"`
	Id    string    `json:"id"`
	Tasks *[]string `json:"tasks"`
}

type GetPagesConfig struct {
	UserId          string
	UserAccessToken string
}

func (service *Service) GetPages(config *GetPagesConfig) (*[]ManagePage, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetPagesConfig must not be a nil pointer")
	}

	var limit = 100

	var values = url.Values{}
	values.Set("access_token", config.UserAccessToken)
	values.Set("limit", fmt.Sprintf("%v", limit))

	var pages []ManagePage

	for {
		var response struct {
			Data   []ManagePage `json:"data"`
			Paging Paging       `json:"paging"`
		}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.urlV20(fmt.Sprintf("%s/accounts?%s", config.UserId, values.Encode())),
			ResponseModel: &response,
		}
		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		if len(response.Data) > 0 {
			pages = append(pages, response.Data...)
		}

		if response.Paging.Next == "" {
			break
		}
		values.Set("after", fmt.Sprintf("%v", response.Paging.Cursors.After))

	}

	return &pages, nil
}
