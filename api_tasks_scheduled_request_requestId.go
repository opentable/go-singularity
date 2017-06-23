package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetScheduledTasksForRequest(requestId string, useWebCache bool) (response dtos.SingularityTaskRequestList, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(dtos.SingularityTaskRequestList, 0)
	err = client.DTORequest(&response, "GET", "/api/tasks/scheduled/request/{requestId}", pathParamMap, queryParamMap)

	return
}
