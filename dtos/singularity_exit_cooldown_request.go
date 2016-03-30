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

type SingularityExitCooldownRequestList []*SingularityExitCooldownRequest

func (list SingularityExitCooldownRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list SingularityExitCooldownRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list SingularityExitCooldownRequestList) FormatJSON() string {
	return FormatJSON(list)
}
