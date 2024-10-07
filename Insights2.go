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

type Insight2sResponse struct {
	Data   []Insight2s `json:"data"`
	Paging *Paging     `json:"paging"`
}

type Insight2s struct {
	Name   string `json:"name"`
	Period string `json:"period"`
	Values []struct {
		Value   int    `json:"value"`
		EndTime string `json:"end_time"`
	} `json:"values"`
	Title                 string  `json:"title"`
	Description           string  `json:"description"`
	DescriptionFromApiDoc *string `json:"description_from_api_doc"`
	Id                    string  `json:"id"`
}

type Period string

const (
	PeriodDay            Period = "day"
	PeriodWeek           Period = "week"
	PeriodDays28         Period = "days_28"
	PeriodMonth          Period = "month"
	PeriodLifetime       Period = "lifetime"
	PeriodTotalOverRange Period = "total_over_range"
)

type GetInsight2sConfig struct {
	Id                        string
	DatePreset                *DatePreset
	Metrics                   []string
	Period                    *Period
	ShowDescriptionFromApiDoc *bool
	Since                     *time.Time
	Until                     *time.Time
}

func (service *Service) GetInsight2s(config *GetInsight2sConfig) (*[]Insight2s, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetInsight2sConfig must not be a nil pointer")
	}

	values := url.Values{}

	if config.DatePreset != nil {
		values.Set("date_preset", string(*config.DatePreset))
	}
	if len(config.Metrics) > 0 {
		values.Set("metric", strings.Join(config.Metrics, ","))
	}
	if config.Period != nil {
		values.Set("period", string(*config.Period))
	}
	if config.ShowDescriptionFromApiDoc != nil {
		values.Set("show_description_from_api_doc", fmt.Sprintf("%v", *config.ShowDescriptionFromApiDoc))
	}
	if config.Since != nil {
		values.Set("since", fmt.Sprintf("%v", config.Since.Unix()))
	}
	if config.Until != nil {
		values.Set("until", fmt.Sprintf("%v", config.Until.Unix()))
	}

	var insight2s []Insight2s

	url_ := service.urlV20(fmt.Sprintf("%s/insights?%s", config.Id, values.Encode()))

	for {
		insight2sResponse := Insight2sResponse{}
		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           url_,
			ResponseModel: &insight2sResponse,
		}
		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		insight2s = append(insight2s, insight2sResponse.Data...)

		if insight2sResponse.Paging == nil {
			break
		}

		if insight2sResponse.Paging.Next == "" {
			break
		}

		url_ = insight2sResponse.Paging.Next
	}

	return &insight2s, nil
}
