package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getCleaningTasks() (response *dtos.SingularityTaskCleanup, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskCleanup)
	err = client.DTORequest(response, "GET", "/api/tasks/cleaning", pathParamMap, queryParamMap)

	return
}
