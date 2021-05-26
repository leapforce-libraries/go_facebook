package facebook

type xBusinessUseCaseUsage map[string][]struct {
	Type                        string `json:"type"`
	CallCount                   int    `json:"call_count"`
	TotalCPUTime                int    `json:"total_cputime"`
	TotalTime                   int    `json:"total_time"`
	EstimatedTimeToRegainAccess int    `json:"estimated_time_to_regain_access"`
}
