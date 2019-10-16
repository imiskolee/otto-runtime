package http

import (
	"net/http"
	otto_runtime "github.com/imiskolee/otto-runtime"
)

func init() {
	otto_runtime.Register("NewHttpClient",NewHttpClient)
}

type HttpClient struct {
	httpClient *http.Client
	err        error
}

func NewHttpClient() *HttpClient {
	return &HttpClient{httpClient: http.DefaultClient}
}
func (h *HttpClient) Header(k, v string) *HttpClient {
	//todo
	return h
}

func (h *HttpClient) Get(url string) string {
	return ""
}

func (h *HttpClient) Post(url string, body string, contentType string) string {
	return ""
}

func (h *HttpClient) Put(url string, body string, contentType string) string {
	return ""
}
