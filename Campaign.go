package facebook

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	f_types "github.com/leapforce-libraries/go_facebook/types"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type CampaignResponse struct {
	Data   []Campaign `json:"data"`
	Paging *Paging    `json:"paging"`
}

type Campaign struct {
	ID                       go_types.Int64String    `json:"id"`
	AccountID                *go_types.Int64String   `json:"account_id"`
	AdStrategyID             *go_types.Int64String   `json:"ad_strategy_id"`
	AdLabels                 json.RawMessage         `json:"adlabels"`
	BidStrategy              *string                 `json:"bid_strategy"`
	BoostedObjectID          *go_types.Int64String   `json:"boosted_object_id"`
	BrandLiftStudies         json.RawMessage         `json:"brand_lift_studies"`
	BudgetRebalanceFlag      *bool                   `json:"budget_rebalance_flag"`
	BudgetRemaining          *go_types.Int64String   `json:"budget_remaining"`
	BuyingType               *string                 `json:"buying_type"`
	CanCreateBrandLiftStudy  *bool                   `json:"can_create_brand_lift_study"`
	CanUseSpendCap           *bool                   `json:"can_use_spend_cap"`
	ConfiguredStatus         *string                 `json:"configured_status"`
	CreatedTime              *f_types.DateTimeString `json:"created_time"`
	DailyBudget              *go_types.Int64String   `json:"daily_budget"`
	EffectiveStatus          *string                 `json:"effective_status"`
	IsSKAdNetworkAttribution *bool                   `json:"is_skadnetwork_attribution"`
	IssuesInfo               json.RawMessage         `json:"issues_info"`
	LastBudgetTogglingTime   *f_types.DateTimeString `json:"last_budget_toggling_time"`
	LifetimeBudget           *go_types.Int64String   `json:"lifetime_budget"`
	Name                     *string                 `json:"name"`
	Objective                *string                 `json:"objective"`
	PacingType               *[]string               `json:"pacing_type"`
	PromotedObject           json.RawMessage         `json:"promoted_object"`
	Recommendations          json.RawMessage         `json:"recommendations"`
	SmartPromotionType       *string                 `json:"smart_promotion_type"`
	SourceCampaign           *Campaign               `json:"source_campaign"`
	SourceCampaignID         *go_types.Int64String   `json:"source_campaign_id"`
	SpecialAdCategories      *[]string               `json:"special_ad_categories"`
	SpecialAdCategory        *string                 `json:"special_ad_category"`
	SpecialAdCategoryCountry *[]string               `json:"special_ad_category_country"`
	SpendCap                 *go_types.Int64String   `json:"spend_cap"`
	StartTime                *f_types.DateTimeString `json:"start_time"`
	Status                   *string                 `json:"status"`
	StopTime                 *f_types.DateTimeString `json:"stop_time"`
	ToplineID                *go_types.Int64String   `json:"topline_id"`
	UpdatedTime              *f_types.DateTimeString `json:"updated_time"`
}

type CampaignField string

const (
	CampaignFieldID                       CampaignField = "id"
	CampaignFieldAccountID                CampaignField = "account_id"
	CampaignFieldAdStrategyID             CampaignField = "ad_strategy_id"
	CampaignFieldAdLabels                 CampaignField = "adlabels"
	CampaignFieldBidStrategy              CampaignField = "bid_strategy"
	CampaignFieldBoostedObjectID          CampaignField = "boosted_object_id"
	CampaignFieldBrandLiftStudies         CampaignField = "brand_lift_studies"
	CampaignFieldBudgetRebalanceFlag      CampaignField = "budget_rebalance_flag"
	CampaignFieldBudgetRemaining          CampaignField = "budget_remaining"
	CampaignFieldBuyingType               CampaignField = "buying_type"
	CampaignFieldCanCreateBrandLiftStudy  CampaignField = "can_create_brand_lift_study"
	CampaignFieldCanUseSpendCap           CampaignField = "can_use_spend_cap"
	CampaignFieldConfiguredStatus         CampaignField = "configured_status"
	CampaignFieldCreatedTime              CampaignField = "created_time"
	CampaignFieldDailyBudget              CampaignField = "daily_budget"
	CampaignFieldEffectiveStatus          CampaignField = "effective_status"
	CampaignFieldIsSKAdNetworkAttribution CampaignField = "is_skadnetwork_attribution"
	CampaignFieldIssuesInfo               CampaignField = "issues_info"
	CampaignFieldLastBudgetTogglingTime   CampaignField = "last_budget_toggling_time"
	CampaignFieldLifetimeBudget           CampaignField = "lifetime_budget"
	CampaignFieldName                     CampaignField = "name"
	CampaignFieldObjective                CampaignField = "objective"
	CampaignFieldPacingType               CampaignField = "pacing_type"
	CampaignFieldPromotedObject           CampaignField = "promoted_object"
	CampaignFieldRecommendations          CampaignField = "recommendations"
	CampaignFieldSmartPromotionType       CampaignField = "smart_promotion_type"
	CampaignFieldSourceCampaign           CampaignField = "source_campaign"
	CampaignFieldSourceCampaignID         CampaignField = "source_campaign_id"
	CampaignFieldSpecialAdCategories      CampaignField = "special_ad_categories"
	CampaignFieldSpecialAdCategory        CampaignField = "special_ad_category"
	CampaignFieldSpecialAdCategoryCountry CampaignField = "special_ad_category_country"
	CampaignFieldSpendCap                 CampaignField = "spend_cap"
	CampaignFieldStartTime                CampaignField = "start_time"
	CampaignFieldStatus                   CampaignField = "status"
	CampaignFieldStopTime                 CampaignField = "stop_time"
	CampaignFieldToplineID                CampaignField = "topline_id"
	CampaignFieldUpdatedTime              CampaignField = "updated_time"
)

type GetCampaignsConfig struct {
	AccountID int64
	Fields    []CampaignField
}

func (service *Service) GetCampaigns(config *GetCampaignsConfig) (*[]Campaign, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetCampaignsConfig must not be a nil pointer")
	}

	values := url.Values{}
	fields := []string{}
	if len(config.Fields) == 0 {
		fields = append(fields, string(CampaignFieldID))
	} else {
		for _, field := range config.Fields {
			fields = append(fields, string(field))
		}
	}
	values.Set("fields", strings.Join(fields, ","))

	campaigns := []Campaign{}

	url := service.url(fmt.Sprintf("act_%v/campaigns?%s", config.AccountID, values.Encode()))

	for {
		campaignResponse := CampaignResponse{}
		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           url,
			ResponseModel: &campaignResponse,
		}
		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		campaigns = append(campaigns, campaignResponse.Data...)

		if campaignResponse.Paging == nil {
			break
		}

		if campaignResponse.Paging.Next == "" {
			break
		}

		url = campaignResponse.Paging.Next
	}

	return &campaigns, nil
}
