package facebook

import (
	"bytes"
	"encoding/json"
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

type InitVideoUploadSessionConfig struct {
	PageId          string
	FileSize        int64
	PageAccessToken string
}

type InitVideoUploadSessionResponse struct {
	VideoId         string               `json:"video_id"`
	StartOffset     go_types.Int64String `json:"start_offset"`
	EndOffset       go_types.Int64String `json:"end_offset"`
	UploadSessionId string               `json:"upload_session_id"`
}

func (service *Service) InitVideoUploadSession(config *InitVideoUploadSessionConfig) (*InitVideoUploadSessionResponse, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("InitVideoUploadSessionConfig must not be a nil pointer")
	}

	var values = url.Values{}
	values.Set("upload_phase", "start")
	values.Set("access_token", config.PageAccessToken)
	values.Set("file_size", fmt.Sprintf("%v", config.FileSize))

	var response InitVideoUploadSessionResponse

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.urlV20(fmt.Sprintf("%s/videos?%s", config.PageId, values.Encode())),
		ResponseModel: &response,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &response, nil
}

type UploadVideoChunkConfig struct {
	PageId          string
	UploadSessionId string
	StartOffset     int64
	VideoFileChunk  string
	PageAccessToken string
}

type UploadVideoChunkResponse struct {
	StartOffset go_types.Int64String `json:"start_offset"`
	EndOffset   go_types.Int64String `json:"end_offset"`
}

func (service *Service) UploadVideoChunk(config *UploadVideoChunkConfig) (*UploadVideoChunkResponse, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("UploadVideoChunkResponse must not be a nil pointer")
	}

	file, _ := os.Open(config.VideoFileChunk)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("upload_phase", "transfer")
	_ = writer.WriteField("upload_session_id", config.UploadSessionId)
	_ = writer.WriteField("access_token", config.PageAccessToken)
	_ = writer.WriteField("start_offset", fmt.Sprintf("%v", config.StartOffset))

	part, _ := writer.CreateFormFile("video_file_chunk", file.Name())
	io.Copy(part, file)
	writer.Close()

	r, err := http.NewRequest("POST", service.urlV20(fmt.Sprintf("%s/videos", config.PageId)), body)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	r.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	re, err := client.Do(r)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	var response UploadVideoChunkResponse

	b, err := io.ReadAll(re.Body)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	return &response, nil
}

type EndVideoUploadSessionConfig struct {
	PageId          string
	UploadSessionId string
	PageAccessToken string
	Description     *string
	Title           *string
	Published       *bool
}

type EndVideoUploadSessionResponse struct {
	Success bool `json:"success"`
}

func (service *Service) EndVideoUploadSession(config *EndVideoUploadSessionConfig) (*EndVideoUploadSessionResponse, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("EndVideoUploadSessionConfig must not be a nil pointer")
	}

	var values = url.Values{}
	values.Set("upload_phase", "finish")
	values.Set("upload_session_id", config.UploadSessionId)
	values.Set("access_token", config.PageAccessToken)
	if config.Description != nil {
		values.Set("description", *config.Description)
	}
	if config.Title != nil {
		values.Set("title", *config.Title)
	}
	if config.Published != nil {
		values.Set("published", fmt.Sprintf("%v", *config.Published))
	}

	var response EndVideoUploadSessionResponse

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.urlV20(fmt.Sprintf("%s/videos?%s", config.PageId, values.Encode())),
		ResponseModel: &response,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &response, nil
}
