package client

import "github.com/opentable/singularity/dtos"

func (client *Client) freezeSlave(slaveId string, body *dtos.SingularityMachineChangeRequest) (err error) {
	pathParamMap := map[string]interface{}{
		"slaveId": slaveId,
	}
	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/slaves/slave/{slaveId}/freeze", pathParamMap, queryParamMap, body)

	return
}
