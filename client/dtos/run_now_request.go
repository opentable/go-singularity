package dtos

import "io"

type RunNowRequest struct {
	CommandLineArgs []string
	Message string
	RunId string
	SkipHealthchecks bool
}

func (self *RunNowRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *RunNowRequest) FormatText() string {
	return FormatText(self)
}

func (self *RunNowRequest) FormatJSON() string {
	return FormatJSON(self)
}
