package singularity

import "github.com/opentable/go-singularity/dtos"

// DeleteActivePriorityFreeze is invalid
func (client *Client) CreatePriorityFreeze(body *dtos.SingularityPriorityFreeze) (response *dtos.SingularityPriorityFreezeParent, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityPriorityFreezeParent)
	err = client.DTORequest("singularity-createpriorityfreeze", response, "POST", "/api/priority/freeze", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) GetActivePriorityFreeze() (response *dtos.SingularityPriorityFreezeParent, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityPriorityFreezeParent)
	err = client.DTORequest("singularity-getactivepriorityfreeze", response, "GET", "/api/priority/freeze", pathParamMap, queryParamMap)

	return
}
