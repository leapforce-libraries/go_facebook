package facebook

type IgWebhookStoryInsights struct {
	MediaId     string `json:"media_id"`
	Impressions *int64 `json:"impressions"`
	Reach       *int64 `json:"reach"`
	TapsForward *int64 `json:"taps_forward"`
	TapsBack    *int64 `json:"taps_back"`
	Exits       *int64 `json:"exits"`
	Replies     *int64 `json:"replies"`
}
