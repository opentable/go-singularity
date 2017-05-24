package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) AddTaskCredits(credits int32) (err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"credits": credits,
	}

	_, err = client.Request("POST", "/api/disasters/task-credits", pathParamMap, queryParamMap)

	return
}

func (client *Client) DisableTaskCredits() (err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	_, err = client.Request("DELETE", "/api/disasters/task-credits", pathParamMap, queryParamMap)

	return
}

func (client *Client) GetTaskCreditData() (response *dtos.SingularityTaskCredits, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskCredits)
	err = client.DTORequest(response, "GET", "/api/disasters/task-credits", pathParamMap, queryParamMap)

	return
}
