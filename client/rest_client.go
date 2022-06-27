package client

import (
	"errors"
	"net/url"
	"strconv"
	"time"

	"github.com/franela/goreq"
	"github.com/jayhyuk/opsgenie-go-sdk/logging"
)

// OpsGenieRestClient is the data type to make requests.
type RestClient struct {
	OpsGenieClient
}

// SetOpsGenieClient sets the embedded OpsGenieClient type of the OpsGenieAlertClient.
func (cli *RestClient) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

func (cli *RestClient) writeBody(resp *goreq.Response, body interface{}) error {
	if err := resp.Body.FromJsonTo(&body); err != nil {
		message := "Server response can not be parsed, " + err.Error()
		logging.Logger().Warn(message)
		return errors.New(message)
	}
	return nil
}

func (cli *RestClient) sendGetRequest(req Request, response Response) error {
	path, params, err := req.GenerateUrl()
	if err != nil {
		return err
	}

	request := cli.buildGetRequest(cli.generateFullPathWithParams(path, params), nil)
	cli.setApiKey(&request, req.GetApiKey())
	httpResponse, err := cli.sendRequestWithRetries(request)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()

	err = cli.writeBody(httpResponse, &response)
	if err != nil {
		return err
	}
	cli.setResponseMeta(httpResponse, response)
	return nil
}

func (cli *RestClient) sendPatchRequest(req Request, response Response) error {
	path, params, err := req.GenerateUrl()
	if err != nil {
		return err
	}

	request := cli.buildPatchRequest(cli.generateFullPathWithParams(path, params), req)
	cli.setApiKey(&request, req.GetApiKey())
	httpResponse, err := cli.sendRequestWithRetries(request)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()

	err = cli.writeBody(httpResponse, &response)
	if err != nil {
		return err
	}
	cli.setResponseMeta(httpResponse, response)
	return nil
}

func (cli *RestClient) sendPutRequest(req Request, response Response) error {
	path, params, err := req.GenerateUrl()
	if err != nil {
		return err
	}

	request := cli.buildPutRequest(cli.generateFullPathWithParams(path, params), req)
	cli.setApiKey(&request, req.GetApiKey())
	httpResponse, err := cli.sendRequestWithRetries(request)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()

	err = cli.writeBody(httpResponse, &response)
	if err != nil {
		return err
	}
	cli.setResponseMeta(httpResponse, response)
	return nil
}

func (cli *RestClient) sendAsyncPostRequest(req Request) (*AsyncRequestResponse, error) {
	var response AsyncRequestResponse
	err := cli.sendPostRequest(req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *RestClient) sendPostRequest(req Request, response Response) error {
	path, params, err := req.GenerateUrl()

	if err != nil {
		return err
	}

	path = cli.generateFullPathWithParams(path, params)

	httpRequest := cli.buildPostRequest(path, req)
	cli.setApiKey(&httpRequest, req.GetApiKey())
	httpResponse, err := cli.sendRequestWithRetries(httpRequest)

	if err != nil {
		return err
	}

	defer httpResponse.Body.Close()

	err = cli.writeBody(httpResponse, &response)
	if err != nil {
		return err
	}

	cli.setResponseMeta(httpResponse, response)

	return nil
}

func (cli *RestClient) sendDeleteRequest(req Request, response Response) error {
	path, params, err := req.GenerateUrl()
	if err != nil {
		return err
	}

	path = cli.generateFullPathWithParams(path, params)

	httpRequest := cli.buildDeleteRequest(path, nil)
	cli.setApiKey(&httpRequest, req.GetApiKey())

	httpResponse, err := cli.sendRequestWithRetries(httpRequest)
	if err != nil {
		return err
	}

	defer httpResponse.Body.Close()

	err = cli.writeBody(httpResponse, &response)
	if err != nil {
		return err
	}

	cli.setResponseMeta(httpResponse, response)

	return nil
}

func (cli *RestClient) generateFullPathWithParams(url string, values url.Values) string {
	if values != nil {
		return url + "?" + values.Encode()
	} else {
		return url
	}
}

func (cli *RestClient) setApiKey(req *goreq.Request, fromRequest string) {
	var apiKey string

	if fromRequest == "" {
		apiKey = cli.apiKey
	} else {
		apiKey = fromRequest
	}

	req.AddHeader("Authorization", "GenieKey "+apiKey)
}

func (cli *RestClient) setResponseMeta(httpResponse *goreq.Response, response Response) {
	requestID := httpResponse.Header.Get("X-Request-ID")
	response.SetRequestID(requestID)

	rateLimitState := httpResponse.Header.Get("X-RateLimit-State")
	response.SetRateLimitState(rateLimitState)

	responseTime, err := strconv.ParseFloat(httpResponse.Header.Get("X-Response-Time"), 32)
	if err == nil {
		response.SetResponseTime(float32(responseTime))
	}
}

func (cli *RestClient) sendRequestWithRetries(request goreq.Request) (*goreq.Response, error) {
	var retry = 20
	maxSleepMs := 30000
	var sleepMs = 100
	var httpResponse *goreq.Response
	var err error
	for retry > 0 {
		httpResponse, err := cli.sendRequest(request)
		if err == nil {
			return httpResponse, err
		}
		defer httpResponse.Body.Close()
		retry--
		sleepMs *= 2
		if sleepMs > maxSleepMs {
			sleepMs = maxSleepMs
		}
		time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	}

	return httpResponse, err
}

type Request interface {
	GetApiKey() string
	GenerateUrl() (string, url.Values, error)
}

type Response interface {
	SetRequestID(requestId string)
	SetResponseTime(responseTime float32)
	SetRateLimitState(state string)
}

type ResponseMeta struct {
	RequestID      string
	ResponseTime   float32
	RateLimitState string
}

func (rm *ResponseMeta) SetRequestID(requestID string) {
	rm.RequestID = requestID
}

func (rm *ResponseMeta) SetResponseTime(responseTime float32) {
	rm.ResponseTime = responseTime
}

func (rm *ResponseMeta) SetRateLimitState(state string) {
	rm.RateLimitState = state
}

// Response for async processing requests
type AsyncRequestResponse struct {
	ResponseMeta
	RequestID string `json:"requestId"`
}
