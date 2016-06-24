package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetRequest(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "GET", "/api/requests/request/{requestId}", pathParamMap, queryParamMap)

	return
}

func (client *Client) DeleteRequest(requestId string, body *dtos.SingularityDeleteRequestRequest) (response *dtos.SingularityRequest, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequest)
	err = client.DTORequest(response, "DELETE", "/api/requests/request/{requestId}", pathParamMap, queryParamMap, body)

	return
}
