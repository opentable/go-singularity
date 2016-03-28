package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getDeploy(requestId string, deployId string) (response *dtos.SingularityDeployHistory, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "deployId": deployId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityDeployHistory)
	err = client.DTORequest(response, "GET", "/api/history/request/{requestId}/deploy/{deployId}", pathParamMap, queryParamMap)

	return
}
