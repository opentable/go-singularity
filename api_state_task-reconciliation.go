package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetTaskReconciliationStatistics() (response *dtos.SingularityTaskReconciliationStatistics, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskReconciliationStatistics)
	err = client.DTORequest(response, "GET", "/api/state/task-reconciliation", pathParamMap, queryParamMap)

	return
}
