package client

import "bytes"

func (client *Client) getOverProvisionedRequestIds(skipCache bool) (response string, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{
		"skipCache": skipCache,
	}

	resBody, err := client.Request("GET", "/api/state/requests/over-provisioned", pathParamMap, queryParamMap)
	readBuf := bytes.Buffer{}
	readBuf.ReadFrom(resBody)
	response = string(readBuf.Bytes())
	return
}
