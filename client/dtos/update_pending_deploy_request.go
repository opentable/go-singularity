package dtos

import "io"

type UpdatePendingDeployRequest struct {
	DeployId string
	RequestId string
	TargetActiveInstances int32
}

func (self *UpdatePendingDeployRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *UpdatePendingDeployRequest) FormatText() string {
	return FormatText(self)
}

func (self *UpdatePendingDeployRequest) FormatJSON() string {
	return FormatJSON(self)
}
