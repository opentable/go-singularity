package dtos

import "io"

type SingularityPendingTaskId struct {
	CreatedAt  int64
	DeployId   string
	Id         string
	InstanceNo int32
	NextRunAt  int64
	//	PendingType *PendingType
	RequestId string
}

func (self *SingularityPendingTaskId) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityPendingTaskId) FormatText() string {
	return FormatText(self)
}

func (self *SingularityPendingTaskId) FormatJSON() string {
	return FormatJSON(self)
}
