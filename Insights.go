package facebook

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"cloud.google.com/go/civil"
	errortools "github.com/leapforce-libraries/go_errortools"
	f_types "github.com/leapforce-libraries/go_facebook/types"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type InsightsResponse struct {
	Data   []Insights `json:"data"`
	Paging *Paging    `json:"paging"`
}

type Insights struct {
	AccountCurrency                           *string                 `json:"account_currency"`
	AccountID                                 *go_types.Int64String   `json:"account_id"`
	AccountName                               *string                 `json:"account_name"`
	ActionValues                              *[]AdsActionStats       `json:"action_values"`
	Actions                                   *[]AdsActionStats       `json:"actions"`
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
	CatalogSegmentActions                     *[]AdsActionStats       `json:"catalog_segment_actions"`
	CatalogSegmentValue                       *[]AdsActionStats       `json:"catalog_segment_value"`
	CatalogSegmentValueMobilePurchaseROAS     *[]AdsActionStats       `json:"catalog_segment_value_mobile_purchase_roas"`
	CatalogSegmentValueOmniPurchaseROAS       *[]AdsActionStats       `json:"catalog_segment_value_omni_purchase_roas"`
	CatalogSegmentValueWebsitePurchaseROAS    *[]AdsActionStats       `json:"catalog_segment_value_website_purchase_roas"`
	Clicks                                    *go_types.Int64String   `json:"clicks"`
	ComparisonNode                            json.RawMessage         `json:"comparison_node"`
	ConversionValues                          *[]AdsActionStats       `json:"conversion_values"`
	Conversions                               *[]AdsActionStats       `json:"conversions"`
	ConvertedProductQuantity                  *[]AdsActionStats       `json:"converted_product_quantity"`
	ConvertedProductValue                     *[]AdsActionStats       `json:"converted_product_value"`
	CostPer15SecVideoView                     *[]AdsActionStats       `json:"cost_per_15_sec_video_view"`
	CostPer2SecContinuousVideoView            *[]AdsActionStats       `json:"cost_per_2_sec_continuous_video_view"`
	CostPerActionType                         *[]AdsActionStats       `json:"cost_per_action_type"`
	CostPerAdClick                            *[]AdsActionStats       `json:"cost_per_ad_click"`
	CostPerConversion                         *[]AdsActionStats       `json:"cost_per_conversion"`
	CostPerDDACountByConvs                    *go_types.Float64String `json:"cost_per_dda_countby_convs"`
	CostPerInlineLinkClick                    *go_types.Float64String `json:"cost_per_inline_link_click"`
	CostPerInlinePostEngagement               *go_types.Float64String `json:"cost_per_inline_post_engagement"`
	CostPerOneThousandAdImpression            *[]AdsActionStats       `json:"cost_per_one_thousand_ad_impression"`
	CostPerOutboundClick                      *[]AdsActionStats       `json:"cost_per_outbound_click"`
	CostPerStoreVisitAction                   *[]AdsActionStats       `json:"cost_per_store_visit_action"`
	CostPerThruplay                           *[]AdsActionStats       `json:"cost_per_thruplay"`
	CostPerUniqueActionType                   *[]AdsActionStats       `json:"cost_per_unique_action_type"`
	CostPerUniqueClick                        *go_types.Float64String `json:"cost_per_unique_click"`
	CostPerUniqueConversion                   *[]AdsActionStats       `json:"cost_per_unique_conversion"`
	CostPerUniqueInlineLinkClick              *go_types.Float64String `json:"cost_per_unique_inline_link_click"`
	CostPerUniqueOutboundClick                *[]AdsActionStats       `json:"cost_per_unique_outbound_click"`
	Country                                   *string                 `json:"country"`
	CPC                                       *go_types.Float64String `json:"cpc"`
	CPM                                       *go_types.Float64String `json:"cpm"`
	CPP                                       *go_types.Float64String `json:"cpp"`
	CreatedTime                               *string                 `json:"created_time"`
	CTR                                       *go_types.Float64String `json:"ctr"`
	DateStart                                 f_types.DateString      `json:"date_start"`
	DateStop                                  f_types.DateString      `json:"date_stop"`
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
	InteractiveComponentTap                   *[]AdsActionStats       `json:"interactive_component_tap"`
	Labels                                    *string                 `json:"labels"`
	Location                                  *string                 `json:"location"`
	MediaAsset                                json.RawMessage         `json:"media_asset"`
	MobileAppPurchaseROAS                     *[]AdsActionStats       `json:"mobile_app_purchase_roas"`
	Objective                                 *string                 `json:"objective"`
	OptimizationGoal                          *string                 `json:"optimization_goal"`
	OutboundClicks                            *[]AdsActionStats       `json:"outbound_clicks"`
	OutboundClicksCTR                         *[]AdsActionStats       `json:"outbound_clicks_ctr"`
	PlacePageID                               *string                 `json:"place_page_id"`
	PlacePageName                             *string                 `json:"place_page_name"`
	PlatformPosition                          *string                 `json:"platform_position"`
	ProductID                                 *string                 `json:"product_id"`
	PublisherPlatform                         *string                 `json:"publisher_platform"`
	PurchaseROAS                              *[]AdsActionStats       `json:"purchase_roas"`
	QualifyingQuestionQualifyAnswerRate       *go_types.Float64String `json:"qualifying_question_qualify_answer_rate"`
	Reach                                     *go_types.Float64String `json:"reach"`
	RuleAsset                                 json.RawMessage         `json:"rule_asset"`
	SocialSpend                               *go_types.Float64String `json:"social_spend"`
	Spend                                     *go_types.Float64String `json:"spend"`
	StoreVisitActions                         *[]AdsActionStats       `json:"store_visit_actions"`
	TitleAsset                                json.RawMessage         `json:"title_asset"`
	UniqueActions                             *[]AdsActionStats       `json:"unique_actions"`
	UniqueClicks                              *go_types.Int64String   `json:"unique_clicks"`
	UniqueConversions                         *[]AdsActionStats       `json:"unique_conversions"`
	UniqueCTR                                 *go_types.Float64String `json:"unique_ctr"`
	UniqueInlineLinkClickCTR                  *go_types.Float64String `json:"unique_inline_link_click_ctr"`
	UniqueInlineLinkClicks                    *go_types.Int64String   `json:"unique_inline_link_clicks"`
	UniqueLinkClicksCTR                       *go_types.Float64String `json:"unique_link_clicks_ctr"`
	UniqueOutboundClicks                      *[]AdsActionStats       `json:"unique_outbound_clicks"`
	UniqueOutboundClicksCTR                   *[]AdsActionStats       `json:"unique_outbound_clicks_ctr"`
	UniqueVideoView15Sec                      *[]AdsActionStats       `json:"unique_video_view_15_sec"`
	UpdatedTime                               *string                 `json:"updated_time"`
	Video15SecWatchedActions                  *[]AdsActionStats       `json:"video_15_sec_watched_actions"`
	Video30SecWatchedActions                  *[]AdsActionStats       `json:"video_30_sec_watched_actions"`
	VideoAsset                                json.RawMessage         `json:"video_asset"`
	VideoAvgTimeWatchedActions                *[]AdsActionStats       `json:"video_avg_time_watched_actions"`
	VideoContinuous2SecWatchedActions         *[]AdsActionStats       `json:"video_continuous_2_sec_watched_actions"`
	VideoP100WatchedActions                   *[]AdsActionStats       `json:"video_p100_watched_actions"`
	VideoP25WatchedActions                    *[]AdsActionStats       `json:"video_p25_watched_actions"`
	VideoP50WatchedActions                    *[]AdsActionStats       `json:"video_p50_watched_actions"`
	VideoP75WatchedActions                    *[]AdsActionStats       `json:"video_p75_watched_actions"`
	VideoP95WatchedActions                    *[]AdsActionStats       `json:"video_p95_watched_actions"`
	VideoPlayActions                          *[]AdsActionStats       `json:"video_play_actions"`
	VideoPlayCurveActions                     json.RawMessage         `json:"video_play_curve_actions"`
	VideoPlayRetention0To15sActions           json.RawMessage         `json:"video_play_retention_0_to_15s_actions"`
	VideoPlayRetention20To60sActions          json.RawMessage         `json:"video_play_retention_20_to_60s_actions"`
	VideoPlayRetentionGraphActions            json.RawMessage         `json:"video_play_retention_graph_actions"`
	VideoTimeWatchedActions                   *[]AdsActionStats       `json:"video_time_watched_actions"`
	WebsiteCTR                                *[]AdsActionStats       `json:"website_ctr"`
	WebsitePurchaseROAS                       *[]AdsActionStats       `json:"website_purchase_roas"`
	WishBid                                   *go_types.Float64String `json:"wish_bid"`
}

