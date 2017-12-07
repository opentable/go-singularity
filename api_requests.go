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
	err = client.DTORequest("singularity-getrequest", response, "GET", "/api/requests/request/{requestId}", pathParamMap, queryParamMap)

	return
}
func (client *Client) DeleteRequest(requestId string, body *dtos.SingularityDeleteRequestRequest) (response *dtos.SingularityRequest, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequest)
	err = client.DTORequest("singularity-deleterequest", response, "DELETE", "/api/requests/request/{requestId}", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) GetRequests(useWebCache bool) (response dtos.SingularityRequestParentList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(dtos.SingularityRequestParentList, 0)
	err = client.DTORequest("singularity-getrequests", &response, "GET", "/api/requests", pathParamMap, queryParamMap)

	return
}
func (client *Client) PostRequest(body *dtos.SingularityRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest("singularity-postrequest", response, "POST", "/api/requests", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) Bounce(requestId string, body *dtos.SingularityBounceRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest("singularity-bounce", response, "POST", "/api/requests/request/{requestId}/bounce", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) DeleteExpiringBounce(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest("singularity-deleteexpiringbounce", response, "DELETE", "/api/requests/request/{requestId}/bounce", pathParamMap, queryParamMap)

	return
}
func (client *Client) ScheduleImmediately(requestId string, body *dtos.SingularityRunNowRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest("singularity-scheduleimmediately", response, "POST", "/api/requests/request/{requestId}/run", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) GetTaskByRunId(requestId string, runId string) (response *dtos.SingularityTaskId, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "runId": runId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskId)
	err = client.DTORequest("singularity-gettaskbyrunid", response, "GET", "/api/requests/request/{requestId}/run/{runId}", pathParamMap, queryParamMap)

	return
}
func (client *Client) Pause(requestId string, body *dtos.SingularityPauseRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest("singularity-pause", response, "POST", "/api/requests/request/{requestId}/pause", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) DeleteExpiringPause(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest("singularity-deleteexpiringpause", response, "DELETE", "/api/requests/request/{requestId}/pause", pathParamMap, queryParamMap)

	return
}
func (client *Client) SkipHealthchecks(requestId string, body *dtos.SingularitySkipHealthchecksRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest("singularity-skiphealthchecks", response, "PUT", "/api/requests/request/{requestId}/skip-healthchecks", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) DeleteExpiringSkipHealthchecks(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest("singularity-deleteexpiringskiphealthchecks", response, "DELETE", "/api/requests/request/{requestId}/skip-healthchecks", pathParamMap, queryParamMap)

	return
}
func (client *Client) Unpause(requestId string, body *dtos.SingularityUnpauseRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest("singularity-unpause", response, "POST", "/api/requests/request/{requestId}/unpause", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) ExitCooldown(requestId string, body *dtos.SingularityExitCooldownRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest("singularity-exitcooldown", response, "POST", "/api/requests/request/{requestId}/exit-cooldown", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) GetActiveRequests(useWebCache bool) (response dtos.SingularityRequestParentList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(dtos.SingularityRequestParentList, 0)
	err = client.DTORequest("singularity-getactiverequests", &response, "GET", "/api/requests/active", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetPausedRequests(useWebCache bool) (response dtos.SingularityRequestParentList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(dtos.SingularityRequestParentList, 0)
	err = client.DTORequest("singularity-getpausedrequests", &response, "GET", "/api/requests/paused", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetCooldownRequests(useWebCache bool) (response dtos.SingularityRequestParentList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(dtos.SingularityRequestParentList, 0)
	err = client.DTORequest("singularity-getcooldownrequests", &response, "GET", "/api/requests/cooldown", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetFinishedRequests(useWebCache bool) (response dtos.SingularityRequestParentList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(dtos.SingularityRequestParentList, 0)
	err = client.DTORequest("singularity-getfinishedrequests", &response, "GET", "/api/requests/finished", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetPendingRequests() (response dtos.SingularityPendingRequestList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityPendingRequestList, 0)
	err = client.DTORequest("singularity-getpendingrequests", &response, "GET", "/api/requests/queued/pending", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetCleanupRequests() (response dtos.SingularityRequestCleanupList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityRequestCleanupList, 0)
	err = client.DTORequest("singularity-getcleanuprequests", &response, "GET", "/api/requests/queued/cleanup", pathParamMap, queryParamMap)

	return
}
func (client *Client) DeleteExpiringScale(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest("singularity-deleteexpiringscale", response, "DELETE", "/api/requests/request/{requestId}/scale", pathParamMap, queryParamMap)

	return
}
func (client *Client) Scale(requestId string, body *dtos.SingularityScaleRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest("singularity-scale", response, "PUT", "/api/requests/request/{requestId}/scale", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) DeleteExpiringSkipHealthchecksDeprecated(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest("singularity-deleteexpiringskiphealthchecksdeprecated", response, "DELETE", "/api/requests/request/{requestId}/skipHealthchecks", pathParamMap, queryParamMap)

	return
}
func (client *Client) SkipHealthchecksDeprecated(requestId string, body *dtos.SingularitySkipHealthchecksRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest("singularity-skiphealthchecksdeprecated", response, "PUT", "/api/requests/request/{requestId}/skipHealthchecks", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) GetLbCleanupRequests(useWebCache bool) (response swaggering.StringList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(swaggering.StringList, 0)
	err = client.DTORequest("singularity-getlbcleanuprequests", &response, "GET", "/api/requests/lbcleanup", pathParamMap, queryParamMap)

	return
}
