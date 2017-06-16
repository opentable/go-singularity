package singularity

import (
	"bytes"

	"github.com/opentable/go-singularity/dtos"
)

func (client *Client) GetActiveWebhooks() (response dtos.SingularityWebhookList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityWebhookList, 0)
	err = client.DTORequest(&response, "GET", "/api/webhooks", pathParamMap, queryParamMap)

	return
}

func (client *Client) AddWebhook(body *dtos.SingularityWebhook) (response string, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	resBody, err := client.Request("POST", "/api/webhooks", pathParamMap, queryParamMap, body)
	readBuf := bytes.Buffer{}
	readBuf.ReadFrom(resBody)
	response = string(readBuf.Bytes())
	return
}
