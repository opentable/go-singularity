package dtos

import "io"

type SingularityDeleteRequestRequest struct {
	ActionId string
	Message  string
}

func (self *SingularityDeleteRequestRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityDeleteRequestRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityDeleteRequestRequest) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityDeleteRequestRequestList []*SingularityDeleteRequestRequest

func (list *SingularityDeleteRequestRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityDeleteRequestRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDeleteRequestRequestList) FormatJSON() string {
	return FormatJSON(list)
}
