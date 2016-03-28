package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) decommissionSlave(slaveId string, body *dtos.SingularityMachineChangeRequest) (err error) {
	pathParamMap := map[string]interface{}{
		"slaveId": slaveId,
	}
	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/slaves/slave/{slaveId}/decommission", pathParamMap, queryParamMap, body)

	return
}
