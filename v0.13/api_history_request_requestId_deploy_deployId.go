package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetDeploy(requestId string, deployId string) (response *dtos.SingularityDeployHistory, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "deployId": deployId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityDeployHistory)
	err = client.DTORequest(response, "GET", "/api/history/request/{requestId}/deploy/{deployId}", pathParamMap, queryParamMap)

	return
}
