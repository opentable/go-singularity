package dtos

import "io"

type ScaleRequest struct {
	ActionId string
	DurationMillis int64
	Instances int32
	Message string
	SkipHealthchecks bool
}

func (self *ScaleRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *ScaleRequest) FormatText() string {
	return FormatText(self)
}

func (self *ScaleRequest) FormatJSON() string {
	return FormatJSON(self)
}
