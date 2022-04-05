package facebook

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type IgUserMediaResponse struct {
	Data    []IgMedia           `json:"data"`
	Paging  Paging              `json:"paging"`
	Summary PostCommentsSummary `json:"summary"`
}

type GetIgUserMediasConfig struct {
	IgUserId string
	Fields   *[]IgMediaField
	Since    *time.Time
	Until    *time.Time
}

func (service *Service) GetIgUserMedias(config *GetIgUserMediasConfig) (*[]IgMedia, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetIgUserMediasConfig must not be a nil pointer")
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

	if config.Since != nil {
		values.Set("since", fmt.Sprintf("%v", config.Since.Unix()))
	}

	if config.Until != nil {
		values.Set("until", fmt.Sprintf("%v", config.Until.Unix()))
	}

	url := service.url(fmt.Sprintf("%s/media?%s", config.IgUserId, values.Encode()))

	igUserMedias := []IgMedia{}

	for {
		response := IgUserMediaResponse{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           url,
			ResponseModel: &response,
		}

		fmt.Println(requestConfig.Url)

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		igUserMedias = append(igUserMedias, response.Data...)

		if response.Paging.Next == "" {
			break
		}

		url = response.Paging.Next
	}

	return &igUserMedias, nil
}
