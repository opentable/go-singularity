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

func (client *Client) DeleteRequestGroup(requestGroupId string) (err error) {
	pathParamMap := map[string]interface{}{
		"requestGroupId": requestGroupId,
	}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("DELETE", "/api/groups/group/{requestGroupId}", pathParamMap, queryParamMap)

	return
}
