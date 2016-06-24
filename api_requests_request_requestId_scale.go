package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) Scale(requestId string, body *dtos.SingularityScaleRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "PUT", "/api/requests/request/{requestId}/scale", pathParamMap, queryParamMap, body)

	return
}

func (client *Client) DeleteExpiringScale(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "DELETE", "/api/requests/request/{requestId}/scale", pathParamMap, queryParamMap)

	return
}
