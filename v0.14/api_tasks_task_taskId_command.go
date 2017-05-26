package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) RunShellCommand(taskId string, body *dtos.SingularityShellCommand) (response *dtos.SingularityTaskShellCommandRequest, err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskShellCommandRequest)
	err = client.DTORequest(response, "POST", "/api/tasks/task/{taskId}/command", pathParamMap, queryParamMap, body)

	return
}
