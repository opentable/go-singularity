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
