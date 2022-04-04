package http

import (
	"cli/models"

	"github.com/go-resty/resty/v2"
)

// RestClient is a wrapper for resty client.
type RestClient[T models.GCPMonitoredResource] struct {
	c *resty.Client
}

func NewRestClient[T models.GCPMonitoredResource](client *resty.Client) *RestClient[T] {
	return &RestClient[T]{c: client}
}

func (c *RestClient[T]) SendMonitoringHttpRequest(method string, payload models.GCPMonitoredResource, query map[string]string) (T, error) {
	token, err := obtainToken()
	if err != nil {
		return *new(T), err
	}

	response, err := c.c.R().SetAuthToken(token).
		SetBody(payload).
		SetQueryParams(query).
		SetResult(payload.NewInstance()).
		Execute(method, payload.GetPath())

	return response.Result().(T), err
}

func obtainToken() (string, error) {
	return "", nil
}
