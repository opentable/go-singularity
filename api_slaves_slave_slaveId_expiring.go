package singularity

func (client *Client) DeleteExpiringSlaveStateChange(slaveId string) (err error) {
	pathParamMap := map[string]interface{}{
		"slaveId": slaveId,
	}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("DELETE", "/api/slaves/slave/{slaveId}/expiring", pathParamMap, queryParamMap)

	return
}
