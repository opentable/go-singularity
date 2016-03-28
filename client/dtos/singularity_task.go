package dtos

import "io"

type SingularityTask struct {
	MesosTask   *TaskInfo
	Offer       *Offer
	RackId      string
	TaskId      *SingularityTaskId
	TaskRequest *SingularityTaskRequest
}

func (self *SingularityTask) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityTask) FormatText() string {
	return FormatText(self)
}

func (self *SingularityTask) FormatJSON() string {
	return FormatJSON(self)
}
