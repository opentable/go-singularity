package singularity

import "github.com/opentable/go-singularity/v0.14/dtos"

func (client *Client) GetSlave(slaveId string) (response *dtos.SingularitySlave, err error) {
	pathParamMap := map[string]interface{}{
		"slaveId": slaveId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularitySlave)
	err = client.DTORequest(response, "GET", "/api/slaves/slave/{slaveId}/details", pathParamMap, queryParamMap)

	return
}
