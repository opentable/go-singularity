package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetExpiringSlaveStateChanges() (response dtos.SingularityExpiringMachineStateList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityExpiringMachineStateList, 0)
	err = client.DTORequest(&response, "GET", "/api/slaves/expiring", pathParamMap, queryParamMap)

	return
}
