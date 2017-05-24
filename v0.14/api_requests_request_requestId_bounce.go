package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) Bounce(requestId string, body *dtos.SingularityBounceRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "POST", "/api/requests/request/{requestId}/bounce", pathParamMap, queryParamMap, body)

	return
}

func (client *Client) DeleteExpiringBounce(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "DELETE", "/api/requests/request/{requestId}/bounce", pathParamMap, queryParamMap)

	return
}
