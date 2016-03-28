package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) runShellCommand(taskId string, body *dtos.SingularityShellCommand) (response *dtos.SingularityTaskShellCommandRequest, err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskShellCommandRequest)
	err = client.DTORequest(response, "POST", "/api/tasks/task/{taskId}/command", pathParamMap, queryParamMap, body)

	return
}
