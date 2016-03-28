package client

func (client *Client) setNotLeader() (err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/test/notleader", pathParamMap, queryParamMap)

	return
}
