package dtos

import "io"

type SingularityBounceRequest struct {
	ActionId         string
	DurationMillis   int64
	Incremental      bool
	Message          string
	SkipHealthchecks bool
}

func (self *SingularityBounceRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityBounceRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityBounceRequest) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityBounceRequestList []*SingularityBounceRequest

func (list *SingularityBounceRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityBounceRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityBounceRequestList) FormatJSON() string {
	return FormatJSON(list)
}
