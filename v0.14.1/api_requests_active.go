package singularity

import "github.com/opentable/go-singularity/v0.14.1/dtos"

func (client *Client) GetActiveRequests() (response dtos.SingularityRequestParentList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityRequestParentList, 0)
	err = client.DTORequest(&response, "GET", "/api/requests/active", pathParamMap, queryParamMap)

	return
}
