package singularity

import "github.com/opentable/go-singularity/v0.14.1/dtos"

func (client *Client) GetQueuedDeployUpdatesDeprecated(webhookId string) (response dtos.SingularityDeployUpdateList, err error) {
	pathParamMap := map[string]interface{}{
		"webhookId": webhookId,
	}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityDeployUpdateList, 0)
	err = client.DTORequest(&response, "GET", "/api/webhooks/deploy/{webhookId}", pathParamMap, queryParamMap)

	return
}
