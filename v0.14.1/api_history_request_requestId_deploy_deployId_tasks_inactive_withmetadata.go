package singularity

import "github.com/opentable/go-singularity/v0.14.1/dtos"

func (client *Client) GetInactiveDeployTasksWithMetadata(requestId string, deployId string, count int32, page int32) (response dtos.SingularityTaskIdHistoryList, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "deployId": deployId,
	}

	queryParamMap := map[string]interface{}{
		"count": count, "page": page,
	}

	response = make(dtos.SingularityTaskIdHistoryList, 0)
	err = client.DTORequest(&response, "GET", "/api/history/request/{requestId}/deploy/{deployId}/tasks/inactive/withmetadata", pathParamMap, queryParamMap)

	return
}
