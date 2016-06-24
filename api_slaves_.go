package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetSlaves(state string) (response dtos.SingularitySlaveList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"state": state,
	}

	response = make(dtos.SingularitySlaveList, 0)
	err = client.DTORequest(&response, "GET", "/api/slaves/", pathParamMap, queryParamMap)

	return
}
