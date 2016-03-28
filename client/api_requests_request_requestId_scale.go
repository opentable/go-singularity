package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) scale(requestId string, body *dtos.SingularityScaleRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "PUT", "/api/requests/request/{requestId}/scale", pathParamMap, queryParamMap, body)

	return
}

func (client *Client) deleteExpiringScale(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "DELETE", "/api/requests/request/{requestId}/scale", pathParamMap, queryParamMap)

	return
}
