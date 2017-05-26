package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetPendingTask(pendingTaskId string) (response *dtos.SingularityTaskRequest, err error) {
	pathParamMap := map[string]interface{}{
		"pendingTaskId": pendingTaskId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskRequest)
	err = client.DTORequest(response, "GET", "/api/tasks/scheduled/task/{pendingTaskId}", pathParamMap, queryParamMap)

	return
}
