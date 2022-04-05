package facebook

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type IgUser struct {
	Biography         *string `json:"caption"`
	Id                string  `json:"id"`
	IgId              *int64  `json:"ig_id"`
	FollowersCount    *int64  `json:"followers_count"`
	FollowsCount      *int64  `json:"follows_count"`
	MediaCount        *int64  `json:"media_count"`
	Name              *string `json:"name"`
	ProfilePictureUrl *string `json:"profile_picture_url"`
	Username          *string `json:"username"`
	Website           *string `json:"website"`
}

type IgUserField string

const (
	IgUserFieldBiography         IgUserField = "biography"
	IgUserFieldId                IgUserField = "id"
	IgUserFieldIgId              IgUserField = "ig_id"
	IgUserFieldFollowersCount    IgUserField = "followers_count"
	IgUserFieldFollowsCount      IgUserField = "follows_count"
	IgUserFieldMediaCount        IgUserField = "media_count"
	IgUserFieldName              IgUserField = "name"
	IgUserFieldProfilePictureUrl IgUserField = "profile_picture_url"
	IgUserFieldUsername          IgUserField = "username"
	IgUserFieldWebsite           IgUserField = "website"
)

type GetIgUserConfig struct {
	IgUserId string
	Fields   *[]IgUserField
}

func (service *Service) GetIgUser(config *GetIgUserConfig) (*IgUser, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetIgUserConfig must not be a nil pointer")
	}

	values := url.Values{}

	fields := []string{string(IgUserFieldId)}
	if config.Fields != nil {
		for _, field := range *config.Fields {
			if field == IgUserFieldId {
				continue
			}
			fields = append(fields, string(field))
		}
	}
	values.Set("fields", strings.Join(fields, ","))

	url := service.url(fmt.Sprintf("%s?%s", config.IgUserId, values.Encode()))

	igUser := IgUser{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           url,
		ResponseModel: &igUser,
	}

	fmt.Println(requestConfig.Url)

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &igUser, nil
}
