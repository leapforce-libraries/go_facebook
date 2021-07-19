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

type CreativeResponse struct {
	Data   []Creative `json:"data"`
	Paging *Paging    `json:"paging"`
}

type Creative struct {
	ID                             go_types.Int64String    `json:"id"`
	AccountID                      *go_types.Int64String   `json:"account_id"`
	ActorID                        *go_types.Int64String   `json:"actor_id"`
	AdLabels                       json.RawMessage         `json:"adlabels"`
	ApplinkTreatment               *string                 `json:"applink_treatment"`
	AssetFeedSpec                  json.RawMessage         `json:"asset_feed_spec"`
	AuthorizationCategory          *string                 `json:"authorization_category"`
	Body                           *string                 `json:"body"`
	BrandedContentSponsorPageID    *go_types.Int64String   `json:"branded_content_sponsor_page_id"`
	BundleFolderID                 *go_types.Int64String   `json:"bundle_folder_id"`
	CallToActionType               *string                 `json:"call_to_action_type"`
	CategorizationCriteria         *string                 `json:"categorization_criteria"`
	CategoryMediaSource            *string                 `json:"category_media_source"`
	DestinationSetID               *[]go_types.Int64String `json:"destination_set_id"`
	DynamicAdVoice                 *string                 `json:"dynamic_ad_voice"`
	EffectiveAuthorizationCategory *string                 `json:"effective_authorization_category"`
	EffectiveInstagramMediaID      *go_types.Int64String   `json:"effective_instagram_media_id"`
	EffectiveInstagramStoryID      *go_types.Int64String   `json:"effective_instagram_story_id"`
	EffectiveObjectStoryID         *string                 `json:"effective_object_story_id"`
	EnableDirectInstall            *bool                   `json:"enable_direct_install"`
	EnableLaunchInstantApp         *bool                   `json:"enable_launch_instant_app"`
	ImageCrops                     json.RawMessage         `json:"image_crops"`
	ImageHash                      *string                 `json:"image_hash"`
	ImageURL                       *string                 `json:"image_url"`
	InstagramActorID               *go_types.Int64String   `json:"instagram_actor_id"`
	InstagramPermalinkURL          *string                 `json:"instagram_permalink_url"`
	InstagramStoryID               *go_types.Int64String   `json:"instagram_story_id"`
	InstagramUserID                *go_types.Int64String   `json:"instagram_user_id"`
	InteractiveComponentsSpec      json.RawMessage         `json:"interactive_components_spec"`
	LinkDestinationDisplayURL      *string                 `json:"link_destination_display_url"`
	LinkOpenGraphID                *string                 `json:"link_og_id"`
	LinkURL                        *string                 `json:"link_url"`
	MessengerSponsoredMessage      *string                 `json:"messenger_sponsored_message"`
	Name                           *string                 `json:"name"`
	ObjectID                       *go_types.Int64String   `json:"object_id"`
	ObjectStoreURL                 *string                 `json:"object_store_url"`
	ObjectStoryID                  *string                 `json:"object_story_id"`
	ObjectStorySpec                json.RawMessage         `json:"object_story_spec"`
	ObjectType                     *string                 `json:"object_type"`
	ObjectURL                      *string                 `json:"object_url"`
	PlacePageSetID                 *go_types.Int64String   `json:"place_page_set_id"`
	PlatformCustomizations         json.RawMessage         `json:"platform_customizations"`
	PlayableAssetID                *go_types.Int64String   `json:"playable_asset_id"`
	PortraitCustomizations         json.RawMessage         `json:"portrait_customizations"`
	ProductSetID                   *go_types.Int64String   `json:"product_set_id"`
	RecommenderSettings            json.RawMessage         `json:"recommender_settings"`
	ReferralID                     *go_types.Int64String   `json:"referral_id"`
	SourceInstagramMediaID         *go_types.Int64String   `json:"source_instagram_media_id"`
	Status                         *string                 `json:"status"`
	TemplateURL                    *string                 `json:"template_url"`
	TemplateURLSpec                json.RawMessage         `json:"template_url_spec"`
	ThumbnailURL                   *string                 `json:"thumbnail_url"`
	Title                          *string                 `json:"title"`
	URLTags                        *string                 `json:"url_tags"`
	UsePageActorOverride           *bool                   `json:"use_page_actor_override"`
	VideoID                        *go_types.Int64String   `json:"video_id"`
}

type CreativeField string

