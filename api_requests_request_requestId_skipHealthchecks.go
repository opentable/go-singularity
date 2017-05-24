package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) DeleteExpiringSkipHealthchecksDeprecated(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "DELETE", "/api/requests/request/{requestId}/skipHealthchecks", pathParamMap, queryParamMap)

	return
}

func (client *Client) SkipHealthchecksDeprecated(requestId string, body *dtos.SingularitySkipHealthchecksRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "PUT", "/api/requests/request/{requestId}/skipHealthchecks", pathParamMap, queryParamMap, body)

	return
}
