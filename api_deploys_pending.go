package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetPendingDeploys() (response dtos.SingularityPendingDeployList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityPendingDeployList, 0)
	err = client.DTORequest(&response, "GET", "/api/deploys/pending", pathParamMap, queryParamMap)

	return
}
