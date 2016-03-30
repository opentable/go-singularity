package dtos

import "io"

type SingularityDeployHistory struct {
	Deploy           *SingularityDeploy
	DeployMarker     *SingularityDeployMarker
	DeployResult     *SingularityDeployResult
	DeployStatistics *SingularityDeployStatistics
}

func (self *SingularityDeployHistory) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityDeployHistory) FormatText() string {
	return FormatText(self)
}

func (self *SingularityDeployHistory) FormatJSON() string {
	return FormatJSON(self)
}
