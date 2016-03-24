package dtos

import "io"

type ExitCooldownRequest struct {
	ActionId string
	Message string
	SkipHealthchecks bool
}

func (self *ExitCooldownRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *ExitCooldownRequest) FormatText() string {
	return FormatText(self)
}

func (self *ExitCooldownRequest) FormatJSON() string {
	return FormatJSON(self)
}
