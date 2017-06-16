package singularity

import "github.com/opentable/go-singularity/v0.14/dtos"

func (client *Client) GetS3LogsForRequest(requestId string, start int64, end int64, excludeMetadata bool, list bool, maxPerPage int32) (response dtos.SingularityS3LogList, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{
		"start": start, "end": end, "excludeMetadata": excludeMetadata, "list": list, "maxPerPage": maxPerPage,
	}

	response = make(dtos.SingularityS3LogList, 0)
	err = client.DTORequest(&response, "GET", "/api/logs/request/{requestId}", pathParamMap, queryParamMap)

	return
}
