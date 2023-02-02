package facebook

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	f_types "github.com/leapforce-libraries/go_facebook/types"
	go_http "github.com/leapforce-libraries/go_http"
)

type IgUserInsightsResponse struct {
	Data   []IgUserInsight `json:"data"`
	Paging Paging          `json:"paging"`
}

type IgUserInsightValue struct {
	Value   int64                  `json:"value"`
	EndTime f_types.DateTimeString `json:"end_time"`
}

type IgUserInsight struct {
	Name        IgUserInsightsMetric `json:"name"`
	Period      IgUserInsightsPeriod `json:"period"`
	Values      []IgUserInsightValue `json:"values"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Id          string               `json:"id"`
}

type IgUserInsightsMetric string

const (
	IgUserInsightsMetricAudienceCity        IgUserInsightsMetric = "audience_city"
	IgUserInsightsMetricAudienceCountry     IgUserInsightsMetric = "audience_country"
	IgUserInsightsMetricAudienceGenderAge   IgUserInsightsMetric = "audience_gender_age"
	IgUserInsightsMetricAudienceLocale      IgUserInsightsMetric = "audience_locale"
	IgUserInsightsMetricEmailContacts       IgUserInsightsMetric = "email_contacts"
	IgUserInsightsMetricFollowerCount       IgUserInsightsMetric = "follower_count"
	IgUserInsightsMetricGetDirectionsClicks IgUserInsightsMetric = "get_directions_clicks"
	IgUserInsightsMetricImpressions         IgUserInsightsMetric = "impressions"
	IgUserInsightsMetricOnlineFollowers     IgUserInsightsMetric = "online_followers"
	IgUserInsightsMetricPhoneCallClicks     IgUserInsightsMetric = "phone_call_clicks"
	IgUserInsightsMetricProfileViews        IgUserInsightsMetric = "profile_views"
	IgUserInsightsMetricReach               IgUserInsightsMetric = "reach"
	IgUserInsightsMetricTextMessageClicks   IgUserInsightsMetric = "text_message_clicks"
	IgUserInsightsMetricWebsiteClicks       IgUserInsightsMetric = "website_clicks"
)

type IgUserInsightsPeriod string

const (
	IgUserInsightsPeriodDay      IgUserInsightsPeriod = "day"
	IgUserInsightsPeriodWeek     IgUserInsightsPeriod = "week"
	IgUserInsightsPeriodDays28   IgUserInsightsPeriod = "days_28"
	IgUserInsightsPeriodLifetime IgUserInsightsPeriod = "lifetime"
)

type GetIgUserInsightsConfig struct {
	UserId          string
	UserAccessToken *string
	Metrics         []IgUserInsightsMetric
	Period          IgUserInsightsPeriod
	Since           time.Time
	Until           time.Time
}

func (service *Service) GetIgUserInsights(config *GetIgUserInsightsConfig) (*[]IgUserInsight, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetIgUserInsightsConfig must not be a nil pointer")
	}

	values := url.Values{}
	if config.UserAccessToken != nil {
		values.Set("access-token", *config.UserAccessToken)
	}

	metrics := []string{}
	for _, m := range config.Metrics {
		metrics = append(metrics, string(m))
	}
	values.Set("metric", strings.Join(metrics, ","))
	values.Set("period", string(config.Period))
	values.Set("since", fmt.Sprintf("%v", config.Since.Unix()))
	values.Set("until", fmt.Sprintf("%v", config.Until.Unix()))

	url := service.urlV16(fmt.Sprintf("%s/insights?%s", config.UserId, values.Encode()))

	insights := []IgUserInsight{}

	for {
		insightsResponse := IgUserInsightsResponse{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           url,
			ResponseModel: &insightsResponse,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		insights = append(insights, insightsResponse.Data...)

		if insightsResponse.Paging.Next == "" {
			break
		}

		url = insightsResponse.Paging.Next
	}

	return &insights, nil
}
