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

type AdSetResponse struct {
	Data   []AdSet `json:"data"`
	Paging *Paging `json:"paging"`
}

type AdSet struct {
	ID                            go_types.Int64String    `json:"id"`
	AccountID                     *go_types.Int64String   `json:"account_id"`
	AdLabels                      json.RawMessage         `json:"adlabels"`
	AdSetSchedule                 json.RawMessage         `json:"adset_schedule"`
	AssetFeedID                   *go_types.Int64String   `json:"asset_feed_id"`
	AttributionSpec               json.RawMessage         `json:"attribution_spec"`
	BidAdjustments                json.RawMessage         `json:"bid_adjustments"`
	BidAmount                     *uint32                 `json:"bid_amount"`
	BidConstraints                json.RawMessage         `json:"bid_constraints"`
	BidInfo                       *map[string]uint32      `json:"bid_info"`
	BidStrategy                   *string                 `json:"bid_strategy"`
	BillingEvent                  *string                 `json:"billing_event"`
	BudgetRemaining               *go_types.Int64String   `json:"budget_remaining"`
	Campaign                      *Campaign               `json:"campaign"`
	CampaignID                    *go_types.Int64String   `json:"campaign_id"`
	ConfiguredStatus              *string                 `json:"configured_status"`
	ContextualBundlingSpec        json.RawMessage         `json:"contextual_bundling_spec"`
	CreatedTime                   *f_types.DateTimeString `json:"created_time"`
	CreativeSequence              *[]go_types.Int64String `json:"creative_sequence"`
	DailyBudget                   *go_types.Int64String   `json:"daily_budget"`
	DailyMinSpendTarget           *go_types.Int64String   `json:"daily_min_spend_target"`
	DailySpendCap                 *go_types.Int64String   `json:"daily_spend_cap"`
	DestinationType               *string                 `json:"destination_type"`
	EffectiveStatus               *string                 `json:"effective_status"`
	EndTime                       *f_types.DateTimeString `json:"end_time"`
	FrequencyControlSpecs         json.RawMessage         `json:"frequency_control_specs"`
	InstagramActorID              *go_types.Int64String   `json:"instagram_actor_id"`
	IsDynamicCreative             *bool                   `json:"is_dynamic_creative"`
	IssuesInfo                    json.RawMessage         `json:"issues_info"`
	LearningStageInfo             json.RawMessage         `json:"learning_stage_info"`
	LifetimeBudget                *go_types.Int64String   `json:"lifetime_budget"`
	LifetimeImpressions           *int32                  `json:"lifetime_imps"`
	LifetimeMinSpendTarget        *go_types.Int64String   `json:"lifetime_min_spend_target"`
	LifetimeSpendCap              *go_types.Int64String   `json:"lifetime_spend_cap"`
	MultiOptimizationGoalWeight   *string                 `json:"multi_optimization_goal_weight"`
	Name                          *string                 `json:"name"`
	OptimizationGoal              *string                 `json:"optimization_goal"`
	OptimizationSubEvent          *string                 `json:"optimization_sub_event"`
	PacingType                    *[]string               `json:"pacing_type"`
	PromotedObject                json.RawMessage         `json:"promoted_object"`
	Recommendations               json.RawMessage         `json:"recommendations"`
	RecurringBudgetSemantics      *bool                   `json:"recurring_budget_semantics"`
	ReviewFeedback                *string                 `json:"review_feedback"`
	ReachAndFrequencyPredictionID *int                    `json:"rf_prediction_id"`
	SourceAdSet                   *AdSet                  `json:"source_adset"`
	SourceAdSetID                 *go_types.Int64String   `json:"source_adset_id"`
	StartTime                     *f_types.DateTimeString `json:"start_time"`
	Status                        *string                 `json:"status"`
	Targeting                     json.RawMessage         `json:"targeting"`
	TimeBasedAdRotationIDBlocks   *[][]int                `json:"time_based_ad_rotation_id_blocks"`
	TimeBasedAdRotationIntervals  *[]uint32               `json:"time_based_ad_rotation_intervals"`
	UpdatedTime                   *f_types.DateTimeString `json:"updated_time"`
	UseNewAppClick                *bool                   `json:"use_new_app_click"`
}

type AdSetField string

