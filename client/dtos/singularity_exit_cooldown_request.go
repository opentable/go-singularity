package dtos

import "io"

type SingularityExitCooldownRequest struct {
	ActionId         string
	Message          string
	SkipHealthchecks bool
}

func (self *SingularityExitCooldownRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityExitCooldownRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityExitCooldownRequest) FormatJSON() string {
	return FormatJSON(self)
}
