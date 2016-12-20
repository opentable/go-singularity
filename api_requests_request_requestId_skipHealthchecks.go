package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) SkipHealthchecks(requestId string, body *dtos.SingularitySkipHealthchecksRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "PUT", "/api/requests/request/{requestId}/skipHealthchecks", pathParamMap, queryParamMap, body)

	return
}

func (client *Client) DeleteExpiringSkipHealthchecks(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "DELETE", "/api/requests/request/{requestId}/skipHealthchecks", pathParamMap, queryParamMap)

	return
}
