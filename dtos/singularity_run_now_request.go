package dtos

import "io"

type SingularityRunNowRequest struct {
	CommandLineArgs  StringList `json:"commandLineArgs"`
	Message          string     `json:"message"`
	RunId            string     `json:"runId"`
	SkipHealthchecks bool       `json:"skipHealthchecks"`
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

type SingularityRunNowRequestList []*SingularityRunNowRequest

func (list *SingularityRunNowRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityRunNowRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityRunNowRequestList) FormatJSON() string {
	return FormatJSON(list)
}
