package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getScheduledTasks() (response *dtos.SingularityTaskRequest, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskRequest)
	err = client.DTORequest(response, "GET", "/api/tasks/scheduled", pathParamMap, queryParamMap)

	return
}
