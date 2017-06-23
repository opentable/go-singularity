package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetTaskHistoryForRequest(requestId string, deployId string, runId string, host string, lastTaskStatus string, startedBefore int64, startedAfter int64, updatedBefore int64, updatedAfter int64, orderDirection string, count int32, page int32) (response dtos.SingularityTaskIdHistoryList, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{
		"deployId": deployId, "runId": runId, "host": host, "lastTaskStatus": lastTaskStatus, "startedBefore": startedBefore, "startedAfter": startedAfter, "updatedBefore": updatedBefore, "updatedAfter": updatedAfter, "orderDirection": orderDirection, "count": count, "page": page,
	}

	response = make(dtos.SingularityTaskIdHistoryList, 0)
	err = client.DTORequest(&response, "GET", "/api/history/request/{requestId}/tasks", pathParamMap, queryParamMap)

	return
}
