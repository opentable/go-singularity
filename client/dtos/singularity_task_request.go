package dtos

import "io"

type SingularityTaskRequest struct {
	Deploy      *SingularityDeploy
	PendingTask *SingularityPendingTask
	Request     *SingularityRequest
}

func (self *SingularityTaskRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityTaskRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityTaskRequest) FormatJSON() string {
	return FormatJSON(self)
}
