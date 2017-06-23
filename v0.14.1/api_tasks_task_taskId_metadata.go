package singularity

import "github.com/opentable/go-singularity/v0.14.1/dtos"

func (client *Client) PostTaskMetadata(taskId string, body *dtos.SingularityTaskMetadataRequest) (err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/tasks/task/{taskId}/metadata", pathParamMap, queryParamMap, body)

	return
}
