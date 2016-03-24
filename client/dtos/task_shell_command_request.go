package dtos

import "io"

type TaskShellCommandRequest struct {
	ShellCommand ShellCommand
	TaskId TaskId
	Timestamp int64
	User string
}

func (self *TaskShellCommandRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *TaskShellCommandRequest) FormatText() string {
	return FormatText(self)
}

func (self *TaskShellCommandRequest) FormatJSON() string {
	return FormatJSON(self)
}
