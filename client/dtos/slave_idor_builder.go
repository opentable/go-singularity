package dtos

import "io"

type SlaveIDOrBuilder struct {
	Value string
	ValueBytes ByteString
}

func (self *SlaveIDOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SlaveIDOrBuilder) FormatText() string {
	return FormatText(self)
}

func (self *SlaveIDOrBuilder) FormatJSON() string {
	return FormatJSON(self)
}
