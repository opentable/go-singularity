package dtos

import "io"

type PendingTaskId struct {
	CreatedAt int64
	DeployId string
	Id string
	InstanceNo int32
	NextRunAt int64
//	PendingType PendingType
	RequestId string
}

func (self *PendingTaskId) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *PendingTaskId) FormatText() string {
	return FormatText(self)
}

func (self *PendingTaskId) FormatJSON() string {
	return FormatJSON(self)
}
