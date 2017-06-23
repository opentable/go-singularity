package singularity

import "github.com/opentable/go-singularity/v0.14.1/dtos"

func (client *Client) GetPendingRequests() (response dtos.SingularityPendingRequestList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityPendingRequestList, 0)
	err = client.DTORequest(&response, "GET", "/api/requests/queued/pending", pathParamMap, queryParamMap)

	return
}
