package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/opentable/singularity/dtos"
)

type Client struct {
	baseUrl string
	http    http.Client
}

func New(apiBase string) (client *Client) {
	return &Client{apiBase, http.Client{}}
}

type urlParams map[string]interface{}

func (client *Client) DTORequest(pop dtos.DTO, method, path string, pathParams, queryParams urlParams, body ...dtos.DTO) (err error) {
	resBody, err := client.Request(method, path, pathParams, queryParams, body...)
	if err != nil {
		return
	}
	err = pop.Populate(resBody)
	return
}

func (client *Client) Request(method, path string, pathParams, queryParams urlParams, body ...dtos.DTO) (resBody io.ReadCloser, err error) {
	req, err := client.buildRequest(method, path, pathParams, queryParams, body...)
	if err != nil {
		return
	}
	res, err := client.http.Do(req)
	if err != nil {
		return
	}
	resBody = res.Body
	return
}

func (client *Client) buildRequest(method, path string, pathParams, queryParams urlParams, bodies ...dtos.DTO) (req *http.Request, err error) {
	url, err := url.Parse(client.baseUrl)
	if err != nil {
		return
	}

	path, err = pathRender(path, pathParams)
	if err != nil {
		return
	}
	url.Path = strings.Join([]string{url.Path, path}, "/")
	log.Print(url)

	if len(bodies) > 0 {
		req, err = buildBodyRequest(method, url.String(), bodies[0])
	} else {
		req, err = buildBodilessRequest(method, url.String())
	}
	return
}

func buildBodilessRequest(method, path string) (req *http.Request, err error) {
	req, err = http.NewRequest(method, path, nil)
	return
}

func buildBodyRequest(method, path string, bodyObj dtos.DTO) (req *http.Request, err error) {
	buf := bytes.Buffer{}
	enc := json.NewEncoder(&buf)
	enc.Encode(bodyObj) // XXX Here, consider a goroutine and a PipeWriter
	log.Print(path, buf.String())
	req, err = http.NewRequest(method, path, &buf)
	return
}

func pathRender(path string, params urlParams) (res string, err error) {
	parmRE := regexp.MustCompile(`{([^}]+)}`)
	building := make([]byte, 0, 150)
	pathBytes := []byte(path)

	start := 0
	indices := parmRE.FindAllSubmatchIndex(pathBytes, -1)
	for _, matches := range indices {
		building = append(building, pathBytes[start:matches[0]]...)
		start = matches[len(matches)-1]
		parmName := pathBytes[matches[2]:matches[3]]
		val, ok := params[string(parmName)]
		if !ok {
			err = fmt.Errorf("Path parameter %q not provided for %s", parmName, path)
			return
		}
		building = append(building, fmt.Sprintf("%v", val)...)
	}
	building = append(building, pathBytes[start+1:]...)
	res = string(building)

	return
}

func (client *Client) buildURL(subpath string) (url string) {
	return strings.Join([]string{client.baseUrl, subpath}, "/")
}
