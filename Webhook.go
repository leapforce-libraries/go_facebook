package facebook

import "encoding/json"

type Webhook struct {
	Object string         `json:"object"`
	Entry  []WebhookEntry `json:"entry"`
}

type WebhookEntry struct {
	Id      string          `json:"id"`
	Changes []WebhookChange `json:"changes"`
	Time    int64           `json:"time"`
}

type WebhookChange struct {
	Field string          `json:"field"`
	Value json.RawMessage `json:"value"`
}
