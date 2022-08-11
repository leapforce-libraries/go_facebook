package facebook

type Error struct {
	Message         string `json:"message"`
	Type            string `json:"type"`
	Code            int    `json:"code"`
	ErrorSubcode    int    `json:"error_subcode"`
	IsTransient     int    `json:"is_transient"`
	ErrorUserTitle  string `json:"error_user_title"`
	ErrorUserMsg    string `json:"error_user_msg"`
	FacebookTraceID string `json:"fbtrace_id"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}
