package dtos

import "io"

type SingularityDeployFailure struct {
	Message string
	//	Reason *SingularityDeployFailureReason
	TaskId *SingularityTaskId
}

func (self *SingularityDeployFailure) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityDeployFailure) FormatText() string {
	return FormatText(self)
}

func (self *SingularityDeployFailure) FormatJSON() string {
	return FormatJSON(self)
}
