package dtos

import "io"

type ExpiringPause struct {
	ActionId string
//	ExpiringAPIRequestObject T
	RequestId string
	StartMillis int64
	User string
}

func (self *ExpiringPause) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *ExpiringPause) FormatText() string {
	return FormatText(self)
}

func (self *ExpiringPause) FormatJSON() string {
	return FormatJSON(self)
}
