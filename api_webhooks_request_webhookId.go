package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetQueuedRequestUpdates(webhookId string) (response dtos.SingularityRequestHistoryList, err error) {
	pathParamMap := map[string]interface{}{
		"webhookId": webhookId,
	}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityRequestHistoryList, 0)
	err = client.DTORequest(&response, "GET", "/api/webhooks/request/{webhookId}", pathParamMap, queryParamMap)

	return
}
