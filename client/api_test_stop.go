package client

func (client *Client) stop() (err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/test/stop", pathParamMap, queryParamMap)

	return
}
