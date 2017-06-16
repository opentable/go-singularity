package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetTaskHistoryForActiveRequest(requestId string) (response dtos.SingularityTaskIdHistoryList, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityTaskIdHistoryList, 0)
	err = client.DTORequest(&response, "GET", "/api/history/request/{requestId}/tasks/active", pathParamMap, queryParamMap)

	return
}
