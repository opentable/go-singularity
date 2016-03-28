package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getPendingTask(pendingTaskId string) (response *dtos.SingularityTaskRequest, err error) {
	pathParamMap := map[string]interface{}{
		"pendingTaskId": pendingTaskId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskRequest)
	err = client.DTORequest(response, "GET", "/api/tasks/scheduled/task/{pendingTaskId}", pathParamMap, queryParamMap)

	return
}
