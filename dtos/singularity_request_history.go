package dtos

import "io"

type SingularityRequestHistory struct {
	CreatedAt int64
	//	EventType *RequestHistoryType
	Message string
	Request *SingularityRequest
	User    string
}

func (self *SingularityRequestHistory) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityRequestHistory) FormatText() string {
	return FormatText(self)
}

func (self *SingularityRequestHistory) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityRequestHistoryList []*SingularityRequestHistory

func (list SingularityRequestHistoryList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list SingularityRequestHistoryList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list SingularityRequestHistoryList) FormatJSON() string {
	return FormatJSON(list)
}
