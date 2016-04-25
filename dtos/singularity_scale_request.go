package dtos

import "io"

type SingularityScaleRequest struct {
	ActionId         string `json:"actionId"`
	DurationMillis   int64  `json:"durationMillis"`
	Instances        int32  `json:"instances"`
	Message          string `json:"message"`
	SkipHealthchecks bool   `json:"skipHealthchecks"`
}

func (self *SingularityScaleRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityScaleRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityScaleRequest) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityScaleRequestList []*SingularityScaleRequest

func (list *SingularityScaleRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityScaleRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityScaleRequestList) FormatJSON() string {
	return FormatJSON(list)
}
