package singularity

import "github.com/opentable/go-singularity/dtos"

// ReactivateSlave is invalid
func (client *Client) GetSlaves(state string) (response dtos.SingularitySlaveList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"state": state,
	}

	response = make(dtos.SingularitySlaveList, 0)
	err = client.DTORequest("singularity-getslaves", &response, "GET", "/api/slaves/", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetSlaveHistory(slaveId string) (response dtos.SingularityMachineStateHistoryUpdateList, err error) {
	pathParamMap := map[string]interface{}{
		"slaveId": slaveId,
	}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityMachineStateHistoryUpdateList, 0)
	err = client.DTORequest("singularity-getslavehistory", &response, "GET", "/api/slaves/slave/{slaveId}", pathParamMap, queryParamMap)

	return
}

// RemoveSlave is invalid
func (client *Client) GetSlave(slaveId string) (response *dtos.SingularitySlave, err error) {
	pathParamMap := map[string]interface{}{
		"slaveId": slaveId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularitySlave)
	err = client.DTORequest("singularity-getslave", response, "GET", "/api/slaves/slave/{slaveId}/details", pathParamMap, queryParamMap)

	return
}

// DecommissionSlave is invalid
// FreezeSlave is invalid
// DeleteExpiringSlaveStateChange is invalid
func (client *Client) GetExpiringSlaveStateChanges() (response dtos.SingularityExpiringMachineStateList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityExpiringMachineStateList, 0)
	err = client.DTORequest("singularity-getexpiringslavestatechanges", &response, "GET", "/api/slaves/expiring", pathParamMap, queryParamMap)

	return
}
