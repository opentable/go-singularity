package dtos

import "io"

type SingularityPendingRequest struct {
	ActionId        string
	CmdLineArgsList string
	DeployId        string
	Message         string
	//	PendingType *PendingType
	RequestId        string
	RunId            string
	SkipHealthchecks bool
	Timestamp        int64
	User             string
}

func (self *SingularityPendingRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityPendingRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityPendingRequest) FormatJSON() string {
	return FormatJSON(self)
}
