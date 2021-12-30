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
	ID             string   `json:"id"`
	AccessToken    string   `json:"access_token"`
	Bio            string   `json:"bio"`
	Category       string   `json:"category"`
	Description    string   `json:"description"`
	Emails         []string `json:"emails"`
	FanCount       uint32   `json:"fan_count"`
	FollowersCount uint32   `json:"followers_count"`
	Name           string   `json:"name"`
	Website        string   `json:"website"`
}

type PageField string

const (
	PageFieldID            PageField = "id"
	PageFieldAbout         PageField = "about"
	PageFieldAccessToken   PageField = "access_token"
	PageFieldAdCampaign    PageField = "ad_campaign"
	PageFieldAffiliation   PageField = "affiliation"
	PageFieldAppID         PageField = "app_id"
	PageFieldArtistsWeLike PageField = "artists_we_like"
	PageFieldAttire        PageField = "attire"
	PageFieldAwards        PageField = "awards"
	/*...*/
	PageFieldBio            PageField = "bio"
	PageFieldCategory       PageField = "category"
	PageFieldDescription    PageField = "description"
	PageFieldEmails         PageField = "emails"
	PageFieldFanCount       PageField = "fan_count"
	PageFieldFollowersCount PageField = "followers_count"
	PageFieldName           PageField = "name"
	PageFieldWebsite        PageField = "website"
)

type GetPageConfig struct {
	PageID string
	Fields *[]PageField
}

// GetPages returns Facebook post comments for a post
//
func (service *Service) GetPageRequest(config *GetPageConfig) (*go_http.RequestConfig, *Page, *errortools.Error) {
	if config == nil {
		return nil, nil, errortools.ErrorMessage("GetPageConfig must not be a nil pointer")
	}

	values := url.Values{}
	fields := []string{string(PageFieldID)}
	if config.Fields != nil {
		for _, field := range *config.Fields {
			if field == PageFieldID {
				continue
			}
			fields = append(fields, string(field))
		}
	}
	values.Set("fields", strings.Join(fields, ","))

	response := Page{}
	relativeURL := fmt.Sprintf("%s?%s", config.PageID, values.Encode())
	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		RelativeURL:   relativeURL,
		URL:           service.url(relativeURL),
		ResponseModel: &response,
	}

	return &requestConfig, &response, nil
}

// GetPages returns Facebook post comments for a post
//
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
