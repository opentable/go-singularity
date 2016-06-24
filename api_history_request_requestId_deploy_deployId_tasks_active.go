package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetActiveDeployTasks(requestId string, deployId string) (response dtos.SingularityTaskIdHistoryList, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "deployId": deployId,
	}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityTaskIdHistoryList, 0)
	err = client.DTORequest(&response, "GET", "/api/history/request/{requestId}/deploy/{deployId}/tasks/active", pathParamMap, queryParamMap)

	return
}
