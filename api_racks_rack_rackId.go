package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetRackHistory(rackId string) (response dtos.SingularityMachineStateHistoryUpdateList, err error) {
	pathParamMap := map[string]interface{}{
		"rackId": rackId,
	}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityMachineStateHistoryUpdateList, 0)
	err = client.DTORequest(&response, "GET", "/api/racks/rack/{rackId}", pathParamMap, queryParamMap)

	return
}

func (client *Client) RemoveRack(rackId string) (err error) {
	pathParamMap := map[string]interface{}{
		"rackId": rackId,
	}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("DELETE", "/api/racks/rack/{rackId}", pathParamMap, queryParamMap)

	return
}
