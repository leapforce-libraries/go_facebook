package facebook

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type AdInsightsResponse struct {
	Data   []AdInsights `json:"data"`
	Paging *Paging      `json:"paging"`
}

type AdInsights struct {
	AccountCurrency                           *string                 `json:"account_currency"`
	AccountID                                 *go_types.Int64String   `json:"account_id"`
	AccountName                               *string                 `json:"account_name"`
	ActionValues                              json.RawMessage         `json:"action_values"`
	Actions                                   json.RawMessage         `json:"actions"`
	ActivityRecency                           *string                 `json:"activity_recency"`
	AdClickActions                            json.RawMessage         `json:"ad_click_actions"`
	AdFormatAsset                             *string                 `json:"ad_format_asset"`
	AdID                                      *go_types.Int64String   `json:"ad_id"`
	AdImpressionActions                       json.RawMessage         `json:"ad_impression_actions"`
	AdName                                    *string                 `json:"ad_name"`
	AdsetID                                   *string                 `json:"adset_id"`
	AdsetName                                 *string                 `json:"adset_name"`
	AgeTargeting                              *string                 `json:"age_targeting"`
	AttributionSetting                        *string                 `json:"attribution_setting"`
	AuctionBid                                *go_types.Float64String `json:"auction_bid"`
	AuctionCompetitiveness                    *go_types.Float64String `json:"auction_competitiveness"`
	AuctionMaxCompetitorBid                   *go_types.Float64String `json:"auction_max_competitor_bid"`
	BodyAsset                                 json.RawMessage         `json:"body_asset"`
	BuyingType                                *string                 `json:"buying_type"`
	CampaignID                                *go_types.Int64String   `json:"campaign_id"`
	CampaignName                              *string                 `json:"campaign_name"`
	CanvasAvgViewPercent                      *go_types.Float64String `json:"canvas_avg_view_percent"`
	CanvasAvgViewTime                         *go_types.Float64String `json:"canvas_avg_view_time"`
	CatalogSegmentActions                     json.RawMessage         `json:"catalog_segment_actions"`
	CatalogSegmentValue                       json.RawMessage         `json:"catalog_segment_value"`
	CatalogSegmentValueMobilePurchaseRoas     json.RawMessage         `json:"catalog_segment_value_mobile_purchase_roas"`
	CatalogSegmentValueOmniPurchaseRoas       json.RawMessage         `json:"catalog_segment_value_omni_purchase_roas"`
	CatalogSegmentValueWebsitePurchaseRoas    json.RawMessage         `json:"catalog_segment_value_website_purchase_roas"`
	Clicks                                    *go_types.Int64String   `json:"clicks"`
	ComparisonNode                            json.RawMessage         `json:"comparison_node"`
	ConversionValues                          json.RawMessage         `json:"conversion_values"`
	Conversions                               json.RawMessage         `json:"conversions"`
	ConvertedProductQuantity                  json.RawMessage         `json:"converted_product_quantity"`
	ConvertedProductValue                     json.RawMessage         `json:"converted_product_value"`
	CostPer15SecVideoView                     json.RawMessage         `json:"cost_per_15_sec_video_view"`
	CostPer2SecContinuousVideoView            json.RawMessage         `json:"cost_per_2_sec_continuous_video_view"`
	CostPerActionType                         json.RawMessage         `json:"cost_per_action_type"`
	CostPerAdClick                            json.RawMessage         `json:"cost_per_ad_click"`
	CostPerConversion                         json.RawMessage         `json:"cost_per_conversion"`
	CostPerDDACountByConvs                    *go_types.Float64String `json:"cost_per_dda_countby_convs"`
	CostPerInlineLinkClick                    *go_types.Float64String `json:"cost_per_inline_link_click"`
	CostPerInlinePostEngagement               *go_types.Float64String `json:"cost_per_inline_post_engagement"`
	CostPerOneThousandAdImpression            json.RawMessage         `json:"cost_per_one_thousand_ad_impression"`
	CostPerOutboundClick                      json.RawMessage         `json:"cost_per_outbound_click"`
	CostPerStoreVisitAction                   json.RawMessage         `json:"cost_per_store_visit_action"`
	CostPerThruplay                           json.RawMessage         `json:"cost_per_thruplay"`
	CostPerUniqueActionType                   json.RawMessage         `json:"cost_per_unique_action_type"`
	CostPerUniqueClick                        *go_types.Float64String `json:"cost_per_unique_click"`
	CostPerUniqueConversion                   json.RawMessage         `json:"cost_per_unique_conversion"`
	CostPerUniqueInlineLinkClick              *go_types.Float64String `json:"cost_per_unique_inline_link_click"`
	CostPerUniqueOutboundClick                json.RawMessage         `json:"cost_per_unique_outbound_click"`
	Country                                   *string                 `json:"country"`
	CPC                                       *go_types.Float64String `json:"cpc"`
	CPM                                       *go_types.Float64String `json:"cpm"`
	CPP                                       *go_types.Float64String `json:"cpp"`
	CreatedTime                               *string                 `json:"created_time"`
	CTR                                       *go_types.Float64String `json:"ctr"`
	DateStart                                 *DateString             `json:"date_start"`
	DateStop                                  *DateString             `json:"date_stop"`
	DDACountByConvs                           *go_types.Float64String `json:"dda_countby_convs"`
	DescriptionAsset                          json.RawMessage         `json:"description_asset"`
	DevicePlatform                            *string                 `json:"device_platform"`
	DMA                                       *string                 `json:"dma"`
	EstimatedAdRecallRateLowerBound           *go_types.Float64String `json:"estimated_ad_recall_rate_lower_bound"`
	EstimatedAdRecallRateUpperBound           *go_types.Float64String `json:"estimated_ad_recall_rate_upper_bound"`
	EstimatedAdRecallersLowerBound            *go_types.Float64String `json:"estimated_ad_recallers_lower_bound"`
	EstimatedAdRecallersUpperBound            *go_types.Float64String `json:"estimated_ad_recallers_upper_bound"`
	Frequency                                 *go_types.Float64String `json:"frequency"`
	FrequencyValue                            *string                 `json:"frequency_value"`
	FullViewImpressions                       *go_types.Int64String   `json:"full_view_impressions"`
	FullViewReach                             *go_types.Float64String `json:"full_view_reach"`
	GenderTargeting                           *string                 `json:"gender_targeting"`
	HourlyStatsAggregatedByAdvertiserTimeZone *string                 `json:"hourly_stats_aggregated_by_advertiser_time_zone"`
	HourlyStatsAggregatedByAudienceTimeZone   *string                 `json:"hourly_stats_aggregated_by_audience_time_zone"`
	ImageAsset                                json.RawMessage         `json:"image_asset"`
	ImpressionDevice                          *string                 `json:"impression_device"`
	Impressions                               *go_types.Int64String   `json:"impressions"`
	ImpressionsDummy                          *string                 `json:"impressions_dummy"`
	InlineLinkClickCTR                        *go_types.Float64String `json:"inline_link_click_ctr"`
	InlineLinkClicks                          *go_types.Int64String   `json:"inline_link_clicks"`
	InlinePostEngagement                      *go_types.Float64String `json:"inline_post_engagement"`
	InstantExperienceClicksToOpen             *go_types.Int64String   `json:"instant_experience_clicks_to_open"`
	InstantExperienceClicksToStart            *go_types.Int64String   `json:"instant_experience_clicks_to_start"`
	InstantExperienceOutboundClicks           *go_types.Int64String   `json:"instant_experience_outbound_clicks"`
	InteractiveComponentTap                   json.RawMessage         `json:"interactive_component_tap"`
	Labels                                    *string                 `json:"labels"`
	Location                                  *string                 `json:"location"`
	MediaAsset                                json.RawMessage         `json:"media_asset"`
	MobileAppPurchaseRoas                     json.RawMessage         `json:"mobile_app_purchase_roas"`
	Objective                                 *string                 `json:"objective"`
	OptimizationGoal                          *string                 `json:"optimization_goal"`
	OutboundClicks                            json.RawMessage         `json:"outbound_clicks"`
	OutboundClicksCTR                         json.RawMessage         `json:"outbound_clicks_ctr"`
	PlacePageID                               *string                 `json:"place_page_id"`
	PlacePageName                             *string                 `json:"place_page_name"`
	PlatformPosition                          *string                 `json:"platform_position"`
	ProductID                                 *string                 `json:"product_id"`
	PublisherPlatform                         *string                 `json:"publisher_platform"`
	PurchaseRoas                              json.RawMessage         `json:"purchase_roas"`
	QualifyingQuestionQualifyAnswerRate       *go_types.Float64String `json:"qualifying_question_qualify_answer_rate"`
	Reach                                     *go_types.Float64String `json:"reach"`
	RuleAsset                                 json.RawMessage         `json:"rule_asset"`
	SocialSpend                               *go_types.Float64String `json:"social_spend"`
	Spend                                     *go_types.Float64String `json:"spend"`
	StoreVisitActions                         json.RawMessage         `json:"store_visit_actions"`
	TitleAsset                                json.RawMessage         `json:"title_asset"`
	UniqueActions                             json.RawMessage         `json:"unique_actions"`
	UniqueClicks                              *go_types.Int64String   `json:"unique_clicks"`
	UniqueConversions                         json.RawMessage         `json:"unique_conversions"`
	UniqueCTR                                 *go_types.Float64String `json:"unique_ctr"`
	UniqueInlineLinkClickCTR                  *go_types.Float64String `json:"unique_inline_link_click_ctr"`
	UniqueInlineLinkClicks                    *go_types.Int64String   `json:"unique_inline_link_clicks"`
	UniqueLinkClicksCTR                       *go_types.Float64String `json:"unique_link_clicks_ctr"`
	UniqueOutboundClicks                      json.RawMessage         `json:"unique_outbound_clicks"`
	UniqueOutboundClicksCTR                   json.RawMessage         `json:"unique_outbound_clicks_ctr"`
	UniqueVideoView15Sec                      json.RawMessage         `json:"unique_video_view_15_sec"`
	UpdatedTime                               *string                 `json:"updated_time"`
	Video15SecWatchedActions                  json.RawMessage         `json:"video_15_sec_watched_actions"`
	Video30SecWatchedActions                  json.RawMessage         `json:"video_30_sec_watched_actions"`
	VideoAsset                                json.RawMessage         `json:"video_asset"`
	VideoAvgTimeWatchedActions                json.RawMessage         `json:"video_avg_time_watched_actions"`
	VideoContinuous2SecWatchedActions         json.RawMessage         `json:"video_continuous_2_sec_watched_actions"`
	VideoP100WatchedActions                   json.RawMessage         `json:"video_p100_watched_actions"`
	VideoP25WatchedActions                    json.RawMessage         `json:"video_p25_watched_actions"`
	VideoP50WatchedActions                    json.RawMessage         `json:"video_p50_watched_actions"`
	VideoP75WatchedActions                    json.RawMessage         `json:"video_p75_watched_actions"`
	VideoP95WatchedActions                    json.RawMessage         `json:"video_p95_watched_actions"`
	VideoPlayActions                          json.RawMessage         `json:"video_play_actions"`
	VideoPlayCurveActions                     json.RawMessage         `json:"video_play_curve_actions"`
	VideoPlayRetention0To15sActions           json.RawMessage         `json:"video_play_retention_0_to_15s_actions"`
	VideoPlayRetention20To60sActions          json.RawMessage         `json:"video_play_retention_20_to_60s_actions"`
	VideoPlayRetentionGraphActions            json.RawMessage         `json:"video_play_retention_graph_actions"`
	VideoTimeWatchedActions                   json.RawMessage         `json:"video_time_watched_actions"`
	WebsiteCTR                                json.RawMessage         `json:"website_ctr"`
	WebsitePurchaseRoas                       json.RawMessage         `json:"website_purchase_roas"`
	WishBid                                   *go_types.Float64String `json:"wish_bid"`
}

