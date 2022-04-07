package facebook

type Error struct {
	Message         string `json:"message"`
	Type            string `json:"type"`
	Code            int    `json:"code"`
	ErrorSubcode    int    `json:"error_subcode"`
	FacebookTraceId string `json:"fbtrace_id"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}
