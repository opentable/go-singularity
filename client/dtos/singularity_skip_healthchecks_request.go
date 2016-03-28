package dtos

import "io"

type SingularitySkipHealthchecksRequest struct {
	ActionId         string
	DurationMillis   int64
	Message          string
	SkipHealthchecks bool
}

func (self *SingularitySkipHealthchecksRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularitySkipHealthchecksRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularitySkipHealthchecksRequest) FormatJSON() string {
	return FormatJSON(self)
}
