package dtos

import "io"

type ExpiringBounce struct {
	ActionId string
	DeployId string
//	ExpiringAPIRequestObject T
	RequestId string
	StartMillis int64
	User string
}

func (self *ExpiringBounce) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *ExpiringBounce) FormatText() string {
	return FormatText(self)
}

func (self *ExpiringBounce) FormatJSON() string {
	return FormatJSON(self)
}
