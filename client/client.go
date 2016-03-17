package client

import (
	"io"
	"net/http"
	"strings"
)

type Client struct {
	baseUrl string
	http    http.Client
}

func New() (client *Client) {
	return &Client{"https://singularity.somewhere/api", http.Client{}}
}

func (client *Client) APIGet(path string) (resBody io.ReadCloser, err error) {
	res, err := client.http.Get(client.buildURL(path))
	if err != nil {
		return
	}

	resBody = res.Body
	return
}

func (client *Client) buildURL(subpath string) (url string) {
	return strings.Join([]string{client.baseUrl, subpath}, "/")
}
