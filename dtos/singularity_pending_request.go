package dtos

import "io"

type SingularityPendingRequest struct {
	ActionId        string
	CmdLineArgsList StringList
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

type SingularityPendingRequestList []*SingularityPendingRequest

func (list SingularityPendingRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list SingularityPendingRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list SingularityPendingRequestList) FormatJSON() string {
	return FormatJSON(list)
}
