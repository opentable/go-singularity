package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) DecommissionRack(rackId string, body *dtos.SingularityMachineChangeRequest) (err error) {
	pathParamMap := map[string]interface{}{
		"rackId": rackId,
	}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/racks/rack/{rackId}/decommission", pathParamMap, queryParamMap, body)

	return
}
