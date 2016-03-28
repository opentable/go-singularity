package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getActiveDeployTasks(requestId string, deployId string) (response *dtos.SingularityTaskIdHistory, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "deployId": deployId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskIdHistory)
	err = client.DTORequest(response, "GET", "/api/history/request/{requestId}/deploy/{deployId}/tasks/active", pathParamMap, queryParamMap)

	return
}
