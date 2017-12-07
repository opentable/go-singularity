package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetS3LogsForTask(taskId string, start int64, end int64, excludeMetadata bool, list bool) (response dtos.SingularityS3LogList, err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}

	queryParamMap := map[string]interface{}{
		"start": start, "end": end, "excludeMetadata": excludeMetadata, "list": list,
	}

	response = make(dtos.SingularityS3LogList, 0)
	err = client.DTORequest("singularity-gets3logsfortask", &response, "GET", "/api/logs/task/{taskId}", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetS3LogsForRequest(requestId string, start int64, end int64, excludeMetadata bool, list bool, maxPerPage int32) (response dtos.SingularityS3LogList, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{
		"start": start, "end": end, "excludeMetadata": excludeMetadata, "list": list, "maxPerPage": maxPerPage,
	}

	response = make(dtos.SingularityS3LogList, 0)
	err = client.DTORequest("singularity-gets3logsforrequest", &response, "GET", "/api/logs/request/{requestId}", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetS3LogsForDeploy(requestId string, deployId string, start int64, end int64, excludeMetadata bool, list bool, maxPerPage int32) (response dtos.SingularityS3LogList, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "deployId": deployId,
	}

	queryParamMap := map[string]interface{}{
		"start": start, "end": end, "excludeMetadata": excludeMetadata, "list": list, "maxPerPage": maxPerPage,
	}

	response = make(dtos.SingularityS3LogList, 0)
	err = client.DTORequest("singularity-gets3logsfordeploy", &response, "GET", "/api/logs/request/{requestId}/deploy/{deployId}", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetPaginatedS3Logs(body *dtos.SingularityS3SearchRequest) (response dtos.SingularityS3LogList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityS3LogList, 0)
	err = client.DTORequest("singularity-getpaginateds3logs", &response, "POST", "/api/logs/search", pathParamMap, queryParamMap, body)

	return
}
