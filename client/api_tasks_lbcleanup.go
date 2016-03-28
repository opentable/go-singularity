package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getLbCleanupTasks() (response *dtos.SingularityTaskId, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskId)
	err = client.DTORequest(response, "GET", "/api/tasks/lbcleanup", pathParamMap, queryParamMap)

	return
}
