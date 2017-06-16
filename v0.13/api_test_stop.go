package singularity

func (client *Client) Stop() (err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/test/stop", pathParamMap, queryParamMap)

	return
}
