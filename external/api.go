package external

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"otpelf-local/request"
	"otpelf-local/response"
	"strings"
	"sync"
)

var s IApi
var once sync.Once

type service struct {
	client *http.Client
}

type Response struct {
	Code int
	Body []byte
}

type IApi interface {
	ExecuteRequest(ctx *context.Context, url string, method string, request interface{}) (*Response, error)
	FetchOtp(ctx *context.Context, request *request.OtpRequest) ([]response.OtpInfo, error)
}

func NewRequest(ctx *context.Context, method string, url string, payload io.Reader, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	for header, value := range headers {
		req.Header.Set(header, value)
	}
	return req, nil
}

func NewApiService() IApi {

	var client *http.Client

	client = &http.Client{Transport: http.DefaultTransport}

	once.Do(func() {
		s = &service{
			client: client,
		}
	})

	return s
}

func (s *service) ExecuteRequest(ctx *context.Context, url string, method string, request interface{}) (*Response, error) {
	reqJSON, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	completeUrl := "https://api-dark.razorpay.com/v1" + url

	httpReq, err := NewRequest(ctx,
		method,
		completeUrl,
		strings.NewReader(string(reqJSON)),
		nil)

	httpReq.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	key, secret := "", ""

	httpReq.SetBasicAuth(key, secret)

	response, err := s.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &Response{Code: response.StatusCode, Body: body}, nil
}

func (s *service) FetchOtp(ctx *context.Context, request *request.OtpRequest) ([]response.OtpInfo, error) {
	info := []response.OtpInfo{}

	resp, err := s.ExecuteRequest(ctx, "/terminal_test_otp", "GET", request)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(resp.Body, &info)

	return info, nil
}
