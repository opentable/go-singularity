package dtos

import "io"

type DeployRequest struct {
	Deploy Deploy
	Message string
	UnpauseOnSuccessfulDeploy bool
}

func (self *DeployRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *DeployRequest) FormatText() string {
	return FormatText(self)
}

func (self *DeployRequest) FormatJSON() string {
	return FormatJSON(self)
}
