package singularity

import "github.com/opentable/go-singularity/v0.13/dtos"

func (client *Client) FreezeSlave(slaveId string, body *dtos.SingularityMachineChangeRequest) (err error) {
	pathParamMap := map[string]interface{}{
		"slaveId": slaveId,
	}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/slaves/slave/{slaveId}/freeze", pathParamMap, queryParamMap, body)

	return
}
