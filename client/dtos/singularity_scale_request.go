package dtos

import "io"

type SingularityScaleRequest struct {
	ActionId         string
	DurationMillis   int64
	Instances        int32
	Message          string
	SkipHealthchecks bool
}

func (self *SingularityScaleRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityScaleRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityScaleRequest) FormatJSON() string {
	return FormatJSON(self)
}
