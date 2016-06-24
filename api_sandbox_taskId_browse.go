package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) Browse(taskId string, path string) (response *dtos.SingularitySandbox, err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}

	queryParamMap := map[string]interface{}{
		"path": path,
	}

	response = new(dtos.SingularitySandbox)
	err = client.DTORequest(response, "GET", "/api/sandbox/{taskId}/browse", pathParamMap, queryParamMap)

	return
}
