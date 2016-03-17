package client

import (
	"net/http"
	"strings"

	"github.com/opentable/sous-singularity/client/dtos"
)

type Client struct {
	baseUrl string
	http    http.Client
}

func New(apiBase string) (client *Client) {
	return &Client{apiBase + "/api", http.Client{}}
}

func (client *Client) APIGet(path string, pop dtos.DTO) (err error) {
	url := client.buildURL(path)
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

func (client *Client) GetState() (ss *dtos.State, err error) {
	ss = &dtos.State{}
	err = ss.Get(client)
	return
}

func (client *Client) GetDeploys() (deps *dtos.Deploys, err error) {
	deps = &dtos.Deploys{}
	err = deps.Get(client)
	return
}

func (client *Client) GetRequests() (reqs *dtos.Requests, err error) {
	reqs = &dtos.Requests{}
	err = reqs.Get(client)
	return
}
