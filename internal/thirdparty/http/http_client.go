package httpclient

import "net/http"

type HttpClient struct {
	Client
}

func (httpClient *HttpClient) Do(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
