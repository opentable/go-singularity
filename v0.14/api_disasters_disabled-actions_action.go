package singularity

import "github.com/opentable/go-singularity/v0.14/dtos"

func (client *Client) DisableAction(action, body *dtos.SingularityDisabledActionRequest) (err error) {
	pathParamMap := map[string]interface{}{
		"action": action,
	}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/disasters/disabled-actions/{action}", pathParamMap, queryParamMap, body)

	return
}

func (client *Client) EnableAction(action) (err error) {
	pathParamMap := map[string]interface{}{
		"action": action,
	}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("DELETE", "/api/disasters/disabled-actions/{action}", pathParamMap, queryParamMap)

	return
}
