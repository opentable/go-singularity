package dtos

import "io"

type DeployFailure struct {
	Message string
//	Reason DeployFailureReason
	TaskId TaskId
}

func (self *DeployFailure) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *DeployFailure) FormatText() string {
	return FormatText(self)
}

func (self *DeployFailure) FormatJSON() string {
	return FormatJSON(self)
}
