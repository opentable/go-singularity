package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getActiveTasks() (response *dtos.SingularityTask, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTask)
	err = client.DTORequest(response, "GET", "/api/tasks/active", pathParamMap, queryParamMap)

	return
}
