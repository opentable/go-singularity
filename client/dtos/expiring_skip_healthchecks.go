package dtos

import "io"

type ExpiringSkipHealthchecks struct {
	ActionId string
//	ExpiringAPIRequestObject T
	RequestId string
	RevertToSkipHealthchecks bool
	StartMillis int64
	User string
}

func (self *ExpiringSkipHealthchecks) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *ExpiringSkipHealthchecks) FormatText() string {
	return FormatText(self)
}

func (self *ExpiringSkipHealthchecks) FormatJSON() string {
	return FormatJSON(self)
}
