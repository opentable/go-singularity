package singularity

import "github.com/opentable/go-singularity/dtos"

func (client *Client) GetScheduledTaskIds(useWebCache bool) (response dtos.SingularityPendingTaskIdList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(dtos.SingularityPendingTaskIdList, 0)
	err = client.DTORequest(&response, "GET", "/api/tasks/scheduled/ids", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetScheduledTasks(useWebCache bool) (response dtos.SingularityTaskRequestList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(dtos.SingularityTaskRequestList, 0)
	err = client.DTORequest(&response, "GET", "/api/tasks/scheduled", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetTaskCleanup(taskId string) (response *dtos.SingularityTaskCleanup, err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskCleanup)
	err = client.DTORequest(response, "GET", "/api/tasks/task/{taskId}/cleanup", pathParamMap, queryParamMap)

	return
}
func (client *Client) KillTask(taskId string, body *dtos.SingularityKillTaskRequest) (response *dtos.SingularityTaskCleanup, err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskCleanup)
	err = client.DTORequest(response, "DELETE", "/api/tasks/task/{taskId}", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) GetActiveTask(taskId string) (response *dtos.SingularityTask, err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTask)
	err = client.DTORequest(response, "GET", "/api/tasks/task/{taskId}", pathParamMap, queryParamMap)

	return
}

// PostTaskMetadata is invalid
func (client *Client) RunShellCommand(taskId string, body *dtos.SingularityShellCommand) (response *dtos.SingularityTaskShellCommandRequest, err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskShellCommandRequest)
	err = client.DTORequest(response, "POST", "/api/tasks/task/{taskId}/command", pathParamMap, queryParamMap, body)

	return
}
func (client *Client) GetPendingTask(pendingTaskId string) (response *dtos.SingularityTaskRequest, err error) {
	pathParamMap := map[string]interface{}{
		"pendingTaskId": pendingTaskId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskRequest)
	err = client.DTORequest(response, "GET", "/api/tasks/scheduled/task/{pendingTaskId}", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetScheduledTasksForRequest(requestId string, useWebCache bool) (response dtos.SingularityTaskRequestList, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(dtos.SingularityTaskRequestList, 0)
	err = client.DTORequest(&response, "GET", "/api/tasks/scheduled/request/{requestId}", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetTasksForSlave(slaveId string, useWebCache bool) (response dtos.SingularityTaskList, err error) {
	pathParamMap := map[string]interface{}{
		"slaveId": slaveId,
	}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(dtos.SingularityTaskList, 0)
	err = client.DTORequest(&response, "GET", "/api/tasks/active/slave/{slaveId}", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetActiveTasks(useWebCache bool) (response dtos.SingularityTaskList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(dtos.SingularityTaskList, 0)
	err = client.DTORequest(&response, "GET", "/api/tasks/active", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetCleaningTasks(useWebCache bool) (response dtos.SingularityTaskCleanupList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{
		"useWebCache": useWebCache,
	}

	response = make(dtos.SingularityTaskCleanupList, 0)
	err = client.DTORequest(&response, "GET", "/api/tasks/cleaning", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetKilledTasks() (response dtos.SingularityKilledTaskIdRecordList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityKilledTaskIdRecordList, 0)
	err = client.DTORequest(&response, "GET", "/api/tasks/killed", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetLbCleanupTasks() (response dtos.SingularityTaskIdList, err error) {
	pathParamMap := map[string]interface{}{}

	queryParamMap := map[string]interface{}{}

	response = make(dtos.SingularityTaskIdList, 0)
	err = client.DTORequest(&response, "GET", "/api/tasks/lbcleanup", pathParamMap, queryParamMap)

	return
}
func (client *Client) GetTaskStatistics(taskId string) (response *dtos.MesosTaskStatisticsObject, err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.MesosTaskStatisticsObject)
	err = client.DTORequest(response, "GET", "/api/tasks/task/{taskId}/statistics", pathParamMap, queryParamMap)

	return
}
