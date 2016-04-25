package singularity

func (client *Client) StatusUpdate(taskId string, taskState string) (err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId, "taskState": taskState,
	}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/test/scheduler/statusUpdate/{taskId}/{taskState}", pathParamMap, queryParamMap)

	return
}
