package facebook

/*
import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type InsightsResponse struct {
	Data   []Insights `json:"data"`
	Paging *Paging    `json:"paging"`
}

type Insights struct {
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

type DatePreset string

const (
	DatePresetToday            DatePreset = "today"
	DatePresetYesterday        DatePreset = "yesterday"
	DatePresetThisMonth        DatePreset = "this_month"
	DatePresetLastMonth        DatePreset = "last_month"
	DatePresetThisQuarter      DatePreset = "this_quarter"
	DatePresetMaximum          DatePreset = "maximum"
	DatePresetLast3Days        DatePreset = "last_3d"
	DatePresetLast7Days        DatePreset = "last_7d"
	DatePresetLast14Days       DatePreset = "last_14d"
	DatePresetLast28Days       DatePreset = "last_28d"
	DatePresetLast30Days       DatePreset = "last_30d"
	DatePresetLast90Days       DatePreset = "last_90d"
	DatePresetLastWeekMonSun   DatePreset = "last_week_mon_sun"
	DatePresetLastWeekSunSat   DatePreset = "last_week_sun_sat"
	DatePresetLastQuarter      DatePreset = "last_quarter"
	DatePresetLastYear         DatePreset = "last_year"
	DatePresetThisWeekMonToday DatePreset = "this_week_mon_today"
	DatePresetThisWeekSunToday DatePreset = "this_week_sun_today"
	DatePresetThisYear         DatePreset = "this_year"
)

type Period string

const (
	PeriodDay            Period = "day"
	PeriodWeek           Period = "week"
	PeriodDays28         Period = "days_28"
	PeriodMonth          Period = "month"
	PeriodLifetime       Period = "lifetime"
	PeriodTotalOverRange Period = "total_over_range"
)

type GetInsightsConfig struct {
	Id                        string
	DatePreset                *DatePreset
	Metrics                   []string
	Period                    *Period
	ShowDescriptionFromApiDoc *bool
	Since                     *time.Time
	Until                     *time.Time
}

func (service *Service) GetInsights(config *GetInsightsConfig) (*[]Insights, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetInsightsConfig must not be a nil pointer")
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

	var insights []Insights

	url_ := service.urlV20(fmt.Sprintf("%s/insights?%s", config.Id, values.Encode()))

	for {
		insightsResponse := InsightsResponse{}
		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           url_,
			ResponseModel: &insightsResponse,
		}
		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		insights = append(insights, insightsResponse.Data...)

		if insightsResponse.Paging == nil {
			break
		}

		if insightsResponse.Paging.Next == "" {
			break
		}

		url_ = insightsResponse.Paging.Next
	}

	return &insights, nil
}
*/
