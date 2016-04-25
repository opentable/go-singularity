package singularity

import "github.com/opentable/singularity/dtos"

func (client *Client) GetOverProvisionedRequestIds(skipCache bool) (response dtos.StringList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"skipCache": skipCache,
	}

	response = make(dtos.StringList, 0)
	err = client.DTORequest(&response, "GET", "/api/state/requests/over-provisioned", pathParamMap, queryParamMap)

	return
}
