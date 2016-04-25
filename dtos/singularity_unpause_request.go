package dtos

import "io"

type SingularityUnpauseRequest struct {
	ActionId         string `json:"actionId"`
	Message          string `json:"message"`
	SkipHealthchecks bool   `json:"skipHealthchecks"`
}

func (self *SingularityUnpauseRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityUnpauseRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityUnpauseRequest) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityUnpauseRequestList []*SingularityUnpauseRequest

func (list *SingularityUnpauseRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityUnpauseRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityUnpauseRequestList) FormatJSON() string {
	return FormatJSON(list)
}
