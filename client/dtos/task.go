package dtos

import "io"

type Task struct {
	MesosTask TaskInfo
	Offer Offer
	RackId string
	TaskId TaskId
	TaskRequest TaskRequest
}

func (self *Task) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *Task) FormatText() string {
	return FormatText(self)
}

func (self *Task) FormatJSON() string {
	return FormatJSON(self)
}
