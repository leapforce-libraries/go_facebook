package facebook

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"net/url"
)

func (service *Service) InspectToken(accessToken string) (*InspectedToken, *errortools.Error) {
	values := url.Values{}
	values.Set("input_token", accessToken)

	var inspectedToken InspectedToken

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.urlV18(fmt.Sprintf("debug_token?%s", values.Encode())),
		ResponseModel: &inspectedToken,
	}
	_, _, e := service.httpService.HttpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &inspectedToken, nil
}
