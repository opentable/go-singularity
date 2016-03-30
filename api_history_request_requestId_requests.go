package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getRequestHistoryForRequest(requestId string, count int32, page int32) (response *dtos.SingularityRequestHistory, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}
	queryParamMap := map[string]interface{}{
		"count": count, "page": page,
	}

	response = new(dtos.SingularityRequestHistory)
	err = client.DTORequest(response, "GET", "/api/history/request/{requestId}/requests", pathParamMap, queryParamMap)

	return
}
