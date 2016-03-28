package dtos

import "io"

type SingularityDeployUpdate struct {
	Deploy       *SingularityDeploy
	DeployMarker *SingularityDeployMarker
	DeployResult *SingularityDeployResult
	//	EventType *DeployEventType
}

func (self *SingularityDeployUpdate) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityDeployUpdate) FormatText() string {
	return FormatText(self)
}

func (self *SingularityDeployUpdate) FormatJSON() string {
	return FormatJSON(self)
}
