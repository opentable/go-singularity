package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getQueuedRequestUpdates(webhookId string) (response *dtos.SingularityRequestHistory, err error) {
	pathParamMap := map[string]interface{}{
		"webhookId": webhookId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestHistory)
	err = client.DTORequest(response, "GET", "/api/webhooks/request/{webhookId}", pathParamMap, queryParamMap)

	return
}
