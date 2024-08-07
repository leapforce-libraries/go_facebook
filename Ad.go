package facebook

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	f_types "github.com/leapforce-libraries/go_facebook/types"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type AdResponse struct {
	Data   []Ad    `json:"data"`
	Paging *Paging `json:"paging"`
}

type Ad struct {
	ID                   go_types.Int64String    `json:"id"`
	AccountID            *go_types.Int64String   `json:"account_id"`
	AdReviewFeedback     json.RawMessage         `json:"ad_review_feedback"`
	AdLabels             json.RawMessage         `json:"adlabels"`
	AdSet                *AdSet                  `json:"adset"`
	AdSetID              *go_types.Int64String   `json:"adset_id"`
	BidAmount            *uint32                 `json:"bid_amount"`
	Campaign             *Campaign               `json:"campaign"`
	CampaignID           *go_types.Int64String   `json:"campaign_id"`
	ConfiguredStatus     *string                 `json:"configured_status"`
	ConversionDomain     json.RawMessage         `json:"conversion_domain"`
	CreatedTime          *f_types.DateTimeString `json:"created_time"`
	Creative             *Creative               `json:"creative"`
	EffectiveStatus      *string                 `json:"effective_status"`
	IssuesInfo           json.RawMessage         `json:"issues_info"`
	LastUpdatedByAppID   *int32                  `json:"last_updated_by_app_id"`
	Name                 *string                 `json:"name"`
	PreviewShareableLink *string                 `json:"preview_shareable_link"`
	Recommendations      json.RawMessage         `json:"recommendations"`
	SourceAd             *Ad                     `json:"source_ad"`
	SourceAdID           *go_types.Int64String   `json:"source_ad_id"`
	Status               *string                 `json:"status"`
	TrackingSpecs        json.RawMessage         `json:"tracking_specs"`
	UpdatedTime          *f_types.DateTimeString `json:"updated_time"`
}

type AdField string

const (
	AdFieldID                   AdField = "id"
	AdFieldAccountID            AdField = "account_id"
	AdFieldAdReviewFeedback     AdField = "ad_review_feedback"
	AdFieldAdLabels             AdField = "adlabels"
	AdFieldAdSet                AdField = "adset"
	AdFieldAdSetID              AdField = "adset_id"
	AdFieldBidAmount            AdField = "bid_amount"
	AdFieldCampaign             AdField = "campaign"
	AdFieldCampaignID           AdField = "campaign_id"
	AdFieldConfiguredStatus     AdField = "configured_status"
	AdFieldConversionDomain     AdField = "conversion_domain"
	AdFieldCreatedTime          AdField = "created_time"
	AdFieldCreative             AdField = "creative"
	AdFieldEffectiveStatus      AdField = "effective_status"
	AdFieldIssuesInfo           AdField = "issues_info"
	AdFieldLastUpdatedByAppID   AdField = "last_updated_by_app_id"
	AdFieldName                 AdField = "name"
	AdFieldPreviewShareableLink AdField = "preview_shareable_link"
	AdFieldRecommendations      AdField = "recommendations"
	AdFieldSourceAd             AdField = "source_ad"
	AdFieldSourceAdID           AdField = "source_ad_id"
	AdFieldStatus               AdField = "status"
	AdFieldTrackingSpecs        AdField = "tracking_specs"
	AdFieldUpdatedTime          AdField = "updated_time"
)

type GetAdsConfig struct {
	AccountID int64
	Since     *time.Time
	Fields    []AdField
	Limit     *uint64
}

func (service *Service) GetAds(config *GetAdsConfig) (*[]Ad, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetAdsConfig must not be a nil pointer")
	}

	values := url.Values{}

	fields := []string{}
	if len(config.Fields) == 0 {
		fields = append(fields, string(AdFieldID))
	} else {
		for _, field := range config.Fields {
			fields = append(fields, string(field))
		}
	}
	values.Set("fields", strings.Join(fields, ","))

	limit := limitDefault
	if config.Limit != nil {
		limit = *config.Limit
	}
	values.Set("limit", fmt.Sprintf("%v", limit))

	ads := []Ad{}

	url := service.urlV20(fmt.Sprintf("act_%v/ads?%s", config.AccountID, values.Encode()))

	for {
		adResponse := AdResponse{}
		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           url,
			ResponseModel: &adResponse,
		}
		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		for _, ad := range adResponse.Data {
			if ad.CreatedTime == nil {
				continue
			}

			if config.Since != nil {
				if ad.CreatedTime.Value().Before(*config.Since) {
					continue
				}
			}

			ads = append(ads, ad)
		}

		if adResponse.Paging == nil {
			break
		}

		if adResponse.Paging.Next == "" {
			break
		}

		url = adResponse.Paging.Next
	}

	return &ads, nil
}
