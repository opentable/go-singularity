package singularity

import "github.com/opentable/go-singularity/v0.14/dtos"

func (client *Client) ActivateRack(rackId string, body *dtos.SingularityMachineChangeRequest) (err error) {
	pathParamMap := map[string]interface{}{
		"rackId": rackId,
	}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/racks/rack/{rackId}/activate", pathParamMap, queryParamMap, body)

	return
}
