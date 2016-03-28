package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getScheduledTasksForRequest(requestId string) (response *dtos.SingularityTaskRequest, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskRequest)
	err = client.DTORequest(response, "GET", "/api/tasks/scheduled/request/{requestId}", pathParamMap, queryParamMap)

	return
}
