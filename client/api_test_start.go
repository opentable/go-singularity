package client

func (client *Client) start() (err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/test/start", pathParamMap, queryParamMap)

	return
}
