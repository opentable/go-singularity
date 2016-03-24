package dtos

import "io"

type BounceRequest struct {
	ActionId string
	DurationMillis int64
	Incremental bool
	Message string
	SkipHealthchecks bool
}

func (self *BounceRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *BounceRequest) FormatText() string {
	return FormatText(self)
}

func (self *BounceRequest) FormatJSON() string {
	return FormatJSON(self)
}
