package dtos

import "io"

type DeployHistory struct {
	Deploy Deploy
	DeployMarker DeployMarker
	DeployResult DeployResult
	DeployStatistics DeployStatistics
}

func (self *DeployHistory) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *DeployHistory) FormatText() string {
	return FormatText(self)
}

func (self *DeployHistory) FormatJSON() string {
	return FormatJSON(self)
}
