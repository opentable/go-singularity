package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) skipHealthchecks(requestId string, body *dtos.SingularitySkipHealthchecksRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "PUT", "/api/requests/request/{requestId}/skipHealthchecks", pathParamMap, queryParamMap, body)

	return
}

func (client *Client) deleteExpiringSkipHealthchecks(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "DELETE", "/api/requests/request/{requestId}/skipHealthchecks", pathParamMap, queryParamMap)

	return
}
