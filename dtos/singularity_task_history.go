package dtos

import "io"

type SingularityTaskHistory struct {
	Directory           string
	HealthcheckResults  *SingularityTaskHealthcheckResult
	LoadBalancerUpdates *SingularityLoadBalancerUpdate
	ShellCommandHistory *SingularityTaskShellCommandHistory
	Task                *SingularityTask
	TaskUpdates         *SingularityTaskHistoryUpdate
}

func (self *SingularityTaskHistory) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityTaskHistory) FormatText() string {
	return FormatText(self)
}

func (self *SingularityTaskHistory) FormatJSON() string {
	return FormatJSON(self)
}
