package dtos

import "io"

type SingularityUnpauseRequest struct {
	ActionId         string
	Message          string
	SkipHealthchecks bool
}

func (self *SingularityUnpauseRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityUnpauseRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityUnpauseRequest) FormatJSON() string {
	return FormatJSON(self)
}
