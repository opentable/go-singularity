package singularity

import "github.com/opentable/go-singularity/v0.14.1/dtos"

func (client *Client) GetDeploysWithMetadata(requestId string, count int32, page int32) (response dtos.SingularityDeployHistoryList, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{
		"count": count, "page": page,
	}

	response = make(dtos.SingularityDeployHistoryList, 0)
	err = client.DTORequest(&response, "GET", "/api/history/request/{requestId}/deploys/withmetadata", pathParamMap, queryParamMap)

	return
}
