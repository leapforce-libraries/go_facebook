package facebook

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	Id                             go_types.Int64String    `json:"id"`
	AccountId                      *go_types.Int64String   `json:"account_id"`
	ActorId                        *go_types.Int64String   `json:"actor_id"`
	AdLabels                       json.RawMessage         `json:"adlabels"`
	ApplinkTreatment               *string                 `json:"applink_treatment"`
	AssetFeedSpec                  json.RawMessage         `json:"asset_feed_spec"`
	AuthorizationCategory          *string                 `json:"authorization_category"`
	Body                           *string                 `json:"body"`
	BrandedContentSponsorPageId    *go_types.Int64String   `json:"branded_content_sponsor_page_id"`
	BundleFolderId                 *go_types.Int64String   `json:"bundle_folder_id"`
	CallToActionType               *string                 `json:"call_to_action_type"`
	CategorizationCriteria         *string                 `json:"categorization_criteria"`
	CategoryMediaSource            *string                 `json:"category_media_source"`
	DestinationSetId               *[]go_types.Int64String `json:"destination_set_id"`
	DynamicAdVoice                 *string                 `json:"dynamic_ad_voice"`
	EffectiveAuthorizationCategory *string                 `json:"effective_authorization_category"`
	EffectiveInstagramMediaId      *go_types.Int64String   `json:"effective_instagram_media_id"`
	EffectiveInstagramStoryId      *go_types.Int64String   `json:"effective_instagram_story_id"`
	EffectiveObjectStoryId         *string                 `json:"effective_object_story_id"`
	EnableDirectInstall            *bool                   `json:"enable_direct_install"`
	EnableLaunchInstantApp         *bool                   `json:"enable_launch_instant_app"`
	ImageCrops                     json.RawMessage         `json:"image_crops"`
	ImageHash                      *string                 `json:"image_hash"`
	ImageUrl                       *string                 `json:"image_url"`
	InstagramActorId               *go_types.Int64String   `json:"instagram_actor_id"`
	InstagramPermalinkUrl          *string                 `json:"instagram_permalink_url"`
	InstagramStoryId               *go_types.Int64String   `json:"instagram_story_id"`
	InstagramUserId                *go_types.Int64String   `json:"instagram_user_id"`
	InteractiveComponentsSpec      json.RawMessage         `json:"interactive_components_spec"`
	LinkDestinationDisplayUrl      *string                 `json:"link_destination_display_url"`
	LinkOpenGraphId                *string                 `json:"link_og_id"`
	LinkUrl                        *string                 `json:"link_url"`
	MessengerSponsoredMessage      *string                 `json:"messenger_sponsored_message"`
	Name                           *string                 `json:"name"`
	ObjectId                       *go_types.Int64String   `json:"object_id"`
	ObjectStoreUrl                 *string                 `json:"object_store_url"`
	ObjectStoryId                  *string                 `json:"object_story_id"`
	ObjectStorySpec                json.RawMessage         `json:"object_story_spec"`
	ObjectType                     *string                 `json:"object_type"`
	ObjectUrl                      *string                 `json:"object_url"`
	PlacePageSetId                 *go_types.Int64String   `json:"place_page_set_id"`
	PlatformCustomizations         json.RawMessage         `json:"platform_customizations"`
	PlayableAssetId                *go_types.Int64String   `json:"playable_asset_id"`
	PortraitCustomizations         json.RawMessage         `json:"portrait_customizations"`
	ProductSetId                   *go_types.Int64String   `json:"product_set_id"`
	RecommenderSettings            json.RawMessage         `json:"recommender_settings"`
	ReferralId                     *go_types.Int64String   `json:"referral_id"`
	SourceInstagramMediaId         *go_types.Int64String   `json:"source_instagram_media_id"`
	Status                         *string                 `json:"status"`
	TemplateUrl                    *string                 `json:"template_url"`
	TemplateUrlSpec                json.RawMessage         `json:"template_url_spec"`
	ThumbnailUrl                   *string                 `json:"thumbnail_url"`
	Title                          *string                 `json:"title"`
	UrlTags                        *string                 `json:"url_tags"`
	UsePageActorOverride           *bool                   `json:"use_page_actor_override"`
	VideoId                        *go_types.Int64String   `json:"video_id"`
}

type CreativeField string

