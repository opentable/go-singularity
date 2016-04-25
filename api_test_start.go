package singularity

func (client *Client) Start() (err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/test/start", pathParamMap, queryParamMap)

	return
}
