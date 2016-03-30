package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getSlaveHistory(slaveId string) (response *dtos.SingularityMachineStateHistoryUpdate, err error) {
	pathParamMap := map[string]interface{}{
		"slaveId": slaveId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityMachineStateHistoryUpdate)
	err = client.DTORequest(response, "GET", "/api/slaves/slave/{slaveId}", pathParamMap, queryParamMap)

	return
}

func (client *Client) removeSlave(slaveId string) (err error) {
	pathParamMap := map[string]interface{}{
		"slaveId": slaveId,
	}
	queryParamMap := map[string]interface{}{}

	_, err = client.Request("DELETE", "/api/slaves/slave/{slaveId}", pathParamMap, queryParamMap)

	return
}
