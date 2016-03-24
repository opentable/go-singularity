package dtos

import "io"

type TaskCleanup struct {
	ActionId string
//	CleanupType TaskCleanupType
	Message string
	TaskId TaskId
	Timestamp int64
	User string
}

func (self *TaskCleanup) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *TaskCleanup) FormatText() string {
	return FormatText(self)
}

func (self *TaskCleanup) FormatJSON() string {
	return FormatJSON(self)
}
