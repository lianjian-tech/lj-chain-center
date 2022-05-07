package util

import (
	"errors"
	"github.com/valyala/fasthttp"
	"time"
)

const (
	POST     = "POST"
	GET      = "GET"
	TIME_OUT = 10 * 1000
)

func HttpClient(uri string, method string, data []byte) ([]byte, error) {
	return HttpClientSub(uri, method, data, false, "")
}

func HttpClientSub(uri string, method string, data []byte, needAuth bool, authValue string) ([]byte, error) {
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	client := fasthttp.Client{}

	req.Header.SetMethod(method)
	req.Header.Set("Content-Type", "application/json")
	if needAuth {
		req.Header.Set("Authorization", authValue)
	}

	if data != nil {
		req.SetBody(data)
	}
	req.SetRequestURI(uri)
	if err := client.DoTimeout(req, resp, TIME_OUT*time.Millisecond); err != nil {
		return nil, errors.New("HttpClient DoTimeout")
	}
	return resp.Body(), nil
}
