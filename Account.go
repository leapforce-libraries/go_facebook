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

type Account struct {
	ID                                AccountID               `json:"id"`
	AccountID                         *go_types.Int64String   `json:"account_id"`
	AccountStatus                     *AccountStatus          `json:"account_status"`
	AdAccountPromotableObjects        json.RawMessage         `json:"ad_account_promotable_objects"`
	Age                               *float64                `json:"age"`
	AgencyClientDeclaration           json.RawMessage         `json:"agency_client_declaration"`
	AmountSpent                       *go_types.Int64String   `json:"amount_spent"`
	AttributionSpec                   json.RawMessage         `json:"attribution_spec"`
	Balance                           *go_types.Int64String   `json:"balance"`
	Business                          json.RawMessage         `json:"business"`
	BusinessCity                      *string                 `json:"business_city"`
	BusinessCountryCode               *string                 `json:"business_country_code"`
	BusinessName                      *string                 `json:"business_name"`
	BusinessState                     *string                 `json:"business_state"`
	BusinessStreet                    *string                 `json:"business_street"`
	BusinessStreet2                   *string                 `json:"business_street2"`
	BusinessZIP                       *string                 `json:"business_zip"`
	CanCreateBrandLiftStudy           *bool                   `json:"can_create_brand_lift_study"`
	Capabilities                      *[]string               `json:"capabilities"`
	CreatedTime                       *f_types.DateTimeString `json:"created_time"`
	Currency                          *string                 `json:"currency"`
	DirectDealsToSAccepted            *bool                   `json:"direct_deals_tos_accepted"`
	DisableReason                     *uint32                 `json:"disable_reason"`
	EndAdvertiser                     *go_types.Int64String   `json:"end_advertiser"`
	EndAdvertiserName                 *string                 `json:"end_advertiser_name"`
	ExtendedCreditInvoiceGroup        json.RawMessage         `json:"extended_credit_invoice_group"`
	FacebookEntity                    json.RawMessage         `json:"fb_entity"`
	FundingSource                     *go_types.Int64String   `json:"funding_source"`
	FundingSourceDetails              json.RawMessage         `json:"funding_source_details"`
	HasMigratedPermissions            *bool                   `json:"has_migrated_permissions"`
	HasPageAuthorizedAdAccount        *bool                   `json:"has_page_authorized_adaccount"`
	InsertionOrderNumber              *go_types.Int64String   `json:"io_number"`
	IsAttributionSpecSystemDefault    *bool                   `json:"is_attribution_spec_system_default"`
	IsDirectDealsEnabled              *bool                   `json:"is_direct_deals_enabled"`
	IsIn3dsAuthorizationEnabledMarket *bool                   `json:"is_in_3ds_authorization_enabled_market"`
	IsInMiddleOfLocalEntityMigration  *bool                   `json:"is_in_middle_of_local_entity_migration"`
	IsNotificationsEnabled            *bool                   `json:"is_notifications_enabled"`
	IsPersonal                        *uint32                 `json:"is_personal"`
	IsPrepayAccount                   *bool                   `json:"is_prepay_account"`
	IsTaxIDRequired                   *bool                   `json:"is_tax_id_required"`
	LineNumbers                       *[]int                  `json:"line_numbers"`
	MediaAgency                       *go_types.Int64String   `json:"media_agency"`
	MinCampaignGroupSpendCapap        *go_types.Int64String   `json:"min_campaign_group_spend_cap"`
	MinDailyBudget                    *uint32                 `json:"min_daily_budget"`
	Name                              *string                 `json:"name"`
	OffsitePixelsToSAccepted          *bool                   `json:"offsite_pixels_tos_accepted"`
	Owner                             *go_types.Int64String   `json:"owner"`
	Partner                           *go_types.Int64String   `json:"partner"`
	ReachFrequencySpec                json.RawMessage         `json:"rf_spec"`
	ShowCheckoutExperience            *bool                   `json:"show_checkout_experience"`
	SpendCap                          *go_types.Int64String   `json:"spend_cap"`
	TaxID                             *string                 `json:"tax_id"`
	TaxIDStatus                       *uint32                 `json:"tax_id_status"`
	TaxIDType                         *string                 `json:"tax_id_type"`
	TimezoneID                        *uint32                 `json:"timezone_id"`
	TimezoneName                      *string                 `json:"timezone_name"`
	TimezoneOffsetHoursUTC            *float64                `json:"timezone_offset_hours_utc"`
	ToSAccepted                       *map[string]int32       `json:"tos_accepted"`
	UsersToSAccepted                  *map[string]int32       `json:"user_tos_accepted"`
}

type AccountField string