const (
	CreativeFieldID                             CreativeField = "id"
	CreativeFieldAccountID                      CreativeField = "account_id"
	CreativeFieldActorID                        CreativeField = "actor_id"
	CreativeFieldAdLabels                       CreativeField = "adlabels"
	CreativeFieldApplinkTreatment               CreativeField = "applink_treatment"
	CreativeFieldAssetFeedSpec                  CreativeField = "asset_feed_spec"
	CreativeFieldAuthorizationCategory          CreativeField = "authorization_category"
	CreativeFieldBody                           CreativeField = "body"
	CreativeFieldBrandedContentSponsorPageID    CreativeField = "branded_content_sponsor_page_id"
	CreativeFieldBundleFolderID                 CreativeField = "bundle_folder_id"
	CreativeFieldCallToActionType               CreativeField = "call_to_action_type"
	CreativeFieldCategorizationCriteria         CreativeField = "categorization_criteria"
	CreativeFieldCategoryMediaSource            CreativeField = "category_media_source"
	CreativeFieldDestinationSetID               CreativeField = "destination_set_id"
	CreativeFieldDynamicAdVoice                 CreativeField = "dynamic_ad_voice"
	CreativeFieldEffectiveAuthorizationCategory CreativeField = "effective_authorization_category"
	CreativeFieldEffectiveInstagramMediaID      CreativeField = "effective_instagram_media_id"
	CreativeFieldEffectiveInstagramStoryID      CreativeField = "effective_instagram_story_id"
	CreativeFieldEffectiveObjectStoryID         CreativeField = "effective_object_story_id"
	CreativeFieldEnableDirectInstall            CreativeField = "enable_direct_install"
	CreativeFieldEnableLaunchInstantApp         CreativeField = "enable_launch_instant_app"
	CreativeFieldImageCrops                     CreativeField = "image_crops"
	CreativeFieldImageHash                      CreativeField = "image_hash"
	CreativeFieldImageURL                       CreativeField = "image_url"
	CreativeFieldInstagramActorID               CreativeField = "instagram_actor_id"
	CreativeFieldInstagramPermalinkURL          CreativeField = "instagram_permalink_url"
	CreativeFieldInstagramStoryID               CreativeField = "instagram_story_id"
	CreativeFieldInstagramUserID                CreativeField = "instagram_user_id"
	CreativeFieldInteractiveComponentsSpec      CreativeField = "interactive_components_spec"
	CreativeFieldLinkDestinationDisplayURL      CreativeField = "link_destination_display_url"
	CreativeFieldLinkOpenGraphID                CreativeField = "link_og_id"
	CreativeFieldLinkURL                        CreativeField = "link_url"
	CreativeFieldMessengerSponsoredMessage      CreativeField = "messenger_sponsored_message"
	CreativeFieldName                           CreativeField = "name"
	CreativeFieldObjectID                       CreativeField = "object_id"
	CreativeFieldObjectStoreURL                 CreativeField = "object_store_url"
	CreativeFieldObjectStoryID                  CreativeField = "object_story_id"
	CreativeFieldObjectStorySpec                CreativeField = "object_story_spec"
	CreativeFieldObjectType                     CreativeField = "object_type"
	CreativeFieldObjectURL                      CreativeField = "object_url"
	CreativeFieldPlacePageSetID                 CreativeField = "place_page_set_id"
	CreativeFieldPlatformCustomizations         CreativeField = "platform_customizations"
	CreativeFieldPlayableAssetID                CreativeField = "playable_asset_id"
	CreativeFieldPortraitCustomizations         CreativeField = "portrait_customizations"
	CreativeFieldProductSetID                   CreativeField = "product_set_id"
	CreativeFieldRecommenderSettings            CreativeField = "recommender_settings"
	CreativeFieldReferralID                     CreativeField = "referral_id"
	CreativeFieldSourceInstagramMediaID         CreativeField = "source_instagram_media_id"
	CreativeFieldStatus                         CreativeField = "status"
	CreativeFieldTemplateURL                    CreativeField = "template_url"
	CreativeFieldTemplateURLSpec                CreativeField = "template_url_spec"
	CreativeFieldThumbnailURL                   CreativeField = "thumbnail_url"
	CreativeFieldTitle                          CreativeField = "title"
	CreativeFieldURLTags                        CreativeField = "url_tags"
	CreativeFieldUsePageActorOverride           CreativeField = "use_page_actor_override"
	CreativeFieldVideoID                        CreativeField = "video_id"
)

type GetCreativesConfig struct {
	AccountID int64
	Fields    []CreativeField
}

func (service *Service) GetCreatives(config *GetCreativesConfig) (*[]Creative, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetCreativesConfig must not be a nil pointer")
	}

	values := url.Values{}
	fields := []string{}
	if len(config.Fields) == 0 {
		fields = append(fields, string(CreativeFieldID))
	} else {
		for _, field := range config.Fields {
			fields = append(fields, string(field))
		}
	}
	values.Set("fields", strings.Join(fields, ","))

	creatives := []Creative{}

	url := service.url(fmt.Sprintf("act_%v/adcreatives?%s", config.AccountID, values.Encode()))

	for {
		creativeResponse := CreativeResponse{}
		requestConfig := go_http.RequestConfig{
			URL:           url,
			ResponseModel: &creativeResponse,
		}

		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		creatives = append(creatives, creativeResponse.Data...)

		if creativeResponse.Paging == nil {
			break
		}

		if creativeResponse.Paging.Next == "" {
			break
		}

		url = creativeResponse.Paging.Next
	}

	return &creatives, nil
}

type GetCreativeConfig struct {
	CreativeID int64
	Fields     []CreativeField
}

func (service *Service) GetCreative(config *GetCreativeConfig) (*Creative, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetCreativeConfig must not be a nil pointer")
	}

	values := url.Values{}
	fields := []string{}
	if len(config.Fields) == 0 {
		fields = append(fields, string(CreativeFieldID))
	} else {
		for _, field := range config.Fields {
			fields = append(fields, string(field))
		}
	}
	values.Set("fields", strings.Join(fields, ","))

	creative := Creative{}

	url := service.url(fmt.Sprintf("%v?%s", config.CreativeID, values.Encode()))
	requestConfig := go_http.RequestConfig{
		URL:           url,
		ResponseModel: &creative,
	}

	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &creative, nil
}
