package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetTaskHistory(requestId string, deployId string, host string, lastTaskStatus string, startedAfter int64, startedBefore int64, orderDirection string, count int32, page int32) (response dtos.SingularityTaskIdHistoryList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"requestId": requestId, "deployId": deployId, "host": host, "lastTaskStatus": lastTaskStatus, "startedAfter": startedAfter, "startedBefore": startedBefore, "orderDirection": orderDirection, "count": count, "page": page,
	}

	response = make(dtos.SingularityTaskIdHistoryList, 0)
	err = client.DTORequest(&response, "GET", "/api/history/tasks", pathParamMap, queryParamMap)

	return
}
