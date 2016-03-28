package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getTasksForSlave(slaveId string) (response *dtos.SingularityTask, err error) {
	pathParamMap := map[string]interface{}{
		"slaveId": slaveId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTask)
	err = client.DTORequest(response, "GET", "/api/tasks/active/slave/{slaveId}", pathParamMap, queryParamMap)

	return
}
