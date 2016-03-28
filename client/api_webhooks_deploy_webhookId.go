package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getQueuedDeployUpdates(webhookId string) (response *dtos.SingularityDeployUpdate, err error) {
	pathParamMap := map[string]interface{}{
		"webhookId": webhookId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityDeployUpdate)
	err = client.DTORequest(response, "GET", "/api/webhooks/deploy/{webhookId}", pathParamMap, queryParamMap)

	return
}
