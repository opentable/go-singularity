package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) GetQueuedDeployUpdates(webhookId string) (response dtos.SingularityDeployUpdateList, err error) {
	pathParamMap := map[string]interface{}{
		"webhookId": webhookId,
	}
	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityDeployUpdateList, 0)
	err = client.DTORequest(response, "GET", "/api/webhooks/deploy/{webhookId}", pathParamMap, queryParamMap)

	return
}
