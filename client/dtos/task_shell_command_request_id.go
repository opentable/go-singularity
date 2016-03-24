package dtos

import "io"

type TaskShellCommandRequestId struct {
	Id string
	Name string
	TaskId TaskId
	Timestamp int64
}

func (self *TaskShellCommandRequestId) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *TaskShellCommandRequestId) FormatText() string {
	return FormatText(self)
}

func (self *TaskShellCommandRequestId) FormatJSON() string {
	return FormatJSON(self)
}
