package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getKilledTasks() (response *dtos.SingularityKilledTaskIdRecord, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityKilledTaskIdRecord)
	err = client.DTORequest(response, "GET", "/api/tasks/killed", pathParamMap, queryParamMap)

	return
}
