package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) CancelDeploy(requestId string, deployId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "deployId": deployId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "DELETE", "/api/deploys/deploy/{deployId}/request/{requestId}", pathParamMap, queryParamMap)

	return
}
