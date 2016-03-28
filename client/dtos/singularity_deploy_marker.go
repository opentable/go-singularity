package dtos

import "io"

type SingularityDeployMarker struct {
	DeployId  string
	Message   string
	RequestId string
	Timestamp int64
	User      string
}

func (self *SingularityDeployMarker) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityDeployMarker) FormatText() string {
	return FormatText(self)
}

func (self *SingularityDeployMarker) FormatJSON() string {
	return FormatJSON(self)
}
