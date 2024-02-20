package facebook

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	f_types "github.com/leapforce-libraries/go_facebook/types"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type LeadResponse struct {
	Data   []Lead  `json:"data"`
	Paging *Paging `json:"paging"`
}

type Lead struct {
	CreatedTime *f_types.DateTimeString `json:"created_time"`
	Id          go_types.Int64String    `json:"id"`
	AdId        *go_types.Int64String   `json:"ad_id"`
	FormId      *go_types.Int64String   `json:"form_id"`
	FieldData   []FieldData             `json:"field_data"`
}

type FieldData struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type LeadField string

const (
	LeadFieldId          LeadField = "id"
	LeadFieldCreatedTime LeadField = "created_time"
	LeadFieldAdId        LeadField = "ad_id"
	LeadFieldFormId      LeadField = "form_id"
	LeadFieldFieldData   LeadField = "field_data"
)

type GetLeadsConfig struct {
	AdId   int64
	Fields []LeadField
}

func (service *Service) GetLeads(config *GetLeadsConfig) (*[]Lead, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetLeadsConfig must not be a nil pointer")
	}

	values := url.Values{}
	fields := []string{}
	if len(config.Fields) == 0 {
		fields = append(fields, string(LeadFieldId))
	} else {
		for _, field := range config.Fields {
			fields = append(fields, string(field))
		}
	}
	values.Set("fields", strings.Join(fields, ","))

	leads := []Lead{}

	url := service.urlV18(fmt.Sprintf("%v/leads?%s", config.AdId, values.Encode()))

	for {
		leadResponse := LeadResponse{}
		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           url,
			ResponseModel: &leadResponse,
		}
		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		leads = append(leads, leadResponse.Data...)

		if leadResponse.Paging == nil {
			break
		}

		if leadResponse.Paging.Next == "" {
			break
		}

		url = leadResponse.Paging.Next
	}

	return &leads, nil
}
