package singularity

func (client *Client) RunHistoryPurge() (err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/test/purge-history", pathParamMap, queryParamMap)

	return
}
