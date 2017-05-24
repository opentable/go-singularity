package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetScheduledTaskIds(useWebCache bool) (response dtos.SingularityPendingTaskIdList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(dtos.SingularityPendingTaskIdList, 0)
	err = client.DTORequest(&response, "GET", "/api/tasks/scheduled/ids", pathParamMap, queryParamMap)

	return
}
