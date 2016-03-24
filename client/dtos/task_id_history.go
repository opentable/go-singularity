package dtos

import "io"

type TaskIdHistory struct {
//	LastTaskState ExtendedTaskState
	RunId string
	TaskId TaskId
	UpdatedAt int64
}

func (self *TaskIdHistory) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *TaskIdHistory) FormatText() string {
	return FormatText(self)
}

func (self *TaskIdHistory) FormatJSON() string {
	return FormatJSON(self)
}
