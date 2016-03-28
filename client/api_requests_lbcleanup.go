package client

import "bytes"

func (client *Client) getLbCleanupRequests() (response string, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	resBody, err := client.Request("GET", "/api/requests/lbcleanup", pathParamMap, queryParamMap)
	readBuf := bytes.Buffer{}
	readBuf.ReadFrom(resBody)
	response = string(readBuf.Bytes())
	return
}
