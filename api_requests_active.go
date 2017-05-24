package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetActiveRequests(useWebCache bool) (response dtos.SingularityRequestParentList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(dtos.SingularityRequestParentList, 0)
	err = client.DTORequest(&response, "GET", "/api/requests/active", pathParamMap, queryParamMap)

	return
}
