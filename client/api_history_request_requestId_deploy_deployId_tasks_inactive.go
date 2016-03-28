package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getInactiveDeployTasks(requestId string, deployId string, count int32, page int32) (response *dtos.SingularityTaskIdHistory, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "deployId": deployId,
	}
	queryParamMap := map[string]interface{}{
		"count": count, "page": page,
	}

	response = new(dtos.SingularityTaskIdHistory)
	err = client.DTORequest(response, "GET", "/api/history/request/{requestId}/deploy/{deployId}/tasks/inactive", pathParamMap, queryParamMap)

	return
}
