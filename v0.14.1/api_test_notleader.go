package singularity

func (client *Client) SetNotLeader() (err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/test/notleader", pathParamMap, queryParamMap)

	return
}
