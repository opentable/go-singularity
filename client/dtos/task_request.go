package dtos

import "io"

type TaskRequest struct {
	Deploy Deploy
	PendingTask PendingTask
	Request Request
}

func (self *TaskRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *TaskRequest) FormatText() string {
	return FormatText(self)
}

func (self *TaskRequest) FormatJSON() string {
	return FormatJSON(self)
}
