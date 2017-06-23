package singularity

import "github.com/opentable/go-singularity/v0.14.1/dtos"

func (client *Client) GetLbCleanupTasks() (response dtos.SingularityTaskIdList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityTaskIdList, 0)
	err = client.DTORequest(&response, "GET", "/api/tasks/lbcleanup", pathParamMap, queryParamMap)

	return
}
