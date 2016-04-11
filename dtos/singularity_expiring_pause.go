package dtos

import "io"

type SingularityExpiringPause struct {
	ActionId string
	//	ExpiringAPIRequestObject *T
	RequestId   string
	StartMillis int64
	User        string
}

func (self *SingularityExpiringPause) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityExpiringPause) FormatText() string {
	return FormatText(self)
}

func (self *SingularityExpiringPause) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityExpiringPauseList []*SingularityExpiringPause

func (list *SingularityExpiringPauseList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityExpiringPauseList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityExpiringPauseList) FormatJSON() string {
	return FormatJSON(list)
}
