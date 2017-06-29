package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) Deploy(body *dtos.SingularityDeployRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "POST", "/api/deploys", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) GetPendingDeploys() (response dtos.SingularityPendingDeployList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityPendingDeployList, 0)
	err = client.DTORequest(&response, "GET", "/api/deploys/pending", pathParamMap, queryParamMap)

	return
}
func (client *Client) CancelDeploy(requestId string, deployId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "deployId": deployId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "DELETE", "/api/deploys/deploy/{deployId}/request/{requestId}", pathParamMap, queryParamMap)

	return
}
func (client *Client) UpdatePendingDeploy(body *dtos.SingularityUpdatePendingDeployRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "POST", "/api/deploys/update", pathParamMap, queryParamMap, body)

	return
}
