package dtos

import "io"

type SingularityDeployFailureSingularityDeployFailureReason string

const (
	SingularityDeployFailureSingularityDeployFailureReasonTASK_FAILED_ON_STARTUP         SingularityDeployFailureSingularityDeployFailureReason = "TASK_FAILED_ON_STARTUP"
	SingularityDeployFailureSingularityDeployFailureReasonTASK_FAILED_HEALTH_CHECKS      SingularityDeployFailureSingularityDeployFailureReason = "TASK_FAILED_HEALTH_CHECKS"
	SingularityDeployFailureSingularityDeployFailureReasonTASK_COULD_NOT_BE_SCHEDULED    SingularityDeployFailureSingularityDeployFailureReason = "TASK_COULD_NOT_BE_SCHEDULED"
	SingularityDeployFailureSingularityDeployFailureReasonTASK_NEVER_ENTERED_RUNNING     SingularityDeployFailureSingularityDeployFailureReason = "TASK_NEVER_ENTERED_RUNNING"
	SingularityDeployFailureSingularityDeployFailureReasonTASK_EXPECTED_RUNNING_FINISHED SingularityDeployFailureSingularityDeployFailureReason = "TASK_EXPECTED_RUNNING_FINISHED"
	SingularityDeployFailureSingularityDeployFailureReasonDEPLOY_CANCELLED               SingularityDeployFailureSingularityDeployFailureReason = "DEPLOY_CANCELLED"
	SingularityDeployFailureSingularityDeployFailureReasonDEPLOY_OVERDUE                 SingularityDeployFailureSingularityDeployFailureReason = "DEPLOY_OVERDUE"
	SingularityDeployFailureSingularityDeployFailureReasonFAILED_TO_SAVE_DEPLOY_STATE    SingularityDeployFailureSingularityDeployFailureReason = "FAILED_TO_SAVE_DEPLOY_STATE"
	SingularityDeployFailureSingularityDeployFailureReasonLOAD_BALANCER_UPDATE_FAILED    SingularityDeployFailureSingularityDeployFailureReason = "LOAD_BALANCER_UPDATE_FAILED"
	SingularityDeployFailureSingularityDeployFailureReasonPENDING_DEPLOY_REMOVED         SingularityDeployFailureSingularityDeployFailureReason = "PENDING_DEPLOY_REMOVED"
)

type SingularityDeployFailure struct {
	Message string                                                 `json:"message"`
	Reason  SingularityDeployFailureSingularityDeployFailureReason `json:"reason"`
	TaskId  *SingularityTaskId                                     `json:"taskId"`
}

func (self *SingularityDeployFailure) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityDeployFailure) FormatText() string {
	return FormatText(self)
}

func (self *SingularityDeployFailure) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityDeployFailureList []*SingularityDeployFailure

func (list *SingularityDeployFailureList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityDeployFailureList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDeployFailureList) FormatJSON() string {
	return FormatJSON(list)
}
