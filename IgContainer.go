package facebook

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"net/url"
	"strings"
)

type IgContainer struct {
	Id         string  `json:"id"`
	Status     *string `json:"status"`
	StatusCode *string `json:"status_code"`
}

type IgContainerField string

const (
	IgContainerFieldId         IgContainerField = "id"
	IgContainerFieldStatus     IgContainerField = "status"
	IgContainerFieldStatusCode IgContainerField = "status_code"
)

type GetIgContainerConfig struct {
	IgContainerId string
	Fields        *[]IgContainerField
}

func (service *Service) GetIgContainer(config *GetIgContainerConfig) (*IgContainer, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetIgContainersConfig must not be a nil pointer")
	}

	values := url.Values{}

	fields := []string{string(IgContainerFieldId)}
	if config.Fields != nil {
		for _, field := range *config.Fields {
			if field == IgContainerFieldId {
				continue
			}
			fields = append(fields, string(field))
		}
	}
	values.Set("fields", strings.Join(fields, ","))

	url := service.urlV18(fmt.Sprintf("%s?%s", config.IgContainerId, values.Encode()))

	igContainer := IgContainer{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           url,
		ResponseModel: &igContainer,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &igContainer, nil
}
