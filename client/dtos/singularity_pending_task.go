package dtos

import "io"

type SingularityPendingTask struct {
	CmdLineArgsList  string
	Message          string
	PendingTaskId    *SingularityPendingTaskId
	RunId            string
	SkipHealthchecks bool
	User             string
}

func (self *SingularityPendingTask) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityPendingTask) FormatText() string {
	return FormatText(self)
}

func (self *SingularityPendingTask) FormatJSON() string {
	return FormatJSON(self)
}
