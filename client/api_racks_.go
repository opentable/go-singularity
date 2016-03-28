package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getRacks(state string) (response *dtos.SingularityRack, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{
		"state": state,
	}

	response = new(dtos.SingularityRack)
	err = client.DTORequest(response, "GET", "/api/racks/", pathParamMap, queryParamMap)

	return
}
