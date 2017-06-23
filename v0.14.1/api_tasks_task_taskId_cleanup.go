package singularity

import "github.com/opentable/go-singularity/v0.14.1/dtos"

func (client *Client) GetTaskCleanup(taskId string) (response *dtos.SingularityTaskCleanup, err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskCleanup)
	err = client.DTORequest(response, "GET", "/api/tasks/task/{taskId}/cleanup", pathParamMap, queryParamMap)

	return
}
