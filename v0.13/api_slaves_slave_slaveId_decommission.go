package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) DecommissionSlave(slaveId string, body *dtos.SingularityMachineChangeRequest) (err error) {
	pathParamMap := map[string]interface{}{
		"slaveId": slaveId,
	}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/slaves/slave/{slaveId}/decommission", pathParamMap, queryParamMap, body)

	return
}
