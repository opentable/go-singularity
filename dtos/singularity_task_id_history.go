package dtos

import "io"

type SingularityTaskIdHistory struct {
	//	LastTaskState *ExtendedTaskState
	RunId     string
	TaskId    *SingularityTaskId
	UpdatedAt int64
}

func (self *SingularityTaskIdHistory) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityTaskIdHistory) FormatText() string {
	return FormatText(self)
}

func (self *SingularityTaskIdHistory) FormatJSON() string {
	return FormatJSON(self)
}
