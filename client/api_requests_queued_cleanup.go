package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getCleanupRequests() (response *dtos.SingularityRequestCleanup, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestCleanup)
	err = client.DTORequest(response, "GET", "/api/requests/queued/cleanup", pathParamMap, queryParamMap)

	return
}