const (
	AccountFieldID                                AccountField = "id"
	AccountFieldAccountID                         AccountField = "account_id"
	AccountFieldAccountStatus                     AccountField = "account_status"
	AccountFieldAdAccountPromotableObjects        AccountField = "ad_account_promotable_objects"
	AccountFieldAge                               AccountField = "age"
	AccountFieldAgencyClientDeclaration           AccountField = "agency_client_declaration"
	AccountFieldAmountSpent                       AccountField = "amount_spent"
	AccountFieldAttributionSpec                   AccountField = "attribution_spec"
	AccountFieldBalance                           AccountField = "balance"
	AccountFieldBusiness                          AccountField = "business"
	AccountFieldBusinessCity                      AccountField = "business_city"
	AccountFieldBusinessCountryCode               AccountField = "business_country_code"
	AccountFieldBusinessName                      AccountField = "business_name"
	AccountFieldBusinessState                     AccountField = "business_state"
	AccountFieldBusinessStreet                    AccountField = "business_street"
	AccountFieldBusinessStreet2                   AccountField = "business_street2"
	AccountFieldBusinessZIP                       AccountField = "business_zip"
	AccountFieldCanCreateBrandLiftStudy           AccountField = "can_create_brand_lift_study"
	AccountFieldCapabilities                      AccountField = "capabilities"
	AccountFieldCreatedTime                       AccountField = "created_time"
	AccountFieldCurrency                          AccountField = "currency"
	AccountFieldDirectDealsToSAccepted            AccountField = "direct_deals_tos_accepted"
	AccountFieldDisableReason                     AccountField = "disable_reason"
	AccountFieldEndAdvertiser                     AccountField = "end_advertiser"
	AccountFieldEndAdvertiserName                 AccountField = "end_advertiser_name"
	AccountFieldExtendedCreditInvoiceGroup        AccountField = "extended_credit_invoice_group"
	AccountFieldFacebookEntity                    AccountField = "fb_entity"
	AccountFieldFundingSource                     AccountField = "funding_source"
	AccountFieldFundingSourceDetails              AccountField = "funding_source_details"
	AccountFieldHasMigratedPermissions            AccountField = "has_migrated_permissions"
	AccountFieldHasPageAuthorizedAdAccount        AccountField = "has_page_authorized_adaccount"
	AccountFieldInsertionOrderNumber              AccountField = "io_number"
	AccountFieldIsAttributionSpecSystemDefault    AccountField = "is_attribution_spec_system_default"
	AccountFieldIsDirectDealsEnabled              AccountField = "is_direct_deals_enabled"
	AccountFieldIsIn3dsAuthorizationEnabledMarket AccountField = "is_in_3ds_authorization_enabled_market"
	AccountFieldIsInMiddleOfLocalEntityMigration  AccountField = "is_in_middle_of_local_entity_migration"
	AccountFieldIsNotificationsEnabled            AccountField = "is_notifications_enabled"
	AccountFieldIsPersonal                        AccountField = "is_personal"
	AccountFieldIsPrepayAccount                   AccountField = "is_prepay_account"
	AccountFieldIsTaxIDRequired                   AccountField = "is_tax_id_required"
	AccountFieldLineNumbers                       AccountField = "line_numbers"
	AccountFieldMediaAgency                       AccountField = "media_agency"
	AccountFieldMinCampaignGroupSpendCapap        AccountField = "min_campaign_group_spend_cap"
	AccountFieldMinDailyBudget                    AccountField = "min_daily_budget"
	AccountFieldName                              AccountField = "name"
	AccountFieldOffsitePixelsToSAccepted          AccountField = "offsite_pixels_tos_accepted"
	AccountFieldOwner                             AccountField = "owner"
	AccountFieldPartner                           AccountField = "partner"
	AccountFieldReachFrequencySpec                AccountField = "rf_spec"
	AccountFieldShowCheckoutExperience            AccountField = "show_checkout_experience"
	AccountFieldSpendCap                          AccountField = "spend_cap"
	AccountFieldTaxID                             AccountField = "tax_id"
	AccountFieldTaxIDStatus                       AccountField = "tax_id_status"
	AccountFieldTaxIDType                         AccountField = "tax_id_type"
	AccountFieldTimezoneID                        AccountField = "timezone_id"
	AccountFieldTimezoneName                      AccountField = "timezone_name"
	AccountFieldTimezoneOffsetHoursUTC            AccountField = "timezone_offset_hours_utc"
	AccountFieldToSAccepted                       AccountField = "tos_accepted"
	AccountFieldUsersToSAccepted                  AccountField = "user_tos_accepted"
)

type GetAccountConfig struct {
	AccountID int64
	Fields    *[]AccountField
}

func (service *Service) GetAccount(config *GetAccountConfig) (*Account, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetAccountsConfig must not be a nil pointer")
	}

	values := url.Values{}
	fields := []string{string(AccountFieldID)}
	if config.Fields != nil {
		for _, field := range *config.Fields {
			fields = append(fields, string(field))
		}
	}
	values.Set("fields", strings.Join(fields, ","))

	account := Account{}
	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		URL:           service.url(fmt.Sprintf("act_%v?%s", config.AccountID, values.Encode())),
		ResponseModel: &account,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &account, nil
}
