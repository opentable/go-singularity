package client

import "github.com/opentable/singularity/dtos"

func (client *Client) getHistoryForTask(taskId string) (response *dtos.SingularityTaskHistory, err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskHistory)
	err = client.DTORequest(response, "GET", "/api/history/task/{taskId}", pathParamMap, queryParamMap)

	return
}
