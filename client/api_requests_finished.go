package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getFinishedRequests() (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "GET", "/api/requests/finished", pathParamMap, queryParamMap)

	return
}
