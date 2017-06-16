package singularity

import "github.com/opentable/go-singularity/v0.14/dtos"

func (client *Client) GetRequestGroupIds() (response dtos.SingularityRequestGroupList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityRequestGroupList, 0)
	err = client.DTORequest(&response, "GET", "/api/groups", pathParamMap, queryParamMap)

	return
}

func (client *Client) SaveRequestGroup(body *dtos.SingularityRequestGroup) (response *dtos.SingularityRequestGroup, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestGroup)
	err = client.DTORequest(response, "POST", "/api/groups", pathParamMap, queryParamMap, body)

	return
}
