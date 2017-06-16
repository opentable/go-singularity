package singularity

func (client *Client) DeleteExpiringRackStateChange(rackId string) (err error) {
	pathParamMap := map[string]interface{}{
		"rackId": rackId,
	}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("DELETE", "/api/racks/rack/{rackId}/expiring", pathParamMap, queryParamMap)

	return
}
