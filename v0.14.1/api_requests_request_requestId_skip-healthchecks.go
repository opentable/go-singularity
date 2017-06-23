package singularity

import "github.com/opentable/go-singularity/v0.14.1/dtos"

func (client *Client) SkipHealthchecks(requestId string, body *dtos.SingularitySkipHealthchecksRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "PUT", "/api/requests/request/{requestId}/skip-healthchecks", pathParamMap, queryParamMap, body)

	return
}

func (client *Client) DeleteExpiringSkipHealthchecks(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "DELETE", "/api/requests/request/{requestId}/skip-healthchecks", pathParamMap, queryParamMap)

	return
}
