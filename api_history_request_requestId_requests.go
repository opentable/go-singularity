package singularity

import "github.com/opentable/go-singularity/dtos"

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
