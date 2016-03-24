package dtos

import "io"

type RequestDeployState struct {
	ActiveDeploy DeployMarker
	PendingDeploy DeployMarker
	RequestId string
}

func (self *RequestDeployState) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *RequestDeployState) FormatText() string {
	return FormatText(self)
}

func (self *RequestDeployState) FormatJSON() string {
	return FormatJSON(self)
}
