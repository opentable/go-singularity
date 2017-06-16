package singularity

func (client *Client) StartTaskReconciliation() (err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/test/reconcile", pathParamMap, queryParamMap)

	return
}
