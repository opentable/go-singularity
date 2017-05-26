package singularity

func (client *Client) DisableAutomatedDisasterCreation() (err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/disasters/disable", pathParamMap, queryParamMap)

	return
}
