package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetRequests() (response dtos.SingularityRequestParentList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityRequestParentList, 0)
	err = client.DTORequest(&response, "GET", "/api/requests", pathParamMap, queryParamMap)

	return
}

func (client *Client) PostRequest(body *dtos.SingularityRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "POST", "/api/requests", pathParamMap, queryParamMap, body)

	return
}