const (
	AdSetFieldID                            AdSetField = "id"
	AdSetFieldAccountID                     AdSetField = "account_id"
	AdSetFieldAdLabels                      AdSetField = "adlabels"
	AdSetFieldAdSetSchedule                 AdSetField = "adset_schedule"
	AdSetFieldAssetFeedID                   AdSetField = "asset_feed_id"
	AdSetFieldAttributionSpec               AdSetField = "attribution_spec"
	AdSetFieldBidAdjustments                AdSetField = "bid_adjustments"
	AdSetFieldBidAmount                     AdSetField = "bid_amount"
	AdSetFieldBidConstraints                AdSetField = "bid_constraints"
	AdSetFieldBidInfo                       AdSetField = "bid_info"
	AdSetFieldBidStrategy                   AdSetField = "bid_strategy"
	AdSetFieldBillingEvent                  AdSetField = "billing_event"
	AdSetFieldBudgetRemaining               AdSetField = "budget_remaining"
	AdSetFieldCampaign                      AdSetField = "campaign"
	AdSetFieldCampaignID                    AdSetField = "campaign_id"
	AdSetFieldConfiguredStatus              AdSetField = "configured_status"
	AdSetFieldContextualBundlingSpec        AdSetField = "contextual_bundling_spec"
	AdSetFieldCreatedTime                   AdSetField = "created_time"
	AdSetFieldCreativeSequence              AdSetField = "creative_sequence"
	AdSetFieldDailyBudget                   AdSetField = "daily_budget"
	AdSetFieldDailyMinSpendTarget           AdSetField = "daily_min_spend_target"
	AdSetFieldDailySpendCap                 AdSetField = "daily_spend_cap"
	AdSetFieldDestinationType               AdSetField = "destination_type"
	AdSetFieldEffectiveStatus               AdSetField = "effective_status"
	AdSetFieldEndTime                       AdSetField = "end_time"
	AdSetFieldFrequencyControlSpecs         AdSetField = "frequency_control_specs"
	AdSetFieldInstagramActorID              AdSetField = "instagram_actor_id"
	AdSetFieldIsDynamicCreative             AdSetField = "is_dynamic_creative"
	AdSetFieldIssuesInfo                    AdSetField = "issues_info"
	AdSetFieldLearningStageInfo             AdSetField = "learning_stage_info"
	AdSetFieldLifetimeBudget                AdSetField = "lifetime_budget"
	AdSetFieldLifetimeImpressions           AdSetField = "lifetime_imps"
	AdSetFieldLifetimeMinSpendTarget        AdSetField = "lifetime_min_spend_target"
	AdSetFieldLifetimeSpendCap              AdSetField = "lifetime_spend_cap"
	AdSetFieldMultiOptimizationGoalWeight   AdSetField = "multi_optimization_goal_weight"
	AdSetFieldName                          AdSetField = "name"
	AdSetFieldOptimizationGoal              AdSetField = "optimization_goal"
	AdSetFieldOptimizationSubEvent          AdSetField = "optimization_sub_event"
	AdSetFieldPacingType                    AdSetField = "pacing_type"
	AdSetFieldPromotedObject                AdSetField = "promoted_object"
	AdSetFieldRecommendations               AdSetField = "recommendations"
	AdSetFieldRecurringBudgetSemantics      AdSetField = "recurring_budget_semantics"
	AdSetFieldReviewFeedback                AdSetField = "review_feedback"
	AdSetFieldReachAndFrequencyPredictionID AdSetField = "rf_prediction_id"
	AdSetFieldSourceAdSet                   AdSetField = "source_adset"
	AdSetFieldSourceAdSetID                 AdSetField = "source_adset_id"
	AdSetFieldStartTime                     AdSetField = "start_time"
	AdSetFieldStatus                        AdSetField = "status"
	AdSetFieldTargeting                     AdSetField = "targeting"
	AdSetFieldTimeBasedAdRotationIDBlocks   AdSetField = "time_based_ad_rotation_id_blocks"
	AdSetFieldTimeBasedAdRotationIntervals  AdSetField = "time_based_ad_rotation_intervals"
	AdSetFieldUpdatedTime                   AdSetField = "updated_time"
	AdSetFieldUseNewAppClick                AdSetField = "use_new_app_click"
)

type GetAdSetsConfig struct {
	AccountID int64
	Fields    []AdSetField
}

func (service *Service) GetAdSets(config *GetAdSetsConfig) (*[]AdSet, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetAdSetsConfig must not be a nil pointer")
	}

	values := url.Values{}
	fields := []string{}
	if len(config.Fields) == 0 {
		fields = append(fields, string(AdSetFieldID))
	} else {
		for _, field := range config.Fields {
			fields = append(fields, string(field))
		}
	}
	values.Set("fields", strings.Join(fields, ","))

	adSets := []AdSet{}

	url := service.url(fmt.Sprintf("act_%v/adsets?%s", config.AccountID, values.Encode()))

	for {
		adSetResponse := AdSetResponse{}
		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           url,
			ResponseModel: &adSetResponse,
		}
		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		adSets = append(adSets, adSetResponse.Data...)

		if adSetResponse.Paging == nil {
			break
		}

		if adSetResponse.Paging.Next == "" {
			break
		}

		url = adSetResponse.Paging.Next
	}

	return &adSets, nil
}
