package singularity

import (
	"bytes"

	"github.com/opentable/go-singularity/dtos"
)

func (client *Client) GetActiveWebhooks() (response dtos.SingularityWebhookList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityWebhookList, 0)
	err = client.DTORequest("singularity-getactivewebhooks", &response, "GET", "/api/webhooks", pathParamMap, queryParamMap)

	return
}
func (client *Client) AddWebhook(body *dtos.SingularityWebhook) (response string, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	resBody, err := client.Request("singularity-addwebhook", "POST", "/api/webhooks", pathParamMap, queryParamMap, body)
	readBuf := bytes.Buffer{}
	readBuf.ReadFrom(resBody)
	response = string(readBuf.Bytes())
	return
}
func (client *Client) DeleteWebhook(webhookId string) (response string, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"webhookId": webhookId,
	}

	resBody, err := client.Request("singularity-deletewebhook", "DELETE", "/api/webhooks", pathParamMap, queryParamMap)
	readBuf := bytes.Buffer{}
	readBuf.ReadFrom(resBody)
	response = string(readBuf.Bytes())
	return
}
func (client *Client) GetWebhooksWithQueueSize() (response dtos.SingularityWebhookSummaryList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityWebhookSummaryList, 0)
	err = client.DTORequest("singularity-getwebhookswithqueuesize", &response, "GET", "/api/webhooks/summary", pathParamMap, queryParamMap)

	return
}
func (client *Client) DeleteWebhookDeprecated(webhookId string) (response string, err error) {
	pathParamMap := map[string]interface{}{
		"webhookId": webhookId,
	}

	queryParamMap := map[string]interface{}{}

	resBody, err := client.Request("singularity-deletewebhookdeprecated", "DELETE", "/api/webhooks/{webhookId}", pathParamMap, queryParamMap)
	readBuf := bytes.Buffer{}
	readBuf.ReadFrom(resBody)
	response = string(readBuf.Bytes())
	return
}
func (client *Client) GetQueuedDeployUpdatesDeprecated(webhookId string) (response dtos.SingularityDeployUpdateList, err error) {
	pathParamMap := map[string]interface{}{
		"webhookId": webhookId,
	}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityDeployUpdateList, 0)
	err = client.DTORequest("singularity-getqueueddeployupdatesdeprecated", &response, "GET", "/api/webhooks/deploy/{webhookId}", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetQueuedRequestUpdatesDeprecated(webhookId string) (response dtos.SingularityRequestHistoryList, err error) {
	pathParamMap := map[string]interface{}{
		"webhookId": webhookId,
	}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityRequestHistoryList, 0)
	err = client.DTORequest("singularity-getqueuedrequestupdatesdeprecated", &response, "GET", "/api/webhooks/request/{webhookId}", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetQueuedTaskUpdatesDeprecated(webhookId string) (response dtos.SingularityTaskHistoryUpdateList, err error) {
	pathParamMap := map[string]interface{}{
		"webhookId": webhookId,
	}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityTaskHistoryUpdateList, 0)
	err = client.DTORequest("singularity-getqueuedtaskupdatesdeprecated", &response, "GET", "/api/webhooks/task/{webhookId}", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetQueuedDeployUpdates(webhookId string) (response dtos.SingularityDeployUpdateList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"webhookId": webhookId,
	}

	response = make(dtos.SingularityDeployUpdateList, 0)
	err = client.DTORequest("singularity-getqueueddeployupdates", &response, "GET", "/api/webhooks/deploy", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetQueuedRequestUpdates(webhookId string) (response dtos.SingularityRequestHistoryList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"webhookId": webhookId,
	}

	response = make(dtos.SingularityRequestHistoryList, 0)
	err = client.DTORequest("singularity-getqueuedrequestupdates", &response, "GET", "/api/webhooks/request", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetQueuedTaskUpdates(webhookId string) (response dtos.SingularityTaskHistoryUpdateList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"webhookId": webhookId,
	}

	response = make(dtos.SingularityTaskHistoryUpdateList, 0)
	err = client.DTORequest("singularity-getqueuedtaskupdates", &response, "GET", "/api/webhooks/task", pathParamMap, queryParamMap)

	return
}
