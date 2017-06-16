package singularity

import "github.com/opentable/swaggering"

func (client *Client) GetUnderProvisionedRequestIds(skipCache bool) (response swaggering.StringList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"skipCache": skipCache,
	}

	response = make(swaggering.StringList, 0)
	err = client.DTORequest(&response, "GET", "/api/state/requests/under-provisioned", pathParamMap, queryParamMap)

	return
}
