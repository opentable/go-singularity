package dtos

import "io"

type DeleteRequestRequest struct {
	ActionId string
	Message string
}

func (self *DeleteRequestRequest) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *DeleteRequestRequest) FormatText() string {
	return FormatText(self)
}

func (self *DeleteRequestRequest) FormatJSON() string {
	return FormatJSON(self)
}
