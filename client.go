package singularity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/opentable/singularity/dtos"
)

//go:generate swagger-client-maker api-docs/ .

type Client struct {
	BaseUrl string
	Debug   bool
	http    http.Client
}

type ReqError struct {
	method, path string
	message      string
	body         bytes.Buffer
}

func (e *ReqError) Error() string {
	return fmt.Sprintf("%s %s => %s: \n%s", e.method, e.path, e.message, e.body.String())
}

func NewClient(apiBase string) (client *Client) {
	return &Client{apiBase, false, http.Client{}}
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
	if res.StatusCode > 299 {
		rerr := &ReqError{
			message: res.Status,
			method:  method,
			path:    path,
			body:    bytes.Buffer{},
		}
		rerr.body.ReadFrom(res.Body)
		res.Body.Close()
		err = rerr
		return
	}
	if client.Debug {
		buf := bytes.Buffer{}
		buf.ReadFrom(res.Body)
		log.Printf("%s -> %+v\n", req.URL, buf.String())
		resBody = ioutil.NopCloser(&buf)
	} else {
		resBody = res.Body
	}
	return
}

func (client *Client) buildRequest(method, path string, pathParams, queryParams urlParams, bodies ...dtos.DTO) (req *http.Request, err error) {
	url, err := url.Parse(client.BaseUrl)
	if err != nil {
		return
	}

	path, err = pathRender(path, pathParams)
	if err != nil {
		return
	}
	url.Path = strings.Join([]string{strings.TrimRight(url.Path, "/"), strings.TrimLeft(path, "/")}, "/")

	if client.Debug {
		log.Print(url)
	}

	if len(bodies) > 0 {
		req, err = client.buildBodyRequest(method, url.String(), bodies[0])
	} else {
		req, err = client.buildBodilessRequest(method, url.String())
	}
	return
}

func (cl *Client) buildBodilessRequest(method, path string) (req *http.Request, err error) {
	if cl.Debug {
		log.Print(method, " ", path)
	}
	req, err = http.NewRequest(method, path, nil)
	return
}

func (cl *Client) buildBodyRequest(method, path string, bodyObj dtos.DTO) (req *http.Request, err error) {
	buf := bytes.Buffer{}
	enc := json.NewEncoder(&buf)
	enc.Encode(bodyObj) // XXX Here, consider a goroutine and a PipeWriter
	if cl.Debug {
		log.Print(method, " ", path, " <- ", buf.String())
	}
	req, err = http.NewRequest(method, path, &buf)
	req.Header.Add("Content-Type", "application/json")
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
		start = matches[len(matches)-1] + 1
		parmName := pathBytes[matches[2]:matches[3]]
		val, ok := params[string(parmName)]
		if !ok {
			err = fmt.Errorf("Path parameter %q not provided for %s", parmName, path)
			return
		}
		building = append(building, fmt.Sprintf("%v", val)...)
	}
	building = append(building, pathBytes[start:]...)
	res = string(building)

	return
}

func (client *Client) buildURL(subpath string) (url string) {
	return strings.Join([]string{client.BaseUrl, subpath}, "/")
}
