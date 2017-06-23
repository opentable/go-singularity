package singularity

import "github.com/opentable/go-singularity/v0.14.1/dtos"

func (client *Client) GetQueuedDeployUpdates(webhookId string) (response dtos.SingularityDeployUpdateList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"webhookId": webhookId,
	}

	response = make(dtos.SingularityDeployUpdateList, 0)
	err = client.DTORequest(&response, "GET", "/api/webhooks/deploy", pathParamMap, queryParamMap)

	return
}
