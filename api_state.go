package singularity

import (
	"github.com/opentable/go-singularity/dtos"
	"github.com/opentable/swaggering"
)

func (client *Client) GetUnderProvisionedRequestIds(skipCache bool) (response swaggering.StringList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"skipCache": skipCache,
	}

	response = make(swaggering.StringList, 0)
	err = client.DTORequest(&response, "GET", "/api/state/requests/under-provisioned", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetOverProvisionedRequestIds(skipCache bool) (response swaggering.StringList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"skipCache": skipCache,
	}

	response = make(swaggering.StringList, 0)
	err = client.DTORequest(&response, "GET", "/api/state/requests/over-provisioned", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetTaskReconciliationStatistics() (response *dtos.SingularityTaskReconciliationStatistics, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskReconciliationStatistics)
	err = client.DTORequest(response, "GET", "/api/state/task-reconciliation", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetState(skipCache bool, includeRequestIds bool) (response *dtos.SingularityState, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"skipCache": skipCache, "includeRequestIds": includeRequestIds,
	}

	response = new(dtos.SingularityState)
	err = client.DTORequest(response, "GET", "/api/state", pathParamMap, queryParamMap)

	return
}
