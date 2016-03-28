package dtos

import "io"

type SingularityUpdatePendingDeployRequest struct {
	DeployId              string
	RequestId             string
	TargetActiveInstances int32
}

func (self *SingularityUpdatePendingDeployRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityUpdatePendingDeployRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityUpdatePendingDeployRequest) FormatJSON() string {
	return FormatJSON(self)
}
