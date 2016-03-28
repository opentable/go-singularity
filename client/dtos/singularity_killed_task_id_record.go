package dtos

import "io"

type SingularityKilledTaskIdRecord struct {
	OriginalTimestamp int64
	//	RequestCleanupType *RequestCleanupType
	Retries int32
	//	TaskCleanupType *TaskCleanupType
	TaskId    *SingularityTaskId
	Timestamp int64
}

func (self *SingularityKilledTaskIdRecord) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityKilledTaskIdRecord) FormatText() string {
	return FormatText(self)
}

func (self *SingularityKilledTaskIdRecord) FormatJSON() string {
	return FormatJSON(self)
}
