package singularity

import (
	"github.com/opentable/go-singularity/dtos"
	"github.com/opentable/swaggering"
)

func (client *Client) GetTaskHistory(requestId string, deployId string, runId string, host string, lastTaskStatus string, startedBefore int64, startedAfter int64, updatedBefore int64, updatedAfter int64, orderDirection string, count int32, page int32) (response dtos.SingularityTaskIdHistoryList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"requestId": requestId, "deployId": deployId, "runId": runId, "host": host, "lastTaskStatus": lastTaskStatus, "startedBefore": startedBefore, "startedAfter": startedAfter, "updatedBefore": updatedBefore, "updatedAfter": updatedAfter, "orderDirection": orderDirection, "count": count, "page": page,
	}

	response = make(dtos.SingularityTaskIdHistoryList, 0)
	err = client.DTORequest(&response, "GET", "/api/history/tasks", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetDeploy(requestId string, deployId string) (response *dtos.SingularityDeployHistory, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "deployId": deployId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityDeployHistory)
	err = client.DTORequest(response, "GET", "/api/history/request/{requestId}/deploy/{deployId}", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetHistoryForTask(taskId string) (response *dtos.SingularityTaskHistory, err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskHistory)
	err = client.DTORequest(response, "GET", "/api/history/task/{taskId}", pathParamMap, queryParamMap)

	return
}
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
func (client *Client) GetTaskHistoryForActiveRequest(requestId string) (response dtos.SingularityTaskIdHistoryList, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityTaskIdHistoryList, 0)
	err = client.DTORequest(&response, "GET", "/api/history/request/{requestId}/tasks/active", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetActiveDeployTasks(requestId string, deployId string) (response dtos.SingularityTaskIdHistoryList, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "deployId": deployId,
	}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityTaskIdHistoryList, 0)
	err = client.DTORequest(&response, "GET", "/api/history/request/{requestId}/deploy/{deployId}/tasks/active", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetInactiveDeployTasks(requestId string, deployId string, count int32, page int32) (response dtos.SingularityTaskIdHistoryList, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "deployId": deployId,
	}

	queryParamMap := map[string]interface{}{
		"count": count, "page": page,
	}

	response = make(dtos.SingularityTaskIdHistoryList, 0)
	err = client.DTORequest(&response, "GET", "/api/history/request/{requestId}/deploy/{deployId}/tasks/inactive", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetInactiveDeployTasksWithMetadata(requestId string, deployId string, count int32, page int32) (response dtos.SingularityTaskIdHistoryList, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "deployId": deployId,
	}

	queryParamMap := map[string]interface{}{
		"count": count, "page": page,
	}

	response = make(dtos.SingularityTaskIdHistoryList, 0)
	err = client.DTORequest(&response, "GET", "/api/history/request/{requestId}/deploy/{deployId}/tasks/inactive/withmetadata", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetTaskHistoryWithMetadata(requestId string, deployId string, runId string, host string, lastTaskStatus string, startedBefore int64, startedAfter int64, updatedBefore int64, updatedAfter int64, orderDirection string, count int32, page int32) (response dtos.SingularityTaskIdHistoryList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"requestId": requestId, "deployId": deployId, "runId": runId, "host": host, "lastTaskStatus": lastTaskStatus, "startedBefore": startedBefore, "startedAfter": startedAfter, "updatedBefore": updatedBefore, "updatedAfter": updatedAfter, "orderDirection": orderDirection, "count": count, "page": page,
	}

	response = make(dtos.SingularityTaskIdHistoryList, 0)
	err = client.DTORequest(&response, "GET", "/api/history/tasks/withmetadata", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetTaskHistoryForRequestWithMetadata(requestId string, deployId string, runId string, host string, lastTaskStatus string, startedBefore int64, startedAfter int64, updatedBefore int64, updatedAfter int64, orderDirection string, count int32, page int32) (response dtos.SingularityTaskIdHistoryList, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{
		"deployId": deployId, "runId": runId, "host": host, "lastTaskStatus": lastTaskStatus, "startedBefore": startedBefore, "startedAfter": startedAfter, "updatedBefore": updatedBefore, "updatedAfter": updatedAfter, "orderDirection": orderDirection, "count": count, "page": page,
	}

	response = make(dtos.SingularityTaskIdHistoryList, 0)
	err = client.DTORequest(&response, "GET", "/api/history/request/{requestId}/tasks/withmetadata", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetTaskHistoryForRequestAndRunId(requestId string, runId string) (response *dtos.SingularityTaskIdHistory, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "runId": runId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskIdHistory)
	err = client.DTORequest(response, "GET", "/api/history/request/{requestId}/run/{runId}", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetDeploys(requestId string, count int32, page int32) (response dtos.SingularityDeployHistoryList, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{
		"count": count, "page": page,
	}

	response = make(dtos.SingularityDeployHistoryList, 0)
	err = client.DTORequest(&response, "GET", "/api/history/request/{requestId}/deploys", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetDeploysWithMetadata(requestId string, count int32, page int32) (response dtos.SingularityDeployHistoryList, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{
		"count": count, "page": page,
	}

	response = make(dtos.SingularityDeployHistoryList, 0)
	err = client.DTORequest(&response, "GET", "/api/history/request/{requestId}/deploys/withmetadata", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetRequestHistoryForRequest(requestId string, count int32, page int32) (response dtos.SingularityRequestHistoryList, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{
		"count": count, "page": page,
	}

	response = make(dtos.SingularityRequestHistoryList, 0)
	err = client.DTORequest(&response, "GET", "/api/history/request/{requestId}/requests", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetRequestHistoryForRequestWithMetadata(requestId string, count int32, page int32) (response dtos.SingularityRequestHistoryList, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{
		"count": count, "page": page,
	}

	response = make(dtos.SingularityRequestHistoryList, 0)
	err = client.DTORequest(&response, "GET", "/api/history/request/{requestId}/requests/withmetadata", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetRequestHistoryForRequestLike(requestIdLike string, count int32, page int32, useWebCache bool) (response swaggering.StringList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"requestIdLike": requestIdLike, "count": count, "page": page, "useWebCache": useWebCache,
	}

	response = make(swaggering.StringList, 0)
	err = client.DTORequest(&response, "GET", "/api/history/requests/search", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetRecentCommandLineArgs(requestId string, count int32) (response *dtos.Set, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{
		"count": count,
	}

	response = new(dtos.Set)
	err = client.DTORequest(response, "GET", "/api/history/request/{requestId}/command-line-args", pathParamMap, queryParamMap)

	return
}
