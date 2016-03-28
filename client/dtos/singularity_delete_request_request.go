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
