package dtos

import "io"

type TaskHistoryUpdate struct {
	StatusMessage string
	StatusReason string
	TaskId TaskId
//	TaskState ExtendedTaskState
	Timestamp int64
}

func (self *TaskHistoryUpdate) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *TaskHistoryUpdate) FormatText() string {
	return FormatText(self)
}

func (self *TaskHistoryUpdate) FormatJSON() string {
	return FormatJSON(self)
}
