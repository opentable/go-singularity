package singularity

import "github.com/opentable/go-singularity/v0.14/dtos"

func (client *Client) DisabledActions() (response dtos.SingularityDisabledActionList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityDisabledActionList, 0)
	err = client.DTORequest(&response, "GET", "/api/disasters/disabled-actions", pathParamMap, queryParamMap)

	return
}
