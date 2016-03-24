package dtos

import "io"

type SkipHealthchecksRequest struct {
	ActionId string
	DurationMillis int64
	Message string
	SkipHealthchecks bool
}

func (self *SkipHealthchecksRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SkipHealthchecksRequest) FormatText() string {
	return FormatText(self)
}

func (self *SkipHealthchecksRequest) FormatJSON() string {
	return FormatJSON(self)
}
