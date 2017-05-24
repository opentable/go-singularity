package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetTaskByRunId(requestId string, runId string) (response *dtos.SingularityTaskId, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "runId": runId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskId)
	err = client.DTORequest(response, "GET", "/api/requests/request/{requestId}/run/{runId}", pathParamMap, queryParamMap)

	return
}
