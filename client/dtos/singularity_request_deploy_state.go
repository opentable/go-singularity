package dtos

import "io"

type SingularityRequestDeployState struct {
	ActiveDeploy  *SingularityDeployMarker
	PendingDeploy *SingularityDeployMarker
	RequestId     string
}

func (self *SingularityRequestDeployState) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityRequestDeployState) FormatText() string {
	return FormatText(self)
}

func (self *SingularityRequestDeployState) FormatJSON() string {
	return FormatJSON(self)
}
