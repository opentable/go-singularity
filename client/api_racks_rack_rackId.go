package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getRackHistory(rackId string) (response *dtos.SingularityMachineStateHistoryUpdate, err error) {
	pathParamMap := map[string]interface{}{
		"rackId": rackId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityMachineStateHistoryUpdate)
	err = client.DTORequest(response, "GET", "/api/racks/rack/{rackId}", pathParamMap, queryParamMap)

	return
}

func (client *Client) removeRack(rackId string) (err error) {
	pathParamMap := map[string]interface{}{
		"rackId": rackId,
	}
	queryParamMap := map[string]interface{}{}

	_, err = client.Request("DELETE", "/api/racks/rack/{rackId}", pathParamMap, queryParamMap)

	return
}