type AdsActionStats struct {
	D1Click                   *go_types.Int64String   `json:"1d_click"`
	D1View                    *go_types.Int64String   `json:"1d_view"`
	D28Click                  *go_types.Int64String   `json:"28d_click"`
	D28View                   *go_types.Int64String   `json:"28d_view"`
	D7Click                   *go_types.Int64String   `json:"7d_click"`
	D7View                    *go_types.Int64String   `json:"7d_view"`
	ActionCanvasComponentName *string                 `json:"action_canvas_component_name"`
	ActionCarouselCardID      *string                 `json:"action_carousel_card_id"`
	ActionCarouselCardName    *string                 `json:"action_carousel_card_name"`
	ActionDestination         *string                 `json:"action_destination"`
	ActionDevice              *string                 `json:"action_device"`
	ActionReaction            *string                 `json:"action_reaction"`
	ActionTargetID            *string                 `json:"action_target_id"`
	ActionType                *string                 `json:"action_type"`
	ActionVideoSound          *string                 `json:"action_video_sound"`
	ActionVideoType           *string                 `json:"action_video_type"`
	DDA                       *go_types.Float64String `json:"dda"`
	Inline                    *go_types.Float64String `json:"inline"`
	Value                     *go_types.Float64String `json:"value"`
}

type InsightsField string

