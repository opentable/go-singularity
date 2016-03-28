package client

import (
	"bytes"

	"github.com/opentable/sous-singularity/client/dtos"
)

func (client *Client) getActiveWebhooks() (response *dtos.SingularityWebhook, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityWebhook)
	err = client.DTORequest(response, "GET", "/api/webhooks", pathParamMap, queryParamMap)

	return
}

func (client *Client) addWebhook(body *dtos.SingularityWebhook) (response string, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	resBody, err := client.Request("POST", "/api/webhooks", pathParamMap, queryParamMap, body)
	readBuf := bytes.Buffer{}
	readBuf.ReadFrom(resBody)
	response = string(readBuf.Bytes())
	return
}
