package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetHistoryForTask(taskId string) (response *dtos.SingularityTaskHistory, err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskHistory)
	err = client.DTORequest(response, "GET", "/api/history/task/{taskId}", pathParamMap, queryParamMap)

	return
}
