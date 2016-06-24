package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetLbCleanupRequests() (response dtos.StringList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.StringList, 0)
	err = client.DTORequest(&response, "GET", "/api/requests/lbcleanup", pathParamMap, queryParamMap)

	return
}
