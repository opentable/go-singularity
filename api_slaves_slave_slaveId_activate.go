package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) ReactivateSlave(slaveId string, body *dtos.SingularityMachineChangeRequest) (err error) {
	pathParamMap := map[string]interface{}{
		"slaveId": slaveId,
	}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/slaves/slave/{slaveId}/activate", pathParamMap, queryParamMap, body)

	return
}
