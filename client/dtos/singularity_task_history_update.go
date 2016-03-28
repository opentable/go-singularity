package dtos

import "io"

type SingularityTaskHistoryUpdate struct {
	StatusMessage string
	StatusReason  string
	TaskId        *SingularityTaskId
	//	TaskState *ExtendedTaskState
	Timestamp int64
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
