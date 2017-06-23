package singularity

import "github.com/opentable/go-singularity/v0.14.1/dtos"

func (client *Client) UpdatePendingDeploy(body *dtos.SingularityUpdatePendingDeployRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "POST", "/api/deploys/update", pathParamMap, queryParamMap, body)

	return
}
