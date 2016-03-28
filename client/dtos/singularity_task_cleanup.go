package dtos

import "io"

type SingularityTaskCleanup struct {
	ActionId string
	//	CleanupType *TaskCleanupType
	Message   string
	TaskId    *SingularityTaskId
	Timestamp int64
	User      string
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