type AdInsightsField string

const (
	AdInsightsFieldAccountCurrency                           AdInsightsField = "account_currency"
	AdInsightsFieldAccountID                                 AdInsightsField = "account_id"
	AdInsightsFieldAccountName                               AdInsightsField = "account_name"
	AdInsightsFieldActionValues                              AdInsightsField = "action_values"
	AdInsightsFieldActions                                   AdInsightsField = "actions"
	AdInsightsFieldActivityRecency                           AdInsightsField = "activity_recency"
	AdInsightsFieldAdClickActions                            AdInsightsField = "ad_click_actions"
	AdInsightsFieldAdFormatAsset                             AdInsightsField = "ad_format_asset"
	AdInsightsFieldAdID                                      AdInsightsField = "ad_id"
	AdInsightsFieldAdImpressionActions                       AdInsightsField = "ad_impression_actions"
	AdInsightsFieldAdName                                    AdInsightsField = "ad_name"
	AdInsightsFieldAdsetID                                   AdInsightsField = "adset_id"
	AdInsightsFieldAdsetName                                 AdInsightsField = "adset_name"
	AdInsightsFieldAgeTargeting                              AdInsightsField = "age_targeting"
	AdInsightsFieldAttributionSetting                        AdInsightsField = "attribution_setting"
	AdInsightsFieldAuctionBid                                AdInsightsField = "auction_bid"
	AdInsightsFieldAuctionCompetitiveness                    AdInsightsField = "auction_competitiveness"
	AdInsightsFieldAuctionMaxCompetitorBid                   AdInsightsField = "auction_max_competitor_bid"
	AdInsightsFieldBodyAsset                                 AdInsightsField = "body_asset"
	AdInsightsFieldBuyingType                                AdInsightsField = "buying_type"
	AdInsightsFieldCampaignID                                AdInsightsField = "campaign_id"
	AdInsightsFieldCampaignName                              AdInsightsField = "campaign_name"
	AdInsightsFieldCanvasAvgViewPercent                      AdInsightsField = "canvas_avg_view_percent"
	AdInsightsFieldCanvasAvgViewTime                         AdInsightsField = "canvas_avg_view_time"
	AdInsightsFieldCatalogSegmentActions                     AdInsightsField = "catalog_segment_actions"
	AdInsightsFieldCatalogSegmentValue                       AdInsightsField = "catalog_segment_value"
	AdInsightsFieldCatalogSegmentValueMobilePurchaseRoas     AdInsightsField = "catalog_segment_value_mobile_purchase_roas"
	AdInsightsFieldCatalogSegmentValueOmniPurchaseRoas       AdInsightsField = "catalog_segment_value_omni_purchase_roas"
	AdInsightsFieldCatalogSegmentValueWebsitePurchaseRoas    AdInsightsField = "catalog_segment_value_website_purchase_roas"
	AdInsightsFieldClicks                                    AdInsightsField = "clicks"
	AdInsightsFieldComparisonNode                            AdInsightsField = "comparison_node"
	AdInsightsFieldConversionValues                          AdInsightsField = "conversion_values"
	AdInsightsFieldConversions                               AdInsightsField = "conversions"
	AdInsightsFieldConvertedProductQuantity                  AdInsightsField = "converted_product_quantity"
	AdInsightsFieldConvertedProductValue                     AdInsightsField = "converted_product_value"
	AdInsightsFieldCostPer15SecVideoView                     AdInsightsField = "cost_per_15_sec_video_view"
	AdInsightsFieldCostPer2SecContinuousVideoView            AdInsightsField = "cost_per_2_sec_continuous_video_view"
	AdInsightsFieldCostPerActionType                         AdInsightsField = "cost_per_action_type"
	AdInsightsFieldCostPerAdClick                            AdInsightsField = "cost_per_ad_click"
	AdInsightsFieldCostPerConversion                         AdInsightsField = "cost_per_conversion"
	AdInsightsFieldCostPerDDACountByConvs                    AdInsightsField = "cost_per_dda_countby_convs"
	AdInsightsFieldCostPerInlineLinkClick                    AdInsightsField = "cost_per_inline_link_click"
	AdInsightsFieldCostPerInlinePostEngagement               AdInsightsField = "cost_per_inline_post_engagement"
	AdInsightsFieldCostPerOneThousandAdImpression            AdInsightsField = "cost_per_one_thousand_ad_impression"
	AdInsightsFieldCostPerOutboundClick                      AdInsightsField = "cost_per_outbound_click"
	AdInsightsFieldCostPerStoreVisitAction                   AdInsightsField = "cost_per_store_visit_action"
	AdInsightsFieldCostPerThruplay                           AdInsightsField = "cost_per_thruplay"
	AdInsightsFieldCostPerUniqueActionType                   AdInsightsField = "cost_per_unique_action_type"
	AdInsightsFieldCostPerUniqueClick                        AdInsightsField = "cost_per_unique_click"
	AdInsightsFieldCostPerUniqueConversion                   AdInsightsField = "cost_per_unique_conversion"
	AdInsightsFieldCostPerUniqueInlineLinkClick              AdInsightsField = "cost_per_unique_inline_link_click"
	AdInsightsFieldCostPerUniqueOutboundClick                AdInsightsField = "cost_per_unique_outbound_click"
	AdInsightsFieldCountry                                   AdInsightsField = "country"
	AdInsightsFieldCPC                                       AdInsightsField = "cpc"
	AdInsightsFieldCPM                                       AdInsightsField = "cpm"
	AdInsightsFieldCPP                                       AdInsightsField = "cpp"
	AdInsightsFieldCreatedTime                               AdInsightsField = "created_time"
	AdInsightsFieldCTR                                       AdInsightsField = "ctr"
	AdInsightsFieldDateStart                                 AdInsightsField = "date_start"
	AdInsightsFieldDateStop                                  AdInsightsField = "date_stop"
	AdInsightsFieldDDACountByConvs                           AdInsightsField = "dda_countby_convs"
	AdInsightsFieldDescriptionAsset                          AdInsightsField = "description_asset"
	AdInsightsFieldDevicePlatform                            AdInsightsField = "device_platform"
	AdInsightsFieldDMA                                       AdInsightsField = "dma"
	AdInsightsFieldEstimatedAdRecallRateLowerBound           AdInsightsField = "estimated_ad_recall_rate_lower_bound"
	AdInsightsFieldEstimatedAdRecallRateUpperBound           AdInsightsField = "estimated_ad_recall_rate_upper_bound"
	AdInsightsFieldEstimatedAdRecallersLowerBound            AdInsightsField = "estimated_ad_recallers_lower_bound"
	AdInsightsFieldEstimatedAdRecallersUpperBound            AdInsightsField = "estimated_ad_recallers_upper_bound"
	AdInsightsFieldFrequency                                 AdInsightsField = "frequency"
	AdInsightsFieldFrequencyValue                            AdInsightsField = "frequency_value"
	AdInsightsFieldFullViewImpressions                       AdInsightsField = "full_view_impressions"
	AdInsightsFieldFullViewReach                             AdInsightsField = "full_view_reach"
	AdInsightsFieldGenderTargeting                           AdInsightsField = "gender_targeting"
	AdInsightsFieldHourlyStatsAggregatedByAdvertiserTimeZone AdInsightsField = "hourly_stats_aggregated_by_advertiser_time_zone"
	AdInsightsFieldHourlyStatsAggregatedByAudienceTimeZone   AdInsightsField = "hourly_stats_aggregated_by_audience_time_zone"
	AdInsightsFieldImageAsset                                AdInsightsField = "image_asset"
	AdInsightsFieldImpressionDevice                          AdInsightsField = "impression_device"
	AdInsightsFieldImpressions                               AdInsightsField = "impressions"
	AdInsightsFieldImpressionsDummy                          AdInsightsField = "impressions_dummy"
	AdInsightsFieldInlineLinkClickCTR                        AdInsightsField = "inline_link_click_ctr"
	AdInsightsFieldInlineLinkClicks                          AdInsightsField = "inline_link_clicks"
	AdInsightsFieldInlinePostEngagement                      AdInsightsField = "inline_post_engagement"
	AdInsightsFieldInstantExperienceClicksToOpen             AdInsightsField = "instant_experience_clicks_to_open"
	AdInsightsFieldInstantExperienceClicksToStart            AdInsightsField = "instant_experience_clicks_to_start"
	AdInsightsFieldInstantExperienceOutboundClicks           AdInsightsField = "instant_experience_outbound_clicks"
	AdInsightsFieldInteractiveComponentTap                   AdInsightsField = "interactive_component_tap"
	AdInsightsFieldLabels                                    AdInsightsField = "labels"
	AdInsightsFieldLocation                                  AdInsightsField = "location"
	AdInsightsFieldMediaAsset                                AdInsightsField = "media_asset"
	AdInsightsFieldMobileAppPurchaseRoas                     AdInsightsField = "mobile_app_purchase_roas"
	AdInsightsFieldObjective                                 AdInsightsField = "objective"
	AdInsightsFieldOptimizationGoal                          AdInsightsField = "optimization_goal"
	AdInsightsFieldOutboundClicks                            AdInsightsField = "outbound_clicks"
	AdInsightsFieldOutboundClicksCTR                         AdInsightsField = "outbound_clicks_ctr"
	AdInsightsFieldPlacePageID                               AdInsightsField = "place_page_id"
	AdInsightsFieldPlacePageName                             AdInsightsField = "place_page_name"
	AdInsightsFieldPlatformPosition                          AdInsightsField = "platform_position"
	AdInsightsFieldProductID                                 AdInsightsField = "product_id"
	AdInsightsFieldPublisherPlatform                         AdInsightsField = "publisher_platform"
	AdInsightsFieldPurchaseRoas                              AdInsightsField = "purchase_roas"
	AdInsightsFieldQualifyingQuestionQualifyAnswerRate       AdInsightsField = "qualifying_question_qualify_answer_rate"
	AdInsightsFieldReach                                     AdInsightsField = "reach"
	AdInsightsFieldRuleAsset                                 AdInsightsField = "rule_asset"
	AdInsightsFieldSocialSpend                               AdInsightsField = "social_spend"
	AdInsightsFieldSpend                                     AdInsightsField = "spend"
	AdInsightsFieldStoreVisitActions                         AdInsightsField = "store_visit_actions"
	AdInsightsFieldTitleAsset                                AdInsightsField = "title_asset"
	AdInsightsFieldUniqueActions                             AdInsightsField = "unique_actions"
	AdInsightsFieldUniqueClicks                              AdInsightsField = "unique_clicks"
	AdInsightsFieldUniqueConversions                         AdInsightsField = "unique_conversions"
	AdInsightsFieldUniqueCTR                                 AdInsightsField = "unique_ctr"
	AdInsightsFieldUniqueInlineLinkClickCTR                  AdInsightsField = "unique_inline_link_click_ctr"
	AdInsightsFieldUniqueInlineLinkClicks                    AdInsightsField = "unique_inline_link_clicks"
	AdInsightsFieldUniqueLinkClicksCTR                       AdInsightsField = "unique_link_clicks_ctr"
	AdInsightsFieldUniqueOutboundClicks                      AdInsightsField = "unique_outbound_clicks"
	AdInsightsFieldUniqueOutboundClicksCTR                   AdInsightsField = "unique_outbound_clicks_ctr"
	AdInsightsFieldUniqueVideoView15Sec                      AdInsightsField = "unique_video_view_15_sec"
	AdInsightsFieldUpdatedTime                               AdInsightsField = "updated_time"
	AdInsightsFieldVideo15SecWatchedActions                  AdInsightsField = "video_15_sec_watched_actions"
	AdInsightsFieldVideo30SecWatchedActions                  AdInsightsField = "video_30_sec_watched_actions"
	AdInsightsFieldVideoAsset                                AdInsightsField = "video_asset"
	AdInsightsFieldVideoAvgTimeWatchedActions                AdInsightsField = "video_avg_time_watched_actions"
	AdInsightsFieldVideoContinuous2SecWatchedActions         AdInsightsField = "video_continuous_2_sec_watched_actions"
	AdInsightsFieldVideoP100WatchedActions                   AdInsightsField = "video_p100_watched_actions"
	AdInsightsFieldVideoP25WatchedActions                    AdInsightsField = "video_p25_watched_actions"
	AdInsightsFieldVideoP50WatchedActions                    AdInsightsField = "video_p50_watched_actions"
	AdInsightsFieldVideoP75WatchedActions                    AdInsightsField = "video_p75_watched_actions"
	AdInsightsFieldVideoP95WatchedActions                    AdInsightsField = "video_p95_watched_actions"
	AdInsightsFieldVideoPlayActions                          AdInsightsField = "video_play_actions"
	AdInsightsFieldVideoPlayCurveActions                     AdInsightsField = "video_play_curve_actions"
	AdInsightsFieldVideoPlayRetention0To15sActions           AdInsightsField = "video_play_retention_0_to_15s_actions"
	AdInsightsFieldVideoPlayRetention20To60sActions          AdInsightsField = "video_play_retention_20_to_60s_actions"
	AdInsightsFieldVideoPlayRetentionGraphActions            AdInsightsField = "video_play_retention_graph_actions"
	AdInsightsFieldVideoTimeWatchedActions                   AdInsightsField = "video_time_watched_actions"
	AdInsightsFieldWebsiteCTR                                AdInsightsField = "website_ctr"
	AdInsightsFieldWebsitePurchaseRoas                       AdInsightsField = "website_purchase_roas"
	AdInsightsFieldWishBid                                   AdInsightsField = "wish_bid"
)

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

type GetAdInsightsConfig struct {
	AdID       int64
	DatePreset *DatePreset
	Fields     []AdInsightsField
}

func (service *Service) GetAdInsights(config *GetAdInsightsConfig) (*[]AdInsights, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetAdInsightsConfig must not be a nil pointer")
	}

	values := url.Values{}
	fields := []string{}

	if config.DatePreset != nil {
		values.Set("date_preset", string(*config.DatePreset))
	}
	if len(config.Fields) == 0 {
		fields = append(fields, string(AdInsightsFieldAdID))
	} else {
		for _, field := range config.Fields {
			fields = append(fields, string(field))
		}
	}
	values.Set("fields", strings.Join(fields, ","))

	adInsights := []AdInsights{}

	url := service.url(fmt.Sprintf("%v/insights?%s", config.AdID, values.Encode()))

	for {
		adInsightsResponse := AdInsightsResponse{}
		requestConfig := go_http.RequestConfig{
			URL:           url,
			ResponseModel: &adInsightsResponse,
		}
		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		adInsights = append(adInsights, adInsightsResponse.Data...)

		if adInsightsResponse.Paging == nil {
			break
		}

		url = adInsightsResponse.Paging.Next
	}

	return &adInsights, nil
}
