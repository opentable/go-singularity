package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) Read(taskId string, path string, grep string, offset int64, length int64) (response *dtos.MesosFileChunkObject, err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}

	queryParamMap := map[string]interface{}{
		"path": path, "grep": grep, "offset": offset, "length": length,
	}

	response = new(dtos.MesosFileChunkObject)
	err = client.DTORequest(response, "GET", "/api/sandbox/{taskId}/read", pathParamMap, queryParamMap)

	return
}
