package facebook

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"net/url"
)

type OAuthService struct {
	clientId      string
	clientSecret  string
	redirectUri   string
	httpService   *go_http.Service
	errorResponse ErrorResponse
}

type OAuthServiceConfig struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string
}

func NewOAuthService(serviceConfig *OAuthServiceConfig) (*OAuthService, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	return &OAuthService{
		clientId:     serviceConfig.ClientId,
		clientSecret: serviceConfig.ClientSecret,
		redirectUri:  serviceConfig.RedirectUri,
		httpService:  httpService,
	}, nil
}

func (service *OAuthService) Error() Error {
	return service.errorResponse.Error
}

func (service *OAuthService) LoginUrl(scopes string) string {
	return fmt.Sprintf("%s/dialog/oauth?client_id=%s&redirect_uri=%s&scope=%s&response_type=code", apiUrlWww, service.clientId, service.redirectUri, scopes)
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   *int64 `json:"expires_in"`
}

func (service *OAuthService) TokenFromCode(code string) (*Token, *errortools.Error) {
	values := url.Values{}
	values.Set("client_id", service.clientId)
	values.Set("redirect_uri", service.redirectUri)
	values.Set("client_secret", service.clientSecret)
	values.Set("code", code)

	var token Token

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           fmt.Sprintf("%s/oauth/access_token?%s", apiUrlV21, values.Encode()),
		ResponseModel: &token,
	}
	_, _, e := service.httpService.HttpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &token, nil
}

func (service *OAuthService) ExchangeToken(accessToken string) (*Token, *errortools.Error) {
	values := url.Values{}
	values.Set("grant_type", "fb_exchange_token")
	values.Set("client_id", service.clientId)
	values.Set("client_secret", service.clientSecret)
	values.Set("fb_exchange_token", accessToken)

	var token Token

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           fmt.Sprintf("%s/oauth/access_token?%s", apiUrlV21, values.Encode()),
		ResponseModel: &token,
	}
	_, _, e := service.httpService.HttpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &token, nil
}

type InspectedToken struct {
	Data struct {
		AppId               string   `json:"app_id"`
		Type                string   `json:"type"`
		Application         string   `json:"application"`
		DataAccessExpiresAt int64    `json:"data_access_expires_at"`
		ExpiresAt           int64    `json:"expires_at"`
		IsValid             bool     `json:"is_valid"`
		IssuedAt            int64    `json:"issued_at"`
		Scopes              []string `json:"scopes"`
		GranularScopes      []struct {
			Scope string `json:"scope"`
		} `json:"granular_scopes"`
		UserId string `json:"user_id"`
	} `json:"data"`
}
