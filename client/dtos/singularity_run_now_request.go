package dtos

import "io"

type SingularityRunNowRequest struct {
	CommandLineArgs  string
	Message          string
	RunId            string
	SkipHealthchecks bool
}

func (self *SingularityRunNowRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityRunNowRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityRunNowRequest) FormatJSON() string {
	return FormatJSON(self)
}
