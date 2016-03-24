package dtos

import "io"

type FrameworkIDOrBuilder struct {
	Value string
	ValueBytes ByteString
}

func (self *FrameworkIDOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *FrameworkIDOrBuilder) FormatText() string {
	return FormatText(self)
}

func (self *FrameworkIDOrBuilder) FormatJSON() string {
	return FormatJSON(self)
}
