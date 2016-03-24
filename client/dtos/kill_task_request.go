package dtos

import "io"

type KillTaskRequest struct {
	ActionId string
	Message string
	Override bool
	WaitForReplacementTask bool
}

func (self *KillTaskRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *KillTaskRequest) FormatText() string {
	return FormatText(self)
}

func (self *KillTaskRequest) FormatJSON() string {
	return FormatJSON(self)
}
