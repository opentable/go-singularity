package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetCleanupRequests() (response dtos.SingularityRequestCleanupList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityRequestCleanupList, 0)
	err = client.DTORequest(&response, "GET", "/api/requests/queued/cleanup", pathParamMap, queryParamMap)

	return
}
