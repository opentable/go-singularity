package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) decommissionRack(rackId string, body *dtos.SingularityMachineChangeRequest) (err error) {
	pathParamMap := map[string]interface{}{
		"rackId": rackId,
	}
	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/racks/rack/{rackId}/decommission", pathParamMap, queryParamMap, body)

	return
}
