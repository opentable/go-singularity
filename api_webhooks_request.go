package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetQueuedRequestUpdates(webhookId string) (response dtos.SingularityRequestHistoryList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"webhookId": webhookId,
	}

	response = make(dtos.SingularityRequestHistoryList, 0)
	err = client.DTORequest(&response, "GET", "/api/webhooks/request", pathParamMap, queryParamMap)

	return
}
