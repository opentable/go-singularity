package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) FreezeRack(rackId string, body *dtos.SingularityMachineChangeRequest) (err error) {
	pathParamMap := map[string]interface{}{
		"rackId": rackId,
	}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/racks/rack/{rackId}/freeze", pathParamMap, queryParamMap, body)

	return
}
