package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetRequestGroup(requestGroupId string) (response *dtos.SingularityRequestGroup, err error) {
	pathParamMap := map[string]interface{}{
		"requestGroupId": requestGroupId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestGroup)
	err = client.DTORequest(response, "GET", "/api/groups/group/{requestGroupId}", pathParamMap, queryParamMap)

	return
}

// DeleteRequestGroup is invalid
func (client *Client) GetRequestGroupIds(useWebCache bool) (response dtos.SingularityRequestGroupList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

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
