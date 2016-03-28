package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getRequests() (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "GET", "/api/requests", pathParamMap, queryParamMap)

	return
}

func (client *Client) postRequest(body *dtos.SingularityRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "POST", "/api/requests", pathParamMap, queryParamMap, body)

	return
}
