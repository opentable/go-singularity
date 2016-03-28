package dtos

import "io"

type SingularityTaskId struct {
	DeployId        string
	Host            string
	Id              string
	InstanceNo      int32
	RackId          string
	RequestId       string
	SanitizedHost   string
	SanitizedRackId string
	StartedAt       int64
}

func (self *SingularityTaskId) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityTaskId) FormatText() string {
	return FormatText(self)
}

func (self *SingularityTaskId) FormatJSON() string {
	return FormatJSON(self)
}
