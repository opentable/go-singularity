package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) scheduleImmediately(requestId string, body *dtos.SingularityRunNowRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "POST", "/api/requests/request/{requestId}/run", pathParamMap, queryParamMap, body)

	return
}
