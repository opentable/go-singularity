package singularity

func (client *Client) EnableAutomatedDisasterCreation() (err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/disasters/enable", pathParamMap, queryParamMap)

	return
}
