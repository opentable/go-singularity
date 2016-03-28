package dtos

import "io"

type SingularityBounceRequest struct {
	ActionId         string
	DurationMillis   int64
	Incremental      bool
	Message          string
	SkipHealthchecks bool
}

func (self *SingularityBounceRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityBounceRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityBounceRequest) FormatJSON() string {
	return FormatJSON(self)
}
