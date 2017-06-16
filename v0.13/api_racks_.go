package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetRacks(state string) (response dtos.SingularityRackList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"state": state,
	}

	response = make(dtos.SingularityRackList, 0)
	err = client.DTORequest(&response, "GET", "/api/racks/", pathParamMap, queryParamMap)

	return
}
