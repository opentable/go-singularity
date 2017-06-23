package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetState(skipCache bool, includeRequestIds bool) (response *dtos.SingularityState, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"skipCache": skipCache, "includeRequestIds": includeRequestIds,
	}

	response = new(dtos.SingularityState)
	err = client.DTORequest(response, "GET", "/api/state", pathParamMap, queryParamMap)

	return
}
