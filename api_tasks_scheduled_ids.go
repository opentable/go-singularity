package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetScheduledTaskIds() (response dtos.SingularityPendingTaskIdList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityPendingTaskIdList, 0)
	err = client.DTORequest(&response, "GET", "/api/tasks/scheduled/ids", pathParamMap, queryParamMap)

	return
}
