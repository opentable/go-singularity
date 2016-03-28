package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) bounce(requestId string, body *dtos.SingularityBounceRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "POST", "/api/requests/request/{requestId}/bounce", pathParamMap, queryParamMap, body)

	return
}

func (client *Client) deleteExpiringBounce(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "DELETE", "/api/requests/request/{requestId}/bounce", pathParamMap, queryParamMap)

	return
}
