package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetSlaveHistory(slaveId string) (response dtos.SingularityMachineStateHistoryUpdateList, err error) {
	pathParamMap := map[string]interface{}{
		"slaveId": slaveId,
	}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityMachineStateHistoryUpdateList, 0)
	err = client.DTORequest(&response, "GET", "/api/slaves/slave/{slaveId}", pathParamMap, queryParamMap)

	return
}

func (client *Client) RemoveSlave(slaveId string) (err error) {
	pathParamMap := map[string]interface{}{
		"slaveId": slaveId,
	}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("DELETE", "/api/slaves/slave/{slaveId}", pathParamMap, queryParamMap)

	return
}
