package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetTasksForSlave(slaveId string) (response dtos.SingularityTaskList, err error) {
	pathParamMap := map[string]interface{}{
		"slaveId": slaveId,
	}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityTaskList, 0)
	err = client.DTORequest(&response, "GET", "/api/tasks/active/slave/{slaveId}", pathParamMap, queryParamMap)

	return
}
