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

type SingularityPauseRequestList []*SingularityPauseRequest

func (list *SingularityPauseRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityPauseRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityPauseRequestList) FormatJSON() string {
	return FormatJSON(list)
}
