package singularity

import "github.com/opentable/go-singularity/v0.14.1/dtos"

func (client *Client) DeleteExpiringScale(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "DELETE", "/api/requests/request/{requestId}/scale", pathParamMap, queryParamMap)

	return
}

func (client *Client) Scale(requestId string, body *dtos.SingularityScaleRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "PUT", "/api/requests/request/{requestId}/scale", pathParamMap, queryParamMap, body)

	return
}