const (
	CreativeFieldId                             CreativeField = "id"
	CreativeFieldAccountId                      CreativeField = "account_id"
	CreativeFieldActorId                        CreativeField = "actor_id"
	CreativeFieldAdLabels                       CreativeField = "adlabels"
	CreativeFieldApplinkTreatment               CreativeField = "applink_treatment"
	CreativeFieldAssetFeedSpec                  CreativeField = "asset_feed_spec"
	CreativeFieldAuthorizationCategory          CreativeField = "authorization_category"
	CreativeFieldBody                           CreativeField = "body"
	CreativeFieldBrandedContentSponsorPageId    CreativeField = "branded_content_sponsor_page_id"
	CreativeFieldBundleFolderId                 CreativeField = "bundle_folder_id"
	CreativeFieldCallToActionType               CreativeField = "call_to_action_type"
	CreativeFieldCategorizationCriteria         CreativeField = "categorization_criteria"
	CreativeFieldCategoryMediaSource            CreativeField = "category_media_source"
	CreativeFieldDestinationSetId               CreativeField = "destination_set_id"
	CreativeFieldDynamicAdVoice                 CreativeField = "dynamic_ad_voice"
	CreativeFieldEffectiveAuthorizationCategory CreativeField = "effective_authorization_category"
	CreativeFieldEffectiveInstagramMediaId      CreativeField = "effective_instagram_media_id"
	CreativeFieldEffectiveInstagramStoryId      CreativeField = "effective_instagram_story_id"
	CreativeFieldEffectiveObjectStoryId         CreativeField = "effective_object_story_id"
	CreativeFieldEnableDirectInstall            CreativeField = "enable_direct_install"
	CreativeFieldEnableLaunchInstantApp         CreativeField = "enable_launch_instant_app"
	CreativeFieldImageCrops                     CreativeField = "image_crops"
	CreativeFieldImageHash                      CreativeField = "image_hash"
	CreativeFieldImageUrl                       CreativeField = "image_url"
	CreativeFieldInstagramActorId               CreativeField = "instagram_actor_id"
	CreativeFieldInstagramPermalinkUrl          CreativeField = "instagram_permalink_url"
	CreativeFieldInstagramStoryId               CreativeField = "instagram_story_id"
	CreativeFieldInstagramUserId                CreativeField = "instagram_user_id"
	CreativeFieldInteractiveComponentsSpec      CreativeField = "interactive_components_spec"
	CreativeFieldLinkDestinationDisplayUrl      CreativeField = "link_destination_display_url"
	CreativeFieldLinkOpenGraphId                CreativeField = "link_og_id"
	CreativeFieldLinkUrl                        CreativeField = "link_url"
	CreativeFieldMessengerSponsoredMessage      CreativeField = "messenger_sponsored_message"
	CreativeFieldName                           CreativeField = "name"
	CreativeFieldObjectId                       CreativeField = "object_id"
	CreativeFieldObjectStoreUrl                 CreativeField = "object_store_url"
	CreativeFieldObjectStoryId                  CreativeField = "object_story_id"
	CreativeFieldObjectStorySpec                CreativeField = "object_story_spec"
	CreativeFieldObjectType                     CreativeField = "object_type"
	CreativeFieldObjectUrl                      CreativeField = "object_url"
	CreativeFieldPlacePageSetId                 CreativeField = "place_page_set_id"
	CreativeFieldPlatformCustomizations         CreativeField = "platform_customizations"
	CreativeFieldPlayableAssetId                CreativeField = "playable_asset_id"
	CreativeFieldPortraitCustomizations         CreativeField = "portrait_customizations"
	CreativeFieldProductSetId                   CreativeField = "product_set_id"
	CreativeFieldRecommenderSettings            CreativeField = "recommender_settings"
	CreativeFieldReferralId                     CreativeField = "referral_id"
	CreativeFieldSourceInstagramMediaId         CreativeField = "source_instagram_media_id"
	CreativeFieldStatus                         CreativeField = "status"
	CreativeFieldTemplateUrl                    CreativeField = "template_url"
	CreativeFieldTemplateUrlSpec                CreativeField = "template_url_spec"
	CreativeFieldThumbnailUrl                   CreativeField = "thumbnail_url"
	CreativeFieldTitle                          CreativeField = "title"
	CreativeFieldUrlTags                        CreativeField = "url_tags"
	CreativeFieldUsePageActorOverride           CreativeField = "use_page_actor_override"
	CreativeFieldVideoId                        CreativeField = "video_id"
)

type GetCreativesConfig struct {
	AccountId int64
	Fields    []CreativeField
}

func (service *Service) GetCreatives(config *GetCreativesConfig) (*[]Creative, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetCreativesConfig must not be a nil pointer")
	}

	values := url.Values{}
	fields := []string{}
	if len(config.Fields) == 0 {
		fields = append(fields, string(CreativeFieldId))
	} else {
		for _, field := range config.Fields {
			fields = append(fields, string(field))
		}
	}
	values.Set("fields", strings.Join(fields, ","))

	creatives := []Creative{}

	url := service.urlV20(fmt.Sprintf("act_%v/adcreatives?%s", config.AccountId, values.Encode()))

	for {
		creativeResponse := CreativeResponse{}
		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           url,
			ResponseModel: &creativeResponse,
		}

		_, _, e := service.httpRequest(&requestConfig)
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
	CreativeId int64
	Fields     []CreativeField
}

func (service *Service) GetCreative(config *GetCreativeConfig) (*Creative, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetCreativeConfig must not be a nil pointer")
	}

	values := url.Values{}
	fields := []string{}
	if len(config.Fields) == 0 {
		fields = append(fields, string(CreativeFieldId))
	} else {
		for _, field := range config.Fields {
			fields = append(fields, string(field))
		}
	}
	values.Set("fields", strings.Join(fields, ","))

	creative := Creative{}

	url := service.urlV20(fmt.Sprintf("%v?%s", config.CreativeId, values.Encode()))
	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           url,
		ResponseModel: &creative,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &creative, nil
}
