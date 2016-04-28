package singularity

import "github.com/opentable/singularity/dtos"

func (client *Client) GetRequestHistoryForRequestLike(requestIdLike string, count int32, page int32) (response dtos.StringList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"requestIdLike": requestIdLike, "count": count, "page": page,
	}

	response = make(dtos.StringList, 0)
	err = client.DTORequest(&response, "GET", "/api/history/requests/search", pathParamMap, queryParamMap)

	return
}
