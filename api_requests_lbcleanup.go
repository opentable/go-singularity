package singularity

import "github.com/opentable/swaggering"

func (client *Client) GetLbCleanupRequests(useWebCache bool) (response swaggering.StringList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(swaggering.StringList, 0)
	err = client.DTORequest(&response, "GET", "/api/requests/lbcleanup", pathParamMap, queryParamMap)

	return
}
