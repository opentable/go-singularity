package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetPaginatedS3Logs(body *dtos.SingularityS3SearchRequest) (response dtos.SingularityS3LogList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityS3LogList, 0)
	err = client.DTORequest(&response, "POST", "/api/logs/search", pathParamMap, queryParamMap, body)

	return
}
