package singularity

import "github.com/opentable/go-singularity/dtos"

// DeleteExpiringRackStateChange is invalid
func (client *Client) GetExpiringRackStateChanges() (response dtos.SingularityExpiringMachineStateList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityExpiringMachineStateList, 0)
	err = client.DTORequest(&response, "GET", "/api/racks/expiring", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetRacks(state string) (response dtos.SingularityRackList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"state": state,
	}

	response = make(dtos.SingularityRackList, 0)
	err = client.DTORequest(&response, "GET", "/api/racks/", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetRackHistory(rackId string) (response dtos.SingularityMachineStateHistoryUpdateList, err error) {
	pathParamMap := map[string]interface{}{
		"rackId": rackId,
	}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityMachineStateHistoryUpdateList, 0)
	err = client.DTORequest(&response, "GET", "/api/racks/rack/{rackId}", pathParamMap, queryParamMap)

	return
}

// RemoveRack is invalid
// DecommissionRack is invalid
// FreezeRack is invalid
// ActivateRack is invalid
