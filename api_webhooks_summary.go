package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetWebhooksWithQueueSize() (response dtos.SingularityWebhookSummaryList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityWebhookSummaryList, 0)
	err = client.DTORequest(&response, "GET", "/api/webhooks/summary", pathParamMap, queryParamMap)

	return
}
