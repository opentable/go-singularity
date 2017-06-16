package singularity

import "github.com/opentable/go-singularity/v0.13/dtos"

func (client *Client) GetS3LogsForTask(taskId string, start int64, end int64) (response dtos.SingularityS3LogList, err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}

	queryParamMap := map[string]interface{}{
		"start": start, "end": end,
	}

	response = make(dtos.SingularityS3LogList, 0)
	err = client.DTORequest(&response, "GET", "/api/logs/task/{taskId}", pathParamMap, queryParamMap)

	return
}
