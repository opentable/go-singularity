package dtos

import "io"

type UnpauseRequest struct {
	ActionId string
	Message string
	SkipHealthchecks bool
}

func (self *UnpauseRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *UnpauseRequest) FormatText() string {
	return FormatText(self)
}

func (self *UnpauseRequest) FormatJSON() string {
	return FormatJSON(self)
}
