package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetQueuedTaskUpdates(webhookId string) (response dtos.SingularityTaskHistoryUpdateList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"webhookId": webhookId,
	}

	response = make(dtos.SingularityTaskHistoryUpdateList, 0)
	err = client.DTORequest(&response, "GET", "/api/webhooks/task", pathParamMap, queryParamMap)

	return
}
