package dtos

import "io"

type SingularityTaskHistoryUpdateExtendedTaskState string

const (
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_LAUNCHED        SingularityTaskHistoryUpdateExtendedTaskState = "TASK_LAUNCHED"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_STAGING         SingularityTaskHistoryUpdateExtendedTaskState = "TASK_STAGING"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_STARTING        SingularityTaskHistoryUpdateExtendedTaskState = "TASK_STARTING"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_RUNNING         SingularityTaskHistoryUpdateExtendedTaskState = "TASK_RUNNING"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_CLEANING        SingularityTaskHistoryUpdateExtendedTaskState = "TASK_CLEANING"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_FINISHED        SingularityTaskHistoryUpdateExtendedTaskState = "TASK_FINISHED"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_FAILED          SingularityTaskHistoryUpdateExtendedTaskState = "TASK_FAILED"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_KILLED          SingularityTaskHistoryUpdateExtendedTaskState = "TASK_KILLED"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_LOST            SingularityTaskHistoryUpdateExtendedTaskState = "TASK_LOST"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_LOST_WHILE_DOWN SingularityTaskHistoryUpdateExtendedTaskState = "TASK_LOST_WHILE_DOWN"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_ERROR           SingularityTaskHistoryUpdateExtendedTaskState = "TASK_ERROR"
)

type SingularityTaskHistoryUpdate struct {
	StatusMessage string                                        `json:"statusMessage"`
	StatusReason  string                                        `json:"statusReason"`
	TaskId        *SingularityTaskId                            `json:"taskId"`
	TaskState     SingularityTaskHistoryUpdateExtendedTaskState `json:"taskState"`
	Timestamp     int64                                         `json:"timestamp"`
}

func (self *SingularityTaskHistoryUpdate) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityTaskHistoryUpdate) FormatText() string {
	return FormatText(self)
}

func (self *SingularityTaskHistoryUpdate) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityTaskHistoryUpdateList []*SingularityTaskHistoryUpdate

func (list *SingularityTaskHistoryUpdateList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskHistoryUpdateList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskHistoryUpdateList) FormatJSON() string {
	return FormatJSON(list)
}
