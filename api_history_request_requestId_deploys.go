package singularity

import "github.com/opentable/singularity/dtos"

func (client *Client) getDeploys(requestId string, count int32, page int32) (response *dtos.SingularityDeployHistory, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}
	queryParamMap := map[string]interface{}{
		"count": count, "page": page,
	}

	response = new(dtos.SingularityDeployHistory)
	err = client.DTORequest(response, "GET", "/api/history/request/{requestId}/deploys", pathParamMap, queryParamMap)

	return
}
