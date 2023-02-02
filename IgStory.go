package facebook

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type IgStoryResponse struct {
	Data    []IgMedia           `json:"data"`
	Paging  Paging              `json:"paging"`
	Summary PostCommentsSummary `json:"summary"`
}

type GetIgUserStoriesConfig struct {
	IgUserId string
	Fields   *[]IgMediaField
}

func (service *Service) GetIgUserStories(config *GetIgUserStoriesConfig) (*[]IgMedia, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetIgUserStoriesConfig must not be a nil pointer")
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

	url := service.urlV16(fmt.Sprintf("%s/stories?%s", config.IgUserId, values.Encode()))

	igUserStories := []IgMedia{}

	for {
		response := IgStoryResponse{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           url,
			ResponseModel: &response,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		igUserStories = append(igUserStories, response.Data...)

		if response.Paging.Next == "" {
			break
		}

		url = response.Paging.Next
	}

	return &igUserStories, nil
}
