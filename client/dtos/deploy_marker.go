package dtos

import "io"

type DeployMarker struct {
	DeployId string
	Message string
	RequestId string
	Timestamp int64
	User string
}

func (self *DeployMarker) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *DeployMarker) FormatText() string {
	return FormatText(self)
}

func (self *DeployMarker) FormatJSON() string {
	return FormatJSON(self)
}
