package singularity

import "github.com/opentable/go-singularity/v0.14/dtos"

func (client *Client) DisasterStats() (response *dtos.SingularityDisastersData, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityDisastersData)
	err = client.DTORequest(response, "GET", "/api/disasters/stats", pathParamMap, queryParamMap)

	return
}
