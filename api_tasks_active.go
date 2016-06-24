package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetActiveTasks() (response dtos.SingularityTaskList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityTaskList, 0)
	err = client.DTORequest(&response, "GET", "/api/tasks/active", pathParamMap, queryParamMap)

	return
}
