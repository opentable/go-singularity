package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getSlaves(state string) (response *dtos.SingularitySlave, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{
		"state": state,
	}

	response = new(dtos.SingularitySlave)
	err = client.DTORequest(response, "GET", "/api/slaves/", pathParamMap, queryParamMap)

	return
}
