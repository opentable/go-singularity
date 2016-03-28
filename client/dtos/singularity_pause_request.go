package dtos

import "io"

type SingularityPauseRequest struct {
	ActionId       string
	DurationMillis int64
	KillTasks      bool
	Message        string
}

func (self *SingularityPauseRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityPauseRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityPauseRequest) FormatJSON() string {
	return FormatJSON(self)
}
