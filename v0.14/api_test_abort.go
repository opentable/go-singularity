package singularity

func (client *Client) Abort() (err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/test/abort", pathParamMap, queryParamMap)

	return
}
