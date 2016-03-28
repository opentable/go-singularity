package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getScheduledTaskIds() (response *dtos.SingularityPendingTaskId, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityPendingTaskId)
	err = client.DTORequest(response, "GET", "/api/tasks/scheduled/ids", pathParamMap, queryParamMap)

	return
}
