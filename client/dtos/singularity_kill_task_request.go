package dtos

import "io"

type SingularityKillTaskRequest struct {
	ActionId               string
	Message                string
	Override               bool
	WaitForReplacementTask bool
}

func (self *SingularityKillTaskRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityKillTaskRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityKillTaskRequest) FormatJSON() string {
	return FormatJSON(self)
}
