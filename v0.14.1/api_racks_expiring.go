package singularity

import "github.com/opentable/go-singularity/v0.14.1/dtos"

func (client *Client) GetExpiringRackStateChanges() (response dtos.SingularityExpiringMachineStateList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityExpiringMachineStateList, 0)
	err = client.DTORequest(&response, "GET", "/api/racks/expiring", pathParamMap, queryParamMap)

	return
}
