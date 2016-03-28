package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getQueuedTaskUpdates(webhookId string) (response *dtos.SingularityTaskHistoryUpdate, err error) {
	pathParamMap := map[string]interface{}{
		"webhookId": webhookId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskHistoryUpdate)
	err = client.DTORequest(response, "GET", "/api/webhooks/task/{webhookId}", pathParamMap, queryParamMap)

	return
}
