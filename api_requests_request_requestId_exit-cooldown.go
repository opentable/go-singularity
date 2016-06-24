package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) ExitCooldown(requestId string, body *dtos.SingularityExitCooldownRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "POST", "/api/requests/request/{requestId}/exit-cooldown", pathParamMap, queryParamMap, body)

	return
}
