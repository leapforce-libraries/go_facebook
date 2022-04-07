package facebook

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type IgMediaInsightsResponse struct {
	Data   []IgMediaInsight `json:"data"`
	Paging Paging           `json:"paging"`
}

type IgMediaInsightValue struct {
	Value int64 `json:"value"`
}

type IgMediaInsight struct {
	Name        IgMediaInsightsMetric `json:"name"`
	Period      IgMediaInsightsPeriod `json:"period"`
	Values      []IgMediaInsightValue `json:"values"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Id          string                `json:"id"`
}

type IgMediaInsightsMetric string

const (
	IgMediaInsightsMetricEngagement  IgMediaInsightsMetric = "engagement"
	IgMediaInsightsMetricImpressions IgMediaInsightsMetric = "impressions"
	IgMediaInsightsMetricReach       IgMediaInsightsMetric = "reach"
	IgMediaInsightsMetricSaved       IgMediaInsightsMetric = "saved"
	IgMediaInsightsMetricVideoViews  IgMediaInsightsMetric = "video_views"
)

type IgMediaInsightsPeriod string

const (
	IgMediaInsightsPeriodLifetime IgMediaInsightsPeriod = "lifetime"
)

type GetIgMediaInsightsConfig struct {
	MediaId         string
	UserAccessToken *string
	Metrics         []IgMediaInsightsMetric
}

func (service *Service) GetIgMediaInsights(config *GetIgMediaInsightsConfig) (*[]IgMediaInsight, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetIgMediaInsightsConfig must not be a nil pointer")
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

	url := service.url(fmt.Sprintf("%s/insights?%s", config.MediaId, values.Encode()))

	insights := []IgMediaInsight{}

	for {
		insightsResponse := IgMediaInsightsResponse{}

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
