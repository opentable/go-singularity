package dtos

import "io"

type SingularityDeployRequest struct {
	Deploy                    *SingularityDeploy
	Message                   string
	UnpauseOnSuccessfulDeploy bool
}

func (self *SingularityDeployRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityDeployRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityDeployRequest) FormatJSON() string {
	return FormatJSON(self)
}
