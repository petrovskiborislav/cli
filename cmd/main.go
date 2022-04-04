package main

import (
	"cli/http"
	"cli/models"
	"fmt"

	"github.com/go-resty/resty/v2"
)

const (
	// BaseURL of the API
	BaseURL = "https://6249e966852fe6ebf8820f75.mockapi.io"
)

func main() {
	httpClient := resty.New().SetBaseURL(BaseURL)
	serviceRestClient := http.NewRestClient[*models.Services](httpClient)
	f, err := serviceRestClient.SendMonitoringHttpRequest(resty.MethodGet, &models.Services{}, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)

}
