package client

func (client *Client) abort() (err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/test/abort", pathParamMap, queryParamMap)

	return
}
