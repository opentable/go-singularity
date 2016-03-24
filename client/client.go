package client

import (
	"fmt"
	"net/http"
	"regexp"
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

func (client *Client) Request(method, path string, pathParams, queryParams map[string]string, body ...interface{}) {
	req := buildReqest(method, path, pathParams, queryParams, body...)
	res, err := client.http.Do(req)
}

type urlParams map[string]string

func buildRequest(method, path string, pathParams, queryParams urlParams, body ...interface{}) {
	path = pathRender(path, pathParams)
	http.NewRequest(method, path, body)
}

func pathRender(path string, params urlParams) (res string, err error) {
	parmRE := regexp.MustCompile(`{([^}]+)}`)
	building := make([]byte, 150, 0)
	pathBytes := []byte(path)

	start := 0
	indices := parmRE.FindAllSubmatchIndexes(pathBytes)
	for _, matches := range indices {
		building = append(building, pathBytes[start:matches[0]-1])
		start = matches[len(matches)-1]
		parmName := pathBytes[matches[2]:matches[3]]
		val, ok := params[parmName]
		if !ok {
			err = fmt.Errorf("Path parameter %q not provided for %s", parmName, path)
			return
		}
		building = append(building, val)
	}
	building = append(building, pathBytes[start:])
	res = string(building)

	return
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
