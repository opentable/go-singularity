package client

func (client *Client) setLeader() (err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/test/leader", pathParamMap, queryParamMap)

	return
}
