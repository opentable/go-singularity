package singularity

import "github.com/opentable/go-singularity/v0.13/dtos"

func (client *Client) GetS3LogsForDeploy(requestId string, deployId string, start int64, end int64) (response dtos.SingularityS3LogList, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "deployId": deployId,
	}

	queryParamMap := map[string]interface{}{
		"start": start, "end": end,
	}

	response = make(dtos.SingularityS3LogList, 0)
	err = client.DTORequest(&response, "GET", "/api/logs/request/{requestId}/deploy/{deployId}", pathParamMap, queryParamMap)

	return
}
