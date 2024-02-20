package facebook

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

const (
	maxBatchSize int = 50
)

type BatchResponse struct {
	Code    int64 `json:"code"`
	Headers []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"headers"`
	Body string `json:"body"`
}

type BatchConfig struct {
	requestConfigs []*go_http.RequestConfig
	IncludeHeaders *bool
	AccessToken    string
}

type BatchEntry struct {
	Method      string `json:"method"`
	RelativeUrl string `json:"relative_url"`
}

func (batchConfig BatchConfig) IsFull() bool {
	return len(batchConfig.requestConfigs) == maxBatchSize
}

func (batchConfig *BatchConfig) Clear() {
	batchConfig.requestConfigs = []*go_http.RequestConfig{}
}

func (batchConfig *BatchConfig) AddRequest(requestConfig *go_http.RequestConfig) *errortools.Error {
	if batchConfig.IsFull() {
		return errortools.ErrorMessagef("Batch is full (max %v requests)", maxBatchSize)
	}

	batchConfig.requestConfigs = append(batchConfig.requestConfigs, requestConfig)

	return nil
}

func (service *Service) Batch(config *BatchConfig) (*[]BatchResponse, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("BatchConfig must not be a nil pointer")
	}

	batch := []BatchEntry{}
	for _, requestConfig := range config.requestConfigs {
		if requestConfig == nil {
			continue
		}

		if requestConfig.RelativeUrl == "" {
			return nil, errortools.ErrorMessage("RequestConfig.RelativeUrl should be entered for Batch endpoint")
		}

		batch = append(batch, BatchEntry{
			Method:      requestConfig.Method,
			RelativeUrl: requestConfig.RelativeUrl,
		})
	}

	if len(batch) == 0 {
		return nil, nil
	}

	b, err := json.Marshal(batch)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	//values := urlV18.Values{}
	values := []string{
		fmt.Sprintf("batch=%s", string(b)),
		fmt.Sprintf("access_token=%s", config.AccessToken),
	} // we do not use urlV18.Values since the "batch" param should not be encoded
	if config.IncludeHeaders != nil {
		values = append(values, fmt.Sprintf("include_headers=%v", *config.IncludeHeaders))
	}

	response := []BatchResponse{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           fmt.Sprintf("https://graph.facebook.com/me?%s", strings.Join(values, "&")),
		ResponseModel: &response,
	}
	//fmt.Println(requestConfig.Url)
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	if len(response) != len(config.requestConfigs) {
		return nil, errortools.ErrorMessagef("Batch call returned %v responses while %v calls were passed", len(response), len(config.requestConfigs))
	}

	for i, res := range response {
		err := json.Unmarshal([]byte(res.Body), config.requestConfigs[i].ResponseModel)
		if err != nil {
			return nil, errortools.ErrorMessage(err)
		}
	}

	return &response, nil
}
