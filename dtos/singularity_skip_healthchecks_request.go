package dtos

import "io"

type SingularitySkipHealthchecksRequest struct {
	ActionId         string
	DurationMillis   int64
	Message          string
	SkipHealthchecks bool
}

func (self *SingularitySkipHealthchecksRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularitySkipHealthchecksRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularitySkipHealthchecksRequest) FormatJSON() string {
	return FormatJSON(self)
}

type SingularitySkipHealthchecksRequestList []*SingularitySkipHealthchecksRequest

func (list *SingularitySkipHealthchecksRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularitySkipHealthchecksRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularitySkipHealthchecksRequestList) FormatJSON() string {
	return FormatJSON(list)
}
