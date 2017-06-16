package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetCleaningTasks() (response dtos.SingularityTaskCleanupList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityTaskCleanupList, 0)
	err = client.DTORequest(&response, "GET", "/api/tasks/cleaning", pathParamMap, queryParamMap)

	return
}
