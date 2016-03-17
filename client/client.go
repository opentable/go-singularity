package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	baseUrl string
	http    http.Client
}

type Populater interface {
	Populate(io.ReadCloser) error
}

func New(apiBase string) (client *Client) {
	return &Client{apiBase + "/api", http.Client{}}
}

func (client *Client) APIGet(path string, pop Populater) (err error) {
	url := client.buildURL(path)
	fmt.Printf("Getting from: %s\n\n", url)
	res, err := client.http.Get(url)
	if err != nil {
		return
	}
	resBody := res.Body
	err = pop.Populate(resBody)
	return
}

func (client *Client) buildURL(subpath string) (url string) {
	return strings.Join([]string{client.baseUrl, subpath}, "/")
}
