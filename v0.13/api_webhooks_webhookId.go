package singularity

import "bytes"

func (client *Client) DeleteWebhook(webhookId string) (response string, err error) {
	pathParamMap := map[string]interface{}{
		"webhookId": webhookId,
	}

	queryParamMap := map[string]interface{}{}

	resBody, err := client.Request("DELETE", "/api/webhooks/{webhookId}", pathParamMap, queryParamMap)
	readBuf := bytes.Buffer{}
	readBuf.ReadFrom(resBody)
	response = string(readBuf.Bytes())
	return
}
