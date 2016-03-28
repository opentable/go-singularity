package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getState(skipCache bool, includeRequestIds bool) (response *dtos.SingularityState, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{
		"skipCache": skipCache, "includeRequestIds": includeRequestIds,
	}

	response = new(dtos.SingularityState)
	err = client.DTORequest(response, "GET", "/api/state", pathParamMap, queryParamMap)

	return
}
