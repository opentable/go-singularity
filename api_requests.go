package singularity

import (
	"github.com/opentable/go-singularity/dtos"
	"github.com/opentable/swaggering"
)

func (client *Client) GetRequest(requestId string, useWebCache bool) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "GET", "/api/requests/request/{requestId}", pathParamMap, queryParamMap)

	return
}
func (client *Client) DeleteRequest(requestId string, body *dtos.SingularityDeleteRequestRequest) (response *dtos.SingularityRequest, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequest)
	err = client.DTORequest(response, "DELETE", "/api/requests/request/{requestId}", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) GetRequests(useWebCache bool) (response dtos.SingularityRequestParentList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(dtos.SingularityRequestParentList, 0)
	err = client.DTORequest(&response, "GET", "/api/requests", pathParamMap, queryParamMap)

	return
}
func (client *Client) PostRequest(body *dtos.SingularityRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "POST", "/api/requests", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) Bounce(requestId string, body *dtos.SingularityBounceRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "POST", "/api/requests/request/{requestId}/bounce", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) DeleteExpiringBounce(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "DELETE", "/api/requests/request/{requestId}/bounce", pathParamMap, queryParamMap)

	return
}
func (client *Client) ScheduleImmediately(requestId string, body *dtos.SingularityRunNowRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "POST", "/api/requests/request/{requestId}/run", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) GetTaskByRunId(requestId string, runId string) (response *dtos.SingularityTaskId, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "runId": runId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskId)
	err = client.DTORequest(response, "GET", "/api/requests/request/{requestId}/run/{runId}", pathParamMap, queryParamMap)

	return
}
func (client *Client) Pause(requestId string, body *dtos.SingularityPauseRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "POST", "/api/requests/request/{requestId}/pause", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) DeleteExpiringPause(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "DELETE", "/api/requests/request/{requestId}/pause", pathParamMap, queryParamMap)

	return
}
func (client *Client) SkipHealthchecks(requestId string, body *dtos.SingularitySkipHealthchecksRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "PUT", "/api/requests/request/{requestId}/skip-healthchecks", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) DeleteExpiringSkipHealthchecks(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "DELETE", "/api/requests/request/{requestId}/skip-healthchecks", pathParamMap, queryParamMap)

	return
}
func (client *Client) Unpause(requestId string, body *dtos.SingularityUnpauseRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "POST", "/api/requests/request/{requestId}/unpause", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) ExitCooldown(requestId string, body *dtos.SingularityExitCooldownRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "POST", "/api/requests/request/{requestId}/exit-cooldown", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) GetActiveRequests(useWebCache bool) (response dtos.SingularityRequestParentList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(dtos.SingularityRequestParentList, 0)
	err = client.DTORequest(&response, "GET", "/api/requests/active", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetPausedRequests(useWebCache bool) (response dtos.SingularityRequestParentList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(dtos.SingularityRequestParentList, 0)
	err = client.DTORequest(&response, "GET", "/api/requests/paused", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetCooldownRequests(useWebCache bool) (response dtos.SingularityRequestParentList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(dtos.SingularityRequestParentList, 0)
	err = client.DTORequest(&response, "GET", "/api/requests/cooldown", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetFinishedRequests(useWebCache bool) (response dtos.SingularityRequestParentList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(dtos.SingularityRequestParentList, 0)
	err = client.DTORequest(&response, "GET", "/api/requests/finished", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetPendingRequests() (response dtos.SingularityPendingRequestList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityPendingRequestList, 0)
	err = client.DTORequest(&response, "GET", "/api/requests/queued/pending", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetCleanupRequests() (response dtos.SingularityRequestCleanupList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityRequestCleanupList, 0)
	err = client.DTORequest(&response, "GET", "/api/requests/queued/cleanup", pathParamMap, queryParamMap)

	return
}
func (client *Client) DeleteExpiringScale(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "DELETE", "/api/requests/request/{requestId}/scale", pathParamMap, queryParamMap)

	return
}
func (client *Client) Scale(requestId string, body *dtos.SingularityScaleRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "PUT", "/api/requests/request/{requestId}/scale", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) DeleteExpiringSkipHealthchecksDeprecated(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "DELETE", "/api/requests/request/{requestId}/skipHealthchecks", pathParamMap, queryParamMap)

	return
}
func (client *Client) SkipHealthchecksDeprecated(requestId string, body *dtos.SingularitySkipHealthchecksRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "PUT", "/api/requests/request/{requestId}/skipHealthchecks", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) GetLbCleanupRequests(useWebCache bool) (response swaggering.StringList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(swaggering.StringList, 0)
	err = client.DTORequest(&response, "GET", "/api/requests/lbcleanup", pathParamMap, queryParamMap)

	return
}
