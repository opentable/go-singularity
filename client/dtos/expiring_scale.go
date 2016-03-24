package dtos

import "io"

type ExpiringScale struct {
	ActionId string
//	ExpiringAPIRequestObject T
	RequestId string
	RevertToInstances int32
	StartMillis int64
	User string
}

func (self *ExpiringScale) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *ExpiringScale) FormatText() string {
	return FormatText(self)
}

func (self *ExpiringScale) FormatJSON() string {
	return FormatJSON(self)
}
