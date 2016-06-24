package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetKilledTasks() (response dtos.SingularityKilledTaskIdRecordList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityKilledTaskIdRecordList, 0)
	err = client.DTORequest(&response, "GET", "/api/tasks/killed", pathParamMap, queryParamMap)

	return
}
