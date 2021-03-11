package facebook

// ErrorResponse stores general API error response
//
type ErrorResponse struct {
	Error struct {
		Message         string `json:"message"`
		Type            string `json:"type"`
		Code            int    `json:"code"`
		ErrorSubcode    int    `json:"error_subcode"`
		FacebookTraceID string `json:"fbtrace_id"`
	} `json:"error"`
}
