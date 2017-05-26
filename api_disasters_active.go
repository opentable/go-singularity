package singularity

import "github.com/opentable/swaggering"

func (client *Client) ActiveDisasters() (response swaggering.StringList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(swaggering.StringList, 0)
	err = client.DTORequest(&response, "GET", "/api/disasters/active", pathParamMap, queryParamMap)

	return
}
