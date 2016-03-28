package dtos

import "io"

type SingularityRequestCleanup struct {
	ActionId string
	//	CleanupType *RequestCleanupType
	DeployId         string
	KillTasks        bool
	Message          string
	RequestId        string
	SkipHealthchecks bool
	Timestamp        int64
	User             string
}

func (self *SingularityRequestCleanup) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityRequestCleanup) FormatText() string {
	return FormatText(self)
}

func (self *SingularityRequestCleanup) FormatJSON() string {
	return FormatJSON(self)
}
