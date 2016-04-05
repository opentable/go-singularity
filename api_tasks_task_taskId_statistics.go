package client

import "github.com/opentable/singularity/dtos"

func (client *Client) GetTaskStatistics(taskId string) (response *dtos.MesosTaskStatisticsObject, err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.MesosTaskStatisticsObject)
	err = client.DTORequest(response, "GET", "/api/tasks/task/{taskId}/statistics", pathParamMap, queryParamMap)

	return
}
