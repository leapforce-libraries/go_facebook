package facebook

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
)

type UserAccount struct {
	AccessToken  string `json:"access_token"`
	Category     string `json:"category"`
	CategoryList []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"category_list"`
	Name  string   `json:"name"`
	Id    string   `json:"id"`
	Tasks []string `json:"tasks"`
}

func (service *Service) UserAccounts(userId string) (*[]UserAccount, *errortools.Error) {

	var userAccounts []UserAccount
	url := service.urlV16(fmt.Sprintf("%s/accounts", userId))

	for {
		response := struct {
			Data   []UserAccount `json:"data"`
			Paging Paging        `json:"paging"`
		}{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           url,
			ResponseModel: &response,
		}
		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		userAccounts = append(userAccounts, response.Data...)

		url = response.Paging.Next
		if url == "" {
			break
		}
	}

	return &userAccounts, nil
}

type UserAdAccount struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func (service *Service) UserAdAccounts(userId string) (*[]UserAdAccount, *errortools.Error) {

	var userAccounts []UserAdAccount
	url := service.urlV16(fmt.Sprintf("%s/adaccounts", userId))

	for {
		response := struct {
			Data   []UserAdAccount `json:"data"`
			Paging Paging          `json:"paging"`
		}{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           url,
			ResponseModel: &response,
		}
		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		userAccounts = append(userAccounts, response.Data...)

		url = response.Paging.Next
		if url == "" {
			break
		}
	}

	return &userAccounts, nil
}
