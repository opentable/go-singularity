package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getPendingDeploys() (response *dtos.SingularityPendingDeploy, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityPendingDeploy)
	err = client.DTORequest(response, "GET", "/api/deploys/pending", pathParamMap, queryParamMap)

	return
}
