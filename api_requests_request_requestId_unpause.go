package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) Unpause(requestId string, body *dtos.SingularityUnpauseRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "POST", "/api/requests/request/{requestId}/unpause", pathParamMap, queryParamMap, body)

	return
}
