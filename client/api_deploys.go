package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) deploy(body *dtos.SingularityDeployRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "POST", "/api/deploys", pathParamMap, queryParamMap, body)

	return
}
