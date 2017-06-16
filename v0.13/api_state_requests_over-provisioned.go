package singularity

import "github.com/opentable/swaggering"

func (client *Client) GetOverProvisionedRequestIds(skipCache bool) (response swaggering.StringList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"skipCache": skipCache,
	}

	response = make(swaggering.StringList, 0)
	err = client.DTORequest(&response, "GET", "/api/state/requests/over-provisioned", pathParamMap, queryParamMap)

	return
}
