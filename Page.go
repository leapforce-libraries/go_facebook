package facebook

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Page struct {
	Id                       string   `json:"id"`
	AccessToken              string   `json:"access_token"`
	Bio                      string   `json:"bio"`
	Category                 string   `json:"category"`
	Description              string   `json:"description"`
	Emails                   []string `json:"emails"`
	FanCount                 uint32   `json:"fan_count"`
	FollowersCount           uint32   `json:"followers_count"`
	InstagramBusinessAccount *struct {
		Id string `json:"id"`
	} `json:"instagram_business_account"`
	Name    string `json:"name"`
	Website string `json:"website"`
}

type PageField string

const (
	PageFieldId            PageField = "id"
	PageFieldAbout         PageField = "about"
	PageFieldAccessToken   PageField = "access_token"
	PageFieldAdCampaign    PageField = "ad_campaign"
	PageFieldAffiliation   PageField = "affiliation"
	PageFieldAppId         PageField = "app_id"
	PageFieldArtistsWeLike PageField = "artists_we_like"
	PageFieldAttire        PageField = "attire"
	PageFieldAwards        PageField = "awards"
	/*...*/
	PageFieldBio                      PageField = "bio"
	PageFieldCategory                 PageField = "category"
	PageFieldDescription              PageField = "description"
	PageFieldEmails                   PageField = "emails"
	PageFieldFanCount                 PageField = "fan_count"
	PageFieldFollowersCount           PageField = "followers_count"
	PageFieldInstagramBusinessAccount PageField = "instagram_business_account"
	PageFieldName                     PageField = "name"
	PageFieldWebsite                  PageField = "website"
)

type GetPageConfig struct {
	PageId          string
	Fields          *[]PageField
	PageAccessToken string
}

func (service *Service) GetPageRequest(config *GetPageConfig) (*go_http.RequestConfig, *Page, *errortools.Error) {
	if config == nil {
		return nil, nil, errortools.ErrorMessage("GetPageConfig must not be a nil pointer")
	}

	values := url.Values{}
	fields := []string{string(PageFieldId)}
	if config.Fields != nil {
		for _, field := range *config.Fields {
			if field == PageFieldId {
				continue
			}
			fields = append(fields, string(field))
		}
	}
	values.Set("fields", strings.Join(fields, ","))
	if config.PageAccessToken != "" {
		values.Set("access_token", config.PageAccessToken)
	}

	response := Page{}
	relativeUrl := fmt.Sprintf("%s?%s", config.PageId, values.Encode())
	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		RelativeUrl:   relativeUrl,
		Url:           service.urlV18(relativeUrl),
		ResponseModel: &response,
	}

	return &requestConfig, &response, nil
}

func (service *Service) GetPage(config *GetPageConfig) (*Page, *errortools.Error) {
	requestConfig, response, e := service.GetPageRequest(config)
	if e != nil {
		return nil, e
	}
	_, _, e = service.httpRequest(requestConfig)
	if e != nil {
		return nil, e
	}

	return response, nil
}
