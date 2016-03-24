package dtos

import "io"

type PauseRequest struct {
	ActionId string
	DurationMillis int64
	KillTasks bool
	Message string
}

func (self *PauseRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *PauseRequest) FormatText() string {
	return FormatText(self)
}

func (self *PauseRequest) FormatJSON() string {
	return FormatJSON(self)
}