const (
	InsightsFieldAccountCurrency                           InsightsField = "account_currency"
	InsightsFieldAccountID                                 InsightsField = "account_id"
	InsightsFieldAccountName                               InsightsField = "account_name"
	InsightsFieldActionValues                              InsightsField = "action_values"
	InsightsFieldActions                                   InsightsField = "actions"
	InsightsFieldActivityRecency                           InsightsField = "activity_recency"
	InsightsFieldAdClickActions                            InsightsField = "ad_click_actions"
	InsightsFieldAdFormatAsset                             InsightsField = "ad_format_asset"
	InsightsFieldAdID                                      InsightsField = "ad_id"
	InsightsFieldAdImpressionActions                       InsightsField = "ad_impression_actions"
	InsightsFieldAdName                                    InsightsField = "ad_name"
	InsightsFieldAdsetID                                   InsightsField = "adset_id"
	InsightsFieldAdsetName                                 InsightsField = "adset_name"
	InsightsFieldAgeTargeting                              InsightsField = "age_targeting"
	InsightsFieldAttributionSetting                        InsightsField = "attribution_setting"
	InsightsFieldAuctionBid                                InsightsField = "auction_bid"
	InsightsFieldAuctionCompetitiveness                    InsightsField = "auction_competitiveness"
	InsightsFieldAuctionMaxCompetitorBid                   InsightsField = "auction_max_competitor_bid"
	InsightsFieldBodyAsset                                 InsightsField = "body_asset"
	InsightsFieldBuyingType                                InsightsField = "buying_type"
	InsightsFieldCampaignID                                InsightsField = "campaign_id"
	InsightsFieldCampaignName                              InsightsField = "campaign_name"
	InsightsFieldCanvasAvgViewPercent                      InsightsField = "canvas_avg_view_percent"
	InsightsFieldCanvasAvgViewTime                         InsightsField = "canvas_avg_view_time"
	InsightsFieldCatalogSegmentActions                     InsightsField = "catalog_segment_actions"
	InsightsFieldCatalogSegmentValue                       InsightsField = "catalog_segment_value"
	InsightsFieldCatalogSegmentValueMobilePurchaseROAS     InsightsField = "catalog_segment_value_mobile_purchase_roas"
	InsightsFieldCatalogSegmentValueOmniPurchaseROAS       InsightsField = "catalog_segment_value_omni_purchase_roas"
	InsightsFieldCatalogSegmentValueWebsitePurchaseROAS    InsightsField = "catalog_segment_value_website_purchase_roas"
	InsightsFieldClicks                                    InsightsField = "clicks"
	InsightsFieldComparisonNode                            InsightsField = "comparison_node"
	InsightsFieldConversionValues                          InsightsField = "conversion_values"
	InsightsFieldConversions                               InsightsField = "conversions"
	InsightsFieldConvertedProductQuantity                  InsightsField = "converted_product_quantity"
	InsightsFieldConvertedProductValue                     InsightsField = "converted_product_value"
	InsightsFieldCostPer15SecVideoView                     InsightsField = "cost_per_15_sec_video_view"
	InsightsFieldCostPer2SecContinuousVideoView            InsightsField = "cost_per_2_sec_continuous_video_view"
	InsightsFieldCostPerActionType                         InsightsField = "cost_per_action_type"
	InsightsFieldCostPerAdClick                            InsightsField = "cost_per_ad_click"
	InsightsFieldCostPerConversion                         InsightsField = "cost_per_conversion"
	InsightsFieldCostPerDDACountByConvs                    InsightsField = "cost_per_dda_countby_convs"
	InsightsFieldCostPerInlineLinkClick                    InsightsField = "cost_per_inline_link_click"
	InsightsFieldCostPerInlinePostEngagement               InsightsField = "cost_per_inline_post_engagement"
	InsightsFieldCostPerOneThousandAdImpression            InsightsField = "cost_per_one_thousand_ad_impression"
	InsightsFieldCostPerOutboundClick                      InsightsField = "cost_per_outbound_click"
	InsightsFieldCostPerStoreVisitAction                   InsightsField = "cost_per_store_visit_action"
	InsightsFieldCostPerThruplay                           InsightsField = "cost_per_thruplay"
	InsightsFieldCostPerUniqueActionType                   InsightsField = "cost_per_unique_action_type"
	InsightsFieldCostPerUniqueClick                        InsightsField = "cost_per_unique_click"
	InsightsFieldCostPerUniqueConversion                   InsightsField = "cost_per_unique_conversion"
	InsightsFieldCostPerUniqueInlineLinkClick              InsightsField = "cost_per_unique_inline_link_click"
	InsightsFieldCostPerUniqueOutboundClick                InsightsField = "cost_per_unique_outbound_click"
	InsightsFieldCountry                                   InsightsField = "country"
	InsightsFieldCPC                                       InsightsField = "cpc"
	InsightsFieldCPM                                       InsightsField = "cpm"
	InsightsFieldCPP                                       InsightsField = "cpp"
	InsightsFieldCreatedTime                               InsightsField = "created_time"
	InsightsFieldCTR                                       InsightsField = "ctr"
	InsightsFieldDateStart                                 InsightsField = "date_start"
	InsightsFieldDateStop                                  InsightsField = "date_stop"
	InsightsFieldDDACountByConvs                           InsightsField = "dda_countby_convs"
	InsightsFieldDescriptionAsset                          InsightsField = "description_asset"
	InsightsFieldDevicePlatform                            InsightsField = "device_platform"
	InsightsFieldDMA                                       InsightsField = "dma"
	InsightsFieldEstimatedAdRecallRateLowerBound           InsightsField = "estimated_ad_recall_rate_lower_bound"
	InsightsFieldEstimatedAdRecallRateUpperBound           InsightsField = "estimated_ad_recall_rate_upper_bound"
	InsightsFieldEstimatedAdRecallersLowerBound            InsightsField = "estimated_ad_recallers_lower_bound"
	InsightsFieldEstimatedAdRecallersUpperBound            InsightsField = "estimated_ad_recallers_upper_bound"
	InsightsFieldFrequency                                 InsightsField = "frequency"
	InsightsFieldFrequencyValue                            InsightsField = "frequency_value"
	InsightsFieldFullViewImpressions                       InsightsField = "full_view_impressions"
	InsightsFieldFullViewReach                             InsightsField = "full_view_reach"
	InsightsFieldGenderTargeting                           InsightsField = "gender_targeting"
	InsightsFieldHourlyStatsAggregatedByAdvertiserTimeZone InsightsField = "hourly_stats_aggregated_by_advertiser_time_zone"
	InsightsFieldHourlyStatsAggregatedByAudienceTimeZone   InsightsField = "hourly_stats_aggregated_by_audience_time_zone"
	InsightsFieldImageAsset                                InsightsField = "image_asset"
	InsightsFieldImpressionDevice                          InsightsField = "impression_device"
	InsightsFieldImpressions                               InsightsField = "impressions"
	InsightsFieldImpressionsDummy                          InsightsField = "impressions_dummy"
	InsightsFieldInlineLinkClickCTR                        InsightsField = "inline_link_click_ctr"
	InsightsFieldInlineLinkClicks                          InsightsField = "inline_link_clicks"
	InsightsFieldInlinePostEngagement                      InsightsField = "inline_post_engagement"
	InsightsFieldInstantExperienceClicksToOpen             InsightsField = "instant_experience_clicks_to_open"
	InsightsFieldInstantExperienceClicksToStart            InsightsField = "instant_experience_clicks_to_start"
	InsightsFieldInstantExperienceOutboundClicks           InsightsField = "instant_experience_outbound_clicks"
	InsightsFieldInteractiveComponentTap                   InsightsField = "interactive_component_tap"
	InsightsFieldLabels                                    InsightsField = "labels"
	InsightsFieldLocation                                  InsightsField = "location"
	InsightsFieldMediaAsset                                InsightsField = "media_asset"
	InsightsFieldMobileAppPurchaseROAS                     InsightsField = "mobile_app_purchase_roas"
	InsightsFieldObjective                                 InsightsField = "objective"
	InsightsFieldOptimizationGoal                          InsightsField = "optimization_goal"
	InsightsFieldOutboundClicks                            InsightsField = "outbound_clicks"
	InsightsFieldOutboundClicksCTR                         InsightsField = "outbound_clicks_ctr"
	InsightsFieldPlacePageID                               InsightsField = "place_page_id"
	InsightsFieldPlacePageName                             InsightsField = "place_page_name"
	InsightsFieldPlatformPosition                          InsightsField = "platform_position"
	InsightsFieldProductID                                 InsightsField = "product_id"
	InsightsFieldPublisherPlatform                         InsightsField = "publisher_platform"
	InsightsFieldPurchaseROAS                              InsightsField = "purchase_roas"
	InsightsFieldQualifyingQuestionQualifyAnswerRate       InsightsField = "qualifying_question_qualify_answer_rate"
	InsightsFieldReach                                     InsightsField = "reach"
	InsightsFieldRuleAsset                                 InsightsField = "rule_asset"
	InsightsFieldSocialSpend                               InsightsField = "social_spend"
	InsightsFieldSpend                                     InsightsField = "spend"
	InsightsFieldStoreVisitActions                         InsightsField = "store_visit_actions"
	InsightsFieldTitleAsset                                InsightsField = "title_asset"
	InsightsFieldUniqueActions                             InsightsField = "unique_actions"
	InsightsFieldUniqueClicks                              InsightsField = "unique_clicks"
	InsightsFieldUniqueConversions                         InsightsField = "unique_conversions"
	InsightsFieldUniqueCTR                                 InsightsField = "unique_ctr"
	InsightsFieldUniqueInlineLinkClickCTR                  InsightsField = "unique_inline_link_click_ctr"
	InsightsFieldUniqueInlineLinkClicks                    InsightsField = "unique_inline_link_clicks"
	InsightsFieldUniqueLinkClicksCTR                       InsightsField = "unique_link_clicks_ctr"
	InsightsFieldUniqueOutboundClicks                      InsightsField = "unique_outbound_clicks"
	InsightsFieldUniqueOutboundClicksCTR                   InsightsField = "unique_outbound_clicks_ctr"
	InsightsFieldUniqueVideoView15Sec                      InsightsField = "unique_video_view_15_sec"
	InsightsFieldUpdatedTime                               InsightsField = "updated_time"
	InsightsFieldVideo15SecWatchedActions                  InsightsField = "video_15_sec_watched_actions"
	InsightsFieldVideo30SecWatchedActions                  InsightsField = "video_30_sec_watched_actions"
	InsightsFieldVideoAsset                                InsightsField = "video_asset"
	InsightsFieldVideoAvgTimeWatchedActions                InsightsField = "video_avg_time_watched_actions"
	InsightsFieldVideoContinuous2SecWatchedActions         InsightsField = "video_continuous_2_sec_watched_actions"
	InsightsFieldVideoP100WatchedActions                   InsightsField = "video_p100_watched_actions"
	InsightsFieldVideoP25WatchedActions                    InsightsField = "video_p25_watched_actions"
	InsightsFieldVideoP50WatchedActions                    InsightsField = "video_p50_watched_actions"
	InsightsFieldVideoP75WatchedActions                    InsightsField = "video_p75_watched_actions"
	InsightsFieldVideoP95WatchedActions                    InsightsField = "video_p95_watched_actions"
	InsightsFieldVideoPlayActions                          InsightsField = "video_play_actions"
	InsightsFieldVideoPlayCurveActions                     InsightsField = "video_play_curve_actions"
	InsightsFieldVideoPlayRetention0To15sActions           InsightsField = "video_play_retention_0_to_15s_actions"
	InsightsFieldVideoPlayRetention20To60sActions          InsightsField = "video_play_retention_20_to_60s_actions"
	InsightsFieldVideoPlayRetentionGraphActions            InsightsField = "video_play_retention_graph_actions"
	InsightsFieldVideoTimeWatchedActions                   InsightsField = "video_time_watched_actions"
	InsightsFieldWebsiteCTR                                InsightsField = "website_ctr"
	InsightsFieldWebsitePurchaseROAS                       InsightsField = "website_purchase_roas"
	InsightsFieldWishBid                                   InsightsField = "wish_bid"
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

type ActionAttributionWindow string

const (
	ActionAttributionWindow1dView   ActionAttributionWindow = "1d_view"
	ActionAttributionWindow7dView   ActionAttributionWindow = "7d_view"
	ActionAttributionWindow28dView  ActionAttributionWindow = "28d_view"
	ActionAttributionWindow1dClick  ActionAttributionWindow = "1d_click"
	ActionAttributionWindow7dClick  ActionAttributionWindow = "7d_click"
	ActionAttributionWindow28dClick ActionAttributionWindow = "28d_click"
	ActionAttributionWindowDDA      ActionAttributionWindow = "dda"
	ActionAttributionWindowDefault  ActionAttributionWindow = "default"
)

type ActionBreakdown string

const (
	ActionBreakdownDevice              ActionBreakdown = "action_device"
	ActionBreakdownCanvasComponentName ActionBreakdown = "action_canvas_component_name"
	ActionBreakdownCarouselCardID      ActionBreakdown = "action_carousel_card_id"
	ActionBreakdownCarouselCardName    ActionBreakdown = "action_carousel_card_name"
	ActionBreakdownDestination         ActionBreakdown = "action_destination"
	ActionBreakdownReaction            ActionBreakdown = "action_reaction"
	ActionBreakdownTargetID            ActionBreakdown = "action_target_id"
	ActionBreakdownType                ActionBreakdown = "action_type"
	ActionBreakdownVideoSound          ActionBreakdown = "action_video_sound"
	ActionBreakdownVideoType           ActionBreakdown = "action_video_type"
)

type Breakdown string

const (
	BreakdownAdFormatAsset                             Breakdown = "ad_format_asset"
	BreakdownAge                                       Breakdown = "age"
	BreakdownBodyAsset                                 Breakdown = "body_asset"
	BreakdownCallToActionAsset                         Breakdown = "call_to_action_asset"
	BreakdownCountry                                   Breakdown = "country"
	BreakdownDescriptionAsset                          Breakdown = "description_asset"
	BreakdownGender                                    Breakdown = "gender"
	BreakdownImageAsset                                Breakdown = "image_asset"
	BreakdownLinkURLAsset                              Breakdown = "link_url_asset"
	BreakdownProductID                                 Breakdown = "product_id"
	BreakdownRegion                                    Breakdown = "region"
	BreakdownTitleAsset                                Breakdown = "title_asset"
	BreakdownVideoAsset                                Breakdown = "video_asset"
	BreakdownDMA                                       Breakdown = "dma"
	BreakdownFrequencyValue                            Breakdown = "frequency_value"
	BreakdownHourlyStatsAggregatedByAdvertiserTimeZone Breakdown = "hourly_stats_aggregated_by_advertiser_time_zone"
	BreakdownHourlyStatsAggregatedByAudienceTimeZone   Breakdown = "hourly_stats_aggregated_by_audience_time_zone"
	BreakdownPlacePageID                               Breakdown = "place_page_id"
	BreakdownPublisherPlatform                         Breakdown = "publisher_platform"
	BreakdownPlatformPosition                          Breakdown = "platform_position"
	BreakdownDevicePlatform                            Breakdown = "device_platform"
)

type Level string

const (
	LevelAccount  Level = "account"
	LevelAd       Level = "ad"
	LevelAdSet    Level = "adset"
	LevelCampaign Level = "campaign"
)

type TimeIncrement string

const (
	TimeIncrementMonthly TimeIncrement = "monthly"
	TimeIncrementAllDays TimeIncrement = "all_days"
)

type TimeRange struct {
	Since civil.Date `json:"since"`
	Until civil.Date `json:"until"`
}

type GetInsightsConfig struct {
	ID                       int64
	ActionAttributionWindows *[]ActionAttributionWindow
	ActionBreakdowns         *[]ActionBreakdown
	Breakdowns               *[]Breakdown
	DatePreset               *DatePreset
	Fields                   []InsightsField
	Level                    *Level
	TimeIncrement            *TimeIncrement
	TimeIncrementDays        *uint // 1 to 90
	TimeRange                *TimeRange
}

func (service *Service) GetInsights(config *GetInsightsConfig) (*[]Insights, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetInsightsConfig must not be a nil pointer")
	}

	values := url.Values{}
	fields := []string{}

	if config.ActionAttributionWindows != nil {
		if len(*config.ActionAttributionWindows) > 0 {
			_actionAttributionWindows := []string{}
			for _, actionAttributionWindow := range *config.ActionAttributionWindows {
				_actionAttributionWindows = append(_actionAttributionWindows, string(actionAttributionWindow))
			}
			values.Set("action_attribution_windows", strings.Join(_actionAttributionWindows, ","))
		}
	}
	if config.ActionBreakdowns != nil {
		if len(*config.ActionBreakdowns) > 0 {
			_actionBreakdowns := []string{}
			for _, actionBreakdown := range *config.ActionBreakdowns {
				_actionBreakdowns = append(_actionBreakdowns, string(actionBreakdown))
			}
			values.Set("action_breakdowns", strings.Join(_actionBreakdowns, ","))
		}
	}
	if config.Breakdowns != nil {
		if len(*config.Breakdowns) > 0 {
			_breakdowns := []string{}
			for _, breakdown := range *config.Breakdowns {
				_breakdowns = append(_breakdowns, string(breakdown))
			}
			values.Set("breakdowns", strings.Join(_breakdowns, ","))
		}
	}
	if config.DatePreset != nil {
		values.Set("date_preset", string(*config.DatePreset))
	}
	if len(config.Fields) == 0 {
		fields = append(fields, string(InsightsFieldAdID))
	} else {
		for _, field := range config.Fields {
			fields = append(fields, string(field))
		}
	}
	if config.Level != nil {
		values.Set("level", string(*config.Level))
	}
	if config.TimeIncrement != nil && config.TimeIncrementDays != nil {
		return nil, errortools.ErrorMessage("Do not supply TimeIncrement and TimeIncrementDays at the same time")
	}
	if config.TimeIncrement != nil {
		values.Set("time_increment", string(*config.TimeIncrement))
	} else if config.TimeIncrementDays != nil {
		if *config.TimeIncrementDays < 1 || *config.TimeIncrementDays > 90 {
			return nil, errortools.ErrorMessage("TimeIncrementDays must be a number between 1 and 90")
		}
		values.Set("time_increment", fmt.Sprintf("%v", *config.TimeIncrementDays))
	}

	if config.TimeRange != nil {
		b, err := json.Marshal(config.TimeRange)
		if err != nil {
			return nil, errortools.ErrorMessagef("Error while marshaling TimeRange: %v", err.Error())
		}

		values.Set("time_range", string(b))
	}

	values.Set("fields", strings.Join(fields, ","))

	insights := []Insights{}

	url := service.url(fmt.Sprintf("%v/insights?%s", config.ID, values.Encode()))

	for {
		insightsResponse := InsightsResponse{}
		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			URL:           url,
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

		url = insightsResponse.Paging.Next
	}

	return &insights, nil
}
