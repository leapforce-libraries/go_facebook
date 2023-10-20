package facebook

import (
	"encoding/json"
	"fmt"
	f_types "github.com/leapforce-libraries/go_facebook/types"
	"net/http"
	"net/url"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type CustomConversionResponse struct {
	Data   []CustomConversion `json:"data"`
	Paging *Paging            `json:"paging"`
}

type CustomConversion struct {
	Id                       go_types.Int64String    `json:"id"`
	AccountId                go_types.Int64String    `json:"account_id"`
	Business                 *Business               `json:"business"`
	CreationTime             *f_types.DateTimeString `json:"creation_time"`
	CustomEventType          *string                 `json:"custom_event_type"`
	DataSources              []ConversionDataSource  `json:"data_sources"`
	DefaultConversionValue   *int64                  `json:"default_conversion_value"`
	Description              *string                 `json:"description"`
	EventSourceType          *string                 `json:"event_source_type"`
	FirstFiredTime           *f_types.DateTimeString `json:"first_fired_time"`
	IsArchived               *bool                   `json:"is_archived"`
	IsUnavailable            *bool                   `json:"is_unavailable"`
	LastFiredTime            *f_types.DateTimeString `json:"last_fired_time"`
	Name                     *string                 `json:"name"`
	OfflineConversionDataSet json.RawMessage         `json:"offline_conversion_data_set"`
	Pixel                    *Pixel                  `json:"pixel"`
	RetentionDays            *int64                  `json:"retention_days"`
	Rule                     *string                 `json:"rule"`
}

type Business struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Pixel struct {
	Id string `json:"id"`
}

type ConversionDataSource struct {
	Id         string `json:"id"`
	SourceType string `json:"source_type"`
	Name       string `json:"name"`
}

type CustomConversionField string

const (
	CustomConversionFieldId                       CustomConversionField = "id"
	CustomConversionFieldAccountId                CustomConversionField = "account_id"
	CustomConversionFieldBusiness                 CustomConversionField = "business"
	CustomConversionFieldCreationTime             CustomConversionField = "creation_time"
	CustomConversionFieldCustomEventType          CustomConversionField = "custom_event_type"
	CustomConversionFieldDataSources              CustomConversionField = "data_sources"
	CustomConversionFieldDefaultConversionValue   CustomConversionField = "default_conversion_value"
	CustomConversionFieldDescription              CustomConversionField = "description"
	CustomConversionFieldEventSourceType          CustomConversionField = "event_source_type"
	CustomConversionFieldFirstFiredTime           CustomConversionField = "first_fired_time"
	CustomConversionFieldIsArchived               CustomConversionField = "is_archived"
	CustomConversionFieldIsUnavailable            CustomConversionField = "is_unavailable"
	CustomConversionFieldLastFiredTime            CustomConversionField = "last_fired_time"
	CustomConversionFieldName                     CustomConversionField = "name"
	CustomConversionFieldOfflineConversionDataSet CustomConversionField = "offline_conversion_data_set"
	CustomConversionFieldPixel                    CustomConversionField = "pixel"
	CustomConversionFieldRetentionDays            CustomConversionField = "retention_days"
	CustomConversionFieldRule                     CustomConversionField = "rule"
)

type GetCustomConversionsConfig struct {
	AccountId int64
	Fields    []CustomConversionField
}

func (service *Service) GetCustomConversions(config *GetCustomConversionsConfig) (*[]CustomConversion, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetCustomConversionsConfig must not be a nil pointer")
	}

	values := url.Values{}
	fields := []string{}
	if len(config.Fields) == 0 {
		fields = append(fields, string(CustomConversionFieldId))
	} else {
		for _, field := range config.Fields {
			fields = append(fields, string(field))
		}
	}
	values.Set("fields", strings.Join(fields, ","))

	customConversions := []CustomConversion{}

	url := service.urlV16(fmt.Sprintf("act_%v/customconversions?%s", config.AccountId, values.Encode()))

	for {
		customConversionResponse := CustomConversionResponse{}
		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           url,
			ResponseModel: &customConversionResponse,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		customConversions = append(customConversions, customConversionResponse.Data...)

		if customConversionResponse.Paging == nil {
			break
		}

		if customConversionResponse.Paging.Next == "" {
			break
		}

		url = customConversionResponse.Paging.Next
	}

	return &customConversions, nil
}
