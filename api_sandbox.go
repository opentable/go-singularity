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
	err = client.DTORequest("singularity-browse", response, "GET", "/api/sandbox/{taskId}/browse", pathParamMap, queryParamMap)

	return
}
func (client *Client) Read(taskId string, path string, grep string, offset int64, length int64) (response *dtos.MesosFileChunkObject, err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}

	queryParamMap := map[string]interface{}{
		"path": path, "grep": grep, "offset": offset, "length": length,
	}

	response = new(dtos.MesosFileChunkObject)
	err = client.DTORequest("singularity-read", response, "GET", "/api/sandbox/{taskId}/read", pathParamMap, queryParamMap)

	return
}
