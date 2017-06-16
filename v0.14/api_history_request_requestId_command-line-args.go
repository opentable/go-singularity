package singularity

import "github.com/opentable/go-singularity/v0.14/dtos"

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
