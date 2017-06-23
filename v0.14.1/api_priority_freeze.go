package singularity

import "github.com/opentable/go-singularity/v0.14.1/dtos"

func (client *Client) GetActivePriorityFreeze() (response *dtos.SingularityPriorityFreezeParent, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityPriorityFreezeParent)
	err = client.DTORequest(response, "GET", "/api/priority/freeze", pathParamMap, queryParamMap)

	return
}

func (client *Client) DeleteActivePriorityFreeze() (err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("DELETE", "/api/priority/freeze", pathParamMap, queryParamMap)

	return
}

func (client *Client) CreatePriorityFreeze(body *dtos.SingularityPriorityFreeze) (response *dtos.SingularityPriorityFreezeParent, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityPriorityFreezeParent)
	err = client.DTORequest(response, "POST", "/api/priority/freeze", pathParamMap, queryParamMap, body)

	return
}
