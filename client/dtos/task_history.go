package dtos

import "io"

type TaskHistory struct {
	Directory string
	HealthcheckResults []
	LoadBalancerUpdates []
	ShellCommandHistory []
	Task Task
	TaskUpdates []
}

func (self *TaskHistory) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *TaskHistory) FormatText() string {
	return FormatText(self)
}

func (self *TaskHistory) FormatJSON() string {
	return FormatJSON(self)
}
