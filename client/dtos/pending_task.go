package dtos

import "io"

type PendingTask struct {
	CmdLineArgsList []string
	Message string
	PendingTaskId PendingTaskId
	RunId string
	SkipHealthchecks bool
	User string
}

func (self *PendingTask) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *PendingTask) FormatText() string {
	return FormatText(self)
}

func (self *PendingTask) FormatJSON() string {
	return FormatJSON(self)
}
