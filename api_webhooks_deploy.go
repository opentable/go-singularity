package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetQueuedDeployUpdates(webhookId string) (response dtos.SingularityDeployUpdateList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"webhookId": webhookId,
	}

	response = make(dtos.SingularityDeployUpdateList, 0)
	err = client.DTORequest(&response, "GET", "/api/webhooks/deploy", pathParamMap, queryParamMap)

	return
}
