package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) getPendingRequests() (response *dtos.SingularityPendingRequest, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityPendingRequest)
	err = client.DTORequest(response, "GET", "/api/requests/queued/pending", pathParamMap, queryParamMap)

	return
}
