package dtos

import "io"

type SingularityTaskCleanupTaskCleanupType string

const (
	SingularityTaskCleanupTaskCleanupTypeUSER_REQUESTED             SingularityTaskCleanupTaskCleanupType = "USER_REQUESTED"
	SingularityTaskCleanupTaskCleanupTypeUSER_REQUESTED_TASK_BOUNCE SingularityTaskCleanupTaskCleanupType = "USER_REQUESTED_TASK_BOUNCE"
	SingularityTaskCleanupTaskCleanupTypeDECOMISSIONING             SingularityTaskCleanupTaskCleanupType = "DECOMISSIONING"
	SingularityTaskCleanupTaskCleanupTypeSCALING_DOWN               SingularityTaskCleanupTaskCleanupType = "SCALING_DOWN"
	SingularityTaskCleanupTaskCleanupTypeBOUNCING                   SingularityTaskCleanupTaskCleanupType = "BOUNCING"
	SingularityTaskCleanupTaskCleanupTypeINCREMENTAL_BOUNCE         SingularityTaskCleanupTaskCleanupType = "INCREMENTAL_BOUNCE"
	SingularityTaskCleanupTaskCleanupTypeDEPLOY_FAILED              SingularityTaskCleanupTaskCleanupType = "DEPLOY_FAILED"
	SingularityTaskCleanupTaskCleanupTypeNEW_DEPLOY_SUCCEEDED       SingularityTaskCleanupTaskCleanupType = "NEW_DEPLOY_SUCCEEDED"
	SingularityTaskCleanupTaskCleanupTypeDEPLOY_STEP_FINISHED       SingularityTaskCleanupTaskCleanupType = "DEPLOY_STEP_FINISHED"
	SingularityTaskCleanupTaskCleanupTypeDEPLOY_CANCELED            SingularityTaskCleanupTaskCleanupType = "DEPLOY_CANCELED"
	SingularityTaskCleanupTaskCleanupTypeUNHEALTHY_NEW_TASK         SingularityTaskCleanupTaskCleanupType = "UNHEALTHY_NEW_TASK"
	SingularityTaskCleanupTaskCleanupTypeOVERDUE_NEW_TASK           SingularityTaskCleanupTaskCleanupType = "OVERDUE_NEW_TASK"
)

type SingularityTaskCleanup struct {
	ActionId    string                                `json:"actionId"`
	CleanupType SingularityTaskCleanupTaskCleanupType `json:"cleanupType"`
	Message     string                                `json:"message"`
	TaskId      *SingularityTaskId                    `json:"taskId"`
	Timestamp   int64                                 `json:"timestamp"`
	User        string                                `json:"user"`
}

func (self *SingularityTaskCleanup) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityTaskCleanup) FormatText() string {
	return FormatText(self)
}

func (self *SingularityTaskCleanup) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityTaskCleanupList []*SingularityTaskCleanup

func (list *SingularityTaskCleanupList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskCleanupList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskCleanupList) FormatJSON() string {
	return FormatJSON(list)
}
