package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetQueuedTaskUpdates(webhookId string) (response dtos.SingularityTaskHistoryUpdateList, err error) {
	pathParamMap := map[string]interface{}{
		"webhookId": webhookId,
	}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityTaskHistoryUpdateList, 0)
	err = client.DTORequest(&response, "GET", "/api/webhooks/task/{webhookId}", pathParamMap, queryParamMap)

	return
}
