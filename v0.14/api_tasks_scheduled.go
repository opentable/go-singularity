package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetScheduledTasks() (response dtos.SingularityTaskRequestList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityTaskRequestList, 0)
	err = client.DTORequest(&response, "GET", "/api/tasks/scheduled", pathParamMap, queryParamMap)

	return
}
