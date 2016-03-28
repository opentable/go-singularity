package dtos

import "io"

type ByteString struct {
	Empty     bool
	ValidUtf8 bool
}

func (self *ByteString) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *ByteString) FormatText() string {
	return FormatText(self)
}

func (self *ByteString) FormatJSON() string {
	return FormatJSON(self)
}
